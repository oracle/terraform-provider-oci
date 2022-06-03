// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InternalVnicAttachment Details of a service VNIC attachment or an attachment of a non-service VNIC to a compute instance.
type InternalVnicAttachment struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the VNIC attachment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC.
	Id *string `mandatory:"true" json:"id"`

	// The current state of a VNIC attachment.
	LifecycleState InternalVnicAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The substrate or anycast IP address of the VNICaaS fleet that the VNIC is attached to.
	SubstrateIp *string `mandatory:"true" json:"substrateIp"`

	// The date and time the VNIC attachment was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The slot number of the VNIC.
	SlotId *int `mandatory:"false" json:"slotId"`

	// Shape of VNIC that is used to allocate resource in the data plane.
	VnicShape InternalVnicAttachmentVnicShapeEnum `mandatory:"false" json:"vnicShape,omitempty"`

	// The instance that a VNIC is attached to
	InstanceId *string `mandatory:"false" json:"instanceId"`

	// Composite key created from SubstrateIp, and data plane IDs of VCN and VNIC
	DataPlaneId *string `mandatory:"false" json:"dataPlaneId"`

	// The availability domain of a VNIC attachment
	InternalAvailabilityDomain *string `mandatory:"false" json:"internalAvailabilityDomain"`

	// The Network Address Translated IP to communicate with internal services
	NatIp *string `mandatory:"false" json:"natIp"`

	// The MAC address of the compute instance
	OverlayMac *string `mandatory:"false" json:"overlayMac"`

	// The tag used internally to identify sending VNIC
	VlanTag *int `mandatory:"false" json:"vlanTag"`

	// Index of NIC that VNIC is attached to (OS boot order)
	NicIndex *int `mandatory:"false" json:"nicIndex"`

	MigrationInfo *MigrationInfo `mandatory:"false" json:"migrationInfo"`

	// Property describing customer facing metrics
	MetadataList []CfmMetadata `mandatory:"false" json:"metadataList"`
}

func (m InternalVnicAttachment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternalVnicAttachment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInternalVnicAttachmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetInternalVnicAttachmentLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingInternalVnicAttachmentVnicShapeEnum(string(m.VnicShape)); !ok && m.VnicShape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VnicShape: %s. Supported values are: %s.", m.VnicShape, strings.Join(GetInternalVnicAttachmentVnicShapeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InternalVnicAttachmentLifecycleStateEnum Enum with underlying type: string
type InternalVnicAttachmentLifecycleStateEnum string

// Set of constants representing the allowable values for InternalVnicAttachmentLifecycleStateEnum
const (
	InternalVnicAttachmentLifecycleStateAttaching InternalVnicAttachmentLifecycleStateEnum = "ATTACHING"
	InternalVnicAttachmentLifecycleStateAttached  InternalVnicAttachmentLifecycleStateEnum = "ATTACHED"
	InternalVnicAttachmentLifecycleStateDetaching InternalVnicAttachmentLifecycleStateEnum = "DETACHING"
	InternalVnicAttachmentLifecycleStateDetached  InternalVnicAttachmentLifecycleStateEnum = "DETACHED"
)

var mappingInternalVnicAttachmentLifecycleStateEnum = map[string]InternalVnicAttachmentLifecycleStateEnum{
	"ATTACHING": InternalVnicAttachmentLifecycleStateAttaching,
	"ATTACHED":  InternalVnicAttachmentLifecycleStateAttached,
	"DETACHING": InternalVnicAttachmentLifecycleStateDetaching,
	"DETACHED":  InternalVnicAttachmentLifecycleStateDetached,
}

var mappingInternalVnicAttachmentLifecycleStateEnumLowerCase = map[string]InternalVnicAttachmentLifecycleStateEnum{
	"attaching": InternalVnicAttachmentLifecycleStateAttaching,
	"attached":  InternalVnicAttachmentLifecycleStateAttached,
	"detaching": InternalVnicAttachmentLifecycleStateDetaching,
	"detached":  InternalVnicAttachmentLifecycleStateDetached,
}

// GetInternalVnicAttachmentLifecycleStateEnumValues Enumerates the set of values for InternalVnicAttachmentLifecycleStateEnum
func GetInternalVnicAttachmentLifecycleStateEnumValues() []InternalVnicAttachmentLifecycleStateEnum {
	values := make([]InternalVnicAttachmentLifecycleStateEnum, 0)
	for _, v := range mappingInternalVnicAttachmentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetInternalVnicAttachmentLifecycleStateEnumStringValues Enumerates the set of values in String for InternalVnicAttachmentLifecycleStateEnum
func GetInternalVnicAttachmentLifecycleStateEnumStringValues() []string {
	return []string{
		"ATTACHING",
		"ATTACHED",
		"DETACHING",
		"DETACHED",
	}
}

// GetMappingInternalVnicAttachmentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternalVnicAttachmentLifecycleStateEnum(val string) (InternalVnicAttachmentLifecycleStateEnum, bool) {
	enum, ok := mappingInternalVnicAttachmentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// InternalVnicAttachmentVnicShapeEnum Enum with underlying type: string
type InternalVnicAttachmentVnicShapeEnum string

// Set of constants representing the allowable values for InternalVnicAttachmentVnicShapeEnum
const (
	InternalVnicAttachmentVnicShapeDynamic                         InternalVnicAttachmentVnicShapeEnum = "DYNAMIC"
	InternalVnicAttachmentVnicShapeFixed0040                       InternalVnicAttachmentVnicShapeEnum = "FIXED0040"
	InternalVnicAttachmentVnicShapeFixed0060                       InternalVnicAttachmentVnicShapeEnum = "FIXED0060"
	InternalVnicAttachmentVnicShapeFixed0060Psm                    InternalVnicAttachmentVnicShapeEnum = "FIXED0060_PSM"
	InternalVnicAttachmentVnicShapeFixed0100                       InternalVnicAttachmentVnicShapeEnum = "FIXED0100"
	InternalVnicAttachmentVnicShapeFixed0120                       InternalVnicAttachmentVnicShapeEnum = "FIXED0120"
	InternalVnicAttachmentVnicShapeFixed01202x                     InternalVnicAttachmentVnicShapeEnum = "FIXED0120_2X"
	InternalVnicAttachmentVnicShapeFixed0200                       InternalVnicAttachmentVnicShapeEnum = "FIXED0200"
	InternalVnicAttachmentVnicShapeFixed0240                       InternalVnicAttachmentVnicShapeEnum = "FIXED0240"
	InternalVnicAttachmentVnicShapeFixed0480                       InternalVnicAttachmentVnicShapeEnum = "FIXED0480"
	InternalVnicAttachmentVnicShapeEntirehost                      InternalVnicAttachmentVnicShapeEnum = "ENTIREHOST"
	InternalVnicAttachmentVnicShapeDynamic25g                      InternalVnicAttachmentVnicShapeEnum = "DYNAMIC_25G"
	InternalVnicAttachmentVnicShapeFixed004025g                    InternalVnicAttachmentVnicShapeEnum = "FIXED0040_25G"
	InternalVnicAttachmentVnicShapeFixed010025g                    InternalVnicAttachmentVnicShapeEnum = "FIXED0100_25G"
	InternalVnicAttachmentVnicShapeFixed020025g                    InternalVnicAttachmentVnicShapeEnum = "FIXED0200_25G"
	InternalVnicAttachmentVnicShapeFixed040025g                    InternalVnicAttachmentVnicShapeEnum = "FIXED0400_25G"
	InternalVnicAttachmentVnicShapeFixed080025g                    InternalVnicAttachmentVnicShapeEnum = "FIXED0800_25G"
	InternalVnicAttachmentVnicShapeFixed160025g                    InternalVnicAttachmentVnicShapeEnum = "FIXED1600_25G"
	InternalVnicAttachmentVnicShapeFixed240025g                    InternalVnicAttachmentVnicShapeEnum = "FIXED2400_25G"
	InternalVnicAttachmentVnicShapeEntirehost25g                   InternalVnicAttachmentVnicShapeEnum = "ENTIREHOST_25G"
	InternalVnicAttachmentVnicShapeDynamicE125g                    InternalVnicAttachmentVnicShapeEnum = "DYNAMIC_E1_25G"
	InternalVnicAttachmentVnicShapeFixed0040E125g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0040_E1_25G"
	InternalVnicAttachmentVnicShapeFixed0070E125g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0070_E1_25G"
	InternalVnicAttachmentVnicShapeFixed0140E125g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0140_E1_25G"
	InternalVnicAttachmentVnicShapeFixed0280E125g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0280_E1_25G"
	InternalVnicAttachmentVnicShapeFixed0560E125g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0560_E1_25G"
	InternalVnicAttachmentVnicShapeFixed1120E125g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1120_E1_25G"
	InternalVnicAttachmentVnicShapeFixed1680E125g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1680_E1_25G"
	InternalVnicAttachmentVnicShapeEntirehostE125g                 InternalVnicAttachmentVnicShapeEnum = "ENTIREHOST_E1_25G"
	InternalVnicAttachmentVnicShapeDynamicB125g                    InternalVnicAttachmentVnicShapeEnum = "DYNAMIC_B1_25G"
	InternalVnicAttachmentVnicShapeFixed0040B125g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0040_B1_25G"
	InternalVnicAttachmentVnicShapeFixed0060B125g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0060_B1_25G"
	InternalVnicAttachmentVnicShapeFixed0120B125g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0120_B1_25G"
	InternalVnicAttachmentVnicShapeFixed0240B125g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0240_B1_25G"
	InternalVnicAttachmentVnicShapeFixed0480B125g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0480_B1_25G"
	InternalVnicAttachmentVnicShapeFixed0960B125g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0960_B1_25G"
	InternalVnicAttachmentVnicShapeEntirehostB125g                 InternalVnicAttachmentVnicShapeEnum = "ENTIREHOST_B1_25G"
	InternalVnicAttachmentVnicShapeMicroVmFixed0048E125g           InternalVnicAttachmentVnicShapeEnum = "MICRO_VM_FIXED0048_E1_25G"
	InternalVnicAttachmentVnicShapeMicroLbFixed0001E125g           InternalVnicAttachmentVnicShapeEnum = "MICRO_LB_FIXED0001_E1_25G"
	InternalVnicAttachmentVnicShapeVnicaasFixed0200                InternalVnicAttachmentVnicShapeEnum = "VNICAAS_FIXED0200"
	InternalVnicAttachmentVnicShapeVnicaasFixed0400                InternalVnicAttachmentVnicShapeEnum = "VNICAAS_FIXED0400"
	InternalVnicAttachmentVnicShapeVnicaasFixed0700                InternalVnicAttachmentVnicShapeEnum = "VNICAAS_FIXED0700"
	InternalVnicAttachmentVnicShapeVnicaasNlbApproved10g           InternalVnicAttachmentVnicShapeEnum = "VNICAAS_NLB_APPROVED_10G"
	InternalVnicAttachmentVnicShapeVnicaasNlbApproved25g           InternalVnicAttachmentVnicShapeEnum = "VNICAAS_NLB_APPROVED_25G"
	InternalVnicAttachmentVnicShapeVnicaasTelesis25g               InternalVnicAttachmentVnicShapeEnum = "VNICAAS_TELESIS_25G"
	InternalVnicAttachmentVnicShapeVnicaasTelesis10g               InternalVnicAttachmentVnicShapeEnum = "VNICAAS_TELESIS_10G"
	InternalVnicAttachmentVnicShapeVnicaasAmbassadorFixed0100      InternalVnicAttachmentVnicShapeEnum = "VNICAAS_AMBASSADOR_FIXED0100"
	InternalVnicAttachmentVnicShapeVnicaasTelesisGamma             InternalVnicAttachmentVnicShapeEnum = "VNICAAS_TELESIS_GAMMA"
	InternalVnicAttachmentVnicShapeVnicaasPrivatedns               InternalVnicAttachmentVnicShapeEnum = "VNICAAS_PRIVATEDNS"
	InternalVnicAttachmentVnicShapeVnicaasFwaas                    InternalVnicAttachmentVnicShapeEnum = "VNICAAS_FWAAS"
	InternalVnicAttachmentVnicShapeVnicaasLbaasFree                InternalVnicAttachmentVnicShapeEnum = "VNICAAS_LBAAS_FREE"
	InternalVnicAttachmentVnicShapeVnicaasLbaas8g512k              InternalVnicAttachmentVnicShapeEnum = "VNICAAS_LBAAS_8G_512K"
	InternalVnicAttachmentVnicShapeVnicaasLbaas8g1m                InternalVnicAttachmentVnicShapeEnum = "VNICAAS_LBAAS_8G_1M"
	InternalVnicAttachmentVnicShapeVnicaasLbaas8g2m                InternalVnicAttachmentVnicShapeEnum = "VNICAAS_LBAAS_8G_2M"
	InternalVnicAttachmentVnicShapeVnicaasLbaas8g3m                InternalVnicAttachmentVnicShapeEnum = "VNICAAS_LBAAS_8G_3M"
	InternalVnicAttachmentVnicShapeVnicaasLbaas8g1m8ghost          InternalVnicAttachmentVnicShapeEnum = "VNICAAS_LBAAS_8G_1M_8GHOST"
	InternalVnicAttachmentVnicShapeVnicaasLbaas8g1m16ghost         InternalVnicAttachmentVnicShapeEnum = "VNICAAS_LBAAS_8G_1M_16GHOST"
	InternalVnicAttachmentVnicShapeDynamicE350g                    InternalVnicAttachmentVnicShapeEnum = "DYNAMIC_E3_50G"
	InternalVnicAttachmentVnicShapeFixed0040E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0040_E3_50G"
	InternalVnicAttachmentVnicShapeFixed0100E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0100_E3_50G"
	InternalVnicAttachmentVnicShapeFixed0200E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0200_E3_50G"
	InternalVnicAttachmentVnicShapeFixed0300E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0300_E3_50G"
	InternalVnicAttachmentVnicShapeFixed0400E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0400_E3_50G"
	InternalVnicAttachmentVnicShapeFixed0500E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0500_E3_50G"
	InternalVnicAttachmentVnicShapeFixed0600E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0600_E3_50G"
	InternalVnicAttachmentVnicShapeFixed0700E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0700_E3_50G"
	InternalVnicAttachmentVnicShapeFixed0800E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0800_E3_50G"
	InternalVnicAttachmentVnicShapeFixed0900E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0900_E3_50G"
	InternalVnicAttachmentVnicShapeFixed1000E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1000_E3_50G"
	InternalVnicAttachmentVnicShapeFixed1100E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1100_E3_50G"
	InternalVnicAttachmentVnicShapeFixed1200E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1200_E3_50G"
	InternalVnicAttachmentVnicShapeFixed1300E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1300_E3_50G"
	InternalVnicAttachmentVnicShapeFixed1400E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1400_E3_50G"
	InternalVnicAttachmentVnicShapeFixed1500E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1500_E3_50G"
	InternalVnicAttachmentVnicShapeFixed1600E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1600_E3_50G"
	InternalVnicAttachmentVnicShapeFixed1700E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1700_E3_50G"
	InternalVnicAttachmentVnicShapeFixed1800E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1800_E3_50G"
	InternalVnicAttachmentVnicShapeFixed1900E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1900_E3_50G"
	InternalVnicAttachmentVnicShapeFixed2000E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2000_E3_50G"
	InternalVnicAttachmentVnicShapeFixed2100E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2100_E3_50G"
	InternalVnicAttachmentVnicShapeFixed2200E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2200_E3_50G"
	InternalVnicAttachmentVnicShapeFixed2300E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2300_E3_50G"
	InternalVnicAttachmentVnicShapeFixed2400E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2400_E3_50G"
	InternalVnicAttachmentVnicShapeFixed2500E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2500_E3_50G"
	InternalVnicAttachmentVnicShapeFixed2600E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2600_E3_50G"
	InternalVnicAttachmentVnicShapeFixed2700E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2700_E3_50G"
	InternalVnicAttachmentVnicShapeFixed2800E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2800_E3_50G"
	InternalVnicAttachmentVnicShapeFixed2900E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2900_E3_50G"
	InternalVnicAttachmentVnicShapeFixed3000E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3000_E3_50G"
	InternalVnicAttachmentVnicShapeFixed3100E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3100_E3_50G"
	InternalVnicAttachmentVnicShapeFixed3200E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3200_E3_50G"
	InternalVnicAttachmentVnicShapeFixed3300E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3300_E3_50G"
	InternalVnicAttachmentVnicShapeFixed3400E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3400_E3_50G"
	InternalVnicAttachmentVnicShapeFixed3500E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3500_E3_50G"
	InternalVnicAttachmentVnicShapeFixed3600E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3600_E3_50G"
	InternalVnicAttachmentVnicShapeFixed3700E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3700_E3_50G"
	InternalVnicAttachmentVnicShapeFixed3800E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3800_E3_50G"
	InternalVnicAttachmentVnicShapeFixed3900E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3900_E3_50G"
	InternalVnicAttachmentVnicShapeFixed4000E350g                  InternalVnicAttachmentVnicShapeEnum = "FIXED4000_E3_50G"
	InternalVnicAttachmentVnicShapeEntirehostE350g                 InternalVnicAttachmentVnicShapeEnum = "ENTIREHOST_E3_50G"
	InternalVnicAttachmentVnicShapeDynamicE450g                    InternalVnicAttachmentVnicShapeEnum = "DYNAMIC_E4_50G"
	InternalVnicAttachmentVnicShapeFixed0040E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0040_E4_50G"
	InternalVnicAttachmentVnicShapeFixed0100E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0100_E4_50G"
	InternalVnicAttachmentVnicShapeFixed0200E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0200_E4_50G"
	InternalVnicAttachmentVnicShapeFixed0300E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0300_E4_50G"
	InternalVnicAttachmentVnicShapeFixed0400E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0400_E4_50G"
	InternalVnicAttachmentVnicShapeFixed0500E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0500_E4_50G"
	InternalVnicAttachmentVnicShapeFixed0600E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0600_E4_50G"
	InternalVnicAttachmentVnicShapeFixed0700E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0700_E4_50G"
	InternalVnicAttachmentVnicShapeFixed0800E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0800_E4_50G"
	InternalVnicAttachmentVnicShapeFixed0900E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0900_E4_50G"
	InternalVnicAttachmentVnicShapeFixed1000E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1000_E4_50G"
	InternalVnicAttachmentVnicShapeFixed1100E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1100_E4_50G"
	InternalVnicAttachmentVnicShapeFixed1200E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1200_E4_50G"
	InternalVnicAttachmentVnicShapeFixed1300E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1300_E4_50G"
	InternalVnicAttachmentVnicShapeFixed1400E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1400_E4_50G"
	InternalVnicAttachmentVnicShapeFixed1500E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1500_E4_50G"
	InternalVnicAttachmentVnicShapeFixed1600E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1600_E4_50G"
	InternalVnicAttachmentVnicShapeFixed1700E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1700_E4_50G"
	InternalVnicAttachmentVnicShapeFixed1800E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1800_E4_50G"
	InternalVnicAttachmentVnicShapeFixed1900E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1900_E4_50G"
	InternalVnicAttachmentVnicShapeFixed2000E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2000_E4_50G"
	InternalVnicAttachmentVnicShapeFixed2100E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2100_E4_50G"
	InternalVnicAttachmentVnicShapeFixed2200E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2200_E4_50G"
	InternalVnicAttachmentVnicShapeFixed2300E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2300_E4_50G"
	InternalVnicAttachmentVnicShapeFixed2400E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2400_E4_50G"
	InternalVnicAttachmentVnicShapeFixed2500E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2500_E4_50G"
	InternalVnicAttachmentVnicShapeFixed2600E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2600_E4_50G"
	InternalVnicAttachmentVnicShapeFixed2700E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2700_E4_50G"
	InternalVnicAttachmentVnicShapeFixed2800E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2800_E4_50G"
	InternalVnicAttachmentVnicShapeFixed2900E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2900_E4_50G"
	InternalVnicAttachmentVnicShapeFixed3000E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3000_E4_50G"
	InternalVnicAttachmentVnicShapeFixed3100E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3100_E4_50G"
	InternalVnicAttachmentVnicShapeFixed3200E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3200_E4_50G"
	InternalVnicAttachmentVnicShapeFixed3300E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3300_E4_50G"
	InternalVnicAttachmentVnicShapeFixed3400E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3400_E4_50G"
	InternalVnicAttachmentVnicShapeFixed3500E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3500_E4_50G"
	InternalVnicAttachmentVnicShapeFixed3600E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3600_E4_50G"
	InternalVnicAttachmentVnicShapeFixed3700E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3700_E4_50G"
	InternalVnicAttachmentVnicShapeFixed3800E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3800_E4_50G"
	InternalVnicAttachmentVnicShapeFixed3900E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3900_E4_50G"
	InternalVnicAttachmentVnicShapeFixed4000E450g                  InternalVnicAttachmentVnicShapeEnum = "FIXED4000_E4_50G"
	InternalVnicAttachmentVnicShapeEntirehostE450g                 InternalVnicAttachmentVnicShapeEnum = "ENTIREHOST_E4_50G"
	InternalVnicAttachmentVnicShapeMicroVmFixed0050E350g           InternalVnicAttachmentVnicShapeEnum = "Micro_VM_Fixed0050_E3_50G"
	InternalVnicAttachmentVnicShapeMicroVmFixed0050E450g           InternalVnicAttachmentVnicShapeEnum = "Micro_VM_Fixed0050_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0025E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0025_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0050E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0050_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0075E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0075_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0100E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0100_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0125E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0125_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0150E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0150_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0175E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0175_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0200E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0200_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0225E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0225_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0250E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0250_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0275E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0275_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0300E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0300_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0325E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0325_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0350E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0350_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0375E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0375_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0400E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0400_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0425E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0425_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0450E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0450_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0475E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0475_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0500E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0500_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0525E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0525_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0550E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0550_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0575E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0575_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0600E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0600_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0625E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0625_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0650E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0650_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0675E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0675_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0700E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0700_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0725E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0725_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0750E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0750_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0775E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0775_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0800E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0800_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0825E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0825_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0850E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0850_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0875E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0875_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0900E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0900_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0925E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0925_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0950E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0950_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0975E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0975_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1000E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1000_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1025E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1025_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1050E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1050_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1075E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1075_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1100E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1100_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1125E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1125_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1150E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1150_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1175E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1175_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1200E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1200_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1225E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1225_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1250E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1250_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1275E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1275_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1300E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1300_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1325E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1325_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1350E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1350_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1375E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1375_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1400E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1400_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1425E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1425_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1450E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1450_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1475E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1475_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1500E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1500_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1525E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1525_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1550E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1550_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1575E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1575_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1600E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1600_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1625E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1625_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1650E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1650_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1700E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1700_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1725E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1725_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1750E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1750_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1800E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1800_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1850E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1850_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1875E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1875_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1900E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1900_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1925E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1925_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1950E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1950_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2000E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2000_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2025E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2025_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2050E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2050_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2100E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2100_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2125E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2125_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2150E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2150_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2175E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2175_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2200E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2200_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2250E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2250_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2275E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2275_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2300E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2300_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2325E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2325_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2350E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2350_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2375E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2375_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2400E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2400_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2450E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2450_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2475E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2475_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2500E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2500_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2550E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2550_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2600E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2600_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2625E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2625_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2650E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2650_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2700E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2700_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2750E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2750_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2775E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2775_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2800E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2800_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2850E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2850_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2875E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2875_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2900E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2900_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2925E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2925_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2950E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2950_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2975E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2975_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3000E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3000_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3025E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3025_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3050E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3050_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3075E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3075_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3100E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3100_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3125E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3125_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3150E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3150_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3200E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3200_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3225E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3225_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3250E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3250_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3300E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3300_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3325E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3325_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3375E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3375_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3400E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3400_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3450E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3450_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3500E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3500_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3525E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3525_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3575E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3575_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3600E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3600_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3625E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3625_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3675E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3675_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3700E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3700_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3750E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3750_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3800E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3800_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3825E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3825_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3850E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3850_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3875E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3875_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3900E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3900_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3975E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3975_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4000E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4000_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4025E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4025_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4050E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4050_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4100E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4100_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4125E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4125_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4200E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4200_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4225E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4225_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4250E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4250_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4275E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4275_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4300E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4300_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4350E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4350_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4375E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4375_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4400E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4400_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4425E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4425_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4500E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4500_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4550E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4550_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4575E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4575_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4600E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4600_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4625E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4625_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4650E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4650_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4675E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4675_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4700E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4700_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4725E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4725_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4750E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4750_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4800E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4800_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4875E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4875_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4900E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4900_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4950E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4950_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed5000E350g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED5000_E3_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0025E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0025_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0050E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0050_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0075E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0075_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0100E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0100_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0125E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0125_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0150E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0150_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0175E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0175_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0200E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0200_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0225E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0225_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0250E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0250_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0275E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0275_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0300E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0300_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0325E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0325_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0350E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0350_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0375E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0375_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0400E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0400_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0425E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0425_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0450E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0450_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0475E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0475_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0500E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0500_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0525E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0525_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0550E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0550_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0575E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0575_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0600E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0600_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0625E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0625_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0650E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0650_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0675E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0675_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0700E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0700_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0725E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0725_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0750E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0750_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0775E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0775_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0800E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0800_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0825E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0825_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0850E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0850_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0875E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0875_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0900E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0900_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0925E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0925_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0950E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0950_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0975E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0975_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1000E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1000_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1025E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1025_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1050E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1050_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1075E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1075_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1100E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1100_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1125E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1125_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1150E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1150_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1175E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1175_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1200E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1200_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1225E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1225_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1250E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1250_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1275E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1275_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1300E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1300_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1325E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1325_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1350E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1350_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1375E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1375_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1400E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1400_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1425E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1425_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1450E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1450_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1475E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1475_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1500E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1500_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1525E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1525_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1550E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1550_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1575E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1575_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1600E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1600_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1625E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1625_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1650E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1650_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1700E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1700_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1725E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1725_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1750E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1750_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1800E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1800_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1850E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1850_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1875E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1875_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1900E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1900_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1925E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1925_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1950E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1950_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2000E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2000_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2025E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2025_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2050E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2050_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2100E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2100_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2125E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2125_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2150E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2150_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2175E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2175_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2200E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2200_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2250E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2250_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2275E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2275_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2300E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2300_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2325E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2325_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2350E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2350_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2375E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2375_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2400E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2400_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2450E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2450_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2475E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2475_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2500E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2500_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2550E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2550_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2600E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2600_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2625E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2625_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2650E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2650_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2700E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2700_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2750E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2750_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2775E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2775_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2800E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2800_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2850E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2850_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2875E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2875_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2900E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2900_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2925E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2925_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2950E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2950_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2975E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2975_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3000E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3000_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3025E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3025_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3050E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3050_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3075E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3075_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3100E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3100_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3125E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3125_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3150E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3150_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3200E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3200_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3225E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3225_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3250E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3250_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3300E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3300_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3325E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3325_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3375E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3375_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3400E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3400_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3450E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3450_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3500E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3500_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3525E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3525_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3575E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3575_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3600E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3600_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3625E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3625_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3675E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3675_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3700E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3700_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3750E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3750_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3800E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3800_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3825E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3825_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3850E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3850_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3875E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3875_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3900E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3900_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3975E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3975_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4000E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4000_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4025E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4025_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4050E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4050_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4100E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4100_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4125E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4125_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4200E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4200_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4225E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4225_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4250E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4250_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4275E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4275_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4300E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4300_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4350E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4350_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4375E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4375_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4400E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4400_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4425E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4425_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4500E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4500_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4550E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4550_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4575E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4575_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4600E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4600_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4625E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4625_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4650E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4650_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4675E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4675_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4700E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4700_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4725E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4725_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4750E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4750_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4800E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4800_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4875E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4875_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4900E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4900_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4950E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4950_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed5000E450g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED5000_E4_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0020A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0020_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0040A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0040_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0060A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0060_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0080A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0080_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0100A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0100_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0120A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0120_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0140A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0140_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0160A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0160_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0180A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0180_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0200A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0200_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0220A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0220_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0240A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0240_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0260A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0260_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0280A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0280_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0300A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0300_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0320A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0320_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0340A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0340_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0360A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0360_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0380A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0380_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0400A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0400_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0420A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0420_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0440A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0440_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0460A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0460_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0480A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0480_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0500A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0500_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0520A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0520_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0540A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0540_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0560A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0560_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0580A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0580_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0600A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0600_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0620A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0620_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0640A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0640_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0660A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0660_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0680A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0680_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0700A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0700_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0720A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0720_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0740A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0740_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0760A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0760_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0780A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0780_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0800A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0800_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0820A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0820_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0840A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0840_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0860A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0860_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0880A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0880_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0900A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0900_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0920A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0920_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0940A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0940_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0960A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0960_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0980A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0980_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1000A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1000_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1020A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1020_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1040A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1040_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1060A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1060_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1080A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1080_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1100A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1100_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1120A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1120_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1140A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1140_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1160A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1160_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1180A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1180_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1200A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1200_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1220A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1220_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1240A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1240_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1260A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1260_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1280A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1280_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1300A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1300_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1320A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1320_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1340A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1340_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1360A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1360_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1380A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1380_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1400A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1400_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1420A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1420_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1440A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1440_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1460A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1460_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1480A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1480_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1500A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1500_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1520A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1520_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1540A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1540_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1560A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1560_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1580A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1580_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1600A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1600_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1620A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1620_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1640A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1640_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1660A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1660_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1680A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1680_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1700A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1700_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1720A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1720_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1740A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1740_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1760A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1760_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1780A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1780_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1800A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1800_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1820A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1820_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1840A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1840_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1860A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1860_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1880A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1880_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1900A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1900_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1920A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1920_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1940A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1940_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1960A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1960_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1980A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1980_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2000A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2000_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2020A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2020_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2040A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2040_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2060A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2060_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2080A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2080_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2100A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2100_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2120A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2120_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2140A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2140_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2160A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2160_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2180A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2180_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2200A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2200_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2220A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2220_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2240A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2240_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2260A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2260_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2280A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2280_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2300A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2300_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2320A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2320_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2340A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2340_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2360A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2360_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2380A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2380_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2400A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2400_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2420A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2420_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2440A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2440_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2460A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2460_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2480A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2480_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2500A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2500_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2520A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2520_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2540A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2540_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2560A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2560_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2580A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2580_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2600A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2600_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2620A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2620_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2640A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2640_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2660A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2660_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2680A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2680_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2700A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2700_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2720A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2720_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2740A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2740_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2760A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2760_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2780A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2780_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2800A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2800_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2820A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2820_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2840A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2840_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2860A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2860_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2880A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2880_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2900A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2900_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2920A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2920_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2940A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2940_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2960A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2960_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2980A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2980_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3000A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3000_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3020A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3020_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3040A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3040_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3060A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3060_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3080A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3080_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3100A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3100_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3120A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3120_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3140A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3140_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3160A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3160_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3180A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3180_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3200A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3200_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3220A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3220_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3240A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3240_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3260A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3260_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3280A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3280_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3300A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3300_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3320A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3320_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3340A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3340_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3360A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3360_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3380A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3380_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3400A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3400_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3420A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3420_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3440A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3440_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3460A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3460_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3480A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3480_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3500A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3500_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3520A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3520_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3540A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3540_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3560A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3560_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3580A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3580_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3600A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3600_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3620A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3620_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3640A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3640_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3660A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3660_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3680A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3680_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3700A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3700_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3720A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3720_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3740A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3740_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3760A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3760_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3780A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3780_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3800A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3800_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3820A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3820_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3840A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3840_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3860A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3860_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3880A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3880_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3900A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3900_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3920A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3920_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3940A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3940_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3960A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3960_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3980A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3980_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4000A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4000_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4020A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4020_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4040A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4040_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4060A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4060_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4080A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4080_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4100A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4100_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4120A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4120_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4140A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4140_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4160A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4160_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4180A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4180_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4200A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4200_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4220A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4220_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4240A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4240_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4260A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4260_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4280A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4280_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4300A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4300_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4320A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4320_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4340A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4340_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4360A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4360_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4380A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4380_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4400A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4400_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4420A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4420_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4440A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4440_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4460A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4460_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4480A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4480_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4500A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4500_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4520A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4520_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4540A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4540_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4560A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4560_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4580A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4580_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4600A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4600_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4620A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4620_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4640A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4640_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4660A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4660_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4680A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4680_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4700A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4700_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4720A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4720_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4740A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4740_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4760A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4760_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4780A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4780_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4800A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4800_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4820A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4820_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4840A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4840_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4860A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4860_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4880A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4880_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4900A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4900_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4920A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4920_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4940A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4940_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4960A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4960_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4980A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4980_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed5000A150g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED5000_A1_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0090X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0090_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0180X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0180_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0270X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0270_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0360X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0360_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0450X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0450_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0540X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0540_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0630X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0630_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0720X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0720_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0810X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0810_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0900X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0900_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed0990X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED0990_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1080X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1080_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1170X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1170_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1260X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1260_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1350X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1350_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1440X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1440_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1530X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1530_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1620X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1620_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1710X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1710_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1800X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1800_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1890X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1890_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed1980X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED1980_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2070X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2070_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2160X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2160_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2250X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2250_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2340X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2340_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2430X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2430_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2520X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2520_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2610X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2610_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2700X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2700_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2790X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2790_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2880X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2880_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed2970X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED2970_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3060X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3060_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3150X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3150_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3240X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3240_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3330X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3330_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3420X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3420_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3510X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3510_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3600X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3600_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3690X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3690_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3780X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3780_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3870X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3870_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed3960X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED3960_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4050X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4050_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4140X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4140_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4230X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4230_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4320X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4320_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4410X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4410_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4500X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4500_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4590X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4590_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4680X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4680_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4770X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4770_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4860X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4860_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreVmFixed4950X950g         InternalVnicAttachmentVnicShapeEnum = "SUBCORE_VM_FIXED4950_X9_50G"
	InternalVnicAttachmentVnicShapeDynamicA150g                    InternalVnicAttachmentVnicShapeEnum = "DYNAMIC_A1_50G"
	InternalVnicAttachmentVnicShapeFixed0040A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0040_A1_50G"
	InternalVnicAttachmentVnicShapeFixed0100A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0100_A1_50G"
	InternalVnicAttachmentVnicShapeFixed0200A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0200_A1_50G"
	InternalVnicAttachmentVnicShapeFixed0300A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0300_A1_50G"
	InternalVnicAttachmentVnicShapeFixed0400A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0400_A1_50G"
	InternalVnicAttachmentVnicShapeFixed0500A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0500_A1_50G"
	InternalVnicAttachmentVnicShapeFixed0600A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0600_A1_50G"
	InternalVnicAttachmentVnicShapeFixed0700A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0700_A1_50G"
	InternalVnicAttachmentVnicShapeFixed0800A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0800_A1_50G"
	InternalVnicAttachmentVnicShapeFixed0900A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0900_A1_50G"
	InternalVnicAttachmentVnicShapeFixed1000A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1000_A1_50G"
	InternalVnicAttachmentVnicShapeFixed1100A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1100_A1_50G"
	InternalVnicAttachmentVnicShapeFixed1200A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1200_A1_50G"
	InternalVnicAttachmentVnicShapeFixed1300A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1300_A1_50G"
	InternalVnicAttachmentVnicShapeFixed1400A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1400_A1_50G"
	InternalVnicAttachmentVnicShapeFixed1500A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1500_A1_50G"
	InternalVnicAttachmentVnicShapeFixed1600A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1600_A1_50G"
	InternalVnicAttachmentVnicShapeFixed1700A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1700_A1_50G"
	InternalVnicAttachmentVnicShapeFixed1800A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1800_A1_50G"
	InternalVnicAttachmentVnicShapeFixed1900A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1900_A1_50G"
	InternalVnicAttachmentVnicShapeFixed2000A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2000_A1_50G"
	InternalVnicAttachmentVnicShapeFixed2100A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2100_A1_50G"
	InternalVnicAttachmentVnicShapeFixed2200A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2200_A1_50G"
	InternalVnicAttachmentVnicShapeFixed2300A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2300_A1_50G"
	InternalVnicAttachmentVnicShapeFixed2400A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2400_A1_50G"
	InternalVnicAttachmentVnicShapeFixed2500A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2500_A1_50G"
	InternalVnicAttachmentVnicShapeFixed2600A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2600_A1_50G"
	InternalVnicAttachmentVnicShapeFixed2700A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2700_A1_50G"
	InternalVnicAttachmentVnicShapeFixed2800A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2800_A1_50G"
	InternalVnicAttachmentVnicShapeFixed2900A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2900_A1_50G"
	InternalVnicAttachmentVnicShapeFixed3000A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3000_A1_50G"
	InternalVnicAttachmentVnicShapeFixed3100A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3100_A1_50G"
	InternalVnicAttachmentVnicShapeFixed3200A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3200_A1_50G"
	InternalVnicAttachmentVnicShapeFixed3300A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3300_A1_50G"
	InternalVnicAttachmentVnicShapeFixed3400A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3400_A1_50G"
	InternalVnicAttachmentVnicShapeFixed3500A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3500_A1_50G"
	InternalVnicAttachmentVnicShapeFixed3600A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3600_A1_50G"
	InternalVnicAttachmentVnicShapeFixed3700A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3700_A1_50G"
	InternalVnicAttachmentVnicShapeFixed3800A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3800_A1_50G"
	InternalVnicAttachmentVnicShapeFixed3900A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3900_A1_50G"
	InternalVnicAttachmentVnicShapeFixed4000A150g                  InternalVnicAttachmentVnicShapeEnum = "FIXED4000_A1_50G"
	InternalVnicAttachmentVnicShapeEntirehostA150g                 InternalVnicAttachmentVnicShapeEnum = "ENTIREHOST_A1_50G"
	InternalVnicAttachmentVnicShapeDynamicX950g                    InternalVnicAttachmentVnicShapeEnum = "DYNAMIC_X9_50G"
	InternalVnicAttachmentVnicShapeFixed0040X950g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0040_X9_50G"
	InternalVnicAttachmentVnicShapeFixed0400X950g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0400_X9_50G"
	InternalVnicAttachmentVnicShapeFixed0800X950g                  InternalVnicAttachmentVnicShapeEnum = "FIXED0800_X9_50G"
	InternalVnicAttachmentVnicShapeFixed1200X950g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1200_X9_50G"
	InternalVnicAttachmentVnicShapeFixed1600X950g                  InternalVnicAttachmentVnicShapeEnum = "FIXED1600_X9_50G"
	InternalVnicAttachmentVnicShapeFixed2000X950g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2000_X9_50G"
	InternalVnicAttachmentVnicShapeFixed2400X950g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2400_X9_50G"
	InternalVnicAttachmentVnicShapeFixed2800X950g                  InternalVnicAttachmentVnicShapeEnum = "FIXED2800_X9_50G"
	InternalVnicAttachmentVnicShapeFixed3200X950g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3200_X9_50G"
	InternalVnicAttachmentVnicShapeFixed3600X950g                  InternalVnicAttachmentVnicShapeEnum = "FIXED3600_X9_50G"
	InternalVnicAttachmentVnicShapeFixed4000X950g                  InternalVnicAttachmentVnicShapeEnum = "FIXED4000_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed0100X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED0100_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed0200X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED0200_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed0300X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED0300_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed0400X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED0400_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed0500X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED0500_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed0600X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED0600_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed0700X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED0700_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed0800X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED0800_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed0900X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED0900_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed1000X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED1000_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed1100X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED1100_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed1200X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED1200_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed1300X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED1300_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed1400X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED1400_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed1500X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED1500_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed1600X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED1600_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed1700X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED1700_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed1800X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED1800_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed1900X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED1900_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed2000X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED2000_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed2100X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED2100_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed2200X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED2200_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed2300X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED2300_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed2400X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED2400_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed2500X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED2500_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed2600X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED2600_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed2700X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED2700_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed2800X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED2800_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed2900X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED2900_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed3000X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED3000_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed3100X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED3100_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed3200X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED3200_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed3300X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED3300_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed3400X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED3400_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed3500X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED3500_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed3600X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED3600_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed3700X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED3700_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed3800X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED3800_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed3900X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED3900_X9_50G"
	InternalVnicAttachmentVnicShapeStandardVmFixed4000X950g        InternalVnicAttachmentVnicShapeEnum = "STANDARD_VM_FIXED4000_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0025X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0025_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0050X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0050_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0075X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0075_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0100X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0100_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0125X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0125_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0150X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0150_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0175X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0175_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0200X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0200_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0225X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0225_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0250X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0250_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0275X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0275_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0300X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0300_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0325X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0325_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0350X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0350_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0375X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0375_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0400X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0400_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0425X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0425_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0450X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0450_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0475X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0475_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0500X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0500_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0525X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0525_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0550X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0550_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0575X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0575_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0600X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0600_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0625X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0625_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0650X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0650_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0675X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0675_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0700X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0700_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0725X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0725_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0750X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0750_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0775X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0775_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0800X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0800_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0825X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0825_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0850X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0850_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0875X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0875_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0900X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0900_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0925X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0925_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0950X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0950_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0975X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0975_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1000X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1000_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1025X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1025_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1050X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1050_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1075X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1075_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1100X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1100_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1125X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1125_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1150X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1150_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1175X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1175_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1200X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1200_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1225X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1225_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1250X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1250_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1275X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1275_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1300X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1300_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1325X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1325_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1350X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1350_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1375X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1375_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1400X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1400_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1425X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1425_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1450X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1450_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1475X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1475_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1500X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1500_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1525X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1525_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1550X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1550_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1575X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1575_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1600X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1600_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1625X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1625_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1650X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1650_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1700X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1700_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1725X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1725_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1750X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1750_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1800X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1800_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1850X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1850_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1875X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1875_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1900X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1900_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1925X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1925_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1950X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1950_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2000X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2000_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2025X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2025_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2050X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2050_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2100X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2100_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2125X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2125_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2150X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2150_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2175X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2175_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2200X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2200_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2250X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2250_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2275X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2275_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2300X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2300_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2325X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2325_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2350X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2350_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2375X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2375_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2400X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2400_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2450X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2450_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2475X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2475_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2500X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2500_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2550X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2550_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2600X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2600_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2625X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2625_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2650X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2650_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2700X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2700_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2750X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2750_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2775X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2775_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2800X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2800_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2850X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2850_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2875X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2875_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2900X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2900_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2925X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2925_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2950X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2950_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2975X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2975_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3000X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3000_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3025X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3025_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3050X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3050_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3075X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3075_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3100X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3100_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3125X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3125_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3150X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3150_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3200X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3200_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3225X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3225_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3250X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3250_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3300X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3300_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3325X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3325_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3375X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3375_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3400X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3400_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3450X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3450_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3500X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3500_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3525X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3525_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3575X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3575_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3600X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3600_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3625X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3625_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3675X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3675_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3700X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3700_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3750X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3750_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3800X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3800_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3825X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3825_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3850X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3850_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3875X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3875_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3900X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3900_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3975X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3975_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4000X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4000_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4025X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4025_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4050X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4050_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4100X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4100_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4125X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4125_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4200X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4200_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4225X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4225_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4250X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4250_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4275X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4275_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4300X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4300_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4350X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4350_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4375X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4375_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4400X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4400_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4425X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4425_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4500X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4500_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4550X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4550_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4575X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4575_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4600X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4600_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4625X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4625_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4650X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4650_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4675X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4675_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4700X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4700_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4725X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4725_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4750X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4750_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4800X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4800_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4875X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4875_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4900X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4900_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4950X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4950_X9_50G"
	InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed5000X950g InternalVnicAttachmentVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED5000_X9_50G"
	InternalVnicAttachmentVnicShapeEntirehostX950g                 InternalVnicAttachmentVnicShapeEnum = "ENTIREHOST_X9_50G"
)

var mappingInternalVnicAttachmentVnicShapeEnum = map[string]InternalVnicAttachmentVnicShapeEnum{
	"DYNAMIC":                              InternalVnicAttachmentVnicShapeDynamic,
	"FIXED0040":                            InternalVnicAttachmentVnicShapeFixed0040,
	"FIXED0060":                            InternalVnicAttachmentVnicShapeFixed0060,
	"FIXED0060_PSM":                        InternalVnicAttachmentVnicShapeFixed0060Psm,
	"FIXED0100":                            InternalVnicAttachmentVnicShapeFixed0100,
	"FIXED0120":                            InternalVnicAttachmentVnicShapeFixed0120,
	"FIXED0120_2X":                         InternalVnicAttachmentVnicShapeFixed01202x,
	"FIXED0200":                            InternalVnicAttachmentVnicShapeFixed0200,
	"FIXED0240":                            InternalVnicAttachmentVnicShapeFixed0240,
	"FIXED0480":                            InternalVnicAttachmentVnicShapeFixed0480,
	"ENTIREHOST":                           InternalVnicAttachmentVnicShapeEntirehost,
	"DYNAMIC_25G":                          InternalVnicAttachmentVnicShapeDynamic25g,
	"FIXED0040_25G":                        InternalVnicAttachmentVnicShapeFixed004025g,
	"FIXED0100_25G":                        InternalVnicAttachmentVnicShapeFixed010025g,
	"FIXED0200_25G":                        InternalVnicAttachmentVnicShapeFixed020025g,
	"FIXED0400_25G":                        InternalVnicAttachmentVnicShapeFixed040025g,
	"FIXED0800_25G":                        InternalVnicAttachmentVnicShapeFixed080025g,
	"FIXED1600_25G":                        InternalVnicAttachmentVnicShapeFixed160025g,
	"FIXED2400_25G":                        InternalVnicAttachmentVnicShapeFixed240025g,
	"ENTIREHOST_25G":                       InternalVnicAttachmentVnicShapeEntirehost25g,
	"DYNAMIC_E1_25G":                       InternalVnicAttachmentVnicShapeDynamicE125g,
	"FIXED0040_E1_25G":                     InternalVnicAttachmentVnicShapeFixed0040E125g,
	"FIXED0070_E1_25G":                     InternalVnicAttachmentVnicShapeFixed0070E125g,
	"FIXED0140_E1_25G":                     InternalVnicAttachmentVnicShapeFixed0140E125g,
	"FIXED0280_E1_25G":                     InternalVnicAttachmentVnicShapeFixed0280E125g,
	"FIXED0560_E1_25G":                     InternalVnicAttachmentVnicShapeFixed0560E125g,
	"FIXED1120_E1_25G":                     InternalVnicAttachmentVnicShapeFixed1120E125g,
	"FIXED1680_E1_25G":                     InternalVnicAttachmentVnicShapeFixed1680E125g,
	"ENTIREHOST_E1_25G":                    InternalVnicAttachmentVnicShapeEntirehostE125g,
	"DYNAMIC_B1_25G":                       InternalVnicAttachmentVnicShapeDynamicB125g,
	"FIXED0040_B1_25G":                     InternalVnicAttachmentVnicShapeFixed0040B125g,
	"FIXED0060_B1_25G":                     InternalVnicAttachmentVnicShapeFixed0060B125g,
	"FIXED0120_B1_25G":                     InternalVnicAttachmentVnicShapeFixed0120B125g,
	"FIXED0240_B1_25G":                     InternalVnicAttachmentVnicShapeFixed0240B125g,
	"FIXED0480_B1_25G":                     InternalVnicAttachmentVnicShapeFixed0480B125g,
	"FIXED0960_B1_25G":                     InternalVnicAttachmentVnicShapeFixed0960B125g,
	"ENTIREHOST_B1_25G":                    InternalVnicAttachmentVnicShapeEntirehostB125g,
	"MICRO_VM_FIXED0048_E1_25G":            InternalVnicAttachmentVnicShapeMicroVmFixed0048E125g,
	"MICRO_LB_FIXED0001_E1_25G":            InternalVnicAttachmentVnicShapeMicroLbFixed0001E125g,
	"VNICAAS_FIXED0200":                    InternalVnicAttachmentVnicShapeVnicaasFixed0200,
	"VNICAAS_FIXED0400":                    InternalVnicAttachmentVnicShapeVnicaasFixed0400,
	"VNICAAS_FIXED0700":                    InternalVnicAttachmentVnicShapeVnicaasFixed0700,
	"VNICAAS_NLB_APPROVED_10G":             InternalVnicAttachmentVnicShapeVnicaasNlbApproved10g,
	"VNICAAS_NLB_APPROVED_25G":             InternalVnicAttachmentVnicShapeVnicaasNlbApproved25g,
	"VNICAAS_TELESIS_25G":                  InternalVnicAttachmentVnicShapeVnicaasTelesis25g,
	"VNICAAS_TELESIS_10G":                  InternalVnicAttachmentVnicShapeVnicaasTelesis10g,
	"VNICAAS_AMBASSADOR_FIXED0100":         InternalVnicAttachmentVnicShapeVnicaasAmbassadorFixed0100,
	"VNICAAS_TELESIS_GAMMA":                InternalVnicAttachmentVnicShapeVnicaasTelesisGamma,
	"VNICAAS_PRIVATEDNS":                   InternalVnicAttachmentVnicShapeVnicaasPrivatedns,
	"VNICAAS_FWAAS":                        InternalVnicAttachmentVnicShapeVnicaasFwaas,
	"VNICAAS_LBAAS_FREE":                   InternalVnicAttachmentVnicShapeVnicaasLbaasFree,
	"VNICAAS_LBAAS_8G_512K":                InternalVnicAttachmentVnicShapeVnicaasLbaas8g512k,
	"VNICAAS_LBAAS_8G_1M":                  InternalVnicAttachmentVnicShapeVnicaasLbaas8g1m,
	"VNICAAS_LBAAS_8G_2M":                  InternalVnicAttachmentVnicShapeVnicaasLbaas8g2m,
	"VNICAAS_LBAAS_8G_3M":                  InternalVnicAttachmentVnicShapeVnicaasLbaas8g3m,
	"VNICAAS_LBAAS_8G_1M_8GHOST":           InternalVnicAttachmentVnicShapeVnicaasLbaas8g1m8ghost,
	"VNICAAS_LBAAS_8G_1M_16GHOST":          InternalVnicAttachmentVnicShapeVnicaasLbaas8g1m16ghost,
	"DYNAMIC_E3_50G":                       InternalVnicAttachmentVnicShapeDynamicE350g,
	"FIXED0040_E3_50G":                     InternalVnicAttachmentVnicShapeFixed0040E350g,
	"FIXED0100_E3_50G":                     InternalVnicAttachmentVnicShapeFixed0100E350g,
	"FIXED0200_E3_50G":                     InternalVnicAttachmentVnicShapeFixed0200E350g,
	"FIXED0300_E3_50G":                     InternalVnicAttachmentVnicShapeFixed0300E350g,
	"FIXED0400_E3_50G":                     InternalVnicAttachmentVnicShapeFixed0400E350g,
	"FIXED0500_E3_50G":                     InternalVnicAttachmentVnicShapeFixed0500E350g,
	"FIXED0600_E3_50G":                     InternalVnicAttachmentVnicShapeFixed0600E350g,
	"FIXED0700_E3_50G":                     InternalVnicAttachmentVnicShapeFixed0700E350g,
	"FIXED0800_E3_50G":                     InternalVnicAttachmentVnicShapeFixed0800E350g,
	"FIXED0900_E3_50G":                     InternalVnicAttachmentVnicShapeFixed0900E350g,
	"FIXED1000_E3_50G":                     InternalVnicAttachmentVnicShapeFixed1000E350g,
	"FIXED1100_E3_50G":                     InternalVnicAttachmentVnicShapeFixed1100E350g,
	"FIXED1200_E3_50G":                     InternalVnicAttachmentVnicShapeFixed1200E350g,
	"FIXED1300_E3_50G":                     InternalVnicAttachmentVnicShapeFixed1300E350g,
	"FIXED1400_E3_50G":                     InternalVnicAttachmentVnicShapeFixed1400E350g,
	"FIXED1500_E3_50G":                     InternalVnicAttachmentVnicShapeFixed1500E350g,
	"FIXED1600_E3_50G":                     InternalVnicAttachmentVnicShapeFixed1600E350g,
	"FIXED1700_E3_50G":                     InternalVnicAttachmentVnicShapeFixed1700E350g,
	"FIXED1800_E3_50G":                     InternalVnicAttachmentVnicShapeFixed1800E350g,
	"FIXED1900_E3_50G":                     InternalVnicAttachmentVnicShapeFixed1900E350g,
	"FIXED2000_E3_50G":                     InternalVnicAttachmentVnicShapeFixed2000E350g,
	"FIXED2100_E3_50G":                     InternalVnicAttachmentVnicShapeFixed2100E350g,
	"FIXED2200_E3_50G":                     InternalVnicAttachmentVnicShapeFixed2200E350g,
	"FIXED2300_E3_50G":                     InternalVnicAttachmentVnicShapeFixed2300E350g,
	"FIXED2400_E3_50G":                     InternalVnicAttachmentVnicShapeFixed2400E350g,
	"FIXED2500_E3_50G":                     InternalVnicAttachmentVnicShapeFixed2500E350g,
	"FIXED2600_E3_50G":                     InternalVnicAttachmentVnicShapeFixed2600E350g,
	"FIXED2700_E3_50G":                     InternalVnicAttachmentVnicShapeFixed2700E350g,
	"FIXED2800_E3_50G":                     InternalVnicAttachmentVnicShapeFixed2800E350g,
	"FIXED2900_E3_50G":                     InternalVnicAttachmentVnicShapeFixed2900E350g,
	"FIXED3000_E3_50G":                     InternalVnicAttachmentVnicShapeFixed3000E350g,
	"FIXED3100_E3_50G":                     InternalVnicAttachmentVnicShapeFixed3100E350g,
	"FIXED3200_E3_50G":                     InternalVnicAttachmentVnicShapeFixed3200E350g,
	"FIXED3300_E3_50G":                     InternalVnicAttachmentVnicShapeFixed3300E350g,
	"FIXED3400_E3_50G":                     InternalVnicAttachmentVnicShapeFixed3400E350g,
	"FIXED3500_E3_50G":                     InternalVnicAttachmentVnicShapeFixed3500E350g,
	"FIXED3600_E3_50G":                     InternalVnicAttachmentVnicShapeFixed3600E350g,
	"FIXED3700_E3_50G":                     InternalVnicAttachmentVnicShapeFixed3700E350g,
	"FIXED3800_E3_50G":                     InternalVnicAttachmentVnicShapeFixed3800E350g,
	"FIXED3900_E3_50G":                     InternalVnicAttachmentVnicShapeFixed3900E350g,
	"FIXED4000_E3_50G":                     InternalVnicAttachmentVnicShapeFixed4000E350g,
	"ENTIREHOST_E3_50G":                    InternalVnicAttachmentVnicShapeEntirehostE350g,
	"DYNAMIC_E4_50G":                       InternalVnicAttachmentVnicShapeDynamicE450g,
	"FIXED0040_E4_50G":                     InternalVnicAttachmentVnicShapeFixed0040E450g,
	"FIXED0100_E4_50G":                     InternalVnicAttachmentVnicShapeFixed0100E450g,
	"FIXED0200_E4_50G":                     InternalVnicAttachmentVnicShapeFixed0200E450g,
	"FIXED0300_E4_50G":                     InternalVnicAttachmentVnicShapeFixed0300E450g,
	"FIXED0400_E4_50G":                     InternalVnicAttachmentVnicShapeFixed0400E450g,
	"FIXED0500_E4_50G":                     InternalVnicAttachmentVnicShapeFixed0500E450g,
	"FIXED0600_E4_50G":                     InternalVnicAttachmentVnicShapeFixed0600E450g,
	"FIXED0700_E4_50G":                     InternalVnicAttachmentVnicShapeFixed0700E450g,
	"FIXED0800_E4_50G":                     InternalVnicAttachmentVnicShapeFixed0800E450g,
	"FIXED0900_E4_50G":                     InternalVnicAttachmentVnicShapeFixed0900E450g,
	"FIXED1000_E4_50G":                     InternalVnicAttachmentVnicShapeFixed1000E450g,
	"FIXED1100_E4_50G":                     InternalVnicAttachmentVnicShapeFixed1100E450g,
	"FIXED1200_E4_50G":                     InternalVnicAttachmentVnicShapeFixed1200E450g,
	"FIXED1300_E4_50G":                     InternalVnicAttachmentVnicShapeFixed1300E450g,
	"FIXED1400_E4_50G":                     InternalVnicAttachmentVnicShapeFixed1400E450g,
	"FIXED1500_E4_50G":                     InternalVnicAttachmentVnicShapeFixed1500E450g,
	"FIXED1600_E4_50G":                     InternalVnicAttachmentVnicShapeFixed1600E450g,
	"FIXED1700_E4_50G":                     InternalVnicAttachmentVnicShapeFixed1700E450g,
	"FIXED1800_E4_50G":                     InternalVnicAttachmentVnicShapeFixed1800E450g,
	"FIXED1900_E4_50G":                     InternalVnicAttachmentVnicShapeFixed1900E450g,
	"FIXED2000_E4_50G":                     InternalVnicAttachmentVnicShapeFixed2000E450g,
	"FIXED2100_E4_50G":                     InternalVnicAttachmentVnicShapeFixed2100E450g,
	"FIXED2200_E4_50G":                     InternalVnicAttachmentVnicShapeFixed2200E450g,
	"FIXED2300_E4_50G":                     InternalVnicAttachmentVnicShapeFixed2300E450g,
	"FIXED2400_E4_50G":                     InternalVnicAttachmentVnicShapeFixed2400E450g,
	"FIXED2500_E4_50G":                     InternalVnicAttachmentVnicShapeFixed2500E450g,
	"FIXED2600_E4_50G":                     InternalVnicAttachmentVnicShapeFixed2600E450g,
	"FIXED2700_E4_50G":                     InternalVnicAttachmentVnicShapeFixed2700E450g,
	"FIXED2800_E4_50G":                     InternalVnicAttachmentVnicShapeFixed2800E450g,
	"FIXED2900_E4_50G":                     InternalVnicAttachmentVnicShapeFixed2900E450g,
	"FIXED3000_E4_50G":                     InternalVnicAttachmentVnicShapeFixed3000E450g,
	"FIXED3100_E4_50G":                     InternalVnicAttachmentVnicShapeFixed3100E450g,
	"FIXED3200_E4_50G":                     InternalVnicAttachmentVnicShapeFixed3200E450g,
	"FIXED3300_E4_50G":                     InternalVnicAttachmentVnicShapeFixed3300E450g,
	"FIXED3400_E4_50G":                     InternalVnicAttachmentVnicShapeFixed3400E450g,
	"FIXED3500_E4_50G":                     InternalVnicAttachmentVnicShapeFixed3500E450g,
	"FIXED3600_E4_50G":                     InternalVnicAttachmentVnicShapeFixed3600E450g,
	"FIXED3700_E4_50G":                     InternalVnicAttachmentVnicShapeFixed3700E450g,
	"FIXED3800_E4_50G":                     InternalVnicAttachmentVnicShapeFixed3800E450g,
	"FIXED3900_E4_50G":                     InternalVnicAttachmentVnicShapeFixed3900E450g,
	"FIXED4000_E4_50G":                     InternalVnicAttachmentVnicShapeFixed4000E450g,
	"ENTIREHOST_E4_50G":                    InternalVnicAttachmentVnicShapeEntirehostE450g,
	"Micro_VM_Fixed0050_E3_50G":            InternalVnicAttachmentVnicShapeMicroVmFixed0050E350g,
	"Micro_VM_Fixed0050_E4_50G":            InternalVnicAttachmentVnicShapeMicroVmFixed0050E450g,
	"SUBCORE_VM_FIXED0025_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0025E350g,
	"SUBCORE_VM_FIXED0050_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0050E350g,
	"SUBCORE_VM_FIXED0075_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0075E350g,
	"SUBCORE_VM_FIXED0100_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0100E350g,
	"SUBCORE_VM_FIXED0125_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0125E350g,
	"SUBCORE_VM_FIXED0150_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0150E350g,
	"SUBCORE_VM_FIXED0175_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0175E350g,
	"SUBCORE_VM_FIXED0200_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0200E350g,
	"SUBCORE_VM_FIXED0225_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0225E350g,
	"SUBCORE_VM_FIXED0250_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0250E350g,
	"SUBCORE_VM_FIXED0275_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0275E350g,
	"SUBCORE_VM_FIXED0300_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0300E350g,
	"SUBCORE_VM_FIXED0325_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0325E350g,
	"SUBCORE_VM_FIXED0350_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0350E350g,
	"SUBCORE_VM_FIXED0375_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0375E350g,
	"SUBCORE_VM_FIXED0400_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0400E350g,
	"SUBCORE_VM_FIXED0425_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0425E350g,
	"SUBCORE_VM_FIXED0450_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0450E350g,
	"SUBCORE_VM_FIXED0475_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0475E350g,
	"SUBCORE_VM_FIXED0500_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0500E350g,
	"SUBCORE_VM_FIXED0525_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0525E350g,
	"SUBCORE_VM_FIXED0550_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0550E350g,
	"SUBCORE_VM_FIXED0575_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0575E350g,
	"SUBCORE_VM_FIXED0600_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0600E350g,
	"SUBCORE_VM_FIXED0625_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0625E350g,
	"SUBCORE_VM_FIXED0650_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0650E350g,
	"SUBCORE_VM_FIXED0675_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0675E350g,
	"SUBCORE_VM_FIXED0700_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0700E350g,
	"SUBCORE_VM_FIXED0725_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0725E350g,
	"SUBCORE_VM_FIXED0750_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0750E350g,
	"SUBCORE_VM_FIXED0775_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0775E350g,
	"SUBCORE_VM_FIXED0800_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0800E350g,
	"SUBCORE_VM_FIXED0825_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0825E350g,
	"SUBCORE_VM_FIXED0850_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0850E350g,
	"SUBCORE_VM_FIXED0875_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0875E350g,
	"SUBCORE_VM_FIXED0900_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0900E350g,
	"SUBCORE_VM_FIXED0925_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0925E350g,
	"SUBCORE_VM_FIXED0950_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0950E350g,
	"SUBCORE_VM_FIXED0975_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0975E350g,
	"SUBCORE_VM_FIXED1000_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1000E350g,
	"SUBCORE_VM_FIXED1025_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1025E350g,
	"SUBCORE_VM_FIXED1050_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1050E350g,
	"SUBCORE_VM_FIXED1075_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1075E350g,
	"SUBCORE_VM_FIXED1100_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1100E350g,
	"SUBCORE_VM_FIXED1125_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1125E350g,
	"SUBCORE_VM_FIXED1150_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1150E350g,
	"SUBCORE_VM_FIXED1175_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1175E350g,
	"SUBCORE_VM_FIXED1200_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1200E350g,
	"SUBCORE_VM_FIXED1225_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1225E350g,
	"SUBCORE_VM_FIXED1250_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1250E350g,
	"SUBCORE_VM_FIXED1275_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1275E350g,
	"SUBCORE_VM_FIXED1300_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1300E350g,
	"SUBCORE_VM_FIXED1325_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1325E350g,
	"SUBCORE_VM_FIXED1350_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1350E350g,
	"SUBCORE_VM_FIXED1375_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1375E350g,
	"SUBCORE_VM_FIXED1400_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1400E350g,
	"SUBCORE_VM_FIXED1425_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1425E350g,
	"SUBCORE_VM_FIXED1450_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1450E350g,
	"SUBCORE_VM_FIXED1475_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1475E350g,
	"SUBCORE_VM_FIXED1500_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1500E350g,
	"SUBCORE_VM_FIXED1525_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1525E350g,
	"SUBCORE_VM_FIXED1550_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1550E350g,
	"SUBCORE_VM_FIXED1575_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1575E350g,
	"SUBCORE_VM_FIXED1600_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1600E350g,
	"SUBCORE_VM_FIXED1625_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1625E350g,
	"SUBCORE_VM_FIXED1650_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1650E350g,
	"SUBCORE_VM_FIXED1700_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1700E350g,
	"SUBCORE_VM_FIXED1725_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1725E350g,
	"SUBCORE_VM_FIXED1750_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1750E350g,
	"SUBCORE_VM_FIXED1800_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1800E350g,
	"SUBCORE_VM_FIXED1850_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1850E350g,
	"SUBCORE_VM_FIXED1875_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1875E350g,
	"SUBCORE_VM_FIXED1900_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1900E350g,
	"SUBCORE_VM_FIXED1925_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1925E350g,
	"SUBCORE_VM_FIXED1950_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1950E350g,
	"SUBCORE_VM_FIXED2000_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2000E350g,
	"SUBCORE_VM_FIXED2025_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2025E350g,
	"SUBCORE_VM_FIXED2050_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2050E350g,
	"SUBCORE_VM_FIXED2100_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2100E350g,
	"SUBCORE_VM_FIXED2125_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2125E350g,
	"SUBCORE_VM_FIXED2150_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2150E350g,
	"SUBCORE_VM_FIXED2175_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2175E350g,
	"SUBCORE_VM_FIXED2200_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2200E350g,
	"SUBCORE_VM_FIXED2250_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2250E350g,
	"SUBCORE_VM_FIXED2275_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2275E350g,
	"SUBCORE_VM_FIXED2300_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2300E350g,
	"SUBCORE_VM_FIXED2325_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2325E350g,
	"SUBCORE_VM_FIXED2350_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2350E350g,
	"SUBCORE_VM_FIXED2375_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2375E350g,
	"SUBCORE_VM_FIXED2400_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2400E350g,
	"SUBCORE_VM_FIXED2450_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2450E350g,
	"SUBCORE_VM_FIXED2475_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2475E350g,
	"SUBCORE_VM_FIXED2500_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2500E350g,
	"SUBCORE_VM_FIXED2550_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2550E350g,
	"SUBCORE_VM_FIXED2600_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2600E350g,
	"SUBCORE_VM_FIXED2625_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2625E350g,
	"SUBCORE_VM_FIXED2650_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2650E350g,
	"SUBCORE_VM_FIXED2700_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2700E350g,
	"SUBCORE_VM_FIXED2750_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2750E350g,
	"SUBCORE_VM_FIXED2775_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2775E350g,
	"SUBCORE_VM_FIXED2800_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2800E350g,
	"SUBCORE_VM_FIXED2850_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2850E350g,
	"SUBCORE_VM_FIXED2875_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2875E350g,
	"SUBCORE_VM_FIXED2900_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2900E350g,
	"SUBCORE_VM_FIXED2925_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2925E350g,
	"SUBCORE_VM_FIXED2950_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2950E350g,
	"SUBCORE_VM_FIXED2975_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2975E350g,
	"SUBCORE_VM_FIXED3000_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3000E350g,
	"SUBCORE_VM_FIXED3025_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3025E350g,
	"SUBCORE_VM_FIXED3050_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3050E350g,
	"SUBCORE_VM_FIXED3075_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3075E350g,
	"SUBCORE_VM_FIXED3100_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3100E350g,
	"SUBCORE_VM_FIXED3125_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3125E350g,
	"SUBCORE_VM_FIXED3150_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3150E350g,
	"SUBCORE_VM_FIXED3200_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3200E350g,
	"SUBCORE_VM_FIXED3225_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3225E350g,
	"SUBCORE_VM_FIXED3250_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3250E350g,
	"SUBCORE_VM_FIXED3300_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3300E350g,
	"SUBCORE_VM_FIXED3325_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3325E350g,
	"SUBCORE_VM_FIXED3375_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3375E350g,
	"SUBCORE_VM_FIXED3400_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3400E350g,
	"SUBCORE_VM_FIXED3450_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3450E350g,
	"SUBCORE_VM_FIXED3500_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3500E350g,
	"SUBCORE_VM_FIXED3525_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3525E350g,
	"SUBCORE_VM_FIXED3575_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3575E350g,
	"SUBCORE_VM_FIXED3600_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3600E350g,
	"SUBCORE_VM_FIXED3625_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3625E350g,
	"SUBCORE_VM_FIXED3675_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3675E350g,
	"SUBCORE_VM_FIXED3700_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3700E350g,
	"SUBCORE_VM_FIXED3750_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3750E350g,
	"SUBCORE_VM_FIXED3800_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3800E350g,
	"SUBCORE_VM_FIXED3825_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3825E350g,
	"SUBCORE_VM_FIXED3850_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3850E350g,
	"SUBCORE_VM_FIXED3875_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3875E350g,
	"SUBCORE_VM_FIXED3900_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3900E350g,
	"SUBCORE_VM_FIXED3975_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3975E350g,
	"SUBCORE_VM_FIXED4000_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4000E350g,
	"SUBCORE_VM_FIXED4025_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4025E350g,
	"SUBCORE_VM_FIXED4050_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4050E350g,
	"SUBCORE_VM_FIXED4100_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4100E350g,
	"SUBCORE_VM_FIXED4125_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4125E350g,
	"SUBCORE_VM_FIXED4200_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4200E350g,
	"SUBCORE_VM_FIXED4225_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4225E350g,
	"SUBCORE_VM_FIXED4250_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4250E350g,
	"SUBCORE_VM_FIXED4275_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4275E350g,
	"SUBCORE_VM_FIXED4300_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4300E350g,
	"SUBCORE_VM_FIXED4350_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4350E350g,
	"SUBCORE_VM_FIXED4375_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4375E350g,
	"SUBCORE_VM_FIXED4400_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4400E350g,
	"SUBCORE_VM_FIXED4425_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4425E350g,
	"SUBCORE_VM_FIXED4500_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4500E350g,
	"SUBCORE_VM_FIXED4550_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4550E350g,
	"SUBCORE_VM_FIXED4575_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4575E350g,
	"SUBCORE_VM_FIXED4600_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4600E350g,
	"SUBCORE_VM_FIXED4625_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4625E350g,
	"SUBCORE_VM_FIXED4650_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4650E350g,
	"SUBCORE_VM_FIXED4675_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4675E350g,
	"SUBCORE_VM_FIXED4700_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4700E350g,
	"SUBCORE_VM_FIXED4725_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4725E350g,
	"SUBCORE_VM_FIXED4750_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4750E350g,
	"SUBCORE_VM_FIXED4800_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4800E350g,
	"SUBCORE_VM_FIXED4875_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4875E350g,
	"SUBCORE_VM_FIXED4900_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4900E350g,
	"SUBCORE_VM_FIXED4950_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4950E350g,
	"SUBCORE_VM_FIXED5000_E3_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed5000E350g,
	"SUBCORE_VM_FIXED0025_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0025E450g,
	"SUBCORE_VM_FIXED0050_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0050E450g,
	"SUBCORE_VM_FIXED0075_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0075E450g,
	"SUBCORE_VM_FIXED0100_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0100E450g,
	"SUBCORE_VM_FIXED0125_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0125E450g,
	"SUBCORE_VM_FIXED0150_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0150E450g,
	"SUBCORE_VM_FIXED0175_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0175E450g,
	"SUBCORE_VM_FIXED0200_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0200E450g,
	"SUBCORE_VM_FIXED0225_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0225E450g,
	"SUBCORE_VM_FIXED0250_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0250E450g,
	"SUBCORE_VM_FIXED0275_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0275E450g,
	"SUBCORE_VM_FIXED0300_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0300E450g,
	"SUBCORE_VM_FIXED0325_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0325E450g,
	"SUBCORE_VM_FIXED0350_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0350E450g,
	"SUBCORE_VM_FIXED0375_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0375E450g,
	"SUBCORE_VM_FIXED0400_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0400E450g,
	"SUBCORE_VM_FIXED0425_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0425E450g,
	"SUBCORE_VM_FIXED0450_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0450E450g,
	"SUBCORE_VM_FIXED0475_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0475E450g,
	"SUBCORE_VM_FIXED0500_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0500E450g,
	"SUBCORE_VM_FIXED0525_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0525E450g,
	"SUBCORE_VM_FIXED0550_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0550E450g,
	"SUBCORE_VM_FIXED0575_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0575E450g,
	"SUBCORE_VM_FIXED0600_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0600E450g,
	"SUBCORE_VM_FIXED0625_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0625E450g,
	"SUBCORE_VM_FIXED0650_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0650E450g,
	"SUBCORE_VM_FIXED0675_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0675E450g,
	"SUBCORE_VM_FIXED0700_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0700E450g,
	"SUBCORE_VM_FIXED0725_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0725E450g,
	"SUBCORE_VM_FIXED0750_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0750E450g,
	"SUBCORE_VM_FIXED0775_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0775E450g,
	"SUBCORE_VM_FIXED0800_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0800E450g,
	"SUBCORE_VM_FIXED0825_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0825E450g,
	"SUBCORE_VM_FIXED0850_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0850E450g,
	"SUBCORE_VM_FIXED0875_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0875E450g,
	"SUBCORE_VM_FIXED0900_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0900E450g,
	"SUBCORE_VM_FIXED0925_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0925E450g,
	"SUBCORE_VM_FIXED0950_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0950E450g,
	"SUBCORE_VM_FIXED0975_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0975E450g,
	"SUBCORE_VM_FIXED1000_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1000E450g,
	"SUBCORE_VM_FIXED1025_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1025E450g,
	"SUBCORE_VM_FIXED1050_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1050E450g,
	"SUBCORE_VM_FIXED1075_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1075E450g,
	"SUBCORE_VM_FIXED1100_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1100E450g,
	"SUBCORE_VM_FIXED1125_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1125E450g,
	"SUBCORE_VM_FIXED1150_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1150E450g,
	"SUBCORE_VM_FIXED1175_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1175E450g,
	"SUBCORE_VM_FIXED1200_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1200E450g,
	"SUBCORE_VM_FIXED1225_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1225E450g,
	"SUBCORE_VM_FIXED1250_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1250E450g,
	"SUBCORE_VM_FIXED1275_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1275E450g,
	"SUBCORE_VM_FIXED1300_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1300E450g,
	"SUBCORE_VM_FIXED1325_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1325E450g,
	"SUBCORE_VM_FIXED1350_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1350E450g,
	"SUBCORE_VM_FIXED1375_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1375E450g,
	"SUBCORE_VM_FIXED1400_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1400E450g,
	"SUBCORE_VM_FIXED1425_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1425E450g,
	"SUBCORE_VM_FIXED1450_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1450E450g,
	"SUBCORE_VM_FIXED1475_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1475E450g,
	"SUBCORE_VM_FIXED1500_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1500E450g,
	"SUBCORE_VM_FIXED1525_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1525E450g,
	"SUBCORE_VM_FIXED1550_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1550E450g,
	"SUBCORE_VM_FIXED1575_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1575E450g,
	"SUBCORE_VM_FIXED1600_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1600E450g,
	"SUBCORE_VM_FIXED1625_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1625E450g,
	"SUBCORE_VM_FIXED1650_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1650E450g,
	"SUBCORE_VM_FIXED1700_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1700E450g,
	"SUBCORE_VM_FIXED1725_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1725E450g,
	"SUBCORE_VM_FIXED1750_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1750E450g,
	"SUBCORE_VM_FIXED1800_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1800E450g,
	"SUBCORE_VM_FIXED1850_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1850E450g,
	"SUBCORE_VM_FIXED1875_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1875E450g,
	"SUBCORE_VM_FIXED1900_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1900E450g,
	"SUBCORE_VM_FIXED1925_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1925E450g,
	"SUBCORE_VM_FIXED1950_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1950E450g,
	"SUBCORE_VM_FIXED2000_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2000E450g,
	"SUBCORE_VM_FIXED2025_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2025E450g,
	"SUBCORE_VM_FIXED2050_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2050E450g,
	"SUBCORE_VM_FIXED2100_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2100E450g,
	"SUBCORE_VM_FIXED2125_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2125E450g,
	"SUBCORE_VM_FIXED2150_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2150E450g,
	"SUBCORE_VM_FIXED2175_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2175E450g,
	"SUBCORE_VM_FIXED2200_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2200E450g,
	"SUBCORE_VM_FIXED2250_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2250E450g,
	"SUBCORE_VM_FIXED2275_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2275E450g,
	"SUBCORE_VM_FIXED2300_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2300E450g,
	"SUBCORE_VM_FIXED2325_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2325E450g,
	"SUBCORE_VM_FIXED2350_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2350E450g,
	"SUBCORE_VM_FIXED2375_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2375E450g,
	"SUBCORE_VM_FIXED2400_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2400E450g,
	"SUBCORE_VM_FIXED2450_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2450E450g,
	"SUBCORE_VM_FIXED2475_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2475E450g,
	"SUBCORE_VM_FIXED2500_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2500E450g,
	"SUBCORE_VM_FIXED2550_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2550E450g,
	"SUBCORE_VM_FIXED2600_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2600E450g,
	"SUBCORE_VM_FIXED2625_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2625E450g,
	"SUBCORE_VM_FIXED2650_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2650E450g,
	"SUBCORE_VM_FIXED2700_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2700E450g,
	"SUBCORE_VM_FIXED2750_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2750E450g,
	"SUBCORE_VM_FIXED2775_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2775E450g,
	"SUBCORE_VM_FIXED2800_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2800E450g,
	"SUBCORE_VM_FIXED2850_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2850E450g,
	"SUBCORE_VM_FIXED2875_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2875E450g,
	"SUBCORE_VM_FIXED2900_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2900E450g,
	"SUBCORE_VM_FIXED2925_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2925E450g,
	"SUBCORE_VM_FIXED2950_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2950E450g,
	"SUBCORE_VM_FIXED2975_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2975E450g,
	"SUBCORE_VM_FIXED3000_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3000E450g,
	"SUBCORE_VM_FIXED3025_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3025E450g,
	"SUBCORE_VM_FIXED3050_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3050E450g,
	"SUBCORE_VM_FIXED3075_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3075E450g,
	"SUBCORE_VM_FIXED3100_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3100E450g,
	"SUBCORE_VM_FIXED3125_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3125E450g,
	"SUBCORE_VM_FIXED3150_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3150E450g,
	"SUBCORE_VM_FIXED3200_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3200E450g,
	"SUBCORE_VM_FIXED3225_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3225E450g,
	"SUBCORE_VM_FIXED3250_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3250E450g,
	"SUBCORE_VM_FIXED3300_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3300E450g,
	"SUBCORE_VM_FIXED3325_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3325E450g,
	"SUBCORE_VM_FIXED3375_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3375E450g,
	"SUBCORE_VM_FIXED3400_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3400E450g,
	"SUBCORE_VM_FIXED3450_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3450E450g,
	"SUBCORE_VM_FIXED3500_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3500E450g,
	"SUBCORE_VM_FIXED3525_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3525E450g,
	"SUBCORE_VM_FIXED3575_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3575E450g,
	"SUBCORE_VM_FIXED3600_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3600E450g,
	"SUBCORE_VM_FIXED3625_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3625E450g,
	"SUBCORE_VM_FIXED3675_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3675E450g,
	"SUBCORE_VM_FIXED3700_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3700E450g,
	"SUBCORE_VM_FIXED3750_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3750E450g,
	"SUBCORE_VM_FIXED3800_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3800E450g,
	"SUBCORE_VM_FIXED3825_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3825E450g,
	"SUBCORE_VM_FIXED3850_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3850E450g,
	"SUBCORE_VM_FIXED3875_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3875E450g,
	"SUBCORE_VM_FIXED3900_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3900E450g,
	"SUBCORE_VM_FIXED3975_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3975E450g,
	"SUBCORE_VM_FIXED4000_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4000E450g,
	"SUBCORE_VM_FIXED4025_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4025E450g,
	"SUBCORE_VM_FIXED4050_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4050E450g,
	"SUBCORE_VM_FIXED4100_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4100E450g,
	"SUBCORE_VM_FIXED4125_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4125E450g,
	"SUBCORE_VM_FIXED4200_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4200E450g,
	"SUBCORE_VM_FIXED4225_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4225E450g,
	"SUBCORE_VM_FIXED4250_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4250E450g,
	"SUBCORE_VM_FIXED4275_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4275E450g,
	"SUBCORE_VM_FIXED4300_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4300E450g,
	"SUBCORE_VM_FIXED4350_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4350E450g,
	"SUBCORE_VM_FIXED4375_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4375E450g,
	"SUBCORE_VM_FIXED4400_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4400E450g,
	"SUBCORE_VM_FIXED4425_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4425E450g,
	"SUBCORE_VM_FIXED4500_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4500E450g,
	"SUBCORE_VM_FIXED4550_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4550E450g,
	"SUBCORE_VM_FIXED4575_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4575E450g,
	"SUBCORE_VM_FIXED4600_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4600E450g,
	"SUBCORE_VM_FIXED4625_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4625E450g,
	"SUBCORE_VM_FIXED4650_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4650E450g,
	"SUBCORE_VM_FIXED4675_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4675E450g,
	"SUBCORE_VM_FIXED4700_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4700E450g,
	"SUBCORE_VM_FIXED4725_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4725E450g,
	"SUBCORE_VM_FIXED4750_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4750E450g,
	"SUBCORE_VM_FIXED4800_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4800E450g,
	"SUBCORE_VM_FIXED4875_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4875E450g,
	"SUBCORE_VM_FIXED4900_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4900E450g,
	"SUBCORE_VM_FIXED4950_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4950E450g,
	"SUBCORE_VM_FIXED5000_E4_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed5000E450g,
	"SUBCORE_VM_FIXED0020_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0020A150g,
	"SUBCORE_VM_FIXED0040_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0040A150g,
	"SUBCORE_VM_FIXED0060_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0060A150g,
	"SUBCORE_VM_FIXED0080_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0080A150g,
	"SUBCORE_VM_FIXED0100_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0100A150g,
	"SUBCORE_VM_FIXED0120_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0120A150g,
	"SUBCORE_VM_FIXED0140_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0140A150g,
	"SUBCORE_VM_FIXED0160_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0160A150g,
	"SUBCORE_VM_FIXED0180_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0180A150g,
	"SUBCORE_VM_FIXED0200_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0200A150g,
	"SUBCORE_VM_FIXED0220_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0220A150g,
	"SUBCORE_VM_FIXED0240_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0240A150g,
	"SUBCORE_VM_FIXED0260_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0260A150g,
	"SUBCORE_VM_FIXED0280_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0280A150g,
	"SUBCORE_VM_FIXED0300_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0300A150g,
	"SUBCORE_VM_FIXED0320_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0320A150g,
	"SUBCORE_VM_FIXED0340_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0340A150g,
	"SUBCORE_VM_FIXED0360_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0360A150g,
	"SUBCORE_VM_FIXED0380_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0380A150g,
	"SUBCORE_VM_FIXED0400_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0400A150g,
	"SUBCORE_VM_FIXED0420_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0420A150g,
	"SUBCORE_VM_FIXED0440_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0440A150g,
	"SUBCORE_VM_FIXED0460_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0460A150g,
	"SUBCORE_VM_FIXED0480_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0480A150g,
	"SUBCORE_VM_FIXED0500_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0500A150g,
	"SUBCORE_VM_FIXED0520_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0520A150g,
	"SUBCORE_VM_FIXED0540_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0540A150g,
	"SUBCORE_VM_FIXED0560_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0560A150g,
	"SUBCORE_VM_FIXED0580_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0580A150g,
	"SUBCORE_VM_FIXED0600_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0600A150g,
	"SUBCORE_VM_FIXED0620_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0620A150g,
	"SUBCORE_VM_FIXED0640_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0640A150g,
	"SUBCORE_VM_FIXED0660_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0660A150g,
	"SUBCORE_VM_FIXED0680_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0680A150g,
	"SUBCORE_VM_FIXED0700_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0700A150g,
	"SUBCORE_VM_FIXED0720_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0720A150g,
	"SUBCORE_VM_FIXED0740_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0740A150g,
	"SUBCORE_VM_FIXED0760_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0760A150g,
	"SUBCORE_VM_FIXED0780_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0780A150g,
	"SUBCORE_VM_FIXED0800_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0800A150g,
	"SUBCORE_VM_FIXED0820_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0820A150g,
	"SUBCORE_VM_FIXED0840_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0840A150g,
	"SUBCORE_VM_FIXED0860_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0860A150g,
	"SUBCORE_VM_FIXED0880_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0880A150g,
	"SUBCORE_VM_FIXED0900_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0900A150g,
	"SUBCORE_VM_FIXED0920_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0920A150g,
	"SUBCORE_VM_FIXED0940_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0940A150g,
	"SUBCORE_VM_FIXED0960_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0960A150g,
	"SUBCORE_VM_FIXED0980_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0980A150g,
	"SUBCORE_VM_FIXED1000_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1000A150g,
	"SUBCORE_VM_FIXED1020_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1020A150g,
	"SUBCORE_VM_FIXED1040_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1040A150g,
	"SUBCORE_VM_FIXED1060_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1060A150g,
	"SUBCORE_VM_FIXED1080_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1080A150g,
	"SUBCORE_VM_FIXED1100_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1100A150g,
	"SUBCORE_VM_FIXED1120_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1120A150g,
	"SUBCORE_VM_FIXED1140_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1140A150g,
	"SUBCORE_VM_FIXED1160_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1160A150g,
	"SUBCORE_VM_FIXED1180_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1180A150g,
	"SUBCORE_VM_FIXED1200_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1200A150g,
	"SUBCORE_VM_FIXED1220_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1220A150g,
	"SUBCORE_VM_FIXED1240_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1240A150g,
	"SUBCORE_VM_FIXED1260_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1260A150g,
	"SUBCORE_VM_FIXED1280_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1280A150g,
	"SUBCORE_VM_FIXED1300_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1300A150g,
	"SUBCORE_VM_FIXED1320_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1320A150g,
	"SUBCORE_VM_FIXED1340_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1340A150g,
	"SUBCORE_VM_FIXED1360_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1360A150g,
	"SUBCORE_VM_FIXED1380_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1380A150g,
	"SUBCORE_VM_FIXED1400_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1400A150g,
	"SUBCORE_VM_FIXED1420_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1420A150g,
	"SUBCORE_VM_FIXED1440_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1440A150g,
	"SUBCORE_VM_FIXED1460_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1460A150g,
	"SUBCORE_VM_FIXED1480_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1480A150g,
	"SUBCORE_VM_FIXED1500_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1500A150g,
	"SUBCORE_VM_FIXED1520_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1520A150g,
	"SUBCORE_VM_FIXED1540_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1540A150g,
	"SUBCORE_VM_FIXED1560_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1560A150g,
	"SUBCORE_VM_FIXED1580_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1580A150g,
	"SUBCORE_VM_FIXED1600_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1600A150g,
	"SUBCORE_VM_FIXED1620_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1620A150g,
	"SUBCORE_VM_FIXED1640_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1640A150g,
	"SUBCORE_VM_FIXED1660_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1660A150g,
	"SUBCORE_VM_FIXED1680_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1680A150g,
	"SUBCORE_VM_FIXED1700_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1700A150g,
	"SUBCORE_VM_FIXED1720_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1720A150g,
	"SUBCORE_VM_FIXED1740_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1740A150g,
	"SUBCORE_VM_FIXED1760_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1760A150g,
	"SUBCORE_VM_FIXED1780_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1780A150g,
	"SUBCORE_VM_FIXED1800_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1800A150g,
	"SUBCORE_VM_FIXED1820_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1820A150g,
	"SUBCORE_VM_FIXED1840_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1840A150g,
	"SUBCORE_VM_FIXED1860_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1860A150g,
	"SUBCORE_VM_FIXED1880_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1880A150g,
	"SUBCORE_VM_FIXED1900_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1900A150g,
	"SUBCORE_VM_FIXED1920_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1920A150g,
	"SUBCORE_VM_FIXED1940_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1940A150g,
	"SUBCORE_VM_FIXED1960_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1960A150g,
	"SUBCORE_VM_FIXED1980_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1980A150g,
	"SUBCORE_VM_FIXED2000_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2000A150g,
	"SUBCORE_VM_FIXED2020_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2020A150g,
	"SUBCORE_VM_FIXED2040_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2040A150g,
	"SUBCORE_VM_FIXED2060_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2060A150g,
	"SUBCORE_VM_FIXED2080_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2080A150g,
	"SUBCORE_VM_FIXED2100_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2100A150g,
	"SUBCORE_VM_FIXED2120_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2120A150g,
	"SUBCORE_VM_FIXED2140_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2140A150g,
	"SUBCORE_VM_FIXED2160_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2160A150g,
	"SUBCORE_VM_FIXED2180_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2180A150g,
	"SUBCORE_VM_FIXED2200_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2200A150g,
	"SUBCORE_VM_FIXED2220_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2220A150g,
	"SUBCORE_VM_FIXED2240_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2240A150g,
	"SUBCORE_VM_FIXED2260_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2260A150g,
	"SUBCORE_VM_FIXED2280_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2280A150g,
	"SUBCORE_VM_FIXED2300_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2300A150g,
	"SUBCORE_VM_FIXED2320_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2320A150g,
	"SUBCORE_VM_FIXED2340_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2340A150g,
	"SUBCORE_VM_FIXED2360_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2360A150g,
	"SUBCORE_VM_FIXED2380_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2380A150g,
	"SUBCORE_VM_FIXED2400_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2400A150g,
	"SUBCORE_VM_FIXED2420_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2420A150g,
	"SUBCORE_VM_FIXED2440_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2440A150g,
	"SUBCORE_VM_FIXED2460_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2460A150g,
	"SUBCORE_VM_FIXED2480_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2480A150g,
	"SUBCORE_VM_FIXED2500_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2500A150g,
	"SUBCORE_VM_FIXED2520_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2520A150g,
	"SUBCORE_VM_FIXED2540_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2540A150g,
	"SUBCORE_VM_FIXED2560_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2560A150g,
	"SUBCORE_VM_FIXED2580_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2580A150g,
	"SUBCORE_VM_FIXED2600_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2600A150g,
	"SUBCORE_VM_FIXED2620_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2620A150g,
	"SUBCORE_VM_FIXED2640_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2640A150g,
	"SUBCORE_VM_FIXED2660_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2660A150g,
	"SUBCORE_VM_FIXED2680_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2680A150g,
	"SUBCORE_VM_FIXED2700_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2700A150g,
	"SUBCORE_VM_FIXED2720_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2720A150g,
	"SUBCORE_VM_FIXED2740_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2740A150g,
	"SUBCORE_VM_FIXED2760_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2760A150g,
	"SUBCORE_VM_FIXED2780_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2780A150g,
	"SUBCORE_VM_FIXED2800_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2800A150g,
	"SUBCORE_VM_FIXED2820_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2820A150g,
	"SUBCORE_VM_FIXED2840_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2840A150g,
	"SUBCORE_VM_FIXED2860_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2860A150g,
	"SUBCORE_VM_FIXED2880_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2880A150g,
	"SUBCORE_VM_FIXED2900_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2900A150g,
	"SUBCORE_VM_FIXED2920_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2920A150g,
	"SUBCORE_VM_FIXED2940_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2940A150g,
	"SUBCORE_VM_FIXED2960_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2960A150g,
	"SUBCORE_VM_FIXED2980_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2980A150g,
	"SUBCORE_VM_FIXED3000_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3000A150g,
	"SUBCORE_VM_FIXED3020_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3020A150g,
	"SUBCORE_VM_FIXED3040_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3040A150g,
	"SUBCORE_VM_FIXED3060_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3060A150g,
	"SUBCORE_VM_FIXED3080_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3080A150g,
	"SUBCORE_VM_FIXED3100_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3100A150g,
	"SUBCORE_VM_FIXED3120_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3120A150g,
	"SUBCORE_VM_FIXED3140_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3140A150g,
	"SUBCORE_VM_FIXED3160_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3160A150g,
	"SUBCORE_VM_FIXED3180_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3180A150g,
	"SUBCORE_VM_FIXED3200_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3200A150g,
	"SUBCORE_VM_FIXED3220_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3220A150g,
	"SUBCORE_VM_FIXED3240_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3240A150g,
	"SUBCORE_VM_FIXED3260_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3260A150g,
	"SUBCORE_VM_FIXED3280_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3280A150g,
	"SUBCORE_VM_FIXED3300_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3300A150g,
	"SUBCORE_VM_FIXED3320_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3320A150g,
	"SUBCORE_VM_FIXED3340_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3340A150g,
	"SUBCORE_VM_FIXED3360_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3360A150g,
	"SUBCORE_VM_FIXED3380_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3380A150g,
	"SUBCORE_VM_FIXED3400_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3400A150g,
	"SUBCORE_VM_FIXED3420_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3420A150g,
	"SUBCORE_VM_FIXED3440_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3440A150g,
	"SUBCORE_VM_FIXED3460_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3460A150g,
	"SUBCORE_VM_FIXED3480_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3480A150g,
	"SUBCORE_VM_FIXED3500_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3500A150g,
	"SUBCORE_VM_FIXED3520_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3520A150g,
	"SUBCORE_VM_FIXED3540_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3540A150g,
	"SUBCORE_VM_FIXED3560_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3560A150g,
	"SUBCORE_VM_FIXED3580_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3580A150g,
	"SUBCORE_VM_FIXED3600_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3600A150g,
	"SUBCORE_VM_FIXED3620_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3620A150g,
	"SUBCORE_VM_FIXED3640_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3640A150g,
	"SUBCORE_VM_FIXED3660_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3660A150g,
	"SUBCORE_VM_FIXED3680_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3680A150g,
	"SUBCORE_VM_FIXED3700_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3700A150g,
	"SUBCORE_VM_FIXED3720_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3720A150g,
	"SUBCORE_VM_FIXED3740_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3740A150g,
	"SUBCORE_VM_FIXED3760_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3760A150g,
	"SUBCORE_VM_FIXED3780_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3780A150g,
	"SUBCORE_VM_FIXED3800_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3800A150g,
	"SUBCORE_VM_FIXED3820_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3820A150g,
	"SUBCORE_VM_FIXED3840_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3840A150g,
	"SUBCORE_VM_FIXED3860_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3860A150g,
	"SUBCORE_VM_FIXED3880_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3880A150g,
	"SUBCORE_VM_FIXED3900_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3900A150g,
	"SUBCORE_VM_FIXED3920_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3920A150g,
	"SUBCORE_VM_FIXED3940_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3940A150g,
	"SUBCORE_VM_FIXED3960_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3960A150g,
	"SUBCORE_VM_FIXED3980_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3980A150g,
	"SUBCORE_VM_FIXED4000_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4000A150g,
	"SUBCORE_VM_FIXED4020_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4020A150g,
	"SUBCORE_VM_FIXED4040_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4040A150g,
	"SUBCORE_VM_FIXED4060_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4060A150g,
	"SUBCORE_VM_FIXED4080_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4080A150g,
	"SUBCORE_VM_FIXED4100_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4100A150g,
	"SUBCORE_VM_FIXED4120_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4120A150g,
	"SUBCORE_VM_FIXED4140_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4140A150g,
	"SUBCORE_VM_FIXED4160_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4160A150g,
	"SUBCORE_VM_FIXED4180_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4180A150g,
	"SUBCORE_VM_FIXED4200_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4200A150g,
	"SUBCORE_VM_FIXED4220_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4220A150g,
	"SUBCORE_VM_FIXED4240_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4240A150g,
	"SUBCORE_VM_FIXED4260_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4260A150g,
	"SUBCORE_VM_FIXED4280_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4280A150g,
	"SUBCORE_VM_FIXED4300_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4300A150g,
	"SUBCORE_VM_FIXED4320_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4320A150g,
	"SUBCORE_VM_FIXED4340_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4340A150g,
	"SUBCORE_VM_FIXED4360_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4360A150g,
	"SUBCORE_VM_FIXED4380_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4380A150g,
	"SUBCORE_VM_FIXED4400_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4400A150g,
	"SUBCORE_VM_FIXED4420_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4420A150g,
	"SUBCORE_VM_FIXED4440_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4440A150g,
	"SUBCORE_VM_FIXED4460_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4460A150g,
	"SUBCORE_VM_FIXED4480_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4480A150g,
	"SUBCORE_VM_FIXED4500_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4500A150g,
	"SUBCORE_VM_FIXED4520_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4520A150g,
	"SUBCORE_VM_FIXED4540_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4540A150g,
	"SUBCORE_VM_FIXED4560_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4560A150g,
	"SUBCORE_VM_FIXED4580_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4580A150g,
	"SUBCORE_VM_FIXED4600_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4600A150g,
	"SUBCORE_VM_FIXED4620_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4620A150g,
	"SUBCORE_VM_FIXED4640_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4640A150g,
	"SUBCORE_VM_FIXED4660_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4660A150g,
	"SUBCORE_VM_FIXED4680_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4680A150g,
	"SUBCORE_VM_FIXED4700_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4700A150g,
	"SUBCORE_VM_FIXED4720_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4720A150g,
	"SUBCORE_VM_FIXED4740_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4740A150g,
	"SUBCORE_VM_FIXED4760_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4760A150g,
	"SUBCORE_VM_FIXED4780_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4780A150g,
	"SUBCORE_VM_FIXED4800_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4800A150g,
	"SUBCORE_VM_FIXED4820_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4820A150g,
	"SUBCORE_VM_FIXED4840_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4840A150g,
	"SUBCORE_VM_FIXED4860_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4860A150g,
	"SUBCORE_VM_FIXED4880_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4880A150g,
	"SUBCORE_VM_FIXED4900_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4900A150g,
	"SUBCORE_VM_FIXED4920_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4920A150g,
	"SUBCORE_VM_FIXED4940_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4940A150g,
	"SUBCORE_VM_FIXED4960_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4960A150g,
	"SUBCORE_VM_FIXED4980_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4980A150g,
	"SUBCORE_VM_FIXED5000_A1_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed5000A150g,
	"SUBCORE_VM_FIXED0090_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0090X950g,
	"SUBCORE_VM_FIXED0180_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0180X950g,
	"SUBCORE_VM_FIXED0270_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0270X950g,
	"SUBCORE_VM_FIXED0360_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0360X950g,
	"SUBCORE_VM_FIXED0450_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0450X950g,
	"SUBCORE_VM_FIXED0540_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0540X950g,
	"SUBCORE_VM_FIXED0630_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0630X950g,
	"SUBCORE_VM_FIXED0720_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0720X950g,
	"SUBCORE_VM_FIXED0810_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0810X950g,
	"SUBCORE_VM_FIXED0900_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0900X950g,
	"SUBCORE_VM_FIXED0990_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0990X950g,
	"SUBCORE_VM_FIXED1080_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1080X950g,
	"SUBCORE_VM_FIXED1170_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1170X950g,
	"SUBCORE_VM_FIXED1260_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1260X950g,
	"SUBCORE_VM_FIXED1350_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1350X950g,
	"SUBCORE_VM_FIXED1440_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1440X950g,
	"SUBCORE_VM_FIXED1530_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1530X950g,
	"SUBCORE_VM_FIXED1620_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1620X950g,
	"SUBCORE_VM_FIXED1710_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1710X950g,
	"SUBCORE_VM_FIXED1800_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1800X950g,
	"SUBCORE_VM_FIXED1890_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1890X950g,
	"SUBCORE_VM_FIXED1980_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1980X950g,
	"SUBCORE_VM_FIXED2070_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2070X950g,
	"SUBCORE_VM_FIXED2160_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2160X950g,
	"SUBCORE_VM_FIXED2250_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2250X950g,
	"SUBCORE_VM_FIXED2340_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2340X950g,
	"SUBCORE_VM_FIXED2430_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2430X950g,
	"SUBCORE_VM_FIXED2520_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2520X950g,
	"SUBCORE_VM_FIXED2610_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2610X950g,
	"SUBCORE_VM_FIXED2700_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2700X950g,
	"SUBCORE_VM_FIXED2790_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2790X950g,
	"SUBCORE_VM_FIXED2880_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2880X950g,
	"SUBCORE_VM_FIXED2970_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2970X950g,
	"SUBCORE_VM_FIXED3060_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3060X950g,
	"SUBCORE_VM_FIXED3150_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3150X950g,
	"SUBCORE_VM_FIXED3240_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3240X950g,
	"SUBCORE_VM_FIXED3330_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3330X950g,
	"SUBCORE_VM_FIXED3420_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3420X950g,
	"SUBCORE_VM_FIXED3510_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3510X950g,
	"SUBCORE_VM_FIXED3600_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3600X950g,
	"SUBCORE_VM_FIXED3690_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3690X950g,
	"SUBCORE_VM_FIXED3780_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3780X950g,
	"SUBCORE_VM_FIXED3870_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3870X950g,
	"SUBCORE_VM_FIXED3960_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3960X950g,
	"SUBCORE_VM_FIXED4050_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4050X950g,
	"SUBCORE_VM_FIXED4140_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4140X950g,
	"SUBCORE_VM_FIXED4230_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4230X950g,
	"SUBCORE_VM_FIXED4320_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4320X950g,
	"SUBCORE_VM_FIXED4410_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4410X950g,
	"SUBCORE_VM_FIXED4500_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4500X950g,
	"SUBCORE_VM_FIXED4590_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4590X950g,
	"SUBCORE_VM_FIXED4680_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4680X950g,
	"SUBCORE_VM_FIXED4770_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4770X950g,
	"SUBCORE_VM_FIXED4860_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4860X950g,
	"SUBCORE_VM_FIXED4950_X9_50G":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4950X950g,
	"DYNAMIC_A1_50G":                       InternalVnicAttachmentVnicShapeDynamicA150g,
	"FIXED0040_A1_50G":                     InternalVnicAttachmentVnicShapeFixed0040A150g,
	"FIXED0100_A1_50G":                     InternalVnicAttachmentVnicShapeFixed0100A150g,
	"FIXED0200_A1_50G":                     InternalVnicAttachmentVnicShapeFixed0200A150g,
	"FIXED0300_A1_50G":                     InternalVnicAttachmentVnicShapeFixed0300A150g,
	"FIXED0400_A1_50G":                     InternalVnicAttachmentVnicShapeFixed0400A150g,
	"FIXED0500_A1_50G":                     InternalVnicAttachmentVnicShapeFixed0500A150g,
	"FIXED0600_A1_50G":                     InternalVnicAttachmentVnicShapeFixed0600A150g,
	"FIXED0700_A1_50G":                     InternalVnicAttachmentVnicShapeFixed0700A150g,
	"FIXED0800_A1_50G":                     InternalVnicAttachmentVnicShapeFixed0800A150g,
	"FIXED0900_A1_50G":                     InternalVnicAttachmentVnicShapeFixed0900A150g,
	"FIXED1000_A1_50G":                     InternalVnicAttachmentVnicShapeFixed1000A150g,
	"FIXED1100_A1_50G":                     InternalVnicAttachmentVnicShapeFixed1100A150g,
	"FIXED1200_A1_50G":                     InternalVnicAttachmentVnicShapeFixed1200A150g,
	"FIXED1300_A1_50G":                     InternalVnicAttachmentVnicShapeFixed1300A150g,
	"FIXED1400_A1_50G":                     InternalVnicAttachmentVnicShapeFixed1400A150g,
	"FIXED1500_A1_50G":                     InternalVnicAttachmentVnicShapeFixed1500A150g,
	"FIXED1600_A1_50G":                     InternalVnicAttachmentVnicShapeFixed1600A150g,
	"FIXED1700_A1_50G":                     InternalVnicAttachmentVnicShapeFixed1700A150g,
	"FIXED1800_A1_50G":                     InternalVnicAttachmentVnicShapeFixed1800A150g,
	"FIXED1900_A1_50G":                     InternalVnicAttachmentVnicShapeFixed1900A150g,
	"FIXED2000_A1_50G":                     InternalVnicAttachmentVnicShapeFixed2000A150g,
	"FIXED2100_A1_50G":                     InternalVnicAttachmentVnicShapeFixed2100A150g,
	"FIXED2200_A1_50G":                     InternalVnicAttachmentVnicShapeFixed2200A150g,
	"FIXED2300_A1_50G":                     InternalVnicAttachmentVnicShapeFixed2300A150g,
	"FIXED2400_A1_50G":                     InternalVnicAttachmentVnicShapeFixed2400A150g,
	"FIXED2500_A1_50G":                     InternalVnicAttachmentVnicShapeFixed2500A150g,
	"FIXED2600_A1_50G":                     InternalVnicAttachmentVnicShapeFixed2600A150g,
	"FIXED2700_A1_50G":                     InternalVnicAttachmentVnicShapeFixed2700A150g,
	"FIXED2800_A1_50G":                     InternalVnicAttachmentVnicShapeFixed2800A150g,
	"FIXED2900_A1_50G":                     InternalVnicAttachmentVnicShapeFixed2900A150g,
	"FIXED3000_A1_50G":                     InternalVnicAttachmentVnicShapeFixed3000A150g,
	"FIXED3100_A1_50G":                     InternalVnicAttachmentVnicShapeFixed3100A150g,
	"FIXED3200_A1_50G":                     InternalVnicAttachmentVnicShapeFixed3200A150g,
	"FIXED3300_A1_50G":                     InternalVnicAttachmentVnicShapeFixed3300A150g,
	"FIXED3400_A1_50G":                     InternalVnicAttachmentVnicShapeFixed3400A150g,
	"FIXED3500_A1_50G":                     InternalVnicAttachmentVnicShapeFixed3500A150g,
	"FIXED3600_A1_50G":                     InternalVnicAttachmentVnicShapeFixed3600A150g,
	"FIXED3700_A1_50G":                     InternalVnicAttachmentVnicShapeFixed3700A150g,
	"FIXED3800_A1_50G":                     InternalVnicAttachmentVnicShapeFixed3800A150g,
	"FIXED3900_A1_50G":                     InternalVnicAttachmentVnicShapeFixed3900A150g,
	"FIXED4000_A1_50G":                     InternalVnicAttachmentVnicShapeFixed4000A150g,
	"ENTIREHOST_A1_50G":                    InternalVnicAttachmentVnicShapeEntirehostA150g,
	"DYNAMIC_X9_50G":                       InternalVnicAttachmentVnicShapeDynamicX950g,
	"FIXED0040_X9_50G":                     InternalVnicAttachmentVnicShapeFixed0040X950g,
	"FIXED0400_X9_50G":                     InternalVnicAttachmentVnicShapeFixed0400X950g,
	"FIXED0800_X9_50G":                     InternalVnicAttachmentVnicShapeFixed0800X950g,
	"FIXED1200_X9_50G":                     InternalVnicAttachmentVnicShapeFixed1200X950g,
	"FIXED1600_X9_50G":                     InternalVnicAttachmentVnicShapeFixed1600X950g,
	"FIXED2000_X9_50G":                     InternalVnicAttachmentVnicShapeFixed2000X950g,
	"FIXED2400_X9_50G":                     InternalVnicAttachmentVnicShapeFixed2400X950g,
	"FIXED2800_X9_50G":                     InternalVnicAttachmentVnicShapeFixed2800X950g,
	"FIXED3200_X9_50G":                     InternalVnicAttachmentVnicShapeFixed3200X950g,
	"FIXED3600_X9_50G":                     InternalVnicAttachmentVnicShapeFixed3600X950g,
	"FIXED4000_X9_50G":                     InternalVnicAttachmentVnicShapeFixed4000X950g,
	"STANDARD_VM_FIXED0100_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed0100X950g,
	"STANDARD_VM_FIXED0200_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed0200X950g,
	"STANDARD_VM_FIXED0300_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed0300X950g,
	"STANDARD_VM_FIXED0400_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed0400X950g,
	"STANDARD_VM_FIXED0500_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed0500X950g,
	"STANDARD_VM_FIXED0600_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed0600X950g,
	"STANDARD_VM_FIXED0700_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed0700X950g,
	"STANDARD_VM_FIXED0800_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed0800X950g,
	"STANDARD_VM_FIXED0900_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed0900X950g,
	"STANDARD_VM_FIXED1000_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed1000X950g,
	"STANDARD_VM_FIXED1100_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed1100X950g,
	"STANDARD_VM_FIXED1200_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed1200X950g,
	"STANDARD_VM_FIXED1300_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed1300X950g,
	"STANDARD_VM_FIXED1400_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed1400X950g,
	"STANDARD_VM_FIXED1500_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed1500X950g,
	"STANDARD_VM_FIXED1600_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed1600X950g,
	"STANDARD_VM_FIXED1700_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed1700X950g,
	"STANDARD_VM_FIXED1800_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed1800X950g,
	"STANDARD_VM_FIXED1900_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed1900X950g,
	"STANDARD_VM_FIXED2000_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed2000X950g,
	"STANDARD_VM_FIXED2100_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed2100X950g,
	"STANDARD_VM_FIXED2200_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed2200X950g,
	"STANDARD_VM_FIXED2300_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed2300X950g,
	"STANDARD_VM_FIXED2400_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed2400X950g,
	"STANDARD_VM_FIXED2500_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed2500X950g,
	"STANDARD_VM_FIXED2600_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed2600X950g,
	"STANDARD_VM_FIXED2700_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed2700X950g,
	"STANDARD_VM_FIXED2800_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed2800X950g,
	"STANDARD_VM_FIXED2900_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed2900X950g,
	"STANDARD_VM_FIXED3000_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed3000X950g,
	"STANDARD_VM_FIXED3100_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed3100X950g,
	"STANDARD_VM_FIXED3200_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed3200X950g,
	"STANDARD_VM_FIXED3300_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed3300X950g,
	"STANDARD_VM_FIXED3400_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed3400X950g,
	"STANDARD_VM_FIXED3500_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed3500X950g,
	"STANDARD_VM_FIXED3600_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed3600X950g,
	"STANDARD_VM_FIXED3700_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed3700X950g,
	"STANDARD_VM_FIXED3800_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed3800X950g,
	"STANDARD_VM_FIXED3900_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed3900X950g,
	"STANDARD_VM_FIXED4000_X9_50G":         InternalVnicAttachmentVnicShapeStandardVmFixed4000X950g,
	"SUBCORE_STANDARD_VM_FIXED0025_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0025X950g,
	"SUBCORE_STANDARD_VM_FIXED0050_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0050X950g,
	"SUBCORE_STANDARD_VM_FIXED0075_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0075X950g,
	"SUBCORE_STANDARD_VM_FIXED0100_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0100X950g,
	"SUBCORE_STANDARD_VM_FIXED0125_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0125X950g,
	"SUBCORE_STANDARD_VM_FIXED0150_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0150X950g,
	"SUBCORE_STANDARD_VM_FIXED0175_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0175X950g,
	"SUBCORE_STANDARD_VM_FIXED0200_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0200X950g,
	"SUBCORE_STANDARD_VM_FIXED0225_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0225X950g,
	"SUBCORE_STANDARD_VM_FIXED0250_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0250X950g,
	"SUBCORE_STANDARD_VM_FIXED0275_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0275X950g,
	"SUBCORE_STANDARD_VM_FIXED0300_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0300X950g,
	"SUBCORE_STANDARD_VM_FIXED0325_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0325X950g,
	"SUBCORE_STANDARD_VM_FIXED0350_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0350X950g,
	"SUBCORE_STANDARD_VM_FIXED0375_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0375X950g,
	"SUBCORE_STANDARD_VM_FIXED0400_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0400X950g,
	"SUBCORE_STANDARD_VM_FIXED0425_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0425X950g,
	"SUBCORE_STANDARD_VM_FIXED0450_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0450X950g,
	"SUBCORE_STANDARD_VM_FIXED0475_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0475X950g,
	"SUBCORE_STANDARD_VM_FIXED0500_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0500X950g,
	"SUBCORE_STANDARD_VM_FIXED0525_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0525X950g,
	"SUBCORE_STANDARD_VM_FIXED0550_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0550X950g,
	"SUBCORE_STANDARD_VM_FIXED0575_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0575X950g,
	"SUBCORE_STANDARD_VM_FIXED0600_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0600X950g,
	"SUBCORE_STANDARD_VM_FIXED0625_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0625X950g,
	"SUBCORE_STANDARD_VM_FIXED0650_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0650X950g,
	"SUBCORE_STANDARD_VM_FIXED0675_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0675X950g,
	"SUBCORE_STANDARD_VM_FIXED0700_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0700X950g,
	"SUBCORE_STANDARD_VM_FIXED0725_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0725X950g,
	"SUBCORE_STANDARD_VM_FIXED0750_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0750X950g,
	"SUBCORE_STANDARD_VM_FIXED0775_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0775X950g,
	"SUBCORE_STANDARD_VM_FIXED0800_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0800X950g,
	"SUBCORE_STANDARD_VM_FIXED0825_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0825X950g,
	"SUBCORE_STANDARD_VM_FIXED0850_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0850X950g,
	"SUBCORE_STANDARD_VM_FIXED0875_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0875X950g,
	"SUBCORE_STANDARD_VM_FIXED0900_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0900X950g,
	"SUBCORE_STANDARD_VM_FIXED0925_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0925X950g,
	"SUBCORE_STANDARD_VM_FIXED0950_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0950X950g,
	"SUBCORE_STANDARD_VM_FIXED0975_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0975X950g,
	"SUBCORE_STANDARD_VM_FIXED1000_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1000X950g,
	"SUBCORE_STANDARD_VM_FIXED1025_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1025X950g,
	"SUBCORE_STANDARD_VM_FIXED1050_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1050X950g,
	"SUBCORE_STANDARD_VM_FIXED1075_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1075X950g,
	"SUBCORE_STANDARD_VM_FIXED1100_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1100X950g,
	"SUBCORE_STANDARD_VM_FIXED1125_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1125X950g,
	"SUBCORE_STANDARD_VM_FIXED1150_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1150X950g,
	"SUBCORE_STANDARD_VM_FIXED1175_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1175X950g,
	"SUBCORE_STANDARD_VM_FIXED1200_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1200X950g,
	"SUBCORE_STANDARD_VM_FIXED1225_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1225X950g,
	"SUBCORE_STANDARD_VM_FIXED1250_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1250X950g,
	"SUBCORE_STANDARD_VM_FIXED1275_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1275X950g,
	"SUBCORE_STANDARD_VM_FIXED1300_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1300X950g,
	"SUBCORE_STANDARD_VM_FIXED1325_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1325X950g,
	"SUBCORE_STANDARD_VM_FIXED1350_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1350X950g,
	"SUBCORE_STANDARD_VM_FIXED1375_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1375X950g,
	"SUBCORE_STANDARD_VM_FIXED1400_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1400X950g,
	"SUBCORE_STANDARD_VM_FIXED1425_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1425X950g,
	"SUBCORE_STANDARD_VM_FIXED1450_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1450X950g,
	"SUBCORE_STANDARD_VM_FIXED1475_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1475X950g,
	"SUBCORE_STANDARD_VM_FIXED1500_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1500X950g,
	"SUBCORE_STANDARD_VM_FIXED1525_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1525X950g,
	"SUBCORE_STANDARD_VM_FIXED1550_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1550X950g,
	"SUBCORE_STANDARD_VM_FIXED1575_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1575X950g,
	"SUBCORE_STANDARD_VM_FIXED1600_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1600X950g,
	"SUBCORE_STANDARD_VM_FIXED1625_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1625X950g,
	"SUBCORE_STANDARD_VM_FIXED1650_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1650X950g,
	"SUBCORE_STANDARD_VM_FIXED1700_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1700X950g,
	"SUBCORE_STANDARD_VM_FIXED1725_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1725X950g,
	"SUBCORE_STANDARD_VM_FIXED1750_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1750X950g,
	"SUBCORE_STANDARD_VM_FIXED1800_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1800X950g,
	"SUBCORE_STANDARD_VM_FIXED1850_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1850X950g,
	"SUBCORE_STANDARD_VM_FIXED1875_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1875X950g,
	"SUBCORE_STANDARD_VM_FIXED1900_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1900X950g,
	"SUBCORE_STANDARD_VM_FIXED1925_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1925X950g,
	"SUBCORE_STANDARD_VM_FIXED1950_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1950X950g,
	"SUBCORE_STANDARD_VM_FIXED2000_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2000X950g,
	"SUBCORE_STANDARD_VM_FIXED2025_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2025X950g,
	"SUBCORE_STANDARD_VM_FIXED2050_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2050X950g,
	"SUBCORE_STANDARD_VM_FIXED2100_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2100X950g,
	"SUBCORE_STANDARD_VM_FIXED2125_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2125X950g,
	"SUBCORE_STANDARD_VM_FIXED2150_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2150X950g,
	"SUBCORE_STANDARD_VM_FIXED2175_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2175X950g,
	"SUBCORE_STANDARD_VM_FIXED2200_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2200X950g,
	"SUBCORE_STANDARD_VM_FIXED2250_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2250X950g,
	"SUBCORE_STANDARD_VM_FIXED2275_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2275X950g,
	"SUBCORE_STANDARD_VM_FIXED2300_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2300X950g,
	"SUBCORE_STANDARD_VM_FIXED2325_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2325X950g,
	"SUBCORE_STANDARD_VM_FIXED2350_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2350X950g,
	"SUBCORE_STANDARD_VM_FIXED2375_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2375X950g,
	"SUBCORE_STANDARD_VM_FIXED2400_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2400X950g,
	"SUBCORE_STANDARD_VM_FIXED2450_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2450X950g,
	"SUBCORE_STANDARD_VM_FIXED2475_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2475X950g,
	"SUBCORE_STANDARD_VM_FIXED2500_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2500X950g,
	"SUBCORE_STANDARD_VM_FIXED2550_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2550X950g,
	"SUBCORE_STANDARD_VM_FIXED2600_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2600X950g,
	"SUBCORE_STANDARD_VM_FIXED2625_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2625X950g,
	"SUBCORE_STANDARD_VM_FIXED2650_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2650X950g,
	"SUBCORE_STANDARD_VM_FIXED2700_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2700X950g,
	"SUBCORE_STANDARD_VM_FIXED2750_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2750X950g,
	"SUBCORE_STANDARD_VM_FIXED2775_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2775X950g,
	"SUBCORE_STANDARD_VM_FIXED2800_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2800X950g,
	"SUBCORE_STANDARD_VM_FIXED2850_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2850X950g,
	"SUBCORE_STANDARD_VM_FIXED2875_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2875X950g,
	"SUBCORE_STANDARD_VM_FIXED2900_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2900X950g,
	"SUBCORE_STANDARD_VM_FIXED2925_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2925X950g,
	"SUBCORE_STANDARD_VM_FIXED2950_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2950X950g,
	"SUBCORE_STANDARD_VM_FIXED2975_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2975X950g,
	"SUBCORE_STANDARD_VM_FIXED3000_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3000X950g,
	"SUBCORE_STANDARD_VM_FIXED3025_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3025X950g,
	"SUBCORE_STANDARD_VM_FIXED3050_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3050X950g,
	"SUBCORE_STANDARD_VM_FIXED3075_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3075X950g,
	"SUBCORE_STANDARD_VM_FIXED3100_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3100X950g,
	"SUBCORE_STANDARD_VM_FIXED3125_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3125X950g,
	"SUBCORE_STANDARD_VM_FIXED3150_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3150X950g,
	"SUBCORE_STANDARD_VM_FIXED3200_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3200X950g,
	"SUBCORE_STANDARD_VM_FIXED3225_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3225X950g,
	"SUBCORE_STANDARD_VM_FIXED3250_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3250X950g,
	"SUBCORE_STANDARD_VM_FIXED3300_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3300X950g,
	"SUBCORE_STANDARD_VM_FIXED3325_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3325X950g,
	"SUBCORE_STANDARD_VM_FIXED3375_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3375X950g,
	"SUBCORE_STANDARD_VM_FIXED3400_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3400X950g,
	"SUBCORE_STANDARD_VM_FIXED3450_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3450X950g,
	"SUBCORE_STANDARD_VM_FIXED3500_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3500X950g,
	"SUBCORE_STANDARD_VM_FIXED3525_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3525X950g,
	"SUBCORE_STANDARD_VM_FIXED3575_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3575X950g,
	"SUBCORE_STANDARD_VM_FIXED3600_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3600X950g,
	"SUBCORE_STANDARD_VM_FIXED3625_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3625X950g,
	"SUBCORE_STANDARD_VM_FIXED3675_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3675X950g,
	"SUBCORE_STANDARD_VM_FIXED3700_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3700X950g,
	"SUBCORE_STANDARD_VM_FIXED3750_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3750X950g,
	"SUBCORE_STANDARD_VM_FIXED3800_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3800X950g,
	"SUBCORE_STANDARD_VM_FIXED3825_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3825X950g,
	"SUBCORE_STANDARD_VM_FIXED3850_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3850X950g,
	"SUBCORE_STANDARD_VM_FIXED3875_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3875X950g,
	"SUBCORE_STANDARD_VM_FIXED3900_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3900X950g,
	"SUBCORE_STANDARD_VM_FIXED3975_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3975X950g,
	"SUBCORE_STANDARD_VM_FIXED4000_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4000X950g,
	"SUBCORE_STANDARD_VM_FIXED4025_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4025X950g,
	"SUBCORE_STANDARD_VM_FIXED4050_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4050X950g,
	"SUBCORE_STANDARD_VM_FIXED4100_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4100X950g,
	"SUBCORE_STANDARD_VM_FIXED4125_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4125X950g,
	"SUBCORE_STANDARD_VM_FIXED4200_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4200X950g,
	"SUBCORE_STANDARD_VM_FIXED4225_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4225X950g,
	"SUBCORE_STANDARD_VM_FIXED4250_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4250X950g,
	"SUBCORE_STANDARD_VM_FIXED4275_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4275X950g,
	"SUBCORE_STANDARD_VM_FIXED4300_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4300X950g,
	"SUBCORE_STANDARD_VM_FIXED4350_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4350X950g,
	"SUBCORE_STANDARD_VM_FIXED4375_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4375X950g,
	"SUBCORE_STANDARD_VM_FIXED4400_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4400X950g,
	"SUBCORE_STANDARD_VM_FIXED4425_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4425X950g,
	"SUBCORE_STANDARD_VM_FIXED4500_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4500X950g,
	"SUBCORE_STANDARD_VM_FIXED4550_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4550X950g,
	"SUBCORE_STANDARD_VM_FIXED4575_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4575X950g,
	"SUBCORE_STANDARD_VM_FIXED4600_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4600X950g,
	"SUBCORE_STANDARD_VM_FIXED4625_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4625X950g,
	"SUBCORE_STANDARD_VM_FIXED4650_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4650X950g,
	"SUBCORE_STANDARD_VM_FIXED4675_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4675X950g,
	"SUBCORE_STANDARD_VM_FIXED4700_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4700X950g,
	"SUBCORE_STANDARD_VM_FIXED4725_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4725X950g,
	"SUBCORE_STANDARD_VM_FIXED4750_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4750X950g,
	"SUBCORE_STANDARD_VM_FIXED4800_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4800X950g,
	"SUBCORE_STANDARD_VM_FIXED4875_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4875X950g,
	"SUBCORE_STANDARD_VM_FIXED4900_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4900X950g,
	"SUBCORE_STANDARD_VM_FIXED4950_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4950X950g,
	"SUBCORE_STANDARD_VM_FIXED5000_X9_50G": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed5000X950g,
	"ENTIREHOST_X9_50G":                    InternalVnicAttachmentVnicShapeEntirehostX950g,
}

var mappingInternalVnicAttachmentVnicShapeEnumLowerCase = map[string]InternalVnicAttachmentVnicShapeEnum{
	"dynamic":                              InternalVnicAttachmentVnicShapeDynamic,
	"fixed0040":                            InternalVnicAttachmentVnicShapeFixed0040,
	"fixed0060":                            InternalVnicAttachmentVnicShapeFixed0060,
	"fixed0060_psm":                        InternalVnicAttachmentVnicShapeFixed0060Psm,
	"fixed0100":                            InternalVnicAttachmentVnicShapeFixed0100,
	"fixed0120":                            InternalVnicAttachmentVnicShapeFixed0120,
	"fixed0120_2x":                         InternalVnicAttachmentVnicShapeFixed01202x,
	"fixed0200":                            InternalVnicAttachmentVnicShapeFixed0200,
	"fixed0240":                            InternalVnicAttachmentVnicShapeFixed0240,
	"fixed0480":                            InternalVnicAttachmentVnicShapeFixed0480,
	"entirehost":                           InternalVnicAttachmentVnicShapeEntirehost,
	"dynamic_25g":                          InternalVnicAttachmentVnicShapeDynamic25g,
	"fixed0040_25g":                        InternalVnicAttachmentVnicShapeFixed004025g,
	"fixed0100_25g":                        InternalVnicAttachmentVnicShapeFixed010025g,
	"fixed0200_25g":                        InternalVnicAttachmentVnicShapeFixed020025g,
	"fixed0400_25g":                        InternalVnicAttachmentVnicShapeFixed040025g,
	"fixed0800_25g":                        InternalVnicAttachmentVnicShapeFixed080025g,
	"fixed1600_25g":                        InternalVnicAttachmentVnicShapeFixed160025g,
	"fixed2400_25g":                        InternalVnicAttachmentVnicShapeFixed240025g,
	"entirehost_25g":                       InternalVnicAttachmentVnicShapeEntirehost25g,
	"dynamic_e1_25g":                       InternalVnicAttachmentVnicShapeDynamicE125g,
	"fixed0040_e1_25g":                     InternalVnicAttachmentVnicShapeFixed0040E125g,
	"fixed0070_e1_25g":                     InternalVnicAttachmentVnicShapeFixed0070E125g,
	"fixed0140_e1_25g":                     InternalVnicAttachmentVnicShapeFixed0140E125g,
	"fixed0280_e1_25g":                     InternalVnicAttachmentVnicShapeFixed0280E125g,
	"fixed0560_e1_25g":                     InternalVnicAttachmentVnicShapeFixed0560E125g,
	"fixed1120_e1_25g":                     InternalVnicAttachmentVnicShapeFixed1120E125g,
	"fixed1680_e1_25g":                     InternalVnicAttachmentVnicShapeFixed1680E125g,
	"entirehost_e1_25g":                    InternalVnicAttachmentVnicShapeEntirehostE125g,
	"dynamic_b1_25g":                       InternalVnicAttachmentVnicShapeDynamicB125g,
	"fixed0040_b1_25g":                     InternalVnicAttachmentVnicShapeFixed0040B125g,
	"fixed0060_b1_25g":                     InternalVnicAttachmentVnicShapeFixed0060B125g,
	"fixed0120_b1_25g":                     InternalVnicAttachmentVnicShapeFixed0120B125g,
	"fixed0240_b1_25g":                     InternalVnicAttachmentVnicShapeFixed0240B125g,
	"fixed0480_b1_25g":                     InternalVnicAttachmentVnicShapeFixed0480B125g,
	"fixed0960_b1_25g":                     InternalVnicAttachmentVnicShapeFixed0960B125g,
	"entirehost_b1_25g":                    InternalVnicAttachmentVnicShapeEntirehostB125g,
	"micro_vm_fixed0048_e1_25g":            InternalVnicAttachmentVnicShapeMicroVmFixed0048E125g,
	"micro_lb_fixed0001_e1_25g":            InternalVnicAttachmentVnicShapeMicroLbFixed0001E125g,
	"vnicaas_fixed0200":                    InternalVnicAttachmentVnicShapeVnicaasFixed0200,
	"vnicaas_fixed0400":                    InternalVnicAttachmentVnicShapeVnicaasFixed0400,
	"vnicaas_fixed0700":                    InternalVnicAttachmentVnicShapeVnicaasFixed0700,
	"vnicaas_nlb_approved_10g":             InternalVnicAttachmentVnicShapeVnicaasNlbApproved10g,
	"vnicaas_nlb_approved_25g":             InternalVnicAttachmentVnicShapeVnicaasNlbApproved25g,
	"vnicaas_telesis_25g":                  InternalVnicAttachmentVnicShapeVnicaasTelesis25g,
	"vnicaas_telesis_10g":                  InternalVnicAttachmentVnicShapeVnicaasTelesis10g,
	"vnicaas_ambassador_fixed0100":         InternalVnicAttachmentVnicShapeVnicaasAmbassadorFixed0100,
	"vnicaas_telesis_gamma":                InternalVnicAttachmentVnicShapeVnicaasTelesisGamma,
	"vnicaas_privatedns":                   InternalVnicAttachmentVnicShapeVnicaasPrivatedns,
	"vnicaas_fwaas":                        InternalVnicAttachmentVnicShapeVnicaasFwaas,
	"vnicaas_lbaas_free":                   InternalVnicAttachmentVnicShapeVnicaasLbaasFree,
	"vnicaas_lbaas_8g_512k":                InternalVnicAttachmentVnicShapeVnicaasLbaas8g512k,
	"vnicaas_lbaas_8g_1m":                  InternalVnicAttachmentVnicShapeVnicaasLbaas8g1m,
	"vnicaas_lbaas_8g_2m":                  InternalVnicAttachmentVnicShapeVnicaasLbaas8g2m,
	"vnicaas_lbaas_8g_3m":                  InternalVnicAttachmentVnicShapeVnicaasLbaas8g3m,
	"vnicaas_lbaas_8g_1m_8ghost":           InternalVnicAttachmentVnicShapeVnicaasLbaas8g1m8ghost,
	"vnicaas_lbaas_8g_1m_16ghost":          InternalVnicAttachmentVnicShapeVnicaasLbaas8g1m16ghost,
	"dynamic_e3_50g":                       InternalVnicAttachmentVnicShapeDynamicE350g,
	"fixed0040_e3_50g":                     InternalVnicAttachmentVnicShapeFixed0040E350g,
	"fixed0100_e3_50g":                     InternalVnicAttachmentVnicShapeFixed0100E350g,
	"fixed0200_e3_50g":                     InternalVnicAttachmentVnicShapeFixed0200E350g,
	"fixed0300_e3_50g":                     InternalVnicAttachmentVnicShapeFixed0300E350g,
	"fixed0400_e3_50g":                     InternalVnicAttachmentVnicShapeFixed0400E350g,
	"fixed0500_e3_50g":                     InternalVnicAttachmentVnicShapeFixed0500E350g,
	"fixed0600_e3_50g":                     InternalVnicAttachmentVnicShapeFixed0600E350g,
	"fixed0700_e3_50g":                     InternalVnicAttachmentVnicShapeFixed0700E350g,
	"fixed0800_e3_50g":                     InternalVnicAttachmentVnicShapeFixed0800E350g,
	"fixed0900_e3_50g":                     InternalVnicAttachmentVnicShapeFixed0900E350g,
	"fixed1000_e3_50g":                     InternalVnicAttachmentVnicShapeFixed1000E350g,
	"fixed1100_e3_50g":                     InternalVnicAttachmentVnicShapeFixed1100E350g,
	"fixed1200_e3_50g":                     InternalVnicAttachmentVnicShapeFixed1200E350g,
	"fixed1300_e3_50g":                     InternalVnicAttachmentVnicShapeFixed1300E350g,
	"fixed1400_e3_50g":                     InternalVnicAttachmentVnicShapeFixed1400E350g,
	"fixed1500_e3_50g":                     InternalVnicAttachmentVnicShapeFixed1500E350g,
	"fixed1600_e3_50g":                     InternalVnicAttachmentVnicShapeFixed1600E350g,
	"fixed1700_e3_50g":                     InternalVnicAttachmentVnicShapeFixed1700E350g,
	"fixed1800_e3_50g":                     InternalVnicAttachmentVnicShapeFixed1800E350g,
	"fixed1900_e3_50g":                     InternalVnicAttachmentVnicShapeFixed1900E350g,
	"fixed2000_e3_50g":                     InternalVnicAttachmentVnicShapeFixed2000E350g,
	"fixed2100_e3_50g":                     InternalVnicAttachmentVnicShapeFixed2100E350g,
	"fixed2200_e3_50g":                     InternalVnicAttachmentVnicShapeFixed2200E350g,
	"fixed2300_e3_50g":                     InternalVnicAttachmentVnicShapeFixed2300E350g,
	"fixed2400_e3_50g":                     InternalVnicAttachmentVnicShapeFixed2400E350g,
	"fixed2500_e3_50g":                     InternalVnicAttachmentVnicShapeFixed2500E350g,
	"fixed2600_e3_50g":                     InternalVnicAttachmentVnicShapeFixed2600E350g,
	"fixed2700_e3_50g":                     InternalVnicAttachmentVnicShapeFixed2700E350g,
	"fixed2800_e3_50g":                     InternalVnicAttachmentVnicShapeFixed2800E350g,
	"fixed2900_e3_50g":                     InternalVnicAttachmentVnicShapeFixed2900E350g,
	"fixed3000_e3_50g":                     InternalVnicAttachmentVnicShapeFixed3000E350g,
	"fixed3100_e3_50g":                     InternalVnicAttachmentVnicShapeFixed3100E350g,
	"fixed3200_e3_50g":                     InternalVnicAttachmentVnicShapeFixed3200E350g,
	"fixed3300_e3_50g":                     InternalVnicAttachmentVnicShapeFixed3300E350g,
	"fixed3400_e3_50g":                     InternalVnicAttachmentVnicShapeFixed3400E350g,
	"fixed3500_e3_50g":                     InternalVnicAttachmentVnicShapeFixed3500E350g,
	"fixed3600_e3_50g":                     InternalVnicAttachmentVnicShapeFixed3600E350g,
	"fixed3700_e3_50g":                     InternalVnicAttachmentVnicShapeFixed3700E350g,
	"fixed3800_e3_50g":                     InternalVnicAttachmentVnicShapeFixed3800E350g,
	"fixed3900_e3_50g":                     InternalVnicAttachmentVnicShapeFixed3900E350g,
	"fixed4000_e3_50g":                     InternalVnicAttachmentVnicShapeFixed4000E350g,
	"entirehost_e3_50g":                    InternalVnicAttachmentVnicShapeEntirehostE350g,
	"dynamic_e4_50g":                       InternalVnicAttachmentVnicShapeDynamicE450g,
	"fixed0040_e4_50g":                     InternalVnicAttachmentVnicShapeFixed0040E450g,
	"fixed0100_e4_50g":                     InternalVnicAttachmentVnicShapeFixed0100E450g,
	"fixed0200_e4_50g":                     InternalVnicAttachmentVnicShapeFixed0200E450g,
	"fixed0300_e4_50g":                     InternalVnicAttachmentVnicShapeFixed0300E450g,
	"fixed0400_e4_50g":                     InternalVnicAttachmentVnicShapeFixed0400E450g,
	"fixed0500_e4_50g":                     InternalVnicAttachmentVnicShapeFixed0500E450g,
	"fixed0600_e4_50g":                     InternalVnicAttachmentVnicShapeFixed0600E450g,
	"fixed0700_e4_50g":                     InternalVnicAttachmentVnicShapeFixed0700E450g,
	"fixed0800_e4_50g":                     InternalVnicAttachmentVnicShapeFixed0800E450g,
	"fixed0900_e4_50g":                     InternalVnicAttachmentVnicShapeFixed0900E450g,
	"fixed1000_e4_50g":                     InternalVnicAttachmentVnicShapeFixed1000E450g,
	"fixed1100_e4_50g":                     InternalVnicAttachmentVnicShapeFixed1100E450g,
	"fixed1200_e4_50g":                     InternalVnicAttachmentVnicShapeFixed1200E450g,
	"fixed1300_e4_50g":                     InternalVnicAttachmentVnicShapeFixed1300E450g,
	"fixed1400_e4_50g":                     InternalVnicAttachmentVnicShapeFixed1400E450g,
	"fixed1500_e4_50g":                     InternalVnicAttachmentVnicShapeFixed1500E450g,
	"fixed1600_e4_50g":                     InternalVnicAttachmentVnicShapeFixed1600E450g,
	"fixed1700_e4_50g":                     InternalVnicAttachmentVnicShapeFixed1700E450g,
	"fixed1800_e4_50g":                     InternalVnicAttachmentVnicShapeFixed1800E450g,
	"fixed1900_e4_50g":                     InternalVnicAttachmentVnicShapeFixed1900E450g,
	"fixed2000_e4_50g":                     InternalVnicAttachmentVnicShapeFixed2000E450g,
	"fixed2100_e4_50g":                     InternalVnicAttachmentVnicShapeFixed2100E450g,
	"fixed2200_e4_50g":                     InternalVnicAttachmentVnicShapeFixed2200E450g,
	"fixed2300_e4_50g":                     InternalVnicAttachmentVnicShapeFixed2300E450g,
	"fixed2400_e4_50g":                     InternalVnicAttachmentVnicShapeFixed2400E450g,
	"fixed2500_e4_50g":                     InternalVnicAttachmentVnicShapeFixed2500E450g,
	"fixed2600_e4_50g":                     InternalVnicAttachmentVnicShapeFixed2600E450g,
	"fixed2700_e4_50g":                     InternalVnicAttachmentVnicShapeFixed2700E450g,
	"fixed2800_e4_50g":                     InternalVnicAttachmentVnicShapeFixed2800E450g,
	"fixed2900_e4_50g":                     InternalVnicAttachmentVnicShapeFixed2900E450g,
	"fixed3000_e4_50g":                     InternalVnicAttachmentVnicShapeFixed3000E450g,
	"fixed3100_e4_50g":                     InternalVnicAttachmentVnicShapeFixed3100E450g,
	"fixed3200_e4_50g":                     InternalVnicAttachmentVnicShapeFixed3200E450g,
	"fixed3300_e4_50g":                     InternalVnicAttachmentVnicShapeFixed3300E450g,
	"fixed3400_e4_50g":                     InternalVnicAttachmentVnicShapeFixed3400E450g,
	"fixed3500_e4_50g":                     InternalVnicAttachmentVnicShapeFixed3500E450g,
	"fixed3600_e4_50g":                     InternalVnicAttachmentVnicShapeFixed3600E450g,
	"fixed3700_e4_50g":                     InternalVnicAttachmentVnicShapeFixed3700E450g,
	"fixed3800_e4_50g":                     InternalVnicAttachmentVnicShapeFixed3800E450g,
	"fixed3900_e4_50g":                     InternalVnicAttachmentVnicShapeFixed3900E450g,
	"fixed4000_e4_50g":                     InternalVnicAttachmentVnicShapeFixed4000E450g,
	"entirehost_e4_50g":                    InternalVnicAttachmentVnicShapeEntirehostE450g,
	"micro_vm_fixed0050_e3_50g":            InternalVnicAttachmentVnicShapeMicroVmFixed0050E350g,
	"micro_vm_fixed0050_e4_50g":            InternalVnicAttachmentVnicShapeMicroVmFixed0050E450g,
	"subcore_vm_fixed0025_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0025E350g,
	"subcore_vm_fixed0050_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0050E350g,
	"subcore_vm_fixed0075_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0075E350g,
	"subcore_vm_fixed0100_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0100E350g,
	"subcore_vm_fixed0125_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0125E350g,
	"subcore_vm_fixed0150_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0150E350g,
	"subcore_vm_fixed0175_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0175E350g,
	"subcore_vm_fixed0200_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0200E350g,
	"subcore_vm_fixed0225_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0225E350g,
	"subcore_vm_fixed0250_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0250E350g,
	"subcore_vm_fixed0275_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0275E350g,
	"subcore_vm_fixed0300_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0300E350g,
	"subcore_vm_fixed0325_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0325E350g,
	"subcore_vm_fixed0350_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0350E350g,
	"subcore_vm_fixed0375_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0375E350g,
	"subcore_vm_fixed0400_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0400E350g,
	"subcore_vm_fixed0425_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0425E350g,
	"subcore_vm_fixed0450_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0450E350g,
	"subcore_vm_fixed0475_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0475E350g,
	"subcore_vm_fixed0500_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0500E350g,
	"subcore_vm_fixed0525_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0525E350g,
	"subcore_vm_fixed0550_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0550E350g,
	"subcore_vm_fixed0575_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0575E350g,
	"subcore_vm_fixed0600_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0600E350g,
	"subcore_vm_fixed0625_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0625E350g,
	"subcore_vm_fixed0650_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0650E350g,
	"subcore_vm_fixed0675_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0675E350g,
	"subcore_vm_fixed0700_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0700E350g,
	"subcore_vm_fixed0725_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0725E350g,
	"subcore_vm_fixed0750_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0750E350g,
	"subcore_vm_fixed0775_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0775E350g,
	"subcore_vm_fixed0800_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0800E350g,
	"subcore_vm_fixed0825_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0825E350g,
	"subcore_vm_fixed0850_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0850E350g,
	"subcore_vm_fixed0875_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0875E350g,
	"subcore_vm_fixed0900_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0900E350g,
	"subcore_vm_fixed0925_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0925E350g,
	"subcore_vm_fixed0950_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0950E350g,
	"subcore_vm_fixed0975_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0975E350g,
	"subcore_vm_fixed1000_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1000E350g,
	"subcore_vm_fixed1025_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1025E350g,
	"subcore_vm_fixed1050_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1050E350g,
	"subcore_vm_fixed1075_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1075E350g,
	"subcore_vm_fixed1100_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1100E350g,
	"subcore_vm_fixed1125_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1125E350g,
	"subcore_vm_fixed1150_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1150E350g,
	"subcore_vm_fixed1175_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1175E350g,
	"subcore_vm_fixed1200_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1200E350g,
	"subcore_vm_fixed1225_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1225E350g,
	"subcore_vm_fixed1250_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1250E350g,
	"subcore_vm_fixed1275_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1275E350g,
	"subcore_vm_fixed1300_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1300E350g,
	"subcore_vm_fixed1325_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1325E350g,
	"subcore_vm_fixed1350_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1350E350g,
	"subcore_vm_fixed1375_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1375E350g,
	"subcore_vm_fixed1400_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1400E350g,
	"subcore_vm_fixed1425_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1425E350g,
	"subcore_vm_fixed1450_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1450E350g,
	"subcore_vm_fixed1475_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1475E350g,
	"subcore_vm_fixed1500_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1500E350g,
	"subcore_vm_fixed1525_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1525E350g,
	"subcore_vm_fixed1550_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1550E350g,
	"subcore_vm_fixed1575_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1575E350g,
	"subcore_vm_fixed1600_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1600E350g,
	"subcore_vm_fixed1625_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1625E350g,
	"subcore_vm_fixed1650_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1650E350g,
	"subcore_vm_fixed1700_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1700E350g,
	"subcore_vm_fixed1725_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1725E350g,
	"subcore_vm_fixed1750_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1750E350g,
	"subcore_vm_fixed1800_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1800E350g,
	"subcore_vm_fixed1850_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1850E350g,
	"subcore_vm_fixed1875_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1875E350g,
	"subcore_vm_fixed1900_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1900E350g,
	"subcore_vm_fixed1925_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1925E350g,
	"subcore_vm_fixed1950_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1950E350g,
	"subcore_vm_fixed2000_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2000E350g,
	"subcore_vm_fixed2025_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2025E350g,
	"subcore_vm_fixed2050_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2050E350g,
	"subcore_vm_fixed2100_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2100E350g,
	"subcore_vm_fixed2125_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2125E350g,
	"subcore_vm_fixed2150_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2150E350g,
	"subcore_vm_fixed2175_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2175E350g,
	"subcore_vm_fixed2200_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2200E350g,
	"subcore_vm_fixed2250_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2250E350g,
	"subcore_vm_fixed2275_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2275E350g,
	"subcore_vm_fixed2300_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2300E350g,
	"subcore_vm_fixed2325_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2325E350g,
	"subcore_vm_fixed2350_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2350E350g,
	"subcore_vm_fixed2375_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2375E350g,
	"subcore_vm_fixed2400_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2400E350g,
	"subcore_vm_fixed2450_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2450E350g,
	"subcore_vm_fixed2475_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2475E350g,
	"subcore_vm_fixed2500_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2500E350g,
	"subcore_vm_fixed2550_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2550E350g,
	"subcore_vm_fixed2600_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2600E350g,
	"subcore_vm_fixed2625_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2625E350g,
	"subcore_vm_fixed2650_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2650E350g,
	"subcore_vm_fixed2700_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2700E350g,
	"subcore_vm_fixed2750_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2750E350g,
	"subcore_vm_fixed2775_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2775E350g,
	"subcore_vm_fixed2800_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2800E350g,
	"subcore_vm_fixed2850_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2850E350g,
	"subcore_vm_fixed2875_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2875E350g,
	"subcore_vm_fixed2900_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2900E350g,
	"subcore_vm_fixed2925_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2925E350g,
	"subcore_vm_fixed2950_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2950E350g,
	"subcore_vm_fixed2975_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2975E350g,
	"subcore_vm_fixed3000_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3000E350g,
	"subcore_vm_fixed3025_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3025E350g,
	"subcore_vm_fixed3050_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3050E350g,
	"subcore_vm_fixed3075_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3075E350g,
	"subcore_vm_fixed3100_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3100E350g,
	"subcore_vm_fixed3125_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3125E350g,
	"subcore_vm_fixed3150_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3150E350g,
	"subcore_vm_fixed3200_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3200E350g,
	"subcore_vm_fixed3225_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3225E350g,
	"subcore_vm_fixed3250_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3250E350g,
	"subcore_vm_fixed3300_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3300E350g,
	"subcore_vm_fixed3325_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3325E350g,
	"subcore_vm_fixed3375_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3375E350g,
	"subcore_vm_fixed3400_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3400E350g,
	"subcore_vm_fixed3450_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3450E350g,
	"subcore_vm_fixed3500_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3500E350g,
	"subcore_vm_fixed3525_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3525E350g,
	"subcore_vm_fixed3575_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3575E350g,
	"subcore_vm_fixed3600_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3600E350g,
	"subcore_vm_fixed3625_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3625E350g,
	"subcore_vm_fixed3675_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3675E350g,
	"subcore_vm_fixed3700_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3700E350g,
	"subcore_vm_fixed3750_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3750E350g,
	"subcore_vm_fixed3800_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3800E350g,
	"subcore_vm_fixed3825_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3825E350g,
	"subcore_vm_fixed3850_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3850E350g,
	"subcore_vm_fixed3875_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3875E350g,
	"subcore_vm_fixed3900_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3900E350g,
	"subcore_vm_fixed3975_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3975E350g,
	"subcore_vm_fixed4000_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4000E350g,
	"subcore_vm_fixed4025_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4025E350g,
	"subcore_vm_fixed4050_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4050E350g,
	"subcore_vm_fixed4100_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4100E350g,
	"subcore_vm_fixed4125_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4125E350g,
	"subcore_vm_fixed4200_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4200E350g,
	"subcore_vm_fixed4225_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4225E350g,
	"subcore_vm_fixed4250_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4250E350g,
	"subcore_vm_fixed4275_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4275E350g,
	"subcore_vm_fixed4300_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4300E350g,
	"subcore_vm_fixed4350_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4350E350g,
	"subcore_vm_fixed4375_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4375E350g,
	"subcore_vm_fixed4400_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4400E350g,
	"subcore_vm_fixed4425_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4425E350g,
	"subcore_vm_fixed4500_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4500E350g,
	"subcore_vm_fixed4550_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4550E350g,
	"subcore_vm_fixed4575_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4575E350g,
	"subcore_vm_fixed4600_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4600E350g,
	"subcore_vm_fixed4625_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4625E350g,
	"subcore_vm_fixed4650_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4650E350g,
	"subcore_vm_fixed4675_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4675E350g,
	"subcore_vm_fixed4700_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4700E350g,
	"subcore_vm_fixed4725_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4725E350g,
	"subcore_vm_fixed4750_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4750E350g,
	"subcore_vm_fixed4800_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4800E350g,
	"subcore_vm_fixed4875_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4875E350g,
	"subcore_vm_fixed4900_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4900E350g,
	"subcore_vm_fixed4950_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4950E350g,
	"subcore_vm_fixed5000_e3_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed5000E350g,
	"subcore_vm_fixed0025_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0025E450g,
	"subcore_vm_fixed0050_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0050E450g,
	"subcore_vm_fixed0075_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0075E450g,
	"subcore_vm_fixed0100_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0100E450g,
	"subcore_vm_fixed0125_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0125E450g,
	"subcore_vm_fixed0150_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0150E450g,
	"subcore_vm_fixed0175_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0175E450g,
	"subcore_vm_fixed0200_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0200E450g,
	"subcore_vm_fixed0225_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0225E450g,
	"subcore_vm_fixed0250_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0250E450g,
	"subcore_vm_fixed0275_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0275E450g,
	"subcore_vm_fixed0300_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0300E450g,
	"subcore_vm_fixed0325_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0325E450g,
	"subcore_vm_fixed0350_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0350E450g,
	"subcore_vm_fixed0375_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0375E450g,
	"subcore_vm_fixed0400_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0400E450g,
	"subcore_vm_fixed0425_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0425E450g,
	"subcore_vm_fixed0450_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0450E450g,
	"subcore_vm_fixed0475_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0475E450g,
	"subcore_vm_fixed0500_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0500E450g,
	"subcore_vm_fixed0525_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0525E450g,
	"subcore_vm_fixed0550_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0550E450g,
	"subcore_vm_fixed0575_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0575E450g,
	"subcore_vm_fixed0600_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0600E450g,
	"subcore_vm_fixed0625_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0625E450g,
	"subcore_vm_fixed0650_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0650E450g,
	"subcore_vm_fixed0675_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0675E450g,
	"subcore_vm_fixed0700_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0700E450g,
	"subcore_vm_fixed0725_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0725E450g,
	"subcore_vm_fixed0750_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0750E450g,
	"subcore_vm_fixed0775_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0775E450g,
	"subcore_vm_fixed0800_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0800E450g,
	"subcore_vm_fixed0825_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0825E450g,
	"subcore_vm_fixed0850_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0850E450g,
	"subcore_vm_fixed0875_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0875E450g,
	"subcore_vm_fixed0900_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0900E450g,
	"subcore_vm_fixed0925_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0925E450g,
	"subcore_vm_fixed0950_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0950E450g,
	"subcore_vm_fixed0975_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0975E450g,
	"subcore_vm_fixed1000_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1000E450g,
	"subcore_vm_fixed1025_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1025E450g,
	"subcore_vm_fixed1050_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1050E450g,
	"subcore_vm_fixed1075_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1075E450g,
	"subcore_vm_fixed1100_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1100E450g,
	"subcore_vm_fixed1125_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1125E450g,
	"subcore_vm_fixed1150_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1150E450g,
	"subcore_vm_fixed1175_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1175E450g,
	"subcore_vm_fixed1200_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1200E450g,
	"subcore_vm_fixed1225_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1225E450g,
	"subcore_vm_fixed1250_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1250E450g,
	"subcore_vm_fixed1275_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1275E450g,
	"subcore_vm_fixed1300_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1300E450g,
	"subcore_vm_fixed1325_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1325E450g,
	"subcore_vm_fixed1350_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1350E450g,
	"subcore_vm_fixed1375_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1375E450g,
	"subcore_vm_fixed1400_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1400E450g,
	"subcore_vm_fixed1425_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1425E450g,
	"subcore_vm_fixed1450_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1450E450g,
	"subcore_vm_fixed1475_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1475E450g,
	"subcore_vm_fixed1500_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1500E450g,
	"subcore_vm_fixed1525_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1525E450g,
	"subcore_vm_fixed1550_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1550E450g,
	"subcore_vm_fixed1575_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1575E450g,
	"subcore_vm_fixed1600_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1600E450g,
	"subcore_vm_fixed1625_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1625E450g,
	"subcore_vm_fixed1650_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1650E450g,
	"subcore_vm_fixed1700_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1700E450g,
	"subcore_vm_fixed1725_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1725E450g,
	"subcore_vm_fixed1750_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1750E450g,
	"subcore_vm_fixed1800_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1800E450g,
	"subcore_vm_fixed1850_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1850E450g,
	"subcore_vm_fixed1875_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1875E450g,
	"subcore_vm_fixed1900_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1900E450g,
	"subcore_vm_fixed1925_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1925E450g,
	"subcore_vm_fixed1950_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1950E450g,
	"subcore_vm_fixed2000_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2000E450g,
	"subcore_vm_fixed2025_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2025E450g,
	"subcore_vm_fixed2050_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2050E450g,
	"subcore_vm_fixed2100_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2100E450g,
	"subcore_vm_fixed2125_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2125E450g,
	"subcore_vm_fixed2150_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2150E450g,
	"subcore_vm_fixed2175_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2175E450g,
	"subcore_vm_fixed2200_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2200E450g,
	"subcore_vm_fixed2250_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2250E450g,
	"subcore_vm_fixed2275_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2275E450g,
	"subcore_vm_fixed2300_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2300E450g,
	"subcore_vm_fixed2325_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2325E450g,
	"subcore_vm_fixed2350_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2350E450g,
	"subcore_vm_fixed2375_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2375E450g,
	"subcore_vm_fixed2400_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2400E450g,
	"subcore_vm_fixed2450_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2450E450g,
	"subcore_vm_fixed2475_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2475E450g,
	"subcore_vm_fixed2500_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2500E450g,
	"subcore_vm_fixed2550_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2550E450g,
	"subcore_vm_fixed2600_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2600E450g,
	"subcore_vm_fixed2625_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2625E450g,
	"subcore_vm_fixed2650_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2650E450g,
	"subcore_vm_fixed2700_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2700E450g,
	"subcore_vm_fixed2750_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2750E450g,
	"subcore_vm_fixed2775_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2775E450g,
	"subcore_vm_fixed2800_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2800E450g,
	"subcore_vm_fixed2850_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2850E450g,
	"subcore_vm_fixed2875_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2875E450g,
	"subcore_vm_fixed2900_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2900E450g,
	"subcore_vm_fixed2925_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2925E450g,
	"subcore_vm_fixed2950_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2950E450g,
	"subcore_vm_fixed2975_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2975E450g,
	"subcore_vm_fixed3000_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3000E450g,
	"subcore_vm_fixed3025_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3025E450g,
	"subcore_vm_fixed3050_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3050E450g,
	"subcore_vm_fixed3075_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3075E450g,
	"subcore_vm_fixed3100_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3100E450g,
	"subcore_vm_fixed3125_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3125E450g,
	"subcore_vm_fixed3150_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3150E450g,
	"subcore_vm_fixed3200_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3200E450g,
	"subcore_vm_fixed3225_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3225E450g,
	"subcore_vm_fixed3250_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3250E450g,
	"subcore_vm_fixed3300_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3300E450g,
	"subcore_vm_fixed3325_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3325E450g,
	"subcore_vm_fixed3375_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3375E450g,
	"subcore_vm_fixed3400_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3400E450g,
	"subcore_vm_fixed3450_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3450E450g,
	"subcore_vm_fixed3500_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3500E450g,
	"subcore_vm_fixed3525_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3525E450g,
	"subcore_vm_fixed3575_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3575E450g,
	"subcore_vm_fixed3600_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3600E450g,
	"subcore_vm_fixed3625_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3625E450g,
	"subcore_vm_fixed3675_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3675E450g,
	"subcore_vm_fixed3700_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3700E450g,
	"subcore_vm_fixed3750_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3750E450g,
	"subcore_vm_fixed3800_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3800E450g,
	"subcore_vm_fixed3825_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3825E450g,
	"subcore_vm_fixed3850_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3850E450g,
	"subcore_vm_fixed3875_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3875E450g,
	"subcore_vm_fixed3900_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3900E450g,
	"subcore_vm_fixed3975_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3975E450g,
	"subcore_vm_fixed4000_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4000E450g,
	"subcore_vm_fixed4025_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4025E450g,
	"subcore_vm_fixed4050_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4050E450g,
	"subcore_vm_fixed4100_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4100E450g,
	"subcore_vm_fixed4125_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4125E450g,
	"subcore_vm_fixed4200_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4200E450g,
	"subcore_vm_fixed4225_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4225E450g,
	"subcore_vm_fixed4250_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4250E450g,
	"subcore_vm_fixed4275_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4275E450g,
	"subcore_vm_fixed4300_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4300E450g,
	"subcore_vm_fixed4350_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4350E450g,
	"subcore_vm_fixed4375_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4375E450g,
	"subcore_vm_fixed4400_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4400E450g,
	"subcore_vm_fixed4425_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4425E450g,
	"subcore_vm_fixed4500_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4500E450g,
	"subcore_vm_fixed4550_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4550E450g,
	"subcore_vm_fixed4575_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4575E450g,
	"subcore_vm_fixed4600_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4600E450g,
	"subcore_vm_fixed4625_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4625E450g,
	"subcore_vm_fixed4650_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4650E450g,
	"subcore_vm_fixed4675_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4675E450g,
	"subcore_vm_fixed4700_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4700E450g,
	"subcore_vm_fixed4725_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4725E450g,
	"subcore_vm_fixed4750_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4750E450g,
	"subcore_vm_fixed4800_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4800E450g,
	"subcore_vm_fixed4875_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4875E450g,
	"subcore_vm_fixed4900_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4900E450g,
	"subcore_vm_fixed4950_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4950E450g,
	"subcore_vm_fixed5000_e4_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed5000E450g,
	"subcore_vm_fixed0020_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0020A150g,
	"subcore_vm_fixed0040_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0040A150g,
	"subcore_vm_fixed0060_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0060A150g,
	"subcore_vm_fixed0080_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0080A150g,
	"subcore_vm_fixed0100_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0100A150g,
	"subcore_vm_fixed0120_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0120A150g,
	"subcore_vm_fixed0140_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0140A150g,
	"subcore_vm_fixed0160_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0160A150g,
	"subcore_vm_fixed0180_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0180A150g,
	"subcore_vm_fixed0200_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0200A150g,
	"subcore_vm_fixed0220_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0220A150g,
	"subcore_vm_fixed0240_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0240A150g,
	"subcore_vm_fixed0260_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0260A150g,
	"subcore_vm_fixed0280_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0280A150g,
	"subcore_vm_fixed0300_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0300A150g,
	"subcore_vm_fixed0320_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0320A150g,
	"subcore_vm_fixed0340_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0340A150g,
	"subcore_vm_fixed0360_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0360A150g,
	"subcore_vm_fixed0380_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0380A150g,
	"subcore_vm_fixed0400_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0400A150g,
	"subcore_vm_fixed0420_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0420A150g,
	"subcore_vm_fixed0440_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0440A150g,
	"subcore_vm_fixed0460_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0460A150g,
	"subcore_vm_fixed0480_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0480A150g,
	"subcore_vm_fixed0500_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0500A150g,
	"subcore_vm_fixed0520_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0520A150g,
	"subcore_vm_fixed0540_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0540A150g,
	"subcore_vm_fixed0560_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0560A150g,
	"subcore_vm_fixed0580_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0580A150g,
	"subcore_vm_fixed0600_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0600A150g,
	"subcore_vm_fixed0620_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0620A150g,
	"subcore_vm_fixed0640_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0640A150g,
	"subcore_vm_fixed0660_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0660A150g,
	"subcore_vm_fixed0680_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0680A150g,
	"subcore_vm_fixed0700_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0700A150g,
	"subcore_vm_fixed0720_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0720A150g,
	"subcore_vm_fixed0740_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0740A150g,
	"subcore_vm_fixed0760_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0760A150g,
	"subcore_vm_fixed0780_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0780A150g,
	"subcore_vm_fixed0800_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0800A150g,
	"subcore_vm_fixed0820_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0820A150g,
	"subcore_vm_fixed0840_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0840A150g,
	"subcore_vm_fixed0860_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0860A150g,
	"subcore_vm_fixed0880_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0880A150g,
	"subcore_vm_fixed0900_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0900A150g,
	"subcore_vm_fixed0920_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0920A150g,
	"subcore_vm_fixed0940_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0940A150g,
	"subcore_vm_fixed0960_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0960A150g,
	"subcore_vm_fixed0980_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0980A150g,
	"subcore_vm_fixed1000_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1000A150g,
	"subcore_vm_fixed1020_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1020A150g,
	"subcore_vm_fixed1040_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1040A150g,
	"subcore_vm_fixed1060_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1060A150g,
	"subcore_vm_fixed1080_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1080A150g,
	"subcore_vm_fixed1100_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1100A150g,
	"subcore_vm_fixed1120_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1120A150g,
	"subcore_vm_fixed1140_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1140A150g,
	"subcore_vm_fixed1160_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1160A150g,
	"subcore_vm_fixed1180_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1180A150g,
	"subcore_vm_fixed1200_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1200A150g,
	"subcore_vm_fixed1220_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1220A150g,
	"subcore_vm_fixed1240_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1240A150g,
	"subcore_vm_fixed1260_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1260A150g,
	"subcore_vm_fixed1280_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1280A150g,
	"subcore_vm_fixed1300_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1300A150g,
	"subcore_vm_fixed1320_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1320A150g,
	"subcore_vm_fixed1340_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1340A150g,
	"subcore_vm_fixed1360_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1360A150g,
	"subcore_vm_fixed1380_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1380A150g,
	"subcore_vm_fixed1400_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1400A150g,
	"subcore_vm_fixed1420_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1420A150g,
	"subcore_vm_fixed1440_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1440A150g,
	"subcore_vm_fixed1460_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1460A150g,
	"subcore_vm_fixed1480_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1480A150g,
	"subcore_vm_fixed1500_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1500A150g,
	"subcore_vm_fixed1520_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1520A150g,
	"subcore_vm_fixed1540_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1540A150g,
	"subcore_vm_fixed1560_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1560A150g,
	"subcore_vm_fixed1580_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1580A150g,
	"subcore_vm_fixed1600_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1600A150g,
	"subcore_vm_fixed1620_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1620A150g,
	"subcore_vm_fixed1640_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1640A150g,
	"subcore_vm_fixed1660_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1660A150g,
	"subcore_vm_fixed1680_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1680A150g,
	"subcore_vm_fixed1700_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1700A150g,
	"subcore_vm_fixed1720_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1720A150g,
	"subcore_vm_fixed1740_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1740A150g,
	"subcore_vm_fixed1760_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1760A150g,
	"subcore_vm_fixed1780_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1780A150g,
	"subcore_vm_fixed1800_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1800A150g,
	"subcore_vm_fixed1820_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1820A150g,
	"subcore_vm_fixed1840_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1840A150g,
	"subcore_vm_fixed1860_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1860A150g,
	"subcore_vm_fixed1880_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1880A150g,
	"subcore_vm_fixed1900_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1900A150g,
	"subcore_vm_fixed1920_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1920A150g,
	"subcore_vm_fixed1940_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1940A150g,
	"subcore_vm_fixed1960_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1960A150g,
	"subcore_vm_fixed1980_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1980A150g,
	"subcore_vm_fixed2000_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2000A150g,
	"subcore_vm_fixed2020_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2020A150g,
	"subcore_vm_fixed2040_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2040A150g,
	"subcore_vm_fixed2060_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2060A150g,
	"subcore_vm_fixed2080_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2080A150g,
	"subcore_vm_fixed2100_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2100A150g,
	"subcore_vm_fixed2120_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2120A150g,
	"subcore_vm_fixed2140_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2140A150g,
	"subcore_vm_fixed2160_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2160A150g,
	"subcore_vm_fixed2180_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2180A150g,
	"subcore_vm_fixed2200_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2200A150g,
	"subcore_vm_fixed2220_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2220A150g,
	"subcore_vm_fixed2240_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2240A150g,
	"subcore_vm_fixed2260_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2260A150g,
	"subcore_vm_fixed2280_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2280A150g,
	"subcore_vm_fixed2300_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2300A150g,
	"subcore_vm_fixed2320_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2320A150g,
	"subcore_vm_fixed2340_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2340A150g,
	"subcore_vm_fixed2360_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2360A150g,
	"subcore_vm_fixed2380_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2380A150g,
	"subcore_vm_fixed2400_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2400A150g,
	"subcore_vm_fixed2420_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2420A150g,
	"subcore_vm_fixed2440_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2440A150g,
	"subcore_vm_fixed2460_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2460A150g,
	"subcore_vm_fixed2480_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2480A150g,
	"subcore_vm_fixed2500_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2500A150g,
	"subcore_vm_fixed2520_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2520A150g,
	"subcore_vm_fixed2540_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2540A150g,
	"subcore_vm_fixed2560_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2560A150g,
	"subcore_vm_fixed2580_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2580A150g,
	"subcore_vm_fixed2600_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2600A150g,
	"subcore_vm_fixed2620_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2620A150g,
	"subcore_vm_fixed2640_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2640A150g,
	"subcore_vm_fixed2660_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2660A150g,
	"subcore_vm_fixed2680_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2680A150g,
	"subcore_vm_fixed2700_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2700A150g,
	"subcore_vm_fixed2720_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2720A150g,
	"subcore_vm_fixed2740_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2740A150g,
	"subcore_vm_fixed2760_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2760A150g,
	"subcore_vm_fixed2780_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2780A150g,
	"subcore_vm_fixed2800_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2800A150g,
	"subcore_vm_fixed2820_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2820A150g,
	"subcore_vm_fixed2840_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2840A150g,
	"subcore_vm_fixed2860_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2860A150g,
	"subcore_vm_fixed2880_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2880A150g,
	"subcore_vm_fixed2900_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2900A150g,
	"subcore_vm_fixed2920_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2920A150g,
	"subcore_vm_fixed2940_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2940A150g,
	"subcore_vm_fixed2960_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2960A150g,
	"subcore_vm_fixed2980_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2980A150g,
	"subcore_vm_fixed3000_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3000A150g,
	"subcore_vm_fixed3020_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3020A150g,
	"subcore_vm_fixed3040_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3040A150g,
	"subcore_vm_fixed3060_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3060A150g,
	"subcore_vm_fixed3080_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3080A150g,
	"subcore_vm_fixed3100_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3100A150g,
	"subcore_vm_fixed3120_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3120A150g,
	"subcore_vm_fixed3140_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3140A150g,
	"subcore_vm_fixed3160_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3160A150g,
	"subcore_vm_fixed3180_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3180A150g,
	"subcore_vm_fixed3200_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3200A150g,
	"subcore_vm_fixed3220_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3220A150g,
	"subcore_vm_fixed3240_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3240A150g,
	"subcore_vm_fixed3260_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3260A150g,
	"subcore_vm_fixed3280_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3280A150g,
	"subcore_vm_fixed3300_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3300A150g,
	"subcore_vm_fixed3320_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3320A150g,
	"subcore_vm_fixed3340_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3340A150g,
	"subcore_vm_fixed3360_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3360A150g,
	"subcore_vm_fixed3380_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3380A150g,
	"subcore_vm_fixed3400_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3400A150g,
	"subcore_vm_fixed3420_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3420A150g,
	"subcore_vm_fixed3440_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3440A150g,
	"subcore_vm_fixed3460_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3460A150g,
	"subcore_vm_fixed3480_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3480A150g,
	"subcore_vm_fixed3500_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3500A150g,
	"subcore_vm_fixed3520_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3520A150g,
	"subcore_vm_fixed3540_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3540A150g,
	"subcore_vm_fixed3560_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3560A150g,
	"subcore_vm_fixed3580_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3580A150g,
	"subcore_vm_fixed3600_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3600A150g,
	"subcore_vm_fixed3620_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3620A150g,
	"subcore_vm_fixed3640_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3640A150g,
	"subcore_vm_fixed3660_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3660A150g,
	"subcore_vm_fixed3680_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3680A150g,
	"subcore_vm_fixed3700_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3700A150g,
	"subcore_vm_fixed3720_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3720A150g,
	"subcore_vm_fixed3740_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3740A150g,
	"subcore_vm_fixed3760_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3760A150g,
	"subcore_vm_fixed3780_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3780A150g,
	"subcore_vm_fixed3800_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3800A150g,
	"subcore_vm_fixed3820_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3820A150g,
	"subcore_vm_fixed3840_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3840A150g,
	"subcore_vm_fixed3860_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3860A150g,
	"subcore_vm_fixed3880_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3880A150g,
	"subcore_vm_fixed3900_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3900A150g,
	"subcore_vm_fixed3920_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3920A150g,
	"subcore_vm_fixed3940_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3940A150g,
	"subcore_vm_fixed3960_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3960A150g,
	"subcore_vm_fixed3980_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3980A150g,
	"subcore_vm_fixed4000_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4000A150g,
	"subcore_vm_fixed4020_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4020A150g,
	"subcore_vm_fixed4040_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4040A150g,
	"subcore_vm_fixed4060_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4060A150g,
	"subcore_vm_fixed4080_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4080A150g,
	"subcore_vm_fixed4100_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4100A150g,
	"subcore_vm_fixed4120_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4120A150g,
	"subcore_vm_fixed4140_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4140A150g,
	"subcore_vm_fixed4160_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4160A150g,
	"subcore_vm_fixed4180_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4180A150g,
	"subcore_vm_fixed4200_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4200A150g,
	"subcore_vm_fixed4220_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4220A150g,
	"subcore_vm_fixed4240_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4240A150g,
	"subcore_vm_fixed4260_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4260A150g,
	"subcore_vm_fixed4280_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4280A150g,
	"subcore_vm_fixed4300_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4300A150g,
	"subcore_vm_fixed4320_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4320A150g,
	"subcore_vm_fixed4340_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4340A150g,
	"subcore_vm_fixed4360_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4360A150g,
	"subcore_vm_fixed4380_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4380A150g,
	"subcore_vm_fixed4400_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4400A150g,
	"subcore_vm_fixed4420_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4420A150g,
	"subcore_vm_fixed4440_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4440A150g,
	"subcore_vm_fixed4460_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4460A150g,
	"subcore_vm_fixed4480_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4480A150g,
	"subcore_vm_fixed4500_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4500A150g,
	"subcore_vm_fixed4520_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4520A150g,
	"subcore_vm_fixed4540_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4540A150g,
	"subcore_vm_fixed4560_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4560A150g,
	"subcore_vm_fixed4580_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4580A150g,
	"subcore_vm_fixed4600_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4600A150g,
	"subcore_vm_fixed4620_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4620A150g,
	"subcore_vm_fixed4640_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4640A150g,
	"subcore_vm_fixed4660_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4660A150g,
	"subcore_vm_fixed4680_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4680A150g,
	"subcore_vm_fixed4700_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4700A150g,
	"subcore_vm_fixed4720_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4720A150g,
	"subcore_vm_fixed4740_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4740A150g,
	"subcore_vm_fixed4760_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4760A150g,
	"subcore_vm_fixed4780_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4780A150g,
	"subcore_vm_fixed4800_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4800A150g,
	"subcore_vm_fixed4820_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4820A150g,
	"subcore_vm_fixed4840_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4840A150g,
	"subcore_vm_fixed4860_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4860A150g,
	"subcore_vm_fixed4880_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4880A150g,
	"subcore_vm_fixed4900_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4900A150g,
	"subcore_vm_fixed4920_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4920A150g,
	"subcore_vm_fixed4940_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4940A150g,
	"subcore_vm_fixed4960_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4960A150g,
	"subcore_vm_fixed4980_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4980A150g,
	"subcore_vm_fixed5000_a1_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed5000A150g,
	"subcore_vm_fixed0090_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0090X950g,
	"subcore_vm_fixed0180_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0180X950g,
	"subcore_vm_fixed0270_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0270X950g,
	"subcore_vm_fixed0360_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0360X950g,
	"subcore_vm_fixed0450_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0450X950g,
	"subcore_vm_fixed0540_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0540X950g,
	"subcore_vm_fixed0630_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0630X950g,
	"subcore_vm_fixed0720_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0720X950g,
	"subcore_vm_fixed0810_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0810X950g,
	"subcore_vm_fixed0900_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0900X950g,
	"subcore_vm_fixed0990_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed0990X950g,
	"subcore_vm_fixed1080_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1080X950g,
	"subcore_vm_fixed1170_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1170X950g,
	"subcore_vm_fixed1260_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1260X950g,
	"subcore_vm_fixed1350_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1350X950g,
	"subcore_vm_fixed1440_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1440X950g,
	"subcore_vm_fixed1530_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1530X950g,
	"subcore_vm_fixed1620_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1620X950g,
	"subcore_vm_fixed1710_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1710X950g,
	"subcore_vm_fixed1800_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1800X950g,
	"subcore_vm_fixed1890_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1890X950g,
	"subcore_vm_fixed1980_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed1980X950g,
	"subcore_vm_fixed2070_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2070X950g,
	"subcore_vm_fixed2160_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2160X950g,
	"subcore_vm_fixed2250_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2250X950g,
	"subcore_vm_fixed2340_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2340X950g,
	"subcore_vm_fixed2430_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2430X950g,
	"subcore_vm_fixed2520_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2520X950g,
	"subcore_vm_fixed2610_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2610X950g,
	"subcore_vm_fixed2700_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2700X950g,
	"subcore_vm_fixed2790_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2790X950g,
	"subcore_vm_fixed2880_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2880X950g,
	"subcore_vm_fixed2970_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed2970X950g,
	"subcore_vm_fixed3060_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3060X950g,
	"subcore_vm_fixed3150_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3150X950g,
	"subcore_vm_fixed3240_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3240X950g,
	"subcore_vm_fixed3330_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3330X950g,
	"subcore_vm_fixed3420_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3420X950g,
	"subcore_vm_fixed3510_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3510X950g,
	"subcore_vm_fixed3600_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3600X950g,
	"subcore_vm_fixed3690_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3690X950g,
	"subcore_vm_fixed3780_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3780X950g,
	"subcore_vm_fixed3870_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3870X950g,
	"subcore_vm_fixed3960_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed3960X950g,
	"subcore_vm_fixed4050_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4050X950g,
	"subcore_vm_fixed4140_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4140X950g,
	"subcore_vm_fixed4230_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4230X950g,
	"subcore_vm_fixed4320_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4320X950g,
	"subcore_vm_fixed4410_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4410X950g,
	"subcore_vm_fixed4500_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4500X950g,
	"subcore_vm_fixed4590_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4590X950g,
	"subcore_vm_fixed4680_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4680X950g,
	"subcore_vm_fixed4770_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4770X950g,
	"subcore_vm_fixed4860_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4860X950g,
	"subcore_vm_fixed4950_x9_50g":          InternalVnicAttachmentVnicShapeSubcoreVmFixed4950X950g,
	"dynamic_a1_50g":                       InternalVnicAttachmentVnicShapeDynamicA150g,
	"fixed0040_a1_50g":                     InternalVnicAttachmentVnicShapeFixed0040A150g,
	"fixed0100_a1_50g":                     InternalVnicAttachmentVnicShapeFixed0100A150g,
	"fixed0200_a1_50g":                     InternalVnicAttachmentVnicShapeFixed0200A150g,
	"fixed0300_a1_50g":                     InternalVnicAttachmentVnicShapeFixed0300A150g,
	"fixed0400_a1_50g":                     InternalVnicAttachmentVnicShapeFixed0400A150g,
	"fixed0500_a1_50g":                     InternalVnicAttachmentVnicShapeFixed0500A150g,
	"fixed0600_a1_50g":                     InternalVnicAttachmentVnicShapeFixed0600A150g,
	"fixed0700_a1_50g":                     InternalVnicAttachmentVnicShapeFixed0700A150g,
	"fixed0800_a1_50g":                     InternalVnicAttachmentVnicShapeFixed0800A150g,
	"fixed0900_a1_50g":                     InternalVnicAttachmentVnicShapeFixed0900A150g,
	"fixed1000_a1_50g":                     InternalVnicAttachmentVnicShapeFixed1000A150g,
	"fixed1100_a1_50g":                     InternalVnicAttachmentVnicShapeFixed1100A150g,
	"fixed1200_a1_50g":                     InternalVnicAttachmentVnicShapeFixed1200A150g,
	"fixed1300_a1_50g":                     InternalVnicAttachmentVnicShapeFixed1300A150g,
	"fixed1400_a1_50g":                     InternalVnicAttachmentVnicShapeFixed1400A150g,
	"fixed1500_a1_50g":                     InternalVnicAttachmentVnicShapeFixed1500A150g,
	"fixed1600_a1_50g":                     InternalVnicAttachmentVnicShapeFixed1600A150g,
	"fixed1700_a1_50g":                     InternalVnicAttachmentVnicShapeFixed1700A150g,
	"fixed1800_a1_50g":                     InternalVnicAttachmentVnicShapeFixed1800A150g,
	"fixed1900_a1_50g":                     InternalVnicAttachmentVnicShapeFixed1900A150g,
	"fixed2000_a1_50g":                     InternalVnicAttachmentVnicShapeFixed2000A150g,
	"fixed2100_a1_50g":                     InternalVnicAttachmentVnicShapeFixed2100A150g,
	"fixed2200_a1_50g":                     InternalVnicAttachmentVnicShapeFixed2200A150g,
	"fixed2300_a1_50g":                     InternalVnicAttachmentVnicShapeFixed2300A150g,
	"fixed2400_a1_50g":                     InternalVnicAttachmentVnicShapeFixed2400A150g,
	"fixed2500_a1_50g":                     InternalVnicAttachmentVnicShapeFixed2500A150g,
	"fixed2600_a1_50g":                     InternalVnicAttachmentVnicShapeFixed2600A150g,
	"fixed2700_a1_50g":                     InternalVnicAttachmentVnicShapeFixed2700A150g,
	"fixed2800_a1_50g":                     InternalVnicAttachmentVnicShapeFixed2800A150g,
	"fixed2900_a1_50g":                     InternalVnicAttachmentVnicShapeFixed2900A150g,
	"fixed3000_a1_50g":                     InternalVnicAttachmentVnicShapeFixed3000A150g,
	"fixed3100_a1_50g":                     InternalVnicAttachmentVnicShapeFixed3100A150g,
	"fixed3200_a1_50g":                     InternalVnicAttachmentVnicShapeFixed3200A150g,
	"fixed3300_a1_50g":                     InternalVnicAttachmentVnicShapeFixed3300A150g,
	"fixed3400_a1_50g":                     InternalVnicAttachmentVnicShapeFixed3400A150g,
	"fixed3500_a1_50g":                     InternalVnicAttachmentVnicShapeFixed3500A150g,
	"fixed3600_a1_50g":                     InternalVnicAttachmentVnicShapeFixed3600A150g,
	"fixed3700_a1_50g":                     InternalVnicAttachmentVnicShapeFixed3700A150g,
	"fixed3800_a1_50g":                     InternalVnicAttachmentVnicShapeFixed3800A150g,
	"fixed3900_a1_50g":                     InternalVnicAttachmentVnicShapeFixed3900A150g,
	"fixed4000_a1_50g":                     InternalVnicAttachmentVnicShapeFixed4000A150g,
	"entirehost_a1_50g":                    InternalVnicAttachmentVnicShapeEntirehostA150g,
	"dynamic_x9_50g":                       InternalVnicAttachmentVnicShapeDynamicX950g,
	"fixed0040_x9_50g":                     InternalVnicAttachmentVnicShapeFixed0040X950g,
	"fixed0400_x9_50g":                     InternalVnicAttachmentVnicShapeFixed0400X950g,
	"fixed0800_x9_50g":                     InternalVnicAttachmentVnicShapeFixed0800X950g,
	"fixed1200_x9_50g":                     InternalVnicAttachmentVnicShapeFixed1200X950g,
	"fixed1600_x9_50g":                     InternalVnicAttachmentVnicShapeFixed1600X950g,
	"fixed2000_x9_50g":                     InternalVnicAttachmentVnicShapeFixed2000X950g,
	"fixed2400_x9_50g":                     InternalVnicAttachmentVnicShapeFixed2400X950g,
	"fixed2800_x9_50g":                     InternalVnicAttachmentVnicShapeFixed2800X950g,
	"fixed3200_x9_50g":                     InternalVnicAttachmentVnicShapeFixed3200X950g,
	"fixed3600_x9_50g":                     InternalVnicAttachmentVnicShapeFixed3600X950g,
	"fixed4000_x9_50g":                     InternalVnicAttachmentVnicShapeFixed4000X950g,
	"standard_vm_fixed0100_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed0100X950g,
	"standard_vm_fixed0200_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed0200X950g,
	"standard_vm_fixed0300_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed0300X950g,
	"standard_vm_fixed0400_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed0400X950g,
	"standard_vm_fixed0500_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed0500X950g,
	"standard_vm_fixed0600_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed0600X950g,
	"standard_vm_fixed0700_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed0700X950g,
	"standard_vm_fixed0800_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed0800X950g,
	"standard_vm_fixed0900_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed0900X950g,
	"standard_vm_fixed1000_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed1000X950g,
	"standard_vm_fixed1100_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed1100X950g,
	"standard_vm_fixed1200_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed1200X950g,
	"standard_vm_fixed1300_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed1300X950g,
	"standard_vm_fixed1400_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed1400X950g,
	"standard_vm_fixed1500_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed1500X950g,
	"standard_vm_fixed1600_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed1600X950g,
	"standard_vm_fixed1700_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed1700X950g,
	"standard_vm_fixed1800_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed1800X950g,
	"standard_vm_fixed1900_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed1900X950g,
	"standard_vm_fixed2000_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed2000X950g,
	"standard_vm_fixed2100_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed2100X950g,
	"standard_vm_fixed2200_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed2200X950g,
	"standard_vm_fixed2300_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed2300X950g,
	"standard_vm_fixed2400_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed2400X950g,
	"standard_vm_fixed2500_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed2500X950g,
	"standard_vm_fixed2600_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed2600X950g,
	"standard_vm_fixed2700_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed2700X950g,
	"standard_vm_fixed2800_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed2800X950g,
	"standard_vm_fixed2900_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed2900X950g,
	"standard_vm_fixed3000_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed3000X950g,
	"standard_vm_fixed3100_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed3100X950g,
	"standard_vm_fixed3200_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed3200X950g,
	"standard_vm_fixed3300_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed3300X950g,
	"standard_vm_fixed3400_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed3400X950g,
	"standard_vm_fixed3500_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed3500X950g,
	"standard_vm_fixed3600_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed3600X950g,
	"standard_vm_fixed3700_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed3700X950g,
	"standard_vm_fixed3800_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed3800X950g,
	"standard_vm_fixed3900_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed3900X950g,
	"standard_vm_fixed4000_x9_50g":         InternalVnicAttachmentVnicShapeStandardVmFixed4000X950g,
	"subcore_standard_vm_fixed0025_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0025X950g,
	"subcore_standard_vm_fixed0050_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0050X950g,
	"subcore_standard_vm_fixed0075_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0075X950g,
	"subcore_standard_vm_fixed0100_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0100X950g,
	"subcore_standard_vm_fixed0125_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0125X950g,
	"subcore_standard_vm_fixed0150_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0150X950g,
	"subcore_standard_vm_fixed0175_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0175X950g,
	"subcore_standard_vm_fixed0200_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0200X950g,
	"subcore_standard_vm_fixed0225_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0225X950g,
	"subcore_standard_vm_fixed0250_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0250X950g,
	"subcore_standard_vm_fixed0275_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0275X950g,
	"subcore_standard_vm_fixed0300_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0300X950g,
	"subcore_standard_vm_fixed0325_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0325X950g,
	"subcore_standard_vm_fixed0350_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0350X950g,
	"subcore_standard_vm_fixed0375_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0375X950g,
	"subcore_standard_vm_fixed0400_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0400X950g,
	"subcore_standard_vm_fixed0425_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0425X950g,
	"subcore_standard_vm_fixed0450_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0450X950g,
	"subcore_standard_vm_fixed0475_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0475X950g,
	"subcore_standard_vm_fixed0500_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0500X950g,
	"subcore_standard_vm_fixed0525_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0525X950g,
	"subcore_standard_vm_fixed0550_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0550X950g,
	"subcore_standard_vm_fixed0575_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0575X950g,
	"subcore_standard_vm_fixed0600_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0600X950g,
	"subcore_standard_vm_fixed0625_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0625X950g,
	"subcore_standard_vm_fixed0650_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0650X950g,
	"subcore_standard_vm_fixed0675_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0675X950g,
	"subcore_standard_vm_fixed0700_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0700X950g,
	"subcore_standard_vm_fixed0725_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0725X950g,
	"subcore_standard_vm_fixed0750_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0750X950g,
	"subcore_standard_vm_fixed0775_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0775X950g,
	"subcore_standard_vm_fixed0800_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0800X950g,
	"subcore_standard_vm_fixed0825_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0825X950g,
	"subcore_standard_vm_fixed0850_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0850X950g,
	"subcore_standard_vm_fixed0875_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0875X950g,
	"subcore_standard_vm_fixed0900_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0900X950g,
	"subcore_standard_vm_fixed0925_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0925X950g,
	"subcore_standard_vm_fixed0950_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0950X950g,
	"subcore_standard_vm_fixed0975_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed0975X950g,
	"subcore_standard_vm_fixed1000_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1000X950g,
	"subcore_standard_vm_fixed1025_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1025X950g,
	"subcore_standard_vm_fixed1050_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1050X950g,
	"subcore_standard_vm_fixed1075_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1075X950g,
	"subcore_standard_vm_fixed1100_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1100X950g,
	"subcore_standard_vm_fixed1125_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1125X950g,
	"subcore_standard_vm_fixed1150_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1150X950g,
	"subcore_standard_vm_fixed1175_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1175X950g,
	"subcore_standard_vm_fixed1200_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1200X950g,
	"subcore_standard_vm_fixed1225_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1225X950g,
	"subcore_standard_vm_fixed1250_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1250X950g,
	"subcore_standard_vm_fixed1275_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1275X950g,
	"subcore_standard_vm_fixed1300_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1300X950g,
	"subcore_standard_vm_fixed1325_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1325X950g,
	"subcore_standard_vm_fixed1350_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1350X950g,
	"subcore_standard_vm_fixed1375_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1375X950g,
	"subcore_standard_vm_fixed1400_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1400X950g,
	"subcore_standard_vm_fixed1425_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1425X950g,
	"subcore_standard_vm_fixed1450_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1450X950g,
	"subcore_standard_vm_fixed1475_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1475X950g,
	"subcore_standard_vm_fixed1500_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1500X950g,
	"subcore_standard_vm_fixed1525_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1525X950g,
	"subcore_standard_vm_fixed1550_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1550X950g,
	"subcore_standard_vm_fixed1575_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1575X950g,
	"subcore_standard_vm_fixed1600_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1600X950g,
	"subcore_standard_vm_fixed1625_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1625X950g,
	"subcore_standard_vm_fixed1650_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1650X950g,
	"subcore_standard_vm_fixed1700_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1700X950g,
	"subcore_standard_vm_fixed1725_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1725X950g,
	"subcore_standard_vm_fixed1750_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1750X950g,
	"subcore_standard_vm_fixed1800_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1800X950g,
	"subcore_standard_vm_fixed1850_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1850X950g,
	"subcore_standard_vm_fixed1875_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1875X950g,
	"subcore_standard_vm_fixed1900_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1900X950g,
	"subcore_standard_vm_fixed1925_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1925X950g,
	"subcore_standard_vm_fixed1950_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed1950X950g,
	"subcore_standard_vm_fixed2000_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2000X950g,
	"subcore_standard_vm_fixed2025_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2025X950g,
	"subcore_standard_vm_fixed2050_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2050X950g,
	"subcore_standard_vm_fixed2100_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2100X950g,
	"subcore_standard_vm_fixed2125_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2125X950g,
	"subcore_standard_vm_fixed2150_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2150X950g,
	"subcore_standard_vm_fixed2175_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2175X950g,
	"subcore_standard_vm_fixed2200_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2200X950g,
	"subcore_standard_vm_fixed2250_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2250X950g,
	"subcore_standard_vm_fixed2275_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2275X950g,
	"subcore_standard_vm_fixed2300_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2300X950g,
	"subcore_standard_vm_fixed2325_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2325X950g,
	"subcore_standard_vm_fixed2350_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2350X950g,
	"subcore_standard_vm_fixed2375_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2375X950g,
	"subcore_standard_vm_fixed2400_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2400X950g,
	"subcore_standard_vm_fixed2450_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2450X950g,
	"subcore_standard_vm_fixed2475_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2475X950g,
	"subcore_standard_vm_fixed2500_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2500X950g,
	"subcore_standard_vm_fixed2550_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2550X950g,
	"subcore_standard_vm_fixed2600_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2600X950g,
	"subcore_standard_vm_fixed2625_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2625X950g,
	"subcore_standard_vm_fixed2650_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2650X950g,
	"subcore_standard_vm_fixed2700_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2700X950g,
	"subcore_standard_vm_fixed2750_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2750X950g,
	"subcore_standard_vm_fixed2775_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2775X950g,
	"subcore_standard_vm_fixed2800_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2800X950g,
	"subcore_standard_vm_fixed2850_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2850X950g,
	"subcore_standard_vm_fixed2875_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2875X950g,
	"subcore_standard_vm_fixed2900_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2900X950g,
	"subcore_standard_vm_fixed2925_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2925X950g,
	"subcore_standard_vm_fixed2950_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2950X950g,
	"subcore_standard_vm_fixed2975_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed2975X950g,
	"subcore_standard_vm_fixed3000_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3000X950g,
	"subcore_standard_vm_fixed3025_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3025X950g,
	"subcore_standard_vm_fixed3050_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3050X950g,
	"subcore_standard_vm_fixed3075_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3075X950g,
	"subcore_standard_vm_fixed3100_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3100X950g,
	"subcore_standard_vm_fixed3125_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3125X950g,
	"subcore_standard_vm_fixed3150_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3150X950g,
	"subcore_standard_vm_fixed3200_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3200X950g,
	"subcore_standard_vm_fixed3225_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3225X950g,
	"subcore_standard_vm_fixed3250_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3250X950g,
	"subcore_standard_vm_fixed3300_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3300X950g,
	"subcore_standard_vm_fixed3325_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3325X950g,
	"subcore_standard_vm_fixed3375_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3375X950g,
	"subcore_standard_vm_fixed3400_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3400X950g,
	"subcore_standard_vm_fixed3450_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3450X950g,
	"subcore_standard_vm_fixed3500_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3500X950g,
	"subcore_standard_vm_fixed3525_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3525X950g,
	"subcore_standard_vm_fixed3575_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3575X950g,
	"subcore_standard_vm_fixed3600_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3600X950g,
	"subcore_standard_vm_fixed3625_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3625X950g,
	"subcore_standard_vm_fixed3675_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3675X950g,
	"subcore_standard_vm_fixed3700_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3700X950g,
	"subcore_standard_vm_fixed3750_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3750X950g,
	"subcore_standard_vm_fixed3800_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3800X950g,
	"subcore_standard_vm_fixed3825_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3825X950g,
	"subcore_standard_vm_fixed3850_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3850X950g,
	"subcore_standard_vm_fixed3875_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3875X950g,
	"subcore_standard_vm_fixed3900_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3900X950g,
	"subcore_standard_vm_fixed3975_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed3975X950g,
	"subcore_standard_vm_fixed4000_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4000X950g,
	"subcore_standard_vm_fixed4025_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4025X950g,
	"subcore_standard_vm_fixed4050_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4050X950g,
	"subcore_standard_vm_fixed4100_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4100X950g,
	"subcore_standard_vm_fixed4125_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4125X950g,
	"subcore_standard_vm_fixed4200_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4200X950g,
	"subcore_standard_vm_fixed4225_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4225X950g,
	"subcore_standard_vm_fixed4250_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4250X950g,
	"subcore_standard_vm_fixed4275_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4275X950g,
	"subcore_standard_vm_fixed4300_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4300X950g,
	"subcore_standard_vm_fixed4350_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4350X950g,
	"subcore_standard_vm_fixed4375_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4375X950g,
	"subcore_standard_vm_fixed4400_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4400X950g,
	"subcore_standard_vm_fixed4425_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4425X950g,
	"subcore_standard_vm_fixed4500_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4500X950g,
	"subcore_standard_vm_fixed4550_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4550X950g,
	"subcore_standard_vm_fixed4575_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4575X950g,
	"subcore_standard_vm_fixed4600_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4600X950g,
	"subcore_standard_vm_fixed4625_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4625X950g,
	"subcore_standard_vm_fixed4650_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4650X950g,
	"subcore_standard_vm_fixed4675_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4675X950g,
	"subcore_standard_vm_fixed4700_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4700X950g,
	"subcore_standard_vm_fixed4725_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4725X950g,
	"subcore_standard_vm_fixed4750_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4750X950g,
	"subcore_standard_vm_fixed4800_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4800X950g,
	"subcore_standard_vm_fixed4875_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4875X950g,
	"subcore_standard_vm_fixed4900_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4900X950g,
	"subcore_standard_vm_fixed4950_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed4950X950g,
	"subcore_standard_vm_fixed5000_x9_50g": InternalVnicAttachmentVnicShapeSubcoreStandardVmFixed5000X950g,
	"entirehost_x9_50g":                    InternalVnicAttachmentVnicShapeEntirehostX950g,
}

// GetInternalVnicAttachmentVnicShapeEnumValues Enumerates the set of values for InternalVnicAttachmentVnicShapeEnum
func GetInternalVnicAttachmentVnicShapeEnumValues() []InternalVnicAttachmentVnicShapeEnum {
	values := make([]InternalVnicAttachmentVnicShapeEnum, 0)
	for _, v := range mappingInternalVnicAttachmentVnicShapeEnum {
		values = append(values, v)
	}
	return values
}

// GetInternalVnicAttachmentVnicShapeEnumStringValues Enumerates the set of values in String for InternalVnicAttachmentVnicShapeEnum
func GetInternalVnicAttachmentVnicShapeEnumStringValues() []string {
	return []string{
		"DYNAMIC",
		"FIXED0040",
		"FIXED0060",
		"FIXED0060_PSM",
		"FIXED0100",
		"FIXED0120",
		"FIXED0120_2X",
		"FIXED0200",
		"FIXED0240",
		"FIXED0480",
		"ENTIREHOST",
		"DYNAMIC_25G",
		"FIXED0040_25G",
		"FIXED0100_25G",
		"FIXED0200_25G",
		"FIXED0400_25G",
		"FIXED0800_25G",
		"FIXED1600_25G",
		"FIXED2400_25G",
		"ENTIREHOST_25G",
		"DYNAMIC_E1_25G",
		"FIXED0040_E1_25G",
		"FIXED0070_E1_25G",
		"FIXED0140_E1_25G",
		"FIXED0280_E1_25G",
		"FIXED0560_E1_25G",
		"FIXED1120_E1_25G",
		"FIXED1680_E1_25G",
		"ENTIREHOST_E1_25G",
		"DYNAMIC_B1_25G",
		"FIXED0040_B1_25G",
		"FIXED0060_B1_25G",
		"FIXED0120_B1_25G",
		"FIXED0240_B1_25G",
		"FIXED0480_B1_25G",
		"FIXED0960_B1_25G",
		"ENTIREHOST_B1_25G",
		"MICRO_VM_FIXED0048_E1_25G",
		"MICRO_LB_FIXED0001_E1_25G",
		"VNICAAS_FIXED0200",
		"VNICAAS_FIXED0400",
		"VNICAAS_FIXED0700",
		"VNICAAS_NLB_APPROVED_10G",
		"VNICAAS_NLB_APPROVED_25G",
		"VNICAAS_TELESIS_25G",
		"VNICAAS_TELESIS_10G",
		"VNICAAS_AMBASSADOR_FIXED0100",
		"VNICAAS_TELESIS_GAMMA",
		"VNICAAS_PRIVATEDNS",
		"VNICAAS_FWAAS",
		"VNICAAS_LBAAS_FREE",
		"VNICAAS_LBAAS_8G_512K",
		"VNICAAS_LBAAS_8G_1M",
		"VNICAAS_LBAAS_8G_2M",
		"VNICAAS_LBAAS_8G_3M",
		"VNICAAS_LBAAS_8G_1M_8GHOST",
		"VNICAAS_LBAAS_8G_1M_16GHOST",
		"DYNAMIC_E3_50G",
		"FIXED0040_E3_50G",
		"FIXED0100_E3_50G",
		"FIXED0200_E3_50G",
		"FIXED0300_E3_50G",
		"FIXED0400_E3_50G",
		"FIXED0500_E3_50G",
		"FIXED0600_E3_50G",
		"FIXED0700_E3_50G",
		"FIXED0800_E3_50G",
		"FIXED0900_E3_50G",
		"FIXED1000_E3_50G",
		"FIXED1100_E3_50G",
		"FIXED1200_E3_50G",
		"FIXED1300_E3_50G",
		"FIXED1400_E3_50G",
		"FIXED1500_E3_50G",
		"FIXED1600_E3_50G",
		"FIXED1700_E3_50G",
		"FIXED1800_E3_50G",
		"FIXED1900_E3_50G",
		"FIXED2000_E3_50G",
		"FIXED2100_E3_50G",
		"FIXED2200_E3_50G",
		"FIXED2300_E3_50G",
		"FIXED2400_E3_50G",
		"FIXED2500_E3_50G",
		"FIXED2600_E3_50G",
		"FIXED2700_E3_50G",
		"FIXED2800_E3_50G",
		"FIXED2900_E3_50G",
		"FIXED3000_E3_50G",
		"FIXED3100_E3_50G",
		"FIXED3200_E3_50G",
		"FIXED3300_E3_50G",
		"FIXED3400_E3_50G",
		"FIXED3500_E3_50G",
		"FIXED3600_E3_50G",
		"FIXED3700_E3_50G",
		"FIXED3800_E3_50G",
		"FIXED3900_E3_50G",
		"FIXED4000_E3_50G",
		"ENTIREHOST_E3_50G",
		"DYNAMIC_E4_50G",
		"FIXED0040_E4_50G",
		"FIXED0100_E4_50G",
		"FIXED0200_E4_50G",
		"FIXED0300_E4_50G",
		"FIXED0400_E4_50G",
		"FIXED0500_E4_50G",
		"FIXED0600_E4_50G",
		"FIXED0700_E4_50G",
		"FIXED0800_E4_50G",
		"FIXED0900_E4_50G",
		"FIXED1000_E4_50G",
		"FIXED1100_E4_50G",
		"FIXED1200_E4_50G",
		"FIXED1300_E4_50G",
		"FIXED1400_E4_50G",
		"FIXED1500_E4_50G",
		"FIXED1600_E4_50G",
		"FIXED1700_E4_50G",
		"FIXED1800_E4_50G",
		"FIXED1900_E4_50G",
		"FIXED2000_E4_50G",
		"FIXED2100_E4_50G",
		"FIXED2200_E4_50G",
		"FIXED2300_E4_50G",
		"FIXED2400_E4_50G",
		"FIXED2500_E4_50G",
		"FIXED2600_E4_50G",
		"FIXED2700_E4_50G",
		"FIXED2800_E4_50G",
		"FIXED2900_E4_50G",
		"FIXED3000_E4_50G",
		"FIXED3100_E4_50G",
		"FIXED3200_E4_50G",
		"FIXED3300_E4_50G",
		"FIXED3400_E4_50G",
		"FIXED3500_E4_50G",
		"FIXED3600_E4_50G",
		"FIXED3700_E4_50G",
		"FIXED3800_E4_50G",
		"FIXED3900_E4_50G",
		"FIXED4000_E4_50G",
		"ENTIREHOST_E4_50G",
		"Micro_VM_Fixed0050_E3_50G",
		"Micro_VM_Fixed0050_E4_50G",
		"SUBCORE_VM_FIXED0025_E3_50G",
		"SUBCORE_VM_FIXED0050_E3_50G",
		"SUBCORE_VM_FIXED0075_E3_50G",
		"SUBCORE_VM_FIXED0100_E3_50G",
		"SUBCORE_VM_FIXED0125_E3_50G",
		"SUBCORE_VM_FIXED0150_E3_50G",
		"SUBCORE_VM_FIXED0175_E3_50G",
		"SUBCORE_VM_FIXED0200_E3_50G",
		"SUBCORE_VM_FIXED0225_E3_50G",
		"SUBCORE_VM_FIXED0250_E3_50G",
		"SUBCORE_VM_FIXED0275_E3_50G",
		"SUBCORE_VM_FIXED0300_E3_50G",
		"SUBCORE_VM_FIXED0325_E3_50G",
		"SUBCORE_VM_FIXED0350_E3_50G",
		"SUBCORE_VM_FIXED0375_E3_50G",
		"SUBCORE_VM_FIXED0400_E3_50G",
		"SUBCORE_VM_FIXED0425_E3_50G",
		"SUBCORE_VM_FIXED0450_E3_50G",
		"SUBCORE_VM_FIXED0475_E3_50G",
		"SUBCORE_VM_FIXED0500_E3_50G",
		"SUBCORE_VM_FIXED0525_E3_50G",
		"SUBCORE_VM_FIXED0550_E3_50G",
		"SUBCORE_VM_FIXED0575_E3_50G",
		"SUBCORE_VM_FIXED0600_E3_50G",
		"SUBCORE_VM_FIXED0625_E3_50G",
		"SUBCORE_VM_FIXED0650_E3_50G",
		"SUBCORE_VM_FIXED0675_E3_50G",
		"SUBCORE_VM_FIXED0700_E3_50G",
		"SUBCORE_VM_FIXED0725_E3_50G",
		"SUBCORE_VM_FIXED0750_E3_50G",
		"SUBCORE_VM_FIXED0775_E3_50G",
		"SUBCORE_VM_FIXED0800_E3_50G",
		"SUBCORE_VM_FIXED0825_E3_50G",
		"SUBCORE_VM_FIXED0850_E3_50G",
		"SUBCORE_VM_FIXED0875_E3_50G",
		"SUBCORE_VM_FIXED0900_E3_50G",
		"SUBCORE_VM_FIXED0925_E3_50G",
		"SUBCORE_VM_FIXED0950_E3_50G",
		"SUBCORE_VM_FIXED0975_E3_50G",
		"SUBCORE_VM_FIXED1000_E3_50G",
		"SUBCORE_VM_FIXED1025_E3_50G",
		"SUBCORE_VM_FIXED1050_E3_50G",
		"SUBCORE_VM_FIXED1075_E3_50G",
		"SUBCORE_VM_FIXED1100_E3_50G",
		"SUBCORE_VM_FIXED1125_E3_50G",
		"SUBCORE_VM_FIXED1150_E3_50G",
		"SUBCORE_VM_FIXED1175_E3_50G",
		"SUBCORE_VM_FIXED1200_E3_50G",
		"SUBCORE_VM_FIXED1225_E3_50G",
		"SUBCORE_VM_FIXED1250_E3_50G",
		"SUBCORE_VM_FIXED1275_E3_50G",
		"SUBCORE_VM_FIXED1300_E3_50G",
		"SUBCORE_VM_FIXED1325_E3_50G",
		"SUBCORE_VM_FIXED1350_E3_50G",
		"SUBCORE_VM_FIXED1375_E3_50G",
		"SUBCORE_VM_FIXED1400_E3_50G",
		"SUBCORE_VM_FIXED1425_E3_50G",
		"SUBCORE_VM_FIXED1450_E3_50G",
		"SUBCORE_VM_FIXED1475_E3_50G",
		"SUBCORE_VM_FIXED1500_E3_50G",
		"SUBCORE_VM_FIXED1525_E3_50G",
		"SUBCORE_VM_FIXED1550_E3_50G",
		"SUBCORE_VM_FIXED1575_E3_50G",
		"SUBCORE_VM_FIXED1600_E3_50G",
		"SUBCORE_VM_FIXED1625_E3_50G",
		"SUBCORE_VM_FIXED1650_E3_50G",
		"SUBCORE_VM_FIXED1700_E3_50G",
		"SUBCORE_VM_FIXED1725_E3_50G",
		"SUBCORE_VM_FIXED1750_E3_50G",
		"SUBCORE_VM_FIXED1800_E3_50G",
		"SUBCORE_VM_FIXED1850_E3_50G",
		"SUBCORE_VM_FIXED1875_E3_50G",
		"SUBCORE_VM_FIXED1900_E3_50G",
		"SUBCORE_VM_FIXED1925_E3_50G",
		"SUBCORE_VM_FIXED1950_E3_50G",
		"SUBCORE_VM_FIXED2000_E3_50G",
		"SUBCORE_VM_FIXED2025_E3_50G",
		"SUBCORE_VM_FIXED2050_E3_50G",
		"SUBCORE_VM_FIXED2100_E3_50G",
		"SUBCORE_VM_FIXED2125_E3_50G",
		"SUBCORE_VM_FIXED2150_E3_50G",
		"SUBCORE_VM_FIXED2175_E3_50G",
		"SUBCORE_VM_FIXED2200_E3_50G",
		"SUBCORE_VM_FIXED2250_E3_50G",
		"SUBCORE_VM_FIXED2275_E3_50G",
		"SUBCORE_VM_FIXED2300_E3_50G",
		"SUBCORE_VM_FIXED2325_E3_50G",
		"SUBCORE_VM_FIXED2350_E3_50G",
		"SUBCORE_VM_FIXED2375_E3_50G",
		"SUBCORE_VM_FIXED2400_E3_50G",
		"SUBCORE_VM_FIXED2450_E3_50G",
		"SUBCORE_VM_FIXED2475_E3_50G",
		"SUBCORE_VM_FIXED2500_E3_50G",
		"SUBCORE_VM_FIXED2550_E3_50G",
		"SUBCORE_VM_FIXED2600_E3_50G",
		"SUBCORE_VM_FIXED2625_E3_50G",
		"SUBCORE_VM_FIXED2650_E3_50G",
		"SUBCORE_VM_FIXED2700_E3_50G",
		"SUBCORE_VM_FIXED2750_E3_50G",
		"SUBCORE_VM_FIXED2775_E3_50G",
		"SUBCORE_VM_FIXED2800_E3_50G",
		"SUBCORE_VM_FIXED2850_E3_50G",
		"SUBCORE_VM_FIXED2875_E3_50G",
		"SUBCORE_VM_FIXED2900_E3_50G",
		"SUBCORE_VM_FIXED2925_E3_50G",
		"SUBCORE_VM_FIXED2950_E3_50G",
		"SUBCORE_VM_FIXED2975_E3_50G",
		"SUBCORE_VM_FIXED3000_E3_50G",
		"SUBCORE_VM_FIXED3025_E3_50G",
		"SUBCORE_VM_FIXED3050_E3_50G",
		"SUBCORE_VM_FIXED3075_E3_50G",
		"SUBCORE_VM_FIXED3100_E3_50G",
		"SUBCORE_VM_FIXED3125_E3_50G",
		"SUBCORE_VM_FIXED3150_E3_50G",
		"SUBCORE_VM_FIXED3200_E3_50G",
		"SUBCORE_VM_FIXED3225_E3_50G",
		"SUBCORE_VM_FIXED3250_E3_50G",
		"SUBCORE_VM_FIXED3300_E3_50G",
		"SUBCORE_VM_FIXED3325_E3_50G",
		"SUBCORE_VM_FIXED3375_E3_50G",
		"SUBCORE_VM_FIXED3400_E3_50G",
		"SUBCORE_VM_FIXED3450_E3_50G",
		"SUBCORE_VM_FIXED3500_E3_50G",
		"SUBCORE_VM_FIXED3525_E3_50G",
		"SUBCORE_VM_FIXED3575_E3_50G",
		"SUBCORE_VM_FIXED3600_E3_50G",
		"SUBCORE_VM_FIXED3625_E3_50G",
		"SUBCORE_VM_FIXED3675_E3_50G",
		"SUBCORE_VM_FIXED3700_E3_50G",
		"SUBCORE_VM_FIXED3750_E3_50G",
		"SUBCORE_VM_FIXED3800_E3_50G",
		"SUBCORE_VM_FIXED3825_E3_50G",
		"SUBCORE_VM_FIXED3850_E3_50G",
		"SUBCORE_VM_FIXED3875_E3_50G",
		"SUBCORE_VM_FIXED3900_E3_50G",
		"SUBCORE_VM_FIXED3975_E3_50G",
		"SUBCORE_VM_FIXED4000_E3_50G",
		"SUBCORE_VM_FIXED4025_E3_50G",
		"SUBCORE_VM_FIXED4050_E3_50G",
		"SUBCORE_VM_FIXED4100_E3_50G",
		"SUBCORE_VM_FIXED4125_E3_50G",
		"SUBCORE_VM_FIXED4200_E3_50G",
		"SUBCORE_VM_FIXED4225_E3_50G",
		"SUBCORE_VM_FIXED4250_E3_50G",
		"SUBCORE_VM_FIXED4275_E3_50G",
		"SUBCORE_VM_FIXED4300_E3_50G",
		"SUBCORE_VM_FIXED4350_E3_50G",
		"SUBCORE_VM_FIXED4375_E3_50G",
		"SUBCORE_VM_FIXED4400_E3_50G",
		"SUBCORE_VM_FIXED4425_E3_50G",
		"SUBCORE_VM_FIXED4500_E3_50G",
		"SUBCORE_VM_FIXED4550_E3_50G",
		"SUBCORE_VM_FIXED4575_E3_50G",
		"SUBCORE_VM_FIXED4600_E3_50G",
		"SUBCORE_VM_FIXED4625_E3_50G",
		"SUBCORE_VM_FIXED4650_E3_50G",
		"SUBCORE_VM_FIXED4675_E3_50G",
		"SUBCORE_VM_FIXED4700_E3_50G",
		"SUBCORE_VM_FIXED4725_E3_50G",
		"SUBCORE_VM_FIXED4750_E3_50G",
		"SUBCORE_VM_FIXED4800_E3_50G",
		"SUBCORE_VM_FIXED4875_E3_50G",
		"SUBCORE_VM_FIXED4900_E3_50G",
		"SUBCORE_VM_FIXED4950_E3_50G",
		"SUBCORE_VM_FIXED5000_E3_50G",
		"SUBCORE_VM_FIXED0025_E4_50G",
		"SUBCORE_VM_FIXED0050_E4_50G",
		"SUBCORE_VM_FIXED0075_E4_50G",
		"SUBCORE_VM_FIXED0100_E4_50G",
		"SUBCORE_VM_FIXED0125_E4_50G",
		"SUBCORE_VM_FIXED0150_E4_50G",
		"SUBCORE_VM_FIXED0175_E4_50G",
		"SUBCORE_VM_FIXED0200_E4_50G",
		"SUBCORE_VM_FIXED0225_E4_50G",
		"SUBCORE_VM_FIXED0250_E4_50G",
		"SUBCORE_VM_FIXED0275_E4_50G",
		"SUBCORE_VM_FIXED0300_E4_50G",
		"SUBCORE_VM_FIXED0325_E4_50G",
		"SUBCORE_VM_FIXED0350_E4_50G",
		"SUBCORE_VM_FIXED0375_E4_50G",
		"SUBCORE_VM_FIXED0400_E4_50G",
		"SUBCORE_VM_FIXED0425_E4_50G",
		"SUBCORE_VM_FIXED0450_E4_50G",
		"SUBCORE_VM_FIXED0475_E4_50G",
		"SUBCORE_VM_FIXED0500_E4_50G",
		"SUBCORE_VM_FIXED0525_E4_50G",
		"SUBCORE_VM_FIXED0550_E4_50G",
		"SUBCORE_VM_FIXED0575_E4_50G",
		"SUBCORE_VM_FIXED0600_E4_50G",
		"SUBCORE_VM_FIXED0625_E4_50G",
		"SUBCORE_VM_FIXED0650_E4_50G",
		"SUBCORE_VM_FIXED0675_E4_50G",
		"SUBCORE_VM_FIXED0700_E4_50G",
		"SUBCORE_VM_FIXED0725_E4_50G",
		"SUBCORE_VM_FIXED0750_E4_50G",
		"SUBCORE_VM_FIXED0775_E4_50G",
		"SUBCORE_VM_FIXED0800_E4_50G",
		"SUBCORE_VM_FIXED0825_E4_50G",
		"SUBCORE_VM_FIXED0850_E4_50G",
		"SUBCORE_VM_FIXED0875_E4_50G",
		"SUBCORE_VM_FIXED0900_E4_50G",
		"SUBCORE_VM_FIXED0925_E4_50G",
		"SUBCORE_VM_FIXED0950_E4_50G",
		"SUBCORE_VM_FIXED0975_E4_50G",
		"SUBCORE_VM_FIXED1000_E4_50G",
		"SUBCORE_VM_FIXED1025_E4_50G",
		"SUBCORE_VM_FIXED1050_E4_50G",
		"SUBCORE_VM_FIXED1075_E4_50G",
		"SUBCORE_VM_FIXED1100_E4_50G",
		"SUBCORE_VM_FIXED1125_E4_50G",
		"SUBCORE_VM_FIXED1150_E4_50G",
		"SUBCORE_VM_FIXED1175_E4_50G",
		"SUBCORE_VM_FIXED1200_E4_50G",
		"SUBCORE_VM_FIXED1225_E4_50G",
		"SUBCORE_VM_FIXED1250_E4_50G",
		"SUBCORE_VM_FIXED1275_E4_50G",
		"SUBCORE_VM_FIXED1300_E4_50G",
		"SUBCORE_VM_FIXED1325_E4_50G",
		"SUBCORE_VM_FIXED1350_E4_50G",
		"SUBCORE_VM_FIXED1375_E4_50G",
		"SUBCORE_VM_FIXED1400_E4_50G",
		"SUBCORE_VM_FIXED1425_E4_50G",
		"SUBCORE_VM_FIXED1450_E4_50G",
		"SUBCORE_VM_FIXED1475_E4_50G",
		"SUBCORE_VM_FIXED1500_E4_50G",
		"SUBCORE_VM_FIXED1525_E4_50G",
		"SUBCORE_VM_FIXED1550_E4_50G",
		"SUBCORE_VM_FIXED1575_E4_50G",
		"SUBCORE_VM_FIXED1600_E4_50G",
		"SUBCORE_VM_FIXED1625_E4_50G",
		"SUBCORE_VM_FIXED1650_E4_50G",
		"SUBCORE_VM_FIXED1700_E4_50G",
		"SUBCORE_VM_FIXED1725_E4_50G",
		"SUBCORE_VM_FIXED1750_E4_50G",
		"SUBCORE_VM_FIXED1800_E4_50G",
		"SUBCORE_VM_FIXED1850_E4_50G",
		"SUBCORE_VM_FIXED1875_E4_50G",
		"SUBCORE_VM_FIXED1900_E4_50G",
		"SUBCORE_VM_FIXED1925_E4_50G",
		"SUBCORE_VM_FIXED1950_E4_50G",
		"SUBCORE_VM_FIXED2000_E4_50G",
		"SUBCORE_VM_FIXED2025_E4_50G",
		"SUBCORE_VM_FIXED2050_E4_50G",
		"SUBCORE_VM_FIXED2100_E4_50G",
		"SUBCORE_VM_FIXED2125_E4_50G",
		"SUBCORE_VM_FIXED2150_E4_50G",
		"SUBCORE_VM_FIXED2175_E4_50G",
		"SUBCORE_VM_FIXED2200_E4_50G",
		"SUBCORE_VM_FIXED2250_E4_50G",
		"SUBCORE_VM_FIXED2275_E4_50G",
		"SUBCORE_VM_FIXED2300_E4_50G",
		"SUBCORE_VM_FIXED2325_E4_50G",
		"SUBCORE_VM_FIXED2350_E4_50G",
		"SUBCORE_VM_FIXED2375_E4_50G",
		"SUBCORE_VM_FIXED2400_E4_50G",
		"SUBCORE_VM_FIXED2450_E4_50G",
		"SUBCORE_VM_FIXED2475_E4_50G",
		"SUBCORE_VM_FIXED2500_E4_50G",
		"SUBCORE_VM_FIXED2550_E4_50G",
		"SUBCORE_VM_FIXED2600_E4_50G",
		"SUBCORE_VM_FIXED2625_E4_50G",
		"SUBCORE_VM_FIXED2650_E4_50G",
		"SUBCORE_VM_FIXED2700_E4_50G",
		"SUBCORE_VM_FIXED2750_E4_50G",
		"SUBCORE_VM_FIXED2775_E4_50G",
		"SUBCORE_VM_FIXED2800_E4_50G",
		"SUBCORE_VM_FIXED2850_E4_50G",
		"SUBCORE_VM_FIXED2875_E4_50G",
		"SUBCORE_VM_FIXED2900_E4_50G",
		"SUBCORE_VM_FIXED2925_E4_50G",
		"SUBCORE_VM_FIXED2950_E4_50G",
		"SUBCORE_VM_FIXED2975_E4_50G",
		"SUBCORE_VM_FIXED3000_E4_50G",
		"SUBCORE_VM_FIXED3025_E4_50G",
		"SUBCORE_VM_FIXED3050_E4_50G",
		"SUBCORE_VM_FIXED3075_E4_50G",
		"SUBCORE_VM_FIXED3100_E4_50G",
		"SUBCORE_VM_FIXED3125_E4_50G",
		"SUBCORE_VM_FIXED3150_E4_50G",
		"SUBCORE_VM_FIXED3200_E4_50G",
		"SUBCORE_VM_FIXED3225_E4_50G",
		"SUBCORE_VM_FIXED3250_E4_50G",
		"SUBCORE_VM_FIXED3300_E4_50G",
		"SUBCORE_VM_FIXED3325_E4_50G",
		"SUBCORE_VM_FIXED3375_E4_50G",
		"SUBCORE_VM_FIXED3400_E4_50G",
		"SUBCORE_VM_FIXED3450_E4_50G",
		"SUBCORE_VM_FIXED3500_E4_50G",
		"SUBCORE_VM_FIXED3525_E4_50G",
		"SUBCORE_VM_FIXED3575_E4_50G",
		"SUBCORE_VM_FIXED3600_E4_50G",
		"SUBCORE_VM_FIXED3625_E4_50G",
		"SUBCORE_VM_FIXED3675_E4_50G",
		"SUBCORE_VM_FIXED3700_E4_50G",
		"SUBCORE_VM_FIXED3750_E4_50G",
		"SUBCORE_VM_FIXED3800_E4_50G",
		"SUBCORE_VM_FIXED3825_E4_50G",
		"SUBCORE_VM_FIXED3850_E4_50G",
		"SUBCORE_VM_FIXED3875_E4_50G",
		"SUBCORE_VM_FIXED3900_E4_50G",
		"SUBCORE_VM_FIXED3975_E4_50G",
		"SUBCORE_VM_FIXED4000_E4_50G",
		"SUBCORE_VM_FIXED4025_E4_50G",
		"SUBCORE_VM_FIXED4050_E4_50G",
		"SUBCORE_VM_FIXED4100_E4_50G",
		"SUBCORE_VM_FIXED4125_E4_50G",
		"SUBCORE_VM_FIXED4200_E4_50G",
		"SUBCORE_VM_FIXED4225_E4_50G",
		"SUBCORE_VM_FIXED4250_E4_50G",
		"SUBCORE_VM_FIXED4275_E4_50G",
		"SUBCORE_VM_FIXED4300_E4_50G",
		"SUBCORE_VM_FIXED4350_E4_50G",
		"SUBCORE_VM_FIXED4375_E4_50G",
		"SUBCORE_VM_FIXED4400_E4_50G",
		"SUBCORE_VM_FIXED4425_E4_50G",
		"SUBCORE_VM_FIXED4500_E4_50G",
		"SUBCORE_VM_FIXED4550_E4_50G",
		"SUBCORE_VM_FIXED4575_E4_50G",
		"SUBCORE_VM_FIXED4600_E4_50G",
		"SUBCORE_VM_FIXED4625_E4_50G",
		"SUBCORE_VM_FIXED4650_E4_50G",
		"SUBCORE_VM_FIXED4675_E4_50G",
		"SUBCORE_VM_FIXED4700_E4_50G",
		"SUBCORE_VM_FIXED4725_E4_50G",
		"SUBCORE_VM_FIXED4750_E4_50G",
		"SUBCORE_VM_FIXED4800_E4_50G",
		"SUBCORE_VM_FIXED4875_E4_50G",
		"SUBCORE_VM_FIXED4900_E4_50G",
		"SUBCORE_VM_FIXED4950_E4_50G",
		"SUBCORE_VM_FIXED5000_E4_50G",
		"SUBCORE_VM_FIXED0020_A1_50G",
		"SUBCORE_VM_FIXED0040_A1_50G",
		"SUBCORE_VM_FIXED0060_A1_50G",
		"SUBCORE_VM_FIXED0080_A1_50G",
		"SUBCORE_VM_FIXED0100_A1_50G",
		"SUBCORE_VM_FIXED0120_A1_50G",
		"SUBCORE_VM_FIXED0140_A1_50G",
		"SUBCORE_VM_FIXED0160_A1_50G",
		"SUBCORE_VM_FIXED0180_A1_50G",
		"SUBCORE_VM_FIXED0200_A1_50G",
		"SUBCORE_VM_FIXED0220_A1_50G",
		"SUBCORE_VM_FIXED0240_A1_50G",
		"SUBCORE_VM_FIXED0260_A1_50G",
		"SUBCORE_VM_FIXED0280_A1_50G",
		"SUBCORE_VM_FIXED0300_A1_50G",
		"SUBCORE_VM_FIXED0320_A1_50G",
		"SUBCORE_VM_FIXED0340_A1_50G",
		"SUBCORE_VM_FIXED0360_A1_50G",
		"SUBCORE_VM_FIXED0380_A1_50G",
		"SUBCORE_VM_FIXED0400_A1_50G",
		"SUBCORE_VM_FIXED0420_A1_50G",
		"SUBCORE_VM_FIXED0440_A1_50G",
		"SUBCORE_VM_FIXED0460_A1_50G",
		"SUBCORE_VM_FIXED0480_A1_50G",
		"SUBCORE_VM_FIXED0500_A1_50G",
		"SUBCORE_VM_FIXED0520_A1_50G",
		"SUBCORE_VM_FIXED0540_A1_50G",
		"SUBCORE_VM_FIXED0560_A1_50G",
		"SUBCORE_VM_FIXED0580_A1_50G",
		"SUBCORE_VM_FIXED0600_A1_50G",
		"SUBCORE_VM_FIXED0620_A1_50G",
		"SUBCORE_VM_FIXED0640_A1_50G",
		"SUBCORE_VM_FIXED0660_A1_50G",
		"SUBCORE_VM_FIXED0680_A1_50G",
		"SUBCORE_VM_FIXED0700_A1_50G",
		"SUBCORE_VM_FIXED0720_A1_50G",
		"SUBCORE_VM_FIXED0740_A1_50G",
		"SUBCORE_VM_FIXED0760_A1_50G",
		"SUBCORE_VM_FIXED0780_A1_50G",
		"SUBCORE_VM_FIXED0800_A1_50G",
		"SUBCORE_VM_FIXED0820_A1_50G",
		"SUBCORE_VM_FIXED0840_A1_50G",
		"SUBCORE_VM_FIXED0860_A1_50G",
		"SUBCORE_VM_FIXED0880_A1_50G",
		"SUBCORE_VM_FIXED0900_A1_50G",
		"SUBCORE_VM_FIXED0920_A1_50G",
		"SUBCORE_VM_FIXED0940_A1_50G",
		"SUBCORE_VM_FIXED0960_A1_50G",
		"SUBCORE_VM_FIXED0980_A1_50G",
		"SUBCORE_VM_FIXED1000_A1_50G",
		"SUBCORE_VM_FIXED1020_A1_50G",
		"SUBCORE_VM_FIXED1040_A1_50G",
		"SUBCORE_VM_FIXED1060_A1_50G",
		"SUBCORE_VM_FIXED1080_A1_50G",
		"SUBCORE_VM_FIXED1100_A1_50G",
		"SUBCORE_VM_FIXED1120_A1_50G",
		"SUBCORE_VM_FIXED1140_A1_50G",
		"SUBCORE_VM_FIXED1160_A1_50G",
		"SUBCORE_VM_FIXED1180_A1_50G",
		"SUBCORE_VM_FIXED1200_A1_50G",
		"SUBCORE_VM_FIXED1220_A1_50G",
		"SUBCORE_VM_FIXED1240_A1_50G",
		"SUBCORE_VM_FIXED1260_A1_50G",
		"SUBCORE_VM_FIXED1280_A1_50G",
		"SUBCORE_VM_FIXED1300_A1_50G",
		"SUBCORE_VM_FIXED1320_A1_50G",
		"SUBCORE_VM_FIXED1340_A1_50G",
		"SUBCORE_VM_FIXED1360_A1_50G",
		"SUBCORE_VM_FIXED1380_A1_50G",
		"SUBCORE_VM_FIXED1400_A1_50G",
		"SUBCORE_VM_FIXED1420_A1_50G",
		"SUBCORE_VM_FIXED1440_A1_50G",
		"SUBCORE_VM_FIXED1460_A1_50G",
		"SUBCORE_VM_FIXED1480_A1_50G",
		"SUBCORE_VM_FIXED1500_A1_50G",
		"SUBCORE_VM_FIXED1520_A1_50G",
		"SUBCORE_VM_FIXED1540_A1_50G",
		"SUBCORE_VM_FIXED1560_A1_50G",
		"SUBCORE_VM_FIXED1580_A1_50G",
		"SUBCORE_VM_FIXED1600_A1_50G",
		"SUBCORE_VM_FIXED1620_A1_50G",
		"SUBCORE_VM_FIXED1640_A1_50G",
		"SUBCORE_VM_FIXED1660_A1_50G",
		"SUBCORE_VM_FIXED1680_A1_50G",
		"SUBCORE_VM_FIXED1700_A1_50G",
		"SUBCORE_VM_FIXED1720_A1_50G",
		"SUBCORE_VM_FIXED1740_A1_50G",
		"SUBCORE_VM_FIXED1760_A1_50G",
		"SUBCORE_VM_FIXED1780_A1_50G",
		"SUBCORE_VM_FIXED1800_A1_50G",
		"SUBCORE_VM_FIXED1820_A1_50G",
		"SUBCORE_VM_FIXED1840_A1_50G",
		"SUBCORE_VM_FIXED1860_A1_50G",
		"SUBCORE_VM_FIXED1880_A1_50G",
		"SUBCORE_VM_FIXED1900_A1_50G",
		"SUBCORE_VM_FIXED1920_A1_50G",
		"SUBCORE_VM_FIXED1940_A1_50G",
		"SUBCORE_VM_FIXED1960_A1_50G",
		"SUBCORE_VM_FIXED1980_A1_50G",
		"SUBCORE_VM_FIXED2000_A1_50G",
		"SUBCORE_VM_FIXED2020_A1_50G",
		"SUBCORE_VM_FIXED2040_A1_50G",
		"SUBCORE_VM_FIXED2060_A1_50G",
		"SUBCORE_VM_FIXED2080_A1_50G",
		"SUBCORE_VM_FIXED2100_A1_50G",
		"SUBCORE_VM_FIXED2120_A1_50G",
		"SUBCORE_VM_FIXED2140_A1_50G",
		"SUBCORE_VM_FIXED2160_A1_50G",
		"SUBCORE_VM_FIXED2180_A1_50G",
		"SUBCORE_VM_FIXED2200_A1_50G",
		"SUBCORE_VM_FIXED2220_A1_50G",
		"SUBCORE_VM_FIXED2240_A1_50G",
		"SUBCORE_VM_FIXED2260_A1_50G",
		"SUBCORE_VM_FIXED2280_A1_50G",
		"SUBCORE_VM_FIXED2300_A1_50G",
		"SUBCORE_VM_FIXED2320_A1_50G",
		"SUBCORE_VM_FIXED2340_A1_50G",
		"SUBCORE_VM_FIXED2360_A1_50G",
		"SUBCORE_VM_FIXED2380_A1_50G",
		"SUBCORE_VM_FIXED2400_A1_50G",
		"SUBCORE_VM_FIXED2420_A1_50G",
		"SUBCORE_VM_FIXED2440_A1_50G",
		"SUBCORE_VM_FIXED2460_A1_50G",
		"SUBCORE_VM_FIXED2480_A1_50G",
		"SUBCORE_VM_FIXED2500_A1_50G",
		"SUBCORE_VM_FIXED2520_A1_50G",
		"SUBCORE_VM_FIXED2540_A1_50G",
		"SUBCORE_VM_FIXED2560_A1_50G",
		"SUBCORE_VM_FIXED2580_A1_50G",
		"SUBCORE_VM_FIXED2600_A1_50G",
		"SUBCORE_VM_FIXED2620_A1_50G",
		"SUBCORE_VM_FIXED2640_A1_50G",
		"SUBCORE_VM_FIXED2660_A1_50G",
		"SUBCORE_VM_FIXED2680_A1_50G",
		"SUBCORE_VM_FIXED2700_A1_50G",
		"SUBCORE_VM_FIXED2720_A1_50G",
		"SUBCORE_VM_FIXED2740_A1_50G",
		"SUBCORE_VM_FIXED2760_A1_50G",
		"SUBCORE_VM_FIXED2780_A1_50G",
		"SUBCORE_VM_FIXED2800_A1_50G",
		"SUBCORE_VM_FIXED2820_A1_50G",
		"SUBCORE_VM_FIXED2840_A1_50G",
		"SUBCORE_VM_FIXED2860_A1_50G",
		"SUBCORE_VM_FIXED2880_A1_50G",
		"SUBCORE_VM_FIXED2900_A1_50G",
		"SUBCORE_VM_FIXED2920_A1_50G",
		"SUBCORE_VM_FIXED2940_A1_50G",
		"SUBCORE_VM_FIXED2960_A1_50G",
		"SUBCORE_VM_FIXED2980_A1_50G",
		"SUBCORE_VM_FIXED3000_A1_50G",
		"SUBCORE_VM_FIXED3020_A1_50G",
		"SUBCORE_VM_FIXED3040_A1_50G",
		"SUBCORE_VM_FIXED3060_A1_50G",
		"SUBCORE_VM_FIXED3080_A1_50G",
		"SUBCORE_VM_FIXED3100_A1_50G",
		"SUBCORE_VM_FIXED3120_A1_50G",
		"SUBCORE_VM_FIXED3140_A1_50G",
		"SUBCORE_VM_FIXED3160_A1_50G",
		"SUBCORE_VM_FIXED3180_A1_50G",
		"SUBCORE_VM_FIXED3200_A1_50G",
		"SUBCORE_VM_FIXED3220_A1_50G",
		"SUBCORE_VM_FIXED3240_A1_50G",
		"SUBCORE_VM_FIXED3260_A1_50G",
		"SUBCORE_VM_FIXED3280_A1_50G",
		"SUBCORE_VM_FIXED3300_A1_50G",
		"SUBCORE_VM_FIXED3320_A1_50G",
		"SUBCORE_VM_FIXED3340_A1_50G",
		"SUBCORE_VM_FIXED3360_A1_50G",
		"SUBCORE_VM_FIXED3380_A1_50G",
		"SUBCORE_VM_FIXED3400_A1_50G",
		"SUBCORE_VM_FIXED3420_A1_50G",
		"SUBCORE_VM_FIXED3440_A1_50G",
		"SUBCORE_VM_FIXED3460_A1_50G",
		"SUBCORE_VM_FIXED3480_A1_50G",
		"SUBCORE_VM_FIXED3500_A1_50G",
		"SUBCORE_VM_FIXED3520_A1_50G",
		"SUBCORE_VM_FIXED3540_A1_50G",
		"SUBCORE_VM_FIXED3560_A1_50G",
		"SUBCORE_VM_FIXED3580_A1_50G",
		"SUBCORE_VM_FIXED3600_A1_50G",
		"SUBCORE_VM_FIXED3620_A1_50G",
		"SUBCORE_VM_FIXED3640_A1_50G",
		"SUBCORE_VM_FIXED3660_A1_50G",
		"SUBCORE_VM_FIXED3680_A1_50G",
		"SUBCORE_VM_FIXED3700_A1_50G",
		"SUBCORE_VM_FIXED3720_A1_50G",
		"SUBCORE_VM_FIXED3740_A1_50G",
		"SUBCORE_VM_FIXED3760_A1_50G",
		"SUBCORE_VM_FIXED3780_A1_50G",
		"SUBCORE_VM_FIXED3800_A1_50G",
		"SUBCORE_VM_FIXED3820_A1_50G",
		"SUBCORE_VM_FIXED3840_A1_50G",
		"SUBCORE_VM_FIXED3860_A1_50G",
		"SUBCORE_VM_FIXED3880_A1_50G",
		"SUBCORE_VM_FIXED3900_A1_50G",
		"SUBCORE_VM_FIXED3920_A1_50G",
		"SUBCORE_VM_FIXED3940_A1_50G",
		"SUBCORE_VM_FIXED3960_A1_50G",
		"SUBCORE_VM_FIXED3980_A1_50G",
		"SUBCORE_VM_FIXED4000_A1_50G",
		"SUBCORE_VM_FIXED4020_A1_50G",
		"SUBCORE_VM_FIXED4040_A1_50G",
		"SUBCORE_VM_FIXED4060_A1_50G",
		"SUBCORE_VM_FIXED4080_A1_50G",
		"SUBCORE_VM_FIXED4100_A1_50G",
		"SUBCORE_VM_FIXED4120_A1_50G",
		"SUBCORE_VM_FIXED4140_A1_50G",
		"SUBCORE_VM_FIXED4160_A1_50G",
		"SUBCORE_VM_FIXED4180_A1_50G",
		"SUBCORE_VM_FIXED4200_A1_50G",
		"SUBCORE_VM_FIXED4220_A1_50G",
		"SUBCORE_VM_FIXED4240_A1_50G",
		"SUBCORE_VM_FIXED4260_A1_50G",
		"SUBCORE_VM_FIXED4280_A1_50G",
		"SUBCORE_VM_FIXED4300_A1_50G",
		"SUBCORE_VM_FIXED4320_A1_50G",
		"SUBCORE_VM_FIXED4340_A1_50G",
		"SUBCORE_VM_FIXED4360_A1_50G",
		"SUBCORE_VM_FIXED4380_A1_50G",
		"SUBCORE_VM_FIXED4400_A1_50G",
		"SUBCORE_VM_FIXED4420_A1_50G",
		"SUBCORE_VM_FIXED4440_A1_50G",
		"SUBCORE_VM_FIXED4460_A1_50G",
		"SUBCORE_VM_FIXED4480_A1_50G",
		"SUBCORE_VM_FIXED4500_A1_50G",
		"SUBCORE_VM_FIXED4520_A1_50G",
		"SUBCORE_VM_FIXED4540_A1_50G",
		"SUBCORE_VM_FIXED4560_A1_50G",
		"SUBCORE_VM_FIXED4580_A1_50G",
		"SUBCORE_VM_FIXED4600_A1_50G",
		"SUBCORE_VM_FIXED4620_A1_50G",
		"SUBCORE_VM_FIXED4640_A1_50G",
		"SUBCORE_VM_FIXED4660_A1_50G",
		"SUBCORE_VM_FIXED4680_A1_50G",
		"SUBCORE_VM_FIXED4700_A1_50G",
		"SUBCORE_VM_FIXED4720_A1_50G",
		"SUBCORE_VM_FIXED4740_A1_50G",
		"SUBCORE_VM_FIXED4760_A1_50G",
		"SUBCORE_VM_FIXED4780_A1_50G",
		"SUBCORE_VM_FIXED4800_A1_50G",
		"SUBCORE_VM_FIXED4820_A1_50G",
		"SUBCORE_VM_FIXED4840_A1_50G",
		"SUBCORE_VM_FIXED4860_A1_50G",
		"SUBCORE_VM_FIXED4880_A1_50G",
		"SUBCORE_VM_FIXED4900_A1_50G",
		"SUBCORE_VM_FIXED4920_A1_50G",
		"SUBCORE_VM_FIXED4940_A1_50G",
		"SUBCORE_VM_FIXED4960_A1_50G",
		"SUBCORE_VM_FIXED4980_A1_50G",
		"SUBCORE_VM_FIXED5000_A1_50G",
		"SUBCORE_VM_FIXED0090_X9_50G",
		"SUBCORE_VM_FIXED0180_X9_50G",
		"SUBCORE_VM_FIXED0270_X9_50G",
		"SUBCORE_VM_FIXED0360_X9_50G",
		"SUBCORE_VM_FIXED0450_X9_50G",
		"SUBCORE_VM_FIXED0540_X9_50G",
		"SUBCORE_VM_FIXED0630_X9_50G",
		"SUBCORE_VM_FIXED0720_X9_50G",
		"SUBCORE_VM_FIXED0810_X9_50G",
		"SUBCORE_VM_FIXED0900_X9_50G",
		"SUBCORE_VM_FIXED0990_X9_50G",
		"SUBCORE_VM_FIXED1080_X9_50G",
		"SUBCORE_VM_FIXED1170_X9_50G",
		"SUBCORE_VM_FIXED1260_X9_50G",
		"SUBCORE_VM_FIXED1350_X9_50G",
		"SUBCORE_VM_FIXED1440_X9_50G",
		"SUBCORE_VM_FIXED1530_X9_50G",
		"SUBCORE_VM_FIXED1620_X9_50G",
		"SUBCORE_VM_FIXED1710_X9_50G",
		"SUBCORE_VM_FIXED1800_X9_50G",
		"SUBCORE_VM_FIXED1890_X9_50G",
		"SUBCORE_VM_FIXED1980_X9_50G",
		"SUBCORE_VM_FIXED2070_X9_50G",
		"SUBCORE_VM_FIXED2160_X9_50G",
		"SUBCORE_VM_FIXED2250_X9_50G",
		"SUBCORE_VM_FIXED2340_X9_50G",
		"SUBCORE_VM_FIXED2430_X9_50G",
		"SUBCORE_VM_FIXED2520_X9_50G",
		"SUBCORE_VM_FIXED2610_X9_50G",
		"SUBCORE_VM_FIXED2700_X9_50G",
		"SUBCORE_VM_FIXED2790_X9_50G",
		"SUBCORE_VM_FIXED2880_X9_50G",
		"SUBCORE_VM_FIXED2970_X9_50G",
		"SUBCORE_VM_FIXED3060_X9_50G",
		"SUBCORE_VM_FIXED3150_X9_50G",
		"SUBCORE_VM_FIXED3240_X9_50G",
		"SUBCORE_VM_FIXED3330_X9_50G",
		"SUBCORE_VM_FIXED3420_X9_50G",
		"SUBCORE_VM_FIXED3510_X9_50G",
		"SUBCORE_VM_FIXED3600_X9_50G",
		"SUBCORE_VM_FIXED3690_X9_50G",
		"SUBCORE_VM_FIXED3780_X9_50G",
		"SUBCORE_VM_FIXED3870_X9_50G",
		"SUBCORE_VM_FIXED3960_X9_50G",
		"SUBCORE_VM_FIXED4050_X9_50G",
		"SUBCORE_VM_FIXED4140_X9_50G",
		"SUBCORE_VM_FIXED4230_X9_50G",
		"SUBCORE_VM_FIXED4320_X9_50G",
		"SUBCORE_VM_FIXED4410_X9_50G",
		"SUBCORE_VM_FIXED4500_X9_50G",
		"SUBCORE_VM_FIXED4590_X9_50G",
		"SUBCORE_VM_FIXED4680_X9_50G",
		"SUBCORE_VM_FIXED4770_X9_50G",
		"SUBCORE_VM_FIXED4860_X9_50G",
		"SUBCORE_VM_FIXED4950_X9_50G",
		"DYNAMIC_A1_50G",
		"FIXED0040_A1_50G",
		"FIXED0100_A1_50G",
		"FIXED0200_A1_50G",
		"FIXED0300_A1_50G",
		"FIXED0400_A1_50G",
		"FIXED0500_A1_50G",
		"FIXED0600_A1_50G",
		"FIXED0700_A1_50G",
		"FIXED0800_A1_50G",
		"FIXED0900_A1_50G",
		"FIXED1000_A1_50G",
		"FIXED1100_A1_50G",
		"FIXED1200_A1_50G",
		"FIXED1300_A1_50G",
		"FIXED1400_A1_50G",
		"FIXED1500_A1_50G",
		"FIXED1600_A1_50G",
		"FIXED1700_A1_50G",
		"FIXED1800_A1_50G",
		"FIXED1900_A1_50G",
		"FIXED2000_A1_50G",
		"FIXED2100_A1_50G",
		"FIXED2200_A1_50G",
		"FIXED2300_A1_50G",
		"FIXED2400_A1_50G",
		"FIXED2500_A1_50G",
		"FIXED2600_A1_50G",
		"FIXED2700_A1_50G",
		"FIXED2800_A1_50G",
		"FIXED2900_A1_50G",
		"FIXED3000_A1_50G",
		"FIXED3100_A1_50G",
		"FIXED3200_A1_50G",
		"FIXED3300_A1_50G",
		"FIXED3400_A1_50G",
		"FIXED3500_A1_50G",
		"FIXED3600_A1_50G",
		"FIXED3700_A1_50G",
		"FIXED3800_A1_50G",
		"FIXED3900_A1_50G",
		"FIXED4000_A1_50G",
		"ENTIREHOST_A1_50G",
		"DYNAMIC_X9_50G",
		"FIXED0040_X9_50G",
		"FIXED0400_X9_50G",
		"FIXED0800_X9_50G",
		"FIXED1200_X9_50G",
		"FIXED1600_X9_50G",
		"FIXED2000_X9_50G",
		"FIXED2400_X9_50G",
		"FIXED2800_X9_50G",
		"FIXED3200_X9_50G",
		"FIXED3600_X9_50G",
		"FIXED4000_X9_50G",
		"STANDARD_VM_FIXED0100_X9_50G",
		"STANDARD_VM_FIXED0200_X9_50G",
		"STANDARD_VM_FIXED0300_X9_50G",
		"STANDARD_VM_FIXED0400_X9_50G",
		"STANDARD_VM_FIXED0500_X9_50G",
		"STANDARD_VM_FIXED0600_X9_50G",
		"STANDARD_VM_FIXED0700_X9_50G",
		"STANDARD_VM_FIXED0800_X9_50G",
		"STANDARD_VM_FIXED0900_X9_50G",
		"STANDARD_VM_FIXED1000_X9_50G",
		"STANDARD_VM_FIXED1100_X9_50G",
		"STANDARD_VM_FIXED1200_X9_50G",
		"STANDARD_VM_FIXED1300_X9_50G",
		"STANDARD_VM_FIXED1400_X9_50G",
		"STANDARD_VM_FIXED1500_X9_50G",
		"STANDARD_VM_FIXED1600_X9_50G",
		"STANDARD_VM_FIXED1700_X9_50G",
		"STANDARD_VM_FIXED1800_X9_50G",
		"STANDARD_VM_FIXED1900_X9_50G",
		"STANDARD_VM_FIXED2000_X9_50G",
		"STANDARD_VM_FIXED2100_X9_50G",
		"STANDARD_VM_FIXED2200_X9_50G",
		"STANDARD_VM_FIXED2300_X9_50G",
		"STANDARD_VM_FIXED2400_X9_50G",
		"STANDARD_VM_FIXED2500_X9_50G",
		"STANDARD_VM_FIXED2600_X9_50G",
		"STANDARD_VM_FIXED2700_X9_50G",
		"STANDARD_VM_FIXED2800_X9_50G",
		"STANDARD_VM_FIXED2900_X9_50G",
		"STANDARD_VM_FIXED3000_X9_50G",
		"STANDARD_VM_FIXED3100_X9_50G",
		"STANDARD_VM_FIXED3200_X9_50G",
		"STANDARD_VM_FIXED3300_X9_50G",
		"STANDARD_VM_FIXED3400_X9_50G",
		"STANDARD_VM_FIXED3500_X9_50G",
		"STANDARD_VM_FIXED3600_X9_50G",
		"STANDARD_VM_FIXED3700_X9_50G",
		"STANDARD_VM_FIXED3800_X9_50G",
		"STANDARD_VM_FIXED3900_X9_50G",
		"STANDARD_VM_FIXED4000_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0025_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0050_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0075_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0100_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0125_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0150_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0175_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0200_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0225_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0250_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0275_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0300_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0325_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0350_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0375_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0400_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0425_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0450_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0475_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0500_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0525_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0550_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0575_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0600_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0625_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0650_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0675_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0700_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0725_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0750_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0775_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0800_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0825_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0850_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0875_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0900_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0925_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0950_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0975_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1000_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1025_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1050_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1075_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1100_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1125_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1150_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1175_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1200_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1225_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1250_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1275_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1300_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1325_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1350_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1375_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1400_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1425_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1450_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1475_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1500_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1525_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1550_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1575_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1600_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1625_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1650_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1700_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1725_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1750_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1800_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1850_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1875_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1900_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1925_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1950_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2000_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2025_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2050_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2100_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2125_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2150_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2175_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2200_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2250_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2275_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2300_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2325_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2350_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2375_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2400_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2450_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2475_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2500_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2550_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2600_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2625_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2650_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2700_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2750_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2775_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2800_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2850_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2875_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2900_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2925_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2950_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2975_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3000_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3025_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3050_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3075_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3100_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3125_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3150_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3200_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3225_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3250_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3300_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3325_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3375_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3400_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3450_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3500_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3525_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3575_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3600_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3625_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3675_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3700_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3750_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3800_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3825_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3850_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3875_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3900_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3975_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4000_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4025_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4050_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4100_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4125_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4200_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4225_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4250_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4275_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4300_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4350_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4375_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4400_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4425_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4500_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4550_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4575_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4600_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4625_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4650_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4675_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4700_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4725_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4750_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4800_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4875_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4900_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4950_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED5000_X9_50G",
		"ENTIREHOST_X9_50G",
	}
}

// GetMappingInternalVnicAttachmentVnicShapeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternalVnicAttachmentVnicShapeEnum(val string) (InternalVnicAttachmentVnicShapeEnum, bool) {
	enum, ok := mappingInternalVnicAttachmentVnicShapeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
