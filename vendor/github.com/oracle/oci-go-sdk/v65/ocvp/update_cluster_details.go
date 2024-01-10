// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateClusterDetails The Cluster information to be updated.
// **Important:** Only the `displayName`, `freeFormTags`, and `definedTags` attributes
// affect the existing Cluster. Changing the other attributes affects the `Cluster` object, but not
// the VMware environment currently running on that Cluster. Those other attributes are used
// by the Oracle Cloud VMware Solution *only* for new ESXi hosts that you add to this
// Cluster in the future with CreateEsxiHost.
type UpdateClusterDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Cluster.
	// Cluster name requirements are 1-16 character length limit, Must start with a letter, Must be English letters, numbers, - only, No repeating hyphens, Must be unique within the region.
	DisplayName *string `mandatory:"false" json:"displayName"`

	NetworkConfiguration *NetworkConfiguration `mandatory:"false" json:"networkConfiguration"`

	// The version of bundled VMware software that the Oracle Cloud VMware Solution will
	// install on any new ESXi hosts that you add to this Cluster in the future. To get a list of the available versions, use
	// ListSupportedVmwareSoftwareVersions.
	VmwareSoftwareVersion *string `mandatory:"false" json:"vmwareSoftwareVersion"`

	// The version of bundled ESXi software that the Oracle Cloud VMware Solution will
	// install on any new ESXi hosts that you add to this Cluster in the future unless a specific version is configured on the ESXi level.
	// To get a list of the available versions, use
	// ListSupportedVmwareSoftwareVersions.
	EsxiSoftwareVersion *string `mandatory:"false" json:"esxiSoftwareVersion"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
