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

// DrgStaticRouteUpdate Data plane information about Drg Static Routes.
type DrgStaticRouteUpdate struct {

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

	// The label given to the target drg attachment. If the target drg attachment is of
	// type VCN or Internet attachment, then this will be the label of the corresponding
	// VCN or Internet attachment. If the target drg attachment if of type Remote Peering
	// Connection, then this will be the label of the peered RPC attachment.
	DrgAttachmentLabel *int `mandatory:"false" json:"drgAttachmentLabel"`

	// Unique identifier for the Drg Route Table that has the updated Static Route.
	VrfId *string `mandatory:"false" json:"vrfId"`

	// The label given to the Drg Route Table.
	VrfLabel *int `mandatory:"false" json:"vrfLabel"`

	// The route data of the non-VCN attachment
	RouteCidr *string `mandatory:"false" json:"routeCidr"`

	// Next hop label for the static route
	SubstrateNextHopLabel *int `mandatory:"false" json:"substrateNextHopLabel"`

	// The drg “ingress” redirector vip (always IPv4) in the region of the peered RPC attachment
	// in case of RPC attachment. Only applies to RPC & Internet attachments.
	SubstrateNextHopIpAddress *string `mandatory:"false" json:"substrateNextHopIpAddress"`

	// Tells if the static route has been blackholed due to the deletion of an
	// attachment.
	IsBlackHole *bool `mandatory:"false" json:"isBlackHole"`

	// The preference given to the attachment.
	DrgAttachmentPreferenceV2 *int `mandatory:"false" json:"drgAttachmentPreferenceV2"`

	// The type of the attachment.
	DrgAttachmentType DrgStaticRouteUpdateDrgAttachmentTypeEnum `mandatory:"false" json:"drgAttachmentType,omitempty"`
}

// GetId returns Id
func (m DrgStaticRouteUpdate) GetId() *string {
	return m.Id
}

// GetIsDelete returns IsDelete
func (m DrgStaticRouteUpdate) GetIsDelete() *bool {
	return m.IsDelete
}

// GetTimeUpdated returns TimeUpdated
func (m DrgStaticRouteUpdate) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m DrgStaticRouteUpdate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrgStaticRouteUpdate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDrgStaticRouteUpdateDrgAttachmentTypeEnum(string(m.DrgAttachmentType)); !ok && m.DrgAttachmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DrgAttachmentType: %s. Supported values are: %s.", m.DrgAttachmentType, strings.Join(GetDrgStaticRouteUpdateDrgAttachmentTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DrgStaticRouteUpdate) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDrgStaticRouteUpdate DrgStaticRouteUpdate
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDrgStaticRouteUpdate
	}{
		"DrgStaticRouteUpdate",
		(MarshalTypeDrgStaticRouteUpdate)(m),
	}

	return json.Marshal(&s)
}

// DrgStaticRouteUpdateDrgAttachmentTypeEnum Enum with underlying type: string
type DrgStaticRouteUpdateDrgAttachmentTypeEnum string

// Set of constants representing the allowable values for DrgStaticRouteUpdateDrgAttachmentTypeEnum
const (
	DrgStaticRouteUpdateDrgAttachmentTypeVcn                     DrgStaticRouteUpdateDrgAttachmentTypeEnum = "VCN"
	DrgStaticRouteUpdateDrgAttachmentTypeVirtualCircuit          DrgStaticRouteUpdateDrgAttachmentTypeEnum = "VIRTUAL_CIRCUIT"
	DrgStaticRouteUpdateDrgAttachmentTypeRemotePeeringConnection DrgStaticRouteUpdateDrgAttachmentTypeEnum = "REMOTE_PEERING_CONNECTION"
	DrgStaticRouteUpdateDrgAttachmentTypeIpsecTunnel             DrgStaticRouteUpdateDrgAttachmentTypeEnum = "IPSEC_TUNNEL"
	DrgStaticRouteUpdateDrgAttachmentTypeInternet                DrgStaticRouteUpdateDrgAttachmentTypeEnum = "INTERNET"
	DrgStaticRouteUpdateDrgAttachmentTypeInternalOnly            DrgStaticRouteUpdateDrgAttachmentTypeEnum = "INTERNAL_ONLY"
	DrgStaticRouteUpdateDrgAttachmentTypeLoopback                DrgStaticRouteUpdateDrgAttachmentTypeEnum = "LOOPBACK"
)

var mappingDrgStaticRouteUpdateDrgAttachmentTypeEnum = map[string]DrgStaticRouteUpdateDrgAttachmentTypeEnum{
	"VCN":                       DrgStaticRouteUpdateDrgAttachmentTypeVcn,
	"VIRTUAL_CIRCUIT":           DrgStaticRouteUpdateDrgAttachmentTypeVirtualCircuit,
	"REMOTE_PEERING_CONNECTION": DrgStaticRouteUpdateDrgAttachmentTypeRemotePeeringConnection,
	"IPSEC_TUNNEL":              DrgStaticRouteUpdateDrgAttachmentTypeIpsecTunnel,
	"INTERNET":                  DrgStaticRouteUpdateDrgAttachmentTypeInternet,
	"INTERNAL_ONLY":             DrgStaticRouteUpdateDrgAttachmentTypeInternalOnly,
	"LOOPBACK":                  DrgStaticRouteUpdateDrgAttachmentTypeLoopback,
}

var mappingDrgStaticRouteUpdateDrgAttachmentTypeEnumLowerCase = map[string]DrgStaticRouteUpdateDrgAttachmentTypeEnum{
	"vcn":                       DrgStaticRouteUpdateDrgAttachmentTypeVcn,
	"virtual_circuit":           DrgStaticRouteUpdateDrgAttachmentTypeVirtualCircuit,
	"remote_peering_connection": DrgStaticRouteUpdateDrgAttachmentTypeRemotePeeringConnection,
	"ipsec_tunnel":              DrgStaticRouteUpdateDrgAttachmentTypeIpsecTunnel,
	"internet":                  DrgStaticRouteUpdateDrgAttachmentTypeInternet,
	"internal_only":             DrgStaticRouteUpdateDrgAttachmentTypeInternalOnly,
	"loopback":                  DrgStaticRouteUpdateDrgAttachmentTypeLoopback,
}

// GetDrgStaticRouteUpdateDrgAttachmentTypeEnumValues Enumerates the set of values for DrgStaticRouteUpdateDrgAttachmentTypeEnum
func GetDrgStaticRouteUpdateDrgAttachmentTypeEnumValues() []DrgStaticRouteUpdateDrgAttachmentTypeEnum {
	values := make([]DrgStaticRouteUpdateDrgAttachmentTypeEnum, 0)
	for _, v := range mappingDrgStaticRouteUpdateDrgAttachmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrgStaticRouteUpdateDrgAttachmentTypeEnumStringValues Enumerates the set of values in String for DrgStaticRouteUpdateDrgAttachmentTypeEnum
func GetDrgStaticRouteUpdateDrgAttachmentTypeEnumStringValues() []string {
	return []string{
		"VCN",
		"VIRTUAL_CIRCUIT",
		"REMOTE_PEERING_CONNECTION",
		"IPSEC_TUNNEL",
		"INTERNET",
		"INTERNAL_ONLY",
		"LOOPBACK",
	}
}

// GetMappingDrgStaticRouteUpdateDrgAttachmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrgStaticRouteUpdateDrgAttachmentTypeEnum(val string) (DrgStaticRouteUpdateDrgAttachmentTypeEnum, bool) {
	enum, ok := mappingDrgStaticRouteUpdateDrgAttachmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
