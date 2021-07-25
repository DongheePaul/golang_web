package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadTest(t *testing.T) {
	assert := assert.New(t)
	path := "C:/Users/admin/hi.png"
	file, _ := os.Open(path)
	//defer 키워드 : 해당 블록의 코드들이 다 끝나면 실행됨.
	defer file.Close()

	//NewWriter에 io.writer로 넣어주기 위한 변수.
	buf := &bytes.Buffer{}
	//웹에서 파일을 전송할 때 MIME 포맷을 사용하는데, 이것을 위해 multipart.NewWriter() 사용하고, 이의 리턴값(인스턴스)를 담는다.
	writer := multipart.NewWriter(buf)
	//File을 만든다. filepath.Base를 하면 filename만 잘라줌.  CreateFormFile의 리턴값은 io.writer, error
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))
	//error 있는지 확인.
	assert.NoError(err)
	//file의 데이터를 읽어 multi에 넣어준다.
	io.Copy(multi, file)
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	req.Header.Set("Content-type", writer.FormDataContentType())
	//파일 업로드
	uploadsHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

	uploadsFilePath := "./uploads/" + filepath.Base(path)
	//파일의 정보를 읽어온다.
	_, err = os.Stat(uploadsFilePath)
	assert.NoError(err)
	//NoError를 통과했으면 파일이 있다는 의미이므로, 원본 파일과 업로드한 파일을 읽어온다.
	uploadFile, _ := os.Open(uploadsFilePath)
	originFile, _ := os.Open(path)
	defer uploadFile.Close()
	defer originFile.Close()

	uploadData := []byte{}
	originData := []byte{}
	uploadFile.Read(uploadData)
	originFile.Read(originData)

	assert.Equal(originData, uploadData)

}
