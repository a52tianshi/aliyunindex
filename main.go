package main

import (
	"fmt"
	//	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var MainDirPath string
var HtmlPath string

func init() {
	fmt.Println("start!")
	MainFilePath, _ := filepath.Abs(os.Args[0])
	MainDirPath = filepath.Dir(MainFilePath)
	HtmlPath = MainDirPath + "/static/html"
}
func main() {
	fmt.Println("MainDir", MainDirPath)
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/chenquan/", proxychenquan)
	http.HandleFunc("/sunqixiang/", proxysunqixiang)
	http.HandleFunc("/youwei/", proxyyouwei)
	http.HandleFunc("/jilingyan/", proxyjilingyan)
	fmt.Println(http.ListenAndServe("139.196.225.38:80", nil))
}
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)
	bufs, _ := ioutil.ReadFile(HtmlPath + "/index.html")
	bufs2, _ := ioutil.ReadFile(HtmlPath + r.URL.Path)
	if len(bufs2) > 0 {
		w.Write(bufs2)
		return
	}
	w.Write(bufs)
}
func proxychenquan(w http.ResponseWriter, r *http.Request) {
	res, _ := http.Get("http://127.0.0.10" + r.URL.Path[9:]) ///省略去9位
	buf, _ := ioutil.ReadAll(res.Body)
	w.Write(buf)
	res.Body.Close()
}
func proxysunqixiang(w http.ResponseWriter, r *http.Request) {
	res, _ := http.Get("http://127.0.0.11" + r.URL.Path[11:]) ///省略去11位
	buf, _ := ioutil.ReadAll(res.Body)
	w.Write(buf)
	res.Body.Close()
}
func proxyyouwei(w http.ResponseWriter, r *http.Request) {
	res, _ := http.Get("http://127.0.0.12" + r.URL.Path[7:]) ///省略去7位
	buf, _ := ioutil.ReadAll(res.Body)
	w.Write(buf)
	res.Body.Close()
}

func proxyjilingyan(w http.ResponseWriter, r *http.Request) {
	res, _ := http.Get("http://127.0.0.13" + r.URL.Path[10:]) ///省略去10位
	buf, _ := ioutil.ReadAll(res.Body)
	w.Write(buf)
	res.Body.Close()
}
