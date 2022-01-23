package file

import (
	"inicio.com/src/messages"
	"path/filepath"
	"archive/zip"
	"net/http"
	"strings"
	"io"
	"os"
)

const tmpFolder = "tmp/"
const unzippedFolder = "output"
const zippedLink = "https://eforexcel.com/wp/wp-content/uploads/2017/07/100-Sales-Records.zip"
// const zippedLink = "https://eforexcel.com/wp/wp-content/uploads/2020/09/2m-Sales-Records.zip"
const zippedExtension = ".zip"

func downloadFile(file_name string) {
	out, err := os.Create(file_name)
	defer out.Close()
	if err != nil { panic(messages.ERROR_CREATING_FILE) }

	resp, err := http.Get(zippedLink)
	defer resp.Body.Close()
	if err != nil { panic(messages.ERROR_DOWNLOADING_FILE) }

	_, err = io.Copy(out, resp.Body)
	if err != nil { panic(messages.ERROR_SAVING_FILE) }
}

func unzipFile(file_name string) {
    dst := "output"
    archive, err := zip.OpenReader(file_name)
    if err != nil { panic(err) }
    defer archive.Close()

    for _, f := range archive.File {
        filePath := filepath.Join(dst, f.Name)

        if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
            panic("invalid file path")
        }

        if f.FileInfo().IsDir() {
            os.MkdirAll(filePath, os.ModePerm)
            continue
        }

        if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil { panic(err) }

        dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
        if err != nil { panic(err) }

        fileInArchive, err := f.Open()
        if err != nil { panic(err) }

        if _, err := io.Copy(dstFile, fileInArchive); err != nil { panic(err) }

        dstFile.Close()
        fileInArchive.Close()
    }
}

func GetFile() {
	tmpFile := tmpFodler + "tmp"

	downloadFile(tmpFile + zippedExtension)
	unzipFile(tmpFile + zippedExtension)
}