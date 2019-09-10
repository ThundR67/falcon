[![Go Report Card](https://goreportcard.com/badge/github.com/SonicRoshan/falcon)](https://goreportcard.com/report/github.com/SonicRoshan/falcon) [![GoDoc](https://godoc.org/github.com/SonicRoshan/falcon?status.svg)](https://godoc.org/github.com/SonicRoshan/falcon) [![GoCover](https://gocover.io/_badge/github.com/SonicRoshan/falcon)](https://gocover.io/github.com/SonicRoshan/falcon)
# Falcon
Minimal And Efficient Error Handling In Go


## Simple Example
```go
errHandler := falcon.NewErrorHandler()

errHandler.AddHandler(func(err error, data ...interface{}) interface{} {
    fmt.Println("Generic Error Occured")
    return SomeData
})

errHandler.AddHandler(func(err error, data ...interface{}) interface{} {
    fmt.Println("Error Of Type CustomErrorType Occured")
    return SomeData
}, CustomErrorType{})

SomeData := errHandler.Check(err, SomeRandomData1, SomeRandomData2)
//Data passed after the err will be passed to handlers
```


