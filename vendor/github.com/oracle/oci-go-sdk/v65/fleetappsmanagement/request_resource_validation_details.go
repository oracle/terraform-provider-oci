// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RequestResourceValidationDetails Request to initiate resource validation.
type RequestResourceValidationDetails struct {

	// A boolean flag that decides if all resources within the fleet should be part of the validation.
	IsApplicableToAllResources *bool `mandatory:"false" json:"isApplicableToAllResources"`

	// Resource OCIDS to be included for validation.
	ResourceIds []string `mandatory:"false" json:"resourceIds"`

	// A list of resource validation statuses. Each status represents a specific state in the workflow.
	Statuses []RequestResourceValidationDetailsStatusesEnum `mandatory:"false" json:"statuses,omitempty"`
}

func (m RequestResourceValidationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RequestResourceValidationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.Statuses {
		if _, ok := GetMappingRequestResourceValidationDetailsStatusesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Statuses: %s. Supported values are: %s.", val, strings.Join(GetRequestResourceValidationDetailsStatusesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RequestResourceValidationDetailsStatusesEnum Enum with underlying type: string
type RequestResourceValidationDetailsStatusesEnum string

// Set of constants representing the allowable values for RequestResourceValidationDetailsStatusesEnum
const (
	RequestResourceValidationDetailsStatusesPendingValidation RequestResourceValidationDetailsStatusesEnum = "PENDING_VALIDATION"
	RequestResourceValidationDetailsStatusesValidating        RequestResourceValidationDetailsStatusesEnum = "VALIDATING"
	RequestResourceValidationDetailsStatusesValidated         RequestResourceValidationDetailsStatusesEnum = "VALIDATED"
	RequestResourceValidationDetailsStatusesInvalid           RequestResourceValidationDetailsStatusesEnum = "INVALID"
)

var mappingRequestResourceValidationDetailsStatusesEnum = map[string]RequestResourceValidationDetailsStatusesEnum{
	"PENDING_VALIDATION": RequestResourceValidationDetailsStatusesPendingValidation,
	"VALIDATING":         RequestResourceValidationDetailsStatusesValidating,
	"VALIDATED":          RequestResourceValidationDetailsStatusesValidated,
	"INVALID":            RequestResourceValidationDetailsStatusesInvalid,
}

var mappingRequestResourceValidationDetailsStatusesEnumLowerCase = map[string]RequestResourceValidationDetailsStatusesEnum{
	"pending_validation": RequestResourceValidationDetailsStatusesPendingValidation,
	"validating":         RequestResourceValidationDetailsStatusesValidating,
	"validated":          RequestResourceValidationDetailsStatusesValidated,
	"invalid":            RequestResourceValidationDetailsStatusesInvalid,
}

// GetRequestResourceValidationDetailsStatusesEnumValues Enumerates the set of values for RequestResourceValidationDetailsStatusesEnum
func GetRequestResourceValidationDetailsStatusesEnumValues() []RequestResourceValidationDetailsStatusesEnum {
	values := make([]RequestResourceValidationDetailsStatusesEnum, 0)
	for _, v := range mappingRequestResourceValidationDetailsStatusesEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestResourceValidationDetailsStatusesEnumStringValues Enumerates the set of values in String for RequestResourceValidationDetailsStatusesEnum
func GetRequestResourceValidationDetailsStatusesEnumStringValues() []string {
	return []string{
		"PENDING_VALIDATION",
		"VALIDATING",
		"VALIDATED",
		"INVALID",
	}
}

// GetMappingRequestResourceValidationDetailsStatusesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestResourceValidationDetailsStatusesEnum(val string) (RequestResourceValidationDetailsStatusesEnum, bool) {
	enum, ok := mappingRequestResourceValidationDetailsStatusesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
