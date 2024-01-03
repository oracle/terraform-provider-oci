// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// TestConnectionAssignmentResult The result of the connectivity test performed between the GoldenGate deployment and the associated database / service.
type TestConnectionAssignmentResult struct {

	// Type of the result (i.e. Success, Failure or Timeout).
	ResultType TestConnectionAssignmentResultResultTypeEnum `mandatory:"true" json:"resultType"`

	Error *TestConnectionAssignmentError `mandatory:"false" json:"error"`
}

func (m TestConnectionAssignmentResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TestConnectionAssignmentResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTestConnectionAssignmentResultResultTypeEnum(string(m.ResultType)); !ok && m.ResultType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResultType: %s. Supported values are: %s.", m.ResultType, strings.Join(GetTestConnectionAssignmentResultResultTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TestConnectionAssignmentResultResultTypeEnum Enum with underlying type: string
type TestConnectionAssignmentResultResultTypeEnum string

// Set of constants representing the allowable values for TestConnectionAssignmentResultResultTypeEnum
const (
	TestConnectionAssignmentResultResultTypeSucceeded TestConnectionAssignmentResultResultTypeEnum = "SUCCEEDED"
	TestConnectionAssignmentResultResultTypeFailed    TestConnectionAssignmentResultResultTypeEnum = "FAILED"
	TestConnectionAssignmentResultResultTypeTimedOut  TestConnectionAssignmentResultResultTypeEnum = "TIMED_OUT"
)

var mappingTestConnectionAssignmentResultResultTypeEnum = map[string]TestConnectionAssignmentResultResultTypeEnum{
	"SUCCEEDED": TestConnectionAssignmentResultResultTypeSucceeded,
	"FAILED":    TestConnectionAssignmentResultResultTypeFailed,
	"TIMED_OUT": TestConnectionAssignmentResultResultTypeTimedOut,
}

var mappingTestConnectionAssignmentResultResultTypeEnumLowerCase = map[string]TestConnectionAssignmentResultResultTypeEnum{
	"succeeded": TestConnectionAssignmentResultResultTypeSucceeded,
	"failed":    TestConnectionAssignmentResultResultTypeFailed,
	"timed_out": TestConnectionAssignmentResultResultTypeTimedOut,
}

// GetTestConnectionAssignmentResultResultTypeEnumValues Enumerates the set of values for TestConnectionAssignmentResultResultTypeEnum
func GetTestConnectionAssignmentResultResultTypeEnumValues() []TestConnectionAssignmentResultResultTypeEnum {
	values := make([]TestConnectionAssignmentResultResultTypeEnum, 0)
	for _, v := range mappingTestConnectionAssignmentResultResultTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTestConnectionAssignmentResultResultTypeEnumStringValues Enumerates the set of values in String for TestConnectionAssignmentResultResultTypeEnum
func GetTestConnectionAssignmentResultResultTypeEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
		"TIMED_OUT",
	}
}

// GetMappingTestConnectionAssignmentResultResultTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTestConnectionAssignmentResultResultTypeEnum(val string) (TestConnectionAssignmentResultResultTypeEnum, bool) {
	enum, ok := mappingTestConnectionAssignmentResultResultTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
