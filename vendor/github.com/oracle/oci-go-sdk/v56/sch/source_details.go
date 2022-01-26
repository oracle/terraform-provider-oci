// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Connector Hub API
//
// Use the Service Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Service Connector Hub, see
// Service Connector Hub Overview (https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/overview.htm).
//

package sch

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// SourceDetails An object that represents the source of the flow defined by the service connector.
// An example source is the VCNFlow logs within the NetworkLogs group.
// For more information about flows defined by service connectors, see
// Service Connector Hub Overview (https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/overview.htm).
// For configuration instructions, see
// To create a service connector (https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/managingconnectors.htm#create).
type SourceDetails interface {
}

type sourcedetails struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *sourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersourcedetails sourcedetails
	s := struct {
		Model Unmarshalersourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *sourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "logging":
		mm := LoggingSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "streaming":
		mm := StreamingSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m sourcedetails) String() string {
	return common.PointerString(m)
}

// SourceDetailsKindEnum Enum with underlying type: string
type SourceDetailsKindEnum string

// Set of constants representing the allowable values for SourceDetailsKindEnum
const (
	SourceDetailsKindLogging   SourceDetailsKindEnum = "logging"
	SourceDetailsKindStreaming SourceDetailsKindEnum = "streaming"
)

var mappingSourceDetailsKind = map[string]SourceDetailsKindEnum{
	"logging":   SourceDetailsKindLogging,
	"streaming": SourceDetailsKindStreaming,
}

// GetSourceDetailsKindEnumValues Enumerates the set of values for SourceDetailsKindEnum
func GetSourceDetailsKindEnumValues() []SourceDetailsKindEnum {
	values := make([]SourceDetailsKindEnum, 0)
	for _, v := range mappingSourceDetailsKind {
		values = append(values, v)
	}
	return values
}
