package interceptor

import (
	"fmt"

	"github.com/google/go-safeweb/safehttp"
)

type Interceptor3 struct {}

func Default3() Interceptor3 {
	return Interceptor3{}
}

func (it Interceptor3) Before(w safehttp.ResponseWriter, r *safehttp.IncomingRequest, _ safehttp.InterceptorConfig) safehttp.Result {
	fmt.Println("3: before")
	return safehttp.NotWritten()
}

func (it Interceptor3) Commit(safehttp.ResponseHeadersWriter, *safehttp.IncomingRequest, safehttp.Response, safehttp.InterceptorConfig) {
	fmt.Println("3: commit")
	return
}

// Match returns false since there are no supported configurations.
func (it Interceptor3) Match(safehttp.InterceptorConfig) bool {
	fmt.Println("3: match")
	return false
}
