package panicintc

import (
	"context"
	"fmt"

	"github.com/decvault/library/common/config"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PanicHandlerInterceptor struct {
	fx.Out

	Interceptor grpc.UnaryServerInterceptor `name:"panic_handler"`
}

func newPanicHandlerInterceptor(
	appMeta config.AppMeta,
	logger *logrus.Logger,
) PanicHandlerInterceptor {
	return PanicHandlerInterceptor{
		Interceptor: func(
			ctx context.Context,
			req any,
			info *grpc.UnaryServerInfo,
			handler grpc.UnaryHandler,
		) (resp any, err error) {
			defer func() {
				if r := recover(); r != nil {
					logger.
						WithContext(ctx).
						Errorf("unhandeled panic occured in method %s\n: %+v", info.FullMethod, r)

					message := "internal server error"
					if appMeta.Stage != config.Prod {
						message = fmt.Sprint(r)
					}

					err = status.Errorf(codes.Internal, message)
				}
			}()

			return handler(ctx, req)
		},
	}
}
