// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: autotest-plan.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*TestPlanUpdateByHookRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Content)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestPlanUpdateByHookResponse)(nil)

// TestPlanUpdateByHookRequest implement urlenc.URLValuesUnmarshaler.
func (m *TestPlanUpdateByHookRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "event":
				m.Event = vals[0]
			case "action":
				m.Action = vals[0]
			case "orgID":
				m.OrgID = vals[0]
			case "projectID":
				m.ProjectID = vals[0]
			case "applicationID":
				m.ApplicationID = vals[0]
			case "env":
				m.Env = vals[0]
			case "timestamp":
				m.Timestamp = vals[0]
			case "content":
				if m.Content == nil {
					m.Content = &Content{}
				}
			case "content.testPlanID":
				if m.Content == nil {
					m.Content = &Content{}
				}
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Content.TestPlanID = val
			case "content.executeTime":
				if m.Content == nil {
					m.Content = &Content{}
				}
				m.Content.ExecuteTime = vals[0]
			case "content.passRate":
				if m.Content == nil {
					m.Content = &Content{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Content.PassRate = val
			case "content.executeMinutes":
				if m.Content == nil {
					m.Content = &Content{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Content.ExecuteMinutes = val
			case "content.apiTotalNum":
				if m.Content == nil {
					m.Content = &Content{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Content.ApiTotalNum = val
			}
		}
	}
	return nil
}

// Content implement urlenc.URLValuesUnmarshaler.
func (m *Content) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "testPlanID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.TestPlanID = val
			case "executeTime":
				m.ExecuteTime = vals[0]
			case "passRate":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.PassRate = val
			case "executeMinutes":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.ExecuteMinutes = val
			case "apiTotalNum":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ApiTotalNum = val
			}
		}
	}
	return nil
}

// TestPlanUpdateByHookResponse implement urlenc.URLValuesUnmarshaler.
func (m *TestPlanUpdateByHookResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data = val
			}
		}
	}
	return nil
}