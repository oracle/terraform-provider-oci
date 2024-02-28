// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MaskingPolicyHealthReportLogSummary A log entry related to the pre-masking health check.
type MaskingPolicyHealthReportLogSummary struct {

	// The log entry type.
	MessageType MaskingPolicyHealthReportLogSummaryMessageTypeEnum `mandatory:"true" json:"messageType"`

	// The date and time the log entry was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`

	// A human-readable log entry.
	Message *string `mandatory:"true" json:"message"`

	// A human-readable description for the log entry.
	Description *string `mandatory:"true" json:"description"`

	// A human-readable log entry to remedy any error or warnings in the masking policy.
	Remediation *string `mandatory:"false" json:"remediation"`
}

func (m MaskingPolicyHealthReportLogSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaskingPolicyHealthReportLogSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMaskingPolicyHealthReportLogSummaryMessageTypeEnum(string(m.MessageType)); !ok && m.MessageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MessageType: %s. Supported values are: %s.", m.MessageType, strings.Join(GetMaskingPolicyHealthReportLogSummaryMessageTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MaskingPolicyHealthReportLogSummaryMessageTypeEnum Enum with underlying type: string
type MaskingPolicyHealthReportLogSummaryMessageTypeEnum string

// Set of constants representing the allowable values for MaskingPolicyHealthReportLogSummaryMessageTypeEnum
const (
	MaskingPolicyHealthReportLogSummaryMessageTypePass    MaskingPolicyHealthReportLogSummaryMessageTypeEnum = "PASS"
	MaskingPolicyHealthReportLogSummaryMessageTypeWarning MaskingPolicyHealthReportLogSummaryMessageTypeEnum = "WARNING"
	MaskingPolicyHealthReportLogSummaryMessageTypeError   MaskingPolicyHealthReportLogSummaryMessageTypeEnum = "ERROR"
)

var mappingMaskingPolicyHealthReportLogSummaryMessageTypeEnum = map[string]MaskingPolicyHealthReportLogSummaryMessageTypeEnum{
	"PASS":    MaskingPolicyHealthReportLogSummaryMessageTypePass,
	"WARNING": MaskingPolicyHealthReportLogSummaryMessageTypeWarning,
	"ERROR":   MaskingPolicyHealthReportLogSummaryMessageTypeError,
}

var mappingMaskingPolicyHealthReportLogSummaryMessageTypeEnumLowerCase = map[string]MaskingPolicyHealthReportLogSummaryMessageTypeEnum{
	"pass":    MaskingPolicyHealthReportLogSummaryMessageTypePass,
	"warning": MaskingPolicyHealthReportLogSummaryMessageTypeWarning,
	"error":   MaskingPolicyHealthReportLogSummaryMessageTypeError,
}

// GetMaskingPolicyHealthReportLogSummaryMessageTypeEnumValues Enumerates the set of values for MaskingPolicyHealthReportLogSummaryMessageTypeEnum
func GetMaskingPolicyHealthReportLogSummaryMessageTypeEnumValues() []MaskingPolicyHealthReportLogSummaryMessageTypeEnum {
	values := make([]MaskingPolicyHealthReportLogSummaryMessageTypeEnum, 0)
	for _, v := range mappingMaskingPolicyHealthReportLogSummaryMessageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMaskingPolicyHealthReportLogSummaryMessageTypeEnumStringValues Enumerates the set of values in String for MaskingPolicyHealthReportLogSummaryMessageTypeEnum
func GetMaskingPolicyHealthReportLogSummaryMessageTypeEnumStringValues() []string {
	return []string{
		"PASS",
		"WARNING",
		"ERROR",
	}
}

// GetMappingMaskingPolicyHealthReportLogSummaryMessageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaskingPolicyHealthReportLogSummaryMessageTypeEnum(val string) (MaskingPolicyHealthReportLogSummaryMessageTypeEnum, bool) {
	enum, ok := mappingMaskingPolicyHealthReportLogSummaryMessageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
