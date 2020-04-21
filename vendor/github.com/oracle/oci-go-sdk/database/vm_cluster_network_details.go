// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// VmClusterNetworkDetails Details for a VM cluster network.
type VmClusterNetworkDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the VM cluster network. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The SCAN details.
	Scans []ScanDetails `mandatory:"true" json:"scans"`

	// Details of the client and backup networks.
	VmNetworks []VmNetworkDetails `mandatory:"true" json:"vmNetworks"`

	// The list of DNS server IP addresses. Maximum of 3 allowed.
	Dns []string `mandatory:"false" json:"dns"`

	// The list of NTP server IP addresses. Maximum of 3 allowed.
	Ntp []string `mandatory:"false" json:"ntp"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m VmClusterNetworkDetails) String() string {
	return common.PointerString(m)
}
