package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	//将 request 中带的 header 写入 response header
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			log.Printf("%s=%s", k, v[0])
			w.Header().Set(k, v[0])
		}
	}
	log.Printf("\n\n\n")

	r.ParseForm()
	if len(r.Form) > 0 {
		for k, v := range r.Form {
			log.Printf("%s=%s", k, v[0])
		}
	}
	log.Printf("\n\n\n")

	//获取环境变量version,并写入response header
	OsVer := os.Getenv("VERSION")
	w.Header().Set("VERSION", OsVer)

	//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println("err:", err)
	}

	if net.ParseIP(ip) != nil {
		fmt.Println("ip===>>%s\n", ip)
		log.Println(ip)
	}
	fmt.Println("http Status Code ===>>%s\n", http.StatusOK)
	log.Println(http.StatusOK)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server Accessy, Success!"))
}

//当访问 localhost/healthz 时，应返回 200
func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200")
}
