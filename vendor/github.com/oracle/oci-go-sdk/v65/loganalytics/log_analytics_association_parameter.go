// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogAnalyticsAssociationParameter LogAnalyticsAssociationParameter
type LogAnalyticsAssociationParameter struct {

	// The agent unique identifier.
	AgentId *string `mandatory:"false" json:"agentId"`

	// The entity type.
	EntityType *string `mandatory:"false" json:"entityType"`

	// The entity unique identifier.
	EntityId *string `mandatory:"false" json:"entityId"`

	// The source name.
	SourceId *string `mandatory:"false" json:"sourceId"`

	// The source display name.
	SourceDisplayName *string `mandatory:"false" json:"sourceDisplayName"`

	// The source type.
	SourceType *string `mandatory:"false" json:"sourceType"`

	// The status.  Either FAILED or SUCCEEDED.
	Status LogAnalyticsAssociationParameterStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The status description.
	StatusDescription *string `mandatory:"false" json:"statusDescription"`

	// A list of association properties.
	AssociationProperties []AssociationProperty `mandatory:"false" json:"associationProperties"`

	// A list of missing properties.
	MissingProperties []string `mandatory:"false" json:"missingProperties"`

	// A list of requried properties.
	RequiredProperties []string `mandatory:"false" json:"requiredProperties"`
}

func (m LogAnalyticsAssociationParameter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsAssociationParameter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLogAnalyticsAssociationParameterStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetLogAnalyticsAssociationParameterStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LogAnalyticsAssociationParameterStatusEnum Enum with underlying type: string
type LogAnalyticsAssociationParameterStatusEnum string

// Set of constants representing the allowable values for LogAnalyticsAssociationParameterStatusEnum
const (
	LogAnalyticsAssociationParameterStatusSucceeded LogAnalyticsAssociationParameterStatusEnum = "SUCCEEDED"
	LogAnalyticsAssociationParameterStatusFailed    LogAnalyticsAssociationParameterStatusEnum = "FAILED"
)

var mappingLogAnalyticsAssociationParameterStatusEnum = map[string]LogAnalyticsAssociationParameterStatusEnum{
	"SUCCEEDED": LogAnalyticsAssociationParameterStatusSucceeded,
	"FAILED":    LogAnalyticsAssociationParameterStatusFailed,
}

var mappingLogAnalyticsAssociationParameterStatusEnumLowerCase = map[string]LogAnalyticsAssociationParameterStatusEnum{
	"succeeded": LogAnalyticsAssociationParameterStatusSucceeded,
	"failed":    LogAnalyticsAssociationParameterStatusFailed,
}

// GetLogAnalyticsAssociationParameterStatusEnumValues Enumerates the set of values for LogAnalyticsAssociationParameterStatusEnum
func GetLogAnalyticsAssociationParameterStatusEnumValues() []LogAnalyticsAssociationParameterStatusEnum {
	values := make([]LogAnalyticsAssociationParameterStatusEnum, 0)
	for _, v := range mappingLogAnalyticsAssociationParameterStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetLogAnalyticsAssociationParameterStatusEnumStringValues Enumerates the set of values in String for LogAnalyticsAssociationParameterStatusEnum
func GetLogAnalyticsAssociationParameterStatusEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingLogAnalyticsAssociationParameterStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogAnalyticsAssociationParameterStatusEnum(val string) (LogAnalyticsAssociationParameterStatusEnum, bool) {
	enum, ok := mappingLogAnalyticsAssociationParameterStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
