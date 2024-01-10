// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TestPreferredCredentialStatus The status of the preferred credential test. The status is 'SUCCEEDED' if the preferred credential is working else the status is 'FAILED'.
type TestPreferredCredentialStatus struct {

	// The status of the preferred credential test. The status is 'SUCCEEDED' if the preferred credential is working else the status is 'FAILED'.
	Status TestPreferredCredentialStatusStatusEnum `mandatory:"false" json:"status,omitempty"`

	// An error code that defines the failure of the preferred credential test. The response is 'null' if the preferred credential test was successful.
	ErrorCode *string `mandatory:"false" json:"errorCode"`

	// The error message that indicates the reason for the failure of the preferred credential test. The response is 'null' if the preferred credential test was successful.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`
}

func (m TestPreferredCredentialStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TestPreferredCredentialStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTestPreferredCredentialStatusStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetTestPreferredCredentialStatusStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TestPreferredCredentialStatusStatusEnum Enum with underlying type: string
type TestPreferredCredentialStatusStatusEnum string

// Set of constants representing the allowable values for TestPreferredCredentialStatusStatusEnum
const (
	TestPreferredCredentialStatusStatusSucceeded TestPreferredCredentialStatusStatusEnum = "SUCCEEDED"
	TestPreferredCredentialStatusStatusFailed    TestPreferredCredentialStatusStatusEnum = "FAILED"
)

var mappingTestPreferredCredentialStatusStatusEnum = map[string]TestPreferredCredentialStatusStatusEnum{
	"SUCCEEDED": TestPreferredCredentialStatusStatusSucceeded,
	"FAILED":    TestPreferredCredentialStatusStatusFailed,
}

var mappingTestPreferredCredentialStatusStatusEnumLowerCase = map[string]TestPreferredCredentialStatusStatusEnum{
	"succeeded": TestPreferredCredentialStatusStatusSucceeded,
	"failed":    TestPreferredCredentialStatusStatusFailed,
}

// GetTestPreferredCredentialStatusStatusEnumValues Enumerates the set of values for TestPreferredCredentialStatusStatusEnum
func GetTestPreferredCredentialStatusStatusEnumValues() []TestPreferredCredentialStatusStatusEnum {
	values := make([]TestPreferredCredentialStatusStatusEnum, 0)
	for _, v := range mappingTestPreferredCredentialStatusStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetTestPreferredCredentialStatusStatusEnumStringValues Enumerates the set of values in String for TestPreferredCredentialStatusStatusEnum
func GetTestPreferredCredentialStatusStatusEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingTestPreferredCredentialStatusStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTestPreferredCredentialStatusStatusEnum(val string) (TestPreferredCredentialStatusStatusEnum, bool) {
	enum, ok := mappingTestPreferredCredentialStatusStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
