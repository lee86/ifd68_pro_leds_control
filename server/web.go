package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// startServer 开启控制台网页
func (ifd68 *Ifd68Pro) startServer() {
	router := mux.NewRouter()
	router.HandleFunc("/", ifd68.StartServer)
	router.HandleFunc("/color", ifd68.ColorHandle)
	router.PathPrefix("/exp").HandlerFunc(ifd68.StaticServer)
	srv := &http.Server{
		Handler:      router,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("http://127.0.0.1:8000")
	log.Fatal(srv.ListenAndServe())
}

// ColorHandle api接口
func (ifd68 *Ifd68Pro) ColorHandle(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Read failed:", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("断链失败")
		}
	}(r.Body)
	//fmt.Println(string(b))
	err = json.Unmarshal(b, &ifd68.ColorWeb)
	if err != nil {
		return
	}
	ifd68.Color.ColorType = ifd68.ColorWeb.ColorType
	ifd68.SetColor()
}

func (ifd68 *Ifd68Pro) StartServer(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./"+r.URL.Path)
}

func (ifd68 *Ifd68Pro) StaticServer(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./"+r.URL.Path)
}
