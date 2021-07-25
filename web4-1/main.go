package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadsHandler(w http.ResponseWriter, r *http.Request) {
	//r.FormFile이 inputFormFile 형태로 날라온 값을 읽는다는 의미. 리턴값 : multipart.File, multipart.FileHeader, error
	uploadFile, header, err := r.FormFile("upload_file")
	//에러 처리.
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	dirname := "./uploads"
	//폴더 생성 후 777 권한 준다.
	os.MkdirAll(dirname, 0777)
	//파일 경로
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	//파일을 만든다.
	file, err := os.Create(filepath)
	//file의 handle을 사용해서 파일을 생성하는데, 이 Handle은 os 자원이므로 반납해야함.
	defer file.Close()
	//파일 생성시 에러 발생했다면 처리
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	//업로드 한 파일의 내용을 file에 넣어준다.
	io.Copy(file, uploadFile)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filepath)
}

func main() {
	http.HandleFunc("/uploads", uploadsHandler)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":3000", nil)

}
