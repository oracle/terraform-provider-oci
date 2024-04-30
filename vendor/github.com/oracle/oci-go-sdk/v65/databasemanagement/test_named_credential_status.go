// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TestNamedCredentialStatus The status of the named credential test. The status is 'SUCCEEDED' if the named credential is working or else the status is 'FAILED'.
type TestNamedCredentialStatus struct {

	// The status of the named credential test. The status is 'SUCCEEDED' if the named credential is working or else the status is 'FAILED'.
	Status TestNamedCredentialStatusStatusEnum `mandatory:"true" json:"status"`

	// An error code that defines the failure of the named credential test. The response is 'null' if the named credential test was successful.
	ErrorCode *string `mandatory:"false" json:"errorCode"`

	// The error message that indicates the reason for the failure of the named credential test. The response is 'null' if the named credential test was successful.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`
}

func (m TestNamedCredentialStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TestNamedCredentialStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTestNamedCredentialStatusStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetTestNamedCredentialStatusStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TestNamedCredentialStatusStatusEnum Enum with underlying type: string
type TestNamedCredentialStatusStatusEnum string

// Set of constants representing the allowable values for TestNamedCredentialStatusStatusEnum
const (
	TestNamedCredentialStatusStatusSucceeded TestNamedCredentialStatusStatusEnum = "SUCCEEDED"
	TestNamedCredentialStatusStatusFailed    TestNamedCredentialStatusStatusEnum = "FAILED"
)

var mappingTestNamedCredentialStatusStatusEnum = map[string]TestNamedCredentialStatusStatusEnum{
	"SUCCEEDED": TestNamedCredentialStatusStatusSucceeded,
	"FAILED":    TestNamedCredentialStatusStatusFailed,
}

var mappingTestNamedCredentialStatusStatusEnumLowerCase = map[string]TestNamedCredentialStatusStatusEnum{
	"succeeded": TestNamedCredentialStatusStatusSucceeded,
	"failed":    TestNamedCredentialStatusStatusFailed,
}

// GetTestNamedCredentialStatusStatusEnumValues Enumerates the set of values for TestNamedCredentialStatusStatusEnum
func GetTestNamedCredentialStatusStatusEnumValues() []TestNamedCredentialStatusStatusEnum {
	values := make([]TestNamedCredentialStatusStatusEnum, 0)
	for _, v := range mappingTestNamedCredentialStatusStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetTestNamedCredentialStatusStatusEnumStringValues Enumerates the set of values in String for TestNamedCredentialStatusStatusEnum
func GetTestNamedCredentialStatusStatusEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingTestNamedCredentialStatusStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTestNamedCredentialStatusStatusEnum(val string) (TestNamedCredentialStatusStatusEnum, bool) {
	enum, ok := mappingTestNamedCredentialStatusStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
