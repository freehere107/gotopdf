package until

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type fileError struct {
	err error
}

func GetFile(path string) (string, error) {
	fileExt, err := GetFileExt(path)
	if err != nil {
		return "", err
	}
	mkdir()
	createFile := "files/" + getTimestamp() + "." + fileExt
	f, err := os.OpenFile(createFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	stat, err := f.Stat()
	if err != nil {
		panic(err)
		return "", err
	}
	url := path
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Range", "bytes="+strconv.FormatInt(stat.Size(), 10)+"-")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
		return "", err
	}
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		panic(err)
		return "", err
	}
	return createFile, nil
}

func getTimestamp() string {
	t := time.Now().UTC()
	timestamp := t.UnixNano()
	return fmt.Sprintf("%d", timestamp)
}

func GetFileExt(path string) (string, error) {
	pathInfo := strings.Split(path, ".")
	if len(pathInfo) <= 1 {
		return "", errors.New("文件格式错误")
	} else {
		return pathInfo[len(pathInfo)-1], nil
	}
}

func GetPdfPath(path string) string {
	arr := strings.Split(path, ".")
	return strings.Join(arr[0:len(arr)-1], ".") + ".pdf"
}

func GetUpQiniuKey(path string) string {
	u, err := url.Parse(path)
	if err != nil {
		panic(err)
	}
	path = u.Path[1:]
	arr := strings.Split(path, ".")
	return strings.Join(arr[0:len(arr)-1], ".") + ".pdf"
}

func mkdir() {
	path := "files"
	if _, err := os.Stat(path); err == nil {
		return
	} else {
		err := os.MkdirAll(path, 0711)
		if err != nil {
			return
		}
	}
}
