package interceptor

import (
	"fmt"
	"github.com/google/go-safeweb/safehttp"
)

type Interceptor2 struct {}

func Default2() Interceptor2 {
	return Interceptor2{}
}

func (it Interceptor2) Before(w safehttp.ResponseWriter, r *safehttp.IncomingRequest, _ safehttp.InterceptorConfig) safehttp.Result {
	fmt.Println("2: before")
	return safehttp.NotWritten()
}

func (it Interceptor2) Commit(safehttp.ResponseHeadersWriter, *safehttp.IncomingRequest, safehttp.Response, safehttp.InterceptorConfig) {
	fmt.Println("2: commit")
	return
}

// Match returns false since there are no supported configurations.
func (it Interceptor2) Match(safehttp.InterceptorConfig) bool {
	fmt.Println("2: match")
	return false
}
