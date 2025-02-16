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

// DrgAttachmentInternalInfo Internal operational information about a DRG Attachment.
type DrgAttachmentInternalInfo struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the DRG attachment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
	DrgId *string `mandatory:"true" json:"drgId"`

	// The DRG attachment's Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
	Id *string `mandatory:"true" json:"id"`

	// The DRG attachment's current state.
	LifecycleState DrgAttachmentInternalInfoLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The date and time the DRG attachment was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table that is assigned to this attachment.
	// The DRG route table manages traffic inside the DRG.
	DrgRouteTableId *string `mandatory:"false" json:"drgRouteTableId"`

	NetworkDetails DrgAttachmentNetworkDetails `mandatory:"false" json:"networkDetails"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Security attributes (https://docs.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm#security-attributes) are labels
	// for a resource that can be referenced in a Zero Trust Packet Routing (https://docs.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm)
	// (ZPR) policy to control access to ZPR-supported resources.
	// Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the DRG attachment is using.
	// For information about why you would associate a route table with a DRG attachment, see:
	//   * Transit Routing: Access to Multiple VCNs in Same Region (https://docs.oracle.com/iaas/Content/Network/Tasks/transitrouting.htm)
	//   * Transit Routing: Private Access to Oracle Services (https://docs.oracle.com/iaas/Content/Network/Tasks/transitroutingoracleservices.htm)
	// This field is deprecated. Instead, use the `networkDetails` field to view the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the attached resource.
	RouteTableId *string `mandatory:"false" json:"routeTableId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.
	// This field is deprecated. Instead, use the `networkDetails` field to view the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the attached resource.
	VcnId *string `mandatory:"false" json:"vcnId"`

	// STANDARD applies to all regional resources which are customer visible, GDRG_SERVICE_RESOURCE applies to
	// internal resources created to back GlobalDRGAttachments, and GDRG_MESH_RPC applies to internal RPC Attachments
	// used to facilitate GlobalDRG functionality.
	InternalType DrgAttachmentInternalTypeEnum `mandatory:"false" json:"internalType,omitempty"`

	// Indicates if transitive traffic is enabled for this DRG attachment. This field is
	// only supported for VirtualCircuit and IPSec DRG attachments.
	TransitiveTrafficEnabled DrgAttachmentTransitiveTrafficStateEnum `mandatory:"false" json:"transitiveTrafficEnabled,omitempty"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the export route distribution used to specify how routes in the assigned DRG route table
	// are advertised to the attachment.
	// If this value is null, no routes are advertised through this attachment.
	ExportDrgRouteDistributionId *string `mandatory:"false" json:"exportDrgRouteDistributionId"`

	// Indicates whether the DRG attachment and attached network live in a different tenancy than the DRG.
	// Example: `false`
	IsCrossTenancy *bool `mandatory:"false" json:"isCrossTenancy"`

	// Route Filtering type of the attachment
	VcnRouteType DrgAttachmentInternalInfoVcnRouteTypeEnum `mandatory:"false" json:"vcnRouteType,omitempty"`

	// The DRG's current state.
	DrgVersion DrgAttachmentInternalInfoDrgVersionEnum `mandatory:"false" json:"drgVersion,omitempty"`

	// indicates the VIP type of the associated Drg
	DrgVipType DrgAttachmentInternalInfoDrgVipTypeEnum `mandatory:"false" json:"drgVipType,omitempty"`

	// The time DRG entered into a particular lifecycle state
	// Example: '2016-08-25T21:10:29.600Z'
	TimeLifecycleStateEntered *common.SDKTime `mandatory:"false" json:"timeLifecycleStateEntered"`

	// The OCID of this network attached to the DRG.
	AttachedNetworkId *string `mandatory:"false" json:"attachedNetworkId"`

	// The OCID of IPSEC connection.
	IpsecConnectionId *string `mandatory:"false" json:"ipsecConnectionId"`

	// The OCID of transport attachment (vc attachment).
	TransportAttachmentId *string `mandatory:"false" json:"transportAttachmentId"`

	// The MPLS label which identifies this DRG Attachment in encapsulated
	// traffic sent to either the DRG egress or ingress redirectors.
	// This label is scoped by the egress and ingress redirector IPs.
	Label *int `mandatory:"false" json:"label"`

	// Type of the DRG attachment
	Type DrgAttachmentInternalInfoTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The BGP ASN to use for the IPSec Connection''s route target
	RegionalOciAsn *string `mandatory:"false" json:"regionalOciAsn"`

	// Routes which are exported to the attachment are exported by
	// El Paso as L3VPN routes to the route reflectors
	// with the route target set to the value of the attachment''s export route
	// target. This is the label of the (asn:label) value.
	ExportRouteTargetLabel *int `mandatory:"false" json:"exportRouteTargetLabel"`

	// Indicates whether the create or delete WF is in progress for a particular DRG Attachment
	// This value is necessary to determine where to pick up the create DRG Attachment request on a retried request
	IsWfInProgress *bool `mandatory:"false" json:"isWfInProgress"`

	// Whether the Fast Connect is an FFAB virtualCircuit.
	// Example: `true`
	IsFfab *bool `mandatory:"false" json:"isFfab"`

	// Whether the Fast Connect exists through an edge pop region.
	// Example: `true`
	IsEdgePop *bool `mandatory:"false" json:"isEdgePop"`

	// The OCI region name
	RegionName *string `mandatory:"false" json:"regionName"`

	// Is Global Fast Connect
	// Example: `true`'
	IsGfc *bool `mandatory:"false" json:"isGfc"`

	// Indicates whether the backdoor API was used to override the exportRouteDistributionId value so no
	// export policies are sent for this attachment
	IsBlockExport *bool `mandatory:"false" json:"isBlockExport"`

	// The peer attachment route target.
	PeerAttachmentRouteTarget *string `mandatory:"false" json:"peerAttachmentRouteTarget"`

	// The peer OCI region name
	PeerRegionName *string `mandatory:"false" json:"peerRegionName"`

	// The list of BYOIP Range OCIDs used to be accessible to the
	// internet via this DRG.
	ByoipRangeIds []string `mandatory:"false" json:"byoipRangeIds"`

	// The list of Public IPv4 or IPv6 CIDRs ["100.0.0.0/24"] used to be
	// accessible to the internet via this DRG.
	PublicCidrBlocks []string `mandatory:"false" json:"publicCidrBlocks"`

	// The maximum DrgPathLength which this export policy should advertise.
	// Any route with a longer DrgPathLength should be suppressed.
	// Zero value indicates that there is no configured limit.
	MaxAdvertisedDrgPathLength *int `mandatory:"false" json:"maxAdvertisedDrgPathLength"`

	// Indicates whether the attachment is "disintermediated" or not.
	IsDisintermediated *bool `mandatory:"false" json:"isDisintermediated"`

	// Indicates whether the attachment is "Substrate Access DrgAttachment" or not.
	IsSubstrateAccess *bool `mandatory:"false" json:"isSubstrateAccess"`

	// Indicates whether the route unification is completed for the attachment,
	// i.e., Common RT configured on both PNP and C3 successfully so that we no longer need per Attachment RT.
	IsAttachmentRouteUnificationComplete *bool `mandatory:"false" json:"isAttachmentRouteUnificationComplete"`

	// While migrating existing attachments for route unification,
	// this indicates whether the route unification is completed in C3 or not.
	IsAttachmentRouteUnificationCompleteInC3 *bool `mandatory:"false" json:"isAttachmentRouteUnificationCompleteInC3"`
}

func (m DrgAttachmentInternalInfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrgAttachmentInternalInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDrgAttachmentInternalInfoLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDrgAttachmentInternalInfoLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDrgAttachmentInternalTypeEnum(string(m.InternalType)); !ok && m.InternalType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InternalType: %s. Supported values are: %s.", m.InternalType, strings.Join(GetDrgAttachmentInternalTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDrgAttachmentTransitiveTrafficStateEnum(string(m.TransitiveTrafficEnabled)); !ok && m.TransitiveTrafficEnabled != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TransitiveTrafficEnabled: %s. Supported values are: %s.", m.TransitiveTrafficEnabled, strings.Join(GetDrgAttachmentTransitiveTrafficStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDrgAttachmentInternalInfoVcnRouteTypeEnum(string(m.VcnRouteType)); !ok && m.VcnRouteType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VcnRouteType: %s. Supported values are: %s.", m.VcnRouteType, strings.Join(GetDrgAttachmentInternalInfoVcnRouteTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDrgAttachmentInternalInfoDrgVersionEnum(string(m.DrgVersion)); !ok && m.DrgVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DrgVersion: %s. Supported values are: %s.", m.DrgVersion, strings.Join(GetDrgAttachmentInternalInfoDrgVersionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDrgAttachmentInternalInfoDrgVipTypeEnum(string(m.DrgVipType)); !ok && m.DrgVipType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DrgVipType: %s. Supported values are: %s.", m.DrgVipType, strings.Join(GetDrgAttachmentInternalInfoDrgVipTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDrgAttachmentInternalInfoTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetDrgAttachmentInternalInfoTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DrgAttachmentInternalInfo) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                              *string                                     `json:"displayName"`
		TimeCreated                              *common.SDKTime                             `json:"timeCreated"`
		DrgRouteTableId                          *string                                     `json:"drgRouteTableId"`
		NetworkDetails                           drgattachmentnetworkdetails                 `json:"networkDetails"`
		DefinedTags                              map[string]map[string]interface{}           `json:"definedTags"`
		FreeformTags                             map[string]string                           `json:"freeformTags"`
		SecurityAttributes                       map[string]map[string]interface{}           `json:"securityAttributes"`
		RouteTableId                             *string                                     `json:"routeTableId"`
		VcnId                                    *string                                     `json:"vcnId"`
		InternalType                             DrgAttachmentInternalTypeEnum               `json:"internalType"`
		TransitiveTrafficEnabled                 DrgAttachmentTransitiveTrafficStateEnum     `json:"transitiveTrafficEnabled"`
		ExportDrgRouteDistributionId             *string                                     `json:"exportDrgRouteDistributionId"`
		IsCrossTenancy                           *bool                                       `json:"isCrossTenancy"`
		VcnRouteType                             DrgAttachmentInternalInfoVcnRouteTypeEnum   `json:"vcnRouteType"`
		DrgVersion                               DrgAttachmentInternalInfoDrgVersionEnum     `json:"drgVersion"`
		DrgVipType                               DrgAttachmentInternalInfoDrgVipTypeEnum     `json:"drgVipType"`
		TimeLifecycleStateEntered                *common.SDKTime                             `json:"timeLifecycleStateEntered"`
		AttachedNetworkId                        *string                                     `json:"attachedNetworkId"`
		IpsecConnectionId                        *string                                     `json:"ipsecConnectionId"`
		TransportAttachmentId                    *string                                     `json:"transportAttachmentId"`
		Label                                    *int                                        `json:"label"`
		Type                                     DrgAttachmentInternalInfoTypeEnum           `json:"type"`
		RegionalOciAsn                           *string                                     `json:"regionalOciAsn"`
		ExportRouteTargetLabel                   *int                                        `json:"exportRouteTargetLabel"`
		IsWfInProgress                           *bool                                       `json:"isWfInProgress"`
		IsFfab                                   *bool                                       `json:"isFfab"`
		IsEdgePop                                *bool                                       `json:"isEdgePop"`
		RegionName                               *string                                     `json:"regionName"`
		IsGfc                                    *bool                                       `json:"isGfc"`
		IsBlockExport                            *bool                                       `json:"isBlockExport"`
		PeerAttachmentRouteTarget                *string                                     `json:"peerAttachmentRouteTarget"`
		PeerRegionName                           *string                                     `json:"peerRegionName"`
		ByoipRangeIds                            []string                                    `json:"byoipRangeIds"`
		PublicCidrBlocks                         []string                                    `json:"publicCidrBlocks"`
		MaxAdvertisedDrgPathLength               *int                                        `json:"maxAdvertisedDrgPathLength"`
		IsDisintermediated                       *bool                                       `json:"isDisintermediated"`
		IsSubstrateAccess                        *bool                                       `json:"isSubstrateAccess"`
		IsAttachmentRouteUnificationComplete     *bool                                       `json:"isAttachmentRouteUnificationComplete"`
		IsAttachmentRouteUnificationCompleteInC3 *bool                                       `json:"isAttachmentRouteUnificationCompleteInC3"`
		CompartmentId                            *string                                     `json:"compartmentId"`
		DrgId                                    *string                                     `json:"drgId"`
		Id                                       *string                                     `json:"id"`
		LifecycleState                           DrgAttachmentInternalInfoLifecycleStateEnum `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.TimeCreated = model.TimeCreated

	m.DrgRouteTableId = model.DrgRouteTableId

	nn, e = model.NetworkDetails.UnmarshalPolymorphicJSON(model.NetworkDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.NetworkDetails = nn.(DrgAttachmentNetworkDetails)
	} else {
		m.NetworkDetails = nil
	}

	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	m.SecurityAttributes = model.SecurityAttributes

	m.RouteTableId = model.RouteTableId

	m.VcnId = model.VcnId

	m.InternalType = model.InternalType

	m.TransitiveTrafficEnabled = model.TransitiveTrafficEnabled

	m.ExportDrgRouteDistributionId = model.ExportDrgRouteDistributionId

	m.IsCrossTenancy = model.IsCrossTenancy

	m.VcnRouteType = model.VcnRouteType

	m.DrgVersion = model.DrgVersion

	m.DrgVipType = model.DrgVipType

	m.TimeLifecycleStateEntered = model.TimeLifecycleStateEntered

	m.AttachedNetworkId = model.AttachedNetworkId

	m.IpsecConnectionId = model.IpsecConnectionId

	m.TransportAttachmentId = model.TransportAttachmentId

	m.Label = model.Label

	m.Type = model.Type

	m.RegionalOciAsn = model.RegionalOciAsn

	m.ExportRouteTargetLabel = model.ExportRouteTargetLabel

	m.IsWfInProgress = model.IsWfInProgress

	m.IsFfab = model.IsFfab

	m.IsEdgePop = model.IsEdgePop

	m.RegionName = model.RegionName

	m.IsGfc = model.IsGfc

	m.IsBlockExport = model.IsBlockExport

	m.PeerAttachmentRouteTarget = model.PeerAttachmentRouteTarget

	m.PeerRegionName = model.PeerRegionName

	m.ByoipRangeIds = make([]string, len(model.ByoipRangeIds))
	copy(m.ByoipRangeIds, model.ByoipRangeIds)
	m.PublicCidrBlocks = make([]string, len(model.PublicCidrBlocks))
	copy(m.PublicCidrBlocks, model.PublicCidrBlocks)
	m.MaxAdvertisedDrgPathLength = model.MaxAdvertisedDrgPathLength

	m.IsDisintermediated = model.IsDisintermediated

	m.IsSubstrateAccess = model.IsSubstrateAccess

	m.IsAttachmentRouteUnificationComplete = model.IsAttachmentRouteUnificationComplete

	m.IsAttachmentRouteUnificationCompleteInC3 = model.IsAttachmentRouteUnificationCompleteInC3

	m.CompartmentId = model.CompartmentId

	m.DrgId = model.DrgId

	m.Id = model.Id

	m.LifecycleState = model.LifecycleState

	return
}

// DrgAttachmentInternalInfoLifecycleStateEnum Enum with underlying type: string
type DrgAttachmentInternalInfoLifecycleStateEnum string

// Set of constants representing the allowable values for DrgAttachmentInternalInfoLifecycleStateEnum
const (
	DrgAttachmentInternalInfoLifecycleStateAttaching DrgAttachmentInternalInfoLifecycleStateEnum = "ATTACHING"
	DrgAttachmentInternalInfoLifecycleStateAttached  DrgAttachmentInternalInfoLifecycleStateEnum = "ATTACHED"
	DrgAttachmentInternalInfoLifecycleStateDetaching DrgAttachmentInternalInfoLifecycleStateEnum = "DETACHING"
	DrgAttachmentInternalInfoLifecycleStateDetached  DrgAttachmentInternalInfoLifecycleStateEnum = "DETACHED"
	DrgAttachmentInternalInfoLifecycleStateUpdating  DrgAttachmentInternalInfoLifecycleStateEnum = "UPDATING"
)

var mappingDrgAttachmentInternalInfoLifecycleStateEnum = map[string]DrgAttachmentInternalInfoLifecycleStateEnum{
	"ATTACHING": DrgAttachmentInternalInfoLifecycleStateAttaching,
	"ATTACHED":  DrgAttachmentInternalInfoLifecycleStateAttached,
	"DETACHING": DrgAttachmentInternalInfoLifecycleStateDetaching,
	"DETACHED":  DrgAttachmentInternalInfoLifecycleStateDetached,
	"UPDATING":  DrgAttachmentInternalInfoLifecycleStateUpdating,
}

var mappingDrgAttachmentInternalInfoLifecycleStateEnumLowerCase = map[string]DrgAttachmentInternalInfoLifecycleStateEnum{
	"attaching": DrgAttachmentInternalInfoLifecycleStateAttaching,
	"attached":  DrgAttachmentInternalInfoLifecycleStateAttached,
	"detaching": DrgAttachmentInternalInfoLifecycleStateDetaching,
	"detached":  DrgAttachmentInternalInfoLifecycleStateDetached,
	"updating":  DrgAttachmentInternalInfoLifecycleStateUpdating,
}

// GetDrgAttachmentInternalInfoLifecycleStateEnumValues Enumerates the set of values for DrgAttachmentInternalInfoLifecycleStateEnum
func GetDrgAttachmentInternalInfoLifecycleStateEnumValues() []DrgAttachmentInternalInfoLifecycleStateEnum {
	values := make([]DrgAttachmentInternalInfoLifecycleStateEnum, 0)
	for _, v := range mappingDrgAttachmentInternalInfoLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDrgAttachmentInternalInfoLifecycleStateEnumStringValues Enumerates the set of values in String for DrgAttachmentInternalInfoLifecycleStateEnum
func GetDrgAttachmentInternalInfoLifecycleStateEnumStringValues() []string {
	return []string{
		"ATTACHING",
		"ATTACHED",
		"DETACHING",
		"DETACHED",
		"UPDATING",
	}
}

// GetMappingDrgAttachmentInternalInfoLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrgAttachmentInternalInfoLifecycleStateEnum(val string) (DrgAttachmentInternalInfoLifecycleStateEnum, bool) {
	enum, ok := mappingDrgAttachmentInternalInfoLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DrgAttachmentInternalInfoVcnRouteTypeEnum Enum with underlying type: string
type DrgAttachmentInternalInfoVcnRouteTypeEnum string

// Set of constants representing the allowable values for DrgAttachmentInternalInfoVcnRouteTypeEnum
const (
	DrgAttachmentInternalInfoVcnRouteTypeSubnetCidrs DrgAttachmentInternalInfoVcnRouteTypeEnum = "SUBNET_CIDRS"
	DrgAttachmentInternalInfoVcnRouteTypeVcnCidrs    DrgAttachmentInternalInfoVcnRouteTypeEnum = "VCN_CIDRS"
)

var mappingDrgAttachmentInternalInfoVcnRouteTypeEnum = map[string]DrgAttachmentInternalInfoVcnRouteTypeEnum{
	"SUBNET_CIDRS": DrgAttachmentInternalInfoVcnRouteTypeSubnetCidrs,
	"VCN_CIDRS":    DrgAttachmentInternalInfoVcnRouteTypeVcnCidrs,
}

var mappingDrgAttachmentInternalInfoVcnRouteTypeEnumLowerCase = map[string]DrgAttachmentInternalInfoVcnRouteTypeEnum{
	"subnet_cidrs": DrgAttachmentInternalInfoVcnRouteTypeSubnetCidrs,
	"vcn_cidrs":    DrgAttachmentInternalInfoVcnRouteTypeVcnCidrs,
}

// GetDrgAttachmentInternalInfoVcnRouteTypeEnumValues Enumerates the set of values for DrgAttachmentInternalInfoVcnRouteTypeEnum
func GetDrgAttachmentInternalInfoVcnRouteTypeEnumValues() []DrgAttachmentInternalInfoVcnRouteTypeEnum {
	values := make([]DrgAttachmentInternalInfoVcnRouteTypeEnum, 0)
	for _, v := range mappingDrgAttachmentInternalInfoVcnRouteTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrgAttachmentInternalInfoVcnRouteTypeEnumStringValues Enumerates the set of values in String for DrgAttachmentInternalInfoVcnRouteTypeEnum
func GetDrgAttachmentInternalInfoVcnRouteTypeEnumStringValues() []string {
	return []string{
		"SUBNET_CIDRS",
		"VCN_CIDRS",
	}
}

// GetMappingDrgAttachmentInternalInfoVcnRouteTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrgAttachmentInternalInfoVcnRouteTypeEnum(val string) (DrgAttachmentInternalInfoVcnRouteTypeEnum, bool) {
	enum, ok := mappingDrgAttachmentInternalInfoVcnRouteTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DrgAttachmentInternalInfoDrgVersionEnum Enum with underlying type: string
type DrgAttachmentInternalInfoDrgVersionEnum string

// Set of constants representing the allowable values for DrgAttachmentInternalInfoDrgVersionEnum
const (
	DrgAttachmentInternalInfoDrgVersion1 DrgAttachmentInternalInfoDrgVersionEnum = "VERSION_1"
	DrgAttachmentInternalInfoDrgVersion2 DrgAttachmentInternalInfoDrgVersionEnum = "VERSION_2"
)

var mappingDrgAttachmentInternalInfoDrgVersionEnum = map[string]DrgAttachmentInternalInfoDrgVersionEnum{
	"VERSION_1": DrgAttachmentInternalInfoDrgVersion1,
	"VERSION_2": DrgAttachmentInternalInfoDrgVersion2,
}

var mappingDrgAttachmentInternalInfoDrgVersionEnumLowerCase = map[string]DrgAttachmentInternalInfoDrgVersionEnum{
	"version_1": DrgAttachmentInternalInfoDrgVersion1,
	"version_2": DrgAttachmentInternalInfoDrgVersion2,
}

// GetDrgAttachmentInternalInfoDrgVersionEnumValues Enumerates the set of values for DrgAttachmentInternalInfoDrgVersionEnum
func GetDrgAttachmentInternalInfoDrgVersionEnumValues() []DrgAttachmentInternalInfoDrgVersionEnum {
	values := make([]DrgAttachmentInternalInfoDrgVersionEnum, 0)
	for _, v := range mappingDrgAttachmentInternalInfoDrgVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetDrgAttachmentInternalInfoDrgVersionEnumStringValues Enumerates the set of values in String for DrgAttachmentInternalInfoDrgVersionEnum
func GetDrgAttachmentInternalInfoDrgVersionEnumStringValues() []string {
	return []string{
		"VERSION_1",
		"VERSION_2",
	}
}

// GetMappingDrgAttachmentInternalInfoDrgVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrgAttachmentInternalInfoDrgVersionEnum(val string) (DrgAttachmentInternalInfoDrgVersionEnum, bool) {
	enum, ok := mappingDrgAttachmentInternalInfoDrgVersionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DrgAttachmentInternalInfoDrgVipTypeEnum Enum with underlying type: string
type DrgAttachmentInternalInfoDrgVipTypeEnum string

// Set of constants representing the allowable values for DrgAttachmentInternalInfoDrgVipTypeEnum
const (
	DrgAttachmentInternalInfoDrgVipTypeProd        DrgAttachmentInternalInfoDrgVipTypeEnum = "PROD"
	DrgAttachmentInternalInfoDrgVipTypeGamma       DrgAttachmentInternalInfoDrgVipTypeEnum = "GAMMA"
	DrgAttachmentInternalInfoDrgVipTypeFfabProd    DrgAttachmentInternalInfoDrgVipTypeEnum = "FFAB_PROD"
	DrgAttachmentInternalInfoDrgVipTypeFfabGamma   DrgAttachmentInternalInfoDrgVipTypeEnum = "FFAB_GAMMA"
	DrgAttachmentInternalInfoDrgVipTypeProdPod1    DrgAttachmentInternalInfoDrgVipTypeEnum = "PROD_POD1"
	DrgAttachmentInternalInfoDrgVipTypeAlpha       DrgAttachmentInternalInfoDrgVipTypeEnum = "ALPHA"
	DrgAttachmentInternalInfoDrgVipTypeBeta        DrgAttachmentInternalInfoDrgVipTypeEnum = "BETA"
	DrgAttachmentInternalInfoDrgVipTypeIndigoGamma DrgAttachmentInternalInfoDrgVipTypeEnum = "INDIGO_GAMMA"
	DrgAttachmentInternalInfoDrgVipTypeIndigoProd  DrgAttachmentInternalInfoDrgVipTypeEnum = "INDIGO_PROD"
	DrgAttachmentInternalInfoDrgVipTypeProdPod2    DrgAttachmentInternalInfoDrgVipTypeEnum = "PROD_POD2"
)

var mappingDrgAttachmentInternalInfoDrgVipTypeEnum = map[string]DrgAttachmentInternalInfoDrgVipTypeEnum{
	"PROD":         DrgAttachmentInternalInfoDrgVipTypeProd,
	"GAMMA":        DrgAttachmentInternalInfoDrgVipTypeGamma,
	"FFAB_PROD":    DrgAttachmentInternalInfoDrgVipTypeFfabProd,
	"FFAB_GAMMA":   DrgAttachmentInternalInfoDrgVipTypeFfabGamma,
	"PROD_POD1":    DrgAttachmentInternalInfoDrgVipTypeProdPod1,
	"ALPHA":        DrgAttachmentInternalInfoDrgVipTypeAlpha,
	"BETA":         DrgAttachmentInternalInfoDrgVipTypeBeta,
	"INDIGO_GAMMA": DrgAttachmentInternalInfoDrgVipTypeIndigoGamma,
	"INDIGO_PROD":  DrgAttachmentInternalInfoDrgVipTypeIndigoProd,
	"PROD_POD2":    DrgAttachmentInternalInfoDrgVipTypeProdPod2,
}

var mappingDrgAttachmentInternalInfoDrgVipTypeEnumLowerCase = map[string]DrgAttachmentInternalInfoDrgVipTypeEnum{
	"prod":         DrgAttachmentInternalInfoDrgVipTypeProd,
	"gamma":        DrgAttachmentInternalInfoDrgVipTypeGamma,
	"ffab_prod":    DrgAttachmentInternalInfoDrgVipTypeFfabProd,
	"ffab_gamma":   DrgAttachmentInternalInfoDrgVipTypeFfabGamma,
	"prod_pod1":    DrgAttachmentInternalInfoDrgVipTypeProdPod1,
	"alpha":        DrgAttachmentInternalInfoDrgVipTypeAlpha,
	"beta":         DrgAttachmentInternalInfoDrgVipTypeBeta,
	"indigo_gamma": DrgAttachmentInternalInfoDrgVipTypeIndigoGamma,
	"indigo_prod":  DrgAttachmentInternalInfoDrgVipTypeIndigoProd,
	"prod_pod2":    DrgAttachmentInternalInfoDrgVipTypeProdPod2,
}

// GetDrgAttachmentInternalInfoDrgVipTypeEnumValues Enumerates the set of values for DrgAttachmentInternalInfoDrgVipTypeEnum
func GetDrgAttachmentInternalInfoDrgVipTypeEnumValues() []DrgAttachmentInternalInfoDrgVipTypeEnum {
	values := make([]DrgAttachmentInternalInfoDrgVipTypeEnum, 0)
	for _, v := range mappingDrgAttachmentInternalInfoDrgVipTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrgAttachmentInternalInfoDrgVipTypeEnumStringValues Enumerates the set of values in String for DrgAttachmentInternalInfoDrgVipTypeEnum
func GetDrgAttachmentInternalInfoDrgVipTypeEnumStringValues() []string {
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
		"PROD_POD2",
	}
}

// GetMappingDrgAttachmentInternalInfoDrgVipTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrgAttachmentInternalInfoDrgVipTypeEnum(val string) (DrgAttachmentInternalInfoDrgVipTypeEnum, bool) {
	enum, ok := mappingDrgAttachmentInternalInfoDrgVipTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DrgAttachmentInternalInfoTypeEnum Enum with underlying type: string
type DrgAttachmentInternalInfoTypeEnum string

// Set of constants representing the allowable values for DrgAttachmentInternalInfoTypeEnum
const (
	DrgAttachmentInternalInfoTypeVcn                     DrgAttachmentInternalInfoTypeEnum = "VCN"
	DrgAttachmentInternalInfoTypeVirtualCircuit          DrgAttachmentInternalInfoTypeEnum = "VIRTUAL_CIRCUIT"
	DrgAttachmentInternalInfoTypeRemotePeeringConnection DrgAttachmentInternalInfoTypeEnum = "REMOTE_PEERING_CONNECTION"
	DrgAttachmentInternalInfoTypeIpsecTunnel             DrgAttachmentInternalInfoTypeEnum = "IPSEC_TUNNEL"
	DrgAttachmentInternalInfoTypeInternet                DrgAttachmentInternalInfoTypeEnum = "INTERNET"
	DrgAttachmentInternalInfoTypeInternalOnly            DrgAttachmentInternalInfoTypeEnum = "INTERNAL_ONLY"
	DrgAttachmentInternalInfoTypeLoopback                DrgAttachmentInternalInfoTypeEnum = "LOOPBACK"
)

var mappingDrgAttachmentInternalInfoTypeEnum = map[string]DrgAttachmentInternalInfoTypeEnum{
	"VCN":                       DrgAttachmentInternalInfoTypeVcn,
	"VIRTUAL_CIRCUIT":           DrgAttachmentInternalInfoTypeVirtualCircuit,
	"REMOTE_PEERING_CONNECTION": DrgAttachmentInternalInfoTypeRemotePeeringConnection,
	"IPSEC_TUNNEL":              DrgAttachmentInternalInfoTypeIpsecTunnel,
	"INTERNET":                  DrgAttachmentInternalInfoTypeInternet,
	"INTERNAL_ONLY":             DrgAttachmentInternalInfoTypeInternalOnly,
	"LOOPBACK":                  DrgAttachmentInternalInfoTypeLoopback,
}

var mappingDrgAttachmentInternalInfoTypeEnumLowerCase = map[string]DrgAttachmentInternalInfoTypeEnum{
	"vcn":                       DrgAttachmentInternalInfoTypeVcn,
	"virtual_circuit":           DrgAttachmentInternalInfoTypeVirtualCircuit,
	"remote_peering_connection": DrgAttachmentInternalInfoTypeRemotePeeringConnection,
	"ipsec_tunnel":              DrgAttachmentInternalInfoTypeIpsecTunnel,
	"internet":                  DrgAttachmentInternalInfoTypeInternet,
	"internal_only":             DrgAttachmentInternalInfoTypeInternalOnly,
	"loopback":                  DrgAttachmentInternalInfoTypeLoopback,
}

// GetDrgAttachmentInternalInfoTypeEnumValues Enumerates the set of values for DrgAttachmentInternalInfoTypeEnum
func GetDrgAttachmentInternalInfoTypeEnumValues() []DrgAttachmentInternalInfoTypeEnum {
	values := make([]DrgAttachmentInternalInfoTypeEnum, 0)
	for _, v := range mappingDrgAttachmentInternalInfoTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrgAttachmentInternalInfoTypeEnumStringValues Enumerates the set of values in String for DrgAttachmentInternalInfoTypeEnum
func GetDrgAttachmentInternalInfoTypeEnumStringValues() []string {
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

// GetMappingDrgAttachmentInternalInfoTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrgAttachmentInternalInfoTypeEnum(val string) (DrgAttachmentInternalInfoTypeEnum, bool) {
	enum, ok := mappingDrgAttachmentInternalInfoTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
