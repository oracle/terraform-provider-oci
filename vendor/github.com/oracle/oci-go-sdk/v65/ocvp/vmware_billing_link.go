// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VmwareBillingLink A VMware billing link represents the billing relationship between a VMware customer and OCI tenancy
type VmwareBillingLink struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VMware billing link.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the OCI tenancy linked to the VMware customer.
	TenancyId *string `mandatory:"true" json:"tenancyId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the customer account in VMware service tenancy linked
	// to customer OCI tenancy.
	VmwareAccountId *string `mandatory:"true" json:"vmwareAccountId"`

	// The list of supported shapes that customer can select to provision a SDDC or add an ESXi host to an existing
	// SDDC.
	// ListSupportedHostShapes.
	AllowedShapeNames []string `mandatory:"true" json:"allowedShapeNames"`

	// The current state of the VMware account.
	AccountState VmwareAccountStatesEnum `mandatory:"true" json:"accountState"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment to contain the VMware billing link.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A descriptive name for the VMware billing link. It must be unique, start with a letter, and contain only
	// letters, digits, whitespaces, dashes and underscores.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The list of VMware subscriptions associated with this billing link.
	VmwareSubscriptions []VmwareSubscription `mandatory:"true" json:"vmwareSubscriptions"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The date and time the SDDC was created, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the SDDC was updated, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m VmwareBillingLink) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VmwareBillingLink) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVmwareAccountStatesEnum(string(m.AccountState)); !ok && m.AccountState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccountState: %s. Supported values are: %s.", m.AccountState, strings.Join(GetVmwareAccountStatesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
