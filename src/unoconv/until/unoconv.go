package until

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type request struct {
	filename string
	filetype string
	w        io.Writer
	errChan  chan error
	notify   string
	path     string
}

type Unoconv struct {
	requestChan chan request
}

func InitUnoconv() *Unoconv {
	uno := new(Unoconv)
	uno.requestChan = make(chan request)
	go func(uno *Unoconv) {
		for {
			select {
			case data := <-uno.requestChan:
				cmd := exec.Command("unoconv", "-f", data.filetype, data.filename)
				cmd.Stdout = data.w
				err := cmd.Run()
				os.Remove(data.filename) //删除本地下载的远程文件
				UpQiniu(GetPdfPath(data.filename), data.notify, GetUpQiniuKey(data.path))
				if err != nil {
					ErrorNotify(data.notify)
					fmt.Print("unoconv error", err)
					data.errChan <- err
				} else {
					data.errChan <- nil
				}
			}
		}
	}(uno)
	return uno
}

func (u *Unoconv) Convert(filetype string, w io.Writer, notify, path string) error {
	localPath, fileErr := GetFile(path)
	if fileErr != nil {
		ErrorNotify(notify)
		return fileErr
	} else {
		err := make(chan error)
		req := request{
			localPath,
			filetype,
			w,
			err,
			notify,
			path,
		}

		u.requestChan <- req
		return <-err
	}

}
