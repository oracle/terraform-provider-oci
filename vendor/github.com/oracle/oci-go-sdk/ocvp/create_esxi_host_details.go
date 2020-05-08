// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use this API to manage the Oracle Cloud VMware Solution.
//

package ocvp

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateEsxiHostDetails Details of the ESXi host to add to the SDDC.
type CreateEsxiHostDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the SDDC to add the
	// ESXi host to.
	SddcId *string `mandatory:"true" json:"sddcId"`

	// A descriptive name for the ESXi host. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// If this attribute is not specified, the SDDC's `instanceDisplayNamePrefix` attribute is used
	// to name and incrementally number the ESXi host. For example, if you're creating the fourth
	// ESXi host in the SDDC, and `instanceDisplayNamePrefix` is `MySDDC`, the host's display
	// name is `MySDDC-4`.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateEsxiHostDetails) String() string {
	return common.PointerString(m)
}
