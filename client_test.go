package main

import (
	"io/ioutil"
	"testing"
)

func TestHandle(t *testing.T) {
	mockHandler := &MockHandler{Messages:[][]byte{}}
	client := &Client{Handlers:[]PacketHandler{mockHandler}}

	contents, err := ioutil.ReadFile("test/pcars_udp_0.bin")
	if err != nil {
		t.Error(err)
	}

	client.Handle(contents)

	if len(mockHandler.Messages) != 1 {
		t.Errorf("expected %s != actual %s", len(mockHandler.Messages), 1)
	}
}

type MockHandler struct {
	Messages [][]byte
}

func (mockHandler *MockHandler) Handle(msg []byte) {
	mockHandler.Messages = append(mockHandler.Messages, msg)
}

