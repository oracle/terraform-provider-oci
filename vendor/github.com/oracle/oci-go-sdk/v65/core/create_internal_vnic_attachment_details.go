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
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateInternalVnicAttachmentDetails Details for attaching a service VNIC to VNICaaS fleet.
type CreateInternalVnicAttachmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the VNIC attachment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A private IP address of your choice to assign to the VNIC. Must be an
	// available IP address within the subnet's CIDR. If you don't specify a
	// value, Oracle automatically assigns a private IP address from the subnet.
	// This is the VNIC's *primary* private IP address. The value appears in
	// the Vnic object and also the
	// PrivateIp object returned by
	// ListPrivateIps and
	// GetPrivateIp.
	// Example: `10.0.3.3`
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// The compute instance id
	InstanceId *string `mandatory:"false" json:"instanceId"`

	// Id for a group of vnics sharing same resource pool, e.g. a group id could be a site/gateway id for
	// Ambassador SVNICs under the same site/gateway.
	GroupId *string `mandatory:"false" json:"groupId"`

	// The compute instance id of the resource pool to be used for the vnic
	InstanceIdForResourcePool *string `mandatory:"false" json:"instanceIdForResourcePool"`

	// The availability domain of the VNIC attachment
	InternalAvailabilityDomain *string `mandatory:"false" json:"internalAvailabilityDomain"`

	// The overlay MAC address of the instance
	MacAddress *string `mandatory:"false" json:"macAddress"`

	// index of NIC that VNIC is attaching to (OS boot order)
	NicIndex *int `mandatory:"false" json:"nicIndex"`

	// The tag used internally to identify the sending VNIC. It can be specified in scenarios where a specific
	// tag needs to be assigned. Examples of such scenarios include reboot migration and VMware support.
	VlanTag *int `mandatory:"false" json:"vlanTag"`

	// Shape of VNIC that is used to allocate resource in the data plane.
	VnicShape CreateInternalVnicAttachmentDetailsVnicShapeEnum `mandatory:"false" json:"vnicShape,omitempty"`

	// The substrate IP address of the instance
	SubstrateIp *string `mandatory:"false" json:"substrateIp"`

	// Indicates if vlanTag 0 can be assigned to this vnic or not.
	IsSkipVlanTag0 *bool `mandatory:"false" json:"isSkipVlanTag0"`

	// Specifies the shard to attach the VNIC to
	ShardId *string `mandatory:"false" json:"shardId"`

	// Property describing customer facing metrics
	MetadataList []CfmMetadata `mandatory:"false" json:"metadataList"`

	// Identifier of how the target instance is going be launched. For example, launch from marketplace.
	// STANDARD is the default type if not specified.
	LaunchType CreateInternalVnicAttachmentDetailsLaunchTypeEnum `mandatory:"false" json:"launchType,omitempty"`
}

func (m CreateInternalVnicAttachmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateInternalVnicAttachmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateInternalVnicAttachmentDetailsVnicShapeEnum(string(m.VnicShape)); !ok && m.VnicShape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VnicShape: %s. Supported values are: %s.", m.VnicShape, strings.Join(GetCreateInternalVnicAttachmentDetailsVnicShapeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateInternalVnicAttachmentDetailsLaunchTypeEnum(string(m.LaunchType)); !ok && m.LaunchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LaunchType: %s. Supported values are: %s.", m.LaunchType, strings.Join(GetCreateInternalVnicAttachmentDetailsLaunchTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateInternalVnicAttachmentDetailsVnicShapeEnum Enum with underlying type: string
type CreateInternalVnicAttachmentDetailsVnicShapeEnum string

// Set of constants representing the allowable values for CreateInternalVnicAttachmentDetailsVnicShapeEnum
const (
	CreateInternalVnicAttachmentDetailsVnicShapeDynamic                         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "DYNAMIC"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0040                       CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0040"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0060                       CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0060"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0060Psm                    CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0060_PSM"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0100                       CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0100"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0120                       CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0120"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed01202x                     CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0120_2X"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0200                       CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0200"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0240                       CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0240"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0480                       CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0480"
	CreateInternalVnicAttachmentDetailsVnicShapeEntirehost                      CreateInternalVnicAttachmentDetailsVnicShapeEnum = "ENTIREHOST"
	CreateInternalVnicAttachmentDetailsVnicShapeDynamic25g                      CreateInternalVnicAttachmentDetailsVnicShapeEnum = "DYNAMIC_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed004025g                    CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0040_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed010025g                    CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0100_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed020025g                    CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0200_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed040025g                    CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0400_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed080025g                    CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0800_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed160025g                    CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1600_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed240025g                    CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2400_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeEntirehost25g                   CreateInternalVnicAttachmentDetailsVnicShapeEnum = "ENTIREHOST_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeDynamicE125g                    CreateInternalVnicAttachmentDetailsVnicShapeEnum = "DYNAMIC_E1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0040E125g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0040_E1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0070E125g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0070_E1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0140E125g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0140_E1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0280E125g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0280_E1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0560E125g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0560_E1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1120E125g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1120_E1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1680E125g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1680_E1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeEntirehostE125g                 CreateInternalVnicAttachmentDetailsVnicShapeEnum = "ENTIREHOST_E1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeDynamicB125g                    CreateInternalVnicAttachmentDetailsVnicShapeEnum = "DYNAMIC_B1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0040B125g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0040_B1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0060B125g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0060_B1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0120B125g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0120_B1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0240B125g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0240_B1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0480B125g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0480_B1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0960B125g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0960_B1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeEntirehostB125g                 CreateInternalVnicAttachmentDetailsVnicShapeEnum = "ENTIREHOST_B1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeMicroVmFixed0048E125g           CreateInternalVnicAttachmentDetailsVnicShapeEnum = "MICRO_VM_FIXED0048_E1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeMicroLbFixed0001E125g           CreateInternalVnicAttachmentDetailsVnicShapeEnum = "MICRO_LB_FIXED0001_E1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeVnicaasFixed0200                CreateInternalVnicAttachmentDetailsVnicShapeEnum = "VNICAAS_FIXED0200"
	CreateInternalVnicAttachmentDetailsVnicShapeVnicaasFixed0400                CreateInternalVnicAttachmentDetailsVnicShapeEnum = "VNICAAS_FIXED0400"
	CreateInternalVnicAttachmentDetailsVnicShapeVnicaasFixed0700                CreateInternalVnicAttachmentDetailsVnicShapeEnum = "VNICAAS_FIXED0700"
	CreateInternalVnicAttachmentDetailsVnicShapeVnicaasNlbApproved10g           CreateInternalVnicAttachmentDetailsVnicShapeEnum = "VNICAAS_NLB_APPROVED_10G"
	CreateInternalVnicAttachmentDetailsVnicShapeVnicaasNlbApproved25g           CreateInternalVnicAttachmentDetailsVnicShapeEnum = "VNICAAS_NLB_APPROVED_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeVnicaasTelesis25g               CreateInternalVnicAttachmentDetailsVnicShapeEnum = "VNICAAS_TELESIS_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeVnicaasTelesis10g               CreateInternalVnicAttachmentDetailsVnicShapeEnum = "VNICAAS_TELESIS_10G"
	CreateInternalVnicAttachmentDetailsVnicShapeVnicaasAmbassadorFixed0100      CreateInternalVnicAttachmentDetailsVnicShapeEnum = "VNICAAS_AMBASSADOR_FIXED0100"
	CreateInternalVnicAttachmentDetailsVnicShapeVnicaasTelesisGamma             CreateInternalVnicAttachmentDetailsVnicShapeEnum = "VNICAAS_TELESIS_GAMMA"
	CreateInternalVnicAttachmentDetailsVnicShapeVnicaasPrivatedns               CreateInternalVnicAttachmentDetailsVnicShapeEnum = "VNICAAS_PRIVATEDNS"
	CreateInternalVnicAttachmentDetailsVnicShapeVnicaasFwaas                    CreateInternalVnicAttachmentDetailsVnicShapeEnum = "VNICAAS_FWAAS"
	CreateInternalVnicAttachmentDetailsVnicShapeVnicaasLbaasFree                CreateInternalVnicAttachmentDetailsVnicShapeEnum = "VNICAAS_LBAAS_FREE"
	CreateInternalVnicAttachmentDetailsVnicShapeVnicaasLbaas8g512k              CreateInternalVnicAttachmentDetailsVnicShapeEnum = "VNICAAS_LBAAS_8G_512K"
	CreateInternalVnicAttachmentDetailsVnicShapeVnicaasLbaas8g1m                CreateInternalVnicAttachmentDetailsVnicShapeEnum = "VNICAAS_LBAAS_8G_1M"
	CreateInternalVnicAttachmentDetailsVnicShapeVnicaasLbaas8g2m                CreateInternalVnicAttachmentDetailsVnicShapeEnum = "VNICAAS_LBAAS_8G_2M"
	CreateInternalVnicAttachmentDetailsVnicShapeVnicaasLbaas8g3m                CreateInternalVnicAttachmentDetailsVnicShapeEnum = "VNICAAS_LBAAS_8G_3M"
	CreateInternalVnicAttachmentDetailsVnicShapeVnicaasLbaas8g1m8ghost          CreateInternalVnicAttachmentDetailsVnicShapeEnum = "VNICAAS_LBAAS_8G_1M_8GHOST"
	CreateInternalVnicAttachmentDetailsVnicShapeVnicaasLbaas8g1m16ghost         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "VNICAAS_LBAAS_8G_1M_16GHOST"
	CreateInternalVnicAttachmentDetailsVnicShapeDynamicE350g                    CreateInternalVnicAttachmentDetailsVnicShapeEnum = "DYNAMIC_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0040E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0040_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0100E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0100_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0200E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0200_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0300E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0300_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0400E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0400_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0500E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0500_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0600E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0600_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0700E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0700_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0800E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0800_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0900E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0900_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1000E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1000_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1100E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1100_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1200E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1200_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1300E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1300_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1400E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1400_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1500E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1500_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1600E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1600_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1700E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1700_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1800E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1800_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1900E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1900_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2000E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2000_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2100E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2100_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2200E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2200_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2300E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2300_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2400E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2400_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2500E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2500_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2600E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2600_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2700E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2700_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2800E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2800_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2900E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2900_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3000E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3000_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3100E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3100_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3200E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3200_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3300E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3300_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3400E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3400_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3500E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3500_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3600E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3600_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3700E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3700_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3800E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3800_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3900E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3900_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed4000E350g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED4000_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeEntirehostE350g                 CreateInternalVnicAttachmentDetailsVnicShapeEnum = "ENTIREHOST_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeDynamicE450g                    CreateInternalVnicAttachmentDetailsVnicShapeEnum = "DYNAMIC_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0040E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0040_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0100E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0100_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0200E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0200_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0300E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0300_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0400E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0400_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0500E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0500_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0600E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0600_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0700E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0700_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0800E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0800_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0900E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0900_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1000E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1000_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1100E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1100_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1200E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1200_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1300E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1300_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1400E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1400_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1500E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1500_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1600E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1600_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1700E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1700_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1800E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1800_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1900E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1900_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2000E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2000_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2100E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2100_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2200E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2200_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2300E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2300_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2400E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2400_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2500E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2500_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2600E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2600_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2700E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2700_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2800E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2800_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2900E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2900_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3000E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3000_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3100E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3100_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3200E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3200_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3300E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3300_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3400E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3400_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3500E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3500_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3600E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3600_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3700E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3700_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3800E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3800_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3900E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3900_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed4000E450g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED4000_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeEntirehostE450g                 CreateInternalVnicAttachmentDetailsVnicShapeEnum = "ENTIREHOST_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeMicroVmFixed0050E350g           CreateInternalVnicAttachmentDetailsVnicShapeEnum = "Micro_VM_Fixed0050_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeMicroVmFixed0050E450g           CreateInternalVnicAttachmentDetailsVnicShapeEnum = "Micro_VM_Fixed0050_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0025E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0025_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0050E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0050_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0075E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0075_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0100E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0100_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0125E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0125_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0150E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0150_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0175E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0175_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0200E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0200_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0225E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0225_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0250E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0250_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0275E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0275_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0300E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0300_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0325E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0325_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0350E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0350_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0375E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0375_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0400E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0400_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0425E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0425_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0450E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0450_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0475E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0475_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0500E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0500_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0525E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0525_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0550E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0550_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0575E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0575_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0600E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0600_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0625E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0625_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0650E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0650_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0675E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0675_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0700E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0700_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0725E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0725_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0750E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0750_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0775E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0775_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0800E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0800_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0825E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0825_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0850E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0850_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0875E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0875_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0900E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0900_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0925E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0925_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0950E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0950_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0975E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0975_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1000E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1000_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1025E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1025_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1050E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1050_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1075E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1075_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1100E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1100_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1125E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1125_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1150E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1150_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1175E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1175_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1200E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1200_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1225E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1225_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1250E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1250_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1275E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1275_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1300E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1300_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1325E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1325_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1350E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1350_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1375E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1375_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1400E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1400_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1425E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1425_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1450E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1450_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1475E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1475_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1500E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1500_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1525E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1525_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1550E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1550_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1575E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1575_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1600E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1600_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1625E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1625_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1650E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1650_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1700E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1700_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1725E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1725_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1750E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1750_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1800E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1800_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1850E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1850_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1875E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1875_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1900E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1900_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1925E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1925_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1950E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1950_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2000E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2000_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2025E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2025_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2050E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2050_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2100E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2100_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2125E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2125_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2150E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2150_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2175E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2175_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2200E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2200_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2250E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2250_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2275E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2275_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2300E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2300_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2325E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2325_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2350E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2350_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2375E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2375_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2400E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2400_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2450E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2450_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2475E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2475_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2500E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2500_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2550E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2550_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2600E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2600_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2625E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2625_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2650E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2650_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2700E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2700_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2750E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2750_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2775E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2775_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2800E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2800_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2850E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2850_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2875E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2875_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2900E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2900_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2925E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2925_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2950E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2950_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2975E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2975_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3000E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3000_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3025E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3025_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3050E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3050_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3075E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3075_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3100E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3100_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3125E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3125_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3150E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3150_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3200E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3200_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3225E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3225_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3250E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3250_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3300E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3300_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3325E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3325_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3375E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3375_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3400E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3400_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3450E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3450_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3500E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3500_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3525E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3525_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3575E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3575_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3600E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3600_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3625E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3625_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3675E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3675_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3700E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3700_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3750E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3750_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3800E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3800_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3825E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3825_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3850E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3850_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3875E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3875_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3900E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3900_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3975E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3975_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4000E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4000_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4025E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4025_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4050E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4050_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4100E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4100_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4125E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4125_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4200E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4200_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4225E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4225_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4250E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4250_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4275E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4275_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4300E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4300_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4350E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4350_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4375E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4375_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4400E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4400_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4425E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4425_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4500E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4500_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4550E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4550_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4575E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4575_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4600E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4600_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4625E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4625_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4650E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4650_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4675E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4675_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4700E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4700_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4725E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4725_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4750E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4750_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4800E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4800_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4875E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4875_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4900E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4900_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4950E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4950_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed5000E350g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED5000_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0025E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0025_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0050E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0050_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0075E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0075_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0100E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0100_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0125E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0125_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0150E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0150_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0175E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0175_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0200E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0200_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0225E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0225_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0250E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0250_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0275E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0275_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0300E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0300_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0325E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0325_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0350E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0350_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0375E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0375_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0400E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0400_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0425E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0425_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0450E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0450_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0475E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0475_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0500E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0500_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0525E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0525_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0550E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0550_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0575E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0575_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0600E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0600_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0625E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0625_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0650E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0650_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0675E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0675_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0700E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0700_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0725E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0725_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0750E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0750_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0775E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0775_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0800E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0800_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0825E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0825_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0850E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0850_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0875E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0875_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0900E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0900_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0925E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0925_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0950E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0950_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0975E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0975_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1000E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1000_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1025E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1025_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1050E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1050_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1075E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1075_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1100E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1100_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1125E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1125_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1150E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1150_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1175E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1175_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1200E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1200_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1225E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1225_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1250E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1250_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1275E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1275_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1300E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1300_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1325E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1325_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1350E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1350_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1375E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1375_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1400E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1400_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1425E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1425_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1450E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1450_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1475E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1475_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1500E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1500_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1525E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1525_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1550E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1550_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1575E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1575_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1600E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1600_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1625E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1625_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1650E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1650_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1700E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1700_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1725E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1725_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1750E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1750_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1800E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1800_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1850E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1850_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1875E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1875_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1900E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1900_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1925E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1925_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1950E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1950_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2000E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2000_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2025E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2025_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2050E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2050_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2100E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2100_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2125E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2125_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2150E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2150_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2175E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2175_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2200E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2200_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2250E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2250_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2275E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2275_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2300E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2300_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2325E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2325_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2350E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2350_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2375E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2375_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2400E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2400_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2450E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2450_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2475E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2475_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2500E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2500_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2550E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2550_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2600E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2600_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2625E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2625_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2650E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2650_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2700E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2700_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2750E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2750_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2775E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2775_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2800E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2800_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2850E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2850_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2875E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2875_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2900E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2900_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2925E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2925_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2950E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2950_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2975E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2975_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3000E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3000_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3025E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3025_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3050E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3050_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3075E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3075_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3100E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3100_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3125E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3125_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3150E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3150_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3200E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3200_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3225E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3225_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3250E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3250_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3300E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3300_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3325E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3325_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3375E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3375_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3400E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3400_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3450E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3450_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3500E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3500_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3525E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3525_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3575E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3575_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3600E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3600_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3625E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3625_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3675E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3675_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3700E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3700_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3750E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3750_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3800E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3800_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3825E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3825_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3850E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3850_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3875E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3875_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3900E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3900_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3975E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3975_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4000E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4000_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4025E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4025_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4050E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4050_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4100E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4100_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4125E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4125_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4200E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4200_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4225E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4225_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4250E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4250_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4275E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4275_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4300E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4300_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4350E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4350_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4375E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4375_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4400E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4400_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4425E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4425_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4500E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4500_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4550E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4550_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4575E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4575_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4600E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4600_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4625E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4625_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4650E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4650_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4675E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4675_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4700E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4700_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4725E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4725_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4750E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4750_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4800E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4800_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4875E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4875_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4900E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4900_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4950E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4950_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed5000E450g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED5000_E4_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeDynamicE550g                    CreateInternalVnicAttachmentDetailsVnicShapeEnum = "DYNAMIC_E5_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeDynamicE5100g                   CreateInternalVnicAttachmentDetailsVnicShapeEnum = "DYNAMIC_E5_100G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0020A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0020_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0040A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0040_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0060A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0060_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0080A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0080_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0100A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0100_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0120A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0120_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0140A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0140_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0160A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0160_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0180A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0180_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0200A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0200_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0220A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0220_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0240A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0240_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0260A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0260_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0280A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0280_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0300A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0300_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0320A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0320_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0340A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0340_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0360A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0360_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0380A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0380_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0400A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0400_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0420A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0420_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0440A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0440_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0460A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0460_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0480A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0480_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0500A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0500_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0520A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0520_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0540A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0540_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0560A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0560_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0580A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0580_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0600A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0600_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0620A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0620_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0640A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0640_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0660A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0660_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0680A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0680_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0700A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0700_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0720A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0720_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0740A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0740_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0760A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0760_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0780A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0780_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0800A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0800_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0820A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0820_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0840A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0840_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0860A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0860_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0880A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0880_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0900A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0900_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0920A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0920_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0940A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0940_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0960A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0960_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0980A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0980_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1000A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1000_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1020A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1020_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1040A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1040_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1060A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1060_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1080A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1080_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1100A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1100_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1120A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1120_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1140A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1140_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1160A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1160_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1180A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1180_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1200A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1200_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1220A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1220_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1240A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1240_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1260A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1260_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1280A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1280_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1300A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1300_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1320A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1320_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1340A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1340_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1360A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1360_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1380A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1380_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1400A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1400_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1420A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1420_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1440A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1440_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1460A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1460_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1480A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1480_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1500A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1500_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1520A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1520_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1540A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1540_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1560A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1560_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1580A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1580_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1600A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1600_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1620A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1620_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1640A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1640_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1660A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1660_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1680A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1680_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1700A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1700_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1720A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1720_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1740A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1740_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1760A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1760_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1780A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1780_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1800A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1800_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1820A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1820_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1840A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1840_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1860A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1860_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1880A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1880_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1900A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1900_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1920A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1920_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1940A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1940_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1960A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1960_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1980A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1980_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2000A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2000_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2020A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2020_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2040A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2040_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2060A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2060_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2080A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2080_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2100A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2100_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2120A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2120_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2140A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2140_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2160A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2160_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2180A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2180_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2200A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2200_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2220A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2220_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2240A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2240_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2260A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2260_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2280A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2280_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2300A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2300_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2320A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2320_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2340A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2340_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2360A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2360_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2380A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2380_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2400A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2400_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2420A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2420_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2440A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2440_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2460A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2460_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2480A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2480_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2500A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2500_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2520A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2520_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2540A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2540_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2560A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2560_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2580A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2580_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2600A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2600_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2620A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2620_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2640A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2640_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2660A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2660_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2680A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2680_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2700A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2700_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2720A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2720_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2740A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2740_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2760A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2760_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2780A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2780_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2800A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2800_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2820A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2820_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2840A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2840_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2860A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2860_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2880A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2880_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2900A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2900_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2920A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2920_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2940A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2940_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2960A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2960_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2980A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2980_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3000A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3000_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3020A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3020_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3040A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3040_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3060A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3060_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3080A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3080_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3100A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3100_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3120A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3120_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3140A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3140_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3160A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3160_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3180A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3180_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3200A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3200_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3220A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3220_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3240A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3240_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3260A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3260_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3280A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3280_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3300A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3300_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3320A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3320_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3340A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3340_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3360A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3360_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3380A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3380_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3400A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3400_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3420A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3420_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3440A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3440_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3460A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3460_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3480A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3480_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3500A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3500_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3520A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3520_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3540A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3540_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3560A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3560_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3580A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3580_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3600A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3600_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3620A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3620_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3640A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3640_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3660A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3660_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3680A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3680_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3700A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3700_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3720A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3720_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3740A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3740_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3760A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3760_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3780A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3780_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3800A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3800_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3820A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3820_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3840A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3840_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3860A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3860_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3880A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3880_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3900A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3900_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3920A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3920_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3940A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3940_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3960A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3960_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3980A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3980_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4000A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4000_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4020A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4020_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4040A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4040_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4060A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4060_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4080A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4080_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4100A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4100_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4120A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4120_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4140A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4140_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4160A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4160_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4180A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4180_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4200A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4200_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4220A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4220_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4240A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4240_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4260A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4260_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4280A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4280_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4300A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4300_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4320A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4320_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4340A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4340_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4360A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4360_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4380A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4380_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4400A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4400_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4420A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4420_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4440A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4440_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4460A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4460_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4480A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4480_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4500A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4500_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4520A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4520_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4540A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4540_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4560A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4560_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4580A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4580_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4600A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4600_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4620A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4620_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4640A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4640_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4660A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4660_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4680A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4680_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4700A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4700_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4720A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4720_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4740A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4740_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4760A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4760_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4780A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4780_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4800A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4800_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4820A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4820_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4840A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4840_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4860A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4860_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4880A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4880_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4900A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4900_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4920A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4920_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4940A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4940_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4960A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4960_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4980A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4980_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed5000A150g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED5000_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0090X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0090_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0180X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0180_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0270X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0270_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0360X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0360_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0450X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0450_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0540X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0540_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0630X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0630_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0720X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0720_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0810X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0810_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0900X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0900_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0990X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0990_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1080X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1080_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1170X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1170_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1260X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1260_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1350X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1350_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1440X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1440_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1530X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1530_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1620X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1620_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1710X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1710_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1800X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1800_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1890X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1890_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1980X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1980_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2070X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2070_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2160X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2160_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2250X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2250_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2340X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2340_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2430X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2430_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2520X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2520_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2610X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2610_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2700X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2700_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2790X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2790_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2880X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2880_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2970X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2970_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3060X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3060_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3150X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3150_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3240X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3240_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3330X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3330_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3420X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3420_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3510X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3510_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3600X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3600_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3690X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3690_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3780X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3780_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3870X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3870_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3960X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3960_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4050X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4050_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4140X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4140_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4230X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4230_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4320X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4320_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4410X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4410_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4500X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4500_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4590X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4590_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4680X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4680_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4770X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4770_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4860X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4860_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4950X950g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4950_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeDynamicA150g                    CreateInternalVnicAttachmentDetailsVnicShapeEnum = "DYNAMIC_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0040A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0040_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0100A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0100_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0200A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0200_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0300A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0300_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0400A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0400_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0500A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0500_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0600A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0600_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0700A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0700_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0800A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0800_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0900A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0900_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1000A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1000_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1100A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1100_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1200A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1200_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1300A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1300_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1400A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1400_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1500A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1500_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1600A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1600_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1700A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1700_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1800A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1800_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1900A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1900_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2000A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2000_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2100A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2100_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2200A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2200_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2300A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2300_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2400A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2400_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2500A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2500_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2600A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2600_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2700A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2700_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2800A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2800_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2900A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2900_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3000A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3000_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3100A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3100_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3200A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3200_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3300A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3300_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3400A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3400_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3500A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3500_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3600A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3600_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3700A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3700_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3800A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3800_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3900A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3900_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed4000A150g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED4000_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed5000TelesisA150g           CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED5000_TELESIS_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeEntirehostA150g                 CreateInternalVnicAttachmentDetailsVnicShapeEnum = "ENTIREHOST_A1_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeDynamicX950g                    CreateInternalVnicAttachmentDetailsVnicShapeEnum = "DYNAMIC_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0040X950g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0040_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0400X950g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0400_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0800X950g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0800_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1200X950g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1200_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1600X950g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1600_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2000X950g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2000_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2400X950g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2400_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2800X950g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2800_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3200X950g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3200_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3600X950g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3600_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed4000X950g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED4000_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0100X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED0100_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0200X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED0200_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0300X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED0300_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0400X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED0400_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0500X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED0500_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0600X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED0600_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0700X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED0700_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0800X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED0800_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0900X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED0900_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1000X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED1000_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1100X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED1100_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1200X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED1200_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1300X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED1300_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1400X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED1400_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1500X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED1500_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1600X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED1600_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1700X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED1700_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1800X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED1800_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1900X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED1900_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2000X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED2000_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2100X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED2100_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2200X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED2200_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2300X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED2300_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2400X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED2400_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2500X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED2500_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2600X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED2600_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2700X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED2700_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2800X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED2800_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2900X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED2900_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3000X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED3000_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3100X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED3100_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3200X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED3200_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3300X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED3300_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3400X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED3400_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3500X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED3500_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3600X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED3600_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3700X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED3700_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3800X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED3800_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3900X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED3900_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4000X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED4000_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4100X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED4100_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4200X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED4200_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4300X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED4300_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4400X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED4400_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4500X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED4500_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4600X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED4600_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4700X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED4700_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4800X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED4800_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4900X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED4900_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed5000X950g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "STANDARD_VM_FIXED5000_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0025X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0025_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0050X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0050_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0075X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0075_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0100X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0100_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0125X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0125_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0150X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0150_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0175X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0175_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0200X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0200_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0225X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0225_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0250X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0250_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0275X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0275_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0300X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0300_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0325X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0325_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0350X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0350_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0375X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0375_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0400X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0400_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0425X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0425_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0450X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0450_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0475X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0475_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0500X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0500_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0525X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0525_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0550X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0550_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0575X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0575_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0600X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0600_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0625X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0625_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0650X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0650_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0675X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0675_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0700X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0700_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0725X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0725_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0750X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0750_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0775X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0775_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0800X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0800_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0825X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0825_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0850X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0850_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0875X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0875_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0900X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0900_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0925X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0925_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0950X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0950_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0975X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0975_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1000X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1000_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1025X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1025_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1050X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1050_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1075X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1075_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1100X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1100_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1125X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1125_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1150X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1150_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1175X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1175_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1200X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1200_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1225X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1225_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1250X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1250_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1275X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1275_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1300X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1300_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1325X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1325_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1350X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1350_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1375X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1375_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1400X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1400_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1425X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1425_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1450X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1450_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1475X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1475_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1500X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1500_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1525X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1525_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1550X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1550_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1575X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1575_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1600X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1600_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1625X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1625_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1650X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1650_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1700X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1700_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1725X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1725_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1750X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1750_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1800X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1800_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1850X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1850_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1875X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1875_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1900X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1900_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1925X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1925_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1950X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1950_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2000X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2000_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2025X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2025_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2050X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2050_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2100X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2100_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2125X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2125_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2150X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2150_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2175X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2175_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2200X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2200_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2250X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2250_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2275X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2275_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2300X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2300_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2325X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2325_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2350X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2350_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2375X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2375_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2400X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2400_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2450X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2450_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2475X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2475_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2500X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2500_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2550X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2550_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2600X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2600_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2625X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2625_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2650X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2650_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2700X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2700_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2750X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2750_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2775X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2775_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2800X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2800_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2850X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2850_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2875X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2875_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2900X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2900_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2925X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2925_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2950X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2950_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2975X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2975_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3000X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3000_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3025X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3025_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3050X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3050_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3075X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3075_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3100X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3100_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3125X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3125_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3150X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3150_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3200X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3200_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3225X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3225_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3250X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3250_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3300X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3300_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3325X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3325_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3375X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3375_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3400X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3400_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3450X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3450_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3500X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3500_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3525X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3525_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3575X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3575_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3600X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3600_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3625X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3625_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3675X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3675_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3700X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3700_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3750X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3750_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3800X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3800_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3825X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3825_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3850X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3850_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3875X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3875_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3900X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3900_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3975X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3975_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4000X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4000_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4025X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4025_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4050X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4050_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4100X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4100_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4125X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4125_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4200X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4200_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4225X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4225_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4250X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4250_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4275X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4275_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4300X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4300_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4350X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4350_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4375X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4375_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4400X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4400_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4425X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4425_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4500X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4500_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4550X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4550_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4575X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4575_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4600X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4600_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4625X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4625_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4650X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4650_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4675X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4675_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4700X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4700_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4725X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4725_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4750X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4750_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4800X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4800_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4875X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4875_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4900X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4900_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4950X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4950_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed5000X950g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED5000_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeEntirehostX950g                 CreateInternalVnicAttachmentDetailsVnicShapeEnum = "ENTIREHOST_X9_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeDynamicX9100g                   CreateInternalVnicAttachmentDetailsVnicShapeEnum = "DYNAMIC_X9_100G"
	CreateInternalVnicAttachmentDetailsVnicShapeDynamicX1050g                   CreateInternalVnicAttachmentDetailsVnicShapeEnum = "DYNAMIC_X10_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeDynamicX10100g                  CreateInternalVnicAttachmentDetailsVnicShapeEnum = "DYNAMIC_X10_100G"
)

var mappingCreateInternalVnicAttachmentDetailsVnicShapeEnum = map[string]CreateInternalVnicAttachmentDetailsVnicShapeEnum{
	"DYNAMIC":                              CreateInternalVnicAttachmentDetailsVnicShapeDynamic,
	"FIXED0040":                            CreateInternalVnicAttachmentDetailsVnicShapeFixed0040,
	"FIXED0060":                            CreateInternalVnicAttachmentDetailsVnicShapeFixed0060,
	"FIXED0060_PSM":                        CreateInternalVnicAttachmentDetailsVnicShapeFixed0060Psm,
	"FIXED0100":                            CreateInternalVnicAttachmentDetailsVnicShapeFixed0100,
	"FIXED0120":                            CreateInternalVnicAttachmentDetailsVnicShapeFixed0120,
	"FIXED0120_2X":                         CreateInternalVnicAttachmentDetailsVnicShapeFixed01202x,
	"FIXED0200":                            CreateInternalVnicAttachmentDetailsVnicShapeFixed0200,
	"FIXED0240":                            CreateInternalVnicAttachmentDetailsVnicShapeFixed0240,
	"FIXED0480":                            CreateInternalVnicAttachmentDetailsVnicShapeFixed0480,
	"ENTIREHOST":                           CreateInternalVnicAttachmentDetailsVnicShapeEntirehost,
	"DYNAMIC_25G":                          CreateInternalVnicAttachmentDetailsVnicShapeDynamic25g,
	"FIXED0040_25G":                        CreateInternalVnicAttachmentDetailsVnicShapeFixed004025g,
	"FIXED0100_25G":                        CreateInternalVnicAttachmentDetailsVnicShapeFixed010025g,
	"FIXED0200_25G":                        CreateInternalVnicAttachmentDetailsVnicShapeFixed020025g,
	"FIXED0400_25G":                        CreateInternalVnicAttachmentDetailsVnicShapeFixed040025g,
	"FIXED0800_25G":                        CreateInternalVnicAttachmentDetailsVnicShapeFixed080025g,
	"FIXED1600_25G":                        CreateInternalVnicAttachmentDetailsVnicShapeFixed160025g,
	"FIXED2400_25G":                        CreateInternalVnicAttachmentDetailsVnicShapeFixed240025g,
	"ENTIREHOST_25G":                       CreateInternalVnicAttachmentDetailsVnicShapeEntirehost25g,
	"DYNAMIC_E1_25G":                       CreateInternalVnicAttachmentDetailsVnicShapeDynamicE125g,
	"FIXED0040_E1_25G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0040E125g,
	"FIXED0070_E1_25G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0070E125g,
	"FIXED0140_E1_25G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0140E125g,
	"FIXED0280_E1_25G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0280E125g,
	"FIXED0560_E1_25G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0560E125g,
	"FIXED1120_E1_25G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1120E125g,
	"FIXED1680_E1_25G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1680E125g,
	"ENTIREHOST_E1_25G":                    CreateInternalVnicAttachmentDetailsVnicShapeEntirehostE125g,
	"DYNAMIC_B1_25G":                       CreateInternalVnicAttachmentDetailsVnicShapeDynamicB125g,
	"FIXED0040_B1_25G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0040B125g,
	"FIXED0060_B1_25G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0060B125g,
	"FIXED0120_B1_25G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0120B125g,
	"FIXED0240_B1_25G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0240B125g,
	"FIXED0480_B1_25G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0480B125g,
	"FIXED0960_B1_25G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0960B125g,
	"ENTIREHOST_B1_25G":                    CreateInternalVnicAttachmentDetailsVnicShapeEntirehostB125g,
	"MICRO_VM_FIXED0048_E1_25G":            CreateInternalVnicAttachmentDetailsVnicShapeMicroVmFixed0048E125g,
	"MICRO_LB_FIXED0001_E1_25G":            CreateInternalVnicAttachmentDetailsVnicShapeMicroLbFixed0001E125g,
	"VNICAAS_FIXED0200":                    CreateInternalVnicAttachmentDetailsVnicShapeVnicaasFixed0200,
	"VNICAAS_FIXED0400":                    CreateInternalVnicAttachmentDetailsVnicShapeVnicaasFixed0400,
	"VNICAAS_FIXED0700":                    CreateInternalVnicAttachmentDetailsVnicShapeVnicaasFixed0700,
	"VNICAAS_NLB_APPROVED_10G":             CreateInternalVnicAttachmentDetailsVnicShapeVnicaasNlbApproved10g,
	"VNICAAS_NLB_APPROVED_25G":             CreateInternalVnicAttachmentDetailsVnicShapeVnicaasNlbApproved25g,
	"VNICAAS_TELESIS_25G":                  CreateInternalVnicAttachmentDetailsVnicShapeVnicaasTelesis25g,
	"VNICAAS_TELESIS_10G":                  CreateInternalVnicAttachmentDetailsVnicShapeVnicaasTelesis10g,
	"VNICAAS_AMBASSADOR_FIXED0100":         CreateInternalVnicAttachmentDetailsVnicShapeVnicaasAmbassadorFixed0100,
	"VNICAAS_TELESIS_GAMMA":                CreateInternalVnicAttachmentDetailsVnicShapeVnicaasTelesisGamma,
	"VNICAAS_PRIVATEDNS":                   CreateInternalVnicAttachmentDetailsVnicShapeVnicaasPrivatedns,
	"VNICAAS_FWAAS":                        CreateInternalVnicAttachmentDetailsVnicShapeVnicaasFwaas,
	"VNICAAS_LBAAS_FREE":                   CreateInternalVnicAttachmentDetailsVnicShapeVnicaasLbaasFree,
	"VNICAAS_LBAAS_8G_512K":                CreateInternalVnicAttachmentDetailsVnicShapeVnicaasLbaas8g512k,
	"VNICAAS_LBAAS_8G_1M":                  CreateInternalVnicAttachmentDetailsVnicShapeVnicaasLbaas8g1m,
	"VNICAAS_LBAAS_8G_2M":                  CreateInternalVnicAttachmentDetailsVnicShapeVnicaasLbaas8g2m,
	"VNICAAS_LBAAS_8G_3M":                  CreateInternalVnicAttachmentDetailsVnicShapeVnicaasLbaas8g3m,
	"VNICAAS_LBAAS_8G_1M_8GHOST":           CreateInternalVnicAttachmentDetailsVnicShapeVnicaasLbaas8g1m8ghost,
	"VNICAAS_LBAAS_8G_1M_16GHOST":          CreateInternalVnicAttachmentDetailsVnicShapeVnicaasLbaas8g1m16ghost,
	"DYNAMIC_E3_50G":                       CreateInternalVnicAttachmentDetailsVnicShapeDynamicE350g,
	"FIXED0040_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0040E350g,
	"FIXED0100_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0100E350g,
	"FIXED0200_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0200E350g,
	"FIXED0300_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0300E350g,
	"FIXED0400_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0400E350g,
	"FIXED0500_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0500E350g,
	"FIXED0600_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0600E350g,
	"FIXED0700_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0700E350g,
	"FIXED0800_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0800E350g,
	"FIXED0900_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0900E350g,
	"FIXED1000_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1000E350g,
	"FIXED1100_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1100E350g,
	"FIXED1200_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1200E350g,
	"FIXED1300_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1300E350g,
	"FIXED1400_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1400E350g,
	"FIXED1500_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1500E350g,
	"FIXED1600_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1600E350g,
	"FIXED1700_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1700E350g,
	"FIXED1800_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1800E350g,
	"FIXED1900_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1900E350g,
	"FIXED2000_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2000E350g,
	"FIXED2100_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2100E350g,
	"FIXED2200_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2200E350g,
	"FIXED2300_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2300E350g,
	"FIXED2400_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2400E350g,
	"FIXED2500_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2500E350g,
	"FIXED2600_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2600E350g,
	"FIXED2700_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2700E350g,
	"FIXED2800_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2800E350g,
	"FIXED2900_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2900E350g,
	"FIXED3000_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3000E350g,
	"FIXED3100_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3100E350g,
	"FIXED3200_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3200E350g,
	"FIXED3300_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3300E350g,
	"FIXED3400_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3400E350g,
	"FIXED3500_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3500E350g,
	"FIXED3600_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3600E350g,
	"FIXED3700_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3700E350g,
	"FIXED3800_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3800E350g,
	"FIXED3900_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3900E350g,
	"FIXED4000_E3_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed4000E350g,
	"ENTIREHOST_E3_50G":                    CreateInternalVnicAttachmentDetailsVnicShapeEntirehostE350g,
	"DYNAMIC_E4_50G":                       CreateInternalVnicAttachmentDetailsVnicShapeDynamicE450g,
	"FIXED0040_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0040E450g,
	"FIXED0100_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0100E450g,
	"FIXED0200_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0200E450g,
	"FIXED0300_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0300E450g,
	"FIXED0400_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0400E450g,
	"FIXED0500_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0500E450g,
	"FIXED0600_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0600E450g,
	"FIXED0700_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0700E450g,
	"FIXED0800_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0800E450g,
	"FIXED0900_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0900E450g,
	"FIXED1000_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1000E450g,
	"FIXED1100_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1100E450g,
	"FIXED1200_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1200E450g,
	"FIXED1300_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1300E450g,
	"FIXED1400_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1400E450g,
	"FIXED1500_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1500E450g,
	"FIXED1600_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1600E450g,
	"FIXED1700_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1700E450g,
	"FIXED1800_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1800E450g,
	"FIXED1900_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1900E450g,
	"FIXED2000_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2000E450g,
	"FIXED2100_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2100E450g,
	"FIXED2200_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2200E450g,
	"FIXED2300_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2300E450g,
	"FIXED2400_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2400E450g,
	"FIXED2500_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2500E450g,
	"FIXED2600_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2600E450g,
	"FIXED2700_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2700E450g,
	"FIXED2800_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2800E450g,
	"FIXED2900_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2900E450g,
	"FIXED3000_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3000E450g,
	"FIXED3100_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3100E450g,
	"FIXED3200_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3200E450g,
	"FIXED3300_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3300E450g,
	"FIXED3400_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3400E450g,
	"FIXED3500_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3500E450g,
	"FIXED3600_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3600E450g,
	"FIXED3700_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3700E450g,
	"FIXED3800_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3800E450g,
	"FIXED3900_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3900E450g,
	"FIXED4000_E4_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed4000E450g,
	"ENTIREHOST_E4_50G":                    CreateInternalVnicAttachmentDetailsVnicShapeEntirehostE450g,
	"Micro_VM_Fixed0050_E3_50G":            CreateInternalVnicAttachmentDetailsVnicShapeMicroVmFixed0050E350g,
	"Micro_VM_Fixed0050_E4_50G":            CreateInternalVnicAttachmentDetailsVnicShapeMicroVmFixed0050E450g,
	"SUBCORE_VM_FIXED0025_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0025E350g,
	"SUBCORE_VM_FIXED0050_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0050E350g,
	"SUBCORE_VM_FIXED0075_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0075E350g,
	"SUBCORE_VM_FIXED0100_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0100E350g,
	"SUBCORE_VM_FIXED0125_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0125E350g,
	"SUBCORE_VM_FIXED0150_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0150E350g,
	"SUBCORE_VM_FIXED0175_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0175E350g,
	"SUBCORE_VM_FIXED0200_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0200E350g,
	"SUBCORE_VM_FIXED0225_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0225E350g,
	"SUBCORE_VM_FIXED0250_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0250E350g,
	"SUBCORE_VM_FIXED0275_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0275E350g,
	"SUBCORE_VM_FIXED0300_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0300E350g,
	"SUBCORE_VM_FIXED0325_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0325E350g,
	"SUBCORE_VM_FIXED0350_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0350E350g,
	"SUBCORE_VM_FIXED0375_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0375E350g,
	"SUBCORE_VM_FIXED0400_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0400E350g,
	"SUBCORE_VM_FIXED0425_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0425E350g,
	"SUBCORE_VM_FIXED0450_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0450E350g,
	"SUBCORE_VM_FIXED0475_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0475E350g,
	"SUBCORE_VM_FIXED0500_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0500E350g,
	"SUBCORE_VM_FIXED0525_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0525E350g,
	"SUBCORE_VM_FIXED0550_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0550E350g,
	"SUBCORE_VM_FIXED0575_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0575E350g,
	"SUBCORE_VM_FIXED0600_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0600E350g,
	"SUBCORE_VM_FIXED0625_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0625E350g,
	"SUBCORE_VM_FIXED0650_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0650E350g,
	"SUBCORE_VM_FIXED0675_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0675E350g,
	"SUBCORE_VM_FIXED0700_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0700E350g,
	"SUBCORE_VM_FIXED0725_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0725E350g,
	"SUBCORE_VM_FIXED0750_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0750E350g,
	"SUBCORE_VM_FIXED0775_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0775E350g,
	"SUBCORE_VM_FIXED0800_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0800E350g,
	"SUBCORE_VM_FIXED0825_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0825E350g,
	"SUBCORE_VM_FIXED0850_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0850E350g,
	"SUBCORE_VM_FIXED0875_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0875E350g,
	"SUBCORE_VM_FIXED0900_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0900E350g,
	"SUBCORE_VM_FIXED0925_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0925E350g,
	"SUBCORE_VM_FIXED0950_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0950E350g,
	"SUBCORE_VM_FIXED0975_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0975E350g,
	"SUBCORE_VM_FIXED1000_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1000E350g,
	"SUBCORE_VM_FIXED1025_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1025E350g,
	"SUBCORE_VM_FIXED1050_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1050E350g,
	"SUBCORE_VM_FIXED1075_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1075E350g,
	"SUBCORE_VM_FIXED1100_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1100E350g,
	"SUBCORE_VM_FIXED1125_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1125E350g,
	"SUBCORE_VM_FIXED1150_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1150E350g,
	"SUBCORE_VM_FIXED1175_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1175E350g,
	"SUBCORE_VM_FIXED1200_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1200E350g,
	"SUBCORE_VM_FIXED1225_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1225E350g,
	"SUBCORE_VM_FIXED1250_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1250E350g,
	"SUBCORE_VM_FIXED1275_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1275E350g,
	"SUBCORE_VM_FIXED1300_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1300E350g,
	"SUBCORE_VM_FIXED1325_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1325E350g,
	"SUBCORE_VM_FIXED1350_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1350E350g,
	"SUBCORE_VM_FIXED1375_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1375E350g,
	"SUBCORE_VM_FIXED1400_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1400E350g,
	"SUBCORE_VM_FIXED1425_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1425E350g,
	"SUBCORE_VM_FIXED1450_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1450E350g,
	"SUBCORE_VM_FIXED1475_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1475E350g,
	"SUBCORE_VM_FIXED1500_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1500E350g,
	"SUBCORE_VM_FIXED1525_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1525E350g,
	"SUBCORE_VM_FIXED1550_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1550E350g,
	"SUBCORE_VM_FIXED1575_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1575E350g,
	"SUBCORE_VM_FIXED1600_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1600E350g,
	"SUBCORE_VM_FIXED1625_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1625E350g,
	"SUBCORE_VM_FIXED1650_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1650E350g,
	"SUBCORE_VM_FIXED1700_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1700E350g,
	"SUBCORE_VM_FIXED1725_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1725E350g,
	"SUBCORE_VM_FIXED1750_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1750E350g,
	"SUBCORE_VM_FIXED1800_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1800E350g,
	"SUBCORE_VM_FIXED1850_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1850E350g,
	"SUBCORE_VM_FIXED1875_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1875E350g,
	"SUBCORE_VM_FIXED1900_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1900E350g,
	"SUBCORE_VM_FIXED1925_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1925E350g,
	"SUBCORE_VM_FIXED1950_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1950E350g,
	"SUBCORE_VM_FIXED2000_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2000E350g,
	"SUBCORE_VM_FIXED2025_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2025E350g,
	"SUBCORE_VM_FIXED2050_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2050E350g,
	"SUBCORE_VM_FIXED2100_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2100E350g,
	"SUBCORE_VM_FIXED2125_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2125E350g,
	"SUBCORE_VM_FIXED2150_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2150E350g,
	"SUBCORE_VM_FIXED2175_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2175E350g,
	"SUBCORE_VM_FIXED2200_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2200E350g,
	"SUBCORE_VM_FIXED2250_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2250E350g,
	"SUBCORE_VM_FIXED2275_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2275E350g,
	"SUBCORE_VM_FIXED2300_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2300E350g,
	"SUBCORE_VM_FIXED2325_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2325E350g,
	"SUBCORE_VM_FIXED2350_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2350E350g,
	"SUBCORE_VM_FIXED2375_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2375E350g,
	"SUBCORE_VM_FIXED2400_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2400E350g,
	"SUBCORE_VM_FIXED2450_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2450E350g,
	"SUBCORE_VM_FIXED2475_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2475E350g,
	"SUBCORE_VM_FIXED2500_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2500E350g,
	"SUBCORE_VM_FIXED2550_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2550E350g,
	"SUBCORE_VM_FIXED2600_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2600E350g,
	"SUBCORE_VM_FIXED2625_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2625E350g,
	"SUBCORE_VM_FIXED2650_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2650E350g,
	"SUBCORE_VM_FIXED2700_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2700E350g,
	"SUBCORE_VM_FIXED2750_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2750E350g,
	"SUBCORE_VM_FIXED2775_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2775E350g,
	"SUBCORE_VM_FIXED2800_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2800E350g,
	"SUBCORE_VM_FIXED2850_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2850E350g,
	"SUBCORE_VM_FIXED2875_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2875E350g,
	"SUBCORE_VM_FIXED2900_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2900E350g,
	"SUBCORE_VM_FIXED2925_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2925E350g,
	"SUBCORE_VM_FIXED2950_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2950E350g,
	"SUBCORE_VM_FIXED2975_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2975E350g,
	"SUBCORE_VM_FIXED3000_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3000E350g,
	"SUBCORE_VM_FIXED3025_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3025E350g,
	"SUBCORE_VM_FIXED3050_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3050E350g,
	"SUBCORE_VM_FIXED3075_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3075E350g,
	"SUBCORE_VM_FIXED3100_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3100E350g,
	"SUBCORE_VM_FIXED3125_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3125E350g,
	"SUBCORE_VM_FIXED3150_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3150E350g,
	"SUBCORE_VM_FIXED3200_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3200E350g,
	"SUBCORE_VM_FIXED3225_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3225E350g,
	"SUBCORE_VM_FIXED3250_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3250E350g,
	"SUBCORE_VM_FIXED3300_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3300E350g,
	"SUBCORE_VM_FIXED3325_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3325E350g,
	"SUBCORE_VM_FIXED3375_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3375E350g,
	"SUBCORE_VM_FIXED3400_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3400E350g,
	"SUBCORE_VM_FIXED3450_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3450E350g,
	"SUBCORE_VM_FIXED3500_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3500E350g,
	"SUBCORE_VM_FIXED3525_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3525E350g,
	"SUBCORE_VM_FIXED3575_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3575E350g,
	"SUBCORE_VM_FIXED3600_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3600E350g,
	"SUBCORE_VM_FIXED3625_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3625E350g,
	"SUBCORE_VM_FIXED3675_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3675E350g,
	"SUBCORE_VM_FIXED3700_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3700E350g,
	"SUBCORE_VM_FIXED3750_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3750E350g,
	"SUBCORE_VM_FIXED3800_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3800E350g,
	"SUBCORE_VM_FIXED3825_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3825E350g,
	"SUBCORE_VM_FIXED3850_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3850E350g,
	"SUBCORE_VM_FIXED3875_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3875E350g,
	"SUBCORE_VM_FIXED3900_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3900E350g,
	"SUBCORE_VM_FIXED3975_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3975E350g,
	"SUBCORE_VM_FIXED4000_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4000E350g,
	"SUBCORE_VM_FIXED4025_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4025E350g,
	"SUBCORE_VM_FIXED4050_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4050E350g,
	"SUBCORE_VM_FIXED4100_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4100E350g,
	"SUBCORE_VM_FIXED4125_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4125E350g,
	"SUBCORE_VM_FIXED4200_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4200E350g,
	"SUBCORE_VM_FIXED4225_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4225E350g,
	"SUBCORE_VM_FIXED4250_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4250E350g,
	"SUBCORE_VM_FIXED4275_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4275E350g,
	"SUBCORE_VM_FIXED4300_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4300E350g,
	"SUBCORE_VM_FIXED4350_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4350E350g,
	"SUBCORE_VM_FIXED4375_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4375E350g,
	"SUBCORE_VM_FIXED4400_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4400E350g,
	"SUBCORE_VM_FIXED4425_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4425E350g,
	"SUBCORE_VM_FIXED4500_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4500E350g,
	"SUBCORE_VM_FIXED4550_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4550E350g,
	"SUBCORE_VM_FIXED4575_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4575E350g,
	"SUBCORE_VM_FIXED4600_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4600E350g,
	"SUBCORE_VM_FIXED4625_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4625E350g,
	"SUBCORE_VM_FIXED4650_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4650E350g,
	"SUBCORE_VM_FIXED4675_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4675E350g,
	"SUBCORE_VM_FIXED4700_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4700E350g,
	"SUBCORE_VM_FIXED4725_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4725E350g,
	"SUBCORE_VM_FIXED4750_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4750E350g,
	"SUBCORE_VM_FIXED4800_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4800E350g,
	"SUBCORE_VM_FIXED4875_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4875E350g,
	"SUBCORE_VM_FIXED4900_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4900E350g,
	"SUBCORE_VM_FIXED4950_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4950E350g,
	"SUBCORE_VM_FIXED5000_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed5000E350g,
	"SUBCORE_VM_FIXED0025_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0025E450g,
	"SUBCORE_VM_FIXED0050_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0050E450g,
	"SUBCORE_VM_FIXED0075_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0075E450g,
	"SUBCORE_VM_FIXED0100_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0100E450g,
	"SUBCORE_VM_FIXED0125_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0125E450g,
	"SUBCORE_VM_FIXED0150_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0150E450g,
	"SUBCORE_VM_FIXED0175_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0175E450g,
	"SUBCORE_VM_FIXED0200_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0200E450g,
	"SUBCORE_VM_FIXED0225_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0225E450g,
	"SUBCORE_VM_FIXED0250_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0250E450g,
	"SUBCORE_VM_FIXED0275_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0275E450g,
	"SUBCORE_VM_FIXED0300_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0300E450g,
	"SUBCORE_VM_FIXED0325_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0325E450g,
	"SUBCORE_VM_FIXED0350_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0350E450g,
	"SUBCORE_VM_FIXED0375_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0375E450g,
	"SUBCORE_VM_FIXED0400_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0400E450g,
	"SUBCORE_VM_FIXED0425_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0425E450g,
	"SUBCORE_VM_FIXED0450_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0450E450g,
	"SUBCORE_VM_FIXED0475_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0475E450g,
	"SUBCORE_VM_FIXED0500_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0500E450g,
	"SUBCORE_VM_FIXED0525_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0525E450g,
	"SUBCORE_VM_FIXED0550_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0550E450g,
	"SUBCORE_VM_FIXED0575_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0575E450g,
	"SUBCORE_VM_FIXED0600_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0600E450g,
	"SUBCORE_VM_FIXED0625_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0625E450g,
	"SUBCORE_VM_FIXED0650_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0650E450g,
	"SUBCORE_VM_FIXED0675_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0675E450g,
	"SUBCORE_VM_FIXED0700_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0700E450g,
	"SUBCORE_VM_FIXED0725_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0725E450g,
	"SUBCORE_VM_FIXED0750_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0750E450g,
	"SUBCORE_VM_FIXED0775_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0775E450g,
	"SUBCORE_VM_FIXED0800_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0800E450g,
	"SUBCORE_VM_FIXED0825_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0825E450g,
	"SUBCORE_VM_FIXED0850_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0850E450g,
	"SUBCORE_VM_FIXED0875_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0875E450g,
	"SUBCORE_VM_FIXED0900_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0900E450g,
	"SUBCORE_VM_FIXED0925_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0925E450g,
	"SUBCORE_VM_FIXED0950_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0950E450g,
	"SUBCORE_VM_FIXED0975_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0975E450g,
	"SUBCORE_VM_FIXED1000_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1000E450g,
	"SUBCORE_VM_FIXED1025_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1025E450g,
	"SUBCORE_VM_FIXED1050_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1050E450g,
	"SUBCORE_VM_FIXED1075_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1075E450g,
	"SUBCORE_VM_FIXED1100_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1100E450g,
	"SUBCORE_VM_FIXED1125_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1125E450g,
	"SUBCORE_VM_FIXED1150_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1150E450g,
	"SUBCORE_VM_FIXED1175_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1175E450g,
	"SUBCORE_VM_FIXED1200_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1200E450g,
	"SUBCORE_VM_FIXED1225_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1225E450g,
	"SUBCORE_VM_FIXED1250_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1250E450g,
	"SUBCORE_VM_FIXED1275_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1275E450g,
	"SUBCORE_VM_FIXED1300_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1300E450g,
	"SUBCORE_VM_FIXED1325_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1325E450g,
	"SUBCORE_VM_FIXED1350_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1350E450g,
	"SUBCORE_VM_FIXED1375_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1375E450g,
	"SUBCORE_VM_FIXED1400_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1400E450g,
	"SUBCORE_VM_FIXED1425_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1425E450g,
	"SUBCORE_VM_FIXED1450_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1450E450g,
	"SUBCORE_VM_FIXED1475_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1475E450g,
	"SUBCORE_VM_FIXED1500_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1500E450g,
	"SUBCORE_VM_FIXED1525_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1525E450g,
	"SUBCORE_VM_FIXED1550_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1550E450g,
	"SUBCORE_VM_FIXED1575_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1575E450g,
	"SUBCORE_VM_FIXED1600_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1600E450g,
	"SUBCORE_VM_FIXED1625_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1625E450g,
	"SUBCORE_VM_FIXED1650_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1650E450g,
	"SUBCORE_VM_FIXED1700_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1700E450g,
	"SUBCORE_VM_FIXED1725_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1725E450g,
	"SUBCORE_VM_FIXED1750_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1750E450g,
	"SUBCORE_VM_FIXED1800_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1800E450g,
	"SUBCORE_VM_FIXED1850_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1850E450g,
	"SUBCORE_VM_FIXED1875_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1875E450g,
	"SUBCORE_VM_FIXED1900_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1900E450g,
	"SUBCORE_VM_FIXED1925_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1925E450g,
	"SUBCORE_VM_FIXED1950_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1950E450g,
	"SUBCORE_VM_FIXED2000_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2000E450g,
	"SUBCORE_VM_FIXED2025_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2025E450g,
	"SUBCORE_VM_FIXED2050_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2050E450g,
	"SUBCORE_VM_FIXED2100_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2100E450g,
	"SUBCORE_VM_FIXED2125_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2125E450g,
	"SUBCORE_VM_FIXED2150_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2150E450g,
	"SUBCORE_VM_FIXED2175_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2175E450g,
	"SUBCORE_VM_FIXED2200_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2200E450g,
	"SUBCORE_VM_FIXED2250_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2250E450g,
	"SUBCORE_VM_FIXED2275_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2275E450g,
	"SUBCORE_VM_FIXED2300_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2300E450g,
	"SUBCORE_VM_FIXED2325_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2325E450g,
	"SUBCORE_VM_FIXED2350_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2350E450g,
	"SUBCORE_VM_FIXED2375_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2375E450g,
	"SUBCORE_VM_FIXED2400_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2400E450g,
	"SUBCORE_VM_FIXED2450_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2450E450g,
	"SUBCORE_VM_FIXED2475_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2475E450g,
	"SUBCORE_VM_FIXED2500_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2500E450g,
	"SUBCORE_VM_FIXED2550_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2550E450g,
	"SUBCORE_VM_FIXED2600_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2600E450g,
	"SUBCORE_VM_FIXED2625_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2625E450g,
	"SUBCORE_VM_FIXED2650_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2650E450g,
	"SUBCORE_VM_FIXED2700_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2700E450g,
	"SUBCORE_VM_FIXED2750_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2750E450g,
	"SUBCORE_VM_FIXED2775_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2775E450g,
	"SUBCORE_VM_FIXED2800_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2800E450g,
	"SUBCORE_VM_FIXED2850_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2850E450g,
	"SUBCORE_VM_FIXED2875_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2875E450g,
	"SUBCORE_VM_FIXED2900_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2900E450g,
	"SUBCORE_VM_FIXED2925_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2925E450g,
	"SUBCORE_VM_FIXED2950_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2950E450g,
	"SUBCORE_VM_FIXED2975_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2975E450g,
	"SUBCORE_VM_FIXED3000_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3000E450g,
	"SUBCORE_VM_FIXED3025_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3025E450g,
	"SUBCORE_VM_FIXED3050_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3050E450g,
	"SUBCORE_VM_FIXED3075_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3075E450g,
	"SUBCORE_VM_FIXED3100_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3100E450g,
	"SUBCORE_VM_FIXED3125_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3125E450g,
	"SUBCORE_VM_FIXED3150_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3150E450g,
	"SUBCORE_VM_FIXED3200_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3200E450g,
	"SUBCORE_VM_FIXED3225_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3225E450g,
	"SUBCORE_VM_FIXED3250_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3250E450g,
	"SUBCORE_VM_FIXED3300_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3300E450g,
	"SUBCORE_VM_FIXED3325_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3325E450g,
	"SUBCORE_VM_FIXED3375_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3375E450g,
	"SUBCORE_VM_FIXED3400_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3400E450g,
	"SUBCORE_VM_FIXED3450_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3450E450g,
	"SUBCORE_VM_FIXED3500_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3500E450g,
	"SUBCORE_VM_FIXED3525_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3525E450g,
	"SUBCORE_VM_FIXED3575_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3575E450g,
	"SUBCORE_VM_FIXED3600_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3600E450g,
	"SUBCORE_VM_FIXED3625_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3625E450g,
	"SUBCORE_VM_FIXED3675_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3675E450g,
	"SUBCORE_VM_FIXED3700_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3700E450g,
	"SUBCORE_VM_FIXED3750_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3750E450g,
	"SUBCORE_VM_FIXED3800_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3800E450g,
	"SUBCORE_VM_FIXED3825_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3825E450g,
	"SUBCORE_VM_FIXED3850_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3850E450g,
	"SUBCORE_VM_FIXED3875_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3875E450g,
	"SUBCORE_VM_FIXED3900_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3900E450g,
	"SUBCORE_VM_FIXED3975_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3975E450g,
	"SUBCORE_VM_FIXED4000_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4000E450g,
	"SUBCORE_VM_FIXED4025_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4025E450g,
	"SUBCORE_VM_FIXED4050_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4050E450g,
	"SUBCORE_VM_FIXED4100_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4100E450g,
	"SUBCORE_VM_FIXED4125_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4125E450g,
	"SUBCORE_VM_FIXED4200_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4200E450g,
	"SUBCORE_VM_FIXED4225_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4225E450g,
	"SUBCORE_VM_FIXED4250_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4250E450g,
	"SUBCORE_VM_FIXED4275_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4275E450g,
	"SUBCORE_VM_FIXED4300_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4300E450g,
	"SUBCORE_VM_FIXED4350_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4350E450g,
	"SUBCORE_VM_FIXED4375_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4375E450g,
	"SUBCORE_VM_FIXED4400_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4400E450g,
	"SUBCORE_VM_FIXED4425_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4425E450g,
	"SUBCORE_VM_FIXED4500_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4500E450g,
	"SUBCORE_VM_FIXED4550_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4550E450g,
	"SUBCORE_VM_FIXED4575_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4575E450g,
	"SUBCORE_VM_FIXED4600_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4600E450g,
	"SUBCORE_VM_FIXED4625_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4625E450g,
	"SUBCORE_VM_FIXED4650_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4650E450g,
	"SUBCORE_VM_FIXED4675_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4675E450g,
	"SUBCORE_VM_FIXED4700_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4700E450g,
	"SUBCORE_VM_FIXED4725_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4725E450g,
	"SUBCORE_VM_FIXED4750_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4750E450g,
	"SUBCORE_VM_FIXED4800_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4800E450g,
	"SUBCORE_VM_FIXED4875_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4875E450g,
	"SUBCORE_VM_FIXED4900_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4900E450g,
	"SUBCORE_VM_FIXED4950_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4950E450g,
	"SUBCORE_VM_FIXED5000_E4_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed5000E450g,
	"DYNAMIC_E5_50G":                       CreateInternalVnicAttachmentDetailsVnicShapeDynamicE550g,
	"DYNAMIC_E5_100G":                      CreateInternalVnicAttachmentDetailsVnicShapeDynamicE5100g,
	"SUBCORE_VM_FIXED0020_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0020A150g,
	"SUBCORE_VM_FIXED0040_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0040A150g,
	"SUBCORE_VM_FIXED0060_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0060A150g,
	"SUBCORE_VM_FIXED0080_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0080A150g,
	"SUBCORE_VM_FIXED0100_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0100A150g,
	"SUBCORE_VM_FIXED0120_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0120A150g,
	"SUBCORE_VM_FIXED0140_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0140A150g,
	"SUBCORE_VM_FIXED0160_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0160A150g,
	"SUBCORE_VM_FIXED0180_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0180A150g,
	"SUBCORE_VM_FIXED0200_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0200A150g,
	"SUBCORE_VM_FIXED0220_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0220A150g,
	"SUBCORE_VM_FIXED0240_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0240A150g,
	"SUBCORE_VM_FIXED0260_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0260A150g,
	"SUBCORE_VM_FIXED0280_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0280A150g,
	"SUBCORE_VM_FIXED0300_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0300A150g,
	"SUBCORE_VM_FIXED0320_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0320A150g,
	"SUBCORE_VM_FIXED0340_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0340A150g,
	"SUBCORE_VM_FIXED0360_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0360A150g,
	"SUBCORE_VM_FIXED0380_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0380A150g,
	"SUBCORE_VM_FIXED0400_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0400A150g,
	"SUBCORE_VM_FIXED0420_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0420A150g,
	"SUBCORE_VM_FIXED0440_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0440A150g,
	"SUBCORE_VM_FIXED0460_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0460A150g,
	"SUBCORE_VM_FIXED0480_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0480A150g,
	"SUBCORE_VM_FIXED0500_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0500A150g,
	"SUBCORE_VM_FIXED0520_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0520A150g,
	"SUBCORE_VM_FIXED0540_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0540A150g,
	"SUBCORE_VM_FIXED0560_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0560A150g,
	"SUBCORE_VM_FIXED0580_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0580A150g,
	"SUBCORE_VM_FIXED0600_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0600A150g,
	"SUBCORE_VM_FIXED0620_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0620A150g,
	"SUBCORE_VM_FIXED0640_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0640A150g,
	"SUBCORE_VM_FIXED0660_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0660A150g,
	"SUBCORE_VM_FIXED0680_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0680A150g,
	"SUBCORE_VM_FIXED0700_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0700A150g,
	"SUBCORE_VM_FIXED0720_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0720A150g,
	"SUBCORE_VM_FIXED0740_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0740A150g,
	"SUBCORE_VM_FIXED0760_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0760A150g,
	"SUBCORE_VM_FIXED0780_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0780A150g,
	"SUBCORE_VM_FIXED0800_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0800A150g,
	"SUBCORE_VM_FIXED0820_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0820A150g,
	"SUBCORE_VM_FIXED0840_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0840A150g,
	"SUBCORE_VM_FIXED0860_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0860A150g,
	"SUBCORE_VM_FIXED0880_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0880A150g,
	"SUBCORE_VM_FIXED0900_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0900A150g,
	"SUBCORE_VM_FIXED0920_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0920A150g,
	"SUBCORE_VM_FIXED0940_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0940A150g,
	"SUBCORE_VM_FIXED0960_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0960A150g,
	"SUBCORE_VM_FIXED0980_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0980A150g,
	"SUBCORE_VM_FIXED1000_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1000A150g,
	"SUBCORE_VM_FIXED1020_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1020A150g,
	"SUBCORE_VM_FIXED1040_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1040A150g,
	"SUBCORE_VM_FIXED1060_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1060A150g,
	"SUBCORE_VM_FIXED1080_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1080A150g,
	"SUBCORE_VM_FIXED1100_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1100A150g,
	"SUBCORE_VM_FIXED1120_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1120A150g,
	"SUBCORE_VM_FIXED1140_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1140A150g,
	"SUBCORE_VM_FIXED1160_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1160A150g,
	"SUBCORE_VM_FIXED1180_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1180A150g,
	"SUBCORE_VM_FIXED1200_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1200A150g,
	"SUBCORE_VM_FIXED1220_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1220A150g,
	"SUBCORE_VM_FIXED1240_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1240A150g,
	"SUBCORE_VM_FIXED1260_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1260A150g,
	"SUBCORE_VM_FIXED1280_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1280A150g,
	"SUBCORE_VM_FIXED1300_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1300A150g,
	"SUBCORE_VM_FIXED1320_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1320A150g,
	"SUBCORE_VM_FIXED1340_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1340A150g,
	"SUBCORE_VM_FIXED1360_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1360A150g,
	"SUBCORE_VM_FIXED1380_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1380A150g,
	"SUBCORE_VM_FIXED1400_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1400A150g,
	"SUBCORE_VM_FIXED1420_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1420A150g,
	"SUBCORE_VM_FIXED1440_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1440A150g,
	"SUBCORE_VM_FIXED1460_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1460A150g,
	"SUBCORE_VM_FIXED1480_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1480A150g,
	"SUBCORE_VM_FIXED1500_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1500A150g,
	"SUBCORE_VM_FIXED1520_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1520A150g,
	"SUBCORE_VM_FIXED1540_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1540A150g,
	"SUBCORE_VM_FIXED1560_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1560A150g,
	"SUBCORE_VM_FIXED1580_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1580A150g,
	"SUBCORE_VM_FIXED1600_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1600A150g,
	"SUBCORE_VM_FIXED1620_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1620A150g,
	"SUBCORE_VM_FIXED1640_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1640A150g,
	"SUBCORE_VM_FIXED1660_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1660A150g,
	"SUBCORE_VM_FIXED1680_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1680A150g,
	"SUBCORE_VM_FIXED1700_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1700A150g,
	"SUBCORE_VM_FIXED1720_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1720A150g,
	"SUBCORE_VM_FIXED1740_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1740A150g,
	"SUBCORE_VM_FIXED1760_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1760A150g,
	"SUBCORE_VM_FIXED1780_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1780A150g,
	"SUBCORE_VM_FIXED1800_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1800A150g,
	"SUBCORE_VM_FIXED1820_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1820A150g,
	"SUBCORE_VM_FIXED1840_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1840A150g,
	"SUBCORE_VM_FIXED1860_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1860A150g,
	"SUBCORE_VM_FIXED1880_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1880A150g,
	"SUBCORE_VM_FIXED1900_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1900A150g,
	"SUBCORE_VM_FIXED1920_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1920A150g,
	"SUBCORE_VM_FIXED1940_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1940A150g,
	"SUBCORE_VM_FIXED1960_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1960A150g,
	"SUBCORE_VM_FIXED1980_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1980A150g,
	"SUBCORE_VM_FIXED2000_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2000A150g,
	"SUBCORE_VM_FIXED2020_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2020A150g,
	"SUBCORE_VM_FIXED2040_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2040A150g,
	"SUBCORE_VM_FIXED2060_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2060A150g,
	"SUBCORE_VM_FIXED2080_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2080A150g,
	"SUBCORE_VM_FIXED2100_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2100A150g,
	"SUBCORE_VM_FIXED2120_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2120A150g,
	"SUBCORE_VM_FIXED2140_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2140A150g,
	"SUBCORE_VM_FIXED2160_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2160A150g,
	"SUBCORE_VM_FIXED2180_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2180A150g,
	"SUBCORE_VM_FIXED2200_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2200A150g,
	"SUBCORE_VM_FIXED2220_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2220A150g,
	"SUBCORE_VM_FIXED2240_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2240A150g,
	"SUBCORE_VM_FIXED2260_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2260A150g,
	"SUBCORE_VM_FIXED2280_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2280A150g,
	"SUBCORE_VM_FIXED2300_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2300A150g,
	"SUBCORE_VM_FIXED2320_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2320A150g,
	"SUBCORE_VM_FIXED2340_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2340A150g,
	"SUBCORE_VM_FIXED2360_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2360A150g,
	"SUBCORE_VM_FIXED2380_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2380A150g,
	"SUBCORE_VM_FIXED2400_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2400A150g,
	"SUBCORE_VM_FIXED2420_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2420A150g,
	"SUBCORE_VM_FIXED2440_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2440A150g,
	"SUBCORE_VM_FIXED2460_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2460A150g,
	"SUBCORE_VM_FIXED2480_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2480A150g,
	"SUBCORE_VM_FIXED2500_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2500A150g,
	"SUBCORE_VM_FIXED2520_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2520A150g,
	"SUBCORE_VM_FIXED2540_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2540A150g,
	"SUBCORE_VM_FIXED2560_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2560A150g,
	"SUBCORE_VM_FIXED2580_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2580A150g,
	"SUBCORE_VM_FIXED2600_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2600A150g,
	"SUBCORE_VM_FIXED2620_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2620A150g,
	"SUBCORE_VM_FIXED2640_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2640A150g,
	"SUBCORE_VM_FIXED2660_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2660A150g,
	"SUBCORE_VM_FIXED2680_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2680A150g,
	"SUBCORE_VM_FIXED2700_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2700A150g,
	"SUBCORE_VM_FIXED2720_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2720A150g,
	"SUBCORE_VM_FIXED2740_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2740A150g,
	"SUBCORE_VM_FIXED2760_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2760A150g,
	"SUBCORE_VM_FIXED2780_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2780A150g,
	"SUBCORE_VM_FIXED2800_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2800A150g,
	"SUBCORE_VM_FIXED2820_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2820A150g,
	"SUBCORE_VM_FIXED2840_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2840A150g,
	"SUBCORE_VM_FIXED2860_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2860A150g,
	"SUBCORE_VM_FIXED2880_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2880A150g,
	"SUBCORE_VM_FIXED2900_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2900A150g,
	"SUBCORE_VM_FIXED2920_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2920A150g,
	"SUBCORE_VM_FIXED2940_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2940A150g,
	"SUBCORE_VM_FIXED2960_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2960A150g,
	"SUBCORE_VM_FIXED2980_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2980A150g,
	"SUBCORE_VM_FIXED3000_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3000A150g,
	"SUBCORE_VM_FIXED3020_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3020A150g,
	"SUBCORE_VM_FIXED3040_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3040A150g,
	"SUBCORE_VM_FIXED3060_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3060A150g,
	"SUBCORE_VM_FIXED3080_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3080A150g,
	"SUBCORE_VM_FIXED3100_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3100A150g,
	"SUBCORE_VM_FIXED3120_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3120A150g,
	"SUBCORE_VM_FIXED3140_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3140A150g,
	"SUBCORE_VM_FIXED3160_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3160A150g,
	"SUBCORE_VM_FIXED3180_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3180A150g,
	"SUBCORE_VM_FIXED3200_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3200A150g,
	"SUBCORE_VM_FIXED3220_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3220A150g,
	"SUBCORE_VM_FIXED3240_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3240A150g,
	"SUBCORE_VM_FIXED3260_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3260A150g,
	"SUBCORE_VM_FIXED3280_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3280A150g,
	"SUBCORE_VM_FIXED3300_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3300A150g,
	"SUBCORE_VM_FIXED3320_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3320A150g,
	"SUBCORE_VM_FIXED3340_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3340A150g,
	"SUBCORE_VM_FIXED3360_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3360A150g,
	"SUBCORE_VM_FIXED3380_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3380A150g,
	"SUBCORE_VM_FIXED3400_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3400A150g,
	"SUBCORE_VM_FIXED3420_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3420A150g,
	"SUBCORE_VM_FIXED3440_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3440A150g,
	"SUBCORE_VM_FIXED3460_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3460A150g,
	"SUBCORE_VM_FIXED3480_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3480A150g,
	"SUBCORE_VM_FIXED3500_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3500A150g,
	"SUBCORE_VM_FIXED3520_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3520A150g,
	"SUBCORE_VM_FIXED3540_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3540A150g,
	"SUBCORE_VM_FIXED3560_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3560A150g,
	"SUBCORE_VM_FIXED3580_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3580A150g,
	"SUBCORE_VM_FIXED3600_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3600A150g,
	"SUBCORE_VM_FIXED3620_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3620A150g,
	"SUBCORE_VM_FIXED3640_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3640A150g,
	"SUBCORE_VM_FIXED3660_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3660A150g,
	"SUBCORE_VM_FIXED3680_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3680A150g,
	"SUBCORE_VM_FIXED3700_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3700A150g,
	"SUBCORE_VM_FIXED3720_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3720A150g,
	"SUBCORE_VM_FIXED3740_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3740A150g,
	"SUBCORE_VM_FIXED3760_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3760A150g,
	"SUBCORE_VM_FIXED3780_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3780A150g,
	"SUBCORE_VM_FIXED3800_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3800A150g,
	"SUBCORE_VM_FIXED3820_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3820A150g,
	"SUBCORE_VM_FIXED3840_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3840A150g,
	"SUBCORE_VM_FIXED3860_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3860A150g,
	"SUBCORE_VM_FIXED3880_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3880A150g,
	"SUBCORE_VM_FIXED3900_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3900A150g,
	"SUBCORE_VM_FIXED3920_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3920A150g,
	"SUBCORE_VM_FIXED3940_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3940A150g,
	"SUBCORE_VM_FIXED3960_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3960A150g,
	"SUBCORE_VM_FIXED3980_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3980A150g,
	"SUBCORE_VM_FIXED4000_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4000A150g,
	"SUBCORE_VM_FIXED4020_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4020A150g,
	"SUBCORE_VM_FIXED4040_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4040A150g,
	"SUBCORE_VM_FIXED4060_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4060A150g,
	"SUBCORE_VM_FIXED4080_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4080A150g,
	"SUBCORE_VM_FIXED4100_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4100A150g,
	"SUBCORE_VM_FIXED4120_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4120A150g,
	"SUBCORE_VM_FIXED4140_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4140A150g,
	"SUBCORE_VM_FIXED4160_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4160A150g,
	"SUBCORE_VM_FIXED4180_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4180A150g,
	"SUBCORE_VM_FIXED4200_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4200A150g,
	"SUBCORE_VM_FIXED4220_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4220A150g,
	"SUBCORE_VM_FIXED4240_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4240A150g,
	"SUBCORE_VM_FIXED4260_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4260A150g,
	"SUBCORE_VM_FIXED4280_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4280A150g,
	"SUBCORE_VM_FIXED4300_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4300A150g,
	"SUBCORE_VM_FIXED4320_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4320A150g,
	"SUBCORE_VM_FIXED4340_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4340A150g,
	"SUBCORE_VM_FIXED4360_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4360A150g,
	"SUBCORE_VM_FIXED4380_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4380A150g,
	"SUBCORE_VM_FIXED4400_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4400A150g,
	"SUBCORE_VM_FIXED4420_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4420A150g,
	"SUBCORE_VM_FIXED4440_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4440A150g,
	"SUBCORE_VM_FIXED4460_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4460A150g,
	"SUBCORE_VM_FIXED4480_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4480A150g,
	"SUBCORE_VM_FIXED4500_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4500A150g,
	"SUBCORE_VM_FIXED4520_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4520A150g,
	"SUBCORE_VM_FIXED4540_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4540A150g,
	"SUBCORE_VM_FIXED4560_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4560A150g,
	"SUBCORE_VM_FIXED4580_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4580A150g,
	"SUBCORE_VM_FIXED4600_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4600A150g,
	"SUBCORE_VM_FIXED4620_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4620A150g,
	"SUBCORE_VM_FIXED4640_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4640A150g,
	"SUBCORE_VM_FIXED4660_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4660A150g,
	"SUBCORE_VM_FIXED4680_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4680A150g,
	"SUBCORE_VM_FIXED4700_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4700A150g,
	"SUBCORE_VM_FIXED4720_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4720A150g,
	"SUBCORE_VM_FIXED4740_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4740A150g,
	"SUBCORE_VM_FIXED4760_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4760A150g,
	"SUBCORE_VM_FIXED4780_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4780A150g,
	"SUBCORE_VM_FIXED4800_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4800A150g,
	"SUBCORE_VM_FIXED4820_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4820A150g,
	"SUBCORE_VM_FIXED4840_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4840A150g,
	"SUBCORE_VM_FIXED4860_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4860A150g,
	"SUBCORE_VM_FIXED4880_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4880A150g,
	"SUBCORE_VM_FIXED4900_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4900A150g,
	"SUBCORE_VM_FIXED4920_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4920A150g,
	"SUBCORE_VM_FIXED4940_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4940A150g,
	"SUBCORE_VM_FIXED4960_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4960A150g,
	"SUBCORE_VM_FIXED4980_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4980A150g,
	"SUBCORE_VM_FIXED5000_A1_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed5000A150g,
	"SUBCORE_VM_FIXED0090_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0090X950g,
	"SUBCORE_VM_FIXED0180_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0180X950g,
	"SUBCORE_VM_FIXED0270_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0270X950g,
	"SUBCORE_VM_FIXED0360_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0360X950g,
	"SUBCORE_VM_FIXED0450_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0450X950g,
	"SUBCORE_VM_FIXED0540_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0540X950g,
	"SUBCORE_VM_FIXED0630_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0630X950g,
	"SUBCORE_VM_FIXED0720_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0720X950g,
	"SUBCORE_VM_FIXED0810_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0810X950g,
	"SUBCORE_VM_FIXED0900_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0900X950g,
	"SUBCORE_VM_FIXED0990_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0990X950g,
	"SUBCORE_VM_FIXED1080_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1080X950g,
	"SUBCORE_VM_FIXED1170_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1170X950g,
	"SUBCORE_VM_FIXED1260_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1260X950g,
	"SUBCORE_VM_FIXED1350_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1350X950g,
	"SUBCORE_VM_FIXED1440_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1440X950g,
	"SUBCORE_VM_FIXED1530_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1530X950g,
	"SUBCORE_VM_FIXED1620_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1620X950g,
	"SUBCORE_VM_FIXED1710_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1710X950g,
	"SUBCORE_VM_FIXED1800_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1800X950g,
	"SUBCORE_VM_FIXED1890_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1890X950g,
	"SUBCORE_VM_FIXED1980_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1980X950g,
	"SUBCORE_VM_FIXED2070_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2070X950g,
	"SUBCORE_VM_FIXED2160_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2160X950g,
	"SUBCORE_VM_FIXED2250_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2250X950g,
	"SUBCORE_VM_FIXED2340_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2340X950g,
	"SUBCORE_VM_FIXED2430_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2430X950g,
	"SUBCORE_VM_FIXED2520_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2520X950g,
	"SUBCORE_VM_FIXED2610_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2610X950g,
	"SUBCORE_VM_FIXED2700_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2700X950g,
	"SUBCORE_VM_FIXED2790_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2790X950g,
	"SUBCORE_VM_FIXED2880_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2880X950g,
	"SUBCORE_VM_FIXED2970_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2970X950g,
	"SUBCORE_VM_FIXED3060_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3060X950g,
	"SUBCORE_VM_FIXED3150_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3150X950g,
	"SUBCORE_VM_FIXED3240_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3240X950g,
	"SUBCORE_VM_FIXED3330_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3330X950g,
	"SUBCORE_VM_FIXED3420_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3420X950g,
	"SUBCORE_VM_FIXED3510_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3510X950g,
	"SUBCORE_VM_FIXED3600_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3600X950g,
	"SUBCORE_VM_FIXED3690_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3690X950g,
	"SUBCORE_VM_FIXED3780_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3780X950g,
	"SUBCORE_VM_FIXED3870_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3870X950g,
	"SUBCORE_VM_FIXED3960_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3960X950g,
	"SUBCORE_VM_FIXED4050_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4050X950g,
	"SUBCORE_VM_FIXED4140_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4140X950g,
	"SUBCORE_VM_FIXED4230_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4230X950g,
	"SUBCORE_VM_FIXED4320_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4320X950g,
	"SUBCORE_VM_FIXED4410_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4410X950g,
	"SUBCORE_VM_FIXED4500_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4500X950g,
	"SUBCORE_VM_FIXED4590_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4590X950g,
	"SUBCORE_VM_FIXED4680_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4680X950g,
	"SUBCORE_VM_FIXED4770_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4770X950g,
	"SUBCORE_VM_FIXED4860_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4860X950g,
	"SUBCORE_VM_FIXED4950_X9_50G":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4950X950g,
	"DYNAMIC_A1_50G":                       CreateInternalVnicAttachmentDetailsVnicShapeDynamicA150g,
	"FIXED0040_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0040A150g,
	"FIXED0100_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0100A150g,
	"FIXED0200_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0200A150g,
	"FIXED0300_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0300A150g,
	"FIXED0400_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0400A150g,
	"FIXED0500_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0500A150g,
	"FIXED0600_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0600A150g,
	"FIXED0700_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0700A150g,
	"FIXED0800_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0800A150g,
	"FIXED0900_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0900A150g,
	"FIXED1000_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1000A150g,
	"FIXED1100_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1100A150g,
	"FIXED1200_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1200A150g,
	"FIXED1300_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1300A150g,
	"FIXED1400_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1400A150g,
	"FIXED1500_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1500A150g,
	"FIXED1600_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1600A150g,
	"FIXED1700_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1700A150g,
	"FIXED1800_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1800A150g,
	"FIXED1900_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1900A150g,
	"FIXED2000_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2000A150g,
	"FIXED2100_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2100A150g,
	"FIXED2200_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2200A150g,
	"FIXED2300_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2300A150g,
	"FIXED2400_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2400A150g,
	"FIXED2500_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2500A150g,
	"FIXED2600_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2600A150g,
	"FIXED2700_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2700A150g,
	"FIXED2800_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2800A150g,
	"FIXED2900_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2900A150g,
	"FIXED3000_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3000A150g,
	"FIXED3100_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3100A150g,
	"FIXED3200_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3200A150g,
	"FIXED3300_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3300A150g,
	"FIXED3400_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3400A150g,
	"FIXED3500_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3500A150g,
	"FIXED3600_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3600A150g,
	"FIXED3700_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3700A150g,
	"FIXED3800_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3800A150g,
	"FIXED3900_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3900A150g,
	"FIXED4000_A1_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed4000A150g,
	"FIXED5000_TELESIS_A1_50G":             CreateInternalVnicAttachmentDetailsVnicShapeFixed5000TelesisA150g,
	"ENTIREHOST_A1_50G":                    CreateInternalVnicAttachmentDetailsVnicShapeEntirehostA150g,
	"DYNAMIC_X9_50G":                       CreateInternalVnicAttachmentDetailsVnicShapeDynamicX950g,
	"FIXED0040_X9_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0040X950g,
	"FIXED0400_X9_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0400X950g,
	"FIXED0800_X9_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0800X950g,
	"FIXED1200_X9_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1200X950g,
	"FIXED1600_X9_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1600X950g,
	"FIXED2000_X9_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2000X950g,
	"FIXED2400_X9_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2400X950g,
	"FIXED2800_X9_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2800X950g,
	"FIXED3200_X9_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3200X950g,
	"FIXED3600_X9_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3600X950g,
	"FIXED4000_X9_50G":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed4000X950g,
	"STANDARD_VM_FIXED0100_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0100X950g,
	"STANDARD_VM_FIXED0200_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0200X950g,
	"STANDARD_VM_FIXED0300_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0300X950g,
	"STANDARD_VM_FIXED0400_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0400X950g,
	"STANDARD_VM_FIXED0500_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0500X950g,
	"STANDARD_VM_FIXED0600_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0600X950g,
	"STANDARD_VM_FIXED0700_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0700X950g,
	"STANDARD_VM_FIXED0800_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0800X950g,
	"STANDARD_VM_FIXED0900_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0900X950g,
	"STANDARD_VM_FIXED1000_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1000X950g,
	"STANDARD_VM_FIXED1100_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1100X950g,
	"STANDARD_VM_FIXED1200_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1200X950g,
	"STANDARD_VM_FIXED1300_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1300X950g,
	"STANDARD_VM_FIXED1400_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1400X950g,
	"STANDARD_VM_FIXED1500_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1500X950g,
	"STANDARD_VM_FIXED1600_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1600X950g,
	"STANDARD_VM_FIXED1700_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1700X950g,
	"STANDARD_VM_FIXED1800_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1800X950g,
	"STANDARD_VM_FIXED1900_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1900X950g,
	"STANDARD_VM_FIXED2000_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2000X950g,
	"STANDARD_VM_FIXED2100_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2100X950g,
	"STANDARD_VM_FIXED2200_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2200X950g,
	"STANDARD_VM_FIXED2300_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2300X950g,
	"STANDARD_VM_FIXED2400_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2400X950g,
	"STANDARD_VM_FIXED2500_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2500X950g,
	"STANDARD_VM_FIXED2600_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2600X950g,
	"STANDARD_VM_FIXED2700_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2700X950g,
	"STANDARD_VM_FIXED2800_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2800X950g,
	"STANDARD_VM_FIXED2900_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2900X950g,
	"STANDARD_VM_FIXED3000_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3000X950g,
	"STANDARD_VM_FIXED3100_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3100X950g,
	"STANDARD_VM_FIXED3200_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3200X950g,
	"STANDARD_VM_FIXED3300_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3300X950g,
	"STANDARD_VM_FIXED3400_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3400X950g,
	"STANDARD_VM_FIXED3500_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3500X950g,
	"STANDARD_VM_FIXED3600_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3600X950g,
	"STANDARD_VM_FIXED3700_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3700X950g,
	"STANDARD_VM_FIXED3800_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3800X950g,
	"STANDARD_VM_FIXED3900_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3900X950g,
	"STANDARD_VM_FIXED4000_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4000X950g,
	"STANDARD_VM_FIXED4100_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4100X950g,
	"STANDARD_VM_FIXED4200_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4200X950g,
	"STANDARD_VM_FIXED4300_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4300X950g,
	"STANDARD_VM_FIXED4400_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4400X950g,
	"STANDARD_VM_FIXED4500_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4500X950g,
	"STANDARD_VM_FIXED4600_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4600X950g,
	"STANDARD_VM_FIXED4700_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4700X950g,
	"STANDARD_VM_FIXED4800_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4800X950g,
	"STANDARD_VM_FIXED4900_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4900X950g,
	"STANDARD_VM_FIXED5000_X9_50G":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed5000X950g,
	"SUBCORE_STANDARD_VM_FIXED0025_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0025X950g,
	"SUBCORE_STANDARD_VM_FIXED0050_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0050X950g,
	"SUBCORE_STANDARD_VM_FIXED0075_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0075X950g,
	"SUBCORE_STANDARD_VM_FIXED0100_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0100X950g,
	"SUBCORE_STANDARD_VM_FIXED0125_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0125X950g,
	"SUBCORE_STANDARD_VM_FIXED0150_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0150X950g,
	"SUBCORE_STANDARD_VM_FIXED0175_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0175X950g,
	"SUBCORE_STANDARD_VM_FIXED0200_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0200X950g,
	"SUBCORE_STANDARD_VM_FIXED0225_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0225X950g,
	"SUBCORE_STANDARD_VM_FIXED0250_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0250X950g,
	"SUBCORE_STANDARD_VM_FIXED0275_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0275X950g,
	"SUBCORE_STANDARD_VM_FIXED0300_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0300X950g,
	"SUBCORE_STANDARD_VM_FIXED0325_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0325X950g,
	"SUBCORE_STANDARD_VM_FIXED0350_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0350X950g,
	"SUBCORE_STANDARD_VM_FIXED0375_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0375X950g,
	"SUBCORE_STANDARD_VM_FIXED0400_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0400X950g,
	"SUBCORE_STANDARD_VM_FIXED0425_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0425X950g,
	"SUBCORE_STANDARD_VM_FIXED0450_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0450X950g,
	"SUBCORE_STANDARD_VM_FIXED0475_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0475X950g,
	"SUBCORE_STANDARD_VM_FIXED0500_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0500X950g,
	"SUBCORE_STANDARD_VM_FIXED0525_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0525X950g,
	"SUBCORE_STANDARD_VM_FIXED0550_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0550X950g,
	"SUBCORE_STANDARD_VM_FIXED0575_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0575X950g,
	"SUBCORE_STANDARD_VM_FIXED0600_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0600X950g,
	"SUBCORE_STANDARD_VM_FIXED0625_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0625X950g,
	"SUBCORE_STANDARD_VM_FIXED0650_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0650X950g,
	"SUBCORE_STANDARD_VM_FIXED0675_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0675X950g,
	"SUBCORE_STANDARD_VM_FIXED0700_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0700X950g,
	"SUBCORE_STANDARD_VM_FIXED0725_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0725X950g,
	"SUBCORE_STANDARD_VM_FIXED0750_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0750X950g,
	"SUBCORE_STANDARD_VM_FIXED0775_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0775X950g,
	"SUBCORE_STANDARD_VM_FIXED0800_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0800X950g,
	"SUBCORE_STANDARD_VM_FIXED0825_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0825X950g,
	"SUBCORE_STANDARD_VM_FIXED0850_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0850X950g,
	"SUBCORE_STANDARD_VM_FIXED0875_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0875X950g,
	"SUBCORE_STANDARD_VM_FIXED0900_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0900X950g,
	"SUBCORE_STANDARD_VM_FIXED0925_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0925X950g,
	"SUBCORE_STANDARD_VM_FIXED0950_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0950X950g,
	"SUBCORE_STANDARD_VM_FIXED0975_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0975X950g,
	"SUBCORE_STANDARD_VM_FIXED1000_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1000X950g,
	"SUBCORE_STANDARD_VM_FIXED1025_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1025X950g,
	"SUBCORE_STANDARD_VM_FIXED1050_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1050X950g,
	"SUBCORE_STANDARD_VM_FIXED1075_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1075X950g,
	"SUBCORE_STANDARD_VM_FIXED1100_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1100X950g,
	"SUBCORE_STANDARD_VM_FIXED1125_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1125X950g,
	"SUBCORE_STANDARD_VM_FIXED1150_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1150X950g,
	"SUBCORE_STANDARD_VM_FIXED1175_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1175X950g,
	"SUBCORE_STANDARD_VM_FIXED1200_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1200X950g,
	"SUBCORE_STANDARD_VM_FIXED1225_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1225X950g,
	"SUBCORE_STANDARD_VM_FIXED1250_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1250X950g,
	"SUBCORE_STANDARD_VM_FIXED1275_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1275X950g,
	"SUBCORE_STANDARD_VM_FIXED1300_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1300X950g,
	"SUBCORE_STANDARD_VM_FIXED1325_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1325X950g,
	"SUBCORE_STANDARD_VM_FIXED1350_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1350X950g,
	"SUBCORE_STANDARD_VM_FIXED1375_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1375X950g,
	"SUBCORE_STANDARD_VM_FIXED1400_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1400X950g,
	"SUBCORE_STANDARD_VM_FIXED1425_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1425X950g,
	"SUBCORE_STANDARD_VM_FIXED1450_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1450X950g,
	"SUBCORE_STANDARD_VM_FIXED1475_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1475X950g,
	"SUBCORE_STANDARD_VM_FIXED1500_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1500X950g,
	"SUBCORE_STANDARD_VM_FIXED1525_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1525X950g,
	"SUBCORE_STANDARD_VM_FIXED1550_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1550X950g,
	"SUBCORE_STANDARD_VM_FIXED1575_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1575X950g,
	"SUBCORE_STANDARD_VM_FIXED1600_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1600X950g,
	"SUBCORE_STANDARD_VM_FIXED1625_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1625X950g,
	"SUBCORE_STANDARD_VM_FIXED1650_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1650X950g,
	"SUBCORE_STANDARD_VM_FIXED1700_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1700X950g,
	"SUBCORE_STANDARD_VM_FIXED1725_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1725X950g,
	"SUBCORE_STANDARD_VM_FIXED1750_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1750X950g,
	"SUBCORE_STANDARD_VM_FIXED1800_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1800X950g,
	"SUBCORE_STANDARD_VM_FIXED1850_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1850X950g,
	"SUBCORE_STANDARD_VM_FIXED1875_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1875X950g,
	"SUBCORE_STANDARD_VM_FIXED1900_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1900X950g,
	"SUBCORE_STANDARD_VM_FIXED1925_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1925X950g,
	"SUBCORE_STANDARD_VM_FIXED1950_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1950X950g,
	"SUBCORE_STANDARD_VM_FIXED2000_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2000X950g,
	"SUBCORE_STANDARD_VM_FIXED2025_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2025X950g,
	"SUBCORE_STANDARD_VM_FIXED2050_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2050X950g,
	"SUBCORE_STANDARD_VM_FIXED2100_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2100X950g,
	"SUBCORE_STANDARD_VM_FIXED2125_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2125X950g,
	"SUBCORE_STANDARD_VM_FIXED2150_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2150X950g,
	"SUBCORE_STANDARD_VM_FIXED2175_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2175X950g,
	"SUBCORE_STANDARD_VM_FIXED2200_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2200X950g,
	"SUBCORE_STANDARD_VM_FIXED2250_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2250X950g,
	"SUBCORE_STANDARD_VM_FIXED2275_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2275X950g,
	"SUBCORE_STANDARD_VM_FIXED2300_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2300X950g,
	"SUBCORE_STANDARD_VM_FIXED2325_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2325X950g,
	"SUBCORE_STANDARD_VM_FIXED2350_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2350X950g,
	"SUBCORE_STANDARD_VM_FIXED2375_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2375X950g,
	"SUBCORE_STANDARD_VM_FIXED2400_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2400X950g,
	"SUBCORE_STANDARD_VM_FIXED2450_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2450X950g,
	"SUBCORE_STANDARD_VM_FIXED2475_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2475X950g,
	"SUBCORE_STANDARD_VM_FIXED2500_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2500X950g,
	"SUBCORE_STANDARD_VM_FIXED2550_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2550X950g,
	"SUBCORE_STANDARD_VM_FIXED2600_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2600X950g,
	"SUBCORE_STANDARD_VM_FIXED2625_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2625X950g,
	"SUBCORE_STANDARD_VM_FIXED2650_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2650X950g,
	"SUBCORE_STANDARD_VM_FIXED2700_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2700X950g,
	"SUBCORE_STANDARD_VM_FIXED2750_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2750X950g,
	"SUBCORE_STANDARD_VM_FIXED2775_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2775X950g,
	"SUBCORE_STANDARD_VM_FIXED2800_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2800X950g,
	"SUBCORE_STANDARD_VM_FIXED2850_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2850X950g,
	"SUBCORE_STANDARD_VM_FIXED2875_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2875X950g,
	"SUBCORE_STANDARD_VM_FIXED2900_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2900X950g,
	"SUBCORE_STANDARD_VM_FIXED2925_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2925X950g,
	"SUBCORE_STANDARD_VM_FIXED2950_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2950X950g,
	"SUBCORE_STANDARD_VM_FIXED2975_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2975X950g,
	"SUBCORE_STANDARD_VM_FIXED3000_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3000X950g,
	"SUBCORE_STANDARD_VM_FIXED3025_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3025X950g,
	"SUBCORE_STANDARD_VM_FIXED3050_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3050X950g,
	"SUBCORE_STANDARD_VM_FIXED3075_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3075X950g,
	"SUBCORE_STANDARD_VM_FIXED3100_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3100X950g,
	"SUBCORE_STANDARD_VM_FIXED3125_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3125X950g,
	"SUBCORE_STANDARD_VM_FIXED3150_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3150X950g,
	"SUBCORE_STANDARD_VM_FIXED3200_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3200X950g,
	"SUBCORE_STANDARD_VM_FIXED3225_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3225X950g,
	"SUBCORE_STANDARD_VM_FIXED3250_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3250X950g,
	"SUBCORE_STANDARD_VM_FIXED3300_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3300X950g,
	"SUBCORE_STANDARD_VM_FIXED3325_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3325X950g,
	"SUBCORE_STANDARD_VM_FIXED3375_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3375X950g,
	"SUBCORE_STANDARD_VM_FIXED3400_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3400X950g,
	"SUBCORE_STANDARD_VM_FIXED3450_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3450X950g,
	"SUBCORE_STANDARD_VM_FIXED3500_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3500X950g,
	"SUBCORE_STANDARD_VM_FIXED3525_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3525X950g,
	"SUBCORE_STANDARD_VM_FIXED3575_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3575X950g,
	"SUBCORE_STANDARD_VM_FIXED3600_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3600X950g,
	"SUBCORE_STANDARD_VM_FIXED3625_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3625X950g,
	"SUBCORE_STANDARD_VM_FIXED3675_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3675X950g,
	"SUBCORE_STANDARD_VM_FIXED3700_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3700X950g,
	"SUBCORE_STANDARD_VM_FIXED3750_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3750X950g,
	"SUBCORE_STANDARD_VM_FIXED3800_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3800X950g,
	"SUBCORE_STANDARD_VM_FIXED3825_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3825X950g,
	"SUBCORE_STANDARD_VM_FIXED3850_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3850X950g,
	"SUBCORE_STANDARD_VM_FIXED3875_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3875X950g,
	"SUBCORE_STANDARD_VM_FIXED3900_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3900X950g,
	"SUBCORE_STANDARD_VM_FIXED3975_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3975X950g,
	"SUBCORE_STANDARD_VM_FIXED4000_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4000X950g,
	"SUBCORE_STANDARD_VM_FIXED4025_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4025X950g,
	"SUBCORE_STANDARD_VM_FIXED4050_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4050X950g,
	"SUBCORE_STANDARD_VM_FIXED4100_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4100X950g,
	"SUBCORE_STANDARD_VM_FIXED4125_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4125X950g,
	"SUBCORE_STANDARD_VM_FIXED4200_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4200X950g,
	"SUBCORE_STANDARD_VM_FIXED4225_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4225X950g,
	"SUBCORE_STANDARD_VM_FIXED4250_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4250X950g,
	"SUBCORE_STANDARD_VM_FIXED4275_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4275X950g,
	"SUBCORE_STANDARD_VM_FIXED4300_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4300X950g,
	"SUBCORE_STANDARD_VM_FIXED4350_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4350X950g,
	"SUBCORE_STANDARD_VM_FIXED4375_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4375X950g,
	"SUBCORE_STANDARD_VM_FIXED4400_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4400X950g,
	"SUBCORE_STANDARD_VM_FIXED4425_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4425X950g,
	"SUBCORE_STANDARD_VM_FIXED4500_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4500X950g,
	"SUBCORE_STANDARD_VM_FIXED4550_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4550X950g,
	"SUBCORE_STANDARD_VM_FIXED4575_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4575X950g,
	"SUBCORE_STANDARD_VM_FIXED4600_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4600X950g,
	"SUBCORE_STANDARD_VM_FIXED4625_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4625X950g,
	"SUBCORE_STANDARD_VM_FIXED4650_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4650X950g,
	"SUBCORE_STANDARD_VM_FIXED4675_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4675X950g,
	"SUBCORE_STANDARD_VM_FIXED4700_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4700X950g,
	"SUBCORE_STANDARD_VM_FIXED4725_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4725X950g,
	"SUBCORE_STANDARD_VM_FIXED4750_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4750X950g,
	"SUBCORE_STANDARD_VM_FIXED4800_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4800X950g,
	"SUBCORE_STANDARD_VM_FIXED4875_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4875X950g,
	"SUBCORE_STANDARD_VM_FIXED4900_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4900X950g,
	"SUBCORE_STANDARD_VM_FIXED4950_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4950X950g,
	"SUBCORE_STANDARD_VM_FIXED5000_X9_50G": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed5000X950g,
	"ENTIREHOST_X9_50G":                    CreateInternalVnicAttachmentDetailsVnicShapeEntirehostX950g,
	"DYNAMIC_X9_100G":                      CreateInternalVnicAttachmentDetailsVnicShapeDynamicX9100g,
	"DYNAMIC_X10_50G":                      CreateInternalVnicAttachmentDetailsVnicShapeDynamicX1050g,
	"DYNAMIC_X10_100G":                     CreateInternalVnicAttachmentDetailsVnicShapeDynamicX10100g,
}

var mappingCreateInternalVnicAttachmentDetailsVnicShapeEnumLowerCase = map[string]CreateInternalVnicAttachmentDetailsVnicShapeEnum{
	"dynamic":                              CreateInternalVnicAttachmentDetailsVnicShapeDynamic,
	"fixed0040":                            CreateInternalVnicAttachmentDetailsVnicShapeFixed0040,
	"fixed0060":                            CreateInternalVnicAttachmentDetailsVnicShapeFixed0060,
	"fixed0060_psm":                        CreateInternalVnicAttachmentDetailsVnicShapeFixed0060Psm,
	"fixed0100":                            CreateInternalVnicAttachmentDetailsVnicShapeFixed0100,
	"fixed0120":                            CreateInternalVnicAttachmentDetailsVnicShapeFixed0120,
	"fixed0120_2x":                         CreateInternalVnicAttachmentDetailsVnicShapeFixed01202x,
	"fixed0200":                            CreateInternalVnicAttachmentDetailsVnicShapeFixed0200,
	"fixed0240":                            CreateInternalVnicAttachmentDetailsVnicShapeFixed0240,
	"fixed0480":                            CreateInternalVnicAttachmentDetailsVnicShapeFixed0480,
	"entirehost":                           CreateInternalVnicAttachmentDetailsVnicShapeEntirehost,
	"dynamic_25g":                          CreateInternalVnicAttachmentDetailsVnicShapeDynamic25g,
	"fixed0040_25g":                        CreateInternalVnicAttachmentDetailsVnicShapeFixed004025g,
	"fixed0100_25g":                        CreateInternalVnicAttachmentDetailsVnicShapeFixed010025g,
	"fixed0200_25g":                        CreateInternalVnicAttachmentDetailsVnicShapeFixed020025g,
	"fixed0400_25g":                        CreateInternalVnicAttachmentDetailsVnicShapeFixed040025g,
	"fixed0800_25g":                        CreateInternalVnicAttachmentDetailsVnicShapeFixed080025g,
	"fixed1600_25g":                        CreateInternalVnicAttachmentDetailsVnicShapeFixed160025g,
	"fixed2400_25g":                        CreateInternalVnicAttachmentDetailsVnicShapeFixed240025g,
	"entirehost_25g":                       CreateInternalVnicAttachmentDetailsVnicShapeEntirehost25g,
	"dynamic_e1_25g":                       CreateInternalVnicAttachmentDetailsVnicShapeDynamicE125g,
	"fixed0040_e1_25g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0040E125g,
	"fixed0070_e1_25g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0070E125g,
	"fixed0140_e1_25g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0140E125g,
	"fixed0280_e1_25g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0280E125g,
	"fixed0560_e1_25g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0560E125g,
	"fixed1120_e1_25g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1120E125g,
	"fixed1680_e1_25g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1680E125g,
	"entirehost_e1_25g":                    CreateInternalVnicAttachmentDetailsVnicShapeEntirehostE125g,
	"dynamic_b1_25g":                       CreateInternalVnicAttachmentDetailsVnicShapeDynamicB125g,
	"fixed0040_b1_25g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0040B125g,
	"fixed0060_b1_25g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0060B125g,
	"fixed0120_b1_25g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0120B125g,
	"fixed0240_b1_25g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0240B125g,
	"fixed0480_b1_25g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0480B125g,
	"fixed0960_b1_25g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0960B125g,
	"entirehost_b1_25g":                    CreateInternalVnicAttachmentDetailsVnicShapeEntirehostB125g,
	"micro_vm_fixed0048_e1_25g":            CreateInternalVnicAttachmentDetailsVnicShapeMicroVmFixed0048E125g,
	"micro_lb_fixed0001_e1_25g":            CreateInternalVnicAttachmentDetailsVnicShapeMicroLbFixed0001E125g,
	"vnicaas_fixed0200":                    CreateInternalVnicAttachmentDetailsVnicShapeVnicaasFixed0200,
	"vnicaas_fixed0400":                    CreateInternalVnicAttachmentDetailsVnicShapeVnicaasFixed0400,
	"vnicaas_fixed0700":                    CreateInternalVnicAttachmentDetailsVnicShapeVnicaasFixed0700,
	"vnicaas_nlb_approved_10g":             CreateInternalVnicAttachmentDetailsVnicShapeVnicaasNlbApproved10g,
	"vnicaas_nlb_approved_25g":             CreateInternalVnicAttachmentDetailsVnicShapeVnicaasNlbApproved25g,
	"vnicaas_telesis_25g":                  CreateInternalVnicAttachmentDetailsVnicShapeVnicaasTelesis25g,
	"vnicaas_telesis_10g":                  CreateInternalVnicAttachmentDetailsVnicShapeVnicaasTelesis10g,
	"vnicaas_ambassador_fixed0100":         CreateInternalVnicAttachmentDetailsVnicShapeVnicaasAmbassadorFixed0100,
	"vnicaas_telesis_gamma":                CreateInternalVnicAttachmentDetailsVnicShapeVnicaasTelesisGamma,
	"vnicaas_privatedns":                   CreateInternalVnicAttachmentDetailsVnicShapeVnicaasPrivatedns,
	"vnicaas_fwaas":                        CreateInternalVnicAttachmentDetailsVnicShapeVnicaasFwaas,
	"vnicaas_lbaas_free":                   CreateInternalVnicAttachmentDetailsVnicShapeVnicaasLbaasFree,
	"vnicaas_lbaas_8g_512k":                CreateInternalVnicAttachmentDetailsVnicShapeVnicaasLbaas8g512k,
	"vnicaas_lbaas_8g_1m":                  CreateInternalVnicAttachmentDetailsVnicShapeVnicaasLbaas8g1m,
	"vnicaas_lbaas_8g_2m":                  CreateInternalVnicAttachmentDetailsVnicShapeVnicaasLbaas8g2m,
	"vnicaas_lbaas_8g_3m":                  CreateInternalVnicAttachmentDetailsVnicShapeVnicaasLbaas8g3m,
	"vnicaas_lbaas_8g_1m_8ghost":           CreateInternalVnicAttachmentDetailsVnicShapeVnicaasLbaas8g1m8ghost,
	"vnicaas_lbaas_8g_1m_16ghost":          CreateInternalVnicAttachmentDetailsVnicShapeVnicaasLbaas8g1m16ghost,
	"dynamic_e3_50g":                       CreateInternalVnicAttachmentDetailsVnicShapeDynamicE350g,
	"fixed0040_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0040E350g,
	"fixed0100_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0100E350g,
	"fixed0200_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0200E350g,
	"fixed0300_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0300E350g,
	"fixed0400_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0400E350g,
	"fixed0500_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0500E350g,
	"fixed0600_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0600E350g,
	"fixed0700_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0700E350g,
	"fixed0800_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0800E350g,
	"fixed0900_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0900E350g,
	"fixed1000_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1000E350g,
	"fixed1100_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1100E350g,
	"fixed1200_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1200E350g,
	"fixed1300_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1300E350g,
	"fixed1400_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1400E350g,
	"fixed1500_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1500E350g,
	"fixed1600_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1600E350g,
	"fixed1700_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1700E350g,
	"fixed1800_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1800E350g,
	"fixed1900_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1900E350g,
	"fixed2000_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2000E350g,
	"fixed2100_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2100E350g,
	"fixed2200_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2200E350g,
	"fixed2300_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2300E350g,
	"fixed2400_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2400E350g,
	"fixed2500_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2500E350g,
	"fixed2600_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2600E350g,
	"fixed2700_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2700E350g,
	"fixed2800_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2800E350g,
	"fixed2900_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2900E350g,
	"fixed3000_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3000E350g,
	"fixed3100_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3100E350g,
	"fixed3200_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3200E350g,
	"fixed3300_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3300E350g,
	"fixed3400_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3400E350g,
	"fixed3500_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3500E350g,
	"fixed3600_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3600E350g,
	"fixed3700_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3700E350g,
	"fixed3800_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3800E350g,
	"fixed3900_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3900E350g,
	"fixed4000_e3_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed4000E350g,
	"entirehost_e3_50g":                    CreateInternalVnicAttachmentDetailsVnicShapeEntirehostE350g,
	"dynamic_e4_50g":                       CreateInternalVnicAttachmentDetailsVnicShapeDynamicE450g,
	"fixed0040_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0040E450g,
	"fixed0100_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0100E450g,
	"fixed0200_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0200E450g,
	"fixed0300_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0300E450g,
	"fixed0400_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0400E450g,
	"fixed0500_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0500E450g,
	"fixed0600_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0600E450g,
	"fixed0700_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0700E450g,
	"fixed0800_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0800E450g,
	"fixed0900_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0900E450g,
	"fixed1000_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1000E450g,
	"fixed1100_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1100E450g,
	"fixed1200_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1200E450g,
	"fixed1300_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1300E450g,
	"fixed1400_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1400E450g,
	"fixed1500_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1500E450g,
	"fixed1600_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1600E450g,
	"fixed1700_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1700E450g,
	"fixed1800_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1800E450g,
	"fixed1900_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1900E450g,
	"fixed2000_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2000E450g,
	"fixed2100_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2100E450g,
	"fixed2200_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2200E450g,
	"fixed2300_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2300E450g,
	"fixed2400_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2400E450g,
	"fixed2500_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2500E450g,
	"fixed2600_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2600E450g,
	"fixed2700_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2700E450g,
	"fixed2800_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2800E450g,
	"fixed2900_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2900E450g,
	"fixed3000_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3000E450g,
	"fixed3100_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3100E450g,
	"fixed3200_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3200E450g,
	"fixed3300_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3300E450g,
	"fixed3400_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3400E450g,
	"fixed3500_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3500E450g,
	"fixed3600_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3600E450g,
	"fixed3700_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3700E450g,
	"fixed3800_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3800E450g,
	"fixed3900_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3900E450g,
	"fixed4000_e4_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed4000E450g,
	"entirehost_e4_50g":                    CreateInternalVnicAttachmentDetailsVnicShapeEntirehostE450g,
	"micro_vm_fixed0050_e3_50g":            CreateInternalVnicAttachmentDetailsVnicShapeMicroVmFixed0050E350g,
	"micro_vm_fixed0050_e4_50g":            CreateInternalVnicAttachmentDetailsVnicShapeMicroVmFixed0050E450g,
	"subcore_vm_fixed0025_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0025E350g,
	"subcore_vm_fixed0050_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0050E350g,
	"subcore_vm_fixed0075_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0075E350g,
	"subcore_vm_fixed0100_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0100E350g,
	"subcore_vm_fixed0125_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0125E350g,
	"subcore_vm_fixed0150_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0150E350g,
	"subcore_vm_fixed0175_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0175E350g,
	"subcore_vm_fixed0200_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0200E350g,
	"subcore_vm_fixed0225_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0225E350g,
	"subcore_vm_fixed0250_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0250E350g,
	"subcore_vm_fixed0275_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0275E350g,
	"subcore_vm_fixed0300_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0300E350g,
	"subcore_vm_fixed0325_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0325E350g,
	"subcore_vm_fixed0350_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0350E350g,
	"subcore_vm_fixed0375_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0375E350g,
	"subcore_vm_fixed0400_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0400E350g,
	"subcore_vm_fixed0425_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0425E350g,
	"subcore_vm_fixed0450_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0450E350g,
	"subcore_vm_fixed0475_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0475E350g,
	"subcore_vm_fixed0500_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0500E350g,
	"subcore_vm_fixed0525_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0525E350g,
	"subcore_vm_fixed0550_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0550E350g,
	"subcore_vm_fixed0575_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0575E350g,
	"subcore_vm_fixed0600_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0600E350g,
	"subcore_vm_fixed0625_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0625E350g,
	"subcore_vm_fixed0650_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0650E350g,
	"subcore_vm_fixed0675_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0675E350g,
	"subcore_vm_fixed0700_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0700E350g,
	"subcore_vm_fixed0725_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0725E350g,
	"subcore_vm_fixed0750_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0750E350g,
	"subcore_vm_fixed0775_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0775E350g,
	"subcore_vm_fixed0800_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0800E350g,
	"subcore_vm_fixed0825_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0825E350g,
	"subcore_vm_fixed0850_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0850E350g,
	"subcore_vm_fixed0875_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0875E350g,
	"subcore_vm_fixed0900_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0900E350g,
	"subcore_vm_fixed0925_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0925E350g,
	"subcore_vm_fixed0950_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0950E350g,
	"subcore_vm_fixed0975_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0975E350g,
	"subcore_vm_fixed1000_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1000E350g,
	"subcore_vm_fixed1025_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1025E350g,
	"subcore_vm_fixed1050_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1050E350g,
	"subcore_vm_fixed1075_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1075E350g,
	"subcore_vm_fixed1100_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1100E350g,
	"subcore_vm_fixed1125_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1125E350g,
	"subcore_vm_fixed1150_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1150E350g,
	"subcore_vm_fixed1175_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1175E350g,
	"subcore_vm_fixed1200_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1200E350g,
	"subcore_vm_fixed1225_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1225E350g,
	"subcore_vm_fixed1250_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1250E350g,
	"subcore_vm_fixed1275_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1275E350g,
	"subcore_vm_fixed1300_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1300E350g,
	"subcore_vm_fixed1325_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1325E350g,
	"subcore_vm_fixed1350_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1350E350g,
	"subcore_vm_fixed1375_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1375E350g,
	"subcore_vm_fixed1400_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1400E350g,
	"subcore_vm_fixed1425_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1425E350g,
	"subcore_vm_fixed1450_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1450E350g,
	"subcore_vm_fixed1475_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1475E350g,
	"subcore_vm_fixed1500_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1500E350g,
	"subcore_vm_fixed1525_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1525E350g,
	"subcore_vm_fixed1550_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1550E350g,
	"subcore_vm_fixed1575_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1575E350g,
	"subcore_vm_fixed1600_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1600E350g,
	"subcore_vm_fixed1625_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1625E350g,
	"subcore_vm_fixed1650_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1650E350g,
	"subcore_vm_fixed1700_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1700E350g,
	"subcore_vm_fixed1725_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1725E350g,
	"subcore_vm_fixed1750_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1750E350g,
	"subcore_vm_fixed1800_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1800E350g,
	"subcore_vm_fixed1850_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1850E350g,
	"subcore_vm_fixed1875_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1875E350g,
	"subcore_vm_fixed1900_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1900E350g,
	"subcore_vm_fixed1925_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1925E350g,
	"subcore_vm_fixed1950_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1950E350g,
	"subcore_vm_fixed2000_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2000E350g,
	"subcore_vm_fixed2025_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2025E350g,
	"subcore_vm_fixed2050_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2050E350g,
	"subcore_vm_fixed2100_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2100E350g,
	"subcore_vm_fixed2125_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2125E350g,
	"subcore_vm_fixed2150_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2150E350g,
	"subcore_vm_fixed2175_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2175E350g,
	"subcore_vm_fixed2200_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2200E350g,
	"subcore_vm_fixed2250_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2250E350g,
	"subcore_vm_fixed2275_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2275E350g,
	"subcore_vm_fixed2300_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2300E350g,
	"subcore_vm_fixed2325_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2325E350g,
	"subcore_vm_fixed2350_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2350E350g,
	"subcore_vm_fixed2375_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2375E350g,
	"subcore_vm_fixed2400_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2400E350g,
	"subcore_vm_fixed2450_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2450E350g,
	"subcore_vm_fixed2475_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2475E350g,
	"subcore_vm_fixed2500_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2500E350g,
	"subcore_vm_fixed2550_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2550E350g,
	"subcore_vm_fixed2600_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2600E350g,
	"subcore_vm_fixed2625_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2625E350g,
	"subcore_vm_fixed2650_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2650E350g,
	"subcore_vm_fixed2700_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2700E350g,
	"subcore_vm_fixed2750_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2750E350g,
	"subcore_vm_fixed2775_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2775E350g,
	"subcore_vm_fixed2800_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2800E350g,
	"subcore_vm_fixed2850_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2850E350g,
	"subcore_vm_fixed2875_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2875E350g,
	"subcore_vm_fixed2900_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2900E350g,
	"subcore_vm_fixed2925_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2925E350g,
	"subcore_vm_fixed2950_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2950E350g,
	"subcore_vm_fixed2975_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2975E350g,
	"subcore_vm_fixed3000_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3000E350g,
	"subcore_vm_fixed3025_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3025E350g,
	"subcore_vm_fixed3050_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3050E350g,
	"subcore_vm_fixed3075_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3075E350g,
	"subcore_vm_fixed3100_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3100E350g,
	"subcore_vm_fixed3125_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3125E350g,
	"subcore_vm_fixed3150_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3150E350g,
	"subcore_vm_fixed3200_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3200E350g,
	"subcore_vm_fixed3225_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3225E350g,
	"subcore_vm_fixed3250_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3250E350g,
	"subcore_vm_fixed3300_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3300E350g,
	"subcore_vm_fixed3325_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3325E350g,
	"subcore_vm_fixed3375_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3375E350g,
	"subcore_vm_fixed3400_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3400E350g,
	"subcore_vm_fixed3450_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3450E350g,
	"subcore_vm_fixed3500_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3500E350g,
	"subcore_vm_fixed3525_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3525E350g,
	"subcore_vm_fixed3575_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3575E350g,
	"subcore_vm_fixed3600_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3600E350g,
	"subcore_vm_fixed3625_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3625E350g,
	"subcore_vm_fixed3675_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3675E350g,
	"subcore_vm_fixed3700_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3700E350g,
	"subcore_vm_fixed3750_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3750E350g,
	"subcore_vm_fixed3800_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3800E350g,
	"subcore_vm_fixed3825_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3825E350g,
	"subcore_vm_fixed3850_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3850E350g,
	"subcore_vm_fixed3875_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3875E350g,
	"subcore_vm_fixed3900_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3900E350g,
	"subcore_vm_fixed3975_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3975E350g,
	"subcore_vm_fixed4000_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4000E350g,
	"subcore_vm_fixed4025_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4025E350g,
	"subcore_vm_fixed4050_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4050E350g,
	"subcore_vm_fixed4100_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4100E350g,
	"subcore_vm_fixed4125_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4125E350g,
	"subcore_vm_fixed4200_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4200E350g,
	"subcore_vm_fixed4225_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4225E350g,
	"subcore_vm_fixed4250_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4250E350g,
	"subcore_vm_fixed4275_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4275E350g,
	"subcore_vm_fixed4300_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4300E350g,
	"subcore_vm_fixed4350_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4350E350g,
	"subcore_vm_fixed4375_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4375E350g,
	"subcore_vm_fixed4400_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4400E350g,
	"subcore_vm_fixed4425_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4425E350g,
	"subcore_vm_fixed4500_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4500E350g,
	"subcore_vm_fixed4550_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4550E350g,
	"subcore_vm_fixed4575_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4575E350g,
	"subcore_vm_fixed4600_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4600E350g,
	"subcore_vm_fixed4625_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4625E350g,
	"subcore_vm_fixed4650_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4650E350g,
	"subcore_vm_fixed4675_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4675E350g,
	"subcore_vm_fixed4700_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4700E350g,
	"subcore_vm_fixed4725_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4725E350g,
	"subcore_vm_fixed4750_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4750E350g,
	"subcore_vm_fixed4800_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4800E350g,
	"subcore_vm_fixed4875_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4875E350g,
	"subcore_vm_fixed4900_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4900E350g,
	"subcore_vm_fixed4950_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4950E350g,
	"subcore_vm_fixed5000_e3_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed5000E350g,
	"subcore_vm_fixed0025_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0025E450g,
	"subcore_vm_fixed0050_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0050E450g,
	"subcore_vm_fixed0075_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0075E450g,
	"subcore_vm_fixed0100_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0100E450g,
	"subcore_vm_fixed0125_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0125E450g,
	"subcore_vm_fixed0150_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0150E450g,
	"subcore_vm_fixed0175_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0175E450g,
	"subcore_vm_fixed0200_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0200E450g,
	"subcore_vm_fixed0225_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0225E450g,
	"subcore_vm_fixed0250_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0250E450g,
	"subcore_vm_fixed0275_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0275E450g,
	"subcore_vm_fixed0300_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0300E450g,
	"subcore_vm_fixed0325_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0325E450g,
	"subcore_vm_fixed0350_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0350E450g,
	"subcore_vm_fixed0375_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0375E450g,
	"subcore_vm_fixed0400_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0400E450g,
	"subcore_vm_fixed0425_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0425E450g,
	"subcore_vm_fixed0450_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0450E450g,
	"subcore_vm_fixed0475_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0475E450g,
	"subcore_vm_fixed0500_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0500E450g,
	"subcore_vm_fixed0525_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0525E450g,
	"subcore_vm_fixed0550_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0550E450g,
	"subcore_vm_fixed0575_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0575E450g,
	"subcore_vm_fixed0600_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0600E450g,
	"subcore_vm_fixed0625_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0625E450g,
	"subcore_vm_fixed0650_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0650E450g,
	"subcore_vm_fixed0675_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0675E450g,
	"subcore_vm_fixed0700_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0700E450g,
	"subcore_vm_fixed0725_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0725E450g,
	"subcore_vm_fixed0750_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0750E450g,
	"subcore_vm_fixed0775_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0775E450g,
	"subcore_vm_fixed0800_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0800E450g,
	"subcore_vm_fixed0825_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0825E450g,
	"subcore_vm_fixed0850_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0850E450g,
	"subcore_vm_fixed0875_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0875E450g,
	"subcore_vm_fixed0900_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0900E450g,
	"subcore_vm_fixed0925_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0925E450g,
	"subcore_vm_fixed0950_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0950E450g,
	"subcore_vm_fixed0975_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0975E450g,
	"subcore_vm_fixed1000_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1000E450g,
	"subcore_vm_fixed1025_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1025E450g,
	"subcore_vm_fixed1050_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1050E450g,
	"subcore_vm_fixed1075_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1075E450g,
	"subcore_vm_fixed1100_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1100E450g,
	"subcore_vm_fixed1125_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1125E450g,
	"subcore_vm_fixed1150_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1150E450g,
	"subcore_vm_fixed1175_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1175E450g,
	"subcore_vm_fixed1200_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1200E450g,
	"subcore_vm_fixed1225_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1225E450g,
	"subcore_vm_fixed1250_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1250E450g,
	"subcore_vm_fixed1275_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1275E450g,
	"subcore_vm_fixed1300_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1300E450g,
	"subcore_vm_fixed1325_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1325E450g,
	"subcore_vm_fixed1350_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1350E450g,
	"subcore_vm_fixed1375_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1375E450g,
	"subcore_vm_fixed1400_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1400E450g,
	"subcore_vm_fixed1425_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1425E450g,
	"subcore_vm_fixed1450_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1450E450g,
	"subcore_vm_fixed1475_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1475E450g,
	"subcore_vm_fixed1500_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1500E450g,
	"subcore_vm_fixed1525_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1525E450g,
	"subcore_vm_fixed1550_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1550E450g,
	"subcore_vm_fixed1575_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1575E450g,
	"subcore_vm_fixed1600_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1600E450g,
	"subcore_vm_fixed1625_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1625E450g,
	"subcore_vm_fixed1650_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1650E450g,
	"subcore_vm_fixed1700_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1700E450g,
	"subcore_vm_fixed1725_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1725E450g,
	"subcore_vm_fixed1750_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1750E450g,
	"subcore_vm_fixed1800_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1800E450g,
	"subcore_vm_fixed1850_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1850E450g,
	"subcore_vm_fixed1875_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1875E450g,
	"subcore_vm_fixed1900_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1900E450g,
	"subcore_vm_fixed1925_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1925E450g,
	"subcore_vm_fixed1950_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1950E450g,
	"subcore_vm_fixed2000_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2000E450g,
	"subcore_vm_fixed2025_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2025E450g,
	"subcore_vm_fixed2050_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2050E450g,
	"subcore_vm_fixed2100_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2100E450g,
	"subcore_vm_fixed2125_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2125E450g,
	"subcore_vm_fixed2150_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2150E450g,
	"subcore_vm_fixed2175_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2175E450g,
	"subcore_vm_fixed2200_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2200E450g,
	"subcore_vm_fixed2250_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2250E450g,
	"subcore_vm_fixed2275_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2275E450g,
	"subcore_vm_fixed2300_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2300E450g,
	"subcore_vm_fixed2325_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2325E450g,
	"subcore_vm_fixed2350_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2350E450g,
	"subcore_vm_fixed2375_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2375E450g,
	"subcore_vm_fixed2400_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2400E450g,
	"subcore_vm_fixed2450_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2450E450g,
	"subcore_vm_fixed2475_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2475E450g,
	"subcore_vm_fixed2500_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2500E450g,
	"subcore_vm_fixed2550_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2550E450g,
	"subcore_vm_fixed2600_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2600E450g,
	"subcore_vm_fixed2625_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2625E450g,
	"subcore_vm_fixed2650_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2650E450g,
	"subcore_vm_fixed2700_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2700E450g,
	"subcore_vm_fixed2750_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2750E450g,
	"subcore_vm_fixed2775_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2775E450g,
	"subcore_vm_fixed2800_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2800E450g,
	"subcore_vm_fixed2850_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2850E450g,
	"subcore_vm_fixed2875_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2875E450g,
	"subcore_vm_fixed2900_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2900E450g,
	"subcore_vm_fixed2925_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2925E450g,
	"subcore_vm_fixed2950_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2950E450g,
	"subcore_vm_fixed2975_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2975E450g,
	"subcore_vm_fixed3000_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3000E450g,
	"subcore_vm_fixed3025_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3025E450g,
	"subcore_vm_fixed3050_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3050E450g,
	"subcore_vm_fixed3075_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3075E450g,
	"subcore_vm_fixed3100_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3100E450g,
	"subcore_vm_fixed3125_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3125E450g,
	"subcore_vm_fixed3150_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3150E450g,
	"subcore_vm_fixed3200_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3200E450g,
	"subcore_vm_fixed3225_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3225E450g,
	"subcore_vm_fixed3250_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3250E450g,
	"subcore_vm_fixed3300_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3300E450g,
	"subcore_vm_fixed3325_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3325E450g,
	"subcore_vm_fixed3375_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3375E450g,
	"subcore_vm_fixed3400_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3400E450g,
	"subcore_vm_fixed3450_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3450E450g,
	"subcore_vm_fixed3500_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3500E450g,
	"subcore_vm_fixed3525_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3525E450g,
	"subcore_vm_fixed3575_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3575E450g,
	"subcore_vm_fixed3600_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3600E450g,
	"subcore_vm_fixed3625_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3625E450g,
	"subcore_vm_fixed3675_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3675E450g,
	"subcore_vm_fixed3700_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3700E450g,
	"subcore_vm_fixed3750_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3750E450g,
	"subcore_vm_fixed3800_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3800E450g,
	"subcore_vm_fixed3825_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3825E450g,
	"subcore_vm_fixed3850_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3850E450g,
	"subcore_vm_fixed3875_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3875E450g,
	"subcore_vm_fixed3900_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3900E450g,
	"subcore_vm_fixed3975_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3975E450g,
	"subcore_vm_fixed4000_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4000E450g,
	"subcore_vm_fixed4025_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4025E450g,
	"subcore_vm_fixed4050_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4050E450g,
	"subcore_vm_fixed4100_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4100E450g,
	"subcore_vm_fixed4125_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4125E450g,
	"subcore_vm_fixed4200_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4200E450g,
	"subcore_vm_fixed4225_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4225E450g,
	"subcore_vm_fixed4250_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4250E450g,
	"subcore_vm_fixed4275_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4275E450g,
	"subcore_vm_fixed4300_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4300E450g,
	"subcore_vm_fixed4350_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4350E450g,
	"subcore_vm_fixed4375_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4375E450g,
	"subcore_vm_fixed4400_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4400E450g,
	"subcore_vm_fixed4425_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4425E450g,
	"subcore_vm_fixed4500_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4500E450g,
	"subcore_vm_fixed4550_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4550E450g,
	"subcore_vm_fixed4575_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4575E450g,
	"subcore_vm_fixed4600_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4600E450g,
	"subcore_vm_fixed4625_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4625E450g,
	"subcore_vm_fixed4650_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4650E450g,
	"subcore_vm_fixed4675_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4675E450g,
	"subcore_vm_fixed4700_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4700E450g,
	"subcore_vm_fixed4725_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4725E450g,
	"subcore_vm_fixed4750_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4750E450g,
	"subcore_vm_fixed4800_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4800E450g,
	"subcore_vm_fixed4875_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4875E450g,
	"subcore_vm_fixed4900_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4900E450g,
	"subcore_vm_fixed4950_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4950E450g,
	"subcore_vm_fixed5000_e4_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed5000E450g,
	"dynamic_e5_50g":                       CreateInternalVnicAttachmentDetailsVnicShapeDynamicE550g,
	"dynamic_e5_100g":                      CreateInternalVnicAttachmentDetailsVnicShapeDynamicE5100g,
	"subcore_vm_fixed0020_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0020A150g,
	"subcore_vm_fixed0040_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0040A150g,
	"subcore_vm_fixed0060_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0060A150g,
	"subcore_vm_fixed0080_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0080A150g,
	"subcore_vm_fixed0100_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0100A150g,
	"subcore_vm_fixed0120_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0120A150g,
	"subcore_vm_fixed0140_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0140A150g,
	"subcore_vm_fixed0160_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0160A150g,
	"subcore_vm_fixed0180_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0180A150g,
	"subcore_vm_fixed0200_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0200A150g,
	"subcore_vm_fixed0220_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0220A150g,
	"subcore_vm_fixed0240_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0240A150g,
	"subcore_vm_fixed0260_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0260A150g,
	"subcore_vm_fixed0280_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0280A150g,
	"subcore_vm_fixed0300_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0300A150g,
	"subcore_vm_fixed0320_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0320A150g,
	"subcore_vm_fixed0340_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0340A150g,
	"subcore_vm_fixed0360_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0360A150g,
	"subcore_vm_fixed0380_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0380A150g,
	"subcore_vm_fixed0400_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0400A150g,
	"subcore_vm_fixed0420_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0420A150g,
	"subcore_vm_fixed0440_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0440A150g,
	"subcore_vm_fixed0460_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0460A150g,
	"subcore_vm_fixed0480_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0480A150g,
	"subcore_vm_fixed0500_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0500A150g,
	"subcore_vm_fixed0520_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0520A150g,
	"subcore_vm_fixed0540_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0540A150g,
	"subcore_vm_fixed0560_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0560A150g,
	"subcore_vm_fixed0580_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0580A150g,
	"subcore_vm_fixed0600_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0600A150g,
	"subcore_vm_fixed0620_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0620A150g,
	"subcore_vm_fixed0640_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0640A150g,
	"subcore_vm_fixed0660_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0660A150g,
	"subcore_vm_fixed0680_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0680A150g,
	"subcore_vm_fixed0700_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0700A150g,
	"subcore_vm_fixed0720_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0720A150g,
	"subcore_vm_fixed0740_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0740A150g,
	"subcore_vm_fixed0760_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0760A150g,
	"subcore_vm_fixed0780_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0780A150g,
	"subcore_vm_fixed0800_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0800A150g,
	"subcore_vm_fixed0820_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0820A150g,
	"subcore_vm_fixed0840_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0840A150g,
	"subcore_vm_fixed0860_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0860A150g,
	"subcore_vm_fixed0880_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0880A150g,
	"subcore_vm_fixed0900_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0900A150g,
	"subcore_vm_fixed0920_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0920A150g,
	"subcore_vm_fixed0940_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0940A150g,
	"subcore_vm_fixed0960_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0960A150g,
	"subcore_vm_fixed0980_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0980A150g,
	"subcore_vm_fixed1000_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1000A150g,
	"subcore_vm_fixed1020_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1020A150g,
	"subcore_vm_fixed1040_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1040A150g,
	"subcore_vm_fixed1060_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1060A150g,
	"subcore_vm_fixed1080_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1080A150g,
	"subcore_vm_fixed1100_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1100A150g,
	"subcore_vm_fixed1120_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1120A150g,
	"subcore_vm_fixed1140_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1140A150g,
	"subcore_vm_fixed1160_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1160A150g,
	"subcore_vm_fixed1180_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1180A150g,
	"subcore_vm_fixed1200_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1200A150g,
	"subcore_vm_fixed1220_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1220A150g,
	"subcore_vm_fixed1240_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1240A150g,
	"subcore_vm_fixed1260_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1260A150g,
	"subcore_vm_fixed1280_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1280A150g,
	"subcore_vm_fixed1300_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1300A150g,
	"subcore_vm_fixed1320_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1320A150g,
	"subcore_vm_fixed1340_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1340A150g,
	"subcore_vm_fixed1360_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1360A150g,
	"subcore_vm_fixed1380_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1380A150g,
	"subcore_vm_fixed1400_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1400A150g,
	"subcore_vm_fixed1420_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1420A150g,
	"subcore_vm_fixed1440_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1440A150g,
	"subcore_vm_fixed1460_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1460A150g,
	"subcore_vm_fixed1480_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1480A150g,
	"subcore_vm_fixed1500_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1500A150g,
	"subcore_vm_fixed1520_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1520A150g,
	"subcore_vm_fixed1540_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1540A150g,
	"subcore_vm_fixed1560_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1560A150g,
	"subcore_vm_fixed1580_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1580A150g,
	"subcore_vm_fixed1600_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1600A150g,
	"subcore_vm_fixed1620_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1620A150g,
	"subcore_vm_fixed1640_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1640A150g,
	"subcore_vm_fixed1660_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1660A150g,
	"subcore_vm_fixed1680_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1680A150g,
	"subcore_vm_fixed1700_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1700A150g,
	"subcore_vm_fixed1720_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1720A150g,
	"subcore_vm_fixed1740_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1740A150g,
	"subcore_vm_fixed1760_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1760A150g,
	"subcore_vm_fixed1780_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1780A150g,
	"subcore_vm_fixed1800_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1800A150g,
	"subcore_vm_fixed1820_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1820A150g,
	"subcore_vm_fixed1840_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1840A150g,
	"subcore_vm_fixed1860_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1860A150g,
	"subcore_vm_fixed1880_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1880A150g,
	"subcore_vm_fixed1900_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1900A150g,
	"subcore_vm_fixed1920_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1920A150g,
	"subcore_vm_fixed1940_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1940A150g,
	"subcore_vm_fixed1960_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1960A150g,
	"subcore_vm_fixed1980_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1980A150g,
	"subcore_vm_fixed2000_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2000A150g,
	"subcore_vm_fixed2020_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2020A150g,
	"subcore_vm_fixed2040_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2040A150g,
	"subcore_vm_fixed2060_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2060A150g,
	"subcore_vm_fixed2080_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2080A150g,
	"subcore_vm_fixed2100_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2100A150g,
	"subcore_vm_fixed2120_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2120A150g,
	"subcore_vm_fixed2140_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2140A150g,
	"subcore_vm_fixed2160_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2160A150g,
	"subcore_vm_fixed2180_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2180A150g,
	"subcore_vm_fixed2200_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2200A150g,
	"subcore_vm_fixed2220_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2220A150g,
	"subcore_vm_fixed2240_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2240A150g,
	"subcore_vm_fixed2260_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2260A150g,
	"subcore_vm_fixed2280_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2280A150g,
	"subcore_vm_fixed2300_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2300A150g,
	"subcore_vm_fixed2320_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2320A150g,
	"subcore_vm_fixed2340_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2340A150g,
	"subcore_vm_fixed2360_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2360A150g,
	"subcore_vm_fixed2380_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2380A150g,
	"subcore_vm_fixed2400_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2400A150g,
	"subcore_vm_fixed2420_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2420A150g,
	"subcore_vm_fixed2440_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2440A150g,
	"subcore_vm_fixed2460_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2460A150g,
	"subcore_vm_fixed2480_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2480A150g,
	"subcore_vm_fixed2500_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2500A150g,
	"subcore_vm_fixed2520_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2520A150g,
	"subcore_vm_fixed2540_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2540A150g,
	"subcore_vm_fixed2560_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2560A150g,
	"subcore_vm_fixed2580_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2580A150g,
	"subcore_vm_fixed2600_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2600A150g,
	"subcore_vm_fixed2620_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2620A150g,
	"subcore_vm_fixed2640_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2640A150g,
	"subcore_vm_fixed2660_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2660A150g,
	"subcore_vm_fixed2680_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2680A150g,
	"subcore_vm_fixed2700_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2700A150g,
	"subcore_vm_fixed2720_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2720A150g,
	"subcore_vm_fixed2740_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2740A150g,
	"subcore_vm_fixed2760_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2760A150g,
	"subcore_vm_fixed2780_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2780A150g,
	"subcore_vm_fixed2800_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2800A150g,
	"subcore_vm_fixed2820_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2820A150g,
	"subcore_vm_fixed2840_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2840A150g,
	"subcore_vm_fixed2860_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2860A150g,
	"subcore_vm_fixed2880_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2880A150g,
	"subcore_vm_fixed2900_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2900A150g,
	"subcore_vm_fixed2920_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2920A150g,
	"subcore_vm_fixed2940_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2940A150g,
	"subcore_vm_fixed2960_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2960A150g,
	"subcore_vm_fixed2980_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2980A150g,
	"subcore_vm_fixed3000_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3000A150g,
	"subcore_vm_fixed3020_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3020A150g,
	"subcore_vm_fixed3040_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3040A150g,
	"subcore_vm_fixed3060_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3060A150g,
	"subcore_vm_fixed3080_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3080A150g,
	"subcore_vm_fixed3100_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3100A150g,
	"subcore_vm_fixed3120_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3120A150g,
	"subcore_vm_fixed3140_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3140A150g,
	"subcore_vm_fixed3160_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3160A150g,
	"subcore_vm_fixed3180_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3180A150g,
	"subcore_vm_fixed3200_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3200A150g,
	"subcore_vm_fixed3220_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3220A150g,
	"subcore_vm_fixed3240_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3240A150g,
	"subcore_vm_fixed3260_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3260A150g,
	"subcore_vm_fixed3280_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3280A150g,
	"subcore_vm_fixed3300_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3300A150g,
	"subcore_vm_fixed3320_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3320A150g,
	"subcore_vm_fixed3340_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3340A150g,
	"subcore_vm_fixed3360_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3360A150g,
	"subcore_vm_fixed3380_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3380A150g,
	"subcore_vm_fixed3400_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3400A150g,
	"subcore_vm_fixed3420_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3420A150g,
	"subcore_vm_fixed3440_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3440A150g,
	"subcore_vm_fixed3460_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3460A150g,
	"subcore_vm_fixed3480_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3480A150g,
	"subcore_vm_fixed3500_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3500A150g,
	"subcore_vm_fixed3520_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3520A150g,
	"subcore_vm_fixed3540_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3540A150g,
	"subcore_vm_fixed3560_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3560A150g,
	"subcore_vm_fixed3580_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3580A150g,
	"subcore_vm_fixed3600_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3600A150g,
	"subcore_vm_fixed3620_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3620A150g,
	"subcore_vm_fixed3640_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3640A150g,
	"subcore_vm_fixed3660_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3660A150g,
	"subcore_vm_fixed3680_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3680A150g,
	"subcore_vm_fixed3700_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3700A150g,
	"subcore_vm_fixed3720_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3720A150g,
	"subcore_vm_fixed3740_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3740A150g,
	"subcore_vm_fixed3760_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3760A150g,
	"subcore_vm_fixed3780_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3780A150g,
	"subcore_vm_fixed3800_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3800A150g,
	"subcore_vm_fixed3820_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3820A150g,
	"subcore_vm_fixed3840_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3840A150g,
	"subcore_vm_fixed3860_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3860A150g,
	"subcore_vm_fixed3880_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3880A150g,
	"subcore_vm_fixed3900_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3900A150g,
	"subcore_vm_fixed3920_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3920A150g,
	"subcore_vm_fixed3940_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3940A150g,
	"subcore_vm_fixed3960_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3960A150g,
	"subcore_vm_fixed3980_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3980A150g,
	"subcore_vm_fixed4000_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4000A150g,
	"subcore_vm_fixed4020_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4020A150g,
	"subcore_vm_fixed4040_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4040A150g,
	"subcore_vm_fixed4060_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4060A150g,
	"subcore_vm_fixed4080_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4080A150g,
	"subcore_vm_fixed4100_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4100A150g,
	"subcore_vm_fixed4120_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4120A150g,
	"subcore_vm_fixed4140_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4140A150g,
	"subcore_vm_fixed4160_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4160A150g,
	"subcore_vm_fixed4180_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4180A150g,
	"subcore_vm_fixed4200_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4200A150g,
	"subcore_vm_fixed4220_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4220A150g,
	"subcore_vm_fixed4240_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4240A150g,
	"subcore_vm_fixed4260_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4260A150g,
	"subcore_vm_fixed4280_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4280A150g,
	"subcore_vm_fixed4300_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4300A150g,
	"subcore_vm_fixed4320_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4320A150g,
	"subcore_vm_fixed4340_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4340A150g,
	"subcore_vm_fixed4360_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4360A150g,
	"subcore_vm_fixed4380_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4380A150g,
	"subcore_vm_fixed4400_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4400A150g,
	"subcore_vm_fixed4420_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4420A150g,
	"subcore_vm_fixed4440_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4440A150g,
	"subcore_vm_fixed4460_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4460A150g,
	"subcore_vm_fixed4480_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4480A150g,
	"subcore_vm_fixed4500_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4500A150g,
	"subcore_vm_fixed4520_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4520A150g,
	"subcore_vm_fixed4540_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4540A150g,
	"subcore_vm_fixed4560_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4560A150g,
	"subcore_vm_fixed4580_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4580A150g,
	"subcore_vm_fixed4600_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4600A150g,
	"subcore_vm_fixed4620_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4620A150g,
	"subcore_vm_fixed4640_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4640A150g,
	"subcore_vm_fixed4660_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4660A150g,
	"subcore_vm_fixed4680_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4680A150g,
	"subcore_vm_fixed4700_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4700A150g,
	"subcore_vm_fixed4720_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4720A150g,
	"subcore_vm_fixed4740_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4740A150g,
	"subcore_vm_fixed4760_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4760A150g,
	"subcore_vm_fixed4780_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4780A150g,
	"subcore_vm_fixed4800_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4800A150g,
	"subcore_vm_fixed4820_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4820A150g,
	"subcore_vm_fixed4840_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4840A150g,
	"subcore_vm_fixed4860_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4860A150g,
	"subcore_vm_fixed4880_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4880A150g,
	"subcore_vm_fixed4900_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4900A150g,
	"subcore_vm_fixed4920_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4920A150g,
	"subcore_vm_fixed4940_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4940A150g,
	"subcore_vm_fixed4960_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4960A150g,
	"subcore_vm_fixed4980_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4980A150g,
	"subcore_vm_fixed5000_a1_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed5000A150g,
	"subcore_vm_fixed0090_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0090X950g,
	"subcore_vm_fixed0180_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0180X950g,
	"subcore_vm_fixed0270_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0270X950g,
	"subcore_vm_fixed0360_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0360X950g,
	"subcore_vm_fixed0450_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0450X950g,
	"subcore_vm_fixed0540_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0540X950g,
	"subcore_vm_fixed0630_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0630X950g,
	"subcore_vm_fixed0720_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0720X950g,
	"subcore_vm_fixed0810_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0810X950g,
	"subcore_vm_fixed0900_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0900X950g,
	"subcore_vm_fixed0990_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed0990X950g,
	"subcore_vm_fixed1080_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1080X950g,
	"subcore_vm_fixed1170_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1170X950g,
	"subcore_vm_fixed1260_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1260X950g,
	"subcore_vm_fixed1350_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1350X950g,
	"subcore_vm_fixed1440_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1440X950g,
	"subcore_vm_fixed1530_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1530X950g,
	"subcore_vm_fixed1620_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1620X950g,
	"subcore_vm_fixed1710_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1710X950g,
	"subcore_vm_fixed1800_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1800X950g,
	"subcore_vm_fixed1890_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1890X950g,
	"subcore_vm_fixed1980_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed1980X950g,
	"subcore_vm_fixed2070_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2070X950g,
	"subcore_vm_fixed2160_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2160X950g,
	"subcore_vm_fixed2250_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2250X950g,
	"subcore_vm_fixed2340_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2340X950g,
	"subcore_vm_fixed2430_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2430X950g,
	"subcore_vm_fixed2520_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2520X950g,
	"subcore_vm_fixed2610_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2610X950g,
	"subcore_vm_fixed2700_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2700X950g,
	"subcore_vm_fixed2790_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2790X950g,
	"subcore_vm_fixed2880_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2880X950g,
	"subcore_vm_fixed2970_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed2970X950g,
	"subcore_vm_fixed3060_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3060X950g,
	"subcore_vm_fixed3150_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3150X950g,
	"subcore_vm_fixed3240_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3240X950g,
	"subcore_vm_fixed3330_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3330X950g,
	"subcore_vm_fixed3420_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3420X950g,
	"subcore_vm_fixed3510_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3510X950g,
	"subcore_vm_fixed3600_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3600X950g,
	"subcore_vm_fixed3690_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3690X950g,
	"subcore_vm_fixed3780_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3780X950g,
	"subcore_vm_fixed3870_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3870X950g,
	"subcore_vm_fixed3960_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed3960X950g,
	"subcore_vm_fixed4050_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4050X950g,
	"subcore_vm_fixed4140_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4140X950g,
	"subcore_vm_fixed4230_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4230X950g,
	"subcore_vm_fixed4320_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4320X950g,
	"subcore_vm_fixed4410_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4410X950g,
	"subcore_vm_fixed4500_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4500X950g,
	"subcore_vm_fixed4590_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4590X950g,
	"subcore_vm_fixed4680_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4680X950g,
	"subcore_vm_fixed4770_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4770X950g,
	"subcore_vm_fixed4860_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4860X950g,
	"subcore_vm_fixed4950_x9_50g":          CreateInternalVnicAttachmentDetailsVnicShapeSubcoreVmFixed4950X950g,
	"dynamic_a1_50g":                       CreateInternalVnicAttachmentDetailsVnicShapeDynamicA150g,
	"fixed0040_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0040A150g,
	"fixed0100_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0100A150g,
	"fixed0200_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0200A150g,
	"fixed0300_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0300A150g,
	"fixed0400_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0400A150g,
	"fixed0500_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0500A150g,
	"fixed0600_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0600A150g,
	"fixed0700_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0700A150g,
	"fixed0800_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0800A150g,
	"fixed0900_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0900A150g,
	"fixed1000_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1000A150g,
	"fixed1100_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1100A150g,
	"fixed1200_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1200A150g,
	"fixed1300_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1300A150g,
	"fixed1400_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1400A150g,
	"fixed1500_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1500A150g,
	"fixed1600_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1600A150g,
	"fixed1700_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1700A150g,
	"fixed1800_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1800A150g,
	"fixed1900_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1900A150g,
	"fixed2000_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2000A150g,
	"fixed2100_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2100A150g,
	"fixed2200_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2200A150g,
	"fixed2300_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2300A150g,
	"fixed2400_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2400A150g,
	"fixed2500_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2500A150g,
	"fixed2600_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2600A150g,
	"fixed2700_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2700A150g,
	"fixed2800_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2800A150g,
	"fixed2900_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2900A150g,
	"fixed3000_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3000A150g,
	"fixed3100_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3100A150g,
	"fixed3200_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3200A150g,
	"fixed3300_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3300A150g,
	"fixed3400_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3400A150g,
	"fixed3500_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3500A150g,
	"fixed3600_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3600A150g,
	"fixed3700_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3700A150g,
	"fixed3800_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3800A150g,
	"fixed3900_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3900A150g,
	"fixed4000_a1_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed4000A150g,
	"fixed5000_telesis_a1_50g":             CreateInternalVnicAttachmentDetailsVnicShapeFixed5000TelesisA150g,
	"entirehost_a1_50g":                    CreateInternalVnicAttachmentDetailsVnicShapeEntirehostA150g,
	"dynamic_x9_50g":                       CreateInternalVnicAttachmentDetailsVnicShapeDynamicX950g,
	"fixed0040_x9_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0040X950g,
	"fixed0400_x9_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0400X950g,
	"fixed0800_x9_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed0800X950g,
	"fixed1200_x9_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1200X950g,
	"fixed1600_x9_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed1600X950g,
	"fixed2000_x9_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2000X950g,
	"fixed2400_x9_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2400X950g,
	"fixed2800_x9_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed2800X950g,
	"fixed3200_x9_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3200X950g,
	"fixed3600_x9_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed3600X950g,
	"fixed4000_x9_50g":                     CreateInternalVnicAttachmentDetailsVnicShapeFixed4000X950g,
	"standard_vm_fixed0100_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0100X950g,
	"standard_vm_fixed0200_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0200X950g,
	"standard_vm_fixed0300_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0300X950g,
	"standard_vm_fixed0400_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0400X950g,
	"standard_vm_fixed0500_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0500X950g,
	"standard_vm_fixed0600_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0600X950g,
	"standard_vm_fixed0700_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0700X950g,
	"standard_vm_fixed0800_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0800X950g,
	"standard_vm_fixed0900_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed0900X950g,
	"standard_vm_fixed1000_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1000X950g,
	"standard_vm_fixed1100_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1100X950g,
	"standard_vm_fixed1200_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1200X950g,
	"standard_vm_fixed1300_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1300X950g,
	"standard_vm_fixed1400_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1400X950g,
	"standard_vm_fixed1500_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1500X950g,
	"standard_vm_fixed1600_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1600X950g,
	"standard_vm_fixed1700_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1700X950g,
	"standard_vm_fixed1800_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1800X950g,
	"standard_vm_fixed1900_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed1900X950g,
	"standard_vm_fixed2000_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2000X950g,
	"standard_vm_fixed2100_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2100X950g,
	"standard_vm_fixed2200_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2200X950g,
	"standard_vm_fixed2300_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2300X950g,
	"standard_vm_fixed2400_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2400X950g,
	"standard_vm_fixed2500_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2500X950g,
	"standard_vm_fixed2600_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2600X950g,
	"standard_vm_fixed2700_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2700X950g,
	"standard_vm_fixed2800_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2800X950g,
	"standard_vm_fixed2900_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed2900X950g,
	"standard_vm_fixed3000_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3000X950g,
	"standard_vm_fixed3100_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3100X950g,
	"standard_vm_fixed3200_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3200X950g,
	"standard_vm_fixed3300_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3300X950g,
	"standard_vm_fixed3400_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3400X950g,
	"standard_vm_fixed3500_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3500X950g,
	"standard_vm_fixed3600_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3600X950g,
	"standard_vm_fixed3700_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3700X950g,
	"standard_vm_fixed3800_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3800X950g,
	"standard_vm_fixed3900_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed3900X950g,
	"standard_vm_fixed4000_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4000X950g,
	"standard_vm_fixed4100_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4100X950g,
	"standard_vm_fixed4200_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4200X950g,
	"standard_vm_fixed4300_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4300X950g,
	"standard_vm_fixed4400_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4400X950g,
	"standard_vm_fixed4500_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4500X950g,
	"standard_vm_fixed4600_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4600X950g,
	"standard_vm_fixed4700_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4700X950g,
	"standard_vm_fixed4800_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4800X950g,
	"standard_vm_fixed4900_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed4900X950g,
	"standard_vm_fixed5000_x9_50g":         CreateInternalVnicAttachmentDetailsVnicShapeStandardVmFixed5000X950g,
	"subcore_standard_vm_fixed0025_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0025X950g,
	"subcore_standard_vm_fixed0050_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0050X950g,
	"subcore_standard_vm_fixed0075_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0075X950g,
	"subcore_standard_vm_fixed0100_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0100X950g,
	"subcore_standard_vm_fixed0125_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0125X950g,
	"subcore_standard_vm_fixed0150_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0150X950g,
	"subcore_standard_vm_fixed0175_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0175X950g,
	"subcore_standard_vm_fixed0200_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0200X950g,
	"subcore_standard_vm_fixed0225_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0225X950g,
	"subcore_standard_vm_fixed0250_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0250X950g,
	"subcore_standard_vm_fixed0275_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0275X950g,
	"subcore_standard_vm_fixed0300_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0300X950g,
	"subcore_standard_vm_fixed0325_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0325X950g,
	"subcore_standard_vm_fixed0350_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0350X950g,
	"subcore_standard_vm_fixed0375_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0375X950g,
	"subcore_standard_vm_fixed0400_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0400X950g,
	"subcore_standard_vm_fixed0425_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0425X950g,
	"subcore_standard_vm_fixed0450_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0450X950g,
	"subcore_standard_vm_fixed0475_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0475X950g,
	"subcore_standard_vm_fixed0500_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0500X950g,
	"subcore_standard_vm_fixed0525_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0525X950g,
	"subcore_standard_vm_fixed0550_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0550X950g,
	"subcore_standard_vm_fixed0575_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0575X950g,
	"subcore_standard_vm_fixed0600_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0600X950g,
	"subcore_standard_vm_fixed0625_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0625X950g,
	"subcore_standard_vm_fixed0650_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0650X950g,
	"subcore_standard_vm_fixed0675_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0675X950g,
	"subcore_standard_vm_fixed0700_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0700X950g,
	"subcore_standard_vm_fixed0725_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0725X950g,
	"subcore_standard_vm_fixed0750_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0750X950g,
	"subcore_standard_vm_fixed0775_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0775X950g,
	"subcore_standard_vm_fixed0800_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0800X950g,
	"subcore_standard_vm_fixed0825_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0825X950g,
	"subcore_standard_vm_fixed0850_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0850X950g,
	"subcore_standard_vm_fixed0875_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0875X950g,
	"subcore_standard_vm_fixed0900_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0900X950g,
	"subcore_standard_vm_fixed0925_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0925X950g,
	"subcore_standard_vm_fixed0950_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0950X950g,
	"subcore_standard_vm_fixed0975_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed0975X950g,
	"subcore_standard_vm_fixed1000_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1000X950g,
	"subcore_standard_vm_fixed1025_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1025X950g,
	"subcore_standard_vm_fixed1050_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1050X950g,
	"subcore_standard_vm_fixed1075_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1075X950g,
	"subcore_standard_vm_fixed1100_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1100X950g,
	"subcore_standard_vm_fixed1125_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1125X950g,
	"subcore_standard_vm_fixed1150_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1150X950g,
	"subcore_standard_vm_fixed1175_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1175X950g,
	"subcore_standard_vm_fixed1200_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1200X950g,
	"subcore_standard_vm_fixed1225_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1225X950g,
	"subcore_standard_vm_fixed1250_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1250X950g,
	"subcore_standard_vm_fixed1275_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1275X950g,
	"subcore_standard_vm_fixed1300_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1300X950g,
	"subcore_standard_vm_fixed1325_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1325X950g,
	"subcore_standard_vm_fixed1350_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1350X950g,
	"subcore_standard_vm_fixed1375_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1375X950g,
	"subcore_standard_vm_fixed1400_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1400X950g,
	"subcore_standard_vm_fixed1425_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1425X950g,
	"subcore_standard_vm_fixed1450_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1450X950g,
	"subcore_standard_vm_fixed1475_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1475X950g,
	"subcore_standard_vm_fixed1500_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1500X950g,
	"subcore_standard_vm_fixed1525_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1525X950g,
	"subcore_standard_vm_fixed1550_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1550X950g,
	"subcore_standard_vm_fixed1575_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1575X950g,
	"subcore_standard_vm_fixed1600_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1600X950g,
	"subcore_standard_vm_fixed1625_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1625X950g,
	"subcore_standard_vm_fixed1650_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1650X950g,
	"subcore_standard_vm_fixed1700_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1700X950g,
	"subcore_standard_vm_fixed1725_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1725X950g,
	"subcore_standard_vm_fixed1750_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1750X950g,
	"subcore_standard_vm_fixed1800_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1800X950g,
	"subcore_standard_vm_fixed1850_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1850X950g,
	"subcore_standard_vm_fixed1875_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1875X950g,
	"subcore_standard_vm_fixed1900_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1900X950g,
	"subcore_standard_vm_fixed1925_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1925X950g,
	"subcore_standard_vm_fixed1950_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed1950X950g,
	"subcore_standard_vm_fixed2000_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2000X950g,
	"subcore_standard_vm_fixed2025_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2025X950g,
	"subcore_standard_vm_fixed2050_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2050X950g,
	"subcore_standard_vm_fixed2100_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2100X950g,
	"subcore_standard_vm_fixed2125_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2125X950g,
	"subcore_standard_vm_fixed2150_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2150X950g,
	"subcore_standard_vm_fixed2175_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2175X950g,
	"subcore_standard_vm_fixed2200_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2200X950g,
	"subcore_standard_vm_fixed2250_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2250X950g,
	"subcore_standard_vm_fixed2275_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2275X950g,
	"subcore_standard_vm_fixed2300_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2300X950g,
	"subcore_standard_vm_fixed2325_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2325X950g,
	"subcore_standard_vm_fixed2350_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2350X950g,
	"subcore_standard_vm_fixed2375_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2375X950g,
	"subcore_standard_vm_fixed2400_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2400X950g,
	"subcore_standard_vm_fixed2450_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2450X950g,
	"subcore_standard_vm_fixed2475_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2475X950g,
	"subcore_standard_vm_fixed2500_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2500X950g,
	"subcore_standard_vm_fixed2550_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2550X950g,
	"subcore_standard_vm_fixed2600_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2600X950g,
	"subcore_standard_vm_fixed2625_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2625X950g,
	"subcore_standard_vm_fixed2650_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2650X950g,
	"subcore_standard_vm_fixed2700_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2700X950g,
	"subcore_standard_vm_fixed2750_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2750X950g,
	"subcore_standard_vm_fixed2775_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2775X950g,
	"subcore_standard_vm_fixed2800_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2800X950g,
	"subcore_standard_vm_fixed2850_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2850X950g,
	"subcore_standard_vm_fixed2875_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2875X950g,
	"subcore_standard_vm_fixed2900_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2900X950g,
	"subcore_standard_vm_fixed2925_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2925X950g,
	"subcore_standard_vm_fixed2950_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2950X950g,
	"subcore_standard_vm_fixed2975_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed2975X950g,
	"subcore_standard_vm_fixed3000_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3000X950g,
	"subcore_standard_vm_fixed3025_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3025X950g,
	"subcore_standard_vm_fixed3050_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3050X950g,
	"subcore_standard_vm_fixed3075_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3075X950g,
	"subcore_standard_vm_fixed3100_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3100X950g,
	"subcore_standard_vm_fixed3125_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3125X950g,
	"subcore_standard_vm_fixed3150_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3150X950g,
	"subcore_standard_vm_fixed3200_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3200X950g,
	"subcore_standard_vm_fixed3225_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3225X950g,
	"subcore_standard_vm_fixed3250_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3250X950g,
	"subcore_standard_vm_fixed3300_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3300X950g,
	"subcore_standard_vm_fixed3325_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3325X950g,
	"subcore_standard_vm_fixed3375_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3375X950g,
	"subcore_standard_vm_fixed3400_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3400X950g,
	"subcore_standard_vm_fixed3450_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3450X950g,
	"subcore_standard_vm_fixed3500_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3500X950g,
	"subcore_standard_vm_fixed3525_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3525X950g,
	"subcore_standard_vm_fixed3575_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3575X950g,
	"subcore_standard_vm_fixed3600_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3600X950g,
	"subcore_standard_vm_fixed3625_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3625X950g,
	"subcore_standard_vm_fixed3675_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3675X950g,
	"subcore_standard_vm_fixed3700_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3700X950g,
	"subcore_standard_vm_fixed3750_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3750X950g,
	"subcore_standard_vm_fixed3800_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3800X950g,
	"subcore_standard_vm_fixed3825_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3825X950g,
	"subcore_standard_vm_fixed3850_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3850X950g,
	"subcore_standard_vm_fixed3875_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3875X950g,
	"subcore_standard_vm_fixed3900_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3900X950g,
	"subcore_standard_vm_fixed3975_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed3975X950g,
	"subcore_standard_vm_fixed4000_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4000X950g,
	"subcore_standard_vm_fixed4025_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4025X950g,
	"subcore_standard_vm_fixed4050_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4050X950g,
	"subcore_standard_vm_fixed4100_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4100X950g,
	"subcore_standard_vm_fixed4125_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4125X950g,
	"subcore_standard_vm_fixed4200_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4200X950g,
	"subcore_standard_vm_fixed4225_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4225X950g,
	"subcore_standard_vm_fixed4250_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4250X950g,
	"subcore_standard_vm_fixed4275_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4275X950g,
	"subcore_standard_vm_fixed4300_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4300X950g,
	"subcore_standard_vm_fixed4350_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4350X950g,
	"subcore_standard_vm_fixed4375_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4375X950g,
	"subcore_standard_vm_fixed4400_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4400X950g,
	"subcore_standard_vm_fixed4425_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4425X950g,
	"subcore_standard_vm_fixed4500_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4500X950g,
	"subcore_standard_vm_fixed4550_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4550X950g,
	"subcore_standard_vm_fixed4575_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4575X950g,
	"subcore_standard_vm_fixed4600_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4600X950g,
	"subcore_standard_vm_fixed4625_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4625X950g,
	"subcore_standard_vm_fixed4650_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4650X950g,
	"subcore_standard_vm_fixed4675_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4675X950g,
	"subcore_standard_vm_fixed4700_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4700X950g,
	"subcore_standard_vm_fixed4725_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4725X950g,
	"subcore_standard_vm_fixed4750_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4750X950g,
	"subcore_standard_vm_fixed4800_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4800X950g,
	"subcore_standard_vm_fixed4875_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4875X950g,
	"subcore_standard_vm_fixed4900_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4900X950g,
	"subcore_standard_vm_fixed4950_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed4950X950g,
	"subcore_standard_vm_fixed5000_x9_50g": CreateInternalVnicAttachmentDetailsVnicShapeSubcoreStandardVmFixed5000X950g,
	"entirehost_x9_50g":                    CreateInternalVnicAttachmentDetailsVnicShapeEntirehostX950g,
	"dynamic_x9_100g":                      CreateInternalVnicAttachmentDetailsVnicShapeDynamicX9100g,
	"dynamic_x10_50g":                      CreateInternalVnicAttachmentDetailsVnicShapeDynamicX1050g,
	"dynamic_x10_100g":                     CreateInternalVnicAttachmentDetailsVnicShapeDynamicX10100g,
}

// GetCreateInternalVnicAttachmentDetailsVnicShapeEnumValues Enumerates the set of values for CreateInternalVnicAttachmentDetailsVnicShapeEnum
func GetCreateInternalVnicAttachmentDetailsVnicShapeEnumValues() []CreateInternalVnicAttachmentDetailsVnicShapeEnum {
	values := make([]CreateInternalVnicAttachmentDetailsVnicShapeEnum, 0)
	for _, v := range mappingCreateInternalVnicAttachmentDetailsVnicShapeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateInternalVnicAttachmentDetailsVnicShapeEnumStringValues Enumerates the set of values in String for CreateInternalVnicAttachmentDetailsVnicShapeEnum
func GetCreateInternalVnicAttachmentDetailsVnicShapeEnumStringValues() []string {
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
		"DYNAMIC_E5_50G",
		"DYNAMIC_E5_100G",
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
		"FIXED5000_TELESIS_A1_50G",
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
		"STANDARD_VM_FIXED4100_X9_50G",
		"STANDARD_VM_FIXED4200_X9_50G",
		"STANDARD_VM_FIXED4300_X9_50G",
		"STANDARD_VM_FIXED4400_X9_50G",
		"STANDARD_VM_FIXED4500_X9_50G",
		"STANDARD_VM_FIXED4600_X9_50G",
		"STANDARD_VM_FIXED4700_X9_50G",
		"STANDARD_VM_FIXED4800_X9_50G",
		"STANDARD_VM_FIXED4900_X9_50G",
		"STANDARD_VM_FIXED5000_X9_50G",
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
		"DYNAMIC_X9_100G",
		"DYNAMIC_X10_50G",
		"DYNAMIC_X10_100G",
	}
}

// GetMappingCreateInternalVnicAttachmentDetailsVnicShapeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateInternalVnicAttachmentDetailsVnicShapeEnum(val string) (CreateInternalVnicAttachmentDetailsVnicShapeEnum, bool) {
	enum, ok := mappingCreateInternalVnicAttachmentDetailsVnicShapeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateInternalVnicAttachmentDetailsLaunchTypeEnum Enum with underlying type: string
type CreateInternalVnicAttachmentDetailsLaunchTypeEnum string

// Set of constants representing the allowable values for CreateInternalVnicAttachmentDetailsLaunchTypeEnum
const (
	CreateInternalVnicAttachmentDetailsLaunchTypeMarketplace CreateInternalVnicAttachmentDetailsLaunchTypeEnum = "MARKETPLACE"
	CreateInternalVnicAttachmentDetailsLaunchTypeStandard    CreateInternalVnicAttachmentDetailsLaunchTypeEnum = "STANDARD"
)

var mappingCreateInternalVnicAttachmentDetailsLaunchTypeEnum = map[string]CreateInternalVnicAttachmentDetailsLaunchTypeEnum{
	"MARKETPLACE": CreateInternalVnicAttachmentDetailsLaunchTypeMarketplace,
	"STANDARD":    CreateInternalVnicAttachmentDetailsLaunchTypeStandard,
}

var mappingCreateInternalVnicAttachmentDetailsLaunchTypeEnumLowerCase = map[string]CreateInternalVnicAttachmentDetailsLaunchTypeEnum{
	"marketplace": CreateInternalVnicAttachmentDetailsLaunchTypeMarketplace,
	"standard":    CreateInternalVnicAttachmentDetailsLaunchTypeStandard,
}

// GetCreateInternalVnicAttachmentDetailsLaunchTypeEnumValues Enumerates the set of values for CreateInternalVnicAttachmentDetailsLaunchTypeEnum
func GetCreateInternalVnicAttachmentDetailsLaunchTypeEnumValues() []CreateInternalVnicAttachmentDetailsLaunchTypeEnum {
	values := make([]CreateInternalVnicAttachmentDetailsLaunchTypeEnum, 0)
	for _, v := range mappingCreateInternalVnicAttachmentDetailsLaunchTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateInternalVnicAttachmentDetailsLaunchTypeEnumStringValues Enumerates the set of values in String for CreateInternalVnicAttachmentDetailsLaunchTypeEnum
func GetCreateInternalVnicAttachmentDetailsLaunchTypeEnumStringValues() []string {
	return []string{
		"MARKETPLACE",
		"STANDARD",
	}
}

// GetMappingCreateInternalVnicAttachmentDetailsLaunchTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateInternalVnicAttachmentDetailsLaunchTypeEnum(val string) (CreateInternalVnicAttachmentDetailsLaunchTypeEnum, bool) {
	enum, ok := mappingCreateInternalVnicAttachmentDetailsLaunchTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
