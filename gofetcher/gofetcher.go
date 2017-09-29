package gofetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/boltdb/bolt"
)

var (
	href_regexps = []*regexp.Regexp{
		regexp.MustCompile(`href\s*=\s*"([\S]+)"`),
		regexp.MustCompile(`src\s*=\s*"([\S]+)"`),
		regexp.MustCompile(`url\s*\("?([0-9A-Za-z_\-\./]+)"?\)`),
	}
)

//provider base function to implements curl
type GoFetcher struct {
	*http.Client
	baseDir string
	destUrl string
	host    string
	dirPath string
	ignores []string
	db      *bolt.DB
}

func NewFetcher(urlStr string) *GoFetcher {
	dir, _ := os.Getwd()
	urlStr = strings.TrimRight(urlStr, "/")
	host := getHost(urlStr)
	if host == "" {
		panic("not leaggle url")
	}
	tryMkdir(host)
	db, err := bolt.Open(fmt.Sprintf("%s/my%d.db", host, time.Now().Unix()), 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return &GoFetcher{
		Client:  http.DefaultClient,
		destUrl: urlStr,
		dirPath: path.Join(dir, host),
		host:    host,
		db:      db,
		ignores: []string{},
	}
}
func (spider *GoFetcher) AddIgnore(suffix string) {
	spider.ignores = append(spider.ignores, suffix)
}

func (spider *GoFetcher) Execute() (err error) {
	defer spider.close()
	spider.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("gos_" + spider.host))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	return spider.fetch(spider.destUrl)
}

func (spider *GoFetcher) close() (err error) {
	return spider.db.Close()
}

func (spider *GoFetcher) fetch(urlStr string) (err error) {
	if strings.Index(urlStr, `http`) != 0 {
		urlStr = "http://" + urlStr
	}
	urlStr = strings.TrimRight(urlStr, "/")
	//read write
	var ok bool
	spider.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("gos_" + spider.host))
		bs := b.Get([]byte(urlStr))
		ok = string(bs) == "1"
		return nil
	})
	if !ok {
		spider.db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("gos_" + spider.host))
			return b.Put([]byte(urlStr), []byte("1"))
		})
	} else {
		return
	}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return err
	}
	resp, err := spider.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(resp.Body)
	spider.saveFile(urlStr, bs)
	//recusive
	typ := resp.Header.Get("Content-Type")
	if "text/html" == typ || "text/css" == typ {
		for _, r := range href_regexps {
			matches := r.FindAllStringSubmatch(string(bs), -1)
			for _, m := range matches {
				if len(m) > 1 {
					realUrl, ok := spider.filterURL(urlStr, m[1])
					if ok {
						err = spider.fetch(realUrl)
						if err != nil {
							log.Println("error fetch->" + err.Error())
						}
					}
				}
			}
		}
	}

	return
}

func (spider *GoFetcher) saveFile(urlStr string, content []byte) {
	name, dirPath := getNameAndPath(urlStr, spider.host)
	var ps = path.Join(spider.dirPath, dirPath)
	tryMkdir(ps)
	fname := path.Join(ps, name)
	if _, err := os.Stat(fname); err == nil {
		return
	}
	fs, err := os.OpenFile(fname, os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		log.Print("error create file" + ps + "=>" + name + err.Error())
		return
	}
	fs.Write(content)
}

func (spider *GoFetcher) filterURL(parentURL, href string) (res string, ok bool) {
	res = href
	href = strings.TrimSpace(href)
	if strings.Contains(href, "javascript") {
		return
	}
	for _, suffix := range spider.ignores {
		if strings.HasSuffix(res, suffix) {
			ok = false
			return
		}
	}
	parentURI, _ := url.ParseRequestURI(parentURL)
	res = href
	if strings.Index(href, "http") != 0 {
		if strings.Index(href, "//") == 0 {
			res = parentURI.Scheme + ":" + href
		} else if strings.Index(href, "/") == 0 {
			res = parentURI.Scheme + "://" + parentURI.Host + href
		} else {
			res = parentURI.Scheme + "://" + parentURI.Host + path.Join(path.Dir(parentURI.RequestURI()), href)
			//去掉其他跨域的
			if !strings.Contains(res, parentURI.Host) {
				return
			}
		}
	} else {
		res = href
		if !strings.Contains(res, parentURI.Host) {
			return
		}
	}
	if idx := strings.Index(res, "#"); idx != -1 {
		res = res[:idx]
	}

	ok = true
	return
}

//mkdir -p
func tryMkdir(dir string) {
	info, _ := os.Stat(dir)
	if info == nil || !info.IsDir() {
		if ind := strings.LastIndex(dir, `/`); ind != -1 {
			tryMkdir(dir[:ind])
		}
		os.Mkdir(dir, 0777)
	}
}

func getNameAndPath(urlStr string, host string) (name string, dirPath string) {
	idx := strings.LastIndex(urlStr, ".")
	idx2 := strings.LastIndex(urlStr, "/")

	if idx2 > idx {
		urlStr = urlStr + "/"
	}
	_url, err := url.Parse(urlStr)
	if err != nil {
		println(err.Error(), urlStr)
		return
	}
	dirPath = path.Dir(_url.RequestURI())
	name = strings.Replace(_url.RequestURI(), dirPath, "", -1)
	if name == "/" || name == "" {
		name = "index.html"
	}
	return
}

func getHost(urlStr string) string {
	uri, err := url.Parse(urlStr)
	if err != nil {
		return ""
	}
	return uri.Host
}

//path join for url
func pathJoin(host, urlPath, additionPath string) (res string) {
	if additionPath[0] == '/' {
		res = host + additionPath
		return
	}
	res = pathDir(urlPath) + "/" + additionPath
	return
}

func pathDir(urlPath string) (res string) {
	urlPath = strings.TrimRight(urlPath, "/")
	index := strings.LastIndex(urlPath, "/")
	return urlPath[:index]
}
