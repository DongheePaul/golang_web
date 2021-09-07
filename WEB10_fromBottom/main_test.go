package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPage(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL)

	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("Hello World", string(data))
}

func TestDecoHandler(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	buf := &bytes.Buffer{}
	//버퍼를 만들고 거기에 로그를 담는다.
	log.SetOutput(buf)

	resp, err := http.Get(ts.URL)

	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	//버퍼를 한줄씩 읽어준다.
	r := bufio.NewReader(buf)
	line, _, err := r.ReadLine()
	assert.NoError(err)
	assert.Contains(string(line), "[LOGGER1] Started")

	line, _, err = r.ReadLine()
	assert.NoError(err)
	assert.Contains(string(line), "[LOGGER1] Completed")

}
