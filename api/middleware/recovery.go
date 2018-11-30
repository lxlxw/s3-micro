package middleware

import (
	"runtime"

	"github.com/wothing/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const (
	MAXSTACKSIZE = 4096
)

// Recovery interceptor to handle grpc panic
func Recovery(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// recovery func
	defer func() {
		if r := recover(); r != nil {
			// log stack
			stack := make([]byte, MAXSTACKSIZE)
			stack = stack[:runtime.Stack(stack, false)]
			log.CtxErrorf(ctx, "panic grpc invoke: %s, err=%v, stack:\n%s", info.FullMethod, r, string(stack))

			err = grpc.Errorf(codes.Internal, "panic error: %v", r)
		}
	}()

	return handler(ctx, req)
}
