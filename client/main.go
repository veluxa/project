package main

import (
	"fmt"
	"net/http"

	"yunquk.com/client/files"
)

func main() {

	http.HandleFunc("/file/upload", files.Upload)
	http.HandleFunc("/file/upload/suc", files.UploadSucHandler)
	http.HandleFunc("/file/meta", files.GetFileMeta)
	http.HandleFunc("/file/download", files.Download)
	http.HandleFunc("/file/rename", files.FileRename)
	http.HandleFunc("/file/delete", files.FileDel)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("启动程序出错", err.Error())
	}

}
