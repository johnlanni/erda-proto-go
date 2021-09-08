// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: org_client.proto

package pb

import (
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	url "net/url"
	strconv "strconv"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*ChangeClientLimitResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ChangeClientLimitRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GrantEndpointRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GrantEndpointResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*RevokeEndpointRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*RevokeEndpointResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*UpdateCredentialsResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*UpdateCredentialsRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetCredentialsResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetCredentialsRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteClientRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteClientResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ClientInfo)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateClientRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateClientResponse)(nil)

// ChangeClientLimitResponse implement urlenc.URLValuesUnmarshaler.
func (m *ChangeClientLimitResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data = val
			}
		}
	}
	return nil
}

// ChangeClientLimitRequest implement urlenc.URLValuesUnmarshaler.
func (m *ChangeClientLimitRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "clientId":
				m.ClientId = vals[0]
			case "packageId":
				m.PackageId = vals[0]
			}
		}
	}
	return nil
}

// GrantEndpointRequest implement urlenc.URLValuesUnmarshaler.
func (m *GrantEndpointRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "clientId":
				m.ClientId = vals[0]
			case "packageId":
				m.PackageId = vals[0]
			}
		}
	}
	return nil
}

// GrantEndpointResponse implement urlenc.URLValuesUnmarshaler.
func (m *GrantEndpointResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data = val
			}
		}
	}
	return nil
}

// RevokeEndpointRequest implement urlenc.URLValuesUnmarshaler.
func (m *RevokeEndpointRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "clientId":
				m.ClientId = vals[0]
			case "packageId":
				m.PackageId = vals[0]
			}
		}
	}
	return nil
}

// RevokeEndpointResponse implement urlenc.URLValuesUnmarshaler.
func (m *RevokeEndpointResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data = val
			}
		}
	}
	return nil
}

// UpdateCredentialsResponse implement urlenc.URLValuesUnmarshaler.
func (m *UpdateCredentialsResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &ClientInfo{}
				}
			case "data.clientId":
				if m.Data == nil {
					m.Data = &ClientInfo{}
				}
				m.Data.ClientId = vals[0]
			case "data.clientSecret":
				if m.Data == nil {
					m.Data = &ClientInfo{}
				}
				m.Data.ClientSecret = vals[0]
			}
		}
	}
	return nil
}

// UpdateCredentialsRequest implement urlenc.URLValuesUnmarshaler.
func (m *UpdateCredentialsRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "clientId":
				m.ClientId = vals[0]
			case "clientSecret":
				m.ClientSecret = vals[0]
			}
		}
	}
	return nil
}

// GetCredentialsResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetCredentialsResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &ClientInfo{}
				}
			case "data.clientId":
				if m.Data == nil {
					m.Data = &ClientInfo{}
				}
				m.Data.ClientId = vals[0]
			case "data.clientSecret":
				if m.Data == nil {
					m.Data = &ClientInfo{}
				}
				m.Data.ClientSecret = vals[0]
			}
		}
	}
	return nil
}

// GetCredentialsRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetCredentialsRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "clientId":
				m.ClientId = vals[0]
			}
		}
	}
	return nil
}

// DeleteClientRequest implement urlenc.URLValuesUnmarshaler.
func (m *DeleteClientRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "clientId":
				m.ClientId = vals[0]
			}
		}
	}
	return nil
}

// DeleteClientResponse implement urlenc.URLValuesUnmarshaler.
func (m *DeleteClientResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data = val
			}
		}
	}
	return nil
}

// ClientInfo implement urlenc.URLValuesUnmarshaler.
func (m *ClientInfo) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "clientId":
				m.ClientId = vals[0]
			case "clientSecret":
				m.ClientSecret = vals[0]
			}
		}
	}
	return nil
}

// CreateClientRequest implement urlenc.URLValuesUnmarshaler.
func (m *CreateClientRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "clientName":
				m.ClientName = vals[0]
			}
		}
	}
	return nil
}

// CreateClientResponse implement urlenc.URLValuesUnmarshaler.
func (m *CreateClientResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &ClientInfo{}
				}
			case "data.clientId":
				if m.Data == nil {
					m.Data = &ClientInfo{}
				}
				m.Data.ClientId = vals[0]
			case "data.clientSecret":
				if m.Data == nil {
					m.Data = &ClientInfo{}
				}
				m.Data.ClientSecret = vals[0]
			}
		}
	}
	return nil
}
