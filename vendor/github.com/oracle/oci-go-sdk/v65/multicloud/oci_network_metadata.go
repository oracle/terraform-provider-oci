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

// OciNetworkMetadata Oracle Cloud Infrastructure network anchor related meta data items
type OciNetworkMetadata struct {

	// Defines status of the Network Anchor.
	NetworkAnchorConnectionStatus NetworkAnchorConnectionStatusEnum `mandatory:"true" json:"networkAnchorConnectionStatus"`

	Vcn *OciVcn `mandatory:"false" json:"vcn"`

	Dns *OciDns `mandatory:"false" json:"dns"`

	// Network subnets
	Subnets []OciNetworkSubnet `mandatory:"false" json:"subnets"`

	// The DNS Listener Endpoint Address.
	DnsListeningEndpointIpAddress *string `mandatory:"false" json:"dnsListeningEndpointIpAddress"`

	// The DNS Listener Forwarding Address.
	DnsForwardingEndpointIpAddress *string `mandatory:"false" json:"dnsForwardingEndpointIpAddress"`

	// DNS forward configuration
	DnsForwardingConfig []map[string]string `mandatory:"false" json:"dnsForwardingConfig"`
}

func (m OciNetworkMetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciNetworkMetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNetworkAnchorConnectionStatusEnum(string(m.NetworkAnchorConnectionStatus)); !ok && m.NetworkAnchorConnectionStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetworkAnchorConnectionStatus: %s. Supported values are: %s.", m.NetworkAnchorConnectionStatus, strings.Join(GetNetworkAnchorConnectionStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
