// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: api_policy.proto

package pb

import (
	context "context"
	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	httprule "github.com/erda-project/erda-infra/pkg/transport/http/httprule"
	runtime "github.com/erda-project/erda-infra/pkg/transport/http/runtime"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	http1 "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// ApiPolicyServiceHandler is the server API for ApiPolicyService service.
type ApiPolicyServiceHandler interface {
	// GET /api/gateway/policies/{category}
	GetPolicy(context.Context, *GetPolicyRequest) (*GetPolicyResponse, error)
	// PUT /api/gateway/policies/{category}
	SetPolicy(context.Context, *SetPolicyRequest) (*SetPolicyResponse, error)
}

// RegisterApiPolicyServiceHandler register ApiPolicyServiceHandler to http.Router.
func RegisterApiPolicyServiceHandler(r http.Router, srv ApiPolicyServiceHandler, opts ...http.HandleOption) {
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

	add_GetPolicy := func(method, path string, fn func(context.Context, *GetPolicyRequest) (*GetPolicyResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetPolicyRequest))
		}
		var GetPolicy_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetPolicy_info = transport.NewServiceInfo("erda.core.hepa.api_policy.ApiPolicyService", "GetPolicy", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetPolicy_info)
				}
				r = r.WithContext(ctx)
				var in GetPolicyRequest
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
				if vals := params["packageId"]; len(vals) > 0 {
					in.PackageId = vals[0]
				}
				if vals := params["apiId"]; len(vals) > 0 {
					in.ApiId = vals[0]
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "category":
							in.Category = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_SetPolicy := func(method, path string, fn func(context.Context, *SetPolicyRequest) (*SetPolicyResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*SetPolicyRequest))
		}
		var SetPolicy_info transport.ServiceInfo
		if h.Interceptor != nil {
			SetPolicy_info = transport.NewServiceInfo("erda.core.hepa.api_policy.ApiPolicyService", "SetPolicy", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, SetPolicy_info)
				}
				r = r.WithContext(ctx)
				var in SetPolicyRequest
				if err := h.Decode(r, &in.Body); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["packageId"]; len(vals) > 0 {
					in.PackageId = vals[0]
				}
				if vals := params["apiId"]; len(vals) > 0 {
					in.ApiId = vals[0]
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "category":
							in.Category = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetPolicy("GET", "/api/gateway/policies/{category}", srv.GetPolicy)
	add_SetPolicy("PUT", "/api/gateway/policies/{category}", srv.SetPolicy)
}
