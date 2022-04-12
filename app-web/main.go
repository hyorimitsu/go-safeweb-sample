package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/go-safeweb/safehttp"
	"github.com/google/go-safeweb/safehttp/plugins/csp"
	"github.com/google/go-safeweb/safehttp/plugins/htmlinject"
	"github.com/google/go-safeweb/safesql"
	"github.com/google/safehtml/template"

	"github.com/hyorimitsu/hello-go-safeweb/app-web/interceptor"
)

func main() {
	safeHttp()
	//safeSql()
}

func safeHttp() {
	isSafe := os.Getenv("RUN_SAFE")

	// DOM XSS があるページの例
	if isSafe == "true" {
		// safehttp を利用した場合
		mc := safehttp.NewServeMuxConfig(nil)
		mc.Intercept(csp.Default(""), interceptor.Default1(), interceptor.Default2(), interceptor.Default3())
		mux := mc.Mux()

		safeTmplSrc, _ := template.TrustedSourceFromConstantDir("", template.TrustedSource{}, "index.html")
		safeTmpl := template.Must(htmlinject.LoadFiles(nil, htmlinject.LoadConfig{}, safeTmplSrc))
		mux.Handle("/", safehttp.MethodGet, safehttp.HandlerFunc(handleTemplateSafe(safeTmpl)))

		fmt.Println("starting safe server")
		log.Fatal(http.ListenAndServe(":80", mux))

	} else {
		// safehttp を利用しなかった場合
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			tmpl, err := template.ParseFiles("index.html")
			if err != nil {
				panic(err)
			}
			err = tmpl.Execute(w, nil)
			if err != nil {
				panic(err)
			}
		})

		fmt.Println("starting unsafe server")
		log.Fatal(http.ListenAndServe(":80", nil))
	}
}

// database/sql の wrapper
var db *safesql.DB

func safeSql() {
	// safesql を利用したクエリ作成の例

	// OKパターン
	trustedSqlStr1 := safesql.New("SELECT * FROM users WHERE id = ?")
	db.Query(trustedSqlStr1, "A01")

	// NGパターン（コンパイルエラー）
	//userID := "A01"
	//trustedSqlStr2 := safesql.New("SELECT * FROM users WHERE id = " + userID)
	//db.Query(trustedSqlStr2)
}

func handleTemplateSafe(tmpl safehttp.Template) func(w safehttp.ResponseWriter, req *safehttp.IncomingRequest) safehttp.Result {
	return func(w safehttp.ResponseWriter, req *safehttp.IncomingRequest) safehttp.Result {
		return safehttp.ExecuteTemplate(w, tmpl, nil)
	}
}
