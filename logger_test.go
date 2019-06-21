package awsutil

import (
	"github.com/aphistic/sweet"
	. "github.com/efritz/go-mockgen/matchers"
	"github.com/go-nacelle/nacelle/mocks"
	. "github.com/onsi/gomega"
)

type LoggerSuite struct{}

func (s *LoggerSuite) TestLogAdapter(t sweet.T) {
	logger := mocks.NewMockLogger()
	adapter := NewAWSLogAdapter(logger)
	adapter.Log("Hello")
	Expect(logger.DebugFunc).To(BeCalledWith("Hello"))
}

func (s *LoggerSuite) TestLogAdapterWithArgs(t sweet.T) {
	logger := mocks.NewMockLogger()
	adapter := NewAWSLogAdapter(logger)
	adapter.Log("Hello, %s and %s", "bar", "baz")
	Expect(logger.DebugFunc).To(BeCalledWith("Hello, %s and %s", "bar", "baz"))
}

func (s *LoggerSuite) TestLogAdapterNonStringFormat(t sweet.T) {
	logger := mocks.NewMockLogger()
	adapter := NewAWSLogAdapter(logger)
	adapter.Log([]string{"this", "is", "dangerous"})
	Expect(logger.DebugFunc).To(BeCalledWith("[this is dangerous]"))
}

func (s *LoggerSuite) TestLogAdapterNoArgs(t sweet.T) {
	logger := mocks.NewMockLogger()
	adapter := NewAWSLogAdapter(logger)
	adapter.Log()
	Expect(logger.DebugFunc).NotTo(BeCalled())
}
