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

	// This can be merge to lifecycleState
	// CONNECTED - Partner and CSI information is assigned and MulticloudLink provisioned.
	// DISCONNECTED - Only partner cloud information is assigned.
	// CONNECTING - Oracle Cloud Infrastructure information is assigned and the control plane is provisioning resources.
	// ACTIVE - Network anchor is connected and resources (VNICs) exist within a subnet.
	// ERROR - DRG attach fails during connection.
	// FAILED - Network anchor creation failed
	// NEEDS_ATTENTION - Network anchor is in temporary bad state
	// UPDATING - Network anchor is getting updated.
	// DELETING - Network anchor is getting deleted
	// DELETED - A connected network anchor is deleted.
	NetworkAnchorConnectionStatus OciNetworkMetadataNetworkAnchorConnectionStatusEnum `mandatory:"true" json:"networkAnchorConnectionStatus"`

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
	if _, ok := GetMappingOciNetworkMetadataNetworkAnchorConnectionStatusEnum(string(m.NetworkAnchorConnectionStatus)); !ok && m.NetworkAnchorConnectionStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetworkAnchorConnectionStatus: %s. Supported values are: %s.", m.NetworkAnchorConnectionStatus, strings.Join(GetOciNetworkMetadataNetworkAnchorConnectionStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OciNetworkMetadataNetworkAnchorConnectionStatusEnum Enum with underlying type: string
type OciNetworkMetadataNetworkAnchorConnectionStatusEnum string

// Set of constants representing the allowable values for OciNetworkMetadataNetworkAnchorConnectionStatusEnum
const (
	OciNetworkMetadataNetworkAnchorConnectionStatusConnected      OciNetworkMetadataNetworkAnchorConnectionStatusEnum = "CONNECTED"
	OciNetworkMetadataNetworkAnchorConnectionStatusDisconnected   OciNetworkMetadataNetworkAnchorConnectionStatusEnum = "DISCONNECTED"
	OciNetworkMetadataNetworkAnchorConnectionStatusConnecting     OciNetworkMetadataNetworkAnchorConnectionStatusEnum = "CONNECTING"
	OciNetworkMetadataNetworkAnchorConnectionStatusActive         OciNetworkMetadataNetworkAnchorConnectionStatusEnum = "ACTIVE"
	OciNetworkMetadataNetworkAnchorConnectionStatusError          OciNetworkMetadataNetworkAnchorConnectionStatusEnum = "ERROR"
	OciNetworkMetadataNetworkAnchorConnectionStatusUpdating       OciNetworkMetadataNetworkAnchorConnectionStatusEnum = "UPDATING"
	OciNetworkMetadataNetworkAnchorConnectionStatusNeedsAttention OciNetworkMetadataNetworkAnchorConnectionStatusEnum = "NEEDS_ATTENTION"
	OciNetworkMetadataNetworkAnchorConnectionStatusFailed         OciNetworkMetadataNetworkAnchorConnectionStatusEnum = "FAILED"
	OciNetworkMetadataNetworkAnchorConnectionStatusDeleting       OciNetworkMetadataNetworkAnchorConnectionStatusEnum = "DELETING"
	OciNetworkMetadataNetworkAnchorConnectionStatusDeleted        OciNetworkMetadataNetworkAnchorConnectionStatusEnum = "DELETED"
)

var mappingOciNetworkMetadataNetworkAnchorConnectionStatusEnum = map[string]OciNetworkMetadataNetworkAnchorConnectionStatusEnum{
	"CONNECTED":       OciNetworkMetadataNetworkAnchorConnectionStatusConnected,
	"DISCONNECTED":    OciNetworkMetadataNetworkAnchorConnectionStatusDisconnected,
	"CONNECTING":      OciNetworkMetadataNetworkAnchorConnectionStatusConnecting,
	"ACTIVE":          OciNetworkMetadataNetworkAnchorConnectionStatusActive,
	"ERROR":           OciNetworkMetadataNetworkAnchorConnectionStatusError,
	"UPDATING":        OciNetworkMetadataNetworkAnchorConnectionStatusUpdating,
	"NEEDS_ATTENTION": OciNetworkMetadataNetworkAnchorConnectionStatusNeedsAttention,
	"FAILED":          OciNetworkMetadataNetworkAnchorConnectionStatusFailed,
	"DELETING":        OciNetworkMetadataNetworkAnchorConnectionStatusDeleting,
	"DELETED":         OciNetworkMetadataNetworkAnchorConnectionStatusDeleted,
}

var mappingOciNetworkMetadataNetworkAnchorConnectionStatusEnumLowerCase = map[string]OciNetworkMetadataNetworkAnchorConnectionStatusEnum{
	"connected":       OciNetworkMetadataNetworkAnchorConnectionStatusConnected,
	"disconnected":    OciNetworkMetadataNetworkAnchorConnectionStatusDisconnected,
	"connecting":      OciNetworkMetadataNetworkAnchorConnectionStatusConnecting,
	"active":          OciNetworkMetadataNetworkAnchorConnectionStatusActive,
	"error":           OciNetworkMetadataNetworkAnchorConnectionStatusError,
	"updating":        OciNetworkMetadataNetworkAnchorConnectionStatusUpdating,
	"needs_attention": OciNetworkMetadataNetworkAnchorConnectionStatusNeedsAttention,
	"failed":          OciNetworkMetadataNetworkAnchorConnectionStatusFailed,
	"deleting":        OciNetworkMetadataNetworkAnchorConnectionStatusDeleting,
	"deleted":         OciNetworkMetadataNetworkAnchorConnectionStatusDeleted,
}

// GetOciNetworkMetadataNetworkAnchorConnectionStatusEnumValues Enumerates the set of values for OciNetworkMetadataNetworkAnchorConnectionStatusEnum
func GetOciNetworkMetadataNetworkAnchorConnectionStatusEnumValues() []OciNetworkMetadataNetworkAnchorConnectionStatusEnum {
	values := make([]OciNetworkMetadataNetworkAnchorConnectionStatusEnum, 0)
	for _, v := range mappingOciNetworkMetadataNetworkAnchorConnectionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOciNetworkMetadataNetworkAnchorConnectionStatusEnumStringValues Enumerates the set of values in String for OciNetworkMetadataNetworkAnchorConnectionStatusEnum
func GetOciNetworkMetadataNetworkAnchorConnectionStatusEnumStringValues() []string {
	return []string{
		"CONNECTED",
		"DISCONNECTED",
		"CONNECTING",
		"ACTIVE",
		"ERROR",
		"UPDATING",
		"NEEDS_ATTENTION",
		"FAILED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingOciNetworkMetadataNetworkAnchorConnectionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOciNetworkMetadataNetworkAnchorConnectionStatusEnum(val string) (OciNetworkMetadataNetworkAnchorConnectionStatusEnum, bool) {
	enum, ok := mappingOciNetworkMetadataNetworkAnchorConnectionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
