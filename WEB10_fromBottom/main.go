package main

import (
	"WEB10_fromBottom/decoHandler"
	"WEB10_fromBottom/myapp"
	"log"
	"net/http"
	"time"
)

//데코레이터 함수
func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER1] Started")
	//기존 mux의 handler
	h.ServeHTTP(w, r)
	log.Println("[LOGGER1] Completed", time.Since(start).Microseconds())
}

//데코레이터 추가
func logger2(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER2] Started")
	//h = 기존 mux의 handler. NewDecoHandler를 통해 기존 mux를 DecoHandler로 감싸고, DecoHandler의 메소드 ServeHTTP 수행 ->
	h.ServeHTTP(w, r)
	log.Println("[LOGGER2] Completed", time.Since(start).Microseconds())
}
func NewHandler() http.Handler {
	mux := myapp.NewHandler()
	//데코레이터로 기존 핸들러를 감싼다(wrapping)   logger() 는 데코레이션 함수.
	h := decoHandler.NewDecoHandler(mux, logger)
	h = decoHandler.NewDecoHandler(h, logger2)
	return h
}
func main() {
	mux := NewHandler()
	http.ListenAndServe(":3000", mux)
}
