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

// ClusterSummary A summary of the Cluster.
type ClusterSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that
	// contains the Cluster.
	Id *string `mandatory:"true" json:"id"`

	// The availability domain that the Cluster's ESXi hosts are running in. For Multi-AD Cluster, it is `multi-AD`.
	ComputeAvailabilityDomain *string `mandatory:"true" json:"computeAvailabilityDomain"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the SDDC that the
	// Cluster belongs to.
	SddcId *string `mandatory:"true" json:"sddcId"`

	// A descriptive name for the Cluster. It must be unique, start with a letter, and contain only letters, digits,
	// whitespaces, dashes and underscores.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// In general, this is a specific version of bundled VMware software supported by
	// Oracle Cloud VMware Solution (see
	// ListSupportedVmwareSoftwareVersions).
	// This attribute is not guaranteed to reflect the version of
	// software currently installed on the ESXi hosts in the Cluster. The purpose
	// of this attribute is to show the version of software that the Oracle
	// Cloud VMware Solution will install on any new ESXi hosts that you *add to this
	// Cluster in the future* with CreateEsxiHost.
	// Therefore, if you upgrade the existing ESXi hosts in the Cluster to use a newer
	// version of bundled VMware software supported by the Oracle Cloud VMware Solution, you
	// should use UpdateCluster to update the Cluster's
	// `vmwareSoftwareVersion` with that new version.
	VmwareSoftwareVersion *string `mandatory:"true" json:"vmwareSoftwareVersion"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that
	// contains the Cluster.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The number of ESXi hosts in the Cluster.
	EsxiHostsCount *int `mandatory:"true" json:"esxiHostsCount"`

	// The initial compute shape of the Cluster's ESXi hosts.
	// ListSupportedHostShapes.
	InitialHostShapeName *string `mandatory:"true" json:"initialHostShapeName"`

	// vSphere Cluster types.
	VsphereType VsphereTypesEnum `mandatory:"true" json:"vsphereType"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The date and time the Cluster was created, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the Cluster was updated, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the Cluster.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Indicates whether shielded instance is enabled at the Cluster level.
	IsShieldedInstanceEnabled *bool `mandatory:"false" json:"isShieldedInstanceEnabled"`

	// The initial OCPU count of the Cluster's ESXi hosts.
	InitialHostOcpuCount *float32 `mandatory:"false" json:"initialHostOcpuCount"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ClusterSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ClusterSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVsphereTypesEnum(string(m.VsphereType)); !ok && m.VsphereType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VsphereType: %s. Supported values are: %s.", m.VsphereType, strings.Join(GetVsphereTypesEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
