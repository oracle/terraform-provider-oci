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

// OccCapacityRequestSummary A summary model for the capacity request model.
type OccCapacityRequestSummary struct {

	// The OCID of the capacity request.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy from which the request was made.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the availability catalog against which the capacity request was placed.
	OccAvailabilityCatalogId *string `mandatory:"true" json:"occAvailabilityCatalogId"`

	// The display name of the capacity request.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the OCI service in consideration. For example, Compute, Exadata, and so on.
	Namespace NamespaceEnum `mandatory:"true" json:"namespace"`

	// The OCID of the customer group to which this customer belongs to.
	OccCustomerGroupId *string `mandatory:"true" json:"occCustomerGroupId"`

	// The name of the region for which the capacity request was made.
	Region *string `mandatory:"true" json:"region"`

	// The availability domain (AD) for which the capacity request was made.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The date by which the capacity requested by customers before dateFinalCustomerOrder needs to be fulfilled.
	DateExpectedCapacityHandover *common.SDKTime `mandatory:"true" json:"dateExpectedCapacityHandover"`

	// A list of states through which the capacity request goes by.
	RequestState OccCapacityRequestRequestStateEnum `mandatory:"true" json:"requestState"`

	// The time when the capacity request was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when the capacity request was updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current lifecycle state of the customer group.
	LifecycleState OccCapacityRequestLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Meaningful text about the capacity request.
	Description *string `mandatory:"false" json:"description"`

	// Type of Capacity Request(New or Transfer)
	RequestType OccCapacityRequestRequestTypeEnum `mandatory:"false" json:"requestType,omitempty"`

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

func (m OccCapacityRequestSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OccCapacityRequestSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNamespaceEnum(string(m.Namespace)); !ok && m.Namespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Namespace: %s. Supported values are: %s.", m.Namespace, strings.Join(GetNamespaceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccCapacityRequestRequestStateEnum(string(m.RequestState)); !ok && m.RequestState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RequestState: %s. Supported values are: %s.", m.RequestState, strings.Join(GetOccCapacityRequestRequestStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccCapacityRequestLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOccCapacityRequestLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingOccCapacityRequestRequestTypeEnum(string(m.RequestType)); !ok && m.RequestType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RequestType: %s. Supported values are: %s.", m.RequestType, strings.Join(GetOccCapacityRequestRequestTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
