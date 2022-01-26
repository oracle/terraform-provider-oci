// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Firewall (WAF) API
//
// API for the Web Application Firewall service.
// Use this API to manage regional Web App Firewalls and corresponding policies for protecting HTTP services.
//

package waf

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// HttpResponseBody Type of returned HTTP response body.
type HttpResponseBody interface {
}

type httpresponsebody struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *httpresponsebody) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerhttpresponsebody httpresponsebody
	s := struct {
		Model Unmarshalerhttpresponsebody
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *httpresponsebody) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "STATIC_TEXT":
		mm := StaticTextHttpResponseBody{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m httpresponsebody) String() string {
	return common.PointerString(m)
}

// HttpResponseBodyTypeEnum Enum with underlying type: string
type HttpResponseBodyTypeEnum string

// Set of constants representing the allowable values for HttpResponseBodyTypeEnum
const (
	HttpResponseBodyTypeStaticText HttpResponseBodyTypeEnum = "STATIC_TEXT"
)

var mappingHttpResponseBodyType = map[string]HttpResponseBodyTypeEnum{
	"STATIC_TEXT": HttpResponseBodyTypeStaticText,
}

// GetHttpResponseBodyTypeEnumValues Enumerates the set of values for HttpResponseBodyTypeEnum
func GetHttpResponseBodyTypeEnumValues() []HttpResponseBodyTypeEnum {
	values := make([]HttpResponseBodyTypeEnum, 0)
	for _, v := range mappingHttpResponseBodyType {
		values = append(values, v)
	}
	return values
}
