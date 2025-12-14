package panicintc

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module(
		"panic_handler_interceptor",
		fx.Provide(
			newPanicInterceptor,
		),
	)
}
