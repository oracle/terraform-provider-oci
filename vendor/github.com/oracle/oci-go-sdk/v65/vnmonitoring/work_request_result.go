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

// WorkRequestResult Ephemeral data resulting from an asynchronous operation.
type WorkRequestResult interface {
}

type workrequestresult struct {
	JsonData   []byte
	ResultType string `json:"resultType"`
}

// UnmarshalJSON unmarshals json
func (m *workrequestresult) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerworkrequestresult workrequestresult
	s := struct {
		Model Unmarshalerworkrequestresult
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ResultType = s.Model.ResultType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *workrequestresult) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ResultType {
	case "PATH_ANALYSIS":
		mm := PathAnalysisWorkRequestResult{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for WorkRequestResult: %s.", m.ResultType)
		return *m, nil
	}
}

func (m workrequestresult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m workrequestresult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WorkRequestResultResultTypeEnum Enum with underlying type: string
type WorkRequestResultResultTypeEnum string

// Set of constants representing the allowable values for WorkRequestResultResultTypeEnum
const (
	WorkRequestResultResultTypePathAnalysis WorkRequestResultResultTypeEnum = "PATH_ANALYSIS"
)

var mappingWorkRequestResultResultTypeEnum = map[string]WorkRequestResultResultTypeEnum{
	"PATH_ANALYSIS": WorkRequestResultResultTypePathAnalysis,
}

var mappingWorkRequestResultResultTypeEnumLowerCase = map[string]WorkRequestResultResultTypeEnum{
	"path_analysis": WorkRequestResultResultTypePathAnalysis,
}

// GetWorkRequestResultResultTypeEnumValues Enumerates the set of values for WorkRequestResultResultTypeEnum
func GetWorkRequestResultResultTypeEnumValues() []WorkRequestResultResultTypeEnum {
	values := make([]WorkRequestResultResultTypeEnum, 0)
	for _, v := range mappingWorkRequestResultResultTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestResultResultTypeEnumStringValues Enumerates the set of values in String for WorkRequestResultResultTypeEnum
func GetWorkRequestResultResultTypeEnumStringValues() []string {
	return []string{
		"PATH_ANALYSIS",
	}
}

// GetMappingWorkRequestResultResultTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestResultResultTypeEnum(val string) (WorkRequestResultResultTypeEnum, bool) {
	enum, ok := mappingWorkRequestResultResultTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
