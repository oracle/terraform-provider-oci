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

// OccmDemandSignalDeliverySummary A summary model containing information about the demand signal delivery resources.
type OccmDemandSignalDeliverySummary struct {

	// The OCID of this demand signal delivery resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy from which the demand signal delivery resource is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the demand signal under which this delivery will be grouped.
	DemandSignalId *string `mandatory:"true" json:"demandSignalId"`

	// The OCID of the demand signal item corresponding to which this delivery is made.
	DemandSignalItemId *string `mandatory:"true" json:"demandSignalItemId"`

	// The quantity of the resource that OCI will supply to the customer.
	AcceptedQuantity *int64 `mandatory:"true" json:"acceptedQuantity"`

	// The current lifecycle state of the resource.
	LifecycleState OccmDemandSignalDeliverySummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The enum values corresponding to the various states associated with the delivery resource.
	// ACCEPTED -> OCI has accepted your resource request and will deliver the quantity as specified by acceptance quantity of this resource.
	// DECLINED -> OCI has declined you resource request.
	// DELIVERED -> OCI has delivered the accepted quantity to the customers.
	// Note: Under extreme rare scenarios the delivery state can toggle between ACCEPTED and DECLINED states
	LifecycleDetails OccmDemandSignalDeliverySummaryLifecycleDetailsEnum `mandatory:"true" json:"lifecycleDetails"`

	// This field could be used by OCI to communicate the reason for accepting or declining the request.
	Justification *string `mandatory:"false" json:"justification"`

	// The date on which the OCI delivered the resource to the customers.
	TimeDelivered *common.SDKTime `mandatory:"false" json:"timeDelivered"`

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

func (m OccmDemandSignalDeliverySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OccmDemandSignalDeliverySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOccmDemandSignalDeliverySummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOccmDemandSignalDeliverySummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccmDemandSignalDeliverySummaryLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetOccmDemandSignalDeliverySummaryLifecycleDetailsEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OccmDemandSignalDeliverySummaryLifecycleStateEnum Enum with underlying type: string
type OccmDemandSignalDeliverySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for OccmDemandSignalDeliverySummaryLifecycleStateEnum
const (
	OccmDemandSignalDeliverySummaryLifecycleStateCreating OccmDemandSignalDeliverySummaryLifecycleStateEnum = "CREATING"
	OccmDemandSignalDeliverySummaryLifecycleStateActive   OccmDemandSignalDeliverySummaryLifecycleStateEnum = "ACTIVE"
	OccmDemandSignalDeliverySummaryLifecycleStateUpdating OccmDemandSignalDeliverySummaryLifecycleStateEnum = "UPDATING"
	OccmDemandSignalDeliverySummaryLifecycleStateDeleted  OccmDemandSignalDeliverySummaryLifecycleStateEnum = "DELETED"
	OccmDemandSignalDeliverySummaryLifecycleStateDeleting OccmDemandSignalDeliverySummaryLifecycleStateEnum = "DELETING"
	OccmDemandSignalDeliverySummaryLifecycleStateFailed   OccmDemandSignalDeliverySummaryLifecycleStateEnum = "FAILED"
)

var mappingOccmDemandSignalDeliverySummaryLifecycleStateEnum = map[string]OccmDemandSignalDeliverySummaryLifecycleStateEnum{
	"CREATING": OccmDemandSignalDeliverySummaryLifecycleStateCreating,
	"ACTIVE":   OccmDemandSignalDeliverySummaryLifecycleStateActive,
	"UPDATING": OccmDemandSignalDeliverySummaryLifecycleStateUpdating,
	"DELETED":  OccmDemandSignalDeliverySummaryLifecycleStateDeleted,
	"DELETING": OccmDemandSignalDeliverySummaryLifecycleStateDeleting,
	"FAILED":   OccmDemandSignalDeliverySummaryLifecycleStateFailed,
}

var mappingOccmDemandSignalDeliverySummaryLifecycleStateEnumLowerCase = map[string]OccmDemandSignalDeliverySummaryLifecycleStateEnum{
	"creating": OccmDemandSignalDeliverySummaryLifecycleStateCreating,
	"active":   OccmDemandSignalDeliverySummaryLifecycleStateActive,
	"updating": OccmDemandSignalDeliverySummaryLifecycleStateUpdating,
	"deleted":  OccmDemandSignalDeliverySummaryLifecycleStateDeleted,
	"deleting": OccmDemandSignalDeliverySummaryLifecycleStateDeleting,
	"failed":   OccmDemandSignalDeliverySummaryLifecycleStateFailed,
}

// GetOccmDemandSignalDeliverySummaryLifecycleStateEnumValues Enumerates the set of values for OccmDemandSignalDeliverySummaryLifecycleStateEnum
func GetOccmDemandSignalDeliverySummaryLifecycleStateEnumValues() []OccmDemandSignalDeliverySummaryLifecycleStateEnum {
	values := make([]OccmDemandSignalDeliverySummaryLifecycleStateEnum, 0)
	for _, v := range mappingOccmDemandSignalDeliverySummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOccmDemandSignalDeliverySummaryLifecycleStateEnumStringValues Enumerates the set of values in String for OccmDemandSignalDeliverySummaryLifecycleStateEnum
func GetOccmDemandSignalDeliverySummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETED",
		"DELETING",
		"FAILED",
	}
}

// GetMappingOccmDemandSignalDeliverySummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccmDemandSignalDeliverySummaryLifecycleStateEnum(val string) (OccmDemandSignalDeliverySummaryLifecycleStateEnum, bool) {
	enum, ok := mappingOccmDemandSignalDeliverySummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OccmDemandSignalDeliverySummaryLifecycleDetailsEnum Enum with underlying type: string
type OccmDemandSignalDeliverySummaryLifecycleDetailsEnum string

// Set of constants representing the allowable values for OccmDemandSignalDeliverySummaryLifecycleDetailsEnum
const (
	OccmDemandSignalDeliverySummaryLifecycleDetailsAccepted  OccmDemandSignalDeliverySummaryLifecycleDetailsEnum = "ACCEPTED"
	OccmDemandSignalDeliverySummaryLifecycleDetailsDeclined  OccmDemandSignalDeliverySummaryLifecycleDetailsEnum = "DECLINED"
	OccmDemandSignalDeliverySummaryLifecycleDetailsDelivered OccmDemandSignalDeliverySummaryLifecycleDetailsEnum = "DELIVERED"
)

var mappingOccmDemandSignalDeliverySummaryLifecycleDetailsEnum = map[string]OccmDemandSignalDeliverySummaryLifecycleDetailsEnum{
	"ACCEPTED":  OccmDemandSignalDeliverySummaryLifecycleDetailsAccepted,
	"DECLINED":  OccmDemandSignalDeliverySummaryLifecycleDetailsDeclined,
	"DELIVERED": OccmDemandSignalDeliverySummaryLifecycleDetailsDelivered,
}

var mappingOccmDemandSignalDeliverySummaryLifecycleDetailsEnumLowerCase = map[string]OccmDemandSignalDeliverySummaryLifecycleDetailsEnum{
	"accepted":  OccmDemandSignalDeliverySummaryLifecycleDetailsAccepted,
	"declined":  OccmDemandSignalDeliverySummaryLifecycleDetailsDeclined,
	"delivered": OccmDemandSignalDeliverySummaryLifecycleDetailsDelivered,
}

// GetOccmDemandSignalDeliverySummaryLifecycleDetailsEnumValues Enumerates the set of values for OccmDemandSignalDeliverySummaryLifecycleDetailsEnum
func GetOccmDemandSignalDeliverySummaryLifecycleDetailsEnumValues() []OccmDemandSignalDeliverySummaryLifecycleDetailsEnum {
	values := make([]OccmDemandSignalDeliverySummaryLifecycleDetailsEnum, 0)
	for _, v := range mappingOccmDemandSignalDeliverySummaryLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetOccmDemandSignalDeliverySummaryLifecycleDetailsEnumStringValues Enumerates the set of values in String for OccmDemandSignalDeliverySummaryLifecycleDetailsEnum
func GetOccmDemandSignalDeliverySummaryLifecycleDetailsEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"DECLINED",
		"DELIVERED",
	}
}

// GetMappingOccmDemandSignalDeliverySummaryLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccmDemandSignalDeliverySummaryLifecycleDetailsEnum(val string) (OccmDemandSignalDeliverySummaryLifecycleDetailsEnum, bool) {
	enum, ok := mappingOccmDemandSignalDeliverySummaryLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
