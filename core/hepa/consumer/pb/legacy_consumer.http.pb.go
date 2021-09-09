// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: legacy_consumer.proto

package pb

import (
	context "context"
	http1 "net/http"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// LegacyConsumerServiceHandler is the server API for LegacyConsumerService service.
type LegacyConsumerServiceHandler interface {
	// +publish path:"/api/gateway/consumer"
	// GET /api/gateway/consumer
	GetConsumer(context.Context, *GetConsumerRequest) (*GetConsumerResponse, error)
}

// RegisterLegacyConsumerServiceHandler register LegacyConsumerServiceHandler to http.Router.
func RegisterLegacyConsumerServiceHandler(r http.Router, srv LegacyConsumerServiceHandler, opts ...http.HandleOption) {
	h := http.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}
	encodeFunc := func(fn func(http1.ResponseWriter, *http1.Request) (interface{}, error)) http.HandlerFunc {
		return func(w http1.ResponseWriter, r *http1.Request) {
			out, err := fn(w, r)
			if err != nil {
				h.Error(w, r, err)
				return
			}
			if err := h.Encode(w, r, out); err != nil {
				h.Error(w, r, err)
			}
		}
	}

	add_GetConsumer := func(method, path string, fn func(context.Context, *GetConsumerRequest) (*GetConsumerResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetConsumerRequest))
		}
		var GetConsumer_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetConsumer_info = transport.NewServiceInfo("erda.core.hepa.consumer.LegacyConsumerService", "GetConsumer", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetConsumer_info)
				}
				r = r.WithContext(ctx)
				var in GetConsumerRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["orgId"]; len(vals) > 0 {
					in.OrgId = vals[0]
				}
				if vals := params["projectId"]; len(vals) > 0 {
					in.ProjectId = vals[0]
				}
				if vals := params["env"]; len(vals) > 0 {
					in.Env = vals[0]
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetConsumer("GET", "/api/gateway/consumer", srv.GetConsumer)
}
