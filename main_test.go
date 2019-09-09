package falcon

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var defaultErrOccured bool
var customErrOccured bool
var dataPassed bool

type customErr struct{}

func (ce customErr) Error() string {
	return "custom error"
}

func TestErrHandler(t *testing.T) {
	errHandler := NewErrorHandler()
	assert := assert.New(t)

	assert.NotNil(errHandler)
	assert.IsType(&ErrorHandler{}, errHandler)

	errHandler.AddHandler(func(err error, data ...interface{}) {
		defaultErrOccured = true
		dataPassed = data != nil
	})
	errHandler.AddHandler(func(err error, data ...interface{}) {
		customErrOccured = true
	}, customErr{})

	errHandler.Check(errors.New("test"), "test")
	errHandler.Check(customErr{})

	assert.True(defaultErrOccured)
	assert.True(customErrOccured)
	assert.True(dataPassed)
}
