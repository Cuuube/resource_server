package tool

import (
	"io/ioutil"
	"os"
	"strings"
)

var defaultContentType = "text/plain"
var downloadContentType = "application/octet-stream"
var contentTypeMap = map[string]string{
	// 图片
	"png":  "image/png",
	"jpg":  "application/x-jpg",
	"jpeg": "image/jpeg",
	"tif":  "image/tiff",
	"gif":  "image/gif",
	"ico":  "image/x-icon",
	"pdf":  "application/pdf",
	"svg":  "text/xml",

	// 音视频
	"mp3": "audio/mp3",
	"mp4": "video/mpeg4",
	"mpv": "video/mpg",
	"avi": "video/avi",
	"wmv": "video/x-ms-wmv",
	"flv": "video/x-flv",

	// 文本
	"html": "text/html;charset=utf-8",
	"htm":  "text/html;charset=utf-8",
	"xml":  "application/xml;charset=utf-8",
	"json": "application/json;charset=utf-8",
	"txt":  "text/plain",
}

type FileReader struct {
}

func (fr *FileReader) ListDir(path string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(path)
}

func (fr *FileReader) GetStat(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

func (fr *FileReader) ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func (fr *FileReader) SplitNameAndExt(fileName string) (string, string) {
	var (
		name string
		ext  string
	)
	// 如果是.abc这种格式的文件名，视为文件名为“.abc”，拓展名为“”
	if strings.Index(fileName, ".") == 0 {
		return fileName, ""
	}

	clips := strings.Split(fileName, ".")
	if len(clips) > 1 {
		name = strings.Join(clips[0:len(clips)-1], ".")
		ext = clips[len(clips)-1]
	} else {
		name = fileName
		ext = ""
	}
	return name, ext
}

func (fr *FileReader) GetContentTypeByExt(ext string) string {
	contentType, found := contentTypeMap[ext]
	if found {
		return contentType
	}
	return defaultContentType
}

func Test() {

}
