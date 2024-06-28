// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Capacity Management API
//
// OCI Control Center (OCC) Capacity Management enables you to manage capacity requests in realms where OCI Control Center Capacity Management is available. For more information, see OCI Control Center (https://docs.cloud.oracle.com/iaas/Content/control-center/home.htm).
//

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OccAvailabilityCatalogSummary A catalog containing resource availability details for a customer.
type OccAvailabilityCatalogSummary struct {

	// The OCID of the availability catalog.
	Id *string `mandatory:"true" json:"id"`

	// The name of the OCI service in consideration. For example, Compute, Exadata, and so on.
	Namespace NamespaceEnum `mandatory:"true" json:"namespace"`

	// The OCID of the tenancy where the availability catalog resides.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// An user-friendly name for the availability catalog. Does not have to be unique, and is changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The Customer Group OCID to which the availability catalog belongs.
	OccCustomerGroupId *string `mandatory:"true" json:"occCustomerGroupId"`

	// Represents whether this version of the availability catalog has been made available to the customer. The state is No by default.
	CatalogState OccAvailabilityCatalogCatalogStateEnum `mandatory:"true" json:"catalogState"`

	MetadataDetails *MetadataDetails `mandatory:"true" json:"metadataDetails"`

	// The time when the availability catalog was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when the availability catalog was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current lifecycle state of the customer group.
	LifecycleState OccAvailabilityCatalogLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Text information about the availability catalog.
	Description *string `mandatory:"false" json:"description"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed State.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m OccAvailabilityCatalogSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OccAvailabilityCatalogSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNamespaceEnum(string(m.Namespace)); !ok && m.Namespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Namespace: %s. Supported values are: %s.", m.Namespace, strings.Join(GetNamespaceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccAvailabilityCatalogCatalogStateEnum(string(m.CatalogState)); !ok && m.CatalogState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CatalogState: %s. Supported values are: %s.", m.CatalogState, strings.Join(GetOccAvailabilityCatalogCatalogStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccAvailabilityCatalogLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOccAvailabilityCatalogLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
