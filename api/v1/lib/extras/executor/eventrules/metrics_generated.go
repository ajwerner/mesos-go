package eventrules

// go generate -import github.com/mesos/mesos-go/api/v1/lib/executor -type E:*executor.Event:&executor.Event{} -output metrics_generated.go
// GENERATED CODE FOLLOWS; DO NOT EDIT.

import (
	"context"
	"strings"

	"github.com/mesos/mesos-go/api/v1/lib/extras/metrics"

	"github.com/mesos/mesos-go/api/v1/lib/executor"
)

func Metrics(harness metrics.Harness) Rule {
	return func(ctx context.Context, e *executor.Event, err error, ch Chain) (context.Context, *executor.Event, error) {
		typename := strings.ToLower(e.GetType().String())
		harness(func() error {
			ctx, e, err = ch(ctx, e, err)
			return err
		}, typename)
		return ctx, e, err
	}
}
