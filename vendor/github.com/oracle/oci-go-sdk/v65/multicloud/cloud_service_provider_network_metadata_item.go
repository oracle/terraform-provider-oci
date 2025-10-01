// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Multicloud API
//
// Use the Oracle Multicloud API to retrieve resource anchors and network anchors, and the metadata mappings related a Cloud Service Provider. For more information, see <link to docs>.
//

package multicloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CloudServiceProviderNetworkMetadataItem Cloud Service Provider metadata item.
// Warning - In future this object can change to generic object with future Cloud Service Provider based on
// CloudServiceProvider field. This can be one of CSP provider type Azure, GCP and AWS
type CloudServiceProviderNetworkMetadataItem struct {

	// Azure/GCP/AWS region
	Region *string `mandatory:"true" json:"region"`

	// CSP oracle database network anchor unique ID/name
	OdbNetworkId *string `mandatory:"true" json:"odbNetworkId"`

	// An Azure/GCP/AWS cidrBlocks
	CidrBlocks []string `mandatory:"false" json:"cidrBlocks"`

	// CSP network anchor Uri
	NetworkAnchorUri *string `mandatory:"false" json:"networkAnchorUri"`

	// DNS domain ip mapping forwarding configuration
	DnsForwardingConfig []map[string]string `mandatory:"false" json:"dnsForwardingConfig"`
}

func (m CloudServiceProviderNetworkMetadataItem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudServiceProviderNetworkMetadataItem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
