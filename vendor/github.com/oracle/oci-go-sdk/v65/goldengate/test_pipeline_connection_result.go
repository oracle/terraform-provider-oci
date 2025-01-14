// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TestPipelineConnectionResult Result of the connectivity test performed on a pipeline's assigned connection.
type TestPipelineConnectionResult struct {

	// Type of result, either Succeeded, Failed, or Timed out.
	ResultType TestPipelineConnectionResultResultTypeEnum `mandatory:"true" json:"resultType"`

	Error *TestPipelineConnectionError `mandatory:"false" json:"error"`
}

func (m TestPipelineConnectionResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TestPipelineConnectionResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTestPipelineConnectionResultResultTypeEnum(string(m.ResultType)); !ok && m.ResultType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResultType: %s. Supported values are: %s.", m.ResultType, strings.Join(GetTestPipelineConnectionResultResultTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TestPipelineConnectionResultResultTypeEnum Enum with underlying type: string
type TestPipelineConnectionResultResultTypeEnum string

// Set of constants representing the allowable values for TestPipelineConnectionResultResultTypeEnum
const (
	TestPipelineConnectionResultResultTypeSucceeded TestPipelineConnectionResultResultTypeEnum = "SUCCEEDED"
	TestPipelineConnectionResultResultTypeFailed    TestPipelineConnectionResultResultTypeEnum = "FAILED"
	TestPipelineConnectionResultResultTypeTimedOut  TestPipelineConnectionResultResultTypeEnum = "TIMED_OUT"
)

var mappingTestPipelineConnectionResultResultTypeEnum = map[string]TestPipelineConnectionResultResultTypeEnum{
	"SUCCEEDED": TestPipelineConnectionResultResultTypeSucceeded,
	"FAILED":    TestPipelineConnectionResultResultTypeFailed,
	"TIMED_OUT": TestPipelineConnectionResultResultTypeTimedOut,
}

var mappingTestPipelineConnectionResultResultTypeEnumLowerCase = map[string]TestPipelineConnectionResultResultTypeEnum{
	"succeeded": TestPipelineConnectionResultResultTypeSucceeded,
	"failed":    TestPipelineConnectionResultResultTypeFailed,
	"timed_out": TestPipelineConnectionResultResultTypeTimedOut,
}

// GetTestPipelineConnectionResultResultTypeEnumValues Enumerates the set of values for TestPipelineConnectionResultResultTypeEnum
func GetTestPipelineConnectionResultResultTypeEnumValues() []TestPipelineConnectionResultResultTypeEnum {
	values := make([]TestPipelineConnectionResultResultTypeEnum, 0)
	for _, v := range mappingTestPipelineConnectionResultResultTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTestPipelineConnectionResultResultTypeEnumStringValues Enumerates the set of values in String for TestPipelineConnectionResultResultTypeEnum
func GetTestPipelineConnectionResultResultTypeEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
		"TIMED_OUT",
	}
}

// GetMappingTestPipelineConnectionResultResultTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTestPipelineConnectionResultResultTypeEnum(val string) (TestPipelineConnectionResultResultTypeEnum, bool) {
	enum, ok := mappingTestPipelineConnectionResultResultTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
