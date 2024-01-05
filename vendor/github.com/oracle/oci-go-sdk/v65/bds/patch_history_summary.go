// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchHistorySummary Patch history of this cluster.
type PatchHistorySummary struct {

	// The version of the patch.
	Version *string `mandatory:"true" json:"version"`

	// The status of this patch.
	LifecycleState PatchHistorySummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time when the patch history was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The type of current patch history.
	// DP - Data Plane patch(This history type is internal available only)
	// ODH - Oracle Distribution of Hadoop patch
	// OS - Operating System patch
	PatchType PatchHistorySummaryPatchTypeEnum `mandatory:"true" json:"patchType"`
}

func (m PatchHistorySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchHistorySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchHistorySummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPatchHistorySummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchHistorySummaryPatchTypeEnum(string(m.PatchType)); !ok && m.PatchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchType: %s. Supported values are: %s.", m.PatchType, strings.Join(GetPatchHistorySummaryPatchTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchHistorySummaryLifecycleStateEnum Enum with underlying type: string
type PatchHistorySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for PatchHistorySummaryLifecycleStateEnum
const (
	PatchHistorySummaryLifecycleStateInstalling PatchHistorySummaryLifecycleStateEnum = "INSTALLING"
	PatchHistorySummaryLifecycleStateInstalled  PatchHistorySummaryLifecycleStateEnum = "INSTALLED"
	PatchHistorySummaryLifecycleStateFailed     PatchHistorySummaryLifecycleStateEnum = "FAILED"
)

var mappingPatchHistorySummaryLifecycleStateEnum = map[string]PatchHistorySummaryLifecycleStateEnum{
	"INSTALLING": PatchHistorySummaryLifecycleStateInstalling,
	"INSTALLED":  PatchHistorySummaryLifecycleStateInstalled,
	"FAILED":     PatchHistorySummaryLifecycleStateFailed,
}

var mappingPatchHistorySummaryLifecycleStateEnumLowerCase = map[string]PatchHistorySummaryLifecycleStateEnum{
	"installing": PatchHistorySummaryLifecycleStateInstalling,
	"installed":  PatchHistorySummaryLifecycleStateInstalled,
	"failed":     PatchHistorySummaryLifecycleStateFailed,
}

// GetPatchHistorySummaryLifecycleStateEnumValues Enumerates the set of values for PatchHistorySummaryLifecycleStateEnum
func GetPatchHistorySummaryLifecycleStateEnumValues() []PatchHistorySummaryLifecycleStateEnum {
	values := make([]PatchHistorySummaryLifecycleStateEnum, 0)
	for _, v := range mappingPatchHistorySummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchHistorySummaryLifecycleStateEnumStringValues Enumerates the set of values in String for PatchHistorySummaryLifecycleStateEnum
func GetPatchHistorySummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"INSTALLING",
		"INSTALLED",
		"FAILED",
	}
}

// GetMappingPatchHistorySummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchHistorySummaryLifecycleStateEnum(val string) (PatchHistorySummaryLifecycleStateEnum, bool) {
	enum, ok := mappingPatchHistorySummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PatchHistorySummaryPatchTypeEnum Enum with underlying type: string
type PatchHistorySummaryPatchTypeEnum string

// Set of constants representing the allowable values for PatchHistorySummaryPatchTypeEnum
const (
	PatchHistorySummaryPatchTypeOdh PatchHistorySummaryPatchTypeEnum = "ODH"
	PatchHistorySummaryPatchTypeOs  PatchHistorySummaryPatchTypeEnum = "OS"
)

var mappingPatchHistorySummaryPatchTypeEnum = map[string]PatchHistorySummaryPatchTypeEnum{
	"ODH": PatchHistorySummaryPatchTypeOdh,
	"OS":  PatchHistorySummaryPatchTypeOs,
}

var mappingPatchHistorySummaryPatchTypeEnumLowerCase = map[string]PatchHistorySummaryPatchTypeEnum{
	"odh": PatchHistorySummaryPatchTypeOdh,
	"os":  PatchHistorySummaryPatchTypeOs,
}

// GetPatchHistorySummaryPatchTypeEnumValues Enumerates the set of values for PatchHistorySummaryPatchTypeEnum
func GetPatchHistorySummaryPatchTypeEnumValues() []PatchHistorySummaryPatchTypeEnum {
	values := make([]PatchHistorySummaryPatchTypeEnum, 0)
	for _, v := range mappingPatchHistorySummaryPatchTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchHistorySummaryPatchTypeEnumStringValues Enumerates the set of values in String for PatchHistorySummaryPatchTypeEnum
func GetPatchHistorySummaryPatchTypeEnumStringValues() []string {
	return []string{
		"ODH",
		"OS",
	}
}

// GetMappingPatchHistorySummaryPatchTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchHistorySummaryPatchTypeEnum(val string) (PatchHistorySummaryPatchTypeEnum, bool) {
	enum, ok := mappingPatchHistorySummaryPatchTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
