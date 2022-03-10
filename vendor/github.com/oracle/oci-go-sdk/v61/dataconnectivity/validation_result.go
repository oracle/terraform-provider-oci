// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v61/common"
	"strings"
)

// ValidationResult Validation Result object for a single DataAsset.
type ValidationResult struct {

	// Error text for validation failure
	ErrorMsg *string `mandatory:"false" json:"errorMsg"`

	// Status of the validatio result execution
	Status ValidationResultStatusEnum `mandatory:"false" json:"status,omitempty"`
}

func (m ValidationResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ValidationResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingValidationResultStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetValidationResultStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ValidationResultStatusEnum Enum with underlying type: string
type ValidationResultStatusEnum string

// Set of constants representing the allowable values for ValidationResultStatusEnum
const (
	ValidationResultStatusError   ValidationResultStatusEnum = "ERROR"
	ValidationResultStatusSuccess ValidationResultStatusEnum = "SUCCESS"
)

var mappingValidationResultStatusEnum = map[string]ValidationResultStatusEnum{
	"ERROR":   ValidationResultStatusError,
	"SUCCESS": ValidationResultStatusSuccess,
}

var mappingValidationResultStatusEnumLowerCase = map[string]ValidationResultStatusEnum{
	"error":   ValidationResultStatusError,
	"success": ValidationResultStatusSuccess,
}

// GetValidationResultStatusEnumValues Enumerates the set of values for ValidationResultStatusEnum
func GetValidationResultStatusEnumValues() []ValidationResultStatusEnum {
	values := make([]ValidationResultStatusEnum, 0)
	for _, v := range mappingValidationResultStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetValidationResultStatusEnumStringValues Enumerates the set of values in String for ValidationResultStatusEnum
func GetValidationResultStatusEnumStringValues() []string {
	return []string{
		"ERROR",
		"SUCCESS",
	}
}

// GetMappingValidationResultStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingValidationResultStatusEnum(val string) (ValidationResultStatusEnum, bool) {
	enum, ok := mappingValidationResultStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
