// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchTaskMessage A patch task message.
type PatchTaskMessage struct {

	// Unique identifier for a message that is system generated
	Key *int64 `mandatory:"false" json:"key"`

	// Retry count of the patch task when this message was generated
	TaskRetryCount *int `mandatory:"false" json:"taskRetryCount"`

	// The time the message was generated. An RFC3339 formatted datetime string
	TimeGenerated *common.SDKTime `mandatory:"false" json:"timeGenerated"`

	// Severity of the message
	Severity PatchTaskMessageSeverityEnum `mandatory:"false" json:"severity,omitempty"`

	// The message
	Message *string `mandatory:"false" json:"message"`
}

func (m PatchTaskMessage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchTaskMessage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPatchTaskMessageSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetPatchTaskMessageSeverityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchTaskMessageSeverityEnum Enum with underlying type: string
type PatchTaskMessageSeverityEnum string

// Set of constants representing the allowable values for PatchTaskMessageSeverityEnum
const (
	PatchTaskMessageSeverityInformation PatchTaskMessageSeverityEnum = "INFORMATION"
	PatchTaskMessageSeverityWarning     PatchTaskMessageSeverityEnum = "WARNING"
	PatchTaskMessageSeverityError       PatchTaskMessageSeverityEnum = "ERROR"
)

var mappingPatchTaskMessageSeverityEnum = map[string]PatchTaskMessageSeverityEnum{
	"INFORMATION": PatchTaskMessageSeverityInformation,
	"WARNING":     PatchTaskMessageSeverityWarning,
	"ERROR":       PatchTaskMessageSeverityError,
}

var mappingPatchTaskMessageSeverityEnumLowerCase = map[string]PatchTaskMessageSeverityEnum{
	"information": PatchTaskMessageSeverityInformation,
	"warning":     PatchTaskMessageSeverityWarning,
	"error":       PatchTaskMessageSeverityError,
}

// GetPatchTaskMessageSeverityEnumValues Enumerates the set of values for PatchTaskMessageSeverityEnum
func GetPatchTaskMessageSeverityEnumValues() []PatchTaskMessageSeverityEnum {
	values := make([]PatchTaskMessageSeverityEnum, 0)
	for _, v := range mappingPatchTaskMessageSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchTaskMessageSeverityEnumStringValues Enumerates the set of values in String for PatchTaskMessageSeverityEnum
func GetPatchTaskMessageSeverityEnumStringValues() []string {
	return []string{
		"INFORMATION",
		"WARNING",
		"ERROR",
	}
}

// GetMappingPatchTaskMessageSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchTaskMessageSeverityEnum(val string) (PatchTaskMessageSeverityEnum, bool) {
	enum, ok := mappingPatchTaskMessageSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
