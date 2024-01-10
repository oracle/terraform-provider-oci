// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FusionEnvironmentSummary Summary of the internal FA Environment.
type FusionEnvironmentSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// FusionEnvironment Identifier, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Type of the FusionEnvironment.
	FusionEnvironmentType FusionEnvironmentFusionEnvironmentTypeEnum `mandatory:"true" json:"fusionEnvironmentType"`

	// The current state of the FusionEnvironment.
	LifecycleState FusionEnvironmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The next maintenance for this environment
	TimeUpcomingMaintenance *common.SDKTime `mandatory:"false" json:"timeUpcomingMaintenance"`

	MaintenancePolicy *GetMaintenancePolicyDetails `mandatory:"false" json:"maintenancePolicy"`

	// FusionEnvironmentFamily Identifier
	FusionEnvironmentFamilyId *string `mandatory:"false" json:"fusionEnvironmentFamilyId"`

	// List of subscription IDs.
	SubscriptionIds []string `mandatory:"false" json:"subscriptionIds"`

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

	// The time the the FusionEnvironment was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the FusionEnvironment was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m FusionEnvironmentSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FusionEnvironmentSummary) ValidateEnumValue() (bool, error) {
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
