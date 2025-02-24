// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrgExportPolicyUpdate Data plane information about Drg Attachment export policy on
// a route table.
type DrgExportPolicyUpdate struct {

	// Unique identifier for the primary resource affected by this update,
	// such as its OCID.
	Id *string `mandatory:"false" json:"id"`

	// True iff this update signals deletion of the identified resource.
	// If true, the type-specific fields of this object may be null.
	IsDelete *bool `mandatory:"false" json:"isDelete"`

	// The date and time that the API call was made that led to this Update.
	// The date and time format is defined by RFC3339.
	// Example: '2016-08-25T21:10:29.600Z'
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The label given to the drg attachment
	DrgAttachmentLabel *int `mandatory:"false" json:"drgAttachmentLabel"`

	// Unique identifier for the Drg Route Table that is affected by this update.
	VrfId *string `mandatory:"false" json:"vrfId"`

	// The label given to the Route Table.
	VrfLabel *int `mandatory:"false" json:"vrfLabel"`

	// The route target label of the peer attachment in case of RPC attachment or route
	// target label of the edge pop virtual circuit.
	ExtendedRouteTargetLabel *int `mandatory:"false" json:"extendedRouteTargetLabel"`

	// The name of the peer region peered through RPC attachment or the name of the region
	// of the edge pop virtual circuit.
	ExtendedRegionName *string `mandatory:"false" json:"extendedRegionName"`

	// The BGP ASN of the peered region in case of RPC attachment or BGP ASN of the region
	// of the edge pop virtual circuit.
	ExtendedAsn *string `mandatory:"false" json:"extendedAsn"`

	// The label of the Export Route Target
	ExportRouteTargetLabel *int `mandatory:"false" json:"exportRouteTargetLabel"`

	// Whether the attachment is an FFAB virtualCircuit.
	IsFFAB *bool `mandatory:"false" json:"isFFAB"`

	// The drg redirector vip (always IPv4) to be used when advertising routes out of the
	// attachment corresponding to this export policy.
	SubstrateNextHopIpAddress *string `mandatory:"false" json:"substrateNextHopIpAddress"`

	// The DP ID of the DRG Attachment's DRG
	DrgDpId *int `mandatory:"false" json:"drgDpId"`

	// Indicates whether the attachment is whitelisted for inter-region transit
	IsWhitelistedInterRegTrans *bool `mandatory:"false" json:"isWhitelistedInterRegTrans"`

	// Disaggregate OSN CIDRs under the DRG VCN attachment
	DoDisaggregate *bool `mandatory:"false" json:"doDisaggregate"`

	// The maximum DrgPathLength which this export policy should advertise.
	// Any route with a longer DrgPathLength should be suppressed.
	// Zero value indicates that there is no configured limit.
	MaxAdvertisedDrgPathLength *int `mandatory:"false" json:"maxAdvertisedDrgPathLength"`

	// The type of the attachment.
	DrgAttachmentType DrgExportPolicyUpdateDrgAttachmentTypeEnum `mandatory:"false" json:"drgAttachmentType,omitempty"`
}

// GetId returns Id
func (m DrgExportPolicyUpdate) GetId() *string {
	return m.Id
}

// GetIsDelete returns IsDelete
func (m DrgExportPolicyUpdate) GetIsDelete() *bool {
	return m.IsDelete
}

// GetTimeUpdated returns TimeUpdated
func (m DrgExportPolicyUpdate) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m DrgExportPolicyUpdate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrgExportPolicyUpdate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDrgExportPolicyUpdateDrgAttachmentTypeEnum(string(m.DrgAttachmentType)); !ok && m.DrgAttachmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DrgAttachmentType: %s. Supported values are: %s.", m.DrgAttachmentType, strings.Join(GetDrgExportPolicyUpdateDrgAttachmentTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DrgExportPolicyUpdate) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDrgExportPolicyUpdate DrgExportPolicyUpdate
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDrgExportPolicyUpdate
	}{
		"DrgExportPolicyUpdate",
		(MarshalTypeDrgExportPolicyUpdate)(m),
	}

	return json.Marshal(&s)
}

// DrgExportPolicyUpdateDrgAttachmentTypeEnum Enum with underlying type: string
type DrgExportPolicyUpdateDrgAttachmentTypeEnum string

// Set of constants representing the allowable values for DrgExportPolicyUpdateDrgAttachmentTypeEnum
const (
	DrgExportPolicyUpdateDrgAttachmentTypeVcn                     DrgExportPolicyUpdateDrgAttachmentTypeEnum = "VCN"
	DrgExportPolicyUpdateDrgAttachmentTypeVirtualCircuit          DrgExportPolicyUpdateDrgAttachmentTypeEnum = "VIRTUAL_CIRCUIT"
	DrgExportPolicyUpdateDrgAttachmentTypeRemotePeeringConnection DrgExportPolicyUpdateDrgAttachmentTypeEnum = "REMOTE_PEERING_CONNECTION"
	DrgExportPolicyUpdateDrgAttachmentTypeIpsecTunnel             DrgExportPolicyUpdateDrgAttachmentTypeEnum = "IPSEC_TUNNEL"
	DrgExportPolicyUpdateDrgAttachmentTypeInternalOnly            DrgExportPolicyUpdateDrgAttachmentTypeEnum = "INTERNAL_ONLY"
	DrgExportPolicyUpdateDrgAttachmentTypeLoopback                DrgExportPolicyUpdateDrgAttachmentTypeEnum = "LOOPBACK"
)

var mappingDrgExportPolicyUpdateDrgAttachmentTypeEnum = map[string]DrgExportPolicyUpdateDrgAttachmentTypeEnum{
	"VCN":                       DrgExportPolicyUpdateDrgAttachmentTypeVcn,
	"VIRTUAL_CIRCUIT":           DrgExportPolicyUpdateDrgAttachmentTypeVirtualCircuit,
	"REMOTE_PEERING_CONNECTION": DrgExportPolicyUpdateDrgAttachmentTypeRemotePeeringConnection,
	"IPSEC_TUNNEL":              DrgExportPolicyUpdateDrgAttachmentTypeIpsecTunnel,
	"INTERNAL_ONLY":             DrgExportPolicyUpdateDrgAttachmentTypeInternalOnly,
	"LOOPBACK":                  DrgExportPolicyUpdateDrgAttachmentTypeLoopback,
}

var mappingDrgExportPolicyUpdateDrgAttachmentTypeEnumLowerCase = map[string]DrgExportPolicyUpdateDrgAttachmentTypeEnum{
	"vcn":                       DrgExportPolicyUpdateDrgAttachmentTypeVcn,
	"virtual_circuit":           DrgExportPolicyUpdateDrgAttachmentTypeVirtualCircuit,
	"remote_peering_connection": DrgExportPolicyUpdateDrgAttachmentTypeRemotePeeringConnection,
	"ipsec_tunnel":              DrgExportPolicyUpdateDrgAttachmentTypeIpsecTunnel,
	"internal_only":             DrgExportPolicyUpdateDrgAttachmentTypeInternalOnly,
	"loopback":                  DrgExportPolicyUpdateDrgAttachmentTypeLoopback,
}

// GetDrgExportPolicyUpdateDrgAttachmentTypeEnumValues Enumerates the set of values for DrgExportPolicyUpdateDrgAttachmentTypeEnum
func GetDrgExportPolicyUpdateDrgAttachmentTypeEnumValues() []DrgExportPolicyUpdateDrgAttachmentTypeEnum {
	values := make([]DrgExportPolicyUpdateDrgAttachmentTypeEnum, 0)
	for _, v := range mappingDrgExportPolicyUpdateDrgAttachmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrgExportPolicyUpdateDrgAttachmentTypeEnumStringValues Enumerates the set of values in String for DrgExportPolicyUpdateDrgAttachmentTypeEnum
func GetDrgExportPolicyUpdateDrgAttachmentTypeEnumStringValues() []string {
	return []string{
		"VCN",
		"VIRTUAL_CIRCUIT",
		"REMOTE_PEERING_CONNECTION",
		"IPSEC_TUNNEL",
		"INTERNAL_ONLY",
		"LOOPBACK",
	}
}

// GetMappingDrgExportPolicyUpdateDrgAttachmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrgExportPolicyUpdateDrgAttachmentTypeEnum(val string) (DrgExportPolicyUpdateDrgAttachmentTypeEnum, bool) {
	enum, ok := mappingDrgExportPolicyUpdateDrgAttachmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
