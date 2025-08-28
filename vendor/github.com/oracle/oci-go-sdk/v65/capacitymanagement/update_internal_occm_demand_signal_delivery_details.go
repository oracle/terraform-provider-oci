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

// UpdateInternalOccmDemandSignalDeliveryDetails Details about different fields used to update the demand signal delivery resource.
type UpdateInternalOccmDemandSignalDeliveryDetails struct {

	// The quantity of the resource that OCI will supply to the customer.
	AcceptedQuantity *int64 `mandatory:"false" json:"acceptedQuantity"`

	// The state in which we want to transition the demand signal delivery resource.
	LifecycleDetails InternalOccmDemandSignalDeliveryLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// The date on which the OCI delivered the resource to the customers.
	TimeDelivered *common.SDKTime `mandatory:"false" json:"timeDelivered"`

	// This field could be used by OCI to communicate the reason for declining the request.
	Justification *string `mandatory:"false" json:"justification"`

	// This field acts as a notes section for operators.
	Notes *string `mandatory:"false" json:"notes"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateInternalOccmDemandSignalDeliveryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateInternalOccmDemandSignalDeliveryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInternalOccmDemandSignalDeliveryLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetInternalOccmDemandSignalDeliveryLifecycleDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
