package xlog

import (
	"context"
	"net/http"

	"qiniupkg.com/x/reqid.v7"
)

// ============================================================================

func NewContext(ctx context.Context, xl *Logger) context.Context {

	if xl != nil {
		ctx = reqid.NewContext(ctx, xl.ReqId())
	}
	return ctx
}

func NewContextWithReq(ctx context.Context, req *http.Request) context.Context {

	return NewContext(ctx, NewWithReq(req))
}

// NewContextWith creates a context with:
// 	1. provided req id (if @a is string or reqIder)
// 	2. provided header (if @a is header)
//	3. **DUMMY** trace recorder (if @a cannot record)
//
func NewContextWith(ctx context.Context, a interface{}) context.Context {

	return NewContext(ctx, NewWith(a))
}

func NewContextWithRW(ctx context.Context, w http.ResponseWriter, r *http.Request) context.Context {

	return NewContext(ctx, New(w, r))
}

func FromContext(ctx context.Context) (xl *Logger, ok bool) {

	v, ok := reqid.FromContext(ctx)
	if ok {
		xl = NewWith(v)
	}
	return
}

func FromContextSafe(ctx context.Context) (xl *Logger) {

	v, ok := reqid.FromContext(ctx)
	if ok {
		xl = NewWith(v)
	} else {
		xl = NewDummy()
	}
	return
}

// ============================================================================
