package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Alextt666/resume-api/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	// 健康检查
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	// 简历数据接口
	mux.HandleFunc("/api/resume", handlers.GetResume)

	log.Printf("resume-api 启动，监听 :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
