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

// FusionEnvironmentFamilySummary Summary information for a Fusion environment family.
type FusionEnvironmentFamilySummary struct {

	// The unique identifier (OCID) of the environment family. Can't be changed after creation.
	Id *string `mandatory:"true" json:"id"`

	// A friendly name for the environment family. The name must contain only letters, numbers, dashes, and underscores. Can be changed later.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment where the environment family is located.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The list of the IDs of the applications subscriptions that are associated with the environment family.
	SubscriptionIds []string `mandatory:"true" json:"subscriptionIds"`

	// The current state of the FusionEnvironmentFamily.
	LifecycleState FusionEnvironmentFamilyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	FamilyMaintenancePolicy *FamilyMaintenancePolicy `mandatory:"false" json:"familyMaintenancePolicy"`

	// When set to True, a subscription update is required for the environment family.
	IsSubscriptionUpdateNeeded *bool `mandatory:"false" json:"isSubscriptionUpdateNeeded"`

	// The time the the FusionEnvironmentFamily was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the FusionEnvironmentFamily was updated. An RFC3339 formatted datetime string.
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

func (m FusionEnvironmentFamilySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FusionEnvironmentFamilySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFusionEnvironmentFamilyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFusionEnvironmentFamilyLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
