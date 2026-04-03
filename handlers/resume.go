package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	cachedData []byte
	cachedAt   time.Time
	mu         sync.RWMutex
	dataFile   = "data/resume.json"
	cacheTTL   = 5 * time.Minute
)

// GetResume 读取 resume.json 并返回，带简单内存缓存
func GetResume(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	if cachedData != nil && time.Since(cachedAt) < cacheTTL {
		data := cachedData
		mu.RUnlock()
		writeJSON(w, data)
		return
	}
	mu.RUnlock()

	// 重新读取文件
	data, err := os.ReadFile(dataFile)
	if err != nil {
		log.Printf("读取 resume.json 失败: %v", err)
		http.Error(w, "数据加载失败", http.StatusInternalServerError)
		return
	}

	// 校验 JSON 合法性
	if !json.Valid(data) {
		log.Println("resume.json 格式不合法")
		http.Error(w, "数据格式错误", http.StatusInternalServerError)
		return
	}

	mu.Lock()
	cachedData = data
	cachedAt = time.Now()
	mu.Unlock()

	writeJSON(w, data)
}

func writeJSON(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
