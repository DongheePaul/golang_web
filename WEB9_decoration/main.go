package main

import (
	"WEB9_decoration/lzw"
	"fmt"
)

type Receiver interface {
	Operator(string)
}

type Sender interface {
	Operator(string)
}

// 전역변수
var sentData string
var recvData string

//전송 컴포넌트 구조체 선언.
type SendComponent struct{}

//SendComponent의 Operator 함수 호출.
func (self *SendComponent) Operator(data string) {
	sentData = data
}

//데이터 압축할 컴포넌트 - 첫번째 데코레이터
type ZipComponent struct {
	com Sender
	key string
}

func (self *ZipComponent) Operator(data string) {
	zipData, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(zipData))
}

//EncryptComponent 데코레이터 생성.
// type EncryptComponent struct {
// 	key string
// 	com Component
// }

// func (self *EncryptComponent) Operator(data string) {
// 	encryptData, err := cipher.Encrypt([]byte(data), self.key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	self.com.Operator(string(encryptData))
// }

// 복호화 컴포넌트 - 데코레이터
// type DecryptComponent struct {
// 	key string
// 	com Component
// }

// func (self *DecryptComponent) Operator(data string) {
// 	decryptData, err := cipher.Decrypt([]byte(data), self.key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	self.com.Operator(string(decryptData))
// }

// 압축해제 컴포넌트 - 데코레이터
type UnzipComponent struct {
	com Receiver
}

func (self *UnzipComponent) Operator(data string) {
	unzipData, err := lzw.Read([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(unzipData))
}

//recvData로 최종 응축 시키는 단계
type ReadComponent struct{}

func (self *ReadComponent) Operator(data string) {
	recvData = data
}

func main() {
	//EncryptComponent 컴포넌트를 sender 변수에 할당.
	// sender := &EncryptComponent{
	// 	key: "abcde",
	// 	//EncryptComponent의 컴포넌트에 ZipComponent 할당
	// 	com: &ZipComponent{
	// 		//ZipComponent의 컴포넌트에 SendComponent를 할당.
	// 		com: &SendComponent{},
	// 	},
	// }

	//EncryptComponent를 받는 Operator를 실행하면
	//1. 전달된 매개변수를 암호화 후 컴포넌트에 담긴 Operator() 실행. 이 문맥에서는 ZipComponent 실행
	//2. ZimComponent를 받는 Operator()가 실행되면 데이터를 압축함. 그 후 컴포넌트에 담긴 Operator() 실행. 이 문맥에서는 SendComponent 실행.
	//3. SendComponent를 받는 Operator()가 실행되면 전달받은 데이터를 전역변수 sendData에 할당.
	//sender.Operator("Hello World")
	//fmt.Println(sentData)

	//UnzipComponent를 receiver 변수에 할당.
	//receiver := &UnzipComponent{
	//UnzipComponent의 컴포넌트에 DecryptComponent 할당.
	//com: &DecryptComponent{
	//key: "abcde",
	//DecryptComponent의 컴포넌트에 ReadComponent 할당.
	//com: &ReadComponent{},
	//},
	//}

	//UnzipComponent를 받는 Operator를 실행하면
	//1. 전달된 매개변수를 압축 해제 후 컴포넌트에 담긴 Operator() 실행. 이 문맥에서는 DecryptComponent 실행
	//3. DecryptComponent를 받는 Operator()가 실행되면 데이터를 복호화. 그 후 컴포넌트에 담긴 Operator() 실행. 이 문맥에서는 ReadComponent 실행.
	//4. ReadComponent를 받는 Operator()가 실행되면 전달받은 데이터를 전역변수 recvData에 할당.
	// receiver.Operator(sentData)
	// fmt.Println(recvData)

	// fmt.Println("==============================")

	sender := &ZipComponent{
		//EncryptComponent의 컴포넌트에 ZipComponent 할당
		com: &SendComponent{},
	}
	sender.Operator("Hello World")

	fmt.Println(sentData)

	receiver := &UnzipComponent{
		//UnzipComponent의 컴포넌트에 DecryptComponent 할당.
		com: &ReadComponent{},
	}
	receiver.Operator(sentData)
	fmt.Println(recvData)
}
