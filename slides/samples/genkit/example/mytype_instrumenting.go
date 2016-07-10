// generated by genkit -- DO NOT EDIT
package example

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

func instrumentingMiddleware(
	requestCount metrics.Counter,
	requestLatency metrics.Histogram,
) ServiceMiddleware {
	return func(next UserService) UserService {
		return instrmw{requestCount, requestLatency, next}
	}
}

type instrmw struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	UserService
}

func (mw instrmw) Create(c User) (u string, err error) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "create"}
		errorField := metrics.Field{Key: "error", Value: fmt.Sprintf("%v", err)}
		mw.requestCount.With(methodField).With(errorField).Add(1)
		var durInt int64 = time.Since(begin).Nanoseconds() / 1e6
		mw.requestLatency.With(methodField).With(errorField).Observe(durInt)
	}(time.Now())

	u, err = mw.UserService.Create(c)
	return
}

func (mw instrmw) Get(id string) (m User, err error) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "get"}
		errorField := metrics.Field{Key: "error", Value: fmt.Sprintf("%v", err)}
		mw.requestCount.With(methodField).With(errorField).Add(1)
		var durInt int64 = time.Since(begin).Nanoseconds() / 1e6
		mw.requestLatency.With(methodField).With(errorField).Observe(durInt)
	}(time.Now())

	m, err = mw.UserService.Get(id)
	return
}

func (mw instrmw) Update(m User) (err error) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "update"}
		errorField := metrics.Field{Key: "error", Value: fmt.Sprintf("%v", err)}
		mw.requestCount.With(methodField).With(errorField).Add(1)
		var durInt int64 = time.Since(begin).Nanoseconds() / 1e6
		mw.requestLatency.With(methodField).With(errorField).Observe(durInt)
	}(time.Now())

	err = mw.UserService.Update(m)
	return
}

func (mw instrmw) List() (list []User, err error) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "list"}
		errorField := metrics.Field{Key: "error", Value: fmt.Sprintf("%v", err)}
		mw.requestCount.With(methodField).With(errorField).Add(1)
		var durInt int64 = time.Since(begin).Nanoseconds() / 1e6
		mw.requestLatency.With(methodField).With(errorField).Observe(durInt)
	}(time.Now())

	list, err = mw.UserService.List()
	return
}

func (mw instrmw) Delete(id string) (err error) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "delete"}
		errorField := metrics.Field{Key: "error", Value: fmt.Sprintf("%v", err)}
		mw.requestCount.With(methodField).With(errorField).Add(1)
		var durInt int64 = time.Since(begin).Nanoseconds() / 1e6
		mw.requestLatency.With(methodField).With(errorField).Observe(durInt)
	}(time.Now())

	err = mw.UserService.Delete(id)
	return
}
