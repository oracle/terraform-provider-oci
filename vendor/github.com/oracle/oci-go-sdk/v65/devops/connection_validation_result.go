// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConnectionValidationResult The result of validating the credentials of a connection.
type ConnectionValidationResult struct {

	// The latest result of whether the credentials pass the validation.
	Result ConnectionValidationResultResultEnum `mandatory:"false" json:"result,omitempty"`

	// The latest timestamp when the connection was validated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeValidated *common.SDKTime `mandatory:"false" json:"timeValidated"`

	// A message describing the result of connection validation in more detail.
	Message *string `mandatory:"false" json:"message"`
}

func (m ConnectionValidationResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConnectionValidationResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingConnectionValidationResultResultEnum(string(m.Result)); !ok && m.Result != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Result: %s. Supported values are: %s.", m.Result, strings.Join(GetConnectionValidationResultResultEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConnectionValidationResultResultEnum Enum with underlying type: string
type ConnectionValidationResultResultEnum string

// Set of constants representing the allowable values for ConnectionValidationResultResultEnum
const (
	ConnectionValidationResultResultPass ConnectionValidationResultResultEnum = "PASS"
	ConnectionValidationResultResultFail ConnectionValidationResultResultEnum = "FAIL"
)

var mappingConnectionValidationResultResultEnum = map[string]ConnectionValidationResultResultEnum{
	"PASS": ConnectionValidationResultResultPass,
	"FAIL": ConnectionValidationResultResultFail,
}

var mappingConnectionValidationResultResultEnumLowerCase = map[string]ConnectionValidationResultResultEnum{
	"pass": ConnectionValidationResultResultPass,
	"fail": ConnectionValidationResultResultFail,
}

// GetConnectionValidationResultResultEnumValues Enumerates the set of values for ConnectionValidationResultResultEnum
func GetConnectionValidationResultResultEnumValues() []ConnectionValidationResultResultEnum {
	values := make([]ConnectionValidationResultResultEnum, 0)
	for _, v := range mappingConnectionValidationResultResultEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectionValidationResultResultEnumStringValues Enumerates the set of values in String for ConnectionValidationResultResultEnum
func GetConnectionValidationResultResultEnumStringValues() []string {
	return []string{
		"PASS",
		"FAIL",
	}
}

// GetMappingConnectionValidationResultResultEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectionValidationResultResultEnum(val string) (ConnectionValidationResultResultEnum, bool) {
	enum, ok := mappingConnectionValidationResultResultEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
