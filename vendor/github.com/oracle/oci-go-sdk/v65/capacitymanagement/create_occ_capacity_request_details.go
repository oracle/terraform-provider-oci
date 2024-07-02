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

// CreateOccCapacityRequestDetails Details about the create request for the capacity request.
type CreateOccCapacityRequestDetails struct {

	// Since all resources are at tenancy level hence this will be the ocid of the tenancy where operation is to be performed.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the OCI service in consideration. For example, Compute, Exadata, and so on.
	Namespace NamespaceEnum `mandatory:"true" json:"namespace"`

	// The name of the region for which the capacity request is made.
	Region *string `mandatory:"true" json:"region"`

	// An user-friendly name for the capacity request. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The availability domain (AD) for which the capacity request is made. If this is specified then the capacity will be validated and fulfilled within the scope of this AD.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The date by which the capacity requested by customers before dateFinalCustomerOrder needs to be fulfilled.
	DateExpectedCapacityHandover *common.SDKTime `mandatory:"true" json:"dateExpectedCapacityHandover"`

	// A list of different resources requested by the user.
	Details []OccCapacityRequestBaseDetails `mandatory:"true" json:"details"`

	// The OCID of the availability catalog against which capacity request is made.
	OccAvailabilityCatalogId *string `mandatory:"false" json:"occAvailabilityCatalogId"`

	// Type of Capacity Request(New or Transfer)
	RequestType OccCapacityRequestRequestTypeEnum `mandatory:"false" json:"requestType,omitempty"`

	// Meaningful text about the capacity request.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed State.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The subset of request states available for creating the capacity request.
	RequestState CreateOccCapacityRequestDetailsRequestStateEnum `mandatory:"false" json:"requestState,omitempty"`
}

func (m CreateOccCapacityRequestDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOccCapacityRequestDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNamespaceEnum(string(m.Namespace)); !ok && m.Namespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Namespace: %s. Supported values are: %s.", m.Namespace, strings.Join(GetNamespaceEnumStringValues(), ",")))
	}

	if _, ok := GetMappingOccCapacityRequestRequestTypeEnum(string(m.RequestType)); !ok && m.RequestType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RequestType: %s. Supported values are: %s.", m.RequestType, strings.Join(GetOccCapacityRequestRequestTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateOccCapacityRequestDetailsRequestStateEnum(string(m.RequestState)); !ok && m.RequestState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RequestState: %s. Supported values are: %s.", m.RequestState, strings.Join(GetCreateOccCapacityRequestDetailsRequestStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateOccCapacityRequestDetailsRequestStateEnum Enum with underlying type: string
type CreateOccCapacityRequestDetailsRequestStateEnum string

// Set of constants representing the allowable values for CreateOccCapacityRequestDetailsRequestStateEnum
const (
	CreateOccCapacityRequestDetailsRequestStateCreated   CreateOccCapacityRequestDetailsRequestStateEnum = "CREATED"
	CreateOccCapacityRequestDetailsRequestStateSubmitted CreateOccCapacityRequestDetailsRequestStateEnum = "SUBMITTED"
)

var mappingCreateOccCapacityRequestDetailsRequestStateEnum = map[string]CreateOccCapacityRequestDetailsRequestStateEnum{
	"CREATED":   CreateOccCapacityRequestDetailsRequestStateCreated,
	"SUBMITTED": CreateOccCapacityRequestDetailsRequestStateSubmitted,
}

var mappingCreateOccCapacityRequestDetailsRequestStateEnumLowerCase = map[string]CreateOccCapacityRequestDetailsRequestStateEnum{
	"created":   CreateOccCapacityRequestDetailsRequestStateCreated,
	"submitted": CreateOccCapacityRequestDetailsRequestStateSubmitted,
}

// GetCreateOccCapacityRequestDetailsRequestStateEnumValues Enumerates the set of values for CreateOccCapacityRequestDetailsRequestStateEnum
func GetCreateOccCapacityRequestDetailsRequestStateEnumValues() []CreateOccCapacityRequestDetailsRequestStateEnum {
	values := make([]CreateOccCapacityRequestDetailsRequestStateEnum, 0)
	for _, v := range mappingCreateOccCapacityRequestDetailsRequestStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateOccCapacityRequestDetailsRequestStateEnumStringValues Enumerates the set of values in String for CreateOccCapacityRequestDetailsRequestStateEnum
func GetCreateOccCapacityRequestDetailsRequestStateEnumStringValues() []string {
	return []string{
		"CREATED",
		"SUBMITTED",
	}
}

// GetMappingCreateOccCapacityRequestDetailsRequestStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateOccCapacityRequestDetailsRequestStateEnum(val string) (CreateOccCapacityRequestDetailsRequestStateEnum, bool) {
	enum, ok := mappingCreateOccCapacityRequestDetailsRequestStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
