// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FleetDiagnosisSummary Diagnosis of a resource needed by the fleet.
type FleetDiagnosisSummary struct {

	// The type of the resource needed by the fleet.
	// This is the role of a resource in the fleet. Use the OCID to determine the actual OCI
	// resource type such as log group or log.
	ResourceType FleetDiagnosisSummaryResourceTypeEnum `mandatory:"true" json:"resourceType"`

	// The OCID of the external resouce needed by the fleet.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// The state of the resource. The resource state is ACTIVE when it works properly for the fleet.
	// In case it would cause an issue for the fleet function, the state is INACTIVE.
	// When JMS can't locate the resource, the state is NOT_FOUND.
	// OTHER covers other cases, such as a temporarily network issue that prevents JMS from detecting the
	// resource. Check the resourceDiagnosis for details.
	ResourceState FleetDiagnosisSummaryResourceStateEnum `mandatory:"false" json:"resourceState,omitempty"`

	// The diagnosis message.
	ResourceDiagnosis *string `mandatory:"false" json:"resourceDiagnosis"`
}

func (m FleetDiagnosisSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FleetDiagnosisSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFleetDiagnosisSummaryResourceTypeEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetFleetDiagnosisSummaryResourceTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingFleetDiagnosisSummaryResourceStateEnum(string(m.ResourceState)); !ok && m.ResourceState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceState: %s. Supported values are: %s.", m.ResourceState, strings.Join(GetFleetDiagnosisSummaryResourceStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FleetDiagnosisSummaryResourceTypeEnum Enum with underlying type: string
type FleetDiagnosisSummaryResourceTypeEnum string

// Set of constants representing the allowable values for FleetDiagnosisSummaryResourceTypeEnum
const (
	FleetDiagnosisSummaryResourceTypeInventoryLog        FleetDiagnosisSummaryResourceTypeEnum = "INVENTORY_LOG"
	FleetDiagnosisSummaryResourceTypeOperationLog        FleetDiagnosisSummaryResourceTypeEnum = "OPERATION_LOG"
	FleetDiagnosisSummaryResourceTypeCryptoSummarizedLog FleetDiagnosisSummaryResourceTypeEnum = "CRYPTO_SUMMARIZED_LOG"
	FleetDiagnosisSummaryResourceTypeAnalysisOssBucket   FleetDiagnosisSummaryResourceTypeEnum = "ANALYSIS_OSS_BUCKET"
)

var mappingFleetDiagnosisSummaryResourceTypeEnum = map[string]FleetDiagnosisSummaryResourceTypeEnum{
	"INVENTORY_LOG":         FleetDiagnosisSummaryResourceTypeInventoryLog,
	"OPERATION_LOG":         FleetDiagnosisSummaryResourceTypeOperationLog,
	"CRYPTO_SUMMARIZED_LOG": FleetDiagnosisSummaryResourceTypeCryptoSummarizedLog,
	"ANALYSIS_OSS_BUCKET":   FleetDiagnosisSummaryResourceTypeAnalysisOssBucket,
}

var mappingFleetDiagnosisSummaryResourceTypeEnumLowerCase = map[string]FleetDiagnosisSummaryResourceTypeEnum{
	"inventory_log":         FleetDiagnosisSummaryResourceTypeInventoryLog,
	"operation_log":         FleetDiagnosisSummaryResourceTypeOperationLog,
	"crypto_summarized_log": FleetDiagnosisSummaryResourceTypeCryptoSummarizedLog,
	"analysis_oss_bucket":   FleetDiagnosisSummaryResourceTypeAnalysisOssBucket,
}

// GetFleetDiagnosisSummaryResourceTypeEnumValues Enumerates the set of values for FleetDiagnosisSummaryResourceTypeEnum
func GetFleetDiagnosisSummaryResourceTypeEnumValues() []FleetDiagnosisSummaryResourceTypeEnum {
	values := make([]FleetDiagnosisSummaryResourceTypeEnum, 0)
	for _, v := range mappingFleetDiagnosisSummaryResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFleetDiagnosisSummaryResourceTypeEnumStringValues Enumerates the set of values in String for FleetDiagnosisSummaryResourceTypeEnum
func GetFleetDiagnosisSummaryResourceTypeEnumStringValues() []string {
	return []string{
		"INVENTORY_LOG",
		"OPERATION_LOG",
		"CRYPTO_SUMMARIZED_LOG",
		"ANALYSIS_OSS_BUCKET",
	}
}

// GetMappingFleetDiagnosisSummaryResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFleetDiagnosisSummaryResourceTypeEnum(val string) (FleetDiagnosisSummaryResourceTypeEnum, bool) {
	enum, ok := mappingFleetDiagnosisSummaryResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// FleetDiagnosisSummaryResourceStateEnum Enum with underlying type: string
type FleetDiagnosisSummaryResourceStateEnum string

// Set of constants representing the allowable values for FleetDiagnosisSummaryResourceStateEnum
const (
	FleetDiagnosisSummaryResourceStateActive   FleetDiagnosisSummaryResourceStateEnum = "ACTIVE"
	FleetDiagnosisSummaryResourceStateInactive FleetDiagnosisSummaryResourceStateEnum = "INACTIVE"
	FleetDiagnosisSummaryResourceStateNotFound FleetDiagnosisSummaryResourceStateEnum = "NOT_FOUND"
	FleetDiagnosisSummaryResourceStateOther    FleetDiagnosisSummaryResourceStateEnum = "OTHER"
)

var mappingFleetDiagnosisSummaryResourceStateEnum = map[string]FleetDiagnosisSummaryResourceStateEnum{
	"ACTIVE":    FleetDiagnosisSummaryResourceStateActive,
	"INACTIVE":  FleetDiagnosisSummaryResourceStateInactive,
	"NOT_FOUND": FleetDiagnosisSummaryResourceStateNotFound,
	"OTHER":     FleetDiagnosisSummaryResourceStateOther,
}

var mappingFleetDiagnosisSummaryResourceStateEnumLowerCase = map[string]FleetDiagnosisSummaryResourceStateEnum{
	"active":    FleetDiagnosisSummaryResourceStateActive,
	"inactive":  FleetDiagnosisSummaryResourceStateInactive,
	"not_found": FleetDiagnosisSummaryResourceStateNotFound,
	"other":     FleetDiagnosisSummaryResourceStateOther,
}

// GetFleetDiagnosisSummaryResourceStateEnumValues Enumerates the set of values for FleetDiagnosisSummaryResourceStateEnum
func GetFleetDiagnosisSummaryResourceStateEnumValues() []FleetDiagnosisSummaryResourceStateEnum {
	values := make([]FleetDiagnosisSummaryResourceStateEnum, 0)
	for _, v := range mappingFleetDiagnosisSummaryResourceStateEnum {
		values = append(values, v)
	}
	return values
}

// GetFleetDiagnosisSummaryResourceStateEnumStringValues Enumerates the set of values in String for FleetDiagnosisSummaryResourceStateEnum
func GetFleetDiagnosisSummaryResourceStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"NOT_FOUND",
		"OTHER",
	}
}

// GetMappingFleetDiagnosisSummaryResourceStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFleetDiagnosisSummaryResourceStateEnum(val string) (FleetDiagnosisSummaryResourceStateEnum, bool) {
	enum, ok := mappingFleetDiagnosisSummaryResourceStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
