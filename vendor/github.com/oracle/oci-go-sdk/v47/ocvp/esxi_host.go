// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use this API to manage your Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"github.com/oracle/oci-go-sdk/v47/common"
)

// EsxiHost An ESXi host is a node in an SDDC. At a minimum, each SDDC has 3 ESXi hosts
// that are used to implement a functioning VMware environment.
// In terms of implementation, an ESXi host is a Compute instance that
// is configured with the chosen bundle of VMware software.
// Notice that an `EsxiHost` object has its own OCID (`id`), and a separate
// attribute for the OCID of the Compute instance (`computeInstanceId`).
type EsxiHost struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the ESXi host.
	Id *string `mandatory:"true" json:"id"`

	// A descriptive name for the ESXi host. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the SDDC that the
	// ESXi host belongs to.
	SddcId *string `mandatory:"true" json:"sddcId"`

	// Billing option selected during SDDC creation.
	// Oracle Cloud Infrastructure VMware Solution supports the following billing interval SKUs:
	// HOUR, MONTH, ONE_YEAR, and THREE_YEARS.
	// ListSupportedSkus.
	CurrentSku SkuEnum `mandatory:"true" json:"currentSku"`

	// Billing option to switch to once existing billing cycle ends.
	// ListSupportedSkus.
	NextSku SkuEnum `mandatory:"true" json:"nextSku"`

	// Current billing cycle end date. If nextSku is different from existing SKU, then we switch to newSKu
	// after this contractEndDate
	// Example: `2016-08-25T21:10:29.600Z`
	BillingContractEndDate *common.SDKTime `mandatory:"true" json:"billingContractEndDate"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that
	// contains the SDDC.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// In terms of implementation, an ESXi host is a Compute instance that
	// is configured with the chosen bundle of VMware software. The `computeInstanceId`
	// is the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of that Compute instance.
	ComputeInstanceId *string `mandatory:"false" json:"computeInstanceId"`

	// The date and time the ESXi host was created, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the ESXi host was updated, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the ESXi host.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m EsxiHost) String() string {
	return common.PointerString(m)
}
