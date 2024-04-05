// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DiscoveryJobLogSummary Log of a specific job
type DiscoveryJobLogSummary struct {

	// The OCID of Discovery job
	Id *string `mandatory:"true" json:"id"`

	// Type of log (INFO, WARNING, ERROR or SUCCESS)
	LogType DiscoveryJobLogSummaryLogTypeEnum `mandatory:"true" json:"logType"`

	// Log message
	LogMessage *string `mandatory:"true" json:"logMessage"`

	// Time the Job log was created
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DiscoveryJobLogSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiscoveryJobLogSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDiscoveryJobLogSummaryLogTypeEnum(string(m.LogType)); !ok && m.LogType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LogType: %s. Supported values are: %s.", m.LogType, strings.Join(GetDiscoveryJobLogSummaryLogTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DiscoveryJobLogSummaryLogTypeEnum Enum with underlying type: string
type DiscoveryJobLogSummaryLogTypeEnum string

// Set of constants representing the allowable values for DiscoveryJobLogSummaryLogTypeEnum
const (
	DiscoveryJobLogSummaryLogTypeInfo    DiscoveryJobLogSummaryLogTypeEnum = "INFO"
	DiscoveryJobLogSummaryLogTypeWarning DiscoveryJobLogSummaryLogTypeEnum = "WARNING"
	DiscoveryJobLogSummaryLogTypeError   DiscoveryJobLogSummaryLogTypeEnum = "ERROR"
	DiscoveryJobLogSummaryLogTypeSuccess DiscoveryJobLogSummaryLogTypeEnum = "SUCCESS"
)

var mappingDiscoveryJobLogSummaryLogTypeEnum = map[string]DiscoveryJobLogSummaryLogTypeEnum{
	"INFO":    DiscoveryJobLogSummaryLogTypeInfo,
	"WARNING": DiscoveryJobLogSummaryLogTypeWarning,
	"ERROR":   DiscoveryJobLogSummaryLogTypeError,
	"SUCCESS": DiscoveryJobLogSummaryLogTypeSuccess,
}

var mappingDiscoveryJobLogSummaryLogTypeEnumLowerCase = map[string]DiscoveryJobLogSummaryLogTypeEnum{
	"info":    DiscoveryJobLogSummaryLogTypeInfo,
	"warning": DiscoveryJobLogSummaryLogTypeWarning,
	"error":   DiscoveryJobLogSummaryLogTypeError,
	"success": DiscoveryJobLogSummaryLogTypeSuccess,
}

// GetDiscoveryJobLogSummaryLogTypeEnumValues Enumerates the set of values for DiscoveryJobLogSummaryLogTypeEnum
func GetDiscoveryJobLogSummaryLogTypeEnumValues() []DiscoveryJobLogSummaryLogTypeEnum {
	values := make([]DiscoveryJobLogSummaryLogTypeEnum, 0)
	for _, v := range mappingDiscoveryJobLogSummaryLogTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveryJobLogSummaryLogTypeEnumStringValues Enumerates the set of values in String for DiscoveryJobLogSummaryLogTypeEnum
func GetDiscoveryJobLogSummaryLogTypeEnumStringValues() []string {
	return []string{
		"INFO",
		"WARNING",
		"ERROR",
		"SUCCESS",
	}
}

// GetMappingDiscoveryJobLogSummaryLogTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveryJobLogSummaryLogTypeEnum(val string) (DiscoveryJobLogSummaryLogTypeEnum, bool) {
	enum, ok := mappingDiscoveryJobLogSummaryLogTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
