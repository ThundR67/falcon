# Falcon
Minimal And Efficient Error Handling In Go


## Simple Example
```go
errHandler := falcon.NewErrorHandler()

errHandler.AddHandler(func(err error, data ...interface{}) {
    fmt.Println("Generic Error Occured")
})

errHandler.AddHandler(func(err error, data ...interface{}) {
    fmt.Println("Error Of Type CustomErrorType Occured")
}, CustomErrorType{})

errHandler.Check(err, SomeRandomData1, SomeRandomData2)
//Data passed after the err will be passed to handlers
```

