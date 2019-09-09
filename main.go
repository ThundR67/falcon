package falcon

import (
	"reflect"
)

//NewErrorHandler returnes new error handler
func NewErrorHandler() *ErrorHandler {
	handler := ErrorHandler{}
	handler.Init()
	return &handler
}

//ErrorHandler will error handle
type ErrorHandler struct {
	handlers       map[reflect.Type]func(error, ...interface{})
	defaultHandler func(error, ...interface{})
}

//Init Initializes
func (errorHandler *ErrorHandler) Init() {
	errorHandler.handlers = map[reflect.Type]func(error, ...interface{}){}
}

//AddHandler adds a new error handler
func (errorHandler *ErrorHandler) AddHandler(handleFunc func(error, ...interface{}),
	errTypes ...interface{}) {

	if errTypes == nil {
		errorHandler.defaultHandler = handleFunc
		return
	}
	errType := reflect.TypeOf(errTypes[0])
	errorHandler.handlers[errType] = handleFunc
}

//Check checks err
func (errorHandler ErrorHandler) Check(err error, data ...interface{}) {
	if err == nil {
		return
	}
	errType := reflect.TypeOf(err)
	handler, valid := errorHandler.handlers[errType]
	if !valid {
		handler = errorHandler.defaultHandler
	}
	handler(err, data)
	return
}
