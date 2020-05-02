package module

import (
	// "os"
	"resource_server/tool"
	// "syscall"
)

type HomeResource struct {
	fileReader tool.FileReader
}

// func (this *HomeResource) GetAllMountedDisks() {
// 	// mounts, _ := gofstab.ParseSystem()
// 	var stat syscall.Statfs_t
// 	wd, _ := os.Getwd()
// 	syscall.Statfs(wd, &stat)
// 	fmt.Println(stat.Bavail * uint64(stat.Bsize))
// }

// func (this *HomeResource) getDiskInfo(path string) string {
// 	fs := syscall.Statfs_t{}
// 	err := syscall.Statfs(path, &fs)
// 	if err != nil {
// 		return ""
// 	}
// 	all := fs.Blocks * uint64(fs.Bsize)
// 	free := fs.Bfree * uint64(fs.Bsize)
// 	used := all - free
// 	return fmt.Sprintf("total: %d, used: %d, free: %d", all, used, free)
// }
