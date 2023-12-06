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

// CreateSddcDetails Details of the SDDC.
type CreateSddcDetails struct {

	// The VMware software bundle to install on the ESXi hosts in the SDDC. To get a
	// list of the available versions, use
	// ListSupportedVmwareSoftwareVersions.
	VmwareSoftwareVersion *string `mandatory:"true" json:"vmwareSoftwareVersion"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment to contain the SDDC.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// HCX configuration of the SDDC.
	HcxMode HcxModesEnum `mandatory:"true" json:"hcxMode"`

	InitialConfiguration *InitialConfiguration `mandatory:"true" json:"initialConfiguration"`

	// One or more public SSH keys to be included in the `~/.ssh/authorized_keys` file for
	// the default user on each ESXi host. Use a newline character to separate multiple keys.
	// The SSH keys must be in the format required for the `authorized_keys` file
	SshAuthorizedKeys *string `mandatory:"true" json:"sshAuthorizedKeys"`

	// A descriptive name for the SDDC.
	// SDDC name requirements are 1-16 character length limit, Must start with a letter, Must be English letters, numbers, - only, No repeating hyphens, Must be unique within the region.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The ESXi software bundle to install on the ESXi hosts in the SDDC.
	// Only versions under the same vmwareSoftwareVersion and have been validate by Oracle Cloud VMware Solution will be accepted.
	// To get a list of the available versions, use
	// ListSupportedVmwareSoftwareVersions.
	EsxiSoftwareVersion *string `mandatory:"false" json:"esxiSoftwareVersion"`

	// Indicates whether this SDDC is designated for only single ESXi host.
	IsSingleHostSddc *bool `mandatory:"false" json:"isSingleHostSddc"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateSddcDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSddcDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHcxModesEnum(string(m.HcxMode)); !ok && m.HcxMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HcxMode: %s. Supported values are: %s.", m.HcxMode, strings.Join(GetHcxModesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
