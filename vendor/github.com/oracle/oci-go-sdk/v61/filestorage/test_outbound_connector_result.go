// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// Use the File Storage service API to manage file systems, mount targets, and snapshots.
// For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v61/common"
	"strings"
)

// TestOutboundConnectorResult Result of connecting to a single endpoint
type TestOutboundConnectorResult struct {

	// Endpoint used while trying to connect while using LDAP bind account
	Endpoint *interface{} `mandatory:"true" json:"endpoint"`

	// Status of connection to given endpoint
	Result TestOutboundConnectorResultResultEnum `mandatory:"true" json:"result"`

	// Error from the perspective of the mount target.
	MountTargetErrorString *string `mandatory:"false" json:"mountTargetErrorString"`

	// Error string from underlying component that failed.
	EndpointErrorString *string `mandatory:"false" json:"endpointErrorString"`
}

func (m TestOutboundConnectorResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TestOutboundConnectorResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTestOutboundConnectorResultResultEnum(string(m.Result)); !ok && m.Result != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Result: %s. Supported values are: %s.", m.Result, strings.Join(GetTestOutboundConnectorResultResultEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TestOutboundConnectorResultResultEnum Enum with underlying type: string
type TestOutboundConnectorResultResultEnum string

// Set of constants representing the allowable values for TestOutboundConnectorResultResultEnum
const (
	TestOutboundConnectorResultResultSuccess TestOutboundConnectorResultResultEnum = "SUCCESS"
	TestOutboundConnectorResultResultFail    TestOutboundConnectorResultResultEnum = "FAIL"
	TestOutboundConnectorResultResultNotrun  TestOutboundConnectorResultResultEnum = "NOTRUN"
)

var mappingTestOutboundConnectorResultResultEnum = map[string]TestOutboundConnectorResultResultEnum{
	"SUCCESS": TestOutboundConnectorResultResultSuccess,
	"FAIL":    TestOutboundConnectorResultResultFail,
	"NOTRUN":  TestOutboundConnectorResultResultNotrun,
}

var mappingTestOutboundConnectorResultResultEnumLowerCase = map[string]TestOutboundConnectorResultResultEnum{
	"success": TestOutboundConnectorResultResultSuccess,
	"fail":    TestOutboundConnectorResultResultFail,
	"notrun":  TestOutboundConnectorResultResultNotrun,
}

// GetTestOutboundConnectorResultResultEnumValues Enumerates the set of values for TestOutboundConnectorResultResultEnum
func GetTestOutboundConnectorResultResultEnumValues() []TestOutboundConnectorResultResultEnum {
	values := make([]TestOutboundConnectorResultResultEnum, 0)
	for _, v := range mappingTestOutboundConnectorResultResultEnum {
		values = append(values, v)
	}
	return values
}

// GetTestOutboundConnectorResultResultEnumStringValues Enumerates the set of values in String for TestOutboundConnectorResultResultEnum
func GetTestOutboundConnectorResultResultEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"FAIL",
		"NOTRUN",
	}
}

// GetMappingTestOutboundConnectorResultResultEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTestOutboundConnectorResultResultEnum(val string) (TestOutboundConnectorResultResultEnum, bool) {
	enum, ok := mappingTestOutboundConnectorResultResultEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
