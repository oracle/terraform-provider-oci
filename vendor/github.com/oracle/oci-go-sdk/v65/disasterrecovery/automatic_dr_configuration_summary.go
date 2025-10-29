// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutomaticDrConfigurationSummary The summary of an Automatic DR configuration.
type AutomaticDrConfigurationSummary struct {

	// The OCID of the Automatic DR configuration.
	// Example: `ocid1.automaticdrconfiguration.oc1..uniqueID`
	Id *string `mandatory:"true" json:"id"`

	// The display name of the Automatic DR configuration.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment containing the Automatic DR configuration.
	// Example: `ocid1.compartment.oc1..uniqueID`
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the Automatic DR configuration was created. An RFC3339 formatted datetime string.
	// Example: `2024-03-29T09:36:42Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the Automatic DR configuration was updated. An RFC3339 formatted datetime string.
	// Example: `2024-03-29T09:36:42Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID of the DR protection group to which this Automatic DR configuration belongs.
	// Example: `ocid1.drprotectiongroup.oc1..uniqueID`
	DrProtectionGroupId *string `mandatory:"true" json:"drProtectionGroupId"`

	// The current state of the Automatic DR configuration.
	LifecycleState AutomaticDrConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The unique id of a Switchover DR Plan.
	// Example: `ocid1.drplan.oc1..uniqueID`
	DefaultSwitchoverDrPlanId *string `mandatory:"false" json:"defaultSwitchoverDrPlanId"`

	// The unique id of a Failover DR Plan.
	// Example: `ocid1.drplan.oc1..uniqueID`
	DefaultFailoverDrPlanId *string `mandatory:"false" json:"defaultFailoverDrPlanId"`

	// A message describing the Automatic DR configuration's current state in more detail.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m AutomaticDrConfigurationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutomaticDrConfigurationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutomaticDrConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutomaticDrConfigurationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
