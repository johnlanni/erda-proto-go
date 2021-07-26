// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: tenant.proto

package pb

import (
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	url "net/url"
	strconv "strconv"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*CreateTenantRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateTenantResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetTenantRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetTenantResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Tenant)(nil)

// CreateTenantRequest implement urlenc.URLValuesUnmarshaler.
func (m *CreateTenantRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "projectID":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ProjectID = val
			case "tenantType":
				m.TenantType = vals[0]
			case "workspaces":
				m.Workspaces = vals
			}
		}
	}
	return nil
}

// CreateTenantResponse implement urlenc.URLValuesUnmarshaler.
func (m *CreateTenantResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// GetTenantRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetTenantRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "projectID":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ProjectID = val
			case "tenantType":
				m.TenantType = vals[0]
			case "workspace":
				m.Workspace = vals[0]
			}
		}
	}
	return nil
}

// GetTenantResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetTenantResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &Tenant{}
				}
			case "data.id":
				if m.Data == nil {
					m.Data = &Tenant{}
				}
				m.Data.Id = vals[0]
			case "data.type":
				if m.Data == nil {
					m.Data = &Tenant{}
				}
				m.Data.Type = vals[0]
			case "data.projectID":
				if m.Data == nil {
					m.Data = &Tenant{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.ProjectID = val
			case "data.workspace":
				if m.Data == nil {
					m.Data = &Tenant{}
				}
				m.Data.Workspace = vals[0]
			case "data.createTime":
				if m.Data == nil {
					m.Data = &Tenant{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.CreateTime = val
			case "data.updateTime":
				if m.Data == nil {
					m.Data = &Tenant{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.UpdateTime = val
			case "data.isDeleted":
				if m.Data == nil {
					m.Data = &Tenant{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data.IsDeleted = val
			}
		}
	}
	return nil
}

// Tenant implement urlenc.URLValuesUnmarshaler.
func (m *Tenant) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "type":
				m.Type = vals[0]
			case "projectID":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ProjectID = val
			case "workspace":
				m.Workspace = vals[0]
			case "createTime":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.CreateTime = val
			case "updateTime":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.UpdateTime = val
			case "isDeleted":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.IsDeleted = val
			}
		}
	}
	return nil
}