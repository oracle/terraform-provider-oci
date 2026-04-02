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

// VanityDomainSummary Summary of the VanityDomain
type VanityDomainSummary struct {

	// The unique identifier (OCID) of the VanityDomain. Can't be changed after creation
	Id *string `mandatory:"true" json:"id"`

	// Vanity domain. Can't be changed after creation
	VanityDomain *string `mandatory:"true" json:"vanityDomain"`

	// The OCID of the Fusion environment that the VanityDomain is created on
	FusionEnvironmentId *string `mandatory:"true" json:"fusionEnvironmentId"`

	// The current lifecycleState of the VanityDomain
	LifecycleState VanityDomainLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The current lifecycleDetails of the VanityDomain
	LifecycleDetails VanityDomainLifecycleDetailsEnum `mandatory:"true" json:"lifecycleDetails"`

	// The time the VanityDomain was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the VanityDomain was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The time the VanityDomain is scheduled to enable. An RFC3339 formatted datetime string
	TimeEnabled *common.SDKTime `mandatory:"true" json:"timeEnabled"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The prefix value of the DnsPrefix. Can't be changed after creation
	Prefix *string `mandatory:"false" json:"prefix"`

	// The ID of the VanityDomainActivity is scheduled
	ScheduledActivityId *string `mandatory:"false" json:"scheduledActivityId"`
}

func (m VanityDomainSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VanityDomainSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVanityDomainLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetVanityDomainLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVanityDomainLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetVanityDomainLifecycleDetailsEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
