package module

import (
	"path/filepath"
	"resource_server/tool"
)

type ResourceList struct {
	fileReader tool.FileReader
}

func (this *ResourceList) GetFilesFromDir(dirName string) ([]FileInfo, error) {
	files, err := this.fileReader.ListDir(dirName)
	if err != nil {
		return nil, err
	}

	ret := make([]FileInfo, len(files))
	for index, file := range files {
		fileName := file.Name()

		// fmt.Println(fmt.Sprintf("是否是文件夹：%t；文件名：%s；文件大小：%d", file.IsDir(), fileName, file.Size()))

		name, ext := this.fileReader.SplitNameAndExt(fileName)
		fileInfo := FileInfo{
			IsDir:      file.IsDir(),
			Name:       name,
			Ext:        ext,
			Fullpath:   filepath.Join(dirName, fileName),
			Size:       file.Size(),
			ModifyTime: file.ModTime().Local().Unix(),
		}
		ret[index] = fileInfo
	}
	return ret, nil
}

type FileInfo struct {
	IsDir      bool   `json:"isDir,omitempty"`
	Name       string `json:"name,omitempty"`
	Ext        string `json:"ext,omitempty"`
	Fullpath   string `json:"fullpath,omitempty"`
	Size       int64  `json:"size,omitempty"`
	ModifyTime int64  `json:"modifyTime,omitempty"`
}

// 参考资料：https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter06/06.2.html
