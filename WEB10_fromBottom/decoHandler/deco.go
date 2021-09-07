package decoHandler

import "net/http"

//function type.   DecoratorFunc 타입은 logger() 타입. 즉 main.go 에서 만든 logger()가 데코레이터
type DecoratorFunc func(http.ResponseWriter, *http.Request, http.Handler)

type DecoHandler struct {
	fn DecoratorFunc
	//이 구조체 자체도 http handler 구현 = ServeHTTP
	h http.Handler
}

func (self *DecoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	self.fn(w, r, self.h)
}

//기존 핸들러를 wrapping 한다 (감싼다) 넘어온 http 핸들러를 멤버변수로 가진다.
func NewDecoHandler(h http.Handler, fn DecoratorFunc) http.Handler {
	return &DecoHandler{
		fn: fn,
		h:  h,
	}

}
