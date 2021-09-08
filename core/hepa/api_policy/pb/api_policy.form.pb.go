// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: api_policy.proto

package pb

import (
	base64 "encoding/base64"
	json "encoding/json"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	structpb "google.golang.org/protobuf/types/known/structpb"
	url "net/url"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*SetPolicyResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*SetPolicyRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetPolicyRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetPolicyResponse)(nil)

// SetPolicyResponse implement urlenc.URLValuesUnmarshaler.
func (m *SetPolicyResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Data = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Data = val
					} else {
						m.Data = structpb.NewStringValue(vals[0])
					}
				}
			}
		}
	}
	return nil
}

// SetPolicyRequest implement urlenc.URLValuesUnmarshaler.
func (m *SetPolicyRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "category":
				m.Category = vals[0]
			case "packageId":
				m.PackageId = vals[0]
			case "apiId":
				m.ApiId = vals[0]
			case "body":
				val, err := base64.StdEncoding.DecodeString(vals[0])
				if err != nil {
					return err
				}
				m.Body = val
			}
		}
	}
	return nil
}

// GetPolicyRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetPolicyRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "category":
				m.Category = vals[0]
			case "packageId":
				m.PackageId = vals[0]
			case "apiId":
				m.ApiId = vals[0]
			}
		}
	}
	return nil
}

// GetPolicyResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetPolicyResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Data = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Data = val
					} else {
						m.Data = structpb.NewStringValue(vals[0])
					}
				}
			}
		}
	}
	return nil
}
