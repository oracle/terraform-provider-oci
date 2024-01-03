// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GetPathAnalysisDetails Defines the configuration for getting a path analysis.
type GetPathAnalysisDetails interface {
}

type getpathanalysisdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *getpathanalysisdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalergetpathanalysisdetails getpathanalysisdetails
	s := struct {
		Model Unmarshalergetpathanalysisdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *getpathanalysisdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "ADHOC_QUERY":
		mm := AdhocGetPathAnalysisDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PERSISTED_QUERY":
		mm := PersistedGetPathAnalysisDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for GetPathAnalysisDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m getpathanalysisdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m getpathanalysisdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetPathAnalysisDetailsTypeEnum Enum with underlying type: string
type GetPathAnalysisDetailsTypeEnum string

// Set of constants representing the allowable values for GetPathAnalysisDetailsTypeEnum
const (
	GetPathAnalysisDetailsTypePersistedQuery GetPathAnalysisDetailsTypeEnum = "PERSISTED_QUERY"
	GetPathAnalysisDetailsTypeAdhocQuery     GetPathAnalysisDetailsTypeEnum = "ADHOC_QUERY"
)

var mappingGetPathAnalysisDetailsTypeEnum = map[string]GetPathAnalysisDetailsTypeEnum{
	"PERSISTED_QUERY": GetPathAnalysisDetailsTypePersistedQuery,
	"ADHOC_QUERY":     GetPathAnalysisDetailsTypeAdhocQuery,
}

var mappingGetPathAnalysisDetailsTypeEnumLowerCase = map[string]GetPathAnalysisDetailsTypeEnum{
	"persisted_query": GetPathAnalysisDetailsTypePersistedQuery,
	"adhoc_query":     GetPathAnalysisDetailsTypeAdhocQuery,
}

// GetGetPathAnalysisDetailsTypeEnumValues Enumerates the set of values for GetPathAnalysisDetailsTypeEnum
func GetGetPathAnalysisDetailsTypeEnumValues() []GetPathAnalysisDetailsTypeEnum {
	values := make([]GetPathAnalysisDetailsTypeEnum, 0)
	for _, v := range mappingGetPathAnalysisDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetPathAnalysisDetailsTypeEnumStringValues Enumerates the set of values in String for GetPathAnalysisDetailsTypeEnum
func GetGetPathAnalysisDetailsTypeEnumStringValues() []string {
	return []string{
		"PERSISTED_QUERY",
		"ADHOC_QUERY",
	}
}

// GetMappingGetPathAnalysisDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetPathAnalysisDetailsTypeEnum(val string) (GetPathAnalysisDetailsTypeEnum, bool) {
	enum, ok := mappingGetPathAnalysisDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
