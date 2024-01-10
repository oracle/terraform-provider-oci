// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FusionEnvironment Description of FusionEnvironment.
type FusionEnvironment struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// FusionEnvironment Identifier, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Type of the FusionEnvironment.
	FusionEnvironmentType FusionEnvironmentFusionEnvironmentTypeEnum `mandatory:"true" json:"fusionEnvironmentType"`

	// The current state of the ServiceInstance.
	LifecycleState FusionEnvironmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	MaintenancePolicy *GetMaintenancePolicyDetails `mandatory:"false" json:"maintenancePolicy"`

	// The next maintenance for this environment
	TimeUpcomingMaintenance *common.SDKTime `mandatory:"false" json:"timeUpcomingMaintenance"`

	// FusionEnvironmentFamily Identifier
	FusionEnvironmentFamilyId *string `mandatory:"false" json:"fusionEnvironmentFamilyId"`

	// List of subscription IDs.
	SubscriptionIds []string `mandatory:"false" json:"subscriptionIds"`

	// BYOK key id
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// BYOK key info
	KmsKeyInfo *interface{} `mandatory:"false" json:"kmsKeyInfo"`

	// The IDCS domain created for the fusion instance
	DomainId *string `mandatory:"false" json:"domainId"`

	// The IDCS Domain URL
	IdcsDomainUrl *string `mandatory:"false" json:"idcsDomainUrl"`

	// Patch bundle names
	AppliedPatchBundles []string `mandatory:"false" json:"appliedPatchBundles"`

	// Version of Fusion Apps used by this environment
	Version *string `mandatory:"false" json:"version"`

	// Public URL
	PublicUrl *string `mandatory:"false" json:"publicUrl"`

	// DNS prefix
	DnsPrefix *string `mandatory:"false" json:"dnsPrefix"`

	// Language packs
	AdditionalLanguagePacks []string `mandatory:"false" json:"additionalLanguagePacks"`

	// The lockbox Id of this fusion environment. If there's no lockbox id, this field will be null
	LockboxId *string `mandatory:"false" json:"lockboxId"`

	// If it's true, then the Break Glass feature is enabled
	IsBreakGlassEnabled *bool `mandatory:"false" json:"isBreakGlassEnabled"`

	Refresh *RefreshDetails `mandatory:"false" json:"refresh"`

	// Network Access Control Rules
	Rules []Rule `mandatory:"false" json:"rules"`

	// The time the the FusionEnvironment was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the FusionEnvironment was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Environment Specific Guid/ System Name
	SystemName *string `mandatory:"false" json:"systemName"`

	EnvironmentRole *EnvironmentRole `mandatory:"false" json:"environmentRole"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m FusionEnvironment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FusionEnvironment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFusionEnvironmentFusionEnvironmentTypeEnum(string(m.FusionEnvironmentType)); !ok && m.FusionEnvironmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FusionEnvironmentType: %s. Supported values are: %s.", m.FusionEnvironmentType, strings.Join(GetFusionEnvironmentFusionEnvironmentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFusionEnvironmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFusionEnvironmentLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *FusionEnvironment) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		MaintenancePolicy         *GetMaintenancePolicyDetails               `json:"maintenancePolicy"`
		TimeUpcomingMaintenance   *common.SDKTime                            `json:"timeUpcomingMaintenance"`
		FusionEnvironmentFamilyId *string                                    `json:"fusionEnvironmentFamilyId"`
		SubscriptionIds           []string                                   `json:"subscriptionIds"`
		KmsKeyId                  *string                                    `json:"kmsKeyId"`
		KmsKeyInfo                *interface{}                               `json:"kmsKeyInfo"`
		DomainId                  *string                                    `json:"domainId"`
		IdcsDomainUrl             *string                                    `json:"idcsDomainUrl"`
		AppliedPatchBundles       []string                                   `json:"appliedPatchBundles"`
		Version                   *string                                    `json:"version"`
		PublicUrl                 *string                                    `json:"publicUrl"`
		DnsPrefix                 *string                                    `json:"dnsPrefix"`
		AdditionalLanguagePacks   []string                                   `json:"additionalLanguagePacks"`
		LockboxId                 *string                                    `json:"lockboxId"`
		IsBreakGlassEnabled       *bool                                      `json:"isBreakGlassEnabled"`
		Refresh                   *RefreshDetails                            `json:"refresh"`
		Rules                     []rule                                     `json:"rules"`
		TimeCreated               *common.SDKTime                            `json:"timeCreated"`
		TimeUpdated               *common.SDKTime                            `json:"timeUpdated"`
		LifecycleDetails          *string                                    `json:"lifecycleDetails"`
		SystemName                *string                                    `json:"systemName"`
		EnvironmentRole           *EnvironmentRole                           `json:"environmentRole"`
		FreeformTags              map[string]string                          `json:"freeformTags"`
		DefinedTags               map[string]map[string]interface{}          `json:"definedTags"`
		Id                        *string                                    `json:"id"`
		DisplayName               *string                                    `json:"displayName"`
		CompartmentId             *string                                    `json:"compartmentId"`
		FusionEnvironmentType     FusionEnvironmentFusionEnvironmentTypeEnum `json:"fusionEnvironmentType"`
		LifecycleState            FusionEnvironmentLifecycleStateEnum        `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.MaintenancePolicy = model.MaintenancePolicy

	m.TimeUpcomingMaintenance = model.TimeUpcomingMaintenance

	m.FusionEnvironmentFamilyId = model.FusionEnvironmentFamilyId

	m.SubscriptionIds = make([]string, len(model.SubscriptionIds))
	copy(m.SubscriptionIds, model.SubscriptionIds)
	m.KmsKeyId = model.KmsKeyId

	m.KmsKeyInfo = model.KmsKeyInfo

	m.DomainId = model.DomainId

	m.IdcsDomainUrl = model.IdcsDomainUrl

	m.AppliedPatchBundles = make([]string, len(model.AppliedPatchBundles))
	copy(m.AppliedPatchBundles, model.AppliedPatchBundles)
	m.Version = model.Version

	m.PublicUrl = model.PublicUrl

	m.DnsPrefix = model.DnsPrefix

	m.AdditionalLanguagePacks = make([]string, len(model.AdditionalLanguagePacks))
	copy(m.AdditionalLanguagePacks, model.AdditionalLanguagePacks)
	m.LockboxId = model.LockboxId

	m.IsBreakGlassEnabled = model.IsBreakGlassEnabled

	m.Refresh = model.Refresh

	m.Rules = make([]Rule, len(model.Rules))
	for i, n := range model.Rules {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Rules[i] = nn.(Rule)
		} else {
			m.Rules[i] = nil
		}
	}
	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleDetails = model.LifecycleDetails

	m.SystemName = model.SystemName

	m.EnvironmentRole = model.EnvironmentRole

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.FusionEnvironmentType = model.FusionEnvironmentType

	m.LifecycleState = model.LifecycleState

	return
}

// FusionEnvironmentFusionEnvironmentTypeEnum Enum with underlying type: string
type FusionEnvironmentFusionEnvironmentTypeEnum string

// Set of constants representing the allowable values for FusionEnvironmentFusionEnvironmentTypeEnum
const (
	FusionEnvironmentFusionEnvironmentTypeProduction  FusionEnvironmentFusionEnvironmentTypeEnum = "PRODUCTION"
	FusionEnvironmentFusionEnvironmentTypeTest        FusionEnvironmentFusionEnvironmentTypeEnum = "TEST"
	FusionEnvironmentFusionEnvironmentTypeDevelopment FusionEnvironmentFusionEnvironmentTypeEnum = "DEVELOPMENT"
)

var mappingFusionEnvironmentFusionEnvironmentTypeEnum = map[string]FusionEnvironmentFusionEnvironmentTypeEnum{
	"PRODUCTION":  FusionEnvironmentFusionEnvironmentTypeProduction,
	"TEST":        FusionEnvironmentFusionEnvironmentTypeTest,
	"DEVELOPMENT": FusionEnvironmentFusionEnvironmentTypeDevelopment,
}

var mappingFusionEnvironmentFusionEnvironmentTypeEnumLowerCase = map[string]FusionEnvironmentFusionEnvironmentTypeEnum{
	"production":  FusionEnvironmentFusionEnvironmentTypeProduction,
	"test":        FusionEnvironmentFusionEnvironmentTypeTest,
	"development": FusionEnvironmentFusionEnvironmentTypeDevelopment,
}

// GetFusionEnvironmentFusionEnvironmentTypeEnumValues Enumerates the set of values for FusionEnvironmentFusionEnvironmentTypeEnum
func GetFusionEnvironmentFusionEnvironmentTypeEnumValues() []FusionEnvironmentFusionEnvironmentTypeEnum {
	values := make([]FusionEnvironmentFusionEnvironmentTypeEnum, 0)
	for _, v := range mappingFusionEnvironmentFusionEnvironmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFusionEnvironmentFusionEnvironmentTypeEnumStringValues Enumerates the set of values in String for FusionEnvironmentFusionEnvironmentTypeEnum
func GetFusionEnvironmentFusionEnvironmentTypeEnumStringValues() []string {
	return []string{
		"PRODUCTION",
		"TEST",
		"DEVELOPMENT",
	}
}

// GetMappingFusionEnvironmentFusionEnvironmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFusionEnvironmentFusionEnvironmentTypeEnum(val string) (FusionEnvironmentFusionEnvironmentTypeEnum, bool) {
	enum, ok := mappingFusionEnvironmentFusionEnvironmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// FusionEnvironmentLifecycleStateEnum Enum with underlying type: string
type FusionEnvironmentLifecycleStateEnum string

// Set of constants representing the allowable values for FusionEnvironmentLifecycleStateEnum
const (
	FusionEnvironmentLifecycleStateCreating FusionEnvironmentLifecycleStateEnum = "CREATING"
	FusionEnvironmentLifecycleStateUpdating FusionEnvironmentLifecycleStateEnum = "UPDATING"
	FusionEnvironmentLifecycleStateActive   FusionEnvironmentLifecycleStateEnum = "ACTIVE"
	FusionEnvironmentLifecycleStateInactive FusionEnvironmentLifecycleStateEnum = "INACTIVE"
	FusionEnvironmentLifecycleStateDeleting FusionEnvironmentLifecycleStateEnum = "DELETING"
	FusionEnvironmentLifecycleStateDeleted  FusionEnvironmentLifecycleStateEnum = "DELETED"
	FusionEnvironmentLifecycleStateFailed   FusionEnvironmentLifecycleStateEnum = "FAILED"
)

var mappingFusionEnvironmentLifecycleStateEnum = map[string]FusionEnvironmentLifecycleStateEnum{
	"CREATING": FusionEnvironmentLifecycleStateCreating,
	"UPDATING": FusionEnvironmentLifecycleStateUpdating,
	"ACTIVE":   FusionEnvironmentLifecycleStateActive,
	"INACTIVE": FusionEnvironmentLifecycleStateInactive,
	"DELETING": FusionEnvironmentLifecycleStateDeleting,
	"DELETED":  FusionEnvironmentLifecycleStateDeleted,
	"FAILED":   FusionEnvironmentLifecycleStateFailed,
}

var mappingFusionEnvironmentLifecycleStateEnumLowerCase = map[string]FusionEnvironmentLifecycleStateEnum{
	"creating": FusionEnvironmentLifecycleStateCreating,
	"updating": FusionEnvironmentLifecycleStateUpdating,
	"active":   FusionEnvironmentLifecycleStateActive,
	"inactive": FusionEnvironmentLifecycleStateInactive,
	"deleting": FusionEnvironmentLifecycleStateDeleting,
	"deleted":  FusionEnvironmentLifecycleStateDeleted,
	"failed":   FusionEnvironmentLifecycleStateFailed,
}

// GetFusionEnvironmentLifecycleStateEnumValues Enumerates the set of values for FusionEnvironmentLifecycleStateEnum
func GetFusionEnvironmentLifecycleStateEnumValues() []FusionEnvironmentLifecycleStateEnum {
	values := make([]FusionEnvironmentLifecycleStateEnum, 0)
	for _, v := range mappingFusionEnvironmentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetFusionEnvironmentLifecycleStateEnumStringValues Enumerates the set of values in String for FusionEnvironmentLifecycleStateEnum
func GetFusionEnvironmentLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingFusionEnvironmentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFusionEnvironmentLifecycleStateEnum(val string) (FusionEnvironmentLifecycleStateEnum, bool) {
	enum, ok := mappingFusionEnvironmentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
