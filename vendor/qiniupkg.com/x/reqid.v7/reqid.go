package reqid

import (
	"context"
	"encoding/base64"
	"encoding/binary"
	"net/http"
	"time"
)

// --------------------------------------------------------------------

var pid = uint32(time.Now().UnixNano() % 4294967291)

func genReqID() string {
	var b [12]byte
	binary.LittleEndian.PutUint32(b[:], pid)
	binary.LittleEndian.PutUint64(b[4:], uint64(time.Now().UnixNano()))
	return base64.URLEncoding.EncodeToString(b[:])
}

// --------------------------------------------------------------------

type key int // key is unexported and used for Context

const (
	reqidKey key = 0
)

// NewContext creates a new context with a reqid.
//
func NewContext(ctx context.Context, reqid string) context.Context {
	return context.WithValue(ctx, reqidKey, reqid)
}

// NewContextWith creates a new context which gets reqid from a req.Header object.
//
func NewContextWith(ctx context.Context, w http.ResponseWriter, req *http.Request) context.Context {
	reqid := req.Header.Get("X-Reqid")
	if reqid == "" {
		reqid = genReqID()
		req.Header.Set("X-Reqid", reqid)
	}
	h := w.Header()
	h.Set("X-Reqid", reqid)
	return context.WithValue(ctx, reqidKey, reqid)
}

// FromContext gets reqid from ctx.
//
func FromContext(ctx context.Context) (reqid string, ok bool) {
	reqid, ok = ctx.Value(reqidKey).(string)
	return
}

// --------------------------------------------------------------------
