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

// DrgImportPolicyUpdateV2 Data plane information about Drg Attachment import policy on
// a route table.
type DrgImportPolicyUpdateV2 struct {

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
	DrgAttachmentLabel *int64 `mandatory:"false" json:"drgAttachmentLabel"`

	// Unique identifier for the Drg Route Table that is affected by this update.
	VrfId *string `mandatory:"false" json:"vrfId"`

	// The label given to the Route Table.
	VrfLabel *int64 `mandatory:"false" json:"vrfLabel"`

	// The preference given to the attachment.
	DrgAttachmentPreference *int `mandatory:"false" json:"drgAttachmentPreference"`

	// The preference given to the attachment.
	DrgAttachmentPreferenceV2 *int `mandatory:"false" json:"drgAttachmentPreferenceV2"`

	// Boolean flag indicating whether the import policy is for
	// a route table with ecmp enabled
	IsEcmpEnabled *bool `mandatory:"false" json:"isEcmpEnabled"`

	// Boolean flag indicating whether the tenancy needs to import VCN_CIDR instead of SUBNET_CIDR
	DoImportVcnCidrs *bool `mandatory:"false" json:"doImportVcnCidrs"`

	// The type of the attachment.
	DrgAttachmentType DrgImportPolicyUpdateV2DrgAttachmentTypeEnum `mandatory:"false" json:"drgAttachmentType,omitempty"`
}

// GetId returns Id
func (m DrgImportPolicyUpdateV2) GetId() *string {
	return m.Id
}

// GetIsDelete returns IsDelete
func (m DrgImportPolicyUpdateV2) GetIsDelete() *bool {
	return m.IsDelete
}

// GetTimeUpdated returns TimeUpdated
func (m DrgImportPolicyUpdateV2) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m DrgImportPolicyUpdateV2) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrgImportPolicyUpdateV2) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDrgImportPolicyUpdateV2DrgAttachmentTypeEnum(string(m.DrgAttachmentType)); !ok && m.DrgAttachmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DrgAttachmentType: %s. Supported values are: %s.", m.DrgAttachmentType, strings.Join(GetDrgImportPolicyUpdateV2DrgAttachmentTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DrgImportPolicyUpdateV2) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDrgImportPolicyUpdateV2 DrgImportPolicyUpdateV2
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDrgImportPolicyUpdateV2
	}{
		"DrgImportPolicyUpdateV2",
		(MarshalTypeDrgImportPolicyUpdateV2)(m),
	}

	return json.Marshal(&s)
}

// DrgImportPolicyUpdateV2DrgAttachmentTypeEnum Enum with underlying type: string
type DrgImportPolicyUpdateV2DrgAttachmentTypeEnum string

// Set of constants representing the allowable values for DrgImportPolicyUpdateV2DrgAttachmentTypeEnum
const (
	DrgImportPolicyUpdateV2DrgAttachmentTypeVcn                     DrgImportPolicyUpdateV2DrgAttachmentTypeEnum = "VCN"
	DrgImportPolicyUpdateV2DrgAttachmentTypeVirtualCircuit          DrgImportPolicyUpdateV2DrgAttachmentTypeEnum = "VIRTUAL_CIRCUIT"
	DrgImportPolicyUpdateV2DrgAttachmentTypeRemotePeeringConnection DrgImportPolicyUpdateV2DrgAttachmentTypeEnum = "REMOTE_PEERING_CONNECTION"
	DrgImportPolicyUpdateV2DrgAttachmentTypeIpsecTunnel             DrgImportPolicyUpdateV2DrgAttachmentTypeEnum = "IPSEC_TUNNEL"
	DrgImportPolicyUpdateV2DrgAttachmentTypeInternalOnly            DrgImportPolicyUpdateV2DrgAttachmentTypeEnum = "INTERNAL_ONLY"
	DrgImportPolicyUpdateV2DrgAttachmentTypeLoopback                DrgImportPolicyUpdateV2DrgAttachmentTypeEnum = "LOOPBACK"
)

var mappingDrgImportPolicyUpdateV2DrgAttachmentTypeEnum = map[string]DrgImportPolicyUpdateV2DrgAttachmentTypeEnum{
	"VCN":                       DrgImportPolicyUpdateV2DrgAttachmentTypeVcn,
	"VIRTUAL_CIRCUIT":           DrgImportPolicyUpdateV2DrgAttachmentTypeVirtualCircuit,
	"REMOTE_PEERING_CONNECTION": DrgImportPolicyUpdateV2DrgAttachmentTypeRemotePeeringConnection,
	"IPSEC_TUNNEL":              DrgImportPolicyUpdateV2DrgAttachmentTypeIpsecTunnel,
	"INTERNAL_ONLY":             DrgImportPolicyUpdateV2DrgAttachmentTypeInternalOnly,
	"LOOPBACK":                  DrgImportPolicyUpdateV2DrgAttachmentTypeLoopback,
}

var mappingDrgImportPolicyUpdateV2DrgAttachmentTypeEnumLowerCase = map[string]DrgImportPolicyUpdateV2DrgAttachmentTypeEnum{
	"vcn":                       DrgImportPolicyUpdateV2DrgAttachmentTypeVcn,
	"virtual_circuit":           DrgImportPolicyUpdateV2DrgAttachmentTypeVirtualCircuit,
	"remote_peering_connection": DrgImportPolicyUpdateV2DrgAttachmentTypeRemotePeeringConnection,
	"ipsec_tunnel":              DrgImportPolicyUpdateV2DrgAttachmentTypeIpsecTunnel,
	"internal_only":             DrgImportPolicyUpdateV2DrgAttachmentTypeInternalOnly,
	"loopback":                  DrgImportPolicyUpdateV2DrgAttachmentTypeLoopback,
}

// GetDrgImportPolicyUpdateV2DrgAttachmentTypeEnumValues Enumerates the set of values for DrgImportPolicyUpdateV2DrgAttachmentTypeEnum
func GetDrgImportPolicyUpdateV2DrgAttachmentTypeEnumValues() []DrgImportPolicyUpdateV2DrgAttachmentTypeEnum {
	values := make([]DrgImportPolicyUpdateV2DrgAttachmentTypeEnum, 0)
	for _, v := range mappingDrgImportPolicyUpdateV2DrgAttachmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrgImportPolicyUpdateV2DrgAttachmentTypeEnumStringValues Enumerates the set of values in String for DrgImportPolicyUpdateV2DrgAttachmentTypeEnum
func GetDrgImportPolicyUpdateV2DrgAttachmentTypeEnumStringValues() []string {
	return []string{
		"VCN",
		"VIRTUAL_CIRCUIT",
		"REMOTE_PEERING_CONNECTION",
		"IPSEC_TUNNEL",
		"INTERNAL_ONLY",
		"LOOPBACK",
	}
}

// GetMappingDrgImportPolicyUpdateV2DrgAttachmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrgImportPolicyUpdateV2DrgAttachmentTypeEnum(val string) (DrgImportPolicyUpdateV2DrgAttachmentTypeEnum, bool) {
	enum, ok := mappingDrgImportPolicyUpdateV2DrgAttachmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
