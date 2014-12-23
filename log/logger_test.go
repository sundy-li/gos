package log

import (
	"testing"
)

func TestConsoleLogger(t *testing.T) {
	var log = NewDefaultLogger()
	log.Info("info msg")
	log.Error("error msg")
	log.Debug("debug msg ")
	log.Trace("trace msg fsdajfjdsajkfkjdsafsjkfldsklfjdsafkdsahjflsajhfhljdshjafjdsahfdshafldsahfjdskafhdjsajfkdshafhdsjalk")

	log.Errorf("fuck u %d %s", 33, "fds")
}

func TestRootLogger(t *testing.T) {
	Info("info msg")
	Error("error msg")
	Debug("debug msg")
	Errorf("fsdafsa %s fdsa", "i am error")
}
