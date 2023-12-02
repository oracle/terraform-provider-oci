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

// UpdateSddcDetails The SDDC information to be updated.
// **Important:** Only the `displayName`, `freeFormTags`, and `definedTags` attributes
// affect the existing SDDC. Changing the other attributes affects the `Sddc` object, but not
// the VMware environment currently running on that SDDC. Those other attributes are used
// by the Oracle Cloud VMware Solution *only* for new ESXi hosts that you add to this
// SDDC in the future with CreateEsxiHost.
type UpdateSddcDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the SDDC.
	// SDDC name requirements are 1-16 character length limit, Must start with a letter, Must be English letters, numbers, - only, No repeating hyphens, Must be unique within the region.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The version of bundled VMware software that the Oracle Cloud VMware Solution will
	// install on any new ESXi hosts that you add to this SDDC in the future unless a specific version is configured on the Cluster level.
	// For the list of versions supported by the Oracle Cloud VMware Solution, see
	// ListSupportedVmwareSoftwareVersions).
	VmwareSoftwareVersion *string `mandatory:"false" json:"vmwareSoftwareVersion"`

	// The version of bundled ESXi software that the Oracle Cloud VMware Solution will
	// install on any new ESXi hosts that you add to this SDDC in the future unless a specific version is configured on the Cluster level.
	// For the list of versions supported by the Oracle Cloud VMware Solution, see
	// ListSupportedVmwareSoftwareVersions).
	EsxiSoftwareVersion *string `mandatory:"false" json:"esxiSoftwareVersion"`

	// One or more public SSH keys to be included in the `~/.ssh/authorized_keys` file for
	// the default user on each ESXi host, only when adding new ESXi hosts to this SDDC.
	// Use a newline character to separate multiple keys.
	// The SSH keys must be in the format required for the `authorized_keys` file.
	SshAuthorizedKeys *string `mandatory:"false" json:"sshAuthorizedKeys"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateSddcDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateSddcDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
