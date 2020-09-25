// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// CreateInternalVnicAttachmentDetails Details for attaching a service VNIC to VNICaaS fleet.
type CreateInternalVnicAttachmentDetails struct {

	// The OCID of the compartment to contain the VNIC attachment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique.
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
}

func (m CreateInternalVnicAttachmentDetails) String() string {
	return common.PointerString(m)
}

// CreateInternalVnicAttachmentDetailsVnicShapeEnum Enum with underlying type: string
type CreateInternalVnicAttachmentDetailsVnicShapeEnum string

// Set of constants representing the allowable values for CreateInternalVnicAttachmentDetailsVnicShapeEnum
const (
	CreateInternalVnicAttachmentDetailsVnicShapeDynamic               CreateInternalVnicAttachmentDetailsVnicShapeEnum = "DYNAMIC"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0040             CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0040"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0060             CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0060"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0060Psm          CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0060_PSM"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0100             CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0100"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0120             CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0120"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed01202x           CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0120_2X"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0200             CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0200"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0240             CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0240"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0480             CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0480"
	CreateInternalVnicAttachmentDetailsVnicShapeEntirehost            CreateInternalVnicAttachmentDetailsVnicShapeEnum = "ENTIREHOST"
	CreateInternalVnicAttachmentDetailsVnicShapeDynamic25g            CreateInternalVnicAttachmentDetailsVnicShapeEnum = "DYNAMIC_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed004025g          CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0040_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed010025g          CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0100_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed020025g          CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0200_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed040025g          CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0400_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed080025g          CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0800_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed160025g          CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1600_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed240025g          CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2400_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeEntirehost25g         CreateInternalVnicAttachmentDetailsVnicShapeEnum = "ENTIREHOST_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeDynamicE125g          CreateInternalVnicAttachmentDetailsVnicShapeEnum = "DYNAMIC_E1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0040E125g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0040_E1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0070E125g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0070_E1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0140E125g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0140_E1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0280E125g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0280_E1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0560E125g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0560_E1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1120E125g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1120_E1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1680E125g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1680_E1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeEntirehostE125g       CreateInternalVnicAttachmentDetailsVnicShapeEnum = "ENTIREHOST_E1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeDynamicB125g          CreateInternalVnicAttachmentDetailsVnicShapeEnum = "DYNAMIC_B1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0040B125g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0040_B1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0060B125g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0060_B1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0120B125g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0120_B1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0240B125g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0240_B1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0480B125g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0480_B1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0960B125g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0960_B1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeEntirehostB125g       CreateInternalVnicAttachmentDetailsVnicShapeEnum = "ENTIREHOST_B1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeMicroVmFixed0048E125g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "MICRO_VM_FIXED0048_E1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeMicroLbFixed0001E125g CreateInternalVnicAttachmentDetailsVnicShapeEnum = "MICRO_LB_FIXED0001_E1_25G"
	CreateInternalVnicAttachmentDetailsVnicShapeVnicaasFixed0200      CreateInternalVnicAttachmentDetailsVnicShapeEnum = "VNICAAS_FIXED0200"
	CreateInternalVnicAttachmentDetailsVnicShapeVnicaasFixed0400      CreateInternalVnicAttachmentDetailsVnicShapeEnum = "VNICAAS_FIXED0400"
	CreateInternalVnicAttachmentDetailsVnicShapeDynamicE350g          CreateInternalVnicAttachmentDetailsVnicShapeEnum = "DYNAMIC_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0040E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0040_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0100E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0100_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0200E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0200_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0300E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0300_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0400E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0400_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0500E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0500_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0600E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0600_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0700E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0700_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0800E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0800_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed0900E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED0900_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1000E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1000_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1100E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1100_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1200E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1200_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1300E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1300_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1400E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1400_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1500E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1500_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1600E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1600_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1700E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1700_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1800E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1800_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed1900E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED1900_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2000E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2000_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2100E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2100_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2200E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2200_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2300E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2300_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2400E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2400_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2500E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2500_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2600E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2600_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2700E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2700_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2800E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2800_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed2900E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED2900_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3000E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3000_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3100E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3100_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3200E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3200_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3300E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3300_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3400E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3400_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3500E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3500_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3600E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3600_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3700E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3700_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3800E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3800_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed3900E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED3900_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeFixed4000E350g        CreateInternalVnicAttachmentDetailsVnicShapeEnum = "FIXED4000_E3_50G"
	CreateInternalVnicAttachmentDetailsVnicShapeEntirehostE350g       CreateInternalVnicAttachmentDetailsVnicShapeEnum = "ENTIREHOST_E3_50G"
)

var mappingCreateInternalVnicAttachmentDetailsVnicShape = map[string]CreateInternalVnicAttachmentDetailsVnicShapeEnum{
	"DYNAMIC":                   CreateInternalVnicAttachmentDetailsVnicShapeDynamic,
	"FIXED0040":                 CreateInternalVnicAttachmentDetailsVnicShapeFixed0040,
	"FIXED0060":                 CreateInternalVnicAttachmentDetailsVnicShapeFixed0060,
	"FIXED0060_PSM":             CreateInternalVnicAttachmentDetailsVnicShapeFixed0060Psm,
	"FIXED0100":                 CreateInternalVnicAttachmentDetailsVnicShapeFixed0100,
	"FIXED0120":                 CreateInternalVnicAttachmentDetailsVnicShapeFixed0120,
	"FIXED0120_2X":              CreateInternalVnicAttachmentDetailsVnicShapeFixed01202x,
	"FIXED0200":                 CreateInternalVnicAttachmentDetailsVnicShapeFixed0200,
	"FIXED0240":                 CreateInternalVnicAttachmentDetailsVnicShapeFixed0240,
	"FIXED0480":                 CreateInternalVnicAttachmentDetailsVnicShapeFixed0480,
	"ENTIREHOST":                CreateInternalVnicAttachmentDetailsVnicShapeEntirehost,
	"DYNAMIC_25G":               CreateInternalVnicAttachmentDetailsVnicShapeDynamic25g,
	"FIXED0040_25G":             CreateInternalVnicAttachmentDetailsVnicShapeFixed004025g,
	"FIXED0100_25G":             CreateInternalVnicAttachmentDetailsVnicShapeFixed010025g,
	"FIXED0200_25G":             CreateInternalVnicAttachmentDetailsVnicShapeFixed020025g,
	"FIXED0400_25G":             CreateInternalVnicAttachmentDetailsVnicShapeFixed040025g,
	"FIXED0800_25G":             CreateInternalVnicAttachmentDetailsVnicShapeFixed080025g,
	"FIXED1600_25G":             CreateInternalVnicAttachmentDetailsVnicShapeFixed160025g,
	"FIXED2400_25G":             CreateInternalVnicAttachmentDetailsVnicShapeFixed240025g,
	"ENTIREHOST_25G":            CreateInternalVnicAttachmentDetailsVnicShapeEntirehost25g,
	"DYNAMIC_E1_25G":            CreateInternalVnicAttachmentDetailsVnicShapeDynamicE125g,
	"FIXED0040_E1_25G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed0040E125g,
	"FIXED0070_E1_25G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed0070E125g,
	"FIXED0140_E1_25G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed0140E125g,
	"FIXED0280_E1_25G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed0280E125g,
	"FIXED0560_E1_25G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed0560E125g,
	"FIXED1120_E1_25G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed1120E125g,
	"FIXED1680_E1_25G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed1680E125g,
	"ENTIREHOST_E1_25G":         CreateInternalVnicAttachmentDetailsVnicShapeEntirehostE125g,
	"DYNAMIC_B1_25G":            CreateInternalVnicAttachmentDetailsVnicShapeDynamicB125g,
	"FIXED0040_B1_25G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed0040B125g,
	"FIXED0060_B1_25G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed0060B125g,
	"FIXED0120_B1_25G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed0120B125g,
	"FIXED0240_B1_25G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed0240B125g,
	"FIXED0480_B1_25G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed0480B125g,
	"FIXED0960_B1_25G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed0960B125g,
	"ENTIREHOST_B1_25G":         CreateInternalVnicAttachmentDetailsVnicShapeEntirehostB125g,
	"MICRO_VM_FIXED0048_E1_25G": CreateInternalVnicAttachmentDetailsVnicShapeMicroVmFixed0048E125g,
	"MICRO_LB_FIXED0001_E1_25G": CreateInternalVnicAttachmentDetailsVnicShapeMicroLbFixed0001E125g,
	"VNICAAS_FIXED0200":         CreateInternalVnicAttachmentDetailsVnicShapeVnicaasFixed0200,
	"VNICAAS_FIXED0400":         CreateInternalVnicAttachmentDetailsVnicShapeVnicaasFixed0400,
	"DYNAMIC_E3_50G":            CreateInternalVnicAttachmentDetailsVnicShapeDynamicE350g,
	"FIXED0040_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed0040E350g,
	"FIXED0100_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed0100E350g,
	"FIXED0200_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed0200E350g,
	"FIXED0300_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed0300E350g,
	"FIXED0400_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed0400E350g,
	"FIXED0500_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed0500E350g,
	"FIXED0600_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed0600E350g,
	"FIXED0700_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed0700E350g,
	"FIXED0800_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed0800E350g,
	"FIXED0900_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed0900E350g,
	"FIXED1000_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed1000E350g,
	"FIXED1100_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed1100E350g,
	"FIXED1200_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed1200E350g,
	"FIXED1300_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed1300E350g,
	"FIXED1400_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed1400E350g,
	"FIXED1500_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed1500E350g,
	"FIXED1600_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed1600E350g,
	"FIXED1700_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed1700E350g,
	"FIXED1800_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed1800E350g,
	"FIXED1900_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed1900E350g,
	"FIXED2000_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed2000E350g,
	"FIXED2100_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed2100E350g,
	"FIXED2200_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed2200E350g,
	"FIXED2300_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed2300E350g,
	"FIXED2400_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed2400E350g,
	"FIXED2500_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed2500E350g,
	"FIXED2600_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed2600E350g,
	"FIXED2700_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed2700E350g,
	"FIXED2800_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed2800E350g,
	"FIXED2900_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed2900E350g,
	"FIXED3000_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed3000E350g,
	"FIXED3100_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed3100E350g,
	"FIXED3200_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed3200E350g,
	"FIXED3300_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed3300E350g,
	"FIXED3400_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed3400E350g,
	"FIXED3500_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed3500E350g,
	"FIXED3600_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed3600E350g,
	"FIXED3700_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed3700E350g,
	"FIXED3800_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed3800E350g,
	"FIXED3900_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed3900E350g,
	"FIXED4000_E3_50G":          CreateInternalVnicAttachmentDetailsVnicShapeFixed4000E350g,
	"ENTIREHOST_E3_50G":         CreateInternalVnicAttachmentDetailsVnicShapeEntirehostE350g,
}

// GetCreateInternalVnicAttachmentDetailsVnicShapeEnumValues Enumerates the set of values for CreateInternalVnicAttachmentDetailsVnicShapeEnum
func GetCreateInternalVnicAttachmentDetailsVnicShapeEnumValues() []CreateInternalVnicAttachmentDetailsVnicShapeEnum {
	values := make([]CreateInternalVnicAttachmentDetailsVnicShapeEnum, 0)
	for _, v := range mappingCreateInternalVnicAttachmentDetailsVnicShape {
		values = append(values, v)
	}
	return values
}
