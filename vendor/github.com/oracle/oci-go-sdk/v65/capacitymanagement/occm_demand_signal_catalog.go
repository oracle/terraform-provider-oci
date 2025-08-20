// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Capacity Management API
//
// OCI Control Center (OCC) Capacity Management enables you to manage capacity requests in realms where OCI Control Center Capacity Management is available. For more information, see OCI Control Center (https://docs.oracle.com/iaas/Content/control-center/home.htm).
//

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OccmDemandSignalCatalog A model for the demand signal catalog.
type OccmDemandSignalCatalog struct {

	// The ocid of demand signal catalog.
	Id *string `mandatory:"true" json:"id"`

	// compartment id from where demand signal catalog is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The customer group OCID to which the availability catalog belongs.
	OccCustomerGroupId *string `mandatory:"true" json:"occCustomerGroupId"`

	// displayName of demand signal catalog.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current lifecycle state of the resource.
	LifecycleState OccmDemandSignalCatalogLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time when the demand signal catalog was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when the demand signal catalog was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// description of demand signal catalog.
	Description *string `mandatory:"false" json:"description"`

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

func (m OccmDemandSignalCatalog) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OccmDemandSignalCatalog) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOccmDemandSignalCatalogLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOccmDemandSignalCatalogLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OccmDemandSignalCatalogLifecycleStateEnum Enum with underlying type: string
type OccmDemandSignalCatalogLifecycleStateEnum string

// Set of constants representing the allowable values for OccmDemandSignalCatalogLifecycleStateEnum
const (
	OccmDemandSignalCatalogLifecycleStateCreating OccmDemandSignalCatalogLifecycleStateEnum = "CREATING"
	OccmDemandSignalCatalogLifecycleStateActive   OccmDemandSignalCatalogLifecycleStateEnum = "ACTIVE"
	OccmDemandSignalCatalogLifecycleStateUpdating OccmDemandSignalCatalogLifecycleStateEnum = "UPDATING"
	OccmDemandSignalCatalogLifecycleStateDeleted  OccmDemandSignalCatalogLifecycleStateEnum = "DELETED"
	OccmDemandSignalCatalogLifecycleStateDeleting OccmDemandSignalCatalogLifecycleStateEnum = "DELETING"
	OccmDemandSignalCatalogLifecycleStateFailed   OccmDemandSignalCatalogLifecycleStateEnum = "FAILED"
)

var mappingOccmDemandSignalCatalogLifecycleStateEnum = map[string]OccmDemandSignalCatalogLifecycleStateEnum{
	"CREATING": OccmDemandSignalCatalogLifecycleStateCreating,
	"ACTIVE":   OccmDemandSignalCatalogLifecycleStateActive,
	"UPDATING": OccmDemandSignalCatalogLifecycleStateUpdating,
	"DELETED":  OccmDemandSignalCatalogLifecycleStateDeleted,
	"DELETING": OccmDemandSignalCatalogLifecycleStateDeleting,
	"FAILED":   OccmDemandSignalCatalogLifecycleStateFailed,
}

var mappingOccmDemandSignalCatalogLifecycleStateEnumLowerCase = map[string]OccmDemandSignalCatalogLifecycleStateEnum{
	"creating": OccmDemandSignalCatalogLifecycleStateCreating,
	"active":   OccmDemandSignalCatalogLifecycleStateActive,
	"updating": OccmDemandSignalCatalogLifecycleStateUpdating,
	"deleted":  OccmDemandSignalCatalogLifecycleStateDeleted,
	"deleting": OccmDemandSignalCatalogLifecycleStateDeleting,
	"failed":   OccmDemandSignalCatalogLifecycleStateFailed,
}

// GetOccmDemandSignalCatalogLifecycleStateEnumValues Enumerates the set of values for OccmDemandSignalCatalogLifecycleStateEnum
func GetOccmDemandSignalCatalogLifecycleStateEnumValues() []OccmDemandSignalCatalogLifecycleStateEnum {
	values := make([]OccmDemandSignalCatalogLifecycleStateEnum, 0)
	for _, v := range mappingOccmDemandSignalCatalogLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOccmDemandSignalCatalogLifecycleStateEnumStringValues Enumerates the set of values in String for OccmDemandSignalCatalogLifecycleStateEnum
func GetOccmDemandSignalCatalogLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETED",
		"DELETING",
		"FAILED",
	}
}

// GetMappingOccmDemandSignalCatalogLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccmDemandSignalCatalogLifecycleStateEnum(val string) (OccmDemandSignalCatalogLifecycleStateEnum, bool) {
	enum, ok := mappingOccmDemandSignalCatalogLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
