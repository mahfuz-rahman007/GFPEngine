package scanner

import (
	"os"
	"path/filepath"
	"time"
)

type FileInfo struct{
	Path string
	Size int64
	Ext string
	ModTime time.Time
}

func Scan(root string, recursive bool) ([]FileInfo, error) {
	if recursive { 
		return ScanRecursive(root)
	 } else {
		return ScanFirstChild(root)
	 }
}

func ScanRecursive(root string) ([]FileInfo, error) {
	var files []FileInfo
	err := filepath.WalkDir(root, func (path string, dir os.DirEntry, err error) error  {
		if err != nil { return err }
		if dir.IsDir() { return nil }

		info,err:= dir.Info()

		if err!=nil {return err}

		files = append(files, FileInfo{
			Path: path,
			Size: info.Size(),
			Ext: filepath.Ext(path),
			ModTime: info.ModTime(),
		})
		return nil
	})

	return files, err
}

func ScanFirstChild(root string) ([]FileInfo, error) {
	var files []FileInfo

	allFiles, err := os.ReadDir(root)

	if err != nil {return nil, err}

	for _,file := range allFiles {

		info, err := file.Info()

		if err != nil { return files, err }
		if info.IsDir() {continue}

		fullPath := filepath.Join(root, file.Name())

		files = append(files, FileInfo{
			Path: fullPath,
			Size: info.Size(),
			Ext: filepath.Ext(fullPath),
			ModTime: info.ModTime(),
		})
	}

	return files, err
}