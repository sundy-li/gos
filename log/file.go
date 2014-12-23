package log

import (
	"log"
	"os"
	"time"
)

type FileAppender struct {
	BaseAppender
	AppenderConfig
	file    *os.File
	logDate string
}

func (this *FileAppender) Output(msg string, level int) {
	if ok := this.levels[level]; !ok {
		return
	}
	this.logger.Println(msg)
}
func (this *FileAppender) Close() {
	this.file.Close()
}

func newFileAppender(conf AppenderConfig) (appender *FileAppender) {
	appender = &FileAppender{
		AppenderConfig: conf,
	}
	var err error
	var filePath = conf.FilePath
	if conf.IsDaliy {
		appender.logDate = time.Now().Format("20060102")
		filePath = conf.FilePath + "_" + appender.logDate
	}
	appender.file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		log.Fatalf("cannot create fileappender %s casue :%s", conf.FilePath, err.Error())
		return
	}
	appender.logger = log.New(appender.file, "", log.Ldate|log.Ltime)
	appender.SetLevel(conf.Levels)
	if conf.IsDaliy {
		go appender.logDaily()
	}
	return
}

func (appender *FileAppender) logDaily() {
	var t = time.Now()
	h, m, s := t.Hour(), t.Minute(), t.Second()
	time.AfterFunc(time.Duration(24*3600-h*3600-m*60-s*60), func() {
		ticker := time.NewTicker(time.Second * 5)
		for {
			select {
			case <-ticker.C:
				//wheter is another day
				var date = time.Now().Format("20060102")
				if date != appender.logDate {
					newAppender := newFileAppender(appender.AppenderConfig)
					appender.Close()
					*appender = *newAppender
					log.Printf("New day is coming, starting logging to %s_%s log", appender.AppenderConfig.FilePath, date)
					return
				}
			}
		}
	})
}
