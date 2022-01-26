// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Database Tools APIs to manage Connections and Private Endpoints.
//

package databasetools

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DatabaseToolsAllowedNetworkSources Allow to restrict access to only requests that come from the specified public or virtual source IP addresses.
type DatabaseToolsAllowedNetworkSources struct {

	// A list of allowed public IPs and CIDR blocks.
	PublicSourceList []string `mandatory:"false" json:"publicSourceList"`

	// A list of allowed VCN OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) and IP ranges pairs.
	// Example:`"vcnId": "ocid1.vcn.oc1.iad.aaaaaaaaexampleuniqueID", "ipRanges": [ "129.213.39.0/24" ]`
	VirtualSourceList []DatabaseToolsVirtualSource `mandatory:"false" json:"virtualSourceList"`
}

func (m DatabaseToolsAllowedNetworkSources) String() string {
	return common.PointerString(m)
}
