package main

import (
	"io"
	_ "io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {

	// b := bytes.NewBuffer(make([]byte, 256))
	// r.Write(b)
	// proxyData := bytes.Trim(b.Bytes(), "\x00")

	client := &http.Client{}
	req := new(http.Request)
	*req = *r
	// req, _ := http.ReadRequest(bufio.NewReader(bytes.NewBuffer(proxyData)))

	req.URL, _ = url.Parse(`http://10.95.154.63:8080` + req.URL.Path)
	req.RequestURI = ""
	rep, _ := client.Do(req)

	for key, value := range rep.Header {
		for _, v := range value {
			w.Header().Add(key, v)
		}
	}

	w.WriteHeader(rep.StatusCode)
	io.Copy(w, rep.Body)
	rep.Body.Close()

}

func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
