// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateWorkspaceDetails The information about new Workspace.
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

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A detailed description for the workspace.
	Description *string `mandatory:"false" json:"description"`

	// Whether the private network connection is enabled or disabled.
	IsPrivateNetworkEnabled *bool `mandatory:"false" json:"isPrivateNetworkEnabled"`
}

func (m CreateWorkspaceDetails) String() string {
	return common.PointerString(m)
}
