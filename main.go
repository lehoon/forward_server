package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	route := chi.NewRouter()
	route.Use(middleware.RequestID)
	route.Use(middleware.Logger)
	route.Use(middleware.Recoverer)
	route.Use(middleware.URLFormat)
	route.Use(render.SetContentType(render.ContentTypeJSON))

	route.Post("/api/v1/notify", func(w http.ResponseWriter, r *http.Request) {
		buf, err := io.ReadAll(r.Body)

		if err != nil {
			fmt.Printf("接收到错误的请求,错误信息:%v\n", err)
			return
		}

		fmt.Printf("接收到转发的请求内容:%s\n", string(buf))
		w.Write([]byte("success"))
	})

	fmt.Println("启动内网接收报文服务器测试程序,接收到的请求信息会打印在终端")
	err := http.ListenAndServe("0.0.0.0:9000", route)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
