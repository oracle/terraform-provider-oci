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

// OccmDemandSignalCatalogResource A model containing information about the details of a demand signal catalog resource.
type OccmDemandSignalCatalogResource struct {

	// The OCID of the demand signal catalog resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy from which the demand signal catalog was created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the OCI service in consideration for demand signal submission. For example: COMPUTE, NETWORK, GPU etc.
	Namespace DemandSignalNamespaceEnum `mandatory:"true" json:"namespace"`

	// The name of the OCI resource that you want to request.
	Name *string `mandatory:"true" json:"name"`

	// The current lifecycle state of the resource.
	LifecycleState OccmDemandSignalCatalogResourceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time when the demand signal catalog resource was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when the demand signal catalog resource was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID of the customer tenancy for which this resource will be available for the customer to order against.
	TargetCompartmentId *string `mandatory:"false" json:"targetCompartmentId"`

	// The name of region for which you want to request the OCI resource. This is an optional parameter.
	Region *string `mandatory:"false" json:"region"`

	// The name of the availability domain for which you want to request the OCI resource. This is an optional parameter.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	ResourceProperties *OccmDemandSignalResourcePropertiesCollection `mandatory:"false" json:"resourceProperties"`

	ResourcePropertyConstraints *OccmDemandSignalResourcePropertyConstraintsCollection `mandatory:"false" json:"resourcePropertyConstraints"`

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

func (m OccmDemandSignalCatalogResource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OccmDemandSignalCatalogResource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDemandSignalNamespaceEnum(string(m.Namespace)); !ok && m.Namespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Namespace: %s. Supported values are: %s.", m.Namespace, strings.Join(GetDemandSignalNamespaceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccmDemandSignalCatalogResourceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOccmDemandSignalCatalogResourceLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OccmDemandSignalCatalogResourceLifecycleStateEnum Enum with underlying type: string
type OccmDemandSignalCatalogResourceLifecycleStateEnum string

// Set of constants representing the allowable values for OccmDemandSignalCatalogResourceLifecycleStateEnum
const (
	OccmDemandSignalCatalogResourceLifecycleStateCreating OccmDemandSignalCatalogResourceLifecycleStateEnum = "CREATING"
	OccmDemandSignalCatalogResourceLifecycleStateActive   OccmDemandSignalCatalogResourceLifecycleStateEnum = "ACTIVE"
	OccmDemandSignalCatalogResourceLifecycleStateUpdating OccmDemandSignalCatalogResourceLifecycleStateEnum = "UPDATING"
	OccmDemandSignalCatalogResourceLifecycleStateDeleted  OccmDemandSignalCatalogResourceLifecycleStateEnum = "DELETED"
	OccmDemandSignalCatalogResourceLifecycleStateDeleting OccmDemandSignalCatalogResourceLifecycleStateEnum = "DELETING"
	OccmDemandSignalCatalogResourceLifecycleStateFailed   OccmDemandSignalCatalogResourceLifecycleStateEnum = "FAILED"
)

var mappingOccmDemandSignalCatalogResourceLifecycleStateEnum = map[string]OccmDemandSignalCatalogResourceLifecycleStateEnum{
	"CREATING": OccmDemandSignalCatalogResourceLifecycleStateCreating,
	"ACTIVE":   OccmDemandSignalCatalogResourceLifecycleStateActive,
	"UPDATING": OccmDemandSignalCatalogResourceLifecycleStateUpdating,
	"DELETED":  OccmDemandSignalCatalogResourceLifecycleStateDeleted,
	"DELETING": OccmDemandSignalCatalogResourceLifecycleStateDeleting,
	"FAILED":   OccmDemandSignalCatalogResourceLifecycleStateFailed,
}

var mappingOccmDemandSignalCatalogResourceLifecycleStateEnumLowerCase = map[string]OccmDemandSignalCatalogResourceLifecycleStateEnum{
	"creating": OccmDemandSignalCatalogResourceLifecycleStateCreating,
	"active":   OccmDemandSignalCatalogResourceLifecycleStateActive,
	"updating": OccmDemandSignalCatalogResourceLifecycleStateUpdating,
	"deleted":  OccmDemandSignalCatalogResourceLifecycleStateDeleted,
	"deleting": OccmDemandSignalCatalogResourceLifecycleStateDeleting,
	"failed":   OccmDemandSignalCatalogResourceLifecycleStateFailed,
}

// GetOccmDemandSignalCatalogResourceLifecycleStateEnumValues Enumerates the set of values for OccmDemandSignalCatalogResourceLifecycleStateEnum
func GetOccmDemandSignalCatalogResourceLifecycleStateEnumValues() []OccmDemandSignalCatalogResourceLifecycleStateEnum {
	values := make([]OccmDemandSignalCatalogResourceLifecycleStateEnum, 0)
	for _, v := range mappingOccmDemandSignalCatalogResourceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOccmDemandSignalCatalogResourceLifecycleStateEnumStringValues Enumerates the set of values in String for OccmDemandSignalCatalogResourceLifecycleStateEnum
func GetOccmDemandSignalCatalogResourceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETED",
		"DELETING",
		"FAILED",
	}
}

// GetMappingOccmDemandSignalCatalogResourceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccmDemandSignalCatalogResourceLifecycleStateEnum(val string) (OccmDemandSignalCatalogResourceLifecycleStateEnum, bool) {
	enum, ok := mappingOccmDemandSignalCatalogResourceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
