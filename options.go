package whitebit

import (
	"sync"
	"time"
)

type Endpoint interface {
	Url() string
	IsAuthed() bool
}

type NoAuth struct{}

func (endpoint NoAuth) IsAuthed() bool {
	return false
}

func (endpoint NoAuth) Url() string {
	return ""
}

type AuthParams struct {
	Request     string `json:"request"`
	Nonce       int64  `json:"nonce"`
	NonceWindow bool   `json:"nonceWindow"`
}

var lastNonce int64
var nonceMutex sync.Mutex
var TimeOffset int64

func NewAuthParams(url string) AuthParams {
	nonceMutex.Lock()
	defer nonceMutex.Unlock()

	nonce := time.Now().UnixMilli() - TimeOffset
	if nonce <= lastNonce {
		nonce = lastNonce + 1
	}
	lastNonce = nonce

	return AuthParams{Nonce: nonce, NonceWindow: true, Request: url}
}

func (params AuthParams) IsAuthed() bool {
	return true
}

func (params AuthParams) Url() string {
	return params.Request
}
