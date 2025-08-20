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

// UpdateOccmDemandSignalItemDetails Details about different fields used to update the demand signal item.
type UpdateOccmDemandSignalItemDetails struct {

	// The region for which you want to request the resource for.
	Region *string `mandatory:"false" json:"region"`

	// The name of the availability domain for which you want to request the OCI resource.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The OCID of the tenancy for which you want to request the OCI resource for.
	TargetCompartmentId *string `mandatory:"false" json:"targetCompartmentId"`

	// The quantity of the resource that you want to demand from OCI.
	DemandQuantity *int64 `mandatory:"false" json:"demandQuantity"`

	// the date before which you would ideally like the OCI resource to be delivered to you.
	TimeNeededBefore *common.SDKTime `mandatory:"false" json:"timeNeededBefore"`

	// A map of various properties associated with the OCI resource. This parameter will act as a replace parameter i.e the existing resource properties will be overridden by this update.
	ResourceProperties map[string]string `mandatory:"false" json:"resourceProperties"`

	// This field will serve as notes section for you. You can use this section to convey a message to OCI regarding your resource request.
	// NOTE: The previous value gets overwritten with the new one for this once updated.
	Notes *string `mandatory:"false" json:"notes"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateOccmDemandSignalItemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOccmDemandSignalItemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
