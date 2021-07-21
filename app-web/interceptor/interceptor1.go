package interceptor

import (
	"fmt"
	"github.com/google/go-safeweb/safehttp"
)

type Interceptor1 struct {}

func Default1() Interceptor1 {
	return Interceptor1{}
}

func (it Interceptor1) Before(w safehttp.ResponseWriter, r *safehttp.IncomingRequest, _ safehttp.InterceptorConfig) safehttp.Result {
	fmt.Println("1: before")
	return safehttp.NotWritten()
	//return w.WriteError(NewError(400))
}

func (it Interceptor1) Commit(safehttp.ResponseHeadersWriter, *safehttp.IncomingRequest, safehttp.Response, safehttp.InterceptorConfig) {
	fmt.Println("1: commit")
	return
}

// Match returns false since there are no supported configurations.
func (it Interceptor1) Match(safehttp.InterceptorConfig) bool {
	fmt.Println("1: match")
	return false
}
