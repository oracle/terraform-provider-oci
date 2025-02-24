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

// DrgAttachmentUpdate Data plane information about Drg Attachment label and its associated vrf
// label.
type DrgAttachmentUpdate struct {

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

	// The label given to the associated VRF
	VrfLabel *int `mandatory:"false" json:"vrfLabel"`

	// Indicates whether the attachment is whitelisted for inter-region transit
	IsWhitelistedInterRegTrans *bool `mandatory:"false" json:"isWhitelistedInterRegTrans"`

	// Indicates whether the attachment is whitelisted for inter-region transit
	IsDisintermediated *bool `mandatory:"false" json:"isDisintermediated"`

	// The OCID for the DRG Attachment's compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID for the DRG
	DrgId *string `mandatory:"false" json:"drgId"`

	// The OCID for the DRG Attachment's Route Table
	RouteTableId *string `mandatory:"false" json:"routeTableId"`

	// The value of the tags.
	TagSlug *string `mandatory:"false" json:"tagSlug"`

	// The value of the peered region name.
	PeerRegionName *string `mandatory:"false" json:"peerRegionName"`

	// The OCID for the VCN
	VcnId *string `mandatory:"false" json:"vcnId"`

	// Indicates whether the DRG is whitelisted for Substrate Access DRG
	IsSubstrateAccess *bool `mandatory:"false" json:"isSubstrateAccess"`

	// The type of the attachment.
	DrgAttachmentType DrgAttachmentUpdateDrgAttachmentTypeEnum `mandatory:"false" json:"drgAttachmentType,omitempty"`

	// indicates the VIP type of the associated Drg
	DrgVipType DrgAttachmentUpdateDrgVipTypeEnum `mandatory:"false" json:"drgVipType,omitempty"`
}

// GetId returns Id
func (m DrgAttachmentUpdate) GetId() *string {
	return m.Id
}

// GetIsDelete returns IsDelete
func (m DrgAttachmentUpdate) GetIsDelete() *bool {
	return m.IsDelete
}

// GetTimeUpdated returns TimeUpdated
func (m DrgAttachmentUpdate) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m DrgAttachmentUpdate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrgAttachmentUpdate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDrgAttachmentUpdateDrgAttachmentTypeEnum(string(m.DrgAttachmentType)); !ok && m.DrgAttachmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DrgAttachmentType: %s. Supported values are: %s.", m.DrgAttachmentType, strings.Join(GetDrgAttachmentUpdateDrgAttachmentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDrgAttachmentUpdateDrgVipTypeEnum(string(m.DrgVipType)); !ok && m.DrgVipType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DrgVipType: %s. Supported values are: %s.", m.DrgVipType, strings.Join(GetDrgAttachmentUpdateDrgVipTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DrgAttachmentUpdate) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDrgAttachmentUpdate DrgAttachmentUpdate
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDrgAttachmentUpdate
	}{
		"DrgAttachmentUpdate",
		(MarshalTypeDrgAttachmentUpdate)(m),
	}

	return json.Marshal(&s)
}

// DrgAttachmentUpdateDrgAttachmentTypeEnum Enum with underlying type: string
type DrgAttachmentUpdateDrgAttachmentTypeEnum string

// Set of constants representing the allowable values for DrgAttachmentUpdateDrgAttachmentTypeEnum
const (
	DrgAttachmentUpdateDrgAttachmentTypeVcn                     DrgAttachmentUpdateDrgAttachmentTypeEnum = "VCN"
	DrgAttachmentUpdateDrgAttachmentTypeVirtualCircuit          DrgAttachmentUpdateDrgAttachmentTypeEnum = "VIRTUAL_CIRCUIT"
	DrgAttachmentUpdateDrgAttachmentTypeRemotePeeringConnection DrgAttachmentUpdateDrgAttachmentTypeEnum = "REMOTE_PEERING_CONNECTION"
	DrgAttachmentUpdateDrgAttachmentTypeIpsecTunnel             DrgAttachmentUpdateDrgAttachmentTypeEnum = "IPSEC_TUNNEL"
	DrgAttachmentUpdateDrgAttachmentTypeInternet                DrgAttachmentUpdateDrgAttachmentTypeEnum = "INTERNET"
	DrgAttachmentUpdateDrgAttachmentTypeInternalOnly            DrgAttachmentUpdateDrgAttachmentTypeEnum = "INTERNAL_ONLY"
	DrgAttachmentUpdateDrgAttachmentTypeLoopback                DrgAttachmentUpdateDrgAttachmentTypeEnum = "LOOPBACK"
)

var mappingDrgAttachmentUpdateDrgAttachmentTypeEnum = map[string]DrgAttachmentUpdateDrgAttachmentTypeEnum{
	"VCN":                       DrgAttachmentUpdateDrgAttachmentTypeVcn,
	"VIRTUAL_CIRCUIT":           DrgAttachmentUpdateDrgAttachmentTypeVirtualCircuit,
	"REMOTE_PEERING_CONNECTION": DrgAttachmentUpdateDrgAttachmentTypeRemotePeeringConnection,
	"IPSEC_TUNNEL":              DrgAttachmentUpdateDrgAttachmentTypeIpsecTunnel,
	"INTERNET":                  DrgAttachmentUpdateDrgAttachmentTypeInternet,
	"INTERNAL_ONLY":             DrgAttachmentUpdateDrgAttachmentTypeInternalOnly,
	"LOOPBACK":                  DrgAttachmentUpdateDrgAttachmentTypeLoopback,
}

var mappingDrgAttachmentUpdateDrgAttachmentTypeEnumLowerCase = map[string]DrgAttachmentUpdateDrgAttachmentTypeEnum{
	"vcn":                       DrgAttachmentUpdateDrgAttachmentTypeVcn,
	"virtual_circuit":           DrgAttachmentUpdateDrgAttachmentTypeVirtualCircuit,
	"remote_peering_connection": DrgAttachmentUpdateDrgAttachmentTypeRemotePeeringConnection,
	"ipsec_tunnel":              DrgAttachmentUpdateDrgAttachmentTypeIpsecTunnel,
	"internet":                  DrgAttachmentUpdateDrgAttachmentTypeInternet,
	"internal_only":             DrgAttachmentUpdateDrgAttachmentTypeInternalOnly,
	"loopback":                  DrgAttachmentUpdateDrgAttachmentTypeLoopback,
}

// GetDrgAttachmentUpdateDrgAttachmentTypeEnumValues Enumerates the set of values for DrgAttachmentUpdateDrgAttachmentTypeEnum
func GetDrgAttachmentUpdateDrgAttachmentTypeEnumValues() []DrgAttachmentUpdateDrgAttachmentTypeEnum {
	values := make([]DrgAttachmentUpdateDrgAttachmentTypeEnum, 0)
	for _, v := range mappingDrgAttachmentUpdateDrgAttachmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrgAttachmentUpdateDrgAttachmentTypeEnumStringValues Enumerates the set of values in String for DrgAttachmentUpdateDrgAttachmentTypeEnum
func GetDrgAttachmentUpdateDrgAttachmentTypeEnumStringValues() []string {
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

// GetMappingDrgAttachmentUpdateDrgAttachmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrgAttachmentUpdateDrgAttachmentTypeEnum(val string) (DrgAttachmentUpdateDrgAttachmentTypeEnum, bool) {
	enum, ok := mappingDrgAttachmentUpdateDrgAttachmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DrgAttachmentUpdateDrgVipTypeEnum Enum with underlying type: string
type DrgAttachmentUpdateDrgVipTypeEnum string

// Set of constants representing the allowable values for DrgAttachmentUpdateDrgVipTypeEnum
const (
	DrgAttachmentUpdateDrgVipTypeProd        DrgAttachmentUpdateDrgVipTypeEnum = "PROD"
	DrgAttachmentUpdateDrgVipTypeGamma       DrgAttachmentUpdateDrgVipTypeEnum = "GAMMA"
	DrgAttachmentUpdateDrgVipTypeFfabProd    DrgAttachmentUpdateDrgVipTypeEnum = "FFAB_PROD"
	DrgAttachmentUpdateDrgVipTypeFfabGamma   DrgAttachmentUpdateDrgVipTypeEnum = "FFAB_GAMMA"
	DrgAttachmentUpdateDrgVipTypeProdPod1    DrgAttachmentUpdateDrgVipTypeEnum = "PROD_POD1"
	DrgAttachmentUpdateDrgVipTypeAlpha       DrgAttachmentUpdateDrgVipTypeEnum = "ALPHA"
	DrgAttachmentUpdateDrgVipTypeBeta        DrgAttachmentUpdateDrgVipTypeEnum = "BETA"
	DrgAttachmentUpdateDrgVipTypeIndigoGamma DrgAttachmentUpdateDrgVipTypeEnum = "INDIGO_GAMMA"
	DrgAttachmentUpdateDrgVipTypeIndigoProd  DrgAttachmentUpdateDrgVipTypeEnum = "INDIGO_PROD"
	DrgAttachmentUpdateDrgVipTypeOasisGamma  DrgAttachmentUpdateDrgVipTypeEnum = "OASIS_GAMMA"
	DrgAttachmentUpdateDrgVipTypeOasisProd   DrgAttachmentUpdateDrgVipTypeEnum = "OASIS_PROD"
	DrgAttachmentUpdateDrgVipTypeProdPod2    DrgAttachmentUpdateDrgVipTypeEnum = "PROD_POD2"
)

var mappingDrgAttachmentUpdateDrgVipTypeEnum = map[string]DrgAttachmentUpdateDrgVipTypeEnum{
	"PROD":         DrgAttachmentUpdateDrgVipTypeProd,
	"GAMMA":        DrgAttachmentUpdateDrgVipTypeGamma,
	"FFAB_PROD":    DrgAttachmentUpdateDrgVipTypeFfabProd,
	"FFAB_GAMMA":   DrgAttachmentUpdateDrgVipTypeFfabGamma,
	"PROD_POD1":    DrgAttachmentUpdateDrgVipTypeProdPod1,
	"ALPHA":        DrgAttachmentUpdateDrgVipTypeAlpha,
	"BETA":         DrgAttachmentUpdateDrgVipTypeBeta,
	"INDIGO_GAMMA": DrgAttachmentUpdateDrgVipTypeIndigoGamma,
	"INDIGO_PROD":  DrgAttachmentUpdateDrgVipTypeIndigoProd,
	"OASIS_GAMMA":  DrgAttachmentUpdateDrgVipTypeOasisGamma,
	"OASIS_PROD":   DrgAttachmentUpdateDrgVipTypeOasisProd,
	"PROD_POD2":    DrgAttachmentUpdateDrgVipTypeProdPod2,
}

var mappingDrgAttachmentUpdateDrgVipTypeEnumLowerCase = map[string]DrgAttachmentUpdateDrgVipTypeEnum{
	"prod":         DrgAttachmentUpdateDrgVipTypeProd,
	"gamma":        DrgAttachmentUpdateDrgVipTypeGamma,
	"ffab_prod":    DrgAttachmentUpdateDrgVipTypeFfabProd,
	"ffab_gamma":   DrgAttachmentUpdateDrgVipTypeFfabGamma,
	"prod_pod1":    DrgAttachmentUpdateDrgVipTypeProdPod1,
	"alpha":        DrgAttachmentUpdateDrgVipTypeAlpha,
	"beta":         DrgAttachmentUpdateDrgVipTypeBeta,
	"indigo_gamma": DrgAttachmentUpdateDrgVipTypeIndigoGamma,
	"indigo_prod":  DrgAttachmentUpdateDrgVipTypeIndigoProd,
	"oasis_gamma":  DrgAttachmentUpdateDrgVipTypeOasisGamma,
	"oasis_prod":   DrgAttachmentUpdateDrgVipTypeOasisProd,
	"prod_pod2":    DrgAttachmentUpdateDrgVipTypeProdPod2,
}

// GetDrgAttachmentUpdateDrgVipTypeEnumValues Enumerates the set of values for DrgAttachmentUpdateDrgVipTypeEnum
func GetDrgAttachmentUpdateDrgVipTypeEnumValues() []DrgAttachmentUpdateDrgVipTypeEnum {
	values := make([]DrgAttachmentUpdateDrgVipTypeEnum, 0)
	for _, v := range mappingDrgAttachmentUpdateDrgVipTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrgAttachmentUpdateDrgVipTypeEnumStringValues Enumerates the set of values in String for DrgAttachmentUpdateDrgVipTypeEnum
func GetDrgAttachmentUpdateDrgVipTypeEnumStringValues() []string {
	return []string{
		"PROD",
		"GAMMA",
		"FFAB_PROD",
		"FFAB_GAMMA",
		"PROD_POD1",
		"ALPHA",
		"BETA",
		"INDIGO_GAMMA",
		"INDIGO_PROD",
		"OASIS_GAMMA",
		"OASIS_PROD",
		"PROD_POD2",
	}
}

// GetMappingDrgAttachmentUpdateDrgVipTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrgAttachmentUpdateDrgVipTypeEnum(val string) (DrgAttachmentUpdateDrgVipTypeEnum, bool) {
	enum, ok := mappingDrgAttachmentUpdateDrgVipTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
