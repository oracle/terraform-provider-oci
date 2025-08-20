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

// InternalOccmDemandSignalCatalogResourceSummary A summary model containing information about the details of a demand signal catalog resource.
type InternalOccmDemandSignalCatalogResourceSummary struct {

	// The OCID of the demand signal catalog resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy from which the request to create the demand signal catalog was made.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the customerGroup.
	OccCustomerGroupId *string `mandatory:"true" json:"occCustomerGroupId"`

	// This OCID of the demand signal catalog
	OccmDemandSignalCatalogId *string `mandatory:"true" json:"occmDemandSignalCatalogId"`

	// The name of the OCI service in consideration for demand signal submission. For example: COMPUTE, NETWORK, GPU etc.
	Namespace DemandSignalNamespaceEnum `mandatory:"true" json:"namespace"`

	// The name of the OCI resource that you want to request.
	Name *string `mandatory:"true" json:"name"`

	// The current lifecycle state of the demand signal catalog resource.
	LifecycleState InternalOccmDemandSignalCatalogResourceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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

func (m InternalOccmDemandSignalCatalogResourceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternalOccmDemandSignalCatalogResourceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDemandSignalNamespaceEnum(string(m.Namespace)); !ok && m.Namespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Namespace: %s. Supported values are: %s.", m.Namespace, strings.Join(GetDemandSignalNamespaceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingInternalOccmDemandSignalCatalogResourceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetInternalOccmDemandSignalCatalogResourceLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
