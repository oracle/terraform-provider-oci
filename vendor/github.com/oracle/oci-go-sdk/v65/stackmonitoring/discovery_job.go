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

// DiscoveryJob The DiscoveryJob details.
type DiscoveryJob struct {

	// The OCID of Discovery job
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the Compartment
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Add option submits new discovery Job. Add with retry option to re-submit failed discovery job. Refresh option refreshes the existing discovered resources.
	DiscoveryType DiscoveryJobDiscoveryTypeEnum `mandatory:"false" json:"discoveryType,omitempty"`

	// Specifies the status of the discovery job
	Status DiscoveryJobStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The short summary of the status of the discovery job
	StatusMessage *string `mandatory:"false" json:"statusMessage"`

	// The OCID of Tenant
	TenantId *string `mandatory:"false" json:"tenantId"`

	// The OCID of user in which the job is submitted
	UserId *string `mandatory:"false" json:"userId"`

	// Client who submits discovery job.
	DiscoveryClient *string `mandatory:"false" json:"discoveryClient"`

	DiscoveryDetails *DiscoveryDetails `mandatory:"false" json:"discoveryDetails"`

	// The time the discovery Job was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the DiscoveryJob Resource.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DiscoveryJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiscoveryJob) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDiscoveryJobDiscoveryTypeEnum(string(m.DiscoveryType)); !ok && m.DiscoveryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DiscoveryType: %s. Supported values are: %s.", m.DiscoveryType, strings.Join(GetDiscoveryJobDiscoveryTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDiscoveryJobStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDiscoveryJobStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DiscoveryJobDiscoveryTypeEnum Enum with underlying type: string
type DiscoveryJobDiscoveryTypeEnum string

// Set of constants representing the allowable values for DiscoveryJobDiscoveryTypeEnum
const (
	DiscoveryJobDiscoveryTypeAdd          DiscoveryJobDiscoveryTypeEnum = "ADD"
	DiscoveryJobDiscoveryTypeAddWithRetry DiscoveryJobDiscoveryTypeEnum = "ADD_WITH_RETRY"
	DiscoveryJobDiscoveryTypeRefresh      DiscoveryJobDiscoveryTypeEnum = "REFRESH"
)

var mappingDiscoveryJobDiscoveryTypeEnum = map[string]DiscoveryJobDiscoveryTypeEnum{
	"ADD":            DiscoveryJobDiscoveryTypeAdd,
	"ADD_WITH_RETRY": DiscoveryJobDiscoveryTypeAddWithRetry,
	"REFRESH":        DiscoveryJobDiscoveryTypeRefresh,
}

var mappingDiscoveryJobDiscoveryTypeEnumLowerCase = map[string]DiscoveryJobDiscoveryTypeEnum{
	"add":            DiscoveryJobDiscoveryTypeAdd,
	"add_with_retry": DiscoveryJobDiscoveryTypeAddWithRetry,
	"refresh":        DiscoveryJobDiscoveryTypeRefresh,
}

// GetDiscoveryJobDiscoveryTypeEnumValues Enumerates the set of values for DiscoveryJobDiscoveryTypeEnum
func GetDiscoveryJobDiscoveryTypeEnumValues() []DiscoveryJobDiscoveryTypeEnum {
	values := make([]DiscoveryJobDiscoveryTypeEnum, 0)
	for _, v := range mappingDiscoveryJobDiscoveryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveryJobDiscoveryTypeEnumStringValues Enumerates the set of values in String for DiscoveryJobDiscoveryTypeEnum
func GetDiscoveryJobDiscoveryTypeEnumStringValues() []string {
	return []string{
		"ADD",
		"ADD_WITH_RETRY",
		"REFRESH",
	}
}

// GetMappingDiscoveryJobDiscoveryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveryJobDiscoveryTypeEnum(val string) (DiscoveryJobDiscoveryTypeEnum, bool) {
	enum, ok := mappingDiscoveryJobDiscoveryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DiscoveryJobStatusEnum Enum with underlying type: string
type DiscoveryJobStatusEnum string

// Set of constants representing the allowable values for DiscoveryJobStatusEnum
const (
	DiscoveryJobStatusSuccess    DiscoveryJobStatusEnum = "SUCCESS"
	DiscoveryJobStatusFailure    DiscoveryJobStatusEnum = "FAILURE"
	DiscoveryJobStatusInprogress DiscoveryJobStatusEnum = "INPROGRESS"
	DiscoveryJobStatusInactive   DiscoveryJobStatusEnum = "INACTIVE"
	DiscoveryJobStatusCreated    DiscoveryJobStatusEnum = "CREATED"
	DiscoveryJobStatusDeleted    DiscoveryJobStatusEnum = "DELETED"
)

var mappingDiscoveryJobStatusEnum = map[string]DiscoveryJobStatusEnum{
	"SUCCESS":    DiscoveryJobStatusSuccess,
	"FAILURE":    DiscoveryJobStatusFailure,
	"INPROGRESS": DiscoveryJobStatusInprogress,
	"INACTIVE":   DiscoveryJobStatusInactive,
	"CREATED":    DiscoveryJobStatusCreated,
	"DELETED":    DiscoveryJobStatusDeleted,
}

var mappingDiscoveryJobStatusEnumLowerCase = map[string]DiscoveryJobStatusEnum{
	"success":    DiscoveryJobStatusSuccess,
	"failure":    DiscoveryJobStatusFailure,
	"inprogress": DiscoveryJobStatusInprogress,
	"inactive":   DiscoveryJobStatusInactive,
	"created":    DiscoveryJobStatusCreated,
	"deleted":    DiscoveryJobStatusDeleted,
}

// GetDiscoveryJobStatusEnumValues Enumerates the set of values for DiscoveryJobStatusEnum
func GetDiscoveryJobStatusEnumValues() []DiscoveryJobStatusEnum {
	values := make([]DiscoveryJobStatusEnum, 0)
	for _, v := range mappingDiscoveryJobStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveryJobStatusEnumStringValues Enumerates the set of values in String for DiscoveryJobStatusEnum
func GetDiscoveryJobStatusEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"FAILURE",
		"INPROGRESS",
		"INACTIVE",
		"CREATED",
		"DELETED",
	}
}

// GetMappingDiscoveryJobStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveryJobStatusEnum(val string) (DiscoveryJobStatusEnum, bool) {
	enum, ok := mappingDiscoveryJobStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
