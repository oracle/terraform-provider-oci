// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VanityDomainActivity Vanity Domain Activity resource
type VanityDomainActivity struct {

	// The unique identifier (OCID) of the VanityDomainActivity. Can't be changed after creation
	Id *string `mandatory:"true" json:"id"`

	// Vanity domain ID
	VanityDomainId *string `mandatory:"true" json:"vanityDomainId"`

	// The OCID of the Fusion environment that the VanityDomainActivity is created on
	FusionEnvironmentId *string `mandatory:"true" json:"fusionEnvironmentId"`

	// The current lifecycleState of the VanityDomainActivity
	LifecycleState VanityDomainActivityLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The current lifecycleDetails of the VanityDomainActivity
	LifecycleDetails VanityDomainActivityLifecycleDetailsEnum `mandatory:"true" json:"lifecycleDetails"`

	// The time the VanityDomainactivity was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the VanityDomainactivity was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The time the VanityDomainactivity is scheduled to enable. An RFC3339 formatted datetime string
	TimeEnabled *common.SDKTime `mandatory:"true" json:"timeEnabled"`

	// The time the VanityDomainactivity is scheduled to deactivate. An RFC3339 formatted datetime string
	TimeDeactivated *common.SDKTime `mandatory:"true" json:"timeDeactivated"`

	// Activity start time
	TimeScheduled *common.SDKTime `mandatory:"true" json:"timeScheduled"`

	// The type of operation
	OperationType VanityDomainActivityOperationTypeEnum `mandatory:"true" json:"operationType"`
}

func (m VanityDomainActivity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VanityDomainActivity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVanityDomainActivityLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetVanityDomainActivityLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVanityDomainActivityLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetVanityDomainActivityLifecycleDetailsEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVanityDomainActivityOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetVanityDomainActivityOperationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VanityDomainActivityLifecycleStateEnum Enum with underlying type: string
type VanityDomainActivityLifecycleStateEnum string

// Set of constants representing the allowable values for VanityDomainActivityLifecycleStateEnum
const (
	VanityDomainActivityLifecycleStateAccepted       VanityDomainActivityLifecycleStateEnum = "ACCEPTED"
	VanityDomainActivityLifecycleStateInProgress     VanityDomainActivityLifecycleStateEnum = "IN_PROGRESS"
	VanityDomainActivityLifecycleStateFailed         VanityDomainActivityLifecycleStateEnum = "FAILED"
	VanityDomainActivityLifecycleStateSucceeded      VanityDomainActivityLifecycleStateEnum = "SUCCEEDED"
	VanityDomainActivityLifecycleStateCanceled       VanityDomainActivityLifecycleStateEnum = "CANCELED"
	VanityDomainActivityLifecycleStateNeedsAttention VanityDomainActivityLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingVanityDomainActivityLifecycleStateEnum = map[string]VanityDomainActivityLifecycleStateEnum{
	"ACCEPTED":        VanityDomainActivityLifecycleStateAccepted,
	"IN_PROGRESS":     VanityDomainActivityLifecycleStateInProgress,
	"FAILED":          VanityDomainActivityLifecycleStateFailed,
	"SUCCEEDED":       VanityDomainActivityLifecycleStateSucceeded,
	"CANCELED":        VanityDomainActivityLifecycleStateCanceled,
	"NEEDS_ATTENTION": VanityDomainActivityLifecycleStateNeedsAttention,
}

var mappingVanityDomainActivityLifecycleStateEnumLowerCase = map[string]VanityDomainActivityLifecycleStateEnum{
	"accepted":        VanityDomainActivityLifecycleStateAccepted,
	"in_progress":     VanityDomainActivityLifecycleStateInProgress,
	"failed":          VanityDomainActivityLifecycleStateFailed,
	"succeeded":       VanityDomainActivityLifecycleStateSucceeded,
	"canceled":        VanityDomainActivityLifecycleStateCanceled,
	"needs_attention": VanityDomainActivityLifecycleStateNeedsAttention,
}

// GetVanityDomainActivityLifecycleStateEnumValues Enumerates the set of values for VanityDomainActivityLifecycleStateEnum
func GetVanityDomainActivityLifecycleStateEnumValues() []VanityDomainActivityLifecycleStateEnum {
	values := make([]VanityDomainActivityLifecycleStateEnum, 0)
	for _, v := range mappingVanityDomainActivityLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetVanityDomainActivityLifecycleStateEnumStringValues Enumerates the set of values in String for VanityDomainActivityLifecycleStateEnum
func GetVanityDomainActivityLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingVanityDomainActivityLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVanityDomainActivityLifecycleStateEnum(val string) (VanityDomainActivityLifecycleStateEnum, bool) {
	enum, ok := mappingVanityDomainActivityLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VanityDomainActivityLifecycleDetailsEnum Enum with underlying type: string
type VanityDomainActivityLifecycleDetailsEnum string

// Set of constants representing the allowable values for VanityDomainActivityLifecycleDetailsEnum
const (
	VanityDomainActivityLifecycleDetailsNone VanityDomainActivityLifecycleDetailsEnum = "NONE"
)

var mappingVanityDomainActivityLifecycleDetailsEnum = map[string]VanityDomainActivityLifecycleDetailsEnum{
	"NONE": VanityDomainActivityLifecycleDetailsNone,
}

var mappingVanityDomainActivityLifecycleDetailsEnumLowerCase = map[string]VanityDomainActivityLifecycleDetailsEnum{
	"none": VanityDomainActivityLifecycleDetailsNone,
}

// GetVanityDomainActivityLifecycleDetailsEnumValues Enumerates the set of values for VanityDomainActivityLifecycleDetailsEnum
func GetVanityDomainActivityLifecycleDetailsEnumValues() []VanityDomainActivityLifecycleDetailsEnum {
	values := make([]VanityDomainActivityLifecycleDetailsEnum, 0)
	for _, v := range mappingVanityDomainActivityLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetVanityDomainActivityLifecycleDetailsEnumStringValues Enumerates the set of values in String for VanityDomainActivityLifecycleDetailsEnum
func GetVanityDomainActivityLifecycleDetailsEnumStringValues() []string {
	return []string{
		"NONE",
	}
}

// GetMappingVanityDomainActivityLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVanityDomainActivityLifecycleDetailsEnum(val string) (VanityDomainActivityLifecycleDetailsEnum, bool) {
	enum, ok := mappingVanityDomainActivityLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VanityDomainActivityOperationTypeEnum Enum with underlying type: string
type VanityDomainActivityOperationTypeEnum string

// Set of constants representing the allowable values for VanityDomainActivityOperationTypeEnum
const (
	VanityDomainActivityOperationTypeEnableVanityDomain VanityDomainActivityOperationTypeEnum = "ENABLE_VANITY_DOMAIN"
	VanityDomainActivityOperationTypeDeleteVanityDomain VanityDomainActivityOperationTypeEnum = "DELETE_VANITY_DOMAIN"
)

var mappingVanityDomainActivityOperationTypeEnum = map[string]VanityDomainActivityOperationTypeEnum{
	"ENABLE_VANITY_DOMAIN": VanityDomainActivityOperationTypeEnableVanityDomain,
	"DELETE_VANITY_DOMAIN": VanityDomainActivityOperationTypeDeleteVanityDomain,
}

var mappingVanityDomainActivityOperationTypeEnumLowerCase = map[string]VanityDomainActivityOperationTypeEnum{
	"enable_vanity_domain": VanityDomainActivityOperationTypeEnableVanityDomain,
	"delete_vanity_domain": VanityDomainActivityOperationTypeDeleteVanityDomain,
}

// GetVanityDomainActivityOperationTypeEnumValues Enumerates the set of values for VanityDomainActivityOperationTypeEnum
func GetVanityDomainActivityOperationTypeEnumValues() []VanityDomainActivityOperationTypeEnum {
	values := make([]VanityDomainActivityOperationTypeEnum, 0)
	for _, v := range mappingVanityDomainActivityOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetVanityDomainActivityOperationTypeEnumStringValues Enumerates the set of values in String for VanityDomainActivityOperationTypeEnum
func GetVanityDomainActivityOperationTypeEnumStringValues() []string {
	return []string{
		"ENABLE_VANITY_DOMAIN",
		"DELETE_VANITY_DOMAIN",
	}
}

// GetMappingVanityDomainActivityOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVanityDomainActivityOperationTypeEnum(val string) (VanityDomainActivityOperationTypeEnum, bool) {
	enum, ok := mappingVanityDomainActivityOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
