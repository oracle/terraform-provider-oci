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

// InternalOccmDemandSignalDeliverySummary An internal summary model containing information about the demand signal delivery resources.
type InternalOccmDemandSignalDeliverySummary struct {

	// The OCID of this demand signal delivery resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy from which the demand signal delivery resource is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the corresponding customer group to which this demand signal delivery resource belongs to.
	OccCustomerGroupId *string `mandatory:"true" json:"occCustomerGroupId"`

	// The OCID of the demand signal under which this delivery will be grouped.
	DemandSignalId *string `mandatory:"true" json:"demandSignalId"`

	// The OCID of the demand signal item corresponding to which this delivery is made.
	DemandSignalItemId *string `mandatory:"true" json:"demandSignalItemId"`

	// The quantity of the resource that OCI will supply to the customer.
	AcceptedQuantity *int64 `mandatory:"true" json:"acceptedQuantity"`

	// The current lifecycle state of the resource.
	LifecycleState InternalOccmDemandSignalDeliveryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The enum values corresponding to the various states associated with the delivery resource.
	// SUBMITTED -> The state where operators have started working and thinking on the quantity that OCI can delivery for the corresponding demand signal item.
	// IN_REVIEW -> The operators are waiting on approvals from different teams/folks in this state.
	// ACCEPTED -> OCI has accepted your resource request and will deliver the quantity as specified by acceptance quantity of this resource.
	// DECLINED -> OCI has declined you resource request.
	// DELIVERED -> OCI has delivered the accepted quantity to the customers.
	// NOTE: The resource becomes visible to customers in ACCEPTED, DECLINED or DELIVERED state.
	LifecycleDetails InternalOccmDemandSignalDeliveryLifecycleDetailsEnum `mandatory:"true" json:"lifecycleDetails"`

	// This field could be used by OCI to communicate the reason for declining the request.
	Justification *string `mandatory:"false" json:"justification"`

	// The date on which the OCI delivered the resource to the customers.
	TimeDelivered *common.SDKTime `mandatory:"false" json:"timeDelivered"`

	// This field acts as a notes section for operators.
	Notes *string `mandatory:"false" json:"notes"`

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

func (m InternalOccmDemandSignalDeliverySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternalOccmDemandSignalDeliverySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInternalOccmDemandSignalDeliveryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetInternalOccmDemandSignalDeliveryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingInternalOccmDemandSignalDeliveryLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetInternalOccmDemandSignalDeliveryLifecycleDetailsEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
