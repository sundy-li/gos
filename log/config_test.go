package log

import (
	"encoding/json"
	"testing"
	"time"
)

func TestConfig(t *testing.T) {
	var configStr = `{
        "appenders" :{
                "c" :{
                        "type" : "console",
                        "levels" : ["debug", "warn" , "error"]
                    },
                "f" :{
                     "type" : "file",
                     "filePath" : "aa.log",
                     "levels" : ["debug", "warn" , "error"],
                     "isDaily" : true
                }
            },
        "loggers" : {
            "root" :{
                "appenders" :["c" , "f"],
                "showFileLine" : true
            },
            "mine": {
                "appenders" :["c" , "c" , "c"],
                "showFileLine" : true
            }
        }
    }`

	var conf = &Config{}
	err := json.Unmarshal([]byte(configStr), conf)
	if err != nil {
		t.Error(err.Error())
		return
	}
	LoadConfig(conf)

	Debugf("I am debuging")
	Trace("I am not showing")
	var l = Get("mine")
	l.Errorf("error occurs")

	tc := time.NewTicker(3 * time.Second).C
	for {
		select {
		case <-tc:
			Debugf("3 seconds repeat  myself")
		}
	}
}
