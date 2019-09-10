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

	errHandler.AddHandler(func(err error, data ...interface{}) interface{} {
		defaultErrOccured = true
		dataPassed = data != nil
		return true
	})
	errHandler.AddHandler(func(err error, data ...interface{}) interface{} {
		customErrOccured = true
		return true
	}, customErr{})

	assert.True(errHandler.Check(errors.New("test"), "test").(bool))
	assert.True(errHandler.Check(customErr{}).(bool))

	assert.True(defaultErrOccured)
	assert.True(customErrOccured)
	assert.True(dataPassed)
}
