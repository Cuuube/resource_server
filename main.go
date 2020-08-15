package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	http.HandleFunc("/upload", next(uploadHandler))

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

func uploadHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.Write([]byte("Method Not Allow!"))
		return
	}
	path := req.URL.Query().Get("path")
	if path == "" {
		w.Write([]byte("no path"))
		return
	}

	// 解析form-data
	req.ParseMultipartForm(32 << 20)

	// 获取formdata的文本资料
	println(req.FormValue("name"))

	file, header, err := req.FormFile("file")
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	defer file.Close()

	// 打开新文件
	f, err := os.OpenFile(path+"/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
	}
	defer f.Close()
	// 将收到的文件穿进去
	io.Copy(f, file)
	w.Write([]byte("success!"))
}

// 浏览器上测试命令：fetch("http://127.0.0.1:8081/ping").then(res => res.text()).then(console.log)
