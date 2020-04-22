// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// HeaderManipulationAction An object that represents an action to apply to an HTTP headers.
type HeaderManipulationAction interface {
}

type headermanipulationaction struct {
	JsonData []byte
	Action   string `json:"action"`
}

// UnmarshalJSON unmarshals json
func (m *headermanipulationaction) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerheadermanipulationaction headermanipulationaction
	s := struct {
		Model Unmarshalerheadermanipulationaction
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Action = s.Model.Action

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *headermanipulationaction) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Action {
	case "EXTEND_HTTP_RESPONSE_HEADER":
		mm := ExtendHttpResponseHeaderAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ADD_HTTP_RESPONSE_HEADER":
		mm := AddHttpResponseHeaderAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REMOVE_HTTP_RESPONSE_HEADER":
		mm := RemoveHttpResponseHeaderAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m headermanipulationaction) String() string {
	return common.PointerString(m)
}

// HeaderManipulationActionActionEnum Enum with underlying type: string
type HeaderManipulationActionActionEnum string

// Set of constants representing the allowable values for HeaderManipulationActionActionEnum
const (
	HeaderManipulationActionActionExtendHttpResponseHeader HeaderManipulationActionActionEnum = "EXTEND_HTTP_RESPONSE_HEADER"
	HeaderManipulationActionActionAddHttpResponseHeader    HeaderManipulationActionActionEnum = "ADD_HTTP_RESPONSE_HEADER"
	HeaderManipulationActionActionRemoveHttpResponseHeader HeaderManipulationActionActionEnum = "REMOVE_HTTP_RESPONSE_HEADER"
)

var mappingHeaderManipulationActionAction = map[string]HeaderManipulationActionActionEnum{
	"EXTEND_HTTP_RESPONSE_HEADER": HeaderManipulationActionActionExtendHttpResponseHeader,
	"ADD_HTTP_RESPONSE_HEADER":    HeaderManipulationActionActionAddHttpResponseHeader,
	"REMOVE_HTTP_RESPONSE_HEADER": HeaderManipulationActionActionRemoveHttpResponseHeader,
}

// GetHeaderManipulationActionActionEnumValues Enumerates the set of values for HeaderManipulationActionActionEnum
func GetHeaderManipulationActionActionEnumValues() []HeaderManipulationActionActionEnum {
	values := make([]HeaderManipulationActionActionEnum, 0)
	for _, v := range mappingHeaderManipulationActionAction {
		values = append(values, v)
	}
	return values
}
