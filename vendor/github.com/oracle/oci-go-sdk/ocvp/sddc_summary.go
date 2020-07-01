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

// SddcSummary A summary of the SDDC.
type SddcSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that
	// contains the SDDC.
	Id *string `mandatory:"true" json:"id"`

	// The availability domain that the SDDC's ESXi hosts are running in.
	ComputeAvailabilityDomain *string `mandatory:"true" json:"computeAvailabilityDomain"`

	// A descriptive name for the SDDC. It must be unique, start with a letter, and contain only letters, digits,
	// whitespaces, dashes and underscores.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// In general, this is a specific version of bundled VMware software supported by
	// Oracle Cloud VMware Solution (see
	// ListSupportedVmwareSoftwareVersions).
	// This attribute is not guaranteed to reflect the version of
	// software currently installed on the ESXi hosts in the SDDC. The purpose
	// of this attribute is to show the version of software that the Oracle
	// Cloud VMware Solution will install on any new ESXi hosts that you *add to this
	// SDDC in the future* with CreateEsxiHost.
	// Therefore, if you upgrade the existing ESXi hosts in the SDDC to use a newer
	// version of bundled VMware software supported by the Oracle Cloud VMware Solution, you
	// should use UpdateSddc to update the SDDC's
	// `vmwareSoftwareVersion` with that new version.
	VmwareSoftwareVersion *string `mandatory:"true" json:"vmwareSoftwareVersion"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that
	// contains the SDDC.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The number of ESXi hosts in the SDDC.
	EsxiHostsCount *int `mandatory:"true" json:"esxiHostsCount"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// FQDN for vCenter
	// Example: `vcenter-my-sddc.sddc.us-phoenix-1.oraclecloud.com`
	VcenterFqdn *string `mandatory:"false" json:"vcenterFqdn"`

	// FQDN for NSX Manager
	// Example: `nsx-my-sddc.sddc.us-phoenix-1.oraclecloud.com`
	NsxManagerFqdn *string `mandatory:"false" json:"nsxManagerFqdn"`

	// The date and time the SDDC was created, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the SDDC was updated, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the SDDC.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m SddcSummary) String() string {
	return common.PointerString(m)
}
