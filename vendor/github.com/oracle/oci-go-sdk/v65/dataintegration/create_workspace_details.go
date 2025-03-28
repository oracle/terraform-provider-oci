// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateWorkspaceDetails The information needed to create a new workspace.
type CreateWorkspaceDetails struct {

	// A user-friendly display name for the workspace. Does not have to be unique, and can be modified. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment containing the workspace.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the VCN the subnet is in.
	VcnId *string `mandatory:"false" json:"vcnId"`

	// The OCID of the subnet for customer connected databases.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The IP of the custom DNS.
	DnsServerIp *string `mandatory:"false" json:"dnsServerIp"`

	// The DNS zone of the custom DNS to use to resolve names.
	DnsServerZone *string `mandatory:"false" json:"dnsServerZone"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user defined description for the workspace.
	Description *string `mandatory:"false" json:"description"`

	// Specifies whether the private network connection is enabled or disabled.
	IsPrivateNetworkEnabled *bool `mandatory:"false" json:"isPrivateNetworkEnabled"`

	// DCMS Data Asset Registry ID to which the workspace is associated
	RegistryId *string `mandatory:"false" json:"registryId"`

	// DCMS Private Endpoint ID associated with workspace if the pvt networking is enabled
	EndpointId *string `mandatory:"false" json:"endpointId"`

	// DCMS Data Asset Registry display name
	RegistryName *string `mandatory:"false" json:"registryName"`

	// DCMS Data Asset Registry Compartment Identifier
	RegistryCompartmentId *string `mandatory:"false" json:"registryCompartmentId"`

	// DCMS Private Endpoint Name
	EndpointName *string `mandatory:"false" json:"endpointName"`

	// DCMS PRivate Endpoint Compartment Identifier
	EndpointCompartmentId *string `mandatory:"false" json:"endpointCompartmentId"`

	// Key-values pairs of workspace for storing properties on the workspace.
	WorkspaceProperties map[string]string `mandatory:"false" json:"workspaceProperties"`
}

func (m CreateWorkspaceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateWorkspaceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
