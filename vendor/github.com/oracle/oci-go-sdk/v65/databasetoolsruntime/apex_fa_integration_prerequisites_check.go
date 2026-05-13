// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ApexFaIntegrationPrerequisitesCheck The results of a prerequisites check for APEX FA integration
type ApexFaIntegrationPrerequisitesCheck struct {

	// Status indicating the outcome of the prerequisites check.
	Status ApexFaIntegrationPrerequisitesCheckStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Messages describing the prerequisites check outcome. Messages can provide actionable information when the status indicates a failure.
	StatusDetails []string `mandatory:"false" json:"statusDetails"`
}

func (m ApexFaIntegrationPrerequisitesCheck) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApexFaIntegrationPrerequisitesCheck) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingApexFaIntegrationPrerequisitesCheckStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetApexFaIntegrationPrerequisitesCheckStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApexFaIntegrationPrerequisitesCheckStatusEnum Enum with underlying type: string
type ApexFaIntegrationPrerequisitesCheckStatusEnum string

// Set of constants representing the allowable values for ApexFaIntegrationPrerequisitesCheckStatusEnum
const (
	ApexFaIntegrationPrerequisitesCheckStatusPass    ApexFaIntegrationPrerequisitesCheckStatusEnum = "PASS"
	ApexFaIntegrationPrerequisitesCheckStatusFail    ApexFaIntegrationPrerequisitesCheckStatusEnum = "FAIL"
	ApexFaIntegrationPrerequisitesCheckStatusError   ApexFaIntegrationPrerequisitesCheckStatusEnum = "ERROR"
	ApexFaIntegrationPrerequisitesCheckStatusUnknown ApexFaIntegrationPrerequisitesCheckStatusEnum = "UNKNOWN"
)

var mappingApexFaIntegrationPrerequisitesCheckStatusEnum = map[string]ApexFaIntegrationPrerequisitesCheckStatusEnum{
	"PASS":    ApexFaIntegrationPrerequisitesCheckStatusPass,
	"FAIL":    ApexFaIntegrationPrerequisitesCheckStatusFail,
	"ERROR":   ApexFaIntegrationPrerequisitesCheckStatusError,
	"UNKNOWN": ApexFaIntegrationPrerequisitesCheckStatusUnknown,
}

var mappingApexFaIntegrationPrerequisitesCheckStatusEnumLowerCase = map[string]ApexFaIntegrationPrerequisitesCheckStatusEnum{
	"pass":    ApexFaIntegrationPrerequisitesCheckStatusPass,
	"fail":    ApexFaIntegrationPrerequisitesCheckStatusFail,
	"error":   ApexFaIntegrationPrerequisitesCheckStatusError,
	"unknown": ApexFaIntegrationPrerequisitesCheckStatusUnknown,
}

// GetApexFaIntegrationPrerequisitesCheckStatusEnumValues Enumerates the set of values for ApexFaIntegrationPrerequisitesCheckStatusEnum
func GetApexFaIntegrationPrerequisitesCheckStatusEnumValues() []ApexFaIntegrationPrerequisitesCheckStatusEnum {
	values := make([]ApexFaIntegrationPrerequisitesCheckStatusEnum, 0)
	for _, v := range mappingApexFaIntegrationPrerequisitesCheckStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetApexFaIntegrationPrerequisitesCheckStatusEnumStringValues Enumerates the set of values in String for ApexFaIntegrationPrerequisitesCheckStatusEnum
func GetApexFaIntegrationPrerequisitesCheckStatusEnumStringValues() []string {
	return []string{
		"PASS",
		"FAIL",
		"ERROR",
		"UNKNOWN",
	}
}

// GetMappingApexFaIntegrationPrerequisitesCheckStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApexFaIntegrationPrerequisitesCheckStatusEnum(val string) (ApexFaIntegrationPrerequisitesCheckStatusEnum, bool) {
	enum, ok := mappingApexFaIntegrationPrerequisitesCheckStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
