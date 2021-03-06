package logger

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/gin-gonic/gin"
)

type LoggerTestSuite struct {
	suite.Suite
	Logger *Logger
	Context *gin.Context
}

func (suite *LoggerTestSuite) SetupSuite() {
	logger, err := NewLogger(false, false)
	suite.Nil(err, "No err creating Logger, json=false, debug=false")

	logger, err = NewLogger(false, true)
	suite.Nil(err, "No err creating Logger, json=false, debug=true")

	logger, err = NewLogger(true, false)
	suite.Nil(err, "No err creating Logger, json=true, debug=false")

	logger, err = NewLogger(true, true)
	suite.Nil(err, "No err creating Logger, json=true, debug=true")
	suite.Logger = logger

	context := &gin.Context{}
	context.Set("RequestCount", "1")
	context.Set("RequestId", "xyz")
	suite.Context = context
}

func (suite *LoggerTestSuite) TestLevelcMethods() {
	suite.Logger.Debugc(suite.Context, "Debugc test", "x", "y")
	suite.Logger.Infoc(suite.Context, "Infoc test", "x", "y")
	suite.Logger.Warnc(suite.Context, "Warnc test", "x", "y")
	suite.Logger.Errorc(suite.Context, "Errorc test", "x", "y")
}

func (suite *LoggerTestSuite) TestContextLoggingFn() {
	log := suite.Logger.ContextLoggingFn(suite.Context)
	log(DebugLevel, "ContextLoggingFn debug test", "x", "y")
	log(InfoLevel, "ContextLoggingFn info test", "x", "y")
	log(WarnLevel, "ContextLoggingFn warn test", "x", "y")
	log(ErrorLevel, "ContextLoggingFn error test", "x", "y")
}

func TestLoggerTestSuite(t *testing.T) {
	suite.Run(t, new(LoggerTestSuite))
}
