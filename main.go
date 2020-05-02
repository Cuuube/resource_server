package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"resource_server/module"
	"resource_server/tool"
)

var (
	nearAPI            string
	fileReader         tool.FileReader
	homeResourceModule module.HomeResource
	resourceListModule module.ResourceList
)

func next(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		// 配置跨域
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
		w.Header().Add("Content-Type", "application/json;charset=utf-8")

		nearAPI = req.RequestURI

		handler(w, req)
	}
}

func main() {
	port := "8081"

	http.HandleFunc("/", next(indexHander))
	http.HandleFunc("/ping", next(pingHandler))
	http.HandleFunc("/resources", next(resourcesHandler))

	// go func() {
	// 	ticker := time.NewTicker(time.Second * 30)
	// 	for range ticker.C {
	// 		fmt.Println("最近请求的api是：", nearAPI)
	// 	}
	// }()

	fmt.Println("Server running at port", port)
	http.ListenAndServe(":"+port, nil)
}

func indexHander(w http.ResponseWriter, _ *http.Request) {
	var content = "It works!"
	// homeResourceModule.GetAllMountedDisks()
	w.Write([]byte(content))
}

func pingHandler(w http.ResponseWriter, req *http.Request) {
	var content = "pong"
	w.Write([]byte(content))
}

func resourcesHandler(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Query().Get("path")
	// isDir := req.URL.Query().Get("isDir")

	if path == "" {
		w.Write([]byte("no path"))
		return
	}

	stat, err := fileReader.GetStat(path)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	// 文件夹的处理
	if stat.IsDir() {
		lists, err := resourceListModule.GetFilesFromDir(path)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		ret, _ := json.Marshal(lists)
		w.Write(ret)
		return
	}

	// 文件的处理
	_, ext := fileReader.SplitNameAndExt(stat.Name())
	w.Header().Add("Content-Type", fileReader.GetContentTypeByExt(ext))
	w.Header().Add("Content-Length", fmt.Sprintf("%d", stat.Size()))
	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", stat.Name()))
	fileBytes, err := fileReader.ReadFile(path)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(fileBytes)
}

// 浏览器上测试命令：fetch("http://127.0.0.1:8081/ping").then(res => res.text()).then(console.log)
