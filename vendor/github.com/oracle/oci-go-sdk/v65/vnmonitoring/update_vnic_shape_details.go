// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateVnicShapeDetails This structure is used when updating the shape of VNIC in VNIC attachment.
type UpdateVnicShapeDetails struct {

	// VNIC whose attachments need to be updated to the destination vnic shape.
	VnicId *string `mandatory:"true" json:"vnicId"`

	// Shape of VNIC that will be used to update VNIC attachment.
	VnicShape UpdateVnicShapeDetailsVnicShapeEnum `mandatory:"true" json:"vnicShape"`
}

func (m UpdateVnicShapeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateVnicShapeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpdateVnicShapeDetailsVnicShapeEnum(string(m.VnicShape)); !ok && m.VnicShape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VnicShape: %s. Supported values are: %s.", m.VnicShape, strings.Join(GetUpdateVnicShapeDetailsVnicShapeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateVnicShapeDetailsVnicShapeEnum Enum with underlying type: string
type UpdateVnicShapeDetailsVnicShapeEnum string

// Set of constants representing the allowable values for UpdateVnicShapeDetailsVnicShapeEnum
const (
	UpdateVnicShapeDetailsVnicShapeDynamic                         UpdateVnicShapeDetailsVnicShapeEnum = "DYNAMIC"
	UpdateVnicShapeDetailsVnicShapeFixed0040                       UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0040"
	UpdateVnicShapeDetailsVnicShapeFixed0060                       UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0060"
	UpdateVnicShapeDetailsVnicShapeFixed0060Psm                    UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0060_PSM"
	UpdateVnicShapeDetailsVnicShapeFixed0100                       UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0100"
	UpdateVnicShapeDetailsVnicShapeFixed0120                       UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0120"
	UpdateVnicShapeDetailsVnicShapeFixed01202x                     UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0120_2X"
	UpdateVnicShapeDetailsVnicShapeFixed0200                       UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0200"
	UpdateVnicShapeDetailsVnicShapeFixed0240                       UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0240"
	UpdateVnicShapeDetailsVnicShapeFixed0480                       UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0480"
	UpdateVnicShapeDetailsVnicShapeEntirehost                      UpdateVnicShapeDetailsVnicShapeEnum = "ENTIREHOST"
	UpdateVnicShapeDetailsVnicShapeDynamic25g                      UpdateVnicShapeDetailsVnicShapeEnum = "DYNAMIC_25G"
	UpdateVnicShapeDetailsVnicShapeFixed004025g                    UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0040_25G"
	UpdateVnicShapeDetailsVnicShapeFixed010025g                    UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0100_25G"
	UpdateVnicShapeDetailsVnicShapeFixed020025g                    UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0200_25G"
	UpdateVnicShapeDetailsVnicShapeFixed040025g                    UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0400_25G"
	UpdateVnicShapeDetailsVnicShapeFixed080025g                    UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0800_25G"
	UpdateVnicShapeDetailsVnicShapeFixed160025g                    UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1600_25G"
	UpdateVnicShapeDetailsVnicShapeFixed240025g                    UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2400_25G"
	UpdateVnicShapeDetailsVnicShapeEntirehost25g                   UpdateVnicShapeDetailsVnicShapeEnum = "ENTIREHOST_25G"
	UpdateVnicShapeDetailsVnicShapeDynamicE125g                    UpdateVnicShapeDetailsVnicShapeEnum = "DYNAMIC_E1_25G"
	UpdateVnicShapeDetailsVnicShapeFixed0040E125g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0040_E1_25G"
	UpdateVnicShapeDetailsVnicShapeFixed0070E125g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0070_E1_25G"
	UpdateVnicShapeDetailsVnicShapeFixed0140E125g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0140_E1_25G"
	UpdateVnicShapeDetailsVnicShapeFixed0280E125g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0280_E1_25G"
	UpdateVnicShapeDetailsVnicShapeFixed0560E125g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0560_E1_25G"
	UpdateVnicShapeDetailsVnicShapeFixed1120E125g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1120_E1_25G"
	UpdateVnicShapeDetailsVnicShapeFixed1680E125g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1680_E1_25G"
	UpdateVnicShapeDetailsVnicShapeEntirehostE125g                 UpdateVnicShapeDetailsVnicShapeEnum = "ENTIREHOST_E1_25G"
	UpdateVnicShapeDetailsVnicShapeDynamicB125g                    UpdateVnicShapeDetailsVnicShapeEnum = "DYNAMIC_B1_25G"
	UpdateVnicShapeDetailsVnicShapeFixed0040B125g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0040_B1_25G"
	UpdateVnicShapeDetailsVnicShapeFixed0060B125g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0060_B1_25G"
	UpdateVnicShapeDetailsVnicShapeFixed0120B125g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0120_B1_25G"
	UpdateVnicShapeDetailsVnicShapeFixed0240B125g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0240_B1_25G"
	UpdateVnicShapeDetailsVnicShapeFixed0480B125g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0480_B1_25G"
	UpdateVnicShapeDetailsVnicShapeFixed0960B125g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0960_B1_25G"
	UpdateVnicShapeDetailsVnicShapeEntirehostB125g                 UpdateVnicShapeDetailsVnicShapeEnum = "ENTIREHOST_B1_25G"
	UpdateVnicShapeDetailsVnicShapeMicroVmFixed0048E125g           UpdateVnicShapeDetailsVnicShapeEnum = "MICRO_VM_FIXED0048_E1_25G"
	UpdateVnicShapeDetailsVnicShapeMicroLbFixed0001E125g           UpdateVnicShapeDetailsVnicShapeEnum = "MICRO_LB_FIXED0001_E1_25G"
	UpdateVnicShapeDetailsVnicShapeVnicaasFixed0025                UpdateVnicShapeDetailsVnicShapeEnum = "VNICAAS_FIXED0025"
	UpdateVnicShapeDetailsVnicShapeVnicaasFixed0200                UpdateVnicShapeDetailsVnicShapeEnum = "VNICAAS_FIXED0200"
	UpdateVnicShapeDetailsVnicShapeVnicaasFixed0400                UpdateVnicShapeDetailsVnicShapeEnum = "VNICAAS_FIXED0400"
	UpdateVnicShapeDetailsVnicShapeVnicaasFixed0625                UpdateVnicShapeDetailsVnicShapeEnum = "VNICAAS_FIXED0625"
	UpdateVnicShapeDetailsVnicShapeVnicaasFixed0700                UpdateVnicShapeDetailsVnicShapeEnum = "VNICAAS_FIXED0700"
	UpdateVnicShapeDetailsVnicShapeVnicaasFixed1250                UpdateVnicShapeDetailsVnicShapeEnum = "VNICAAS_FIXED1250"
	UpdateVnicShapeDetailsVnicShapeVnicaasFwaas6g250m              UpdateVnicShapeDetailsVnicShapeEnum = "VNICAAS_FWAAS_6G_250M"
	UpdateVnicShapeDetailsVnicShapeVnicaasNlbApproved10g           UpdateVnicShapeDetailsVnicShapeEnum = "VNICAAS_NLB_APPROVED_10G"
	UpdateVnicShapeDetailsVnicShapeVnicaasNlbApproved25g           UpdateVnicShapeDetailsVnicShapeEnum = "VNICAAS_NLB_APPROVED_25G"
	UpdateVnicShapeDetailsVnicShapeVnicaasTelesis25g               UpdateVnicShapeDetailsVnicShapeEnum = "VNICAAS_TELESIS_25G"
	UpdateVnicShapeDetailsVnicShapeVnicaasTelesis10g               UpdateVnicShapeDetailsVnicShapeEnum = "VNICAAS_TELESIS_10G"
	UpdateVnicShapeDetailsVnicShapeVnicaasAmbassadorFixed0100      UpdateVnicShapeDetailsVnicShapeEnum = "VNICAAS_AMBASSADOR_FIXED0100"
	UpdateVnicShapeDetailsVnicShapeVnicaasTelesisGamma             UpdateVnicShapeDetailsVnicShapeEnum = "VNICAAS_TELESIS_GAMMA"
	UpdateVnicShapeDetailsVnicShapeVnicaasPrivatedns               UpdateVnicShapeDetailsVnicShapeEnum = "VNICAAS_PRIVATEDNS"
	UpdateVnicShapeDetailsVnicShapeVnicaasFwaas                    UpdateVnicShapeDetailsVnicShapeEnum = "VNICAAS_FWAAS"
	UpdateVnicShapeDetailsVnicShapeVnicaasLbaasFree                UpdateVnicShapeDetailsVnicShapeEnum = "VNICAAS_LBAAS_FREE"
	UpdateVnicShapeDetailsVnicShapeVnicaasLbaas8g512k              UpdateVnicShapeDetailsVnicShapeEnum = "VNICAAS_LBAAS_8G_512K"
	UpdateVnicShapeDetailsVnicShapeVnicaasLbaas8g1m                UpdateVnicShapeDetailsVnicShapeEnum = "VNICAAS_LBAAS_8G_1M"
	UpdateVnicShapeDetailsVnicShapeVnicaasLbaas8g2m                UpdateVnicShapeDetailsVnicShapeEnum = "VNICAAS_LBAAS_8G_2M"
	UpdateVnicShapeDetailsVnicShapeVnicaasLbaas8g3m                UpdateVnicShapeDetailsVnicShapeEnum = "VNICAAS_LBAAS_8G_3M"
	UpdateVnicShapeDetailsVnicShapeVnicaasLbaas8g1m8ghost          UpdateVnicShapeDetailsVnicShapeEnum = "VNICAAS_LBAAS_8G_1M_8GHOST"
	UpdateVnicShapeDetailsVnicShapeVnicaasLbaas8g1m16ghost         UpdateVnicShapeDetailsVnicShapeEnum = "VNICAAS_LBAAS_8G_1M_16GHOST"
	UpdateVnicShapeDetailsVnicShapeDynamicE350g                    UpdateVnicShapeDetailsVnicShapeEnum = "DYNAMIC_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0040E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0040_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0100E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0100_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0200E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0200_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0300E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0300_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0400E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0400_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0500E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0500_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0600E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0600_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0700E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0700_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0800E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0800_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0900E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0900_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1000E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1000_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1100E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1100_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1200E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1200_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1300E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1300_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1400E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1400_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1500E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1500_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1600E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1600_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1700E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1700_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1800E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1800_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1900E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1900_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2000E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2000_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2100E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2100_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2200E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2200_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2300E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2300_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2400E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2400_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2500E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2500_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2600E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2600_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2700E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2700_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2800E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2800_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2900E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2900_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3000E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3000_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3100E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3100_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3200E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3200_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3300E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3300_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3400E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3400_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3500E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3500_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3600E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3600_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3700E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3700_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3800E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3800_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3900E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3900_E3_50G"
	UpdateVnicShapeDetailsVnicShapeFixed4000E350g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED4000_E3_50G"
	UpdateVnicShapeDetailsVnicShapeEntirehostE350g                 UpdateVnicShapeDetailsVnicShapeEnum = "ENTIREHOST_E3_50G"
	UpdateVnicShapeDetailsVnicShapeDynamicE450g                    UpdateVnicShapeDetailsVnicShapeEnum = "DYNAMIC_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0040E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0040_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0100E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0100_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0200E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0200_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0300E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0300_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0400E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0400_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0500E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0500_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0600E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0600_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0700E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0700_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0800E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0800_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0900E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0900_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1000E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1000_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1100E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1100_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1200E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1200_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1300E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1300_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1400E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1400_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1500E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1500_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1600E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1600_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1700E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1700_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1800E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1800_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1900E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1900_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2000E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2000_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2100E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2100_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2200E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2200_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2300E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2300_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2400E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2400_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2500E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2500_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2600E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2600_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2700E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2700_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2800E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2800_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2900E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2900_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3000E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3000_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3100E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3100_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3200E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3200_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3300E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3300_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3400E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3400_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3500E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3500_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3600E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3600_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3700E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3700_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3800E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3800_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3900E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3900_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed4000E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED4000_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed4700E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED4700_E4_50G"
	UpdateVnicShapeDetailsVnicShapeFixed4960E450g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED4960_E4_50G"
	UpdateVnicShapeDetailsVnicShapeEntirehostE450g                 UpdateVnicShapeDetailsVnicShapeEnum = "ENTIREHOST_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed01001E450g          UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED0100_1_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed02002E450g          UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED0200_2_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed03003E450g          UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED0300_3_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed04004E450g          UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED0400_4_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed05005E450g          UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED0500_5_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed06006E450g          UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED0600_6_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed07007E450g          UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED0700_7_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed08008E450g          UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED0800_8_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed09009E450g          UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED0900_9_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed100010E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED1000_10_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed110011E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED1100_11_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed120012E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED1200_12_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed130013E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED1300_13_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed140014E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED1400_14_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed150015E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED1500_15_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed160016E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED1600_16_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed170017E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED1700_17_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed180018E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED1800_18_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed190019E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED1900_19_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed200020E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED2000_20_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed210021E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED2100_21_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed220022E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED2200_22_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed230023E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED2300_23_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed240024E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED2400_24_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed250025E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED2500_25_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed260026E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED2600_26_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed270027E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED2700_27_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed280028E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED2800_28_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed290029E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED2900_29_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed300030E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3000_30_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed310031E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3100_31_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320032E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_32_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed330033E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3300_33_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed340034E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3400_34_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed350035E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3500_35_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed360036E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3600_36_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed370037E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3700_37_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed380038E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3800_38_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed390039E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3900_39_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400040E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_40_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400041E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_41_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400042E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_42_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400043E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_43_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400044E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_44_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400045E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_45_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400046E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_46_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400047E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_47_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400048E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_48_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400049E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_49_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400050E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_50_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400051E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_51_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400052E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_52_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400053E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_53_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400054E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_54_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400055E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_55_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400056E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_56_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400057E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_57_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400058E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_58_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400059E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_59_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400060E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_60_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400061E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_61_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400062E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_62_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400063E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_63_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400064E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_64_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400065E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_65_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400066E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_66_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400067E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_67_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400068E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_68_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400069E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_69_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400070E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_70_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400071E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_71_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400072E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_72_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400073E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_73_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400074E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_74_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400075E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_75_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400076E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_76_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400077E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_77_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400078E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_78_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400079E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_79_E4_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400080E450g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED4000_80_E4_50G"
	UpdateVnicShapeDetailsVnicShapeMicroVmFixed0050E350g           UpdateVnicShapeDetailsVnicShapeEnum = "Micro_VM_Fixed0050_E3_50G"
	UpdateVnicShapeDetailsVnicShapeMicroVmFixed0050E450g           UpdateVnicShapeDetailsVnicShapeEnum = "Micro_VM_Fixed0050_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0025E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0025_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0050E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0050_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0075E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0075_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0100E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0100_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0125E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0125_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0150E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0150_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0175E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0175_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0200E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0200_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0225E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0225_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0250E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0250_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0275E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0275_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0300E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0300_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0325E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0325_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0350E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0350_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0375E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0375_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0400E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0400_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0425E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0425_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0450E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0450_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0475E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0475_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0500E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0500_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0525E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0525_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0550E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0550_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0575E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0575_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0600E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0600_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0625E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0625_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0650E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0650_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0675E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0675_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0700E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0700_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0725E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0725_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0750E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0750_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0775E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0775_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0800E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0800_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0825E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0825_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0850E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0850_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0875E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0875_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0900E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0900_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0925E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0925_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0950E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0950_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0975E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0975_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1000E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1000_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1025E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1025_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1050E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1050_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1075E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1075_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1100E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1100_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1125E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1125_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1150E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1150_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1175E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1175_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1200E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1200_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1225E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1225_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1250E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1250_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1275E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1275_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1300E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1300_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1325E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1325_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1350E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1350_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1375E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1375_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1400E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1400_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1425E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1425_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1450E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1450_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1475E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1475_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1500E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1500_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1525E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1525_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1550E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1550_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1575E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1575_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1600E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1600_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1625E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1625_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1650E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1650_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1700E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1700_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1725E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1725_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1750E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1750_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1800E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1800_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1850E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1850_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1875E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1875_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1900E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1900_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1925E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1925_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1950E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1950_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2000E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2000_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2025E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2025_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2050E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2050_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2100E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2100_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2125E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2125_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2150E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2150_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2175E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2175_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2200E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2200_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2250E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2250_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2275E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2275_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2300E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2300_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2325E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2325_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2350E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2350_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2375E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2375_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2400E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2400_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2450E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2450_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2475E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2475_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2500E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2500_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2550E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2550_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2600E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2600_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2625E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2625_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2650E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2650_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2700E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2700_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2750E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2750_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2775E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2775_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2800E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2800_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2850E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2850_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2875E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2875_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2900E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2900_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2925E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2925_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2950E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2950_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2975E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2975_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3000E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3000_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3025E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3025_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3050E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3050_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3075E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3075_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3100E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3100_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3125E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3125_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3150E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3150_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3200E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3200_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3225E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3225_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3250E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3250_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3300E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3300_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3325E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3325_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3375E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3375_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3400E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3400_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3450E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3450_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3500E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3500_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3525E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3525_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3575E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3575_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3600E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3600_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3625E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3625_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3675E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3675_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3700E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3700_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3750E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3750_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3800E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3800_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3825E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3825_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3850E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3850_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3875E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3875_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3900E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3900_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3975E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3975_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4000E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4000_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4025E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4025_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4050E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4050_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4100E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4100_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4125E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4125_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4200E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4200_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4225E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4225_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4250E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4250_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4275E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4275_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4300E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4300_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4350E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4350_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4375E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4375_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4400E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4400_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4425E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4425_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4500E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4500_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4550E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4550_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4575E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4575_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4600E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4600_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4625E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4625_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4650E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4650_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4675E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4675_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4700E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4700_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4725E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4725_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4750E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4750_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4800E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4800_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4875E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4875_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4900E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4900_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4950E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4950_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed5000E350g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED5000_E3_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0025E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0025_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0050E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0050_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0075E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0075_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0100E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0100_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0125E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0125_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0150E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0150_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0175E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0175_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0200E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0200_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0225E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0225_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0250E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0250_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0275E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0275_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0300E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0300_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0325E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0325_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0350E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0350_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0375E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0375_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0400E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0400_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0425E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0425_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0450E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0450_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0475E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0475_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0500E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0500_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0525E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0525_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0550E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0550_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0575E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0575_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0600E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0600_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0625E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0625_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0650E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0650_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0675E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0675_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0700E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0700_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0725E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0725_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0750E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0750_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0775E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0775_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0800E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0800_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0825E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0825_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0850E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0850_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0875E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0875_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0900E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0900_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0925E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0925_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0950E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0950_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0975E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0975_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1000E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1000_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1025E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1025_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1050E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1050_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1075E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1075_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1100E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1100_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1125E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1125_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1150E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1150_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1175E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1175_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1200E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1200_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1225E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1225_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1250E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1250_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1275E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1275_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1300E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1300_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1325E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1325_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1350E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1350_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1375E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1375_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1400E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1400_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1425E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1425_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1450E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1450_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1475E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1475_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1500E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1500_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1525E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1525_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1550E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1550_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1575E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1575_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1600E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1600_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1625E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1625_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1650E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1650_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1700E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1700_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1725E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1725_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1750E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1750_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1800E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1800_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1850E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1850_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1875E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1875_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1900E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1900_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1925E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1925_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1950E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1950_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2000E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2000_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2025E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2025_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2050E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2050_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2100E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2100_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2125E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2125_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2150E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2150_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2175E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2175_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2200E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2200_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2250E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2250_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2275E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2275_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2300E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2300_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2325E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2325_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2350E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2350_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2375E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2375_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2400E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2400_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2450E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2450_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2475E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2475_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2500E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2500_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2550E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2550_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2600E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2600_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2625E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2625_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2650E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2650_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2700E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2700_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2750E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2750_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2775E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2775_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2800E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2800_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2850E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2850_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2875E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2875_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2900E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2900_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2925E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2925_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2950E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2950_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2975E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2975_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3000E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3000_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3025E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3025_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3050E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3050_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3075E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3075_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3100E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3100_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3125E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3125_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3150E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3150_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3200E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3200_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3225E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3225_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3250E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3250_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3300E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3300_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3325E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3325_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3375E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3375_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3400E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3400_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3450E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3450_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3500E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3500_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3525E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3525_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3575E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3575_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3600E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3600_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3625E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3625_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3675E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3675_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3700E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3700_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3750E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3750_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3800E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3800_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3825E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3825_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3850E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3850_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3875E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3875_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3900E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3900_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3975E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3975_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4000E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4000_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4025E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4025_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4050E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4050_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4100E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4100_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4125E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4125_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4200E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4200_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4225E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4225_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4250E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4250_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4275E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4275_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4300E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4300_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4350E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4350_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4375E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4375_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4400E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4400_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4425E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4425_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4500E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4500_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4550E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4550_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4575E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4575_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4600E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4600_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4625E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4625_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4650E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4650_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4675E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4675_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4700E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4700_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4725E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4725_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4750E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4750_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4800E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4800_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4875E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4875_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4900E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4900_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4950E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4950_E4_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed5000E450g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED5000_E4_50G"
	UpdateVnicShapeDetailsVnicShapeDynamicE550g                    UpdateVnicShapeDetailsVnicShapeEnum = "DYNAMIC_E5_50G"
	UpdateVnicShapeDetailsVnicShapeDynamicE5100g                   UpdateVnicShapeDetailsVnicShapeEnum = "DYNAMIC_E5_100G"
	UpdateVnicShapeDetailsVnicShapeEntirehostE550g                 UpdateVnicShapeDetailsVnicShapeEnum = "EntireHost_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0040E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "Fixed0040_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0100E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0100_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0200E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0200_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0300E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0300_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0400E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0400_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0500E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0500_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0600E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0600_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0700E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0700_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0800E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0800_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0900E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0900_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1000E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1000_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1100E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1100_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1200E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1200_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1300E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1300_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1400E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1400_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1500E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1500_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1600E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1600_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1700E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1700_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1800E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1800_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1900E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1900_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2000E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2000_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2100E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2100_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2200E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2200_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2300E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2300_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2400E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2400_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2500E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2500_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2600E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2600_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2700E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2700_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2800E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2800_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2900E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2900_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3000E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3000_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3100E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3100_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3200E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3200_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3300E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3300_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3400E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3400_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3500E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3500_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3600E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3600_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3700E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3700_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3800E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3800_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3900E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3900_E5_50G"
	UpdateVnicShapeDetailsVnicShapeFixed4000E550g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED4000_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0025E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0025_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0050E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0050_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0075E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0075_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0100E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0100_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0125E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0125_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0150E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0150_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0175E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0175_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0200E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0200_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0225E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0225_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0250E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0250_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0275E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0275_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0300E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0300_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0325E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0325_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0350E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0350_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0375E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0375_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0400E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0400_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0425E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0425_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0450E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0450_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0475E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0475_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0500E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0500_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0525E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0525_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0550E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0550_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0575E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0575_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0600E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0600_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0625E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0625_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0650E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0650_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0675E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0675_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0700E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0700_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0725E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0725_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0750E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0750_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0775E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0775_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0800E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0800_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0825E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0825_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0850E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0850_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0875E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0875_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0900E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0900_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0925E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0925_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0950E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0950_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0975E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0975_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1000E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1000_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1025E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1025_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1050E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1050_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1075E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1075_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1100E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1100_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1125E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1125_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1150E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1150_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1175E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1175_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1200E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1200_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1225E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1225_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1250E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1250_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1275E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1275_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1300E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1300_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1325E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1325_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1350E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1350_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1375E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1375_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1400E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1400_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1425E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1425_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1450E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1450_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1475E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1475_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1500E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1500_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1525E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1525_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1550E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1550_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1575E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1575_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1600E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1600_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1625E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1625_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1650E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1650_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1675E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1675_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1700E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1700_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1725E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1725_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1750E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1750_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1775E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1775_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1800E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1800_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1825E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1825_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1850E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1850_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1875E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1875_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1900E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1900_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1925E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1925_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1950E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1950_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1975E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1975_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2000E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2000_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2025E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2025_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2050E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2050_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2075E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2075_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2100E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2100_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2125E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2125_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2150E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2150_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2175E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2175_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2200E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2200_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2225E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2225_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2250E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2250_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2275E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2275_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2300E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2300_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2325E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2325_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2350E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2350_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2375E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2375_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2400E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2400_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2425E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2425_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2450E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2450_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2475E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2475_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2500E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2500_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2525E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2525_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2550E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2550_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2575E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2575_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2600E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2600_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2625E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2625_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2650E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2650_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2675E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2675_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2700E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2700_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2725E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2725_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2750E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2750_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2775E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2775_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2800E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2800_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2825E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2825_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2850E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2850_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2875E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2875_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2900E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2900_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2925E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2925_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2950E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2950_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2975E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2975_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3000E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3000_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3025E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3025_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3050E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3050_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3075E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3075_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3100E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3100_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3125E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3125_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3150E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3150_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3175E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3175_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3200E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3200_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3225E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3225_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3250E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3250_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3275E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3275_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3300E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3300_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3325E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3325_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3350E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3350_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3375E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3375_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3400E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3400_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3425E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3425_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3450E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3450_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3475E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3475_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3500E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3500_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3525E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3525_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3550E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3550_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3575E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3575_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3600E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3600_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3625E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3625_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3650E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3650_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3675E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3675_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3700E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3700_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3725E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3725_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3750E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3750_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3775E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3775_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3800E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3800_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3825E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3825_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3850E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3850_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3875E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3875_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3900E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3900_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3925E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3925_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3950E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3950_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3975E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3975_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4000E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4000_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4025E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4025_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4050E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4050_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4075E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4075_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4100E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4100_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4125E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4125_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4150E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4150_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4175E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4175_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4200E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4200_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4225E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4225_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4250E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4250_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4275E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4275_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4300E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4300_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4325E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4325_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4350E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4350_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4375E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4375_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4400E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4400_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4425E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4425_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4450E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4450_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4475E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4475_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4500E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4500_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4525E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4525_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4550E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4550_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4575E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4575_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4600E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4600_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4625E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4625_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4650E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4650_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4675E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4675_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4700E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4700_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4725E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4725_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4750E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4750_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4775E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4775_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4800E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4800_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4825E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4825_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4850E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4850_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4875E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4875_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4900E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4900_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4925E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4925_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4950E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4950_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4975E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4975_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed5000E550g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED5000_E5_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0020A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0020_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0040A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0040_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0060A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0060_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0080A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0080_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0100A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0100_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0120A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0120_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0140A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0140_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0160A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0160_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0180A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0180_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0200A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0200_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0220A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0220_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0240A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0240_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0260A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0260_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0280A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0280_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0300A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0300_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0320A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0320_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0340A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0340_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0360A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0360_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0380A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0380_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0400A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0400_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0420A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0420_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0440A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0440_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0460A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0460_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0480A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0480_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0500A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0500_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0520A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0520_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0540A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0540_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0560A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0560_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0580A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0580_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0600A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0600_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0620A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0620_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0640A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0640_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0660A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0660_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0680A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0680_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0700A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0700_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0720A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0720_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0740A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0740_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0760A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0760_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0780A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0780_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0800A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0800_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0820A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0820_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0840A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0840_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0860A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0860_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0880A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0880_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0900A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0900_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0920A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0920_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0940A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0940_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0960A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0960_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0980A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0980_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1000A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1000_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1020A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1020_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1040A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1040_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1060A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1060_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1080A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1080_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1100A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1100_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1120A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1120_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1140A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1140_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1160A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1160_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1180A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1180_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1200A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1200_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1220A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1220_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1240A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1240_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1260A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1260_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1280A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1280_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1300A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1300_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1320A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1320_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1340A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1340_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1360A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1360_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1380A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1380_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1400A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1400_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1420A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1420_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1440A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1440_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1460A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1460_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1480A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1480_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1500A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1500_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1520A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1520_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1540A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1540_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1560A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1560_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1580A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1580_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1600A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1600_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1620A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1620_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1640A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1640_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1660A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1660_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1680A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1680_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1700A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1700_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1720A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1720_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1740A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1740_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1760A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1760_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1780A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1780_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1800A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1800_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1820A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1820_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1840A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1840_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1860A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1860_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1880A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1880_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1900A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1900_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1920A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1920_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1940A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1940_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1960A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1960_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1980A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1980_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2000A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2000_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2020A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2020_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2040A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2040_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2060A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2060_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2080A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2080_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2100A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2100_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2120A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2120_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2140A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2140_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2160A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2160_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2180A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2180_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2200A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2200_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2220A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2220_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2240A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2240_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2260A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2260_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2280A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2280_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2300A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2300_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2320A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2320_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2340A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2340_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2360A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2360_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2380A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2380_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2400A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2400_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2420A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2420_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2440A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2440_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2460A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2460_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2480A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2480_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2500A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2500_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2520A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2520_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2540A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2540_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2560A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2560_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2580A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2580_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2600A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2600_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2620A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2620_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2640A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2640_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2660A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2660_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2680A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2680_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2700A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2700_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2720A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2720_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2740A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2740_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2760A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2760_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2780A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2780_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2800A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2800_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2820A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2820_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2840A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2840_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2860A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2860_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2880A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2880_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2900A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2900_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2920A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2920_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2940A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2940_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2960A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2960_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2980A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2980_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3000A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3000_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3020A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3020_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3040A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3040_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3060A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3060_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3080A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3080_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3100A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3100_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3120A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3120_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3140A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3140_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3160A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3160_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3180A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3180_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3200A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3200_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3220A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3220_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3240A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3240_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3260A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3260_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3280A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3280_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3300A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3300_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3320A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3320_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3340A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3340_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3360A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3360_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3380A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3380_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3400A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3400_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3420A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3420_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3440A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3440_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3460A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3460_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3480A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3480_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3500A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3500_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3520A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3520_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3540A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3540_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3560A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3560_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3580A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3580_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3600A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3600_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3620A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3620_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3640A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3640_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3660A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3660_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3680A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3680_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3700A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3700_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3720A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3720_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3740A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3740_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3760A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3760_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3780A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3780_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3800A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3800_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3820A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3820_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3840A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3840_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3860A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3860_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3880A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3880_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3900A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3900_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3920A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3920_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3940A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3940_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3960A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3960_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3980A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3980_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4000A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4000_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4020A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4020_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4040A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4040_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4060A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4060_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4080A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4080_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4100A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4100_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4120A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4120_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4140A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4140_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4160A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4160_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4180A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4180_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4200A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4200_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4220A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4220_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4240A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4240_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4260A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4260_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4280A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4280_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4300A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4300_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4320A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4320_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4340A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4340_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4360A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4360_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4380A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4380_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4400A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4400_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4420A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4420_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4440A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4440_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4460A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4460_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4480A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4480_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4500A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4500_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4520A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4520_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4540A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4540_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4560A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4560_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4580A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4580_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4600A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4600_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4620A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4620_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4640A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4640_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4660A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4660_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4680A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4680_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4700A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4700_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4720A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4720_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4740A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4740_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4760A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4760_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4780A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4780_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4800A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4800_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4820A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4820_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4840A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4840_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4860A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4860_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4880A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4880_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4900A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4900_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4920A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4920_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4940A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4940_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4960A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4960_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4980A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4980_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed5000A150g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED5000_A1_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0090X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0090_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0180X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0180_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0270X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0270_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0360X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0360_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0450X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0450_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0540X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0540_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0630X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0630_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0720X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0720_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0810X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0810_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0900X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0900_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0990X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0990_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1080X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1080_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1170X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1170_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1260X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1260_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1350X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1350_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1440X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1440_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1530X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1530_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1620X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1620_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1710X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1710_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1800X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1800_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1890X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1890_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1980X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1980_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2070X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2070_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2160X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2160_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2250X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2250_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2340X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2340_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2430X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2430_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2520X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2520_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2610X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2610_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2700X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2700_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2790X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2790_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2880X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2880_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2970X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2970_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3060X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3060_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3150X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3150_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3240X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3240_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3330X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3330_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3420X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3420_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3510X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3510_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3600X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3600_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3690X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3690_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3780X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3780_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3870X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3870_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3960X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3960_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4050X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4050_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4140X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4140_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4230X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4230_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4320X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4320_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4410X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4410_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4500X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4500_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4590X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4590_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4680X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4680_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4770X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4770_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4860X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4860_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4950X950g         UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4950_X9_50G"
	UpdateVnicShapeDetailsVnicShapeDynamicA150g                    UpdateVnicShapeDetailsVnicShapeEnum = "DYNAMIC_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0040A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0040_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0100A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0100_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0200A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0200_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0300A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0300_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0400A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0400_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0500A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0500_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0600A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0600_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0700A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0700_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0800A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0800_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0900A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0900_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1000A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1000_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1100A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1100_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1200A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1200_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1300A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1300_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1400A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1400_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1500A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1500_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1600A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1600_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1700A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1700_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1800A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1800_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1900A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1900_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2000A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2000_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2100A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2100_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2200A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2200_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2300A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2300_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2400A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2400_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2500A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2500_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2600A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2600_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2700A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2700_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2800A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2800_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2900A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2900_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3000A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3000_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3100A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3100_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3200A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3200_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3300A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3300_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3400A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3400_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3500A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3500_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3600A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3600_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3700A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3700_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3800A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3800_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3900A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3900_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed4000A150g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED4000_A1_50G"
	UpdateVnicShapeDetailsVnicShapeFixed5000TelesisA150g           UpdateVnicShapeDetailsVnicShapeEnum = "FIXED5000_TELESIS_A1_50G"
	UpdateVnicShapeDetailsVnicShapeEntirehostA150g                 UpdateVnicShapeDetailsVnicShapeEnum = "ENTIREHOST_A1_50G"
	UpdateVnicShapeDetailsVnicShapeDynamicX725g                    UpdateVnicShapeDetailsVnicShapeEnum = "DYNAMIC_X7_25G"
	UpdateVnicShapeDetailsVnicShapeFixed0040X725g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0040_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed0100X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED0100_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed0200X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED0200_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed0300X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED0300_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed0400X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED0400_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed0500X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED0500_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed0600X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED0600_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed0700X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED0700_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed0800X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED0800_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed0900X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED0900_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed1000X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED1000_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed1100X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED1100_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed1200X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED1200_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed1300X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED1300_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed1400X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED1400_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed1500X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED1500_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed1600X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED1600_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed1700X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED1700_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed1800X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED1800_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed1900X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED1900_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed2000X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED2000_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed2100X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED2100_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed2200X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED2200_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed2300X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED2300_X7_25G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed2400X725g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED2400_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0025X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0025_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0050X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0050_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0075X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0075_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0100X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0100_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0125X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0125_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0150X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0150_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0175X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0175_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0200X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0200_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0225X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0225_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0250X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0250_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0275X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0275_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0300X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0300_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0325X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0325_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0350X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0350_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0375X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0375_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0400X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0400_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0425X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0425_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0450X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0450_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0475X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0475_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0500X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0500_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0525X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0525_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0550X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0550_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0575X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0575_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0600X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0600_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0625X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0625_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0650X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0650_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0675X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0675_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0700X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0700_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0725X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0725_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0750X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0750_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0775X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0775_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0800X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0800_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0825X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0825_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0850X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0850_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0875X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0875_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0900X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0900_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0925X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0925_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0950X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0950_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0975X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0975_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1000X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1000_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1025X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1025_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1050X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1050_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1075X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1075_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1100X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1100_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1125X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1125_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1150X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1150_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1175X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1175_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1200X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1200_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1225X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1225_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1250X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1250_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1275X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1275_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1300X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1300_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1325X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1325_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1350X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1350_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1375X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1375_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1400X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1400_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1425X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1425_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1450X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1450_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1475X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1475_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1500X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1500_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1525X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1525_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1550X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1550_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1575X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1575_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1600X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1600_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1625X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1625_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1650X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1650_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1675X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1675_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1700X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1700_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1725X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1725_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1750X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1750_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1775X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1775_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1800X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1800_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1825X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1825_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1850X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1850_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1875X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1875_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1900X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1900_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1925X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1925_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1950X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1950_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1975X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1975_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2000X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2000_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2025X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2025_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2050X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2050_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2075X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2075_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2100X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2100_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2125X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2125_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2150X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2150_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2175X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2175_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2200X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2200_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2225X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2225_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2250X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2250_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2275X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2275_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2300X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2300_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2325X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2325_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2350X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2350_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2375X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2375_X7_25G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2400X725g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2400_X7_25G"
	UpdateVnicShapeDetailsVnicShapeEntirehostX725g                 UpdateVnicShapeDetailsVnicShapeEnum = "ENTIREHOST_X7_25G"
	UpdateVnicShapeDetailsVnicShapeDynamicX950g                    UpdateVnicShapeDetailsVnicShapeEnum = "DYNAMIC_X9_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0040X950g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0040_X9_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0400X950g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0400_X9_50G"
	UpdateVnicShapeDetailsVnicShapeFixed0800X950g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED0800_X9_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1200X950g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1200_X9_50G"
	UpdateVnicShapeDetailsVnicShapeFixed1600X950g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED1600_X9_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2000X950g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2000_X9_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2400X950g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2400_X9_50G"
	UpdateVnicShapeDetailsVnicShapeFixed2800X950g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED2800_X9_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3200X950g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3200_X9_50G"
	UpdateVnicShapeDetailsVnicShapeFixed3600X950g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED3600_X9_50G"
	UpdateVnicShapeDetailsVnicShapeFixed4000X950g                  UpdateVnicShapeDetailsVnicShapeEnum = "FIXED4000_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed0100X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED0100_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed0200X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED0200_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed0300X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED0300_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed0400X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED0400_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed0500X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED0500_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed0600X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED0600_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed0700X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED0700_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed0800X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED0800_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed0900X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED0900_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed1000X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED1000_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed1100X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED1100_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed1200X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED1200_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed1300X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED1300_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed1400X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED1400_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed1500X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED1500_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed1600X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED1600_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed1700X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED1700_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed1800X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED1800_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed1900X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED1900_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed2000X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED2000_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed2100X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED2100_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed2200X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED2200_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed2300X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED2300_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed2400X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED2400_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed2500X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED2500_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed2600X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED2600_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed2700X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED2700_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed2800X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED2800_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed2900X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED2900_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed3000X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED3000_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed3100X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED3100_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed3200X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED3200_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed3300X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED3300_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed3400X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED3400_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed3500X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED3500_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed3600X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED3600_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed3700X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED3700_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed3800X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED3800_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed3900X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED3900_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed4000X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED4000_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed4100X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED4100_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed4200X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED4200_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed4300X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED4300_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed4400X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED4400_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed4500X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED4500_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed4600X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED4600_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed4700X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED4700_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed4800X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED4800_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed4900X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED4900_X9_50G"
	UpdateVnicShapeDetailsVnicShapeStandardVmFixed5000X950g        UpdateVnicShapeDetailsVnicShapeEnum = "STANDARD_VM_FIXED5000_X9_50G"
	UpdateVnicShapeDetailsVnicShapeA10GpuVmFixed2400X950g          UpdateVnicShapeDetailsVnicShapeEnum = "A10_GPU_VM_FIXED2400_X9_50G"
	UpdateVnicShapeDetailsVnicShapeA10GpuVmFixed4800X950g          UpdateVnicShapeDetailsVnicShapeEnum = "A10_GPU_VM_FIXED4800_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed01001X950g          UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED0100_1_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed02002X950g          UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED0200_2_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed03003X950g          UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED0300_3_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed04004X950g          UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED0400_4_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed05005X950g          UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED0500_5_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed06006X950g          UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED0600_6_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed07007X950g          UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED0700_7_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed08008X950g          UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED0800_8_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed09009X950g          UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED0900_9_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed100010X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED1000_10_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed110011X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED1100_11_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed120012X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED1200_12_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed130013X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED1300_13_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed140014X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED1400_14_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed150015X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED1500_15_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed160016X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED1600_16_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed170017X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED1700_17_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed180018X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED1800_18_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed190019X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED1900_19_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed200020X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED2000_20_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed210021X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED2100_21_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed220022X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED2200_22_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed230023X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED2300_23_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed240024X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED2400_24_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed250025X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED2500_25_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed260026X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED2600_26_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed270027X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED2700_27_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed280028X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED2800_28_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed290029X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED2900_29_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed300030X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3000_30_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed310031X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3100_31_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320032X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_32_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320033X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_33_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320034X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_34_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320035X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_35_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320036X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_36_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320037X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_37_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320038X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_38_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320039X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_39_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320040X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_40_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320041X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_41_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320042X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_42_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320043X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_43_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320044X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_44_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320045X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_45_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320046X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_46_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320047X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_47_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320048X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_48_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320049X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_49_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320050X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_50_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320051X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_51_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320052X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_52_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320053X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_53_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320054X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_54_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320055X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_55_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320056X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_56_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320057X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_57_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320058X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_58_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320059X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_59_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320060X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_60_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320061X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_61_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320062X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_62_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320063X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_63_X9_50G"
	UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320064X950g         UpdateVnicShapeDetailsVnicShapeEnum = "VNUMA_VM_FIXED3200_64_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0025X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0025_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0050X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0050_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0075X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0075_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0100X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0100_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0125X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0125_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0150X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0150_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0175X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0175_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0200X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0200_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0225X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0225_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0250X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0250_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0275X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0275_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0300X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0300_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0325X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0325_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0350X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0350_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0375X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0375_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0400X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0400_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0425X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0425_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0450X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0450_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0475X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0475_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0500X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0500_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0525X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0525_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0550X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0550_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0575X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0575_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0600X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0600_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0625X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0625_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0650X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0650_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0675X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0675_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0700X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0700_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0725X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0725_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0750X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0750_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0775X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0775_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0800X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0800_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0825X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0825_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0850X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0850_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0875X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0875_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0900X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0900_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0925X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0925_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0950X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0950_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0975X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0975_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1000X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1000_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1025X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1025_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1050X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1050_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1075X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1075_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1100X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1100_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1125X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1125_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1150X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1150_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1175X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1175_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1200X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1200_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1225X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1225_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1250X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1250_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1275X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1275_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1300X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1300_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1325X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1325_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1350X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1350_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1375X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1375_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1400X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1400_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1425X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1425_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1450X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1450_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1475X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1475_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1500X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1500_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1525X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1525_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1550X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1550_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1575X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1575_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1600X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1600_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1625X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1625_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1650X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1650_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1700X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1700_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1725X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1725_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1750X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1750_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1800X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1800_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1850X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1850_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1875X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1875_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1900X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1900_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1925X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1925_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1950X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1950_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2000X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2000_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2025X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2025_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2050X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2050_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2100X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2100_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2125X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2125_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2150X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2150_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2175X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2175_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2200X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2200_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2250X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2250_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2275X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2275_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2300X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2300_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2325X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2325_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2350X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2350_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2375X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2375_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2400X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2400_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2450X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2450_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2475X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2475_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2500X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2500_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2550X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2550_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2600X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2600_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2625X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2625_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2650X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2650_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2700X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2700_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2750X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2750_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2775X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2775_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2800X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2800_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2850X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2850_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2875X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2875_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2900X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2900_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2925X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2925_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2950X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2950_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2975X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2975_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3000X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3000_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3025X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3025_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3050X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3050_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3075X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3075_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3100X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3100_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3125X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3125_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3150X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3150_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3200X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3200_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3225X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3225_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3250X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3250_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3300X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3300_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3325X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3325_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3375X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3375_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3400X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3400_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3450X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3450_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3500X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3500_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3525X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3525_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3575X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3575_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3600X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3600_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3625X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3625_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3675X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3675_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3700X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3700_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3750X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3750_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3800X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3800_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3825X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3825_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3850X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3850_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3875X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3875_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3900X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3900_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3975X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3975_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4000X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4000_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4025X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4025_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4050X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4050_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4100X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4100_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4125X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4125_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4200X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4200_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4225X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4225_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4250X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4250_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4275X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4275_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4300X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4300_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4350X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4350_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4375X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4375_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4400X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4400_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4425X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4425_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4500X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4500_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4550X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4550_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4575X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4575_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4600X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4600_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4625X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4625_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4650X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4650_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4675X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4675_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4700X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4700_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4725X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4725_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4750X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4750_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4800X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4800_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4875X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4875_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4900X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4900_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4950X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4950_X9_50G"
	UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed5000X950g UpdateVnicShapeDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED5000_X9_50G"
	UpdateVnicShapeDetailsVnicShapeEntirehostX950g                 UpdateVnicShapeDetailsVnicShapeEnum = "ENTIREHOST_X9_50G"
	UpdateVnicShapeDetailsVnicShapeDynamicX9100g                   UpdateVnicShapeDetailsVnicShapeEnum = "DYNAMIC_X9_100G"
	UpdateVnicShapeDetailsVnicShapeDynamicX1050g                   UpdateVnicShapeDetailsVnicShapeEnum = "DYNAMIC_X10_50G"
	UpdateVnicShapeDetailsVnicShapeDynamicX10100g                  UpdateVnicShapeDetailsVnicShapeEnum = "DYNAMIC_X10_100G"
)

var mappingUpdateVnicShapeDetailsVnicShapeEnum = map[string]UpdateVnicShapeDetailsVnicShapeEnum{
	"DYNAMIC":                              UpdateVnicShapeDetailsVnicShapeDynamic,
	"FIXED0040":                            UpdateVnicShapeDetailsVnicShapeFixed0040,
	"FIXED0060":                            UpdateVnicShapeDetailsVnicShapeFixed0060,
	"FIXED0060_PSM":                        UpdateVnicShapeDetailsVnicShapeFixed0060Psm,
	"FIXED0100":                            UpdateVnicShapeDetailsVnicShapeFixed0100,
	"FIXED0120":                            UpdateVnicShapeDetailsVnicShapeFixed0120,
	"FIXED0120_2X":                         UpdateVnicShapeDetailsVnicShapeFixed01202x,
	"FIXED0200":                            UpdateVnicShapeDetailsVnicShapeFixed0200,
	"FIXED0240":                            UpdateVnicShapeDetailsVnicShapeFixed0240,
	"FIXED0480":                            UpdateVnicShapeDetailsVnicShapeFixed0480,
	"ENTIREHOST":                           UpdateVnicShapeDetailsVnicShapeEntirehost,
	"DYNAMIC_25G":                          UpdateVnicShapeDetailsVnicShapeDynamic25g,
	"FIXED0040_25G":                        UpdateVnicShapeDetailsVnicShapeFixed004025g,
	"FIXED0100_25G":                        UpdateVnicShapeDetailsVnicShapeFixed010025g,
	"FIXED0200_25G":                        UpdateVnicShapeDetailsVnicShapeFixed020025g,
	"FIXED0400_25G":                        UpdateVnicShapeDetailsVnicShapeFixed040025g,
	"FIXED0800_25G":                        UpdateVnicShapeDetailsVnicShapeFixed080025g,
	"FIXED1600_25G":                        UpdateVnicShapeDetailsVnicShapeFixed160025g,
	"FIXED2400_25G":                        UpdateVnicShapeDetailsVnicShapeFixed240025g,
	"ENTIREHOST_25G":                       UpdateVnicShapeDetailsVnicShapeEntirehost25g,
	"DYNAMIC_E1_25G":                       UpdateVnicShapeDetailsVnicShapeDynamicE125g,
	"FIXED0040_E1_25G":                     UpdateVnicShapeDetailsVnicShapeFixed0040E125g,
	"FIXED0070_E1_25G":                     UpdateVnicShapeDetailsVnicShapeFixed0070E125g,
	"FIXED0140_E1_25G":                     UpdateVnicShapeDetailsVnicShapeFixed0140E125g,
	"FIXED0280_E1_25G":                     UpdateVnicShapeDetailsVnicShapeFixed0280E125g,
	"FIXED0560_E1_25G":                     UpdateVnicShapeDetailsVnicShapeFixed0560E125g,
	"FIXED1120_E1_25G":                     UpdateVnicShapeDetailsVnicShapeFixed1120E125g,
	"FIXED1680_E1_25G":                     UpdateVnicShapeDetailsVnicShapeFixed1680E125g,
	"ENTIREHOST_E1_25G":                    UpdateVnicShapeDetailsVnicShapeEntirehostE125g,
	"DYNAMIC_B1_25G":                       UpdateVnicShapeDetailsVnicShapeDynamicB125g,
	"FIXED0040_B1_25G":                     UpdateVnicShapeDetailsVnicShapeFixed0040B125g,
	"FIXED0060_B1_25G":                     UpdateVnicShapeDetailsVnicShapeFixed0060B125g,
	"FIXED0120_B1_25G":                     UpdateVnicShapeDetailsVnicShapeFixed0120B125g,
	"FIXED0240_B1_25G":                     UpdateVnicShapeDetailsVnicShapeFixed0240B125g,
	"FIXED0480_B1_25G":                     UpdateVnicShapeDetailsVnicShapeFixed0480B125g,
	"FIXED0960_B1_25G":                     UpdateVnicShapeDetailsVnicShapeFixed0960B125g,
	"ENTIREHOST_B1_25G":                    UpdateVnicShapeDetailsVnicShapeEntirehostB125g,
	"MICRO_VM_FIXED0048_E1_25G":            UpdateVnicShapeDetailsVnicShapeMicroVmFixed0048E125g,
	"MICRO_LB_FIXED0001_E1_25G":            UpdateVnicShapeDetailsVnicShapeMicroLbFixed0001E125g,
	"VNICAAS_FIXED0025":                    UpdateVnicShapeDetailsVnicShapeVnicaasFixed0025,
	"VNICAAS_FIXED0200":                    UpdateVnicShapeDetailsVnicShapeVnicaasFixed0200,
	"VNICAAS_FIXED0400":                    UpdateVnicShapeDetailsVnicShapeVnicaasFixed0400,
	"VNICAAS_FIXED0625":                    UpdateVnicShapeDetailsVnicShapeVnicaasFixed0625,
	"VNICAAS_FIXED0700":                    UpdateVnicShapeDetailsVnicShapeVnicaasFixed0700,
	"VNICAAS_FIXED1250":                    UpdateVnicShapeDetailsVnicShapeVnicaasFixed1250,
	"VNICAAS_FWAAS_6G_250M":                UpdateVnicShapeDetailsVnicShapeVnicaasFwaas6g250m,
	"VNICAAS_NLB_APPROVED_10G":             UpdateVnicShapeDetailsVnicShapeVnicaasNlbApproved10g,
	"VNICAAS_NLB_APPROVED_25G":             UpdateVnicShapeDetailsVnicShapeVnicaasNlbApproved25g,
	"VNICAAS_TELESIS_25G":                  UpdateVnicShapeDetailsVnicShapeVnicaasTelesis25g,
	"VNICAAS_TELESIS_10G":                  UpdateVnicShapeDetailsVnicShapeVnicaasTelesis10g,
	"VNICAAS_AMBASSADOR_FIXED0100":         UpdateVnicShapeDetailsVnicShapeVnicaasAmbassadorFixed0100,
	"VNICAAS_TELESIS_GAMMA":                UpdateVnicShapeDetailsVnicShapeVnicaasTelesisGamma,
	"VNICAAS_PRIVATEDNS":                   UpdateVnicShapeDetailsVnicShapeVnicaasPrivatedns,
	"VNICAAS_FWAAS":                        UpdateVnicShapeDetailsVnicShapeVnicaasFwaas,
	"VNICAAS_LBAAS_FREE":                   UpdateVnicShapeDetailsVnicShapeVnicaasLbaasFree,
	"VNICAAS_LBAAS_8G_512K":                UpdateVnicShapeDetailsVnicShapeVnicaasLbaas8g512k,
	"VNICAAS_LBAAS_8G_1M":                  UpdateVnicShapeDetailsVnicShapeVnicaasLbaas8g1m,
	"VNICAAS_LBAAS_8G_2M":                  UpdateVnicShapeDetailsVnicShapeVnicaasLbaas8g2m,
	"VNICAAS_LBAAS_8G_3M":                  UpdateVnicShapeDetailsVnicShapeVnicaasLbaas8g3m,
	"VNICAAS_LBAAS_8G_1M_8GHOST":           UpdateVnicShapeDetailsVnicShapeVnicaasLbaas8g1m8ghost,
	"VNICAAS_LBAAS_8G_1M_16GHOST":          UpdateVnicShapeDetailsVnicShapeVnicaasLbaas8g1m16ghost,
	"DYNAMIC_E3_50G":                       UpdateVnicShapeDetailsVnicShapeDynamicE350g,
	"FIXED0040_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0040E350g,
	"FIXED0100_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0100E350g,
	"FIXED0200_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0200E350g,
	"FIXED0300_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0300E350g,
	"FIXED0400_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0400E350g,
	"FIXED0500_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0500E350g,
	"FIXED0600_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0600E350g,
	"FIXED0700_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0700E350g,
	"FIXED0800_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0800E350g,
	"FIXED0900_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0900E350g,
	"FIXED1000_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1000E350g,
	"FIXED1100_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1100E350g,
	"FIXED1200_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1200E350g,
	"FIXED1300_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1300E350g,
	"FIXED1400_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1400E350g,
	"FIXED1500_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1500E350g,
	"FIXED1600_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1600E350g,
	"FIXED1700_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1700E350g,
	"FIXED1800_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1800E350g,
	"FIXED1900_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1900E350g,
	"FIXED2000_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2000E350g,
	"FIXED2100_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2100E350g,
	"FIXED2200_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2200E350g,
	"FIXED2300_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2300E350g,
	"FIXED2400_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2400E350g,
	"FIXED2500_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2500E350g,
	"FIXED2600_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2600E350g,
	"FIXED2700_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2700E350g,
	"FIXED2800_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2800E350g,
	"FIXED2900_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2900E350g,
	"FIXED3000_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3000E350g,
	"FIXED3100_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3100E350g,
	"FIXED3200_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3200E350g,
	"FIXED3300_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3300E350g,
	"FIXED3400_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3400E350g,
	"FIXED3500_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3500E350g,
	"FIXED3600_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3600E350g,
	"FIXED3700_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3700E350g,
	"FIXED3800_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3800E350g,
	"FIXED3900_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3900E350g,
	"FIXED4000_E3_50G":                     UpdateVnicShapeDetailsVnicShapeFixed4000E350g,
	"ENTIREHOST_E3_50G":                    UpdateVnicShapeDetailsVnicShapeEntirehostE350g,
	"DYNAMIC_E4_50G":                       UpdateVnicShapeDetailsVnicShapeDynamicE450g,
	"FIXED0040_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0040E450g,
	"FIXED0100_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0100E450g,
	"FIXED0200_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0200E450g,
	"FIXED0300_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0300E450g,
	"FIXED0400_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0400E450g,
	"FIXED0500_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0500E450g,
	"FIXED0600_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0600E450g,
	"FIXED0700_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0700E450g,
	"FIXED0800_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0800E450g,
	"FIXED0900_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0900E450g,
	"FIXED1000_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1000E450g,
	"FIXED1100_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1100E450g,
	"FIXED1200_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1200E450g,
	"FIXED1300_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1300E450g,
	"FIXED1400_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1400E450g,
	"FIXED1500_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1500E450g,
	"FIXED1600_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1600E450g,
	"FIXED1700_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1700E450g,
	"FIXED1800_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1800E450g,
	"FIXED1900_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1900E450g,
	"FIXED2000_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2000E450g,
	"FIXED2100_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2100E450g,
	"FIXED2200_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2200E450g,
	"FIXED2300_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2300E450g,
	"FIXED2400_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2400E450g,
	"FIXED2500_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2500E450g,
	"FIXED2600_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2600E450g,
	"FIXED2700_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2700E450g,
	"FIXED2800_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2800E450g,
	"FIXED2900_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2900E450g,
	"FIXED3000_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3000E450g,
	"FIXED3100_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3100E450g,
	"FIXED3200_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3200E450g,
	"FIXED3300_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3300E450g,
	"FIXED3400_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3400E450g,
	"FIXED3500_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3500E450g,
	"FIXED3600_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3600E450g,
	"FIXED3700_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3700E450g,
	"FIXED3800_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3800E450g,
	"FIXED3900_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3900E450g,
	"FIXED4000_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed4000E450g,
	"FIXED4700_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed4700E450g,
	"FIXED4960_E4_50G":                     UpdateVnicShapeDetailsVnicShapeFixed4960E450g,
	"ENTIREHOST_E4_50G":                    UpdateVnicShapeDetailsVnicShapeEntirehostE450g,
	"VNUMA_VM_FIXED0100_1_E4_50G":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed01001E450g,
	"VNUMA_VM_FIXED0200_2_E4_50G":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed02002E450g,
	"VNUMA_VM_FIXED0300_3_E4_50G":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed03003E450g,
	"VNUMA_VM_FIXED0400_4_E4_50G":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed04004E450g,
	"VNUMA_VM_FIXED0500_5_E4_50G":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed05005E450g,
	"VNUMA_VM_FIXED0600_6_E4_50G":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed06006E450g,
	"VNUMA_VM_FIXED0700_7_E4_50G":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed07007E450g,
	"VNUMA_VM_FIXED0800_8_E4_50G":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed08008E450g,
	"VNUMA_VM_FIXED0900_9_E4_50G":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed09009E450g,
	"VNUMA_VM_FIXED1000_10_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed100010E450g,
	"VNUMA_VM_FIXED1100_11_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed110011E450g,
	"VNUMA_VM_FIXED1200_12_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed120012E450g,
	"VNUMA_VM_FIXED1300_13_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed130013E450g,
	"VNUMA_VM_FIXED1400_14_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed140014E450g,
	"VNUMA_VM_FIXED1500_15_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed150015E450g,
	"VNUMA_VM_FIXED1600_16_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed160016E450g,
	"VNUMA_VM_FIXED1700_17_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed170017E450g,
	"VNUMA_VM_FIXED1800_18_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed180018E450g,
	"VNUMA_VM_FIXED1900_19_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed190019E450g,
	"VNUMA_VM_FIXED2000_20_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed200020E450g,
	"VNUMA_VM_FIXED2100_21_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed210021E450g,
	"VNUMA_VM_FIXED2200_22_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed220022E450g,
	"VNUMA_VM_FIXED2300_23_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed230023E450g,
	"VNUMA_VM_FIXED2400_24_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed240024E450g,
	"VNUMA_VM_FIXED2500_25_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed250025E450g,
	"VNUMA_VM_FIXED2600_26_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed260026E450g,
	"VNUMA_VM_FIXED2700_27_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed270027E450g,
	"VNUMA_VM_FIXED2800_28_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed280028E450g,
	"VNUMA_VM_FIXED2900_29_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed290029E450g,
	"VNUMA_VM_FIXED3000_30_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed300030E450g,
	"VNUMA_VM_FIXED3100_31_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed310031E450g,
	"VNUMA_VM_FIXED3200_32_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320032E450g,
	"VNUMA_VM_FIXED3300_33_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed330033E450g,
	"VNUMA_VM_FIXED3400_34_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed340034E450g,
	"VNUMA_VM_FIXED3500_35_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed350035E450g,
	"VNUMA_VM_FIXED3600_36_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed360036E450g,
	"VNUMA_VM_FIXED3700_37_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed370037E450g,
	"VNUMA_VM_FIXED3800_38_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed380038E450g,
	"VNUMA_VM_FIXED3900_39_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed390039E450g,
	"VNUMA_VM_FIXED4000_40_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400040E450g,
	"VNUMA_VM_FIXED4000_41_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400041E450g,
	"VNUMA_VM_FIXED4000_42_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400042E450g,
	"VNUMA_VM_FIXED4000_43_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400043E450g,
	"VNUMA_VM_FIXED4000_44_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400044E450g,
	"VNUMA_VM_FIXED4000_45_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400045E450g,
	"VNUMA_VM_FIXED4000_46_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400046E450g,
	"VNUMA_VM_FIXED4000_47_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400047E450g,
	"VNUMA_VM_FIXED4000_48_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400048E450g,
	"VNUMA_VM_FIXED4000_49_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400049E450g,
	"VNUMA_VM_FIXED4000_50_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400050E450g,
	"VNUMA_VM_FIXED4000_51_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400051E450g,
	"VNUMA_VM_FIXED4000_52_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400052E450g,
	"VNUMA_VM_FIXED4000_53_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400053E450g,
	"VNUMA_VM_FIXED4000_54_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400054E450g,
	"VNUMA_VM_FIXED4000_55_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400055E450g,
	"VNUMA_VM_FIXED4000_56_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400056E450g,
	"VNUMA_VM_FIXED4000_57_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400057E450g,
	"VNUMA_VM_FIXED4000_58_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400058E450g,
	"VNUMA_VM_FIXED4000_59_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400059E450g,
	"VNUMA_VM_FIXED4000_60_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400060E450g,
	"VNUMA_VM_FIXED4000_61_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400061E450g,
	"VNUMA_VM_FIXED4000_62_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400062E450g,
	"VNUMA_VM_FIXED4000_63_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400063E450g,
	"VNUMA_VM_FIXED4000_64_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400064E450g,
	"VNUMA_VM_FIXED4000_65_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400065E450g,
	"VNUMA_VM_FIXED4000_66_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400066E450g,
	"VNUMA_VM_FIXED4000_67_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400067E450g,
	"VNUMA_VM_FIXED4000_68_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400068E450g,
	"VNUMA_VM_FIXED4000_69_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400069E450g,
	"VNUMA_VM_FIXED4000_70_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400070E450g,
	"VNUMA_VM_FIXED4000_71_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400071E450g,
	"VNUMA_VM_FIXED4000_72_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400072E450g,
	"VNUMA_VM_FIXED4000_73_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400073E450g,
	"VNUMA_VM_FIXED4000_74_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400074E450g,
	"VNUMA_VM_FIXED4000_75_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400075E450g,
	"VNUMA_VM_FIXED4000_76_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400076E450g,
	"VNUMA_VM_FIXED4000_77_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400077E450g,
	"VNUMA_VM_FIXED4000_78_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400078E450g,
	"VNUMA_VM_FIXED4000_79_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400079E450g,
	"VNUMA_VM_FIXED4000_80_E4_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400080E450g,
	"Micro_VM_Fixed0050_E3_50G":            UpdateVnicShapeDetailsVnicShapeMicroVmFixed0050E350g,
	"Micro_VM_Fixed0050_E4_50G":            UpdateVnicShapeDetailsVnicShapeMicroVmFixed0050E450g,
	"SUBCORE_VM_FIXED0025_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0025E350g,
	"SUBCORE_VM_FIXED0050_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0050E350g,
	"SUBCORE_VM_FIXED0075_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0075E350g,
	"SUBCORE_VM_FIXED0100_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0100E350g,
	"SUBCORE_VM_FIXED0125_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0125E350g,
	"SUBCORE_VM_FIXED0150_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0150E350g,
	"SUBCORE_VM_FIXED0175_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0175E350g,
	"SUBCORE_VM_FIXED0200_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0200E350g,
	"SUBCORE_VM_FIXED0225_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0225E350g,
	"SUBCORE_VM_FIXED0250_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0250E350g,
	"SUBCORE_VM_FIXED0275_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0275E350g,
	"SUBCORE_VM_FIXED0300_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0300E350g,
	"SUBCORE_VM_FIXED0325_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0325E350g,
	"SUBCORE_VM_FIXED0350_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0350E350g,
	"SUBCORE_VM_FIXED0375_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0375E350g,
	"SUBCORE_VM_FIXED0400_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0400E350g,
	"SUBCORE_VM_FIXED0425_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0425E350g,
	"SUBCORE_VM_FIXED0450_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0450E350g,
	"SUBCORE_VM_FIXED0475_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0475E350g,
	"SUBCORE_VM_FIXED0500_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0500E350g,
	"SUBCORE_VM_FIXED0525_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0525E350g,
	"SUBCORE_VM_FIXED0550_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0550E350g,
	"SUBCORE_VM_FIXED0575_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0575E350g,
	"SUBCORE_VM_FIXED0600_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0600E350g,
	"SUBCORE_VM_FIXED0625_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0625E350g,
	"SUBCORE_VM_FIXED0650_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0650E350g,
	"SUBCORE_VM_FIXED0675_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0675E350g,
	"SUBCORE_VM_FIXED0700_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0700E350g,
	"SUBCORE_VM_FIXED0725_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0725E350g,
	"SUBCORE_VM_FIXED0750_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0750E350g,
	"SUBCORE_VM_FIXED0775_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0775E350g,
	"SUBCORE_VM_FIXED0800_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0800E350g,
	"SUBCORE_VM_FIXED0825_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0825E350g,
	"SUBCORE_VM_FIXED0850_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0850E350g,
	"SUBCORE_VM_FIXED0875_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0875E350g,
	"SUBCORE_VM_FIXED0900_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0900E350g,
	"SUBCORE_VM_FIXED0925_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0925E350g,
	"SUBCORE_VM_FIXED0950_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0950E350g,
	"SUBCORE_VM_FIXED0975_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0975E350g,
	"SUBCORE_VM_FIXED1000_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1000E350g,
	"SUBCORE_VM_FIXED1025_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1025E350g,
	"SUBCORE_VM_FIXED1050_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1050E350g,
	"SUBCORE_VM_FIXED1075_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1075E350g,
	"SUBCORE_VM_FIXED1100_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1100E350g,
	"SUBCORE_VM_FIXED1125_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1125E350g,
	"SUBCORE_VM_FIXED1150_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1150E350g,
	"SUBCORE_VM_FIXED1175_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1175E350g,
	"SUBCORE_VM_FIXED1200_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1200E350g,
	"SUBCORE_VM_FIXED1225_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1225E350g,
	"SUBCORE_VM_FIXED1250_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1250E350g,
	"SUBCORE_VM_FIXED1275_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1275E350g,
	"SUBCORE_VM_FIXED1300_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1300E350g,
	"SUBCORE_VM_FIXED1325_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1325E350g,
	"SUBCORE_VM_FIXED1350_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1350E350g,
	"SUBCORE_VM_FIXED1375_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1375E350g,
	"SUBCORE_VM_FIXED1400_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1400E350g,
	"SUBCORE_VM_FIXED1425_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1425E350g,
	"SUBCORE_VM_FIXED1450_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1450E350g,
	"SUBCORE_VM_FIXED1475_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1475E350g,
	"SUBCORE_VM_FIXED1500_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1500E350g,
	"SUBCORE_VM_FIXED1525_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1525E350g,
	"SUBCORE_VM_FIXED1550_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1550E350g,
	"SUBCORE_VM_FIXED1575_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1575E350g,
	"SUBCORE_VM_FIXED1600_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1600E350g,
	"SUBCORE_VM_FIXED1625_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1625E350g,
	"SUBCORE_VM_FIXED1650_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1650E350g,
	"SUBCORE_VM_FIXED1700_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1700E350g,
	"SUBCORE_VM_FIXED1725_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1725E350g,
	"SUBCORE_VM_FIXED1750_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1750E350g,
	"SUBCORE_VM_FIXED1800_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1800E350g,
	"SUBCORE_VM_FIXED1850_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1850E350g,
	"SUBCORE_VM_FIXED1875_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1875E350g,
	"SUBCORE_VM_FIXED1900_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1900E350g,
	"SUBCORE_VM_FIXED1925_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1925E350g,
	"SUBCORE_VM_FIXED1950_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1950E350g,
	"SUBCORE_VM_FIXED2000_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2000E350g,
	"SUBCORE_VM_FIXED2025_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2025E350g,
	"SUBCORE_VM_FIXED2050_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2050E350g,
	"SUBCORE_VM_FIXED2100_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2100E350g,
	"SUBCORE_VM_FIXED2125_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2125E350g,
	"SUBCORE_VM_FIXED2150_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2150E350g,
	"SUBCORE_VM_FIXED2175_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2175E350g,
	"SUBCORE_VM_FIXED2200_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2200E350g,
	"SUBCORE_VM_FIXED2250_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2250E350g,
	"SUBCORE_VM_FIXED2275_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2275E350g,
	"SUBCORE_VM_FIXED2300_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2300E350g,
	"SUBCORE_VM_FIXED2325_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2325E350g,
	"SUBCORE_VM_FIXED2350_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2350E350g,
	"SUBCORE_VM_FIXED2375_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2375E350g,
	"SUBCORE_VM_FIXED2400_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2400E350g,
	"SUBCORE_VM_FIXED2450_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2450E350g,
	"SUBCORE_VM_FIXED2475_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2475E350g,
	"SUBCORE_VM_FIXED2500_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2500E350g,
	"SUBCORE_VM_FIXED2550_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2550E350g,
	"SUBCORE_VM_FIXED2600_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2600E350g,
	"SUBCORE_VM_FIXED2625_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2625E350g,
	"SUBCORE_VM_FIXED2650_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2650E350g,
	"SUBCORE_VM_FIXED2700_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2700E350g,
	"SUBCORE_VM_FIXED2750_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2750E350g,
	"SUBCORE_VM_FIXED2775_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2775E350g,
	"SUBCORE_VM_FIXED2800_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2800E350g,
	"SUBCORE_VM_FIXED2850_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2850E350g,
	"SUBCORE_VM_FIXED2875_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2875E350g,
	"SUBCORE_VM_FIXED2900_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2900E350g,
	"SUBCORE_VM_FIXED2925_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2925E350g,
	"SUBCORE_VM_FIXED2950_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2950E350g,
	"SUBCORE_VM_FIXED2975_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2975E350g,
	"SUBCORE_VM_FIXED3000_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3000E350g,
	"SUBCORE_VM_FIXED3025_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3025E350g,
	"SUBCORE_VM_FIXED3050_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3050E350g,
	"SUBCORE_VM_FIXED3075_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3075E350g,
	"SUBCORE_VM_FIXED3100_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3100E350g,
	"SUBCORE_VM_FIXED3125_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3125E350g,
	"SUBCORE_VM_FIXED3150_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3150E350g,
	"SUBCORE_VM_FIXED3200_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3200E350g,
	"SUBCORE_VM_FIXED3225_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3225E350g,
	"SUBCORE_VM_FIXED3250_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3250E350g,
	"SUBCORE_VM_FIXED3300_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3300E350g,
	"SUBCORE_VM_FIXED3325_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3325E350g,
	"SUBCORE_VM_FIXED3375_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3375E350g,
	"SUBCORE_VM_FIXED3400_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3400E350g,
	"SUBCORE_VM_FIXED3450_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3450E350g,
	"SUBCORE_VM_FIXED3500_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3500E350g,
	"SUBCORE_VM_FIXED3525_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3525E350g,
	"SUBCORE_VM_FIXED3575_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3575E350g,
	"SUBCORE_VM_FIXED3600_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3600E350g,
	"SUBCORE_VM_FIXED3625_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3625E350g,
	"SUBCORE_VM_FIXED3675_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3675E350g,
	"SUBCORE_VM_FIXED3700_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3700E350g,
	"SUBCORE_VM_FIXED3750_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3750E350g,
	"SUBCORE_VM_FIXED3800_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3800E350g,
	"SUBCORE_VM_FIXED3825_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3825E350g,
	"SUBCORE_VM_FIXED3850_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3850E350g,
	"SUBCORE_VM_FIXED3875_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3875E350g,
	"SUBCORE_VM_FIXED3900_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3900E350g,
	"SUBCORE_VM_FIXED3975_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3975E350g,
	"SUBCORE_VM_FIXED4000_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4000E350g,
	"SUBCORE_VM_FIXED4025_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4025E350g,
	"SUBCORE_VM_FIXED4050_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4050E350g,
	"SUBCORE_VM_FIXED4100_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4100E350g,
	"SUBCORE_VM_FIXED4125_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4125E350g,
	"SUBCORE_VM_FIXED4200_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4200E350g,
	"SUBCORE_VM_FIXED4225_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4225E350g,
	"SUBCORE_VM_FIXED4250_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4250E350g,
	"SUBCORE_VM_FIXED4275_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4275E350g,
	"SUBCORE_VM_FIXED4300_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4300E350g,
	"SUBCORE_VM_FIXED4350_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4350E350g,
	"SUBCORE_VM_FIXED4375_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4375E350g,
	"SUBCORE_VM_FIXED4400_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4400E350g,
	"SUBCORE_VM_FIXED4425_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4425E350g,
	"SUBCORE_VM_FIXED4500_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4500E350g,
	"SUBCORE_VM_FIXED4550_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4550E350g,
	"SUBCORE_VM_FIXED4575_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4575E350g,
	"SUBCORE_VM_FIXED4600_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4600E350g,
	"SUBCORE_VM_FIXED4625_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4625E350g,
	"SUBCORE_VM_FIXED4650_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4650E350g,
	"SUBCORE_VM_FIXED4675_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4675E350g,
	"SUBCORE_VM_FIXED4700_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4700E350g,
	"SUBCORE_VM_FIXED4725_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4725E350g,
	"SUBCORE_VM_FIXED4750_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4750E350g,
	"SUBCORE_VM_FIXED4800_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4800E350g,
	"SUBCORE_VM_FIXED4875_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4875E350g,
	"SUBCORE_VM_FIXED4900_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4900E350g,
	"SUBCORE_VM_FIXED4950_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4950E350g,
	"SUBCORE_VM_FIXED5000_E3_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed5000E350g,
	"SUBCORE_VM_FIXED0025_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0025E450g,
	"SUBCORE_VM_FIXED0050_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0050E450g,
	"SUBCORE_VM_FIXED0075_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0075E450g,
	"SUBCORE_VM_FIXED0100_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0100E450g,
	"SUBCORE_VM_FIXED0125_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0125E450g,
	"SUBCORE_VM_FIXED0150_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0150E450g,
	"SUBCORE_VM_FIXED0175_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0175E450g,
	"SUBCORE_VM_FIXED0200_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0200E450g,
	"SUBCORE_VM_FIXED0225_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0225E450g,
	"SUBCORE_VM_FIXED0250_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0250E450g,
	"SUBCORE_VM_FIXED0275_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0275E450g,
	"SUBCORE_VM_FIXED0300_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0300E450g,
	"SUBCORE_VM_FIXED0325_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0325E450g,
	"SUBCORE_VM_FIXED0350_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0350E450g,
	"SUBCORE_VM_FIXED0375_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0375E450g,
	"SUBCORE_VM_FIXED0400_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0400E450g,
	"SUBCORE_VM_FIXED0425_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0425E450g,
	"SUBCORE_VM_FIXED0450_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0450E450g,
	"SUBCORE_VM_FIXED0475_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0475E450g,
	"SUBCORE_VM_FIXED0500_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0500E450g,
	"SUBCORE_VM_FIXED0525_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0525E450g,
	"SUBCORE_VM_FIXED0550_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0550E450g,
	"SUBCORE_VM_FIXED0575_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0575E450g,
	"SUBCORE_VM_FIXED0600_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0600E450g,
	"SUBCORE_VM_FIXED0625_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0625E450g,
	"SUBCORE_VM_FIXED0650_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0650E450g,
	"SUBCORE_VM_FIXED0675_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0675E450g,
	"SUBCORE_VM_FIXED0700_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0700E450g,
	"SUBCORE_VM_FIXED0725_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0725E450g,
	"SUBCORE_VM_FIXED0750_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0750E450g,
	"SUBCORE_VM_FIXED0775_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0775E450g,
	"SUBCORE_VM_FIXED0800_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0800E450g,
	"SUBCORE_VM_FIXED0825_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0825E450g,
	"SUBCORE_VM_FIXED0850_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0850E450g,
	"SUBCORE_VM_FIXED0875_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0875E450g,
	"SUBCORE_VM_FIXED0900_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0900E450g,
	"SUBCORE_VM_FIXED0925_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0925E450g,
	"SUBCORE_VM_FIXED0950_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0950E450g,
	"SUBCORE_VM_FIXED0975_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0975E450g,
	"SUBCORE_VM_FIXED1000_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1000E450g,
	"SUBCORE_VM_FIXED1025_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1025E450g,
	"SUBCORE_VM_FIXED1050_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1050E450g,
	"SUBCORE_VM_FIXED1075_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1075E450g,
	"SUBCORE_VM_FIXED1100_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1100E450g,
	"SUBCORE_VM_FIXED1125_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1125E450g,
	"SUBCORE_VM_FIXED1150_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1150E450g,
	"SUBCORE_VM_FIXED1175_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1175E450g,
	"SUBCORE_VM_FIXED1200_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1200E450g,
	"SUBCORE_VM_FIXED1225_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1225E450g,
	"SUBCORE_VM_FIXED1250_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1250E450g,
	"SUBCORE_VM_FIXED1275_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1275E450g,
	"SUBCORE_VM_FIXED1300_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1300E450g,
	"SUBCORE_VM_FIXED1325_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1325E450g,
	"SUBCORE_VM_FIXED1350_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1350E450g,
	"SUBCORE_VM_FIXED1375_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1375E450g,
	"SUBCORE_VM_FIXED1400_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1400E450g,
	"SUBCORE_VM_FIXED1425_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1425E450g,
	"SUBCORE_VM_FIXED1450_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1450E450g,
	"SUBCORE_VM_FIXED1475_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1475E450g,
	"SUBCORE_VM_FIXED1500_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1500E450g,
	"SUBCORE_VM_FIXED1525_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1525E450g,
	"SUBCORE_VM_FIXED1550_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1550E450g,
	"SUBCORE_VM_FIXED1575_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1575E450g,
	"SUBCORE_VM_FIXED1600_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1600E450g,
	"SUBCORE_VM_FIXED1625_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1625E450g,
	"SUBCORE_VM_FIXED1650_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1650E450g,
	"SUBCORE_VM_FIXED1700_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1700E450g,
	"SUBCORE_VM_FIXED1725_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1725E450g,
	"SUBCORE_VM_FIXED1750_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1750E450g,
	"SUBCORE_VM_FIXED1800_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1800E450g,
	"SUBCORE_VM_FIXED1850_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1850E450g,
	"SUBCORE_VM_FIXED1875_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1875E450g,
	"SUBCORE_VM_FIXED1900_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1900E450g,
	"SUBCORE_VM_FIXED1925_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1925E450g,
	"SUBCORE_VM_FIXED1950_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1950E450g,
	"SUBCORE_VM_FIXED2000_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2000E450g,
	"SUBCORE_VM_FIXED2025_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2025E450g,
	"SUBCORE_VM_FIXED2050_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2050E450g,
	"SUBCORE_VM_FIXED2100_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2100E450g,
	"SUBCORE_VM_FIXED2125_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2125E450g,
	"SUBCORE_VM_FIXED2150_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2150E450g,
	"SUBCORE_VM_FIXED2175_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2175E450g,
	"SUBCORE_VM_FIXED2200_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2200E450g,
	"SUBCORE_VM_FIXED2250_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2250E450g,
	"SUBCORE_VM_FIXED2275_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2275E450g,
	"SUBCORE_VM_FIXED2300_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2300E450g,
	"SUBCORE_VM_FIXED2325_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2325E450g,
	"SUBCORE_VM_FIXED2350_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2350E450g,
	"SUBCORE_VM_FIXED2375_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2375E450g,
	"SUBCORE_VM_FIXED2400_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2400E450g,
	"SUBCORE_VM_FIXED2450_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2450E450g,
	"SUBCORE_VM_FIXED2475_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2475E450g,
	"SUBCORE_VM_FIXED2500_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2500E450g,
	"SUBCORE_VM_FIXED2550_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2550E450g,
	"SUBCORE_VM_FIXED2600_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2600E450g,
	"SUBCORE_VM_FIXED2625_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2625E450g,
	"SUBCORE_VM_FIXED2650_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2650E450g,
	"SUBCORE_VM_FIXED2700_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2700E450g,
	"SUBCORE_VM_FIXED2750_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2750E450g,
	"SUBCORE_VM_FIXED2775_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2775E450g,
	"SUBCORE_VM_FIXED2800_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2800E450g,
	"SUBCORE_VM_FIXED2850_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2850E450g,
	"SUBCORE_VM_FIXED2875_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2875E450g,
	"SUBCORE_VM_FIXED2900_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2900E450g,
	"SUBCORE_VM_FIXED2925_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2925E450g,
	"SUBCORE_VM_FIXED2950_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2950E450g,
	"SUBCORE_VM_FIXED2975_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2975E450g,
	"SUBCORE_VM_FIXED3000_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3000E450g,
	"SUBCORE_VM_FIXED3025_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3025E450g,
	"SUBCORE_VM_FIXED3050_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3050E450g,
	"SUBCORE_VM_FIXED3075_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3075E450g,
	"SUBCORE_VM_FIXED3100_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3100E450g,
	"SUBCORE_VM_FIXED3125_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3125E450g,
	"SUBCORE_VM_FIXED3150_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3150E450g,
	"SUBCORE_VM_FIXED3200_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3200E450g,
	"SUBCORE_VM_FIXED3225_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3225E450g,
	"SUBCORE_VM_FIXED3250_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3250E450g,
	"SUBCORE_VM_FIXED3300_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3300E450g,
	"SUBCORE_VM_FIXED3325_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3325E450g,
	"SUBCORE_VM_FIXED3375_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3375E450g,
	"SUBCORE_VM_FIXED3400_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3400E450g,
	"SUBCORE_VM_FIXED3450_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3450E450g,
	"SUBCORE_VM_FIXED3500_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3500E450g,
	"SUBCORE_VM_FIXED3525_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3525E450g,
	"SUBCORE_VM_FIXED3575_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3575E450g,
	"SUBCORE_VM_FIXED3600_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3600E450g,
	"SUBCORE_VM_FIXED3625_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3625E450g,
	"SUBCORE_VM_FIXED3675_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3675E450g,
	"SUBCORE_VM_FIXED3700_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3700E450g,
	"SUBCORE_VM_FIXED3750_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3750E450g,
	"SUBCORE_VM_FIXED3800_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3800E450g,
	"SUBCORE_VM_FIXED3825_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3825E450g,
	"SUBCORE_VM_FIXED3850_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3850E450g,
	"SUBCORE_VM_FIXED3875_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3875E450g,
	"SUBCORE_VM_FIXED3900_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3900E450g,
	"SUBCORE_VM_FIXED3975_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3975E450g,
	"SUBCORE_VM_FIXED4000_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4000E450g,
	"SUBCORE_VM_FIXED4025_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4025E450g,
	"SUBCORE_VM_FIXED4050_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4050E450g,
	"SUBCORE_VM_FIXED4100_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4100E450g,
	"SUBCORE_VM_FIXED4125_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4125E450g,
	"SUBCORE_VM_FIXED4200_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4200E450g,
	"SUBCORE_VM_FIXED4225_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4225E450g,
	"SUBCORE_VM_FIXED4250_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4250E450g,
	"SUBCORE_VM_FIXED4275_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4275E450g,
	"SUBCORE_VM_FIXED4300_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4300E450g,
	"SUBCORE_VM_FIXED4350_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4350E450g,
	"SUBCORE_VM_FIXED4375_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4375E450g,
	"SUBCORE_VM_FIXED4400_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4400E450g,
	"SUBCORE_VM_FIXED4425_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4425E450g,
	"SUBCORE_VM_FIXED4500_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4500E450g,
	"SUBCORE_VM_FIXED4550_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4550E450g,
	"SUBCORE_VM_FIXED4575_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4575E450g,
	"SUBCORE_VM_FIXED4600_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4600E450g,
	"SUBCORE_VM_FIXED4625_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4625E450g,
	"SUBCORE_VM_FIXED4650_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4650E450g,
	"SUBCORE_VM_FIXED4675_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4675E450g,
	"SUBCORE_VM_FIXED4700_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4700E450g,
	"SUBCORE_VM_FIXED4725_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4725E450g,
	"SUBCORE_VM_FIXED4750_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4750E450g,
	"SUBCORE_VM_FIXED4800_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4800E450g,
	"SUBCORE_VM_FIXED4875_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4875E450g,
	"SUBCORE_VM_FIXED4900_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4900E450g,
	"SUBCORE_VM_FIXED4950_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4950E450g,
	"SUBCORE_VM_FIXED5000_E4_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed5000E450g,
	"DYNAMIC_E5_50G":                       UpdateVnicShapeDetailsVnicShapeDynamicE550g,
	"DYNAMIC_E5_100G":                      UpdateVnicShapeDetailsVnicShapeDynamicE5100g,
	"EntireHost_E5_50G":                    UpdateVnicShapeDetailsVnicShapeEntirehostE550g,
	"Fixed0040_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0040E550g,
	"FIXED0100_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0100E550g,
	"FIXED0200_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0200E550g,
	"FIXED0300_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0300E550g,
	"FIXED0400_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0400E550g,
	"FIXED0500_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0500E550g,
	"FIXED0600_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0600E550g,
	"FIXED0700_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0700E550g,
	"FIXED0800_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0800E550g,
	"FIXED0900_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0900E550g,
	"FIXED1000_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1000E550g,
	"FIXED1100_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1100E550g,
	"FIXED1200_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1200E550g,
	"FIXED1300_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1300E550g,
	"FIXED1400_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1400E550g,
	"FIXED1500_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1500E550g,
	"FIXED1600_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1600E550g,
	"FIXED1700_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1700E550g,
	"FIXED1800_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1800E550g,
	"FIXED1900_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1900E550g,
	"FIXED2000_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2000E550g,
	"FIXED2100_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2100E550g,
	"FIXED2200_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2200E550g,
	"FIXED2300_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2300E550g,
	"FIXED2400_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2400E550g,
	"FIXED2500_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2500E550g,
	"FIXED2600_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2600E550g,
	"FIXED2700_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2700E550g,
	"FIXED2800_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2800E550g,
	"FIXED2900_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2900E550g,
	"FIXED3000_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3000E550g,
	"FIXED3100_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3100E550g,
	"FIXED3200_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3200E550g,
	"FIXED3300_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3300E550g,
	"FIXED3400_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3400E550g,
	"FIXED3500_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3500E550g,
	"FIXED3600_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3600E550g,
	"FIXED3700_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3700E550g,
	"FIXED3800_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3800E550g,
	"FIXED3900_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3900E550g,
	"FIXED4000_E5_50G":                     UpdateVnicShapeDetailsVnicShapeFixed4000E550g,
	"SUBCORE_VM_FIXED0025_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0025E550g,
	"SUBCORE_VM_FIXED0050_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0050E550g,
	"SUBCORE_VM_FIXED0075_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0075E550g,
	"SUBCORE_VM_FIXED0100_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0100E550g,
	"SUBCORE_VM_FIXED0125_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0125E550g,
	"SUBCORE_VM_FIXED0150_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0150E550g,
	"SUBCORE_VM_FIXED0175_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0175E550g,
	"SUBCORE_VM_FIXED0200_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0200E550g,
	"SUBCORE_VM_FIXED0225_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0225E550g,
	"SUBCORE_VM_FIXED0250_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0250E550g,
	"SUBCORE_VM_FIXED0275_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0275E550g,
	"SUBCORE_VM_FIXED0300_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0300E550g,
	"SUBCORE_VM_FIXED0325_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0325E550g,
	"SUBCORE_VM_FIXED0350_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0350E550g,
	"SUBCORE_VM_FIXED0375_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0375E550g,
	"SUBCORE_VM_FIXED0400_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0400E550g,
	"SUBCORE_VM_FIXED0425_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0425E550g,
	"SUBCORE_VM_FIXED0450_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0450E550g,
	"SUBCORE_VM_FIXED0475_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0475E550g,
	"SUBCORE_VM_FIXED0500_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0500E550g,
	"SUBCORE_VM_FIXED0525_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0525E550g,
	"SUBCORE_VM_FIXED0550_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0550E550g,
	"SUBCORE_VM_FIXED0575_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0575E550g,
	"SUBCORE_VM_FIXED0600_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0600E550g,
	"SUBCORE_VM_FIXED0625_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0625E550g,
	"SUBCORE_VM_FIXED0650_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0650E550g,
	"SUBCORE_VM_FIXED0675_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0675E550g,
	"SUBCORE_VM_FIXED0700_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0700E550g,
	"SUBCORE_VM_FIXED0725_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0725E550g,
	"SUBCORE_VM_FIXED0750_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0750E550g,
	"SUBCORE_VM_FIXED0775_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0775E550g,
	"SUBCORE_VM_FIXED0800_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0800E550g,
	"SUBCORE_VM_FIXED0825_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0825E550g,
	"SUBCORE_VM_FIXED0850_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0850E550g,
	"SUBCORE_VM_FIXED0875_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0875E550g,
	"SUBCORE_VM_FIXED0900_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0900E550g,
	"SUBCORE_VM_FIXED0925_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0925E550g,
	"SUBCORE_VM_FIXED0950_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0950E550g,
	"SUBCORE_VM_FIXED0975_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0975E550g,
	"SUBCORE_VM_FIXED1000_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1000E550g,
	"SUBCORE_VM_FIXED1025_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1025E550g,
	"SUBCORE_VM_FIXED1050_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1050E550g,
	"SUBCORE_VM_FIXED1075_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1075E550g,
	"SUBCORE_VM_FIXED1100_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1100E550g,
	"SUBCORE_VM_FIXED1125_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1125E550g,
	"SUBCORE_VM_FIXED1150_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1150E550g,
	"SUBCORE_VM_FIXED1175_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1175E550g,
	"SUBCORE_VM_FIXED1200_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1200E550g,
	"SUBCORE_VM_FIXED1225_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1225E550g,
	"SUBCORE_VM_FIXED1250_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1250E550g,
	"SUBCORE_VM_FIXED1275_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1275E550g,
	"SUBCORE_VM_FIXED1300_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1300E550g,
	"SUBCORE_VM_FIXED1325_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1325E550g,
	"SUBCORE_VM_FIXED1350_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1350E550g,
	"SUBCORE_VM_FIXED1375_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1375E550g,
	"SUBCORE_VM_FIXED1400_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1400E550g,
	"SUBCORE_VM_FIXED1425_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1425E550g,
	"SUBCORE_VM_FIXED1450_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1450E550g,
	"SUBCORE_VM_FIXED1475_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1475E550g,
	"SUBCORE_VM_FIXED1500_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1500E550g,
	"SUBCORE_VM_FIXED1525_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1525E550g,
	"SUBCORE_VM_FIXED1550_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1550E550g,
	"SUBCORE_VM_FIXED1575_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1575E550g,
	"SUBCORE_VM_FIXED1600_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1600E550g,
	"SUBCORE_VM_FIXED1625_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1625E550g,
	"SUBCORE_VM_FIXED1650_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1650E550g,
	"SUBCORE_VM_FIXED1675_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1675E550g,
	"SUBCORE_VM_FIXED1700_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1700E550g,
	"SUBCORE_VM_FIXED1725_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1725E550g,
	"SUBCORE_VM_FIXED1750_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1750E550g,
	"SUBCORE_VM_FIXED1775_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1775E550g,
	"SUBCORE_VM_FIXED1800_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1800E550g,
	"SUBCORE_VM_FIXED1825_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1825E550g,
	"SUBCORE_VM_FIXED1850_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1850E550g,
	"SUBCORE_VM_FIXED1875_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1875E550g,
	"SUBCORE_VM_FIXED1900_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1900E550g,
	"SUBCORE_VM_FIXED1925_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1925E550g,
	"SUBCORE_VM_FIXED1950_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1950E550g,
	"SUBCORE_VM_FIXED1975_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1975E550g,
	"SUBCORE_VM_FIXED2000_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2000E550g,
	"SUBCORE_VM_FIXED2025_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2025E550g,
	"SUBCORE_VM_FIXED2050_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2050E550g,
	"SUBCORE_VM_FIXED2075_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2075E550g,
	"SUBCORE_VM_FIXED2100_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2100E550g,
	"SUBCORE_VM_FIXED2125_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2125E550g,
	"SUBCORE_VM_FIXED2150_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2150E550g,
	"SUBCORE_VM_FIXED2175_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2175E550g,
	"SUBCORE_VM_FIXED2200_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2200E550g,
	"SUBCORE_VM_FIXED2225_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2225E550g,
	"SUBCORE_VM_FIXED2250_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2250E550g,
	"SUBCORE_VM_FIXED2275_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2275E550g,
	"SUBCORE_VM_FIXED2300_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2300E550g,
	"SUBCORE_VM_FIXED2325_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2325E550g,
	"SUBCORE_VM_FIXED2350_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2350E550g,
	"SUBCORE_VM_FIXED2375_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2375E550g,
	"SUBCORE_VM_FIXED2400_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2400E550g,
	"SUBCORE_VM_FIXED2425_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2425E550g,
	"SUBCORE_VM_FIXED2450_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2450E550g,
	"SUBCORE_VM_FIXED2475_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2475E550g,
	"SUBCORE_VM_FIXED2500_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2500E550g,
	"SUBCORE_VM_FIXED2525_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2525E550g,
	"SUBCORE_VM_FIXED2550_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2550E550g,
	"SUBCORE_VM_FIXED2575_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2575E550g,
	"SUBCORE_VM_FIXED2600_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2600E550g,
	"SUBCORE_VM_FIXED2625_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2625E550g,
	"SUBCORE_VM_FIXED2650_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2650E550g,
	"SUBCORE_VM_FIXED2675_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2675E550g,
	"SUBCORE_VM_FIXED2700_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2700E550g,
	"SUBCORE_VM_FIXED2725_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2725E550g,
	"SUBCORE_VM_FIXED2750_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2750E550g,
	"SUBCORE_VM_FIXED2775_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2775E550g,
	"SUBCORE_VM_FIXED2800_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2800E550g,
	"SUBCORE_VM_FIXED2825_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2825E550g,
	"SUBCORE_VM_FIXED2850_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2850E550g,
	"SUBCORE_VM_FIXED2875_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2875E550g,
	"SUBCORE_VM_FIXED2900_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2900E550g,
	"SUBCORE_VM_FIXED2925_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2925E550g,
	"SUBCORE_VM_FIXED2950_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2950E550g,
	"SUBCORE_VM_FIXED2975_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2975E550g,
	"SUBCORE_VM_FIXED3000_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3000E550g,
	"SUBCORE_VM_FIXED3025_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3025E550g,
	"SUBCORE_VM_FIXED3050_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3050E550g,
	"SUBCORE_VM_FIXED3075_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3075E550g,
	"SUBCORE_VM_FIXED3100_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3100E550g,
	"SUBCORE_VM_FIXED3125_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3125E550g,
	"SUBCORE_VM_FIXED3150_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3150E550g,
	"SUBCORE_VM_FIXED3175_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3175E550g,
	"SUBCORE_VM_FIXED3200_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3200E550g,
	"SUBCORE_VM_FIXED3225_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3225E550g,
	"SUBCORE_VM_FIXED3250_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3250E550g,
	"SUBCORE_VM_FIXED3275_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3275E550g,
	"SUBCORE_VM_FIXED3300_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3300E550g,
	"SUBCORE_VM_FIXED3325_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3325E550g,
	"SUBCORE_VM_FIXED3350_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3350E550g,
	"SUBCORE_VM_FIXED3375_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3375E550g,
	"SUBCORE_VM_FIXED3400_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3400E550g,
	"SUBCORE_VM_FIXED3425_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3425E550g,
	"SUBCORE_VM_FIXED3450_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3450E550g,
	"SUBCORE_VM_FIXED3475_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3475E550g,
	"SUBCORE_VM_FIXED3500_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3500E550g,
	"SUBCORE_VM_FIXED3525_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3525E550g,
	"SUBCORE_VM_FIXED3550_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3550E550g,
	"SUBCORE_VM_FIXED3575_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3575E550g,
	"SUBCORE_VM_FIXED3600_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3600E550g,
	"SUBCORE_VM_FIXED3625_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3625E550g,
	"SUBCORE_VM_FIXED3650_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3650E550g,
	"SUBCORE_VM_FIXED3675_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3675E550g,
	"SUBCORE_VM_FIXED3700_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3700E550g,
	"SUBCORE_VM_FIXED3725_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3725E550g,
	"SUBCORE_VM_FIXED3750_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3750E550g,
	"SUBCORE_VM_FIXED3775_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3775E550g,
	"SUBCORE_VM_FIXED3800_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3800E550g,
	"SUBCORE_VM_FIXED3825_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3825E550g,
	"SUBCORE_VM_FIXED3850_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3850E550g,
	"SUBCORE_VM_FIXED3875_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3875E550g,
	"SUBCORE_VM_FIXED3900_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3900E550g,
	"SUBCORE_VM_FIXED3925_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3925E550g,
	"SUBCORE_VM_FIXED3950_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3950E550g,
	"SUBCORE_VM_FIXED3975_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3975E550g,
	"SUBCORE_VM_FIXED4000_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4000E550g,
	"SUBCORE_VM_FIXED4025_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4025E550g,
	"SUBCORE_VM_FIXED4050_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4050E550g,
	"SUBCORE_VM_FIXED4075_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4075E550g,
	"SUBCORE_VM_FIXED4100_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4100E550g,
	"SUBCORE_VM_FIXED4125_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4125E550g,
	"SUBCORE_VM_FIXED4150_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4150E550g,
	"SUBCORE_VM_FIXED4175_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4175E550g,
	"SUBCORE_VM_FIXED4200_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4200E550g,
	"SUBCORE_VM_FIXED4225_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4225E550g,
	"SUBCORE_VM_FIXED4250_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4250E550g,
	"SUBCORE_VM_FIXED4275_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4275E550g,
	"SUBCORE_VM_FIXED4300_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4300E550g,
	"SUBCORE_VM_FIXED4325_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4325E550g,
	"SUBCORE_VM_FIXED4350_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4350E550g,
	"SUBCORE_VM_FIXED4375_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4375E550g,
	"SUBCORE_VM_FIXED4400_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4400E550g,
	"SUBCORE_VM_FIXED4425_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4425E550g,
	"SUBCORE_VM_FIXED4450_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4450E550g,
	"SUBCORE_VM_FIXED4475_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4475E550g,
	"SUBCORE_VM_FIXED4500_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4500E550g,
	"SUBCORE_VM_FIXED4525_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4525E550g,
	"SUBCORE_VM_FIXED4550_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4550E550g,
	"SUBCORE_VM_FIXED4575_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4575E550g,
	"SUBCORE_VM_FIXED4600_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4600E550g,
	"SUBCORE_VM_FIXED4625_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4625E550g,
	"SUBCORE_VM_FIXED4650_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4650E550g,
	"SUBCORE_VM_FIXED4675_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4675E550g,
	"SUBCORE_VM_FIXED4700_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4700E550g,
	"SUBCORE_VM_FIXED4725_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4725E550g,
	"SUBCORE_VM_FIXED4750_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4750E550g,
	"SUBCORE_VM_FIXED4775_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4775E550g,
	"SUBCORE_VM_FIXED4800_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4800E550g,
	"SUBCORE_VM_FIXED4825_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4825E550g,
	"SUBCORE_VM_FIXED4850_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4850E550g,
	"SUBCORE_VM_FIXED4875_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4875E550g,
	"SUBCORE_VM_FIXED4900_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4900E550g,
	"SUBCORE_VM_FIXED4925_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4925E550g,
	"SUBCORE_VM_FIXED4950_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4950E550g,
	"SUBCORE_VM_FIXED4975_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4975E550g,
	"SUBCORE_VM_FIXED5000_E5_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed5000E550g,
	"SUBCORE_VM_FIXED0020_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0020A150g,
	"SUBCORE_VM_FIXED0040_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0040A150g,
	"SUBCORE_VM_FIXED0060_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0060A150g,
	"SUBCORE_VM_FIXED0080_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0080A150g,
	"SUBCORE_VM_FIXED0100_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0100A150g,
	"SUBCORE_VM_FIXED0120_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0120A150g,
	"SUBCORE_VM_FIXED0140_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0140A150g,
	"SUBCORE_VM_FIXED0160_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0160A150g,
	"SUBCORE_VM_FIXED0180_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0180A150g,
	"SUBCORE_VM_FIXED0200_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0200A150g,
	"SUBCORE_VM_FIXED0220_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0220A150g,
	"SUBCORE_VM_FIXED0240_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0240A150g,
	"SUBCORE_VM_FIXED0260_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0260A150g,
	"SUBCORE_VM_FIXED0280_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0280A150g,
	"SUBCORE_VM_FIXED0300_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0300A150g,
	"SUBCORE_VM_FIXED0320_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0320A150g,
	"SUBCORE_VM_FIXED0340_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0340A150g,
	"SUBCORE_VM_FIXED0360_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0360A150g,
	"SUBCORE_VM_FIXED0380_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0380A150g,
	"SUBCORE_VM_FIXED0400_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0400A150g,
	"SUBCORE_VM_FIXED0420_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0420A150g,
	"SUBCORE_VM_FIXED0440_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0440A150g,
	"SUBCORE_VM_FIXED0460_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0460A150g,
	"SUBCORE_VM_FIXED0480_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0480A150g,
	"SUBCORE_VM_FIXED0500_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0500A150g,
	"SUBCORE_VM_FIXED0520_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0520A150g,
	"SUBCORE_VM_FIXED0540_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0540A150g,
	"SUBCORE_VM_FIXED0560_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0560A150g,
	"SUBCORE_VM_FIXED0580_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0580A150g,
	"SUBCORE_VM_FIXED0600_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0600A150g,
	"SUBCORE_VM_FIXED0620_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0620A150g,
	"SUBCORE_VM_FIXED0640_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0640A150g,
	"SUBCORE_VM_FIXED0660_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0660A150g,
	"SUBCORE_VM_FIXED0680_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0680A150g,
	"SUBCORE_VM_FIXED0700_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0700A150g,
	"SUBCORE_VM_FIXED0720_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0720A150g,
	"SUBCORE_VM_FIXED0740_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0740A150g,
	"SUBCORE_VM_FIXED0760_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0760A150g,
	"SUBCORE_VM_FIXED0780_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0780A150g,
	"SUBCORE_VM_FIXED0800_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0800A150g,
	"SUBCORE_VM_FIXED0820_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0820A150g,
	"SUBCORE_VM_FIXED0840_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0840A150g,
	"SUBCORE_VM_FIXED0860_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0860A150g,
	"SUBCORE_VM_FIXED0880_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0880A150g,
	"SUBCORE_VM_FIXED0900_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0900A150g,
	"SUBCORE_VM_FIXED0920_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0920A150g,
	"SUBCORE_VM_FIXED0940_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0940A150g,
	"SUBCORE_VM_FIXED0960_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0960A150g,
	"SUBCORE_VM_FIXED0980_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0980A150g,
	"SUBCORE_VM_FIXED1000_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1000A150g,
	"SUBCORE_VM_FIXED1020_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1020A150g,
	"SUBCORE_VM_FIXED1040_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1040A150g,
	"SUBCORE_VM_FIXED1060_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1060A150g,
	"SUBCORE_VM_FIXED1080_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1080A150g,
	"SUBCORE_VM_FIXED1100_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1100A150g,
	"SUBCORE_VM_FIXED1120_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1120A150g,
	"SUBCORE_VM_FIXED1140_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1140A150g,
	"SUBCORE_VM_FIXED1160_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1160A150g,
	"SUBCORE_VM_FIXED1180_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1180A150g,
	"SUBCORE_VM_FIXED1200_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1200A150g,
	"SUBCORE_VM_FIXED1220_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1220A150g,
	"SUBCORE_VM_FIXED1240_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1240A150g,
	"SUBCORE_VM_FIXED1260_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1260A150g,
	"SUBCORE_VM_FIXED1280_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1280A150g,
	"SUBCORE_VM_FIXED1300_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1300A150g,
	"SUBCORE_VM_FIXED1320_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1320A150g,
	"SUBCORE_VM_FIXED1340_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1340A150g,
	"SUBCORE_VM_FIXED1360_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1360A150g,
	"SUBCORE_VM_FIXED1380_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1380A150g,
	"SUBCORE_VM_FIXED1400_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1400A150g,
	"SUBCORE_VM_FIXED1420_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1420A150g,
	"SUBCORE_VM_FIXED1440_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1440A150g,
	"SUBCORE_VM_FIXED1460_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1460A150g,
	"SUBCORE_VM_FIXED1480_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1480A150g,
	"SUBCORE_VM_FIXED1500_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1500A150g,
	"SUBCORE_VM_FIXED1520_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1520A150g,
	"SUBCORE_VM_FIXED1540_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1540A150g,
	"SUBCORE_VM_FIXED1560_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1560A150g,
	"SUBCORE_VM_FIXED1580_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1580A150g,
	"SUBCORE_VM_FIXED1600_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1600A150g,
	"SUBCORE_VM_FIXED1620_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1620A150g,
	"SUBCORE_VM_FIXED1640_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1640A150g,
	"SUBCORE_VM_FIXED1660_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1660A150g,
	"SUBCORE_VM_FIXED1680_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1680A150g,
	"SUBCORE_VM_FIXED1700_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1700A150g,
	"SUBCORE_VM_FIXED1720_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1720A150g,
	"SUBCORE_VM_FIXED1740_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1740A150g,
	"SUBCORE_VM_FIXED1760_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1760A150g,
	"SUBCORE_VM_FIXED1780_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1780A150g,
	"SUBCORE_VM_FIXED1800_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1800A150g,
	"SUBCORE_VM_FIXED1820_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1820A150g,
	"SUBCORE_VM_FIXED1840_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1840A150g,
	"SUBCORE_VM_FIXED1860_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1860A150g,
	"SUBCORE_VM_FIXED1880_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1880A150g,
	"SUBCORE_VM_FIXED1900_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1900A150g,
	"SUBCORE_VM_FIXED1920_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1920A150g,
	"SUBCORE_VM_FIXED1940_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1940A150g,
	"SUBCORE_VM_FIXED1960_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1960A150g,
	"SUBCORE_VM_FIXED1980_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1980A150g,
	"SUBCORE_VM_FIXED2000_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2000A150g,
	"SUBCORE_VM_FIXED2020_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2020A150g,
	"SUBCORE_VM_FIXED2040_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2040A150g,
	"SUBCORE_VM_FIXED2060_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2060A150g,
	"SUBCORE_VM_FIXED2080_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2080A150g,
	"SUBCORE_VM_FIXED2100_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2100A150g,
	"SUBCORE_VM_FIXED2120_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2120A150g,
	"SUBCORE_VM_FIXED2140_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2140A150g,
	"SUBCORE_VM_FIXED2160_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2160A150g,
	"SUBCORE_VM_FIXED2180_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2180A150g,
	"SUBCORE_VM_FIXED2200_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2200A150g,
	"SUBCORE_VM_FIXED2220_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2220A150g,
	"SUBCORE_VM_FIXED2240_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2240A150g,
	"SUBCORE_VM_FIXED2260_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2260A150g,
	"SUBCORE_VM_FIXED2280_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2280A150g,
	"SUBCORE_VM_FIXED2300_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2300A150g,
	"SUBCORE_VM_FIXED2320_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2320A150g,
	"SUBCORE_VM_FIXED2340_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2340A150g,
	"SUBCORE_VM_FIXED2360_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2360A150g,
	"SUBCORE_VM_FIXED2380_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2380A150g,
	"SUBCORE_VM_FIXED2400_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2400A150g,
	"SUBCORE_VM_FIXED2420_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2420A150g,
	"SUBCORE_VM_FIXED2440_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2440A150g,
	"SUBCORE_VM_FIXED2460_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2460A150g,
	"SUBCORE_VM_FIXED2480_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2480A150g,
	"SUBCORE_VM_FIXED2500_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2500A150g,
	"SUBCORE_VM_FIXED2520_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2520A150g,
	"SUBCORE_VM_FIXED2540_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2540A150g,
	"SUBCORE_VM_FIXED2560_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2560A150g,
	"SUBCORE_VM_FIXED2580_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2580A150g,
	"SUBCORE_VM_FIXED2600_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2600A150g,
	"SUBCORE_VM_FIXED2620_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2620A150g,
	"SUBCORE_VM_FIXED2640_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2640A150g,
	"SUBCORE_VM_FIXED2660_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2660A150g,
	"SUBCORE_VM_FIXED2680_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2680A150g,
	"SUBCORE_VM_FIXED2700_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2700A150g,
	"SUBCORE_VM_FIXED2720_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2720A150g,
	"SUBCORE_VM_FIXED2740_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2740A150g,
	"SUBCORE_VM_FIXED2760_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2760A150g,
	"SUBCORE_VM_FIXED2780_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2780A150g,
	"SUBCORE_VM_FIXED2800_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2800A150g,
	"SUBCORE_VM_FIXED2820_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2820A150g,
	"SUBCORE_VM_FIXED2840_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2840A150g,
	"SUBCORE_VM_FIXED2860_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2860A150g,
	"SUBCORE_VM_FIXED2880_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2880A150g,
	"SUBCORE_VM_FIXED2900_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2900A150g,
	"SUBCORE_VM_FIXED2920_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2920A150g,
	"SUBCORE_VM_FIXED2940_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2940A150g,
	"SUBCORE_VM_FIXED2960_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2960A150g,
	"SUBCORE_VM_FIXED2980_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2980A150g,
	"SUBCORE_VM_FIXED3000_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3000A150g,
	"SUBCORE_VM_FIXED3020_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3020A150g,
	"SUBCORE_VM_FIXED3040_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3040A150g,
	"SUBCORE_VM_FIXED3060_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3060A150g,
	"SUBCORE_VM_FIXED3080_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3080A150g,
	"SUBCORE_VM_FIXED3100_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3100A150g,
	"SUBCORE_VM_FIXED3120_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3120A150g,
	"SUBCORE_VM_FIXED3140_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3140A150g,
	"SUBCORE_VM_FIXED3160_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3160A150g,
	"SUBCORE_VM_FIXED3180_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3180A150g,
	"SUBCORE_VM_FIXED3200_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3200A150g,
	"SUBCORE_VM_FIXED3220_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3220A150g,
	"SUBCORE_VM_FIXED3240_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3240A150g,
	"SUBCORE_VM_FIXED3260_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3260A150g,
	"SUBCORE_VM_FIXED3280_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3280A150g,
	"SUBCORE_VM_FIXED3300_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3300A150g,
	"SUBCORE_VM_FIXED3320_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3320A150g,
	"SUBCORE_VM_FIXED3340_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3340A150g,
	"SUBCORE_VM_FIXED3360_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3360A150g,
	"SUBCORE_VM_FIXED3380_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3380A150g,
	"SUBCORE_VM_FIXED3400_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3400A150g,
	"SUBCORE_VM_FIXED3420_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3420A150g,
	"SUBCORE_VM_FIXED3440_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3440A150g,
	"SUBCORE_VM_FIXED3460_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3460A150g,
	"SUBCORE_VM_FIXED3480_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3480A150g,
	"SUBCORE_VM_FIXED3500_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3500A150g,
	"SUBCORE_VM_FIXED3520_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3520A150g,
	"SUBCORE_VM_FIXED3540_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3540A150g,
	"SUBCORE_VM_FIXED3560_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3560A150g,
	"SUBCORE_VM_FIXED3580_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3580A150g,
	"SUBCORE_VM_FIXED3600_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3600A150g,
	"SUBCORE_VM_FIXED3620_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3620A150g,
	"SUBCORE_VM_FIXED3640_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3640A150g,
	"SUBCORE_VM_FIXED3660_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3660A150g,
	"SUBCORE_VM_FIXED3680_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3680A150g,
	"SUBCORE_VM_FIXED3700_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3700A150g,
	"SUBCORE_VM_FIXED3720_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3720A150g,
	"SUBCORE_VM_FIXED3740_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3740A150g,
	"SUBCORE_VM_FIXED3760_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3760A150g,
	"SUBCORE_VM_FIXED3780_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3780A150g,
	"SUBCORE_VM_FIXED3800_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3800A150g,
	"SUBCORE_VM_FIXED3820_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3820A150g,
	"SUBCORE_VM_FIXED3840_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3840A150g,
	"SUBCORE_VM_FIXED3860_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3860A150g,
	"SUBCORE_VM_FIXED3880_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3880A150g,
	"SUBCORE_VM_FIXED3900_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3900A150g,
	"SUBCORE_VM_FIXED3920_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3920A150g,
	"SUBCORE_VM_FIXED3940_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3940A150g,
	"SUBCORE_VM_FIXED3960_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3960A150g,
	"SUBCORE_VM_FIXED3980_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3980A150g,
	"SUBCORE_VM_FIXED4000_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4000A150g,
	"SUBCORE_VM_FIXED4020_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4020A150g,
	"SUBCORE_VM_FIXED4040_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4040A150g,
	"SUBCORE_VM_FIXED4060_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4060A150g,
	"SUBCORE_VM_FIXED4080_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4080A150g,
	"SUBCORE_VM_FIXED4100_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4100A150g,
	"SUBCORE_VM_FIXED4120_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4120A150g,
	"SUBCORE_VM_FIXED4140_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4140A150g,
	"SUBCORE_VM_FIXED4160_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4160A150g,
	"SUBCORE_VM_FIXED4180_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4180A150g,
	"SUBCORE_VM_FIXED4200_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4200A150g,
	"SUBCORE_VM_FIXED4220_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4220A150g,
	"SUBCORE_VM_FIXED4240_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4240A150g,
	"SUBCORE_VM_FIXED4260_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4260A150g,
	"SUBCORE_VM_FIXED4280_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4280A150g,
	"SUBCORE_VM_FIXED4300_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4300A150g,
	"SUBCORE_VM_FIXED4320_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4320A150g,
	"SUBCORE_VM_FIXED4340_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4340A150g,
	"SUBCORE_VM_FIXED4360_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4360A150g,
	"SUBCORE_VM_FIXED4380_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4380A150g,
	"SUBCORE_VM_FIXED4400_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4400A150g,
	"SUBCORE_VM_FIXED4420_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4420A150g,
	"SUBCORE_VM_FIXED4440_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4440A150g,
	"SUBCORE_VM_FIXED4460_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4460A150g,
	"SUBCORE_VM_FIXED4480_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4480A150g,
	"SUBCORE_VM_FIXED4500_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4500A150g,
	"SUBCORE_VM_FIXED4520_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4520A150g,
	"SUBCORE_VM_FIXED4540_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4540A150g,
	"SUBCORE_VM_FIXED4560_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4560A150g,
	"SUBCORE_VM_FIXED4580_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4580A150g,
	"SUBCORE_VM_FIXED4600_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4600A150g,
	"SUBCORE_VM_FIXED4620_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4620A150g,
	"SUBCORE_VM_FIXED4640_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4640A150g,
	"SUBCORE_VM_FIXED4660_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4660A150g,
	"SUBCORE_VM_FIXED4680_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4680A150g,
	"SUBCORE_VM_FIXED4700_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4700A150g,
	"SUBCORE_VM_FIXED4720_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4720A150g,
	"SUBCORE_VM_FIXED4740_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4740A150g,
	"SUBCORE_VM_FIXED4760_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4760A150g,
	"SUBCORE_VM_FIXED4780_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4780A150g,
	"SUBCORE_VM_FIXED4800_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4800A150g,
	"SUBCORE_VM_FIXED4820_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4820A150g,
	"SUBCORE_VM_FIXED4840_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4840A150g,
	"SUBCORE_VM_FIXED4860_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4860A150g,
	"SUBCORE_VM_FIXED4880_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4880A150g,
	"SUBCORE_VM_FIXED4900_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4900A150g,
	"SUBCORE_VM_FIXED4920_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4920A150g,
	"SUBCORE_VM_FIXED4940_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4940A150g,
	"SUBCORE_VM_FIXED4960_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4960A150g,
	"SUBCORE_VM_FIXED4980_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4980A150g,
	"SUBCORE_VM_FIXED5000_A1_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed5000A150g,
	"SUBCORE_VM_FIXED0090_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0090X950g,
	"SUBCORE_VM_FIXED0180_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0180X950g,
	"SUBCORE_VM_FIXED0270_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0270X950g,
	"SUBCORE_VM_FIXED0360_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0360X950g,
	"SUBCORE_VM_FIXED0450_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0450X950g,
	"SUBCORE_VM_FIXED0540_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0540X950g,
	"SUBCORE_VM_FIXED0630_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0630X950g,
	"SUBCORE_VM_FIXED0720_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0720X950g,
	"SUBCORE_VM_FIXED0810_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0810X950g,
	"SUBCORE_VM_FIXED0900_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0900X950g,
	"SUBCORE_VM_FIXED0990_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0990X950g,
	"SUBCORE_VM_FIXED1080_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1080X950g,
	"SUBCORE_VM_FIXED1170_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1170X950g,
	"SUBCORE_VM_FIXED1260_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1260X950g,
	"SUBCORE_VM_FIXED1350_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1350X950g,
	"SUBCORE_VM_FIXED1440_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1440X950g,
	"SUBCORE_VM_FIXED1530_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1530X950g,
	"SUBCORE_VM_FIXED1620_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1620X950g,
	"SUBCORE_VM_FIXED1710_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1710X950g,
	"SUBCORE_VM_FIXED1800_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1800X950g,
	"SUBCORE_VM_FIXED1890_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1890X950g,
	"SUBCORE_VM_FIXED1980_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1980X950g,
	"SUBCORE_VM_FIXED2070_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2070X950g,
	"SUBCORE_VM_FIXED2160_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2160X950g,
	"SUBCORE_VM_FIXED2250_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2250X950g,
	"SUBCORE_VM_FIXED2340_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2340X950g,
	"SUBCORE_VM_FIXED2430_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2430X950g,
	"SUBCORE_VM_FIXED2520_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2520X950g,
	"SUBCORE_VM_FIXED2610_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2610X950g,
	"SUBCORE_VM_FIXED2700_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2700X950g,
	"SUBCORE_VM_FIXED2790_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2790X950g,
	"SUBCORE_VM_FIXED2880_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2880X950g,
	"SUBCORE_VM_FIXED2970_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2970X950g,
	"SUBCORE_VM_FIXED3060_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3060X950g,
	"SUBCORE_VM_FIXED3150_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3150X950g,
	"SUBCORE_VM_FIXED3240_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3240X950g,
	"SUBCORE_VM_FIXED3330_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3330X950g,
	"SUBCORE_VM_FIXED3420_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3420X950g,
	"SUBCORE_VM_FIXED3510_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3510X950g,
	"SUBCORE_VM_FIXED3600_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3600X950g,
	"SUBCORE_VM_FIXED3690_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3690X950g,
	"SUBCORE_VM_FIXED3780_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3780X950g,
	"SUBCORE_VM_FIXED3870_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3870X950g,
	"SUBCORE_VM_FIXED3960_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3960X950g,
	"SUBCORE_VM_FIXED4050_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4050X950g,
	"SUBCORE_VM_FIXED4140_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4140X950g,
	"SUBCORE_VM_FIXED4230_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4230X950g,
	"SUBCORE_VM_FIXED4320_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4320X950g,
	"SUBCORE_VM_FIXED4410_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4410X950g,
	"SUBCORE_VM_FIXED4500_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4500X950g,
	"SUBCORE_VM_FIXED4590_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4590X950g,
	"SUBCORE_VM_FIXED4680_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4680X950g,
	"SUBCORE_VM_FIXED4770_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4770X950g,
	"SUBCORE_VM_FIXED4860_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4860X950g,
	"SUBCORE_VM_FIXED4950_X9_50G":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4950X950g,
	"DYNAMIC_A1_50G":                       UpdateVnicShapeDetailsVnicShapeDynamicA150g,
	"FIXED0040_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0040A150g,
	"FIXED0100_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0100A150g,
	"FIXED0200_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0200A150g,
	"FIXED0300_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0300A150g,
	"FIXED0400_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0400A150g,
	"FIXED0500_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0500A150g,
	"FIXED0600_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0600A150g,
	"FIXED0700_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0700A150g,
	"FIXED0800_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0800A150g,
	"FIXED0900_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0900A150g,
	"FIXED1000_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1000A150g,
	"FIXED1100_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1100A150g,
	"FIXED1200_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1200A150g,
	"FIXED1300_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1300A150g,
	"FIXED1400_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1400A150g,
	"FIXED1500_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1500A150g,
	"FIXED1600_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1600A150g,
	"FIXED1700_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1700A150g,
	"FIXED1800_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1800A150g,
	"FIXED1900_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1900A150g,
	"FIXED2000_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2000A150g,
	"FIXED2100_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2100A150g,
	"FIXED2200_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2200A150g,
	"FIXED2300_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2300A150g,
	"FIXED2400_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2400A150g,
	"FIXED2500_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2500A150g,
	"FIXED2600_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2600A150g,
	"FIXED2700_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2700A150g,
	"FIXED2800_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2800A150g,
	"FIXED2900_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2900A150g,
	"FIXED3000_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3000A150g,
	"FIXED3100_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3100A150g,
	"FIXED3200_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3200A150g,
	"FIXED3300_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3300A150g,
	"FIXED3400_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3400A150g,
	"FIXED3500_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3500A150g,
	"FIXED3600_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3600A150g,
	"FIXED3700_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3700A150g,
	"FIXED3800_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3800A150g,
	"FIXED3900_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3900A150g,
	"FIXED4000_A1_50G":                     UpdateVnicShapeDetailsVnicShapeFixed4000A150g,
	"FIXED5000_TELESIS_A1_50G":             UpdateVnicShapeDetailsVnicShapeFixed5000TelesisA150g,
	"ENTIREHOST_A1_50G":                    UpdateVnicShapeDetailsVnicShapeEntirehostA150g,
	"DYNAMIC_X7_25G":                       UpdateVnicShapeDetailsVnicShapeDynamicX725g,
	"FIXED0040_X7_25G":                     UpdateVnicShapeDetailsVnicShapeFixed0040X725g,
	"STANDARD_VM_FIXED0100_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0100X725g,
	"STANDARD_VM_FIXED0200_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0200X725g,
	"STANDARD_VM_FIXED0300_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0300X725g,
	"STANDARD_VM_FIXED0400_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0400X725g,
	"STANDARD_VM_FIXED0500_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0500X725g,
	"STANDARD_VM_FIXED0600_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0600X725g,
	"STANDARD_VM_FIXED0700_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0700X725g,
	"STANDARD_VM_FIXED0800_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0800X725g,
	"STANDARD_VM_FIXED0900_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0900X725g,
	"STANDARD_VM_FIXED1000_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1000X725g,
	"STANDARD_VM_FIXED1100_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1100X725g,
	"STANDARD_VM_FIXED1200_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1200X725g,
	"STANDARD_VM_FIXED1300_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1300X725g,
	"STANDARD_VM_FIXED1400_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1400X725g,
	"STANDARD_VM_FIXED1500_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1500X725g,
	"STANDARD_VM_FIXED1600_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1600X725g,
	"STANDARD_VM_FIXED1700_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1700X725g,
	"STANDARD_VM_FIXED1800_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1800X725g,
	"STANDARD_VM_FIXED1900_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1900X725g,
	"STANDARD_VM_FIXED2000_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2000X725g,
	"STANDARD_VM_FIXED2100_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2100X725g,
	"STANDARD_VM_FIXED2200_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2200X725g,
	"STANDARD_VM_FIXED2300_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2300X725g,
	"STANDARD_VM_FIXED2400_X7_25G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2400X725g,
	"SUBCORE_STANDARD_VM_FIXED0025_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0025X725g,
	"SUBCORE_STANDARD_VM_FIXED0050_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0050X725g,
	"SUBCORE_STANDARD_VM_FIXED0075_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0075X725g,
	"SUBCORE_STANDARD_VM_FIXED0100_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0100X725g,
	"SUBCORE_STANDARD_VM_FIXED0125_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0125X725g,
	"SUBCORE_STANDARD_VM_FIXED0150_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0150X725g,
	"SUBCORE_STANDARD_VM_FIXED0175_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0175X725g,
	"SUBCORE_STANDARD_VM_FIXED0200_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0200X725g,
	"SUBCORE_STANDARD_VM_FIXED0225_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0225X725g,
	"SUBCORE_STANDARD_VM_FIXED0250_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0250X725g,
	"SUBCORE_STANDARD_VM_FIXED0275_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0275X725g,
	"SUBCORE_STANDARD_VM_FIXED0300_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0300X725g,
	"SUBCORE_STANDARD_VM_FIXED0325_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0325X725g,
	"SUBCORE_STANDARD_VM_FIXED0350_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0350X725g,
	"SUBCORE_STANDARD_VM_FIXED0375_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0375X725g,
	"SUBCORE_STANDARD_VM_FIXED0400_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0400X725g,
	"SUBCORE_STANDARD_VM_FIXED0425_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0425X725g,
	"SUBCORE_STANDARD_VM_FIXED0450_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0450X725g,
	"SUBCORE_STANDARD_VM_FIXED0475_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0475X725g,
	"SUBCORE_STANDARD_VM_FIXED0500_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0500X725g,
	"SUBCORE_STANDARD_VM_FIXED0525_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0525X725g,
	"SUBCORE_STANDARD_VM_FIXED0550_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0550X725g,
	"SUBCORE_STANDARD_VM_FIXED0575_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0575X725g,
	"SUBCORE_STANDARD_VM_FIXED0600_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0600X725g,
	"SUBCORE_STANDARD_VM_FIXED0625_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0625X725g,
	"SUBCORE_STANDARD_VM_FIXED0650_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0650X725g,
	"SUBCORE_STANDARD_VM_FIXED0675_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0675X725g,
	"SUBCORE_STANDARD_VM_FIXED0700_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0700X725g,
	"SUBCORE_STANDARD_VM_FIXED0725_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0725X725g,
	"SUBCORE_STANDARD_VM_FIXED0750_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0750X725g,
	"SUBCORE_STANDARD_VM_FIXED0775_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0775X725g,
	"SUBCORE_STANDARD_VM_FIXED0800_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0800X725g,
	"SUBCORE_STANDARD_VM_FIXED0825_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0825X725g,
	"SUBCORE_STANDARD_VM_FIXED0850_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0850X725g,
	"SUBCORE_STANDARD_VM_FIXED0875_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0875X725g,
	"SUBCORE_STANDARD_VM_FIXED0900_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0900X725g,
	"SUBCORE_STANDARD_VM_FIXED0925_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0925X725g,
	"SUBCORE_STANDARD_VM_FIXED0950_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0950X725g,
	"SUBCORE_STANDARD_VM_FIXED0975_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0975X725g,
	"SUBCORE_STANDARD_VM_FIXED1000_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1000X725g,
	"SUBCORE_STANDARD_VM_FIXED1025_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1025X725g,
	"SUBCORE_STANDARD_VM_FIXED1050_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1050X725g,
	"SUBCORE_STANDARD_VM_FIXED1075_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1075X725g,
	"SUBCORE_STANDARD_VM_FIXED1100_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1100X725g,
	"SUBCORE_STANDARD_VM_FIXED1125_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1125X725g,
	"SUBCORE_STANDARD_VM_FIXED1150_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1150X725g,
	"SUBCORE_STANDARD_VM_FIXED1175_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1175X725g,
	"SUBCORE_STANDARD_VM_FIXED1200_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1200X725g,
	"SUBCORE_STANDARD_VM_FIXED1225_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1225X725g,
	"SUBCORE_STANDARD_VM_FIXED1250_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1250X725g,
	"SUBCORE_STANDARD_VM_FIXED1275_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1275X725g,
	"SUBCORE_STANDARD_VM_FIXED1300_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1300X725g,
	"SUBCORE_STANDARD_VM_FIXED1325_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1325X725g,
	"SUBCORE_STANDARD_VM_FIXED1350_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1350X725g,
	"SUBCORE_STANDARD_VM_FIXED1375_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1375X725g,
	"SUBCORE_STANDARD_VM_FIXED1400_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1400X725g,
	"SUBCORE_STANDARD_VM_FIXED1425_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1425X725g,
	"SUBCORE_STANDARD_VM_FIXED1450_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1450X725g,
	"SUBCORE_STANDARD_VM_FIXED1475_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1475X725g,
	"SUBCORE_STANDARD_VM_FIXED1500_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1500X725g,
	"SUBCORE_STANDARD_VM_FIXED1525_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1525X725g,
	"SUBCORE_STANDARD_VM_FIXED1550_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1550X725g,
	"SUBCORE_STANDARD_VM_FIXED1575_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1575X725g,
	"SUBCORE_STANDARD_VM_FIXED1600_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1600X725g,
	"SUBCORE_STANDARD_VM_FIXED1625_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1625X725g,
	"SUBCORE_STANDARD_VM_FIXED1650_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1650X725g,
	"SUBCORE_STANDARD_VM_FIXED1675_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1675X725g,
	"SUBCORE_STANDARD_VM_FIXED1700_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1700X725g,
	"SUBCORE_STANDARD_VM_FIXED1725_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1725X725g,
	"SUBCORE_STANDARD_VM_FIXED1750_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1750X725g,
	"SUBCORE_STANDARD_VM_FIXED1775_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1775X725g,
	"SUBCORE_STANDARD_VM_FIXED1800_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1800X725g,
	"SUBCORE_STANDARD_VM_FIXED1825_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1825X725g,
	"SUBCORE_STANDARD_VM_FIXED1850_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1850X725g,
	"SUBCORE_STANDARD_VM_FIXED1875_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1875X725g,
	"SUBCORE_STANDARD_VM_FIXED1900_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1900X725g,
	"SUBCORE_STANDARD_VM_FIXED1925_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1925X725g,
	"SUBCORE_STANDARD_VM_FIXED1950_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1950X725g,
	"SUBCORE_STANDARD_VM_FIXED1975_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1975X725g,
	"SUBCORE_STANDARD_VM_FIXED2000_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2000X725g,
	"SUBCORE_STANDARD_VM_FIXED2025_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2025X725g,
	"SUBCORE_STANDARD_VM_FIXED2050_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2050X725g,
	"SUBCORE_STANDARD_VM_FIXED2075_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2075X725g,
	"SUBCORE_STANDARD_VM_FIXED2100_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2100X725g,
	"SUBCORE_STANDARD_VM_FIXED2125_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2125X725g,
	"SUBCORE_STANDARD_VM_FIXED2150_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2150X725g,
	"SUBCORE_STANDARD_VM_FIXED2175_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2175X725g,
	"SUBCORE_STANDARD_VM_FIXED2200_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2200X725g,
	"SUBCORE_STANDARD_VM_FIXED2225_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2225X725g,
	"SUBCORE_STANDARD_VM_FIXED2250_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2250X725g,
	"SUBCORE_STANDARD_VM_FIXED2275_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2275X725g,
	"SUBCORE_STANDARD_VM_FIXED2300_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2300X725g,
	"SUBCORE_STANDARD_VM_FIXED2325_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2325X725g,
	"SUBCORE_STANDARD_VM_FIXED2350_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2350X725g,
	"SUBCORE_STANDARD_VM_FIXED2375_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2375X725g,
	"SUBCORE_STANDARD_VM_FIXED2400_X7_25G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2400X725g,
	"ENTIREHOST_X7_25G":                    UpdateVnicShapeDetailsVnicShapeEntirehostX725g,
	"DYNAMIC_X9_50G":                       UpdateVnicShapeDetailsVnicShapeDynamicX950g,
	"FIXED0040_X9_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0040X950g,
	"FIXED0400_X9_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0400X950g,
	"FIXED0800_X9_50G":                     UpdateVnicShapeDetailsVnicShapeFixed0800X950g,
	"FIXED1200_X9_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1200X950g,
	"FIXED1600_X9_50G":                     UpdateVnicShapeDetailsVnicShapeFixed1600X950g,
	"FIXED2000_X9_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2000X950g,
	"FIXED2400_X9_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2400X950g,
	"FIXED2800_X9_50G":                     UpdateVnicShapeDetailsVnicShapeFixed2800X950g,
	"FIXED3200_X9_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3200X950g,
	"FIXED3600_X9_50G":                     UpdateVnicShapeDetailsVnicShapeFixed3600X950g,
	"FIXED4000_X9_50G":                     UpdateVnicShapeDetailsVnicShapeFixed4000X950g,
	"STANDARD_VM_FIXED0100_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0100X950g,
	"STANDARD_VM_FIXED0200_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0200X950g,
	"STANDARD_VM_FIXED0300_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0300X950g,
	"STANDARD_VM_FIXED0400_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0400X950g,
	"STANDARD_VM_FIXED0500_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0500X950g,
	"STANDARD_VM_FIXED0600_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0600X950g,
	"STANDARD_VM_FIXED0700_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0700X950g,
	"STANDARD_VM_FIXED0800_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0800X950g,
	"STANDARD_VM_FIXED0900_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0900X950g,
	"STANDARD_VM_FIXED1000_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1000X950g,
	"STANDARD_VM_FIXED1100_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1100X950g,
	"STANDARD_VM_FIXED1200_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1200X950g,
	"STANDARD_VM_FIXED1300_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1300X950g,
	"STANDARD_VM_FIXED1400_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1400X950g,
	"STANDARD_VM_FIXED1500_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1500X950g,
	"STANDARD_VM_FIXED1600_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1600X950g,
	"STANDARD_VM_FIXED1700_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1700X950g,
	"STANDARD_VM_FIXED1800_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1800X950g,
	"STANDARD_VM_FIXED1900_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1900X950g,
	"STANDARD_VM_FIXED2000_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2000X950g,
	"STANDARD_VM_FIXED2100_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2100X950g,
	"STANDARD_VM_FIXED2200_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2200X950g,
	"STANDARD_VM_FIXED2300_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2300X950g,
	"STANDARD_VM_FIXED2400_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2400X950g,
	"STANDARD_VM_FIXED2500_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2500X950g,
	"STANDARD_VM_FIXED2600_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2600X950g,
	"STANDARD_VM_FIXED2700_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2700X950g,
	"STANDARD_VM_FIXED2800_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2800X950g,
	"STANDARD_VM_FIXED2900_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2900X950g,
	"STANDARD_VM_FIXED3000_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed3000X950g,
	"STANDARD_VM_FIXED3100_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed3100X950g,
	"STANDARD_VM_FIXED3200_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed3200X950g,
	"STANDARD_VM_FIXED3300_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed3300X950g,
	"STANDARD_VM_FIXED3400_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed3400X950g,
	"STANDARD_VM_FIXED3500_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed3500X950g,
	"STANDARD_VM_FIXED3600_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed3600X950g,
	"STANDARD_VM_FIXED3700_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed3700X950g,
	"STANDARD_VM_FIXED3800_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed3800X950g,
	"STANDARD_VM_FIXED3900_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed3900X950g,
	"STANDARD_VM_FIXED4000_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed4000X950g,
	"STANDARD_VM_FIXED4100_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed4100X950g,
	"STANDARD_VM_FIXED4200_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed4200X950g,
	"STANDARD_VM_FIXED4300_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed4300X950g,
	"STANDARD_VM_FIXED4400_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed4400X950g,
	"STANDARD_VM_FIXED4500_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed4500X950g,
	"STANDARD_VM_FIXED4600_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed4600X950g,
	"STANDARD_VM_FIXED4700_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed4700X950g,
	"STANDARD_VM_FIXED4800_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed4800X950g,
	"STANDARD_VM_FIXED4900_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed4900X950g,
	"STANDARD_VM_FIXED5000_X9_50G":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed5000X950g,
	"A10_GPU_VM_FIXED2400_X9_50G":          UpdateVnicShapeDetailsVnicShapeA10GpuVmFixed2400X950g,
	"A10_GPU_VM_FIXED4800_X9_50G":          UpdateVnicShapeDetailsVnicShapeA10GpuVmFixed4800X950g,
	"VNUMA_VM_FIXED0100_1_X9_50G":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed01001X950g,
	"VNUMA_VM_FIXED0200_2_X9_50G":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed02002X950g,
	"VNUMA_VM_FIXED0300_3_X9_50G":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed03003X950g,
	"VNUMA_VM_FIXED0400_4_X9_50G":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed04004X950g,
	"VNUMA_VM_FIXED0500_5_X9_50G":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed05005X950g,
	"VNUMA_VM_FIXED0600_6_X9_50G":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed06006X950g,
	"VNUMA_VM_FIXED0700_7_X9_50G":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed07007X950g,
	"VNUMA_VM_FIXED0800_8_X9_50G":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed08008X950g,
	"VNUMA_VM_FIXED0900_9_X9_50G":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed09009X950g,
	"VNUMA_VM_FIXED1000_10_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed100010X950g,
	"VNUMA_VM_FIXED1100_11_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed110011X950g,
	"VNUMA_VM_FIXED1200_12_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed120012X950g,
	"VNUMA_VM_FIXED1300_13_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed130013X950g,
	"VNUMA_VM_FIXED1400_14_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed140014X950g,
	"VNUMA_VM_FIXED1500_15_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed150015X950g,
	"VNUMA_VM_FIXED1600_16_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed160016X950g,
	"VNUMA_VM_FIXED1700_17_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed170017X950g,
	"VNUMA_VM_FIXED1800_18_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed180018X950g,
	"VNUMA_VM_FIXED1900_19_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed190019X950g,
	"VNUMA_VM_FIXED2000_20_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed200020X950g,
	"VNUMA_VM_FIXED2100_21_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed210021X950g,
	"VNUMA_VM_FIXED2200_22_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed220022X950g,
	"VNUMA_VM_FIXED2300_23_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed230023X950g,
	"VNUMA_VM_FIXED2400_24_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed240024X950g,
	"VNUMA_VM_FIXED2500_25_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed250025X950g,
	"VNUMA_VM_FIXED2600_26_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed260026X950g,
	"VNUMA_VM_FIXED2700_27_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed270027X950g,
	"VNUMA_VM_FIXED2800_28_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed280028X950g,
	"VNUMA_VM_FIXED2900_29_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed290029X950g,
	"VNUMA_VM_FIXED3000_30_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed300030X950g,
	"VNUMA_VM_FIXED3100_31_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed310031X950g,
	"VNUMA_VM_FIXED3200_32_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320032X950g,
	"VNUMA_VM_FIXED3200_33_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320033X950g,
	"VNUMA_VM_FIXED3200_34_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320034X950g,
	"VNUMA_VM_FIXED3200_35_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320035X950g,
	"VNUMA_VM_FIXED3200_36_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320036X950g,
	"VNUMA_VM_FIXED3200_37_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320037X950g,
	"VNUMA_VM_FIXED3200_38_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320038X950g,
	"VNUMA_VM_FIXED3200_39_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320039X950g,
	"VNUMA_VM_FIXED3200_40_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320040X950g,
	"VNUMA_VM_FIXED3200_41_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320041X950g,
	"VNUMA_VM_FIXED3200_42_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320042X950g,
	"VNUMA_VM_FIXED3200_43_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320043X950g,
	"VNUMA_VM_FIXED3200_44_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320044X950g,
	"VNUMA_VM_FIXED3200_45_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320045X950g,
	"VNUMA_VM_FIXED3200_46_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320046X950g,
	"VNUMA_VM_FIXED3200_47_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320047X950g,
	"VNUMA_VM_FIXED3200_48_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320048X950g,
	"VNUMA_VM_FIXED3200_49_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320049X950g,
	"VNUMA_VM_FIXED3200_50_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320050X950g,
	"VNUMA_VM_FIXED3200_51_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320051X950g,
	"VNUMA_VM_FIXED3200_52_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320052X950g,
	"VNUMA_VM_FIXED3200_53_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320053X950g,
	"VNUMA_VM_FIXED3200_54_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320054X950g,
	"VNUMA_VM_FIXED3200_55_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320055X950g,
	"VNUMA_VM_FIXED3200_56_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320056X950g,
	"VNUMA_VM_FIXED3200_57_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320057X950g,
	"VNUMA_VM_FIXED3200_58_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320058X950g,
	"VNUMA_VM_FIXED3200_59_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320059X950g,
	"VNUMA_VM_FIXED3200_60_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320060X950g,
	"VNUMA_VM_FIXED3200_61_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320061X950g,
	"VNUMA_VM_FIXED3200_62_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320062X950g,
	"VNUMA_VM_FIXED3200_63_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320063X950g,
	"VNUMA_VM_FIXED3200_64_X9_50G":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320064X950g,
	"SUBCORE_STANDARD_VM_FIXED0025_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0025X950g,
	"SUBCORE_STANDARD_VM_FIXED0050_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0050X950g,
	"SUBCORE_STANDARD_VM_FIXED0075_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0075X950g,
	"SUBCORE_STANDARD_VM_FIXED0100_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0100X950g,
	"SUBCORE_STANDARD_VM_FIXED0125_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0125X950g,
	"SUBCORE_STANDARD_VM_FIXED0150_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0150X950g,
	"SUBCORE_STANDARD_VM_FIXED0175_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0175X950g,
	"SUBCORE_STANDARD_VM_FIXED0200_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0200X950g,
	"SUBCORE_STANDARD_VM_FIXED0225_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0225X950g,
	"SUBCORE_STANDARD_VM_FIXED0250_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0250X950g,
	"SUBCORE_STANDARD_VM_FIXED0275_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0275X950g,
	"SUBCORE_STANDARD_VM_FIXED0300_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0300X950g,
	"SUBCORE_STANDARD_VM_FIXED0325_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0325X950g,
	"SUBCORE_STANDARD_VM_FIXED0350_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0350X950g,
	"SUBCORE_STANDARD_VM_FIXED0375_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0375X950g,
	"SUBCORE_STANDARD_VM_FIXED0400_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0400X950g,
	"SUBCORE_STANDARD_VM_FIXED0425_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0425X950g,
	"SUBCORE_STANDARD_VM_FIXED0450_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0450X950g,
	"SUBCORE_STANDARD_VM_FIXED0475_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0475X950g,
	"SUBCORE_STANDARD_VM_FIXED0500_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0500X950g,
	"SUBCORE_STANDARD_VM_FIXED0525_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0525X950g,
	"SUBCORE_STANDARD_VM_FIXED0550_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0550X950g,
	"SUBCORE_STANDARD_VM_FIXED0575_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0575X950g,
	"SUBCORE_STANDARD_VM_FIXED0600_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0600X950g,
	"SUBCORE_STANDARD_VM_FIXED0625_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0625X950g,
	"SUBCORE_STANDARD_VM_FIXED0650_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0650X950g,
	"SUBCORE_STANDARD_VM_FIXED0675_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0675X950g,
	"SUBCORE_STANDARD_VM_FIXED0700_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0700X950g,
	"SUBCORE_STANDARD_VM_FIXED0725_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0725X950g,
	"SUBCORE_STANDARD_VM_FIXED0750_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0750X950g,
	"SUBCORE_STANDARD_VM_FIXED0775_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0775X950g,
	"SUBCORE_STANDARD_VM_FIXED0800_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0800X950g,
	"SUBCORE_STANDARD_VM_FIXED0825_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0825X950g,
	"SUBCORE_STANDARD_VM_FIXED0850_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0850X950g,
	"SUBCORE_STANDARD_VM_FIXED0875_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0875X950g,
	"SUBCORE_STANDARD_VM_FIXED0900_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0900X950g,
	"SUBCORE_STANDARD_VM_FIXED0925_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0925X950g,
	"SUBCORE_STANDARD_VM_FIXED0950_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0950X950g,
	"SUBCORE_STANDARD_VM_FIXED0975_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0975X950g,
	"SUBCORE_STANDARD_VM_FIXED1000_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1000X950g,
	"SUBCORE_STANDARD_VM_FIXED1025_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1025X950g,
	"SUBCORE_STANDARD_VM_FIXED1050_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1050X950g,
	"SUBCORE_STANDARD_VM_FIXED1075_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1075X950g,
	"SUBCORE_STANDARD_VM_FIXED1100_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1100X950g,
	"SUBCORE_STANDARD_VM_FIXED1125_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1125X950g,
	"SUBCORE_STANDARD_VM_FIXED1150_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1150X950g,
	"SUBCORE_STANDARD_VM_FIXED1175_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1175X950g,
	"SUBCORE_STANDARD_VM_FIXED1200_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1200X950g,
	"SUBCORE_STANDARD_VM_FIXED1225_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1225X950g,
	"SUBCORE_STANDARD_VM_FIXED1250_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1250X950g,
	"SUBCORE_STANDARD_VM_FIXED1275_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1275X950g,
	"SUBCORE_STANDARD_VM_FIXED1300_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1300X950g,
	"SUBCORE_STANDARD_VM_FIXED1325_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1325X950g,
	"SUBCORE_STANDARD_VM_FIXED1350_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1350X950g,
	"SUBCORE_STANDARD_VM_FIXED1375_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1375X950g,
	"SUBCORE_STANDARD_VM_FIXED1400_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1400X950g,
	"SUBCORE_STANDARD_VM_FIXED1425_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1425X950g,
	"SUBCORE_STANDARD_VM_FIXED1450_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1450X950g,
	"SUBCORE_STANDARD_VM_FIXED1475_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1475X950g,
	"SUBCORE_STANDARD_VM_FIXED1500_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1500X950g,
	"SUBCORE_STANDARD_VM_FIXED1525_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1525X950g,
	"SUBCORE_STANDARD_VM_FIXED1550_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1550X950g,
	"SUBCORE_STANDARD_VM_FIXED1575_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1575X950g,
	"SUBCORE_STANDARD_VM_FIXED1600_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1600X950g,
	"SUBCORE_STANDARD_VM_FIXED1625_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1625X950g,
	"SUBCORE_STANDARD_VM_FIXED1650_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1650X950g,
	"SUBCORE_STANDARD_VM_FIXED1700_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1700X950g,
	"SUBCORE_STANDARD_VM_FIXED1725_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1725X950g,
	"SUBCORE_STANDARD_VM_FIXED1750_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1750X950g,
	"SUBCORE_STANDARD_VM_FIXED1800_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1800X950g,
	"SUBCORE_STANDARD_VM_FIXED1850_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1850X950g,
	"SUBCORE_STANDARD_VM_FIXED1875_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1875X950g,
	"SUBCORE_STANDARD_VM_FIXED1900_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1900X950g,
	"SUBCORE_STANDARD_VM_FIXED1925_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1925X950g,
	"SUBCORE_STANDARD_VM_FIXED1950_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1950X950g,
	"SUBCORE_STANDARD_VM_FIXED2000_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2000X950g,
	"SUBCORE_STANDARD_VM_FIXED2025_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2025X950g,
	"SUBCORE_STANDARD_VM_FIXED2050_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2050X950g,
	"SUBCORE_STANDARD_VM_FIXED2100_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2100X950g,
	"SUBCORE_STANDARD_VM_FIXED2125_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2125X950g,
	"SUBCORE_STANDARD_VM_FIXED2150_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2150X950g,
	"SUBCORE_STANDARD_VM_FIXED2175_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2175X950g,
	"SUBCORE_STANDARD_VM_FIXED2200_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2200X950g,
	"SUBCORE_STANDARD_VM_FIXED2250_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2250X950g,
	"SUBCORE_STANDARD_VM_FIXED2275_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2275X950g,
	"SUBCORE_STANDARD_VM_FIXED2300_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2300X950g,
	"SUBCORE_STANDARD_VM_FIXED2325_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2325X950g,
	"SUBCORE_STANDARD_VM_FIXED2350_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2350X950g,
	"SUBCORE_STANDARD_VM_FIXED2375_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2375X950g,
	"SUBCORE_STANDARD_VM_FIXED2400_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2400X950g,
	"SUBCORE_STANDARD_VM_FIXED2450_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2450X950g,
	"SUBCORE_STANDARD_VM_FIXED2475_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2475X950g,
	"SUBCORE_STANDARD_VM_FIXED2500_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2500X950g,
	"SUBCORE_STANDARD_VM_FIXED2550_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2550X950g,
	"SUBCORE_STANDARD_VM_FIXED2600_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2600X950g,
	"SUBCORE_STANDARD_VM_FIXED2625_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2625X950g,
	"SUBCORE_STANDARD_VM_FIXED2650_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2650X950g,
	"SUBCORE_STANDARD_VM_FIXED2700_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2700X950g,
	"SUBCORE_STANDARD_VM_FIXED2750_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2750X950g,
	"SUBCORE_STANDARD_VM_FIXED2775_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2775X950g,
	"SUBCORE_STANDARD_VM_FIXED2800_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2800X950g,
	"SUBCORE_STANDARD_VM_FIXED2850_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2850X950g,
	"SUBCORE_STANDARD_VM_FIXED2875_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2875X950g,
	"SUBCORE_STANDARD_VM_FIXED2900_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2900X950g,
	"SUBCORE_STANDARD_VM_FIXED2925_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2925X950g,
	"SUBCORE_STANDARD_VM_FIXED2950_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2950X950g,
	"SUBCORE_STANDARD_VM_FIXED2975_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2975X950g,
	"SUBCORE_STANDARD_VM_FIXED3000_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3000X950g,
	"SUBCORE_STANDARD_VM_FIXED3025_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3025X950g,
	"SUBCORE_STANDARD_VM_FIXED3050_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3050X950g,
	"SUBCORE_STANDARD_VM_FIXED3075_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3075X950g,
	"SUBCORE_STANDARD_VM_FIXED3100_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3100X950g,
	"SUBCORE_STANDARD_VM_FIXED3125_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3125X950g,
	"SUBCORE_STANDARD_VM_FIXED3150_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3150X950g,
	"SUBCORE_STANDARD_VM_FIXED3200_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3200X950g,
	"SUBCORE_STANDARD_VM_FIXED3225_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3225X950g,
	"SUBCORE_STANDARD_VM_FIXED3250_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3250X950g,
	"SUBCORE_STANDARD_VM_FIXED3300_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3300X950g,
	"SUBCORE_STANDARD_VM_FIXED3325_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3325X950g,
	"SUBCORE_STANDARD_VM_FIXED3375_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3375X950g,
	"SUBCORE_STANDARD_VM_FIXED3400_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3400X950g,
	"SUBCORE_STANDARD_VM_FIXED3450_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3450X950g,
	"SUBCORE_STANDARD_VM_FIXED3500_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3500X950g,
	"SUBCORE_STANDARD_VM_FIXED3525_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3525X950g,
	"SUBCORE_STANDARD_VM_FIXED3575_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3575X950g,
	"SUBCORE_STANDARD_VM_FIXED3600_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3600X950g,
	"SUBCORE_STANDARD_VM_FIXED3625_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3625X950g,
	"SUBCORE_STANDARD_VM_FIXED3675_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3675X950g,
	"SUBCORE_STANDARD_VM_FIXED3700_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3700X950g,
	"SUBCORE_STANDARD_VM_FIXED3750_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3750X950g,
	"SUBCORE_STANDARD_VM_FIXED3800_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3800X950g,
	"SUBCORE_STANDARD_VM_FIXED3825_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3825X950g,
	"SUBCORE_STANDARD_VM_FIXED3850_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3850X950g,
	"SUBCORE_STANDARD_VM_FIXED3875_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3875X950g,
	"SUBCORE_STANDARD_VM_FIXED3900_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3900X950g,
	"SUBCORE_STANDARD_VM_FIXED3975_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3975X950g,
	"SUBCORE_STANDARD_VM_FIXED4000_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4000X950g,
	"SUBCORE_STANDARD_VM_FIXED4025_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4025X950g,
	"SUBCORE_STANDARD_VM_FIXED4050_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4050X950g,
	"SUBCORE_STANDARD_VM_FIXED4100_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4100X950g,
	"SUBCORE_STANDARD_VM_FIXED4125_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4125X950g,
	"SUBCORE_STANDARD_VM_FIXED4200_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4200X950g,
	"SUBCORE_STANDARD_VM_FIXED4225_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4225X950g,
	"SUBCORE_STANDARD_VM_FIXED4250_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4250X950g,
	"SUBCORE_STANDARD_VM_FIXED4275_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4275X950g,
	"SUBCORE_STANDARD_VM_FIXED4300_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4300X950g,
	"SUBCORE_STANDARD_VM_FIXED4350_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4350X950g,
	"SUBCORE_STANDARD_VM_FIXED4375_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4375X950g,
	"SUBCORE_STANDARD_VM_FIXED4400_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4400X950g,
	"SUBCORE_STANDARD_VM_FIXED4425_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4425X950g,
	"SUBCORE_STANDARD_VM_FIXED4500_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4500X950g,
	"SUBCORE_STANDARD_VM_FIXED4550_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4550X950g,
	"SUBCORE_STANDARD_VM_FIXED4575_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4575X950g,
	"SUBCORE_STANDARD_VM_FIXED4600_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4600X950g,
	"SUBCORE_STANDARD_VM_FIXED4625_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4625X950g,
	"SUBCORE_STANDARD_VM_FIXED4650_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4650X950g,
	"SUBCORE_STANDARD_VM_FIXED4675_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4675X950g,
	"SUBCORE_STANDARD_VM_FIXED4700_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4700X950g,
	"SUBCORE_STANDARD_VM_FIXED4725_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4725X950g,
	"SUBCORE_STANDARD_VM_FIXED4750_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4750X950g,
	"SUBCORE_STANDARD_VM_FIXED4800_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4800X950g,
	"SUBCORE_STANDARD_VM_FIXED4875_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4875X950g,
	"SUBCORE_STANDARD_VM_FIXED4900_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4900X950g,
	"SUBCORE_STANDARD_VM_FIXED4950_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4950X950g,
	"SUBCORE_STANDARD_VM_FIXED5000_X9_50G": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed5000X950g,
	"ENTIREHOST_X9_50G":                    UpdateVnicShapeDetailsVnicShapeEntirehostX950g,
	"DYNAMIC_X9_100G":                      UpdateVnicShapeDetailsVnicShapeDynamicX9100g,
	"DYNAMIC_X10_50G":                      UpdateVnicShapeDetailsVnicShapeDynamicX1050g,
	"DYNAMIC_X10_100G":                     UpdateVnicShapeDetailsVnicShapeDynamicX10100g,
}

var mappingUpdateVnicShapeDetailsVnicShapeEnumLowerCase = map[string]UpdateVnicShapeDetailsVnicShapeEnum{
	"dynamic":                              UpdateVnicShapeDetailsVnicShapeDynamic,
	"fixed0040":                            UpdateVnicShapeDetailsVnicShapeFixed0040,
	"fixed0060":                            UpdateVnicShapeDetailsVnicShapeFixed0060,
	"fixed0060_psm":                        UpdateVnicShapeDetailsVnicShapeFixed0060Psm,
	"fixed0100":                            UpdateVnicShapeDetailsVnicShapeFixed0100,
	"fixed0120":                            UpdateVnicShapeDetailsVnicShapeFixed0120,
	"fixed0120_2x":                         UpdateVnicShapeDetailsVnicShapeFixed01202x,
	"fixed0200":                            UpdateVnicShapeDetailsVnicShapeFixed0200,
	"fixed0240":                            UpdateVnicShapeDetailsVnicShapeFixed0240,
	"fixed0480":                            UpdateVnicShapeDetailsVnicShapeFixed0480,
	"entirehost":                           UpdateVnicShapeDetailsVnicShapeEntirehost,
	"dynamic_25g":                          UpdateVnicShapeDetailsVnicShapeDynamic25g,
	"fixed0040_25g":                        UpdateVnicShapeDetailsVnicShapeFixed004025g,
	"fixed0100_25g":                        UpdateVnicShapeDetailsVnicShapeFixed010025g,
	"fixed0200_25g":                        UpdateVnicShapeDetailsVnicShapeFixed020025g,
	"fixed0400_25g":                        UpdateVnicShapeDetailsVnicShapeFixed040025g,
	"fixed0800_25g":                        UpdateVnicShapeDetailsVnicShapeFixed080025g,
	"fixed1600_25g":                        UpdateVnicShapeDetailsVnicShapeFixed160025g,
	"fixed2400_25g":                        UpdateVnicShapeDetailsVnicShapeFixed240025g,
	"entirehost_25g":                       UpdateVnicShapeDetailsVnicShapeEntirehost25g,
	"dynamic_e1_25g":                       UpdateVnicShapeDetailsVnicShapeDynamicE125g,
	"fixed0040_e1_25g":                     UpdateVnicShapeDetailsVnicShapeFixed0040E125g,
	"fixed0070_e1_25g":                     UpdateVnicShapeDetailsVnicShapeFixed0070E125g,
	"fixed0140_e1_25g":                     UpdateVnicShapeDetailsVnicShapeFixed0140E125g,
	"fixed0280_e1_25g":                     UpdateVnicShapeDetailsVnicShapeFixed0280E125g,
	"fixed0560_e1_25g":                     UpdateVnicShapeDetailsVnicShapeFixed0560E125g,
	"fixed1120_e1_25g":                     UpdateVnicShapeDetailsVnicShapeFixed1120E125g,
	"fixed1680_e1_25g":                     UpdateVnicShapeDetailsVnicShapeFixed1680E125g,
	"entirehost_e1_25g":                    UpdateVnicShapeDetailsVnicShapeEntirehostE125g,
	"dynamic_b1_25g":                       UpdateVnicShapeDetailsVnicShapeDynamicB125g,
	"fixed0040_b1_25g":                     UpdateVnicShapeDetailsVnicShapeFixed0040B125g,
	"fixed0060_b1_25g":                     UpdateVnicShapeDetailsVnicShapeFixed0060B125g,
	"fixed0120_b1_25g":                     UpdateVnicShapeDetailsVnicShapeFixed0120B125g,
	"fixed0240_b1_25g":                     UpdateVnicShapeDetailsVnicShapeFixed0240B125g,
	"fixed0480_b1_25g":                     UpdateVnicShapeDetailsVnicShapeFixed0480B125g,
	"fixed0960_b1_25g":                     UpdateVnicShapeDetailsVnicShapeFixed0960B125g,
	"entirehost_b1_25g":                    UpdateVnicShapeDetailsVnicShapeEntirehostB125g,
	"micro_vm_fixed0048_e1_25g":            UpdateVnicShapeDetailsVnicShapeMicroVmFixed0048E125g,
	"micro_lb_fixed0001_e1_25g":            UpdateVnicShapeDetailsVnicShapeMicroLbFixed0001E125g,
	"vnicaas_fixed0025":                    UpdateVnicShapeDetailsVnicShapeVnicaasFixed0025,
	"vnicaas_fixed0200":                    UpdateVnicShapeDetailsVnicShapeVnicaasFixed0200,
	"vnicaas_fixed0400":                    UpdateVnicShapeDetailsVnicShapeVnicaasFixed0400,
	"vnicaas_fixed0625":                    UpdateVnicShapeDetailsVnicShapeVnicaasFixed0625,
	"vnicaas_fixed0700":                    UpdateVnicShapeDetailsVnicShapeVnicaasFixed0700,
	"vnicaas_fixed1250":                    UpdateVnicShapeDetailsVnicShapeVnicaasFixed1250,
	"vnicaas_fwaas_6g_250m":                UpdateVnicShapeDetailsVnicShapeVnicaasFwaas6g250m,
	"vnicaas_nlb_approved_10g":             UpdateVnicShapeDetailsVnicShapeVnicaasNlbApproved10g,
	"vnicaas_nlb_approved_25g":             UpdateVnicShapeDetailsVnicShapeVnicaasNlbApproved25g,
	"vnicaas_telesis_25g":                  UpdateVnicShapeDetailsVnicShapeVnicaasTelesis25g,
	"vnicaas_telesis_10g":                  UpdateVnicShapeDetailsVnicShapeVnicaasTelesis10g,
	"vnicaas_ambassador_fixed0100":         UpdateVnicShapeDetailsVnicShapeVnicaasAmbassadorFixed0100,
	"vnicaas_telesis_gamma":                UpdateVnicShapeDetailsVnicShapeVnicaasTelesisGamma,
	"vnicaas_privatedns":                   UpdateVnicShapeDetailsVnicShapeVnicaasPrivatedns,
	"vnicaas_fwaas":                        UpdateVnicShapeDetailsVnicShapeVnicaasFwaas,
	"vnicaas_lbaas_free":                   UpdateVnicShapeDetailsVnicShapeVnicaasLbaasFree,
	"vnicaas_lbaas_8g_512k":                UpdateVnicShapeDetailsVnicShapeVnicaasLbaas8g512k,
	"vnicaas_lbaas_8g_1m":                  UpdateVnicShapeDetailsVnicShapeVnicaasLbaas8g1m,
	"vnicaas_lbaas_8g_2m":                  UpdateVnicShapeDetailsVnicShapeVnicaasLbaas8g2m,
	"vnicaas_lbaas_8g_3m":                  UpdateVnicShapeDetailsVnicShapeVnicaasLbaas8g3m,
	"vnicaas_lbaas_8g_1m_8ghost":           UpdateVnicShapeDetailsVnicShapeVnicaasLbaas8g1m8ghost,
	"vnicaas_lbaas_8g_1m_16ghost":          UpdateVnicShapeDetailsVnicShapeVnicaasLbaas8g1m16ghost,
	"dynamic_e3_50g":                       UpdateVnicShapeDetailsVnicShapeDynamicE350g,
	"fixed0040_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0040E350g,
	"fixed0100_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0100E350g,
	"fixed0200_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0200E350g,
	"fixed0300_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0300E350g,
	"fixed0400_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0400E350g,
	"fixed0500_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0500E350g,
	"fixed0600_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0600E350g,
	"fixed0700_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0700E350g,
	"fixed0800_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0800E350g,
	"fixed0900_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0900E350g,
	"fixed1000_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1000E350g,
	"fixed1100_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1100E350g,
	"fixed1200_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1200E350g,
	"fixed1300_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1300E350g,
	"fixed1400_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1400E350g,
	"fixed1500_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1500E350g,
	"fixed1600_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1600E350g,
	"fixed1700_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1700E350g,
	"fixed1800_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1800E350g,
	"fixed1900_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1900E350g,
	"fixed2000_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2000E350g,
	"fixed2100_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2100E350g,
	"fixed2200_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2200E350g,
	"fixed2300_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2300E350g,
	"fixed2400_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2400E350g,
	"fixed2500_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2500E350g,
	"fixed2600_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2600E350g,
	"fixed2700_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2700E350g,
	"fixed2800_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2800E350g,
	"fixed2900_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2900E350g,
	"fixed3000_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3000E350g,
	"fixed3100_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3100E350g,
	"fixed3200_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3200E350g,
	"fixed3300_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3300E350g,
	"fixed3400_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3400E350g,
	"fixed3500_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3500E350g,
	"fixed3600_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3600E350g,
	"fixed3700_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3700E350g,
	"fixed3800_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3800E350g,
	"fixed3900_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3900E350g,
	"fixed4000_e3_50g":                     UpdateVnicShapeDetailsVnicShapeFixed4000E350g,
	"entirehost_e3_50g":                    UpdateVnicShapeDetailsVnicShapeEntirehostE350g,
	"dynamic_e4_50g":                       UpdateVnicShapeDetailsVnicShapeDynamicE450g,
	"fixed0040_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0040E450g,
	"fixed0100_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0100E450g,
	"fixed0200_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0200E450g,
	"fixed0300_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0300E450g,
	"fixed0400_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0400E450g,
	"fixed0500_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0500E450g,
	"fixed0600_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0600E450g,
	"fixed0700_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0700E450g,
	"fixed0800_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0800E450g,
	"fixed0900_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0900E450g,
	"fixed1000_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1000E450g,
	"fixed1100_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1100E450g,
	"fixed1200_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1200E450g,
	"fixed1300_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1300E450g,
	"fixed1400_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1400E450g,
	"fixed1500_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1500E450g,
	"fixed1600_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1600E450g,
	"fixed1700_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1700E450g,
	"fixed1800_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1800E450g,
	"fixed1900_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1900E450g,
	"fixed2000_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2000E450g,
	"fixed2100_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2100E450g,
	"fixed2200_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2200E450g,
	"fixed2300_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2300E450g,
	"fixed2400_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2400E450g,
	"fixed2500_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2500E450g,
	"fixed2600_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2600E450g,
	"fixed2700_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2700E450g,
	"fixed2800_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2800E450g,
	"fixed2900_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2900E450g,
	"fixed3000_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3000E450g,
	"fixed3100_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3100E450g,
	"fixed3200_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3200E450g,
	"fixed3300_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3300E450g,
	"fixed3400_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3400E450g,
	"fixed3500_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3500E450g,
	"fixed3600_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3600E450g,
	"fixed3700_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3700E450g,
	"fixed3800_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3800E450g,
	"fixed3900_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3900E450g,
	"fixed4000_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed4000E450g,
	"fixed4700_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed4700E450g,
	"fixed4960_e4_50g":                     UpdateVnicShapeDetailsVnicShapeFixed4960E450g,
	"entirehost_e4_50g":                    UpdateVnicShapeDetailsVnicShapeEntirehostE450g,
	"vnuma_vm_fixed0100_1_e4_50g":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed01001E450g,
	"vnuma_vm_fixed0200_2_e4_50g":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed02002E450g,
	"vnuma_vm_fixed0300_3_e4_50g":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed03003E450g,
	"vnuma_vm_fixed0400_4_e4_50g":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed04004E450g,
	"vnuma_vm_fixed0500_5_e4_50g":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed05005E450g,
	"vnuma_vm_fixed0600_6_e4_50g":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed06006E450g,
	"vnuma_vm_fixed0700_7_e4_50g":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed07007E450g,
	"vnuma_vm_fixed0800_8_e4_50g":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed08008E450g,
	"vnuma_vm_fixed0900_9_e4_50g":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed09009E450g,
	"vnuma_vm_fixed1000_10_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed100010E450g,
	"vnuma_vm_fixed1100_11_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed110011E450g,
	"vnuma_vm_fixed1200_12_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed120012E450g,
	"vnuma_vm_fixed1300_13_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed130013E450g,
	"vnuma_vm_fixed1400_14_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed140014E450g,
	"vnuma_vm_fixed1500_15_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed150015E450g,
	"vnuma_vm_fixed1600_16_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed160016E450g,
	"vnuma_vm_fixed1700_17_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed170017E450g,
	"vnuma_vm_fixed1800_18_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed180018E450g,
	"vnuma_vm_fixed1900_19_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed190019E450g,
	"vnuma_vm_fixed2000_20_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed200020E450g,
	"vnuma_vm_fixed2100_21_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed210021E450g,
	"vnuma_vm_fixed2200_22_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed220022E450g,
	"vnuma_vm_fixed2300_23_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed230023E450g,
	"vnuma_vm_fixed2400_24_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed240024E450g,
	"vnuma_vm_fixed2500_25_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed250025E450g,
	"vnuma_vm_fixed2600_26_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed260026E450g,
	"vnuma_vm_fixed2700_27_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed270027E450g,
	"vnuma_vm_fixed2800_28_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed280028E450g,
	"vnuma_vm_fixed2900_29_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed290029E450g,
	"vnuma_vm_fixed3000_30_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed300030E450g,
	"vnuma_vm_fixed3100_31_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed310031E450g,
	"vnuma_vm_fixed3200_32_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320032E450g,
	"vnuma_vm_fixed3300_33_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed330033E450g,
	"vnuma_vm_fixed3400_34_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed340034E450g,
	"vnuma_vm_fixed3500_35_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed350035E450g,
	"vnuma_vm_fixed3600_36_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed360036E450g,
	"vnuma_vm_fixed3700_37_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed370037E450g,
	"vnuma_vm_fixed3800_38_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed380038E450g,
	"vnuma_vm_fixed3900_39_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed390039E450g,
	"vnuma_vm_fixed4000_40_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400040E450g,
	"vnuma_vm_fixed4000_41_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400041E450g,
	"vnuma_vm_fixed4000_42_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400042E450g,
	"vnuma_vm_fixed4000_43_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400043E450g,
	"vnuma_vm_fixed4000_44_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400044E450g,
	"vnuma_vm_fixed4000_45_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400045E450g,
	"vnuma_vm_fixed4000_46_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400046E450g,
	"vnuma_vm_fixed4000_47_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400047E450g,
	"vnuma_vm_fixed4000_48_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400048E450g,
	"vnuma_vm_fixed4000_49_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400049E450g,
	"vnuma_vm_fixed4000_50_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400050E450g,
	"vnuma_vm_fixed4000_51_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400051E450g,
	"vnuma_vm_fixed4000_52_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400052E450g,
	"vnuma_vm_fixed4000_53_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400053E450g,
	"vnuma_vm_fixed4000_54_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400054E450g,
	"vnuma_vm_fixed4000_55_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400055E450g,
	"vnuma_vm_fixed4000_56_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400056E450g,
	"vnuma_vm_fixed4000_57_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400057E450g,
	"vnuma_vm_fixed4000_58_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400058E450g,
	"vnuma_vm_fixed4000_59_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400059E450g,
	"vnuma_vm_fixed4000_60_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400060E450g,
	"vnuma_vm_fixed4000_61_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400061E450g,
	"vnuma_vm_fixed4000_62_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400062E450g,
	"vnuma_vm_fixed4000_63_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400063E450g,
	"vnuma_vm_fixed4000_64_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400064E450g,
	"vnuma_vm_fixed4000_65_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400065E450g,
	"vnuma_vm_fixed4000_66_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400066E450g,
	"vnuma_vm_fixed4000_67_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400067E450g,
	"vnuma_vm_fixed4000_68_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400068E450g,
	"vnuma_vm_fixed4000_69_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400069E450g,
	"vnuma_vm_fixed4000_70_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400070E450g,
	"vnuma_vm_fixed4000_71_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400071E450g,
	"vnuma_vm_fixed4000_72_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400072E450g,
	"vnuma_vm_fixed4000_73_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400073E450g,
	"vnuma_vm_fixed4000_74_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400074E450g,
	"vnuma_vm_fixed4000_75_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400075E450g,
	"vnuma_vm_fixed4000_76_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400076E450g,
	"vnuma_vm_fixed4000_77_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400077E450g,
	"vnuma_vm_fixed4000_78_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400078E450g,
	"vnuma_vm_fixed4000_79_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400079E450g,
	"vnuma_vm_fixed4000_80_e4_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed400080E450g,
	"micro_vm_fixed0050_e3_50g":            UpdateVnicShapeDetailsVnicShapeMicroVmFixed0050E350g,
	"micro_vm_fixed0050_e4_50g":            UpdateVnicShapeDetailsVnicShapeMicroVmFixed0050E450g,
	"subcore_vm_fixed0025_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0025E350g,
	"subcore_vm_fixed0050_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0050E350g,
	"subcore_vm_fixed0075_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0075E350g,
	"subcore_vm_fixed0100_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0100E350g,
	"subcore_vm_fixed0125_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0125E350g,
	"subcore_vm_fixed0150_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0150E350g,
	"subcore_vm_fixed0175_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0175E350g,
	"subcore_vm_fixed0200_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0200E350g,
	"subcore_vm_fixed0225_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0225E350g,
	"subcore_vm_fixed0250_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0250E350g,
	"subcore_vm_fixed0275_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0275E350g,
	"subcore_vm_fixed0300_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0300E350g,
	"subcore_vm_fixed0325_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0325E350g,
	"subcore_vm_fixed0350_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0350E350g,
	"subcore_vm_fixed0375_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0375E350g,
	"subcore_vm_fixed0400_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0400E350g,
	"subcore_vm_fixed0425_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0425E350g,
	"subcore_vm_fixed0450_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0450E350g,
	"subcore_vm_fixed0475_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0475E350g,
	"subcore_vm_fixed0500_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0500E350g,
	"subcore_vm_fixed0525_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0525E350g,
	"subcore_vm_fixed0550_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0550E350g,
	"subcore_vm_fixed0575_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0575E350g,
	"subcore_vm_fixed0600_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0600E350g,
	"subcore_vm_fixed0625_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0625E350g,
	"subcore_vm_fixed0650_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0650E350g,
	"subcore_vm_fixed0675_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0675E350g,
	"subcore_vm_fixed0700_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0700E350g,
	"subcore_vm_fixed0725_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0725E350g,
	"subcore_vm_fixed0750_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0750E350g,
	"subcore_vm_fixed0775_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0775E350g,
	"subcore_vm_fixed0800_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0800E350g,
	"subcore_vm_fixed0825_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0825E350g,
	"subcore_vm_fixed0850_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0850E350g,
	"subcore_vm_fixed0875_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0875E350g,
	"subcore_vm_fixed0900_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0900E350g,
	"subcore_vm_fixed0925_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0925E350g,
	"subcore_vm_fixed0950_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0950E350g,
	"subcore_vm_fixed0975_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0975E350g,
	"subcore_vm_fixed1000_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1000E350g,
	"subcore_vm_fixed1025_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1025E350g,
	"subcore_vm_fixed1050_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1050E350g,
	"subcore_vm_fixed1075_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1075E350g,
	"subcore_vm_fixed1100_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1100E350g,
	"subcore_vm_fixed1125_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1125E350g,
	"subcore_vm_fixed1150_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1150E350g,
	"subcore_vm_fixed1175_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1175E350g,
	"subcore_vm_fixed1200_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1200E350g,
	"subcore_vm_fixed1225_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1225E350g,
	"subcore_vm_fixed1250_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1250E350g,
	"subcore_vm_fixed1275_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1275E350g,
	"subcore_vm_fixed1300_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1300E350g,
	"subcore_vm_fixed1325_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1325E350g,
	"subcore_vm_fixed1350_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1350E350g,
	"subcore_vm_fixed1375_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1375E350g,
	"subcore_vm_fixed1400_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1400E350g,
	"subcore_vm_fixed1425_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1425E350g,
	"subcore_vm_fixed1450_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1450E350g,
	"subcore_vm_fixed1475_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1475E350g,
	"subcore_vm_fixed1500_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1500E350g,
	"subcore_vm_fixed1525_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1525E350g,
	"subcore_vm_fixed1550_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1550E350g,
	"subcore_vm_fixed1575_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1575E350g,
	"subcore_vm_fixed1600_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1600E350g,
	"subcore_vm_fixed1625_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1625E350g,
	"subcore_vm_fixed1650_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1650E350g,
	"subcore_vm_fixed1700_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1700E350g,
	"subcore_vm_fixed1725_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1725E350g,
	"subcore_vm_fixed1750_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1750E350g,
	"subcore_vm_fixed1800_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1800E350g,
	"subcore_vm_fixed1850_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1850E350g,
	"subcore_vm_fixed1875_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1875E350g,
	"subcore_vm_fixed1900_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1900E350g,
	"subcore_vm_fixed1925_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1925E350g,
	"subcore_vm_fixed1950_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1950E350g,
	"subcore_vm_fixed2000_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2000E350g,
	"subcore_vm_fixed2025_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2025E350g,
	"subcore_vm_fixed2050_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2050E350g,
	"subcore_vm_fixed2100_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2100E350g,
	"subcore_vm_fixed2125_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2125E350g,
	"subcore_vm_fixed2150_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2150E350g,
	"subcore_vm_fixed2175_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2175E350g,
	"subcore_vm_fixed2200_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2200E350g,
	"subcore_vm_fixed2250_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2250E350g,
	"subcore_vm_fixed2275_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2275E350g,
	"subcore_vm_fixed2300_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2300E350g,
	"subcore_vm_fixed2325_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2325E350g,
	"subcore_vm_fixed2350_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2350E350g,
	"subcore_vm_fixed2375_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2375E350g,
	"subcore_vm_fixed2400_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2400E350g,
	"subcore_vm_fixed2450_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2450E350g,
	"subcore_vm_fixed2475_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2475E350g,
	"subcore_vm_fixed2500_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2500E350g,
	"subcore_vm_fixed2550_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2550E350g,
	"subcore_vm_fixed2600_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2600E350g,
	"subcore_vm_fixed2625_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2625E350g,
	"subcore_vm_fixed2650_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2650E350g,
	"subcore_vm_fixed2700_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2700E350g,
	"subcore_vm_fixed2750_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2750E350g,
	"subcore_vm_fixed2775_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2775E350g,
	"subcore_vm_fixed2800_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2800E350g,
	"subcore_vm_fixed2850_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2850E350g,
	"subcore_vm_fixed2875_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2875E350g,
	"subcore_vm_fixed2900_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2900E350g,
	"subcore_vm_fixed2925_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2925E350g,
	"subcore_vm_fixed2950_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2950E350g,
	"subcore_vm_fixed2975_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2975E350g,
	"subcore_vm_fixed3000_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3000E350g,
	"subcore_vm_fixed3025_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3025E350g,
	"subcore_vm_fixed3050_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3050E350g,
	"subcore_vm_fixed3075_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3075E350g,
	"subcore_vm_fixed3100_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3100E350g,
	"subcore_vm_fixed3125_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3125E350g,
	"subcore_vm_fixed3150_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3150E350g,
	"subcore_vm_fixed3200_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3200E350g,
	"subcore_vm_fixed3225_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3225E350g,
	"subcore_vm_fixed3250_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3250E350g,
	"subcore_vm_fixed3300_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3300E350g,
	"subcore_vm_fixed3325_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3325E350g,
	"subcore_vm_fixed3375_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3375E350g,
	"subcore_vm_fixed3400_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3400E350g,
	"subcore_vm_fixed3450_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3450E350g,
	"subcore_vm_fixed3500_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3500E350g,
	"subcore_vm_fixed3525_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3525E350g,
	"subcore_vm_fixed3575_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3575E350g,
	"subcore_vm_fixed3600_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3600E350g,
	"subcore_vm_fixed3625_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3625E350g,
	"subcore_vm_fixed3675_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3675E350g,
	"subcore_vm_fixed3700_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3700E350g,
	"subcore_vm_fixed3750_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3750E350g,
	"subcore_vm_fixed3800_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3800E350g,
	"subcore_vm_fixed3825_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3825E350g,
	"subcore_vm_fixed3850_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3850E350g,
	"subcore_vm_fixed3875_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3875E350g,
	"subcore_vm_fixed3900_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3900E350g,
	"subcore_vm_fixed3975_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3975E350g,
	"subcore_vm_fixed4000_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4000E350g,
	"subcore_vm_fixed4025_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4025E350g,
	"subcore_vm_fixed4050_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4050E350g,
	"subcore_vm_fixed4100_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4100E350g,
	"subcore_vm_fixed4125_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4125E350g,
	"subcore_vm_fixed4200_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4200E350g,
	"subcore_vm_fixed4225_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4225E350g,
	"subcore_vm_fixed4250_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4250E350g,
	"subcore_vm_fixed4275_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4275E350g,
	"subcore_vm_fixed4300_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4300E350g,
	"subcore_vm_fixed4350_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4350E350g,
	"subcore_vm_fixed4375_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4375E350g,
	"subcore_vm_fixed4400_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4400E350g,
	"subcore_vm_fixed4425_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4425E350g,
	"subcore_vm_fixed4500_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4500E350g,
	"subcore_vm_fixed4550_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4550E350g,
	"subcore_vm_fixed4575_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4575E350g,
	"subcore_vm_fixed4600_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4600E350g,
	"subcore_vm_fixed4625_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4625E350g,
	"subcore_vm_fixed4650_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4650E350g,
	"subcore_vm_fixed4675_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4675E350g,
	"subcore_vm_fixed4700_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4700E350g,
	"subcore_vm_fixed4725_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4725E350g,
	"subcore_vm_fixed4750_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4750E350g,
	"subcore_vm_fixed4800_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4800E350g,
	"subcore_vm_fixed4875_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4875E350g,
	"subcore_vm_fixed4900_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4900E350g,
	"subcore_vm_fixed4950_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4950E350g,
	"subcore_vm_fixed5000_e3_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed5000E350g,
	"subcore_vm_fixed0025_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0025E450g,
	"subcore_vm_fixed0050_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0050E450g,
	"subcore_vm_fixed0075_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0075E450g,
	"subcore_vm_fixed0100_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0100E450g,
	"subcore_vm_fixed0125_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0125E450g,
	"subcore_vm_fixed0150_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0150E450g,
	"subcore_vm_fixed0175_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0175E450g,
	"subcore_vm_fixed0200_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0200E450g,
	"subcore_vm_fixed0225_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0225E450g,
	"subcore_vm_fixed0250_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0250E450g,
	"subcore_vm_fixed0275_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0275E450g,
	"subcore_vm_fixed0300_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0300E450g,
	"subcore_vm_fixed0325_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0325E450g,
	"subcore_vm_fixed0350_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0350E450g,
	"subcore_vm_fixed0375_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0375E450g,
	"subcore_vm_fixed0400_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0400E450g,
	"subcore_vm_fixed0425_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0425E450g,
	"subcore_vm_fixed0450_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0450E450g,
	"subcore_vm_fixed0475_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0475E450g,
	"subcore_vm_fixed0500_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0500E450g,
	"subcore_vm_fixed0525_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0525E450g,
	"subcore_vm_fixed0550_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0550E450g,
	"subcore_vm_fixed0575_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0575E450g,
	"subcore_vm_fixed0600_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0600E450g,
	"subcore_vm_fixed0625_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0625E450g,
	"subcore_vm_fixed0650_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0650E450g,
	"subcore_vm_fixed0675_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0675E450g,
	"subcore_vm_fixed0700_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0700E450g,
	"subcore_vm_fixed0725_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0725E450g,
	"subcore_vm_fixed0750_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0750E450g,
	"subcore_vm_fixed0775_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0775E450g,
	"subcore_vm_fixed0800_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0800E450g,
	"subcore_vm_fixed0825_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0825E450g,
	"subcore_vm_fixed0850_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0850E450g,
	"subcore_vm_fixed0875_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0875E450g,
	"subcore_vm_fixed0900_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0900E450g,
	"subcore_vm_fixed0925_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0925E450g,
	"subcore_vm_fixed0950_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0950E450g,
	"subcore_vm_fixed0975_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0975E450g,
	"subcore_vm_fixed1000_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1000E450g,
	"subcore_vm_fixed1025_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1025E450g,
	"subcore_vm_fixed1050_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1050E450g,
	"subcore_vm_fixed1075_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1075E450g,
	"subcore_vm_fixed1100_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1100E450g,
	"subcore_vm_fixed1125_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1125E450g,
	"subcore_vm_fixed1150_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1150E450g,
	"subcore_vm_fixed1175_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1175E450g,
	"subcore_vm_fixed1200_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1200E450g,
	"subcore_vm_fixed1225_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1225E450g,
	"subcore_vm_fixed1250_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1250E450g,
	"subcore_vm_fixed1275_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1275E450g,
	"subcore_vm_fixed1300_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1300E450g,
	"subcore_vm_fixed1325_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1325E450g,
	"subcore_vm_fixed1350_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1350E450g,
	"subcore_vm_fixed1375_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1375E450g,
	"subcore_vm_fixed1400_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1400E450g,
	"subcore_vm_fixed1425_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1425E450g,
	"subcore_vm_fixed1450_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1450E450g,
	"subcore_vm_fixed1475_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1475E450g,
	"subcore_vm_fixed1500_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1500E450g,
	"subcore_vm_fixed1525_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1525E450g,
	"subcore_vm_fixed1550_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1550E450g,
	"subcore_vm_fixed1575_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1575E450g,
	"subcore_vm_fixed1600_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1600E450g,
	"subcore_vm_fixed1625_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1625E450g,
	"subcore_vm_fixed1650_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1650E450g,
	"subcore_vm_fixed1700_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1700E450g,
	"subcore_vm_fixed1725_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1725E450g,
	"subcore_vm_fixed1750_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1750E450g,
	"subcore_vm_fixed1800_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1800E450g,
	"subcore_vm_fixed1850_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1850E450g,
	"subcore_vm_fixed1875_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1875E450g,
	"subcore_vm_fixed1900_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1900E450g,
	"subcore_vm_fixed1925_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1925E450g,
	"subcore_vm_fixed1950_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1950E450g,
	"subcore_vm_fixed2000_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2000E450g,
	"subcore_vm_fixed2025_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2025E450g,
	"subcore_vm_fixed2050_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2050E450g,
	"subcore_vm_fixed2100_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2100E450g,
	"subcore_vm_fixed2125_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2125E450g,
	"subcore_vm_fixed2150_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2150E450g,
	"subcore_vm_fixed2175_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2175E450g,
	"subcore_vm_fixed2200_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2200E450g,
	"subcore_vm_fixed2250_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2250E450g,
	"subcore_vm_fixed2275_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2275E450g,
	"subcore_vm_fixed2300_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2300E450g,
	"subcore_vm_fixed2325_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2325E450g,
	"subcore_vm_fixed2350_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2350E450g,
	"subcore_vm_fixed2375_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2375E450g,
	"subcore_vm_fixed2400_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2400E450g,
	"subcore_vm_fixed2450_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2450E450g,
	"subcore_vm_fixed2475_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2475E450g,
	"subcore_vm_fixed2500_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2500E450g,
	"subcore_vm_fixed2550_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2550E450g,
	"subcore_vm_fixed2600_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2600E450g,
	"subcore_vm_fixed2625_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2625E450g,
	"subcore_vm_fixed2650_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2650E450g,
	"subcore_vm_fixed2700_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2700E450g,
	"subcore_vm_fixed2750_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2750E450g,
	"subcore_vm_fixed2775_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2775E450g,
	"subcore_vm_fixed2800_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2800E450g,
	"subcore_vm_fixed2850_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2850E450g,
	"subcore_vm_fixed2875_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2875E450g,
	"subcore_vm_fixed2900_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2900E450g,
	"subcore_vm_fixed2925_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2925E450g,
	"subcore_vm_fixed2950_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2950E450g,
	"subcore_vm_fixed2975_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2975E450g,
	"subcore_vm_fixed3000_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3000E450g,
	"subcore_vm_fixed3025_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3025E450g,
	"subcore_vm_fixed3050_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3050E450g,
	"subcore_vm_fixed3075_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3075E450g,
	"subcore_vm_fixed3100_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3100E450g,
	"subcore_vm_fixed3125_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3125E450g,
	"subcore_vm_fixed3150_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3150E450g,
	"subcore_vm_fixed3200_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3200E450g,
	"subcore_vm_fixed3225_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3225E450g,
	"subcore_vm_fixed3250_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3250E450g,
	"subcore_vm_fixed3300_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3300E450g,
	"subcore_vm_fixed3325_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3325E450g,
	"subcore_vm_fixed3375_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3375E450g,
	"subcore_vm_fixed3400_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3400E450g,
	"subcore_vm_fixed3450_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3450E450g,
	"subcore_vm_fixed3500_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3500E450g,
	"subcore_vm_fixed3525_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3525E450g,
	"subcore_vm_fixed3575_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3575E450g,
	"subcore_vm_fixed3600_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3600E450g,
	"subcore_vm_fixed3625_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3625E450g,
	"subcore_vm_fixed3675_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3675E450g,
	"subcore_vm_fixed3700_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3700E450g,
	"subcore_vm_fixed3750_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3750E450g,
	"subcore_vm_fixed3800_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3800E450g,
	"subcore_vm_fixed3825_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3825E450g,
	"subcore_vm_fixed3850_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3850E450g,
	"subcore_vm_fixed3875_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3875E450g,
	"subcore_vm_fixed3900_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3900E450g,
	"subcore_vm_fixed3975_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3975E450g,
	"subcore_vm_fixed4000_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4000E450g,
	"subcore_vm_fixed4025_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4025E450g,
	"subcore_vm_fixed4050_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4050E450g,
	"subcore_vm_fixed4100_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4100E450g,
	"subcore_vm_fixed4125_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4125E450g,
	"subcore_vm_fixed4200_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4200E450g,
	"subcore_vm_fixed4225_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4225E450g,
	"subcore_vm_fixed4250_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4250E450g,
	"subcore_vm_fixed4275_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4275E450g,
	"subcore_vm_fixed4300_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4300E450g,
	"subcore_vm_fixed4350_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4350E450g,
	"subcore_vm_fixed4375_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4375E450g,
	"subcore_vm_fixed4400_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4400E450g,
	"subcore_vm_fixed4425_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4425E450g,
	"subcore_vm_fixed4500_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4500E450g,
	"subcore_vm_fixed4550_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4550E450g,
	"subcore_vm_fixed4575_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4575E450g,
	"subcore_vm_fixed4600_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4600E450g,
	"subcore_vm_fixed4625_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4625E450g,
	"subcore_vm_fixed4650_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4650E450g,
	"subcore_vm_fixed4675_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4675E450g,
	"subcore_vm_fixed4700_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4700E450g,
	"subcore_vm_fixed4725_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4725E450g,
	"subcore_vm_fixed4750_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4750E450g,
	"subcore_vm_fixed4800_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4800E450g,
	"subcore_vm_fixed4875_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4875E450g,
	"subcore_vm_fixed4900_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4900E450g,
	"subcore_vm_fixed4950_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4950E450g,
	"subcore_vm_fixed5000_e4_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed5000E450g,
	"dynamic_e5_50g":                       UpdateVnicShapeDetailsVnicShapeDynamicE550g,
	"dynamic_e5_100g":                      UpdateVnicShapeDetailsVnicShapeDynamicE5100g,
	"entirehost_e5_50g":                    UpdateVnicShapeDetailsVnicShapeEntirehostE550g,
	"fixed0040_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0040E550g,
	"fixed0100_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0100E550g,
	"fixed0200_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0200E550g,
	"fixed0300_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0300E550g,
	"fixed0400_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0400E550g,
	"fixed0500_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0500E550g,
	"fixed0600_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0600E550g,
	"fixed0700_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0700E550g,
	"fixed0800_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0800E550g,
	"fixed0900_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0900E550g,
	"fixed1000_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1000E550g,
	"fixed1100_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1100E550g,
	"fixed1200_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1200E550g,
	"fixed1300_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1300E550g,
	"fixed1400_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1400E550g,
	"fixed1500_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1500E550g,
	"fixed1600_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1600E550g,
	"fixed1700_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1700E550g,
	"fixed1800_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1800E550g,
	"fixed1900_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1900E550g,
	"fixed2000_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2000E550g,
	"fixed2100_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2100E550g,
	"fixed2200_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2200E550g,
	"fixed2300_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2300E550g,
	"fixed2400_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2400E550g,
	"fixed2500_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2500E550g,
	"fixed2600_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2600E550g,
	"fixed2700_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2700E550g,
	"fixed2800_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2800E550g,
	"fixed2900_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2900E550g,
	"fixed3000_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3000E550g,
	"fixed3100_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3100E550g,
	"fixed3200_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3200E550g,
	"fixed3300_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3300E550g,
	"fixed3400_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3400E550g,
	"fixed3500_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3500E550g,
	"fixed3600_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3600E550g,
	"fixed3700_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3700E550g,
	"fixed3800_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3800E550g,
	"fixed3900_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3900E550g,
	"fixed4000_e5_50g":                     UpdateVnicShapeDetailsVnicShapeFixed4000E550g,
	"subcore_vm_fixed0025_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0025E550g,
	"subcore_vm_fixed0050_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0050E550g,
	"subcore_vm_fixed0075_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0075E550g,
	"subcore_vm_fixed0100_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0100E550g,
	"subcore_vm_fixed0125_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0125E550g,
	"subcore_vm_fixed0150_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0150E550g,
	"subcore_vm_fixed0175_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0175E550g,
	"subcore_vm_fixed0200_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0200E550g,
	"subcore_vm_fixed0225_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0225E550g,
	"subcore_vm_fixed0250_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0250E550g,
	"subcore_vm_fixed0275_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0275E550g,
	"subcore_vm_fixed0300_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0300E550g,
	"subcore_vm_fixed0325_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0325E550g,
	"subcore_vm_fixed0350_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0350E550g,
	"subcore_vm_fixed0375_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0375E550g,
	"subcore_vm_fixed0400_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0400E550g,
	"subcore_vm_fixed0425_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0425E550g,
	"subcore_vm_fixed0450_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0450E550g,
	"subcore_vm_fixed0475_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0475E550g,
	"subcore_vm_fixed0500_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0500E550g,
	"subcore_vm_fixed0525_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0525E550g,
	"subcore_vm_fixed0550_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0550E550g,
	"subcore_vm_fixed0575_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0575E550g,
	"subcore_vm_fixed0600_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0600E550g,
	"subcore_vm_fixed0625_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0625E550g,
	"subcore_vm_fixed0650_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0650E550g,
	"subcore_vm_fixed0675_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0675E550g,
	"subcore_vm_fixed0700_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0700E550g,
	"subcore_vm_fixed0725_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0725E550g,
	"subcore_vm_fixed0750_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0750E550g,
	"subcore_vm_fixed0775_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0775E550g,
	"subcore_vm_fixed0800_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0800E550g,
	"subcore_vm_fixed0825_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0825E550g,
	"subcore_vm_fixed0850_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0850E550g,
	"subcore_vm_fixed0875_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0875E550g,
	"subcore_vm_fixed0900_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0900E550g,
	"subcore_vm_fixed0925_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0925E550g,
	"subcore_vm_fixed0950_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0950E550g,
	"subcore_vm_fixed0975_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0975E550g,
	"subcore_vm_fixed1000_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1000E550g,
	"subcore_vm_fixed1025_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1025E550g,
	"subcore_vm_fixed1050_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1050E550g,
	"subcore_vm_fixed1075_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1075E550g,
	"subcore_vm_fixed1100_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1100E550g,
	"subcore_vm_fixed1125_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1125E550g,
	"subcore_vm_fixed1150_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1150E550g,
	"subcore_vm_fixed1175_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1175E550g,
	"subcore_vm_fixed1200_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1200E550g,
	"subcore_vm_fixed1225_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1225E550g,
	"subcore_vm_fixed1250_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1250E550g,
	"subcore_vm_fixed1275_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1275E550g,
	"subcore_vm_fixed1300_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1300E550g,
	"subcore_vm_fixed1325_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1325E550g,
	"subcore_vm_fixed1350_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1350E550g,
	"subcore_vm_fixed1375_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1375E550g,
	"subcore_vm_fixed1400_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1400E550g,
	"subcore_vm_fixed1425_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1425E550g,
	"subcore_vm_fixed1450_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1450E550g,
	"subcore_vm_fixed1475_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1475E550g,
	"subcore_vm_fixed1500_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1500E550g,
	"subcore_vm_fixed1525_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1525E550g,
	"subcore_vm_fixed1550_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1550E550g,
	"subcore_vm_fixed1575_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1575E550g,
	"subcore_vm_fixed1600_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1600E550g,
	"subcore_vm_fixed1625_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1625E550g,
	"subcore_vm_fixed1650_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1650E550g,
	"subcore_vm_fixed1675_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1675E550g,
	"subcore_vm_fixed1700_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1700E550g,
	"subcore_vm_fixed1725_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1725E550g,
	"subcore_vm_fixed1750_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1750E550g,
	"subcore_vm_fixed1775_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1775E550g,
	"subcore_vm_fixed1800_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1800E550g,
	"subcore_vm_fixed1825_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1825E550g,
	"subcore_vm_fixed1850_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1850E550g,
	"subcore_vm_fixed1875_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1875E550g,
	"subcore_vm_fixed1900_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1900E550g,
	"subcore_vm_fixed1925_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1925E550g,
	"subcore_vm_fixed1950_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1950E550g,
	"subcore_vm_fixed1975_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1975E550g,
	"subcore_vm_fixed2000_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2000E550g,
	"subcore_vm_fixed2025_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2025E550g,
	"subcore_vm_fixed2050_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2050E550g,
	"subcore_vm_fixed2075_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2075E550g,
	"subcore_vm_fixed2100_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2100E550g,
	"subcore_vm_fixed2125_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2125E550g,
	"subcore_vm_fixed2150_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2150E550g,
	"subcore_vm_fixed2175_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2175E550g,
	"subcore_vm_fixed2200_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2200E550g,
	"subcore_vm_fixed2225_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2225E550g,
	"subcore_vm_fixed2250_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2250E550g,
	"subcore_vm_fixed2275_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2275E550g,
	"subcore_vm_fixed2300_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2300E550g,
	"subcore_vm_fixed2325_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2325E550g,
	"subcore_vm_fixed2350_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2350E550g,
	"subcore_vm_fixed2375_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2375E550g,
	"subcore_vm_fixed2400_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2400E550g,
	"subcore_vm_fixed2425_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2425E550g,
	"subcore_vm_fixed2450_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2450E550g,
	"subcore_vm_fixed2475_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2475E550g,
	"subcore_vm_fixed2500_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2500E550g,
	"subcore_vm_fixed2525_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2525E550g,
	"subcore_vm_fixed2550_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2550E550g,
	"subcore_vm_fixed2575_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2575E550g,
	"subcore_vm_fixed2600_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2600E550g,
	"subcore_vm_fixed2625_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2625E550g,
	"subcore_vm_fixed2650_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2650E550g,
	"subcore_vm_fixed2675_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2675E550g,
	"subcore_vm_fixed2700_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2700E550g,
	"subcore_vm_fixed2725_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2725E550g,
	"subcore_vm_fixed2750_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2750E550g,
	"subcore_vm_fixed2775_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2775E550g,
	"subcore_vm_fixed2800_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2800E550g,
	"subcore_vm_fixed2825_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2825E550g,
	"subcore_vm_fixed2850_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2850E550g,
	"subcore_vm_fixed2875_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2875E550g,
	"subcore_vm_fixed2900_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2900E550g,
	"subcore_vm_fixed2925_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2925E550g,
	"subcore_vm_fixed2950_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2950E550g,
	"subcore_vm_fixed2975_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2975E550g,
	"subcore_vm_fixed3000_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3000E550g,
	"subcore_vm_fixed3025_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3025E550g,
	"subcore_vm_fixed3050_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3050E550g,
	"subcore_vm_fixed3075_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3075E550g,
	"subcore_vm_fixed3100_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3100E550g,
	"subcore_vm_fixed3125_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3125E550g,
	"subcore_vm_fixed3150_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3150E550g,
	"subcore_vm_fixed3175_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3175E550g,
	"subcore_vm_fixed3200_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3200E550g,
	"subcore_vm_fixed3225_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3225E550g,
	"subcore_vm_fixed3250_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3250E550g,
	"subcore_vm_fixed3275_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3275E550g,
	"subcore_vm_fixed3300_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3300E550g,
	"subcore_vm_fixed3325_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3325E550g,
	"subcore_vm_fixed3350_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3350E550g,
	"subcore_vm_fixed3375_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3375E550g,
	"subcore_vm_fixed3400_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3400E550g,
	"subcore_vm_fixed3425_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3425E550g,
	"subcore_vm_fixed3450_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3450E550g,
	"subcore_vm_fixed3475_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3475E550g,
	"subcore_vm_fixed3500_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3500E550g,
	"subcore_vm_fixed3525_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3525E550g,
	"subcore_vm_fixed3550_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3550E550g,
	"subcore_vm_fixed3575_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3575E550g,
	"subcore_vm_fixed3600_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3600E550g,
	"subcore_vm_fixed3625_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3625E550g,
	"subcore_vm_fixed3650_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3650E550g,
	"subcore_vm_fixed3675_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3675E550g,
	"subcore_vm_fixed3700_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3700E550g,
	"subcore_vm_fixed3725_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3725E550g,
	"subcore_vm_fixed3750_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3750E550g,
	"subcore_vm_fixed3775_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3775E550g,
	"subcore_vm_fixed3800_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3800E550g,
	"subcore_vm_fixed3825_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3825E550g,
	"subcore_vm_fixed3850_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3850E550g,
	"subcore_vm_fixed3875_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3875E550g,
	"subcore_vm_fixed3900_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3900E550g,
	"subcore_vm_fixed3925_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3925E550g,
	"subcore_vm_fixed3950_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3950E550g,
	"subcore_vm_fixed3975_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3975E550g,
	"subcore_vm_fixed4000_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4000E550g,
	"subcore_vm_fixed4025_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4025E550g,
	"subcore_vm_fixed4050_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4050E550g,
	"subcore_vm_fixed4075_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4075E550g,
	"subcore_vm_fixed4100_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4100E550g,
	"subcore_vm_fixed4125_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4125E550g,
	"subcore_vm_fixed4150_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4150E550g,
	"subcore_vm_fixed4175_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4175E550g,
	"subcore_vm_fixed4200_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4200E550g,
	"subcore_vm_fixed4225_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4225E550g,
	"subcore_vm_fixed4250_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4250E550g,
	"subcore_vm_fixed4275_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4275E550g,
	"subcore_vm_fixed4300_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4300E550g,
	"subcore_vm_fixed4325_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4325E550g,
	"subcore_vm_fixed4350_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4350E550g,
	"subcore_vm_fixed4375_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4375E550g,
	"subcore_vm_fixed4400_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4400E550g,
	"subcore_vm_fixed4425_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4425E550g,
	"subcore_vm_fixed4450_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4450E550g,
	"subcore_vm_fixed4475_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4475E550g,
	"subcore_vm_fixed4500_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4500E550g,
	"subcore_vm_fixed4525_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4525E550g,
	"subcore_vm_fixed4550_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4550E550g,
	"subcore_vm_fixed4575_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4575E550g,
	"subcore_vm_fixed4600_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4600E550g,
	"subcore_vm_fixed4625_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4625E550g,
	"subcore_vm_fixed4650_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4650E550g,
	"subcore_vm_fixed4675_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4675E550g,
	"subcore_vm_fixed4700_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4700E550g,
	"subcore_vm_fixed4725_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4725E550g,
	"subcore_vm_fixed4750_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4750E550g,
	"subcore_vm_fixed4775_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4775E550g,
	"subcore_vm_fixed4800_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4800E550g,
	"subcore_vm_fixed4825_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4825E550g,
	"subcore_vm_fixed4850_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4850E550g,
	"subcore_vm_fixed4875_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4875E550g,
	"subcore_vm_fixed4900_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4900E550g,
	"subcore_vm_fixed4925_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4925E550g,
	"subcore_vm_fixed4950_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4950E550g,
	"subcore_vm_fixed4975_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4975E550g,
	"subcore_vm_fixed5000_e5_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed5000E550g,
	"subcore_vm_fixed0020_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0020A150g,
	"subcore_vm_fixed0040_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0040A150g,
	"subcore_vm_fixed0060_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0060A150g,
	"subcore_vm_fixed0080_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0080A150g,
	"subcore_vm_fixed0100_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0100A150g,
	"subcore_vm_fixed0120_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0120A150g,
	"subcore_vm_fixed0140_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0140A150g,
	"subcore_vm_fixed0160_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0160A150g,
	"subcore_vm_fixed0180_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0180A150g,
	"subcore_vm_fixed0200_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0200A150g,
	"subcore_vm_fixed0220_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0220A150g,
	"subcore_vm_fixed0240_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0240A150g,
	"subcore_vm_fixed0260_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0260A150g,
	"subcore_vm_fixed0280_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0280A150g,
	"subcore_vm_fixed0300_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0300A150g,
	"subcore_vm_fixed0320_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0320A150g,
	"subcore_vm_fixed0340_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0340A150g,
	"subcore_vm_fixed0360_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0360A150g,
	"subcore_vm_fixed0380_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0380A150g,
	"subcore_vm_fixed0400_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0400A150g,
	"subcore_vm_fixed0420_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0420A150g,
	"subcore_vm_fixed0440_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0440A150g,
	"subcore_vm_fixed0460_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0460A150g,
	"subcore_vm_fixed0480_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0480A150g,
	"subcore_vm_fixed0500_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0500A150g,
	"subcore_vm_fixed0520_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0520A150g,
	"subcore_vm_fixed0540_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0540A150g,
	"subcore_vm_fixed0560_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0560A150g,
	"subcore_vm_fixed0580_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0580A150g,
	"subcore_vm_fixed0600_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0600A150g,
	"subcore_vm_fixed0620_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0620A150g,
	"subcore_vm_fixed0640_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0640A150g,
	"subcore_vm_fixed0660_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0660A150g,
	"subcore_vm_fixed0680_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0680A150g,
	"subcore_vm_fixed0700_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0700A150g,
	"subcore_vm_fixed0720_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0720A150g,
	"subcore_vm_fixed0740_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0740A150g,
	"subcore_vm_fixed0760_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0760A150g,
	"subcore_vm_fixed0780_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0780A150g,
	"subcore_vm_fixed0800_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0800A150g,
	"subcore_vm_fixed0820_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0820A150g,
	"subcore_vm_fixed0840_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0840A150g,
	"subcore_vm_fixed0860_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0860A150g,
	"subcore_vm_fixed0880_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0880A150g,
	"subcore_vm_fixed0900_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0900A150g,
	"subcore_vm_fixed0920_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0920A150g,
	"subcore_vm_fixed0940_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0940A150g,
	"subcore_vm_fixed0960_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0960A150g,
	"subcore_vm_fixed0980_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0980A150g,
	"subcore_vm_fixed1000_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1000A150g,
	"subcore_vm_fixed1020_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1020A150g,
	"subcore_vm_fixed1040_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1040A150g,
	"subcore_vm_fixed1060_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1060A150g,
	"subcore_vm_fixed1080_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1080A150g,
	"subcore_vm_fixed1100_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1100A150g,
	"subcore_vm_fixed1120_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1120A150g,
	"subcore_vm_fixed1140_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1140A150g,
	"subcore_vm_fixed1160_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1160A150g,
	"subcore_vm_fixed1180_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1180A150g,
	"subcore_vm_fixed1200_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1200A150g,
	"subcore_vm_fixed1220_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1220A150g,
	"subcore_vm_fixed1240_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1240A150g,
	"subcore_vm_fixed1260_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1260A150g,
	"subcore_vm_fixed1280_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1280A150g,
	"subcore_vm_fixed1300_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1300A150g,
	"subcore_vm_fixed1320_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1320A150g,
	"subcore_vm_fixed1340_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1340A150g,
	"subcore_vm_fixed1360_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1360A150g,
	"subcore_vm_fixed1380_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1380A150g,
	"subcore_vm_fixed1400_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1400A150g,
	"subcore_vm_fixed1420_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1420A150g,
	"subcore_vm_fixed1440_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1440A150g,
	"subcore_vm_fixed1460_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1460A150g,
	"subcore_vm_fixed1480_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1480A150g,
	"subcore_vm_fixed1500_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1500A150g,
	"subcore_vm_fixed1520_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1520A150g,
	"subcore_vm_fixed1540_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1540A150g,
	"subcore_vm_fixed1560_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1560A150g,
	"subcore_vm_fixed1580_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1580A150g,
	"subcore_vm_fixed1600_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1600A150g,
	"subcore_vm_fixed1620_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1620A150g,
	"subcore_vm_fixed1640_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1640A150g,
	"subcore_vm_fixed1660_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1660A150g,
	"subcore_vm_fixed1680_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1680A150g,
	"subcore_vm_fixed1700_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1700A150g,
	"subcore_vm_fixed1720_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1720A150g,
	"subcore_vm_fixed1740_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1740A150g,
	"subcore_vm_fixed1760_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1760A150g,
	"subcore_vm_fixed1780_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1780A150g,
	"subcore_vm_fixed1800_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1800A150g,
	"subcore_vm_fixed1820_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1820A150g,
	"subcore_vm_fixed1840_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1840A150g,
	"subcore_vm_fixed1860_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1860A150g,
	"subcore_vm_fixed1880_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1880A150g,
	"subcore_vm_fixed1900_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1900A150g,
	"subcore_vm_fixed1920_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1920A150g,
	"subcore_vm_fixed1940_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1940A150g,
	"subcore_vm_fixed1960_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1960A150g,
	"subcore_vm_fixed1980_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1980A150g,
	"subcore_vm_fixed2000_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2000A150g,
	"subcore_vm_fixed2020_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2020A150g,
	"subcore_vm_fixed2040_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2040A150g,
	"subcore_vm_fixed2060_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2060A150g,
	"subcore_vm_fixed2080_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2080A150g,
	"subcore_vm_fixed2100_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2100A150g,
	"subcore_vm_fixed2120_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2120A150g,
	"subcore_vm_fixed2140_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2140A150g,
	"subcore_vm_fixed2160_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2160A150g,
	"subcore_vm_fixed2180_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2180A150g,
	"subcore_vm_fixed2200_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2200A150g,
	"subcore_vm_fixed2220_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2220A150g,
	"subcore_vm_fixed2240_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2240A150g,
	"subcore_vm_fixed2260_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2260A150g,
	"subcore_vm_fixed2280_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2280A150g,
	"subcore_vm_fixed2300_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2300A150g,
	"subcore_vm_fixed2320_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2320A150g,
	"subcore_vm_fixed2340_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2340A150g,
	"subcore_vm_fixed2360_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2360A150g,
	"subcore_vm_fixed2380_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2380A150g,
	"subcore_vm_fixed2400_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2400A150g,
	"subcore_vm_fixed2420_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2420A150g,
	"subcore_vm_fixed2440_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2440A150g,
	"subcore_vm_fixed2460_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2460A150g,
	"subcore_vm_fixed2480_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2480A150g,
	"subcore_vm_fixed2500_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2500A150g,
	"subcore_vm_fixed2520_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2520A150g,
	"subcore_vm_fixed2540_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2540A150g,
	"subcore_vm_fixed2560_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2560A150g,
	"subcore_vm_fixed2580_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2580A150g,
	"subcore_vm_fixed2600_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2600A150g,
	"subcore_vm_fixed2620_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2620A150g,
	"subcore_vm_fixed2640_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2640A150g,
	"subcore_vm_fixed2660_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2660A150g,
	"subcore_vm_fixed2680_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2680A150g,
	"subcore_vm_fixed2700_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2700A150g,
	"subcore_vm_fixed2720_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2720A150g,
	"subcore_vm_fixed2740_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2740A150g,
	"subcore_vm_fixed2760_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2760A150g,
	"subcore_vm_fixed2780_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2780A150g,
	"subcore_vm_fixed2800_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2800A150g,
	"subcore_vm_fixed2820_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2820A150g,
	"subcore_vm_fixed2840_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2840A150g,
	"subcore_vm_fixed2860_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2860A150g,
	"subcore_vm_fixed2880_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2880A150g,
	"subcore_vm_fixed2900_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2900A150g,
	"subcore_vm_fixed2920_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2920A150g,
	"subcore_vm_fixed2940_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2940A150g,
	"subcore_vm_fixed2960_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2960A150g,
	"subcore_vm_fixed2980_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2980A150g,
	"subcore_vm_fixed3000_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3000A150g,
	"subcore_vm_fixed3020_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3020A150g,
	"subcore_vm_fixed3040_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3040A150g,
	"subcore_vm_fixed3060_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3060A150g,
	"subcore_vm_fixed3080_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3080A150g,
	"subcore_vm_fixed3100_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3100A150g,
	"subcore_vm_fixed3120_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3120A150g,
	"subcore_vm_fixed3140_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3140A150g,
	"subcore_vm_fixed3160_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3160A150g,
	"subcore_vm_fixed3180_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3180A150g,
	"subcore_vm_fixed3200_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3200A150g,
	"subcore_vm_fixed3220_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3220A150g,
	"subcore_vm_fixed3240_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3240A150g,
	"subcore_vm_fixed3260_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3260A150g,
	"subcore_vm_fixed3280_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3280A150g,
	"subcore_vm_fixed3300_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3300A150g,
	"subcore_vm_fixed3320_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3320A150g,
	"subcore_vm_fixed3340_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3340A150g,
	"subcore_vm_fixed3360_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3360A150g,
	"subcore_vm_fixed3380_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3380A150g,
	"subcore_vm_fixed3400_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3400A150g,
	"subcore_vm_fixed3420_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3420A150g,
	"subcore_vm_fixed3440_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3440A150g,
	"subcore_vm_fixed3460_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3460A150g,
	"subcore_vm_fixed3480_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3480A150g,
	"subcore_vm_fixed3500_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3500A150g,
	"subcore_vm_fixed3520_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3520A150g,
	"subcore_vm_fixed3540_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3540A150g,
	"subcore_vm_fixed3560_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3560A150g,
	"subcore_vm_fixed3580_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3580A150g,
	"subcore_vm_fixed3600_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3600A150g,
	"subcore_vm_fixed3620_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3620A150g,
	"subcore_vm_fixed3640_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3640A150g,
	"subcore_vm_fixed3660_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3660A150g,
	"subcore_vm_fixed3680_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3680A150g,
	"subcore_vm_fixed3700_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3700A150g,
	"subcore_vm_fixed3720_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3720A150g,
	"subcore_vm_fixed3740_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3740A150g,
	"subcore_vm_fixed3760_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3760A150g,
	"subcore_vm_fixed3780_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3780A150g,
	"subcore_vm_fixed3800_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3800A150g,
	"subcore_vm_fixed3820_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3820A150g,
	"subcore_vm_fixed3840_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3840A150g,
	"subcore_vm_fixed3860_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3860A150g,
	"subcore_vm_fixed3880_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3880A150g,
	"subcore_vm_fixed3900_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3900A150g,
	"subcore_vm_fixed3920_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3920A150g,
	"subcore_vm_fixed3940_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3940A150g,
	"subcore_vm_fixed3960_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3960A150g,
	"subcore_vm_fixed3980_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3980A150g,
	"subcore_vm_fixed4000_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4000A150g,
	"subcore_vm_fixed4020_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4020A150g,
	"subcore_vm_fixed4040_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4040A150g,
	"subcore_vm_fixed4060_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4060A150g,
	"subcore_vm_fixed4080_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4080A150g,
	"subcore_vm_fixed4100_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4100A150g,
	"subcore_vm_fixed4120_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4120A150g,
	"subcore_vm_fixed4140_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4140A150g,
	"subcore_vm_fixed4160_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4160A150g,
	"subcore_vm_fixed4180_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4180A150g,
	"subcore_vm_fixed4200_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4200A150g,
	"subcore_vm_fixed4220_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4220A150g,
	"subcore_vm_fixed4240_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4240A150g,
	"subcore_vm_fixed4260_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4260A150g,
	"subcore_vm_fixed4280_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4280A150g,
	"subcore_vm_fixed4300_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4300A150g,
	"subcore_vm_fixed4320_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4320A150g,
	"subcore_vm_fixed4340_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4340A150g,
	"subcore_vm_fixed4360_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4360A150g,
	"subcore_vm_fixed4380_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4380A150g,
	"subcore_vm_fixed4400_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4400A150g,
	"subcore_vm_fixed4420_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4420A150g,
	"subcore_vm_fixed4440_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4440A150g,
	"subcore_vm_fixed4460_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4460A150g,
	"subcore_vm_fixed4480_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4480A150g,
	"subcore_vm_fixed4500_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4500A150g,
	"subcore_vm_fixed4520_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4520A150g,
	"subcore_vm_fixed4540_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4540A150g,
	"subcore_vm_fixed4560_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4560A150g,
	"subcore_vm_fixed4580_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4580A150g,
	"subcore_vm_fixed4600_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4600A150g,
	"subcore_vm_fixed4620_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4620A150g,
	"subcore_vm_fixed4640_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4640A150g,
	"subcore_vm_fixed4660_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4660A150g,
	"subcore_vm_fixed4680_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4680A150g,
	"subcore_vm_fixed4700_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4700A150g,
	"subcore_vm_fixed4720_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4720A150g,
	"subcore_vm_fixed4740_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4740A150g,
	"subcore_vm_fixed4760_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4760A150g,
	"subcore_vm_fixed4780_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4780A150g,
	"subcore_vm_fixed4800_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4800A150g,
	"subcore_vm_fixed4820_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4820A150g,
	"subcore_vm_fixed4840_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4840A150g,
	"subcore_vm_fixed4860_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4860A150g,
	"subcore_vm_fixed4880_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4880A150g,
	"subcore_vm_fixed4900_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4900A150g,
	"subcore_vm_fixed4920_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4920A150g,
	"subcore_vm_fixed4940_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4940A150g,
	"subcore_vm_fixed4960_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4960A150g,
	"subcore_vm_fixed4980_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4980A150g,
	"subcore_vm_fixed5000_a1_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed5000A150g,
	"subcore_vm_fixed0090_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0090X950g,
	"subcore_vm_fixed0180_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0180X950g,
	"subcore_vm_fixed0270_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0270X950g,
	"subcore_vm_fixed0360_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0360X950g,
	"subcore_vm_fixed0450_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0450X950g,
	"subcore_vm_fixed0540_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0540X950g,
	"subcore_vm_fixed0630_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0630X950g,
	"subcore_vm_fixed0720_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0720X950g,
	"subcore_vm_fixed0810_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0810X950g,
	"subcore_vm_fixed0900_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0900X950g,
	"subcore_vm_fixed0990_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed0990X950g,
	"subcore_vm_fixed1080_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1080X950g,
	"subcore_vm_fixed1170_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1170X950g,
	"subcore_vm_fixed1260_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1260X950g,
	"subcore_vm_fixed1350_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1350X950g,
	"subcore_vm_fixed1440_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1440X950g,
	"subcore_vm_fixed1530_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1530X950g,
	"subcore_vm_fixed1620_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1620X950g,
	"subcore_vm_fixed1710_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1710X950g,
	"subcore_vm_fixed1800_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1800X950g,
	"subcore_vm_fixed1890_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1890X950g,
	"subcore_vm_fixed1980_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed1980X950g,
	"subcore_vm_fixed2070_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2070X950g,
	"subcore_vm_fixed2160_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2160X950g,
	"subcore_vm_fixed2250_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2250X950g,
	"subcore_vm_fixed2340_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2340X950g,
	"subcore_vm_fixed2430_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2430X950g,
	"subcore_vm_fixed2520_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2520X950g,
	"subcore_vm_fixed2610_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2610X950g,
	"subcore_vm_fixed2700_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2700X950g,
	"subcore_vm_fixed2790_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2790X950g,
	"subcore_vm_fixed2880_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2880X950g,
	"subcore_vm_fixed2970_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed2970X950g,
	"subcore_vm_fixed3060_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3060X950g,
	"subcore_vm_fixed3150_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3150X950g,
	"subcore_vm_fixed3240_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3240X950g,
	"subcore_vm_fixed3330_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3330X950g,
	"subcore_vm_fixed3420_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3420X950g,
	"subcore_vm_fixed3510_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3510X950g,
	"subcore_vm_fixed3600_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3600X950g,
	"subcore_vm_fixed3690_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3690X950g,
	"subcore_vm_fixed3780_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3780X950g,
	"subcore_vm_fixed3870_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3870X950g,
	"subcore_vm_fixed3960_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed3960X950g,
	"subcore_vm_fixed4050_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4050X950g,
	"subcore_vm_fixed4140_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4140X950g,
	"subcore_vm_fixed4230_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4230X950g,
	"subcore_vm_fixed4320_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4320X950g,
	"subcore_vm_fixed4410_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4410X950g,
	"subcore_vm_fixed4500_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4500X950g,
	"subcore_vm_fixed4590_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4590X950g,
	"subcore_vm_fixed4680_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4680X950g,
	"subcore_vm_fixed4770_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4770X950g,
	"subcore_vm_fixed4860_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4860X950g,
	"subcore_vm_fixed4950_x9_50g":          UpdateVnicShapeDetailsVnicShapeSubcoreVmFixed4950X950g,
	"dynamic_a1_50g":                       UpdateVnicShapeDetailsVnicShapeDynamicA150g,
	"fixed0040_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0040A150g,
	"fixed0100_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0100A150g,
	"fixed0200_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0200A150g,
	"fixed0300_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0300A150g,
	"fixed0400_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0400A150g,
	"fixed0500_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0500A150g,
	"fixed0600_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0600A150g,
	"fixed0700_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0700A150g,
	"fixed0800_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0800A150g,
	"fixed0900_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0900A150g,
	"fixed1000_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1000A150g,
	"fixed1100_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1100A150g,
	"fixed1200_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1200A150g,
	"fixed1300_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1300A150g,
	"fixed1400_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1400A150g,
	"fixed1500_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1500A150g,
	"fixed1600_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1600A150g,
	"fixed1700_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1700A150g,
	"fixed1800_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1800A150g,
	"fixed1900_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1900A150g,
	"fixed2000_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2000A150g,
	"fixed2100_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2100A150g,
	"fixed2200_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2200A150g,
	"fixed2300_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2300A150g,
	"fixed2400_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2400A150g,
	"fixed2500_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2500A150g,
	"fixed2600_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2600A150g,
	"fixed2700_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2700A150g,
	"fixed2800_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2800A150g,
	"fixed2900_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2900A150g,
	"fixed3000_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3000A150g,
	"fixed3100_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3100A150g,
	"fixed3200_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3200A150g,
	"fixed3300_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3300A150g,
	"fixed3400_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3400A150g,
	"fixed3500_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3500A150g,
	"fixed3600_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3600A150g,
	"fixed3700_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3700A150g,
	"fixed3800_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3800A150g,
	"fixed3900_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3900A150g,
	"fixed4000_a1_50g":                     UpdateVnicShapeDetailsVnicShapeFixed4000A150g,
	"fixed5000_telesis_a1_50g":             UpdateVnicShapeDetailsVnicShapeFixed5000TelesisA150g,
	"entirehost_a1_50g":                    UpdateVnicShapeDetailsVnicShapeEntirehostA150g,
	"dynamic_x7_25g":                       UpdateVnicShapeDetailsVnicShapeDynamicX725g,
	"fixed0040_x7_25g":                     UpdateVnicShapeDetailsVnicShapeFixed0040X725g,
	"standard_vm_fixed0100_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0100X725g,
	"standard_vm_fixed0200_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0200X725g,
	"standard_vm_fixed0300_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0300X725g,
	"standard_vm_fixed0400_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0400X725g,
	"standard_vm_fixed0500_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0500X725g,
	"standard_vm_fixed0600_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0600X725g,
	"standard_vm_fixed0700_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0700X725g,
	"standard_vm_fixed0800_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0800X725g,
	"standard_vm_fixed0900_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0900X725g,
	"standard_vm_fixed1000_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1000X725g,
	"standard_vm_fixed1100_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1100X725g,
	"standard_vm_fixed1200_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1200X725g,
	"standard_vm_fixed1300_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1300X725g,
	"standard_vm_fixed1400_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1400X725g,
	"standard_vm_fixed1500_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1500X725g,
	"standard_vm_fixed1600_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1600X725g,
	"standard_vm_fixed1700_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1700X725g,
	"standard_vm_fixed1800_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1800X725g,
	"standard_vm_fixed1900_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1900X725g,
	"standard_vm_fixed2000_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2000X725g,
	"standard_vm_fixed2100_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2100X725g,
	"standard_vm_fixed2200_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2200X725g,
	"standard_vm_fixed2300_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2300X725g,
	"standard_vm_fixed2400_x7_25g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2400X725g,
	"subcore_standard_vm_fixed0025_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0025X725g,
	"subcore_standard_vm_fixed0050_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0050X725g,
	"subcore_standard_vm_fixed0075_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0075X725g,
	"subcore_standard_vm_fixed0100_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0100X725g,
	"subcore_standard_vm_fixed0125_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0125X725g,
	"subcore_standard_vm_fixed0150_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0150X725g,
	"subcore_standard_vm_fixed0175_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0175X725g,
	"subcore_standard_vm_fixed0200_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0200X725g,
	"subcore_standard_vm_fixed0225_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0225X725g,
	"subcore_standard_vm_fixed0250_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0250X725g,
	"subcore_standard_vm_fixed0275_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0275X725g,
	"subcore_standard_vm_fixed0300_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0300X725g,
	"subcore_standard_vm_fixed0325_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0325X725g,
	"subcore_standard_vm_fixed0350_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0350X725g,
	"subcore_standard_vm_fixed0375_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0375X725g,
	"subcore_standard_vm_fixed0400_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0400X725g,
	"subcore_standard_vm_fixed0425_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0425X725g,
	"subcore_standard_vm_fixed0450_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0450X725g,
	"subcore_standard_vm_fixed0475_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0475X725g,
	"subcore_standard_vm_fixed0500_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0500X725g,
	"subcore_standard_vm_fixed0525_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0525X725g,
	"subcore_standard_vm_fixed0550_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0550X725g,
	"subcore_standard_vm_fixed0575_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0575X725g,
	"subcore_standard_vm_fixed0600_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0600X725g,
	"subcore_standard_vm_fixed0625_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0625X725g,
	"subcore_standard_vm_fixed0650_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0650X725g,
	"subcore_standard_vm_fixed0675_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0675X725g,
	"subcore_standard_vm_fixed0700_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0700X725g,
	"subcore_standard_vm_fixed0725_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0725X725g,
	"subcore_standard_vm_fixed0750_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0750X725g,
	"subcore_standard_vm_fixed0775_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0775X725g,
	"subcore_standard_vm_fixed0800_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0800X725g,
	"subcore_standard_vm_fixed0825_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0825X725g,
	"subcore_standard_vm_fixed0850_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0850X725g,
	"subcore_standard_vm_fixed0875_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0875X725g,
	"subcore_standard_vm_fixed0900_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0900X725g,
	"subcore_standard_vm_fixed0925_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0925X725g,
	"subcore_standard_vm_fixed0950_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0950X725g,
	"subcore_standard_vm_fixed0975_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0975X725g,
	"subcore_standard_vm_fixed1000_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1000X725g,
	"subcore_standard_vm_fixed1025_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1025X725g,
	"subcore_standard_vm_fixed1050_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1050X725g,
	"subcore_standard_vm_fixed1075_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1075X725g,
	"subcore_standard_vm_fixed1100_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1100X725g,
	"subcore_standard_vm_fixed1125_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1125X725g,
	"subcore_standard_vm_fixed1150_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1150X725g,
	"subcore_standard_vm_fixed1175_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1175X725g,
	"subcore_standard_vm_fixed1200_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1200X725g,
	"subcore_standard_vm_fixed1225_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1225X725g,
	"subcore_standard_vm_fixed1250_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1250X725g,
	"subcore_standard_vm_fixed1275_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1275X725g,
	"subcore_standard_vm_fixed1300_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1300X725g,
	"subcore_standard_vm_fixed1325_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1325X725g,
	"subcore_standard_vm_fixed1350_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1350X725g,
	"subcore_standard_vm_fixed1375_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1375X725g,
	"subcore_standard_vm_fixed1400_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1400X725g,
	"subcore_standard_vm_fixed1425_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1425X725g,
	"subcore_standard_vm_fixed1450_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1450X725g,
	"subcore_standard_vm_fixed1475_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1475X725g,
	"subcore_standard_vm_fixed1500_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1500X725g,
	"subcore_standard_vm_fixed1525_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1525X725g,
	"subcore_standard_vm_fixed1550_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1550X725g,
	"subcore_standard_vm_fixed1575_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1575X725g,
	"subcore_standard_vm_fixed1600_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1600X725g,
	"subcore_standard_vm_fixed1625_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1625X725g,
	"subcore_standard_vm_fixed1650_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1650X725g,
	"subcore_standard_vm_fixed1675_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1675X725g,
	"subcore_standard_vm_fixed1700_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1700X725g,
	"subcore_standard_vm_fixed1725_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1725X725g,
	"subcore_standard_vm_fixed1750_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1750X725g,
	"subcore_standard_vm_fixed1775_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1775X725g,
	"subcore_standard_vm_fixed1800_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1800X725g,
	"subcore_standard_vm_fixed1825_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1825X725g,
	"subcore_standard_vm_fixed1850_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1850X725g,
	"subcore_standard_vm_fixed1875_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1875X725g,
	"subcore_standard_vm_fixed1900_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1900X725g,
	"subcore_standard_vm_fixed1925_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1925X725g,
	"subcore_standard_vm_fixed1950_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1950X725g,
	"subcore_standard_vm_fixed1975_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1975X725g,
	"subcore_standard_vm_fixed2000_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2000X725g,
	"subcore_standard_vm_fixed2025_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2025X725g,
	"subcore_standard_vm_fixed2050_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2050X725g,
	"subcore_standard_vm_fixed2075_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2075X725g,
	"subcore_standard_vm_fixed2100_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2100X725g,
	"subcore_standard_vm_fixed2125_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2125X725g,
	"subcore_standard_vm_fixed2150_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2150X725g,
	"subcore_standard_vm_fixed2175_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2175X725g,
	"subcore_standard_vm_fixed2200_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2200X725g,
	"subcore_standard_vm_fixed2225_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2225X725g,
	"subcore_standard_vm_fixed2250_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2250X725g,
	"subcore_standard_vm_fixed2275_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2275X725g,
	"subcore_standard_vm_fixed2300_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2300X725g,
	"subcore_standard_vm_fixed2325_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2325X725g,
	"subcore_standard_vm_fixed2350_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2350X725g,
	"subcore_standard_vm_fixed2375_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2375X725g,
	"subcore_standard_vm_fixed2400_x7_25g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2400X725g,
	"entirehost_x7_25g":                    UpdateVnicShapeDetailsVnicShapeEntirehostX725g,
	"dynamic_x9_50g":                       UpdateVnicShapeDetailsVnicShapeDynamicX950g,
	"fixed0040_x9_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0040X950g,
	"fixed0400_x9_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0400X950g,
	"fixed0800_x9_50g":                     UpdateVnicShapeDetailsVnicShapeFixed0800X950g,
	"fixed1200_x9_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1200X950g,
	"fixed1600_x9_50g":                     UpdateVnicShapeDetailsVnicShapeFixed1600X950g,
	"fixed2000_x9_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2000X950g,
	"fixed2400_x9_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2400X950g,
	"fixed2800_x9_50g":                     UpdateVnicShapeDetailsVnicShapeFixed2800X950g,
	"fixed3200_x9_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3200X950g,
	"fixed3600_x9_50g":                     UpdateVnicShapeDetailsVnicShapeFixed3600X950g,
	"fixed4000_x9_50g":                     UpdateVnicShapeDetailsVnicShapeFixed4000X950g,
	"standard_vm_fixed0100_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0100X950g,
	"standard_vm_fixed0200_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0200X950g,
	"standard_vm_fixed0300_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0300X950g,
	"standard_vm_fixed0400_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0400X950g,
	"standard_vm_fixed0500_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0500X950g,
	"standard_vm_fixed0600_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0600X950g,
	"standard_vm_fixed0700_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0700X950g,
	"standard_vm_fixed0800_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0800X950g,
	"standard_vm_fixed0900_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed0900X950g,
	"standard_vm_fixed1000_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1000X950g,
	"standard_vm_fixed1100_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1100X950g,
	"standard_vm_fixed1200_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1200X950g,
	"standard_vm_fixed1300_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1300X950g,
	"standard_vm_fixed1400_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1400X950g,
	"standard_vm_fixed1500_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1500X950g,
	"standard_vm_fixed1600_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1600X950g,
	"standard_vm_fixed1700_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1700X950g,
	"standard_vm_fixed1800_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1800X950g,
	"standard_vm_fixed1900_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed1900X950g,
	"standard_vm_fixed2000_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2000X950g,
	"standard_vm_fixed2100_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2100X950g,
	"standard_vm_fixed2200_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2200X950g,
	"standard_vm_fixed2300_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2300X950g,
	"standard_vm_fixed2400_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2400X950g,
	"standard_vm_fixed2500_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2500X950g,
	"standard_vm_fixed2600_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2600X950g,
	"standard_vm_fixed2700_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2700X950g,
	"standard_vm_fixed2800_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2800X950g,
	"standard_vm_fixed2900_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed2900X950g,
	"standard_vm_fixed3000_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed3000X950g,
	"standard_vm_fixed3100_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed3100X950g,
	"standard_vm_fixed3200_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed3200X950g,
	"standard_vm_fixed3300_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed3300X950g,
	"standard_vm_fixed3400_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed3400X950g,
	"standard_vm_fixed3500_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed3500X950g,
	"standard_vm_fixed3600_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed3600X950g,
	"standard_vm_fixed3700_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed3700X950g,
	"standard_vm_fixed3800_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed3800X950g,
	"standard_vm_fixed3900_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed3900X950g,
	"standard_vm_fixed4000_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed4000X950g,
	"standard_vm_fixed4100_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed4100X950g,
	"standard_vm_fixed4200_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed4200X950g,
	"standard_vm_fixed4300_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed4300X950g,
	"standard_vm_fixed4400_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed4400X950g,
	"standard_vm_fixed4500_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed4500X950g,
	"standard_vm_fixed4600_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed4600X950g,
	"standard_vm_fixed4700_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed4700X950g,
	"standard_vm_fixed4800_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed4800X950g,
	"standard_vm_fixed4900_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed4900X950g,
	"standard_vm_fixed5000_x9_50g":         UpdateVnicShapeDetailsVnicShapeStandardVmFixed5000X950g,
	"a10_gpu_vm_fixed2400_x9_50g":          UpdateVnicShapeDetailsVnicShapeA10GpuVmFixed2400X950g,
	"a10_gpu_vm_fixed4800_x9_50g":          UpdateVnicShapeDetailsVnicShapeA10GpuVmFixed4800X950g,
	"vnuma_vm_fixed0100_1_x9_50g":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed01001X950g,
	"vnuma_vm_fixed0200_2_x9_50g":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed02002X950g,
	"vnuma_vm_fixed0300_3_x9_50g":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed03003X950g,
	"vnuma_vm_fixed0400_4_x9_50g":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed04004X950g,
	"vnuma_vm_fixed0500_5_x9_50g":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed05005X950g,
	"vnuma_vm_fixed0600_6_x9_50g":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed06006X950g,
	"vnuma_vm_fixed0700_7_x9_50g":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed07007X950g,
	"vnuma_vm_fixed0800_8_x9_50g":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed08008X950g,
	"vnuma_vm_fixed0900_9_x9_50g":          UpdateVnicShapeDetailsVnicShapeVnumaVmFixed09009X950g,
	"vnuma_vm_fixed1000_10_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed100010X950g,
	"vnuma_vm_fixed1100_11_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed110011X950g,
	"vnuma_vm_fixed1200_12_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed120012X950g,
	"vnuma_vm_fixed1300_13_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed130013X950g,
	"vnuma_vm_fixed1400_14_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed140014X950g,
	"vnuma_vm_fixed1500_15_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed150015X950g,
	"vnuma_vm_fixed1600_16_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed160016X950g,
	"vnuma_vm_fixed1700_17_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed170017X950g,
	"vnuma_vm_fixed1800_18_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed180018X950g,
	"vnuma_vm_fixed1900_19_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed190019X950g,
	"vnuma_vm_fixed2000_20_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed200020X950g,
	"vnuma_vm_fixed2100_21_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed210021X950g,
	"vnuma_vm_fixed2200_22_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed220022X950g,
	"vnuma_vm_fixed2300_23_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed230023X950g,
	"vnuma_vm_fixed2400_24_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed240024X950g,
	"vnuma_vm_fixed2500_25_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed250025X950g,
	"vnuma_vm_fixed2600_26_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed260026X950g,
	"vnuma_vm_fixed2700_27_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed270027X950g,
	"vnuma_vm_fixed2800_28_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed280028X950g,
	"vnuma_vm_fixed2900_29_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed290029X950g,
	"vnuma_vm_fixed3000_30_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed300030X950g,
	"vnuma_vm_fixed3100_31_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed310031X950g,
	"vnuma_vm_fixed3200_32_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320032X950g,
	"vnuma_vm_fixed3200_33_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320033X950g,
	"vnuma_vm_fixed3200_34_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320034X950g,
	"vnuma_vm_fixed3200_35_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320035X950g,
	"vnuma_vm_fixed3200_36_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320036X950g,
	"vnuma_vm_fixed3200_37_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320037X950g,
	"vnuma_vm_fixed3200_38_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320038X950g,
	"vnuma_vm_fixed3200_39_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320039X950g,
	"vnuma_vm_fixed3200_40_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320040X950g,
	"vnuma_vm_fixed3200_41_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320041X950g,
	"vnuma_vm_fixed3200_42_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320042X950g,
	"vnuma_vm_fixed3200_43_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320043X950g,
	"vnuma_vm_fixed3200_44_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320044X950g,
	"vnuma_vm_fixed3200_45_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320045X950g,
	"vnuma_vm_fixed3200_46_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320046X950g,
	"vnuma_vm_fixed3200_47_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320047X950g,
	"vnuma_vm_fixed3200_48_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320048X950g,
	"vnuma_vm_fixed3200_49_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320049X950g,
	"vnuma_vm_fixed3200_50_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320050X950g,
	"vnuma_vm_fixed3200_51_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320051X950g,
	"vnuma_vm_fixed3200_52_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320052X950g,
	"vnuma_vm_fixed3200_53_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320053X950g,
	"vnuma_vm_fixed3200_54_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320054X950g,
	"vnuma_vm_fixed3200_55_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320055X950g,
	"vnuma_vm_fixed3200_56_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320056X950g,
	"vnuma_vm_fixed3200_57_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320057X950g,
	"vnuma_vm_fixed3200_58_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320058X950g,
	"vnuma_vm_fixed3200_59_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320059X950g,
	"vnuma_vm_fixed3200_60_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320060X950g,
	"vnuma_vm_fixed3200_61_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320061X950g,
	"vnuma_vm_fixed3200_62_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320062X950g,
	"vnuma_vm_fixed3200_63_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320063X950g,
	"vnuma_vm_fixed3200_64_x9_50g":         UpdateVnicShapeDetailsVnicShapeVnumaVmFixed320064X950g,
	"subcore_standard_vm_fixed0025_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0025X950g,
	"subcore_standard_vm_fixed0050_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0050X950g,
	"subcore_standard_vm_fixed0075_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0075X950g,
	"subcore_standard_vm_fixed0100_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0100X950g,
	"subcore_standard_vm_fixed0125_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0125X950g,
	"subcore_standard_vm_fixed0150_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0150X950g,
	"subcore_standard_vm_fixed0175_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0175X950g,
	"subcore_standard_vm_fixed0200_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0200X950g,
	"subcore_standard_vm_fixed0225_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0225X950g,
	"subcore_standard_vm_fixed0250_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0250X950g,
	"subcore_standard_vm_fixed0275_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0275X950g,
	"subcore_standard_vm_fixed0300_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0300X950g,
	"subcore_standard_vm_fixed0325_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0325X950g,
	"subcore_standard_vm_fixed0350_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0350X950g,
	"subcore_standard_vm_fixed0375_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0375X950g,
	"subcore_standard_vm_fixed0400_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0400X950g,
	"subcore_standard_vm_fixed0425_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0425X950g,
	"subcore_standard_vm_fixed0450_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0450X950g,
	"subcore_standard_vm_fixed0475_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0475X950g,
	"subcore_standard_vm_fixed0500_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0500X950g,
	"subcore_standard_vm_fixed0525_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0525X950g,
	"subcore_standard_vm_fixed0550_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0550X950g,
	"subcore_standard_vm_fixed0575_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0575X950g,
	"subcore_standard_vm_fixed0600_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0600X950g,
	"subcore_standard_vm_fixed0625_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0625X950g,
	"subcore_standard_vm_fixed0650_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0650X950g,
	"subcore_standard_vm_fixed0675_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0675X950g,
	"subcore_standard_vm_fixed0700_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0700X950g,
	"subcore_standard_vm_fixed0725_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0725X950g,
	"subcore_standard_vm_fixed0750_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0750X950g,
	"subcore_standard_vm_fixed0775_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0775X950g,
	"subcore_standard_vm_fixed0800_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0800X950g,
	"subcore_standard_vm_fixed0825_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0825X950g,
	"subcore_standard_vm_fixed0850_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0850X950g,
	"subcore_standard_vm_fixed0875_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0875X950g,
	"subcore_standard_vm_fixed0900_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0900X950g,
	"subcore_standard_vm_fixed0925_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0925X950g,
	"subcore_standard_vm_fixed0950_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0950X950g,
	"subcore_standard_vm_fixed0975_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed0975X950g,
	"subcore_standard_vm_fixed1000_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1000X950g,
	"subcore_standard_vm_fixed1025_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1025X950g,
	"subcore_standard_vm_fixed1050_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1050X950g,
	"subcore_standard_vm_fixed1075_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1075X950g,
	"subcore_standard_vm_fixed1100_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1100X950g,
	"subcore_standard_vm_fixed1125_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1125X950g,
	"subcore_standard_vm_fixed1150_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1150X950g,
	"subcore_standard_vm_fixed1175_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1175X950g,
	"subcore_standard_vm_fixed1200_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1200X950g,
	"subcore_standard_vm_fixed1225_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1225X950g,
	"subcore_standard_vm_fixed1250_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1250X950g,
	"subcore_standard_vm_fixed1275_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1275X950g,
	"subcore_standard_vm_fixed1300_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1300X950g,
	"subcore_standard_vm_fixed1325_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1325X950g,
	"subcore_standard_vm_fixed1350_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1350X950g,
	"subcore_standard_vm_fixed1375_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1375X950g,
	"subcore_standard_vm_fixed1400_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1400X950g,
	"subcore_standard_vm_fixed1425_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1425X950g,
	"subcore_standard_vm_fixed1450_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1450X950g,
	"subcore_standard_vm_fixed1475_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1475X950g,
	"subcore_standard_vm_fixed1500_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1500X950g,
	"subcore_standard_vm_fixed1525_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1525X950g,
	"subcore_standard_vm_fixed1550_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1550X950g,
	"subcore_standard_vm_fixed1575_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1575X950g,
	"subcore_standard_vm_fixed1600_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1600X950g,
	"subcore_standard_vm_fixed1625_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1625X950g,
	"subcore_standard_vm_fixed1650_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1650X950g,
	"subcore_standard_vm_fixed1700_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1700X950g,
	"subcore_standard_vm_fixed1725_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1725X950g,
	"subcore_standard_vm_fixed1750_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1750X950g,
	"subcore_standard_vm_fixed1800_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1800X950g,
	"subcore_standard_vm_fixed1850_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1850X950g,
	"subcore_standard_vm_fixed1875_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1875X950g,
	"subcore_standard_vm_fixed1900_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1900X950g,
	"subcore_standard_vm_fixed1925_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1925X950g,
	"subcore_standard_vm_fixed1950_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed1950X950g,
	"subcore_standard_vm_fixed2000_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2000X950g,
	"subcore_standard_vm_fixed2025_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2025X950g,
	"subcore_standard_vm_fixed2050_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2050X950g,
	"subcore_standard_vm_fixed2100_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2100X950g,
	"subcore_standard_vm_fixed2125_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2125X950g,
	"subcore_standard_vm_fixed2150_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2150X950g,
	"subcore_standard_vm_fixed2175_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2175X950g,
	"subcore_standard_vm_fixed2200_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2200X950g,
	"subcore_standard_vm_fixed2250_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2250X950g,
	"subcore_standard_vm_fixed2275_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2275X950g,
	"subcore_standard_vm_fixed2300_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2300X950g,
	"subcore_standard_vm_fixed2325_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2325X950g,
	"subcore_standard_vm_fixed2350_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2350X950g,
	"subcore_standard_vm_fixed2375_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2375X950g,
	"subcore_standard_vm_fixed2400_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2400X950g,
	"subcore_standard_vm_fixed2450_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2450X950g,
	"subcore_standard_vm_fixed2475_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2475X950g,
	"subcore_standard_vm_fixed2500_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2500X950g,
	"subcore_standard_vm_fixed2550_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2550X950g,
	"subcore_standard_vm_fixed2600_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2600X950g,
	"subcore_standard_vm_fixed2625_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2625X950g,
	"subcore_standard_vm_fixed2650_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2650X950g,
	"subcore_standard_vm_fixed2700_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2700X950g,
	"subcore_standard_vm_fixed2750_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2750X950g,
	"subcore_standard_vm_fixed2775_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2775X950g,
	"subcore_standard_vm_fixed2800_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2800X950g,
	"subcore_standard_vm_fixed2850_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2850X950g,
	"subcore_standard_vm_fixed2875_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2875X950g,
	"subcore_standard_vm_fixed2900_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2900X950g,
	"subcore_standard_vm_fixed2925_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2925X950g,
	"subcore_standard_vm_fixed2950_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2950X950g,
	"subcore_standard_vm_fixed2975_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed2975X950g,
	"subcore_standard_vm_fixed3000_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3000X950g,
	"subcore_standard_vm_fixed3025_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3025X950g,
	"subcore_standard_vm_fixed3050_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3050X950g,
	"subcore_standard_vm_fixed3075_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3075X950g,
	"subcore_standard_vm_fixed3100_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3100X950g,
	"subcore_standard_vm_fixed3125_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3125X950g,
	"subcore_standard_vm_fixed3150_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3150X950g,
	"subcore_standard_vm_fixed3200_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3200X950g,
	"subcore_standard_vm_fixed3225_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3225X950g,
	"subcore_standard_vm_fixed3250_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3250X950g,
	"subcore_standard_vm_fixed3300_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3300X950g,
	"subcore_standard_vm_fixed3325_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3325X950g,
	"subcore_standard_vm_fixed3375_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3375X950g,
	"subcore_standard_vm_fixed3400_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3400X950g,
	"subcore_standard_vm_fixed3450_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3450X950g,
	"subcore_standard_vm_fixed3500_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3500X950g,
	"subcore_standard_vm_fixed3525_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3525X950g,
	"subcore_standard_vm_fixed3575_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3575X950g,
	"subcore_standard_vm_fixed3600_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3600X950g,
	"subcore_standard_vm_fixed3625_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3625X950g,
	"subcore_standard_vm_fixed3675_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3675X950g,
	"subcore_standard_vm_fixed3700_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3700X950g,
	"subcore_standard_vm_fixed3750_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3750X950g,
	"subcore_standard_vm_fixed3800_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3800X950g,
	"subcore_standard_vm_fixed3825_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3825X950g,
	"subcore_standard_vm_fixed3850_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3850X950g,
	"subcore_standard_vm_fixed3875_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3875X950g,
	"subcore_standard_vm_fixed3900_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3900X950g,
	"subcore_standard_vm_fixed3975_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed3975X950g,
	"subcore_standard_vm_fixed4000_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4000X950g,
	"subcore_standard_vm_fixed4025_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4025X950g,
	"subcore_standard_vm_fixed4050_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4050X950g,
	"subcore_standard_vm_fixed4100_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4100X950g,
	"subcore_standard_vm_fixed4125_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4125X950g,
	"subcore_standard_vm_fixed4200_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4200X950g,
	"subcore_standard_vm_fixed4225_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4225X950g,
	"subcore_standard_vm_fixed4250_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4250X950g,
	"subcore_standard_vm_fixed4275_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4275X950g,
	"subcore_standard_vm_fixed4300_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4300X950g,
	"subcore_standard_vm_fixed4350_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4350X950g,
	"subcore_standard_vm_fixed4375_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4375X950g,
	"subcore_standard_vm_fixed4400_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4400X950g,
	"subcore_standard_vm_fixed4425_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4425X950g,
	"subcore_standard_vm_fixed4500_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4500X950g,
	"subcore_standard_vm_fixed4550_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4550X950g,
	"subcore_standard_vm_fixed4575_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4575X950g,
	"subcore_standard_vm_fixed4600_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4600X950g,
	"subcore_standard_vm_fixed4625_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4625X950g,
	"subcore_standard_vm_fixed4650_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4650X950g,
	"subcore_standard_vm_fixed4675_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4675X950g,
	"subcore_standard_vm_fixed4700_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4700X950g,
	"subcore_standard_vm_fixed4725_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4725X950g,
	"subcore_standard_vm_fixed4750_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4750X950g,
	"subcore_standard_vm_fixed4800_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4800X950g,
	"subcore_standard_vm_fixed4875_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4875X950g,
	"subcore_standard_vm_fixed4900_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4900X950g,
	"subcore_standard_vm_fixed4950_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed4950X950g,
	"subcore_standard_vm_fixed5000_x9_50g": UpdateVnicShapeDetailsVnicShapeSubcoreStandardVmFixed5000X950g,
	"entirehost_x9_50g":                    UpdateVnicShapeDetailsVnicShapeEntirehostX950g,
	"dynamic_x9_100g":                      UpdateVnicShapeDetailsVnicShapeDynamicX9100g,
	"dynamic_x10_50g":                      UpdateVnicShapeDetailsVnicShapeDynamicX1050g,
	"dynamic_x10_100g":                     UpdateVnicShapeDetailsVnicShapeDynamicX10100g,
}

// GetUpdateVnicShapeDetailsVnicShapeEnumValues Enumerates the set of values for UpdateVnicShapeDetailsVnicShapeEnum
func GetUpdateVnicShapeDetailsVnicShapeEnumValues() []UpdateVnicShapeDetailsVnicShapeEnum {
	values := make([]UpdateVnicShapeDetailsVnicShapeEnum, 0)
	for _, v := range mappingUpdateVnicShapeDetailsVnicShapeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateVnicShapeDetailsVnicShapeEnumStringValues Enumerates the set of values in String for UpdateVnicShapeDetailsVnicShapeEnum
func GetUpdateVnicShapeDetailsVnicShapeEnumStringValues() []string {
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
		"VNICAAS_FIXED0025",
		"VNICAAS_FIXED0200",
		"VNICAAS_FIXED0400",
		"VNICAAS_FIXED0625",
		"VNICAAS_FIXED0700",
		"VNICAAS_FIXED1250",
		"VNICAAS_FWAAS_6G_250M",
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
		"FIXED4700_E4_50G",
		"FIXED4960_E4_50G",
		"ENTIREHOST_E4_50G",
		"VNUMA_VM_FIXED0100_1_E4_50G",
		"VNUMA_VM_FIXED0200_2_E4_50G",
		"VNUMA_VM_FIXED0300_3_E4_50G",
		"VNUMA_VM_FIXED0400_4_E4_50G",
		"VNUMA_VM_FIXED0500_5_E4_50G",
		"VNUMA_VM_FIXED0600_6_E4_50G",
		"VNUMA_VM_FIXED0700_7_E4_50G",
		"VNUMA_VM_FIXED0800_8_E4_50G",
		"VNUMA_VM_FIXED0900_9_E4_50G",
		"VNUMA_VM_FIXED1000_10_E4_50G",
		"VNUMA_VM_FIXED1100_11_E4_50G",
		"VNUMA_VM_FIXED1200_12_E4_50G",
		"VNUMA_VM_FIXED1300_13_E4_50G",
		"VNUMA_VM_FIXED1400_14_E4_50G",
		"VNUMA_VM_FIXED1500_15_E4_50G",
		"VNUMA_VM_FIXED1600_16_E4_50G",
		"VNUMA_VM_FIXED1700_17_E4_50G",
		"VNUMA_VM_FIXED1800_18_E4_50G",
		"VNUMA_VM_FIXED1900_19_E4_50G",
		"VNUMA_VM_FIXED2000_20_E4_50G",
		"VNUMA_VM_FIXED2100_21_E4_50G",
		"VNUMA_VM_FIXED2200_22_E4_50G",
		"VNUMA_VM_FIXED2300_23_E4_50G",
		"VNUMA_VM_FIXED2400_24_E4_50G",
		"VNUMA_VM_FIXED2500_25_E4_50G",
		"VNUMA_VM_FIXED2600_26_E4_50G",
		"VNUMA_VM_FIXED2700_27_E4_50G",
		"VNUMA_VM_FIXED2800_28_E4_50G",
		"VNUMA_VM_FIXED2900_29_E4_50G",
		"VNUMA_VM_FIXED3000_30_E4_50G",
		"VNUMA_VM_FIXED3100_31_E4_50G",
		"VNUMA_VM_FIXED3200_32_E4_50G",
		"VNUMA_VM_FIXED3300_33_E4_50G",
		"VNUMA_VM_FIXED3400_34_E4_50G",
		"VNUMA_VM_FIXED3500_35_E4_50G",
		"VNUMA_VM_FIXED3600_36_E4_50G",
		"VNUMA_VM_FIXED3700_37_E4_50G",
		"VNUMA_VM_FIXED3800_38_E4_50G",
		"VNUMA_VM_FIXED3900_39_E4_50G",
		"VNUMA_VM_FIXED4000_40_E4_50G",
		"VNUMA_VM_FIXED4000_41_E4_50G",
		"VNUMA_VM_FIXED4000_42_E4_50G",
		"VNUMA_VM_FIXED4000_43_E4_50G",
		"VNUMA_VM_FIXED4000_44_E4_50G",
		"VNUMA_VM_FIXED4000_45_E4_50G",
		"VNUMA_VM_FIXED4000_46_E4_50G",
		"VNUMA_VM_FIXED4000_47_E4_50G",
		"VNUMA_VM_FIXED4000_48_E4_50G",
		"VNUMA_VM_FIXED4000_49_E4_50G",
		"VNUMA_VM_FIXED4000_50_E4_50G",
		"VNUMA_VM_FIXED4000_51_E4_50G",
		"VNUMA_VM_FIXED4000_52_E4_50G",
		"VNUMA_VM_FIXED4000_53_E4_50G",
		"VNUMA_VM_FIXED4000_54_E4_50G",
		"VNUMA_VM_FIXED4000_55_E4_50G",
		"VNUMA_VM_FIXED4000_56_E4_50G",
		"VNUMA_VM_FIXED4000_57_E4_50G",
		"VNUMA_VM_FIXED4000_58_E4_50G",
		"VNUMA_VM_FIXED4000_59_E4_50G",
		"VNUMA_VM_FIXED4000_60_E4_50G",
		"VNUMA_VM_FIXED4000_61_E4_50G",
		"VNUMA_VM_FIXED4000_62_E4_50G",
		"VNUMA_VM_FIXED4000_63_E4_50G",
		"VNUMA_VM_FIXED4000_64_E4_50G",
		"VNUMA_VM_FIXED4000_65_E4_50G",
		"VNUMA_VM_FIXED4000_66_E4_50G",
		"VNUMA_VM_FIXED4000_67_E4_50G",
		"VNUMA_VM_FIXED4000_68_E4_50G",
		"VNUMA_VM_FIXED4000_69_E4_50G",
		"VNUMA_VM_FIXED4000_70_E4_50G",
		"VNUMA_VM_FIXED4000_71_E4_50G",
		"VNUMA_VM_FIXED4000_72_E4_50G",
		"VNUMA_VM_FIXED4000_73_E4_50G",
		"VNUMA_VM_FIXED4000_74_E4_50G",
		"VNUMA_VM_FIXED4000_75_E4_50G",
		"VNUMA_VM_FIXED4000_76_E4_50G",
		"VNUMA_VM_FIXED4000_77_E4_50G",
		"VNUMA_VM_FIXED4000_78_E4_50G",
		"VNUMA_VM_FIXED4000_79_E4_50G",
		"VNUMA_VM_FIXED4000_80_E4_50G",
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
		"EntireHost_E5_50G",
		"Fixed0040_E5_50G",
		"FIXED0100_E5_50G",
		"FIXED0200_E5_50G",
		"FIXED0300_E5_50G",
		"FIXED0400_E5_50G",
		"FIXED0500_E5_50G",
		"FIXED0600_E5_50G",
		"FIXED0700_E5_50G",
		"FIXED0800_E5_50G",
		"FIXED0900_E5_50G",
		"FIXED1000_E5_50G",
		"FIXED1100_E5_50G",
		"FIXED1200_E5_50G",
		"FIXED1300_E5_50G",
		"FIXED1400_E5_50G",
		"FIXED1500_E5_50G",
		"FIXED1600_E5_50G",
		"FIXED1700_E5_50G",
		"FIXED1800_E5_50G",
		"FIXED1900_E5_50G",
		"FIXED2000_E5_50G",
		"FIXED2100_E5_50G",
		"FIXED2200_E5_50G",
		"FIXED2300_E5_50G",
		"FIXED2400_E5_50G",
		"FIXED2500_E5_50G",
		"FIXED2600_E5_50G",
		"FIXED2700_E5_50G",
		"FIXED2800_E5_50G",
		"FIXED2900_E5_50G",
		"FIXED3000_E5_50G",
		"FIXED3100_E5_50G",
		"FIXED3200_E5_50G",
		"FIXED3300_E5_50G",
		"FIXED3400_E5_50G",
		"FIXED3500_E5_50G",
		"FIXED3600_E5_50G",
		"FIXED3700_E5_50G",
		"FIXED3800_E5_50G",
		"FIXED3900_E5_50G",
		"FIXED4000_E5_50G",
		"SUBCORE_VM_FIXED0025_E5_50G",
		"SUBCORE_VM_FIXED0050_E5_50G",
		"SUBCORE_VM_FIXED0075_E5_50G",
		"SUBCORE_VM_FIXED0100_E5_50G",
		"SUBCORE_VM_FIXED0125_E5_50G",
		"SUBCORE_VM_FIXED0150_E5_50G",
		"SUBCORE_VM_FIXED0175_E5_50G",
		"SUBCORE_VM_FIXED0200_E5_50G",
		"SUBCORE_VM_FIXED0225_E5_50G",
		"SUBCORE_VM_FIXED0250_E5_50G",
		"SUBCORE_VM_FIXED0275_E5_50G",
		"SUBCORE_VM_FIXED0300_E5_50G",
		"SUBCORE_VM_FIXED0325_E5_50G",
		"SUBCORE_VM_FIXED0350_E5_50G",
		"SUBCORE_VM_FIXED0375_E5_50G",
		"SUBCORE_VM_FIXED0400_E5_50G",
		"SUBCORE_VM_FIXED0425_E5_50G",
		"SUBCORE_VM_FIXED0450_E5_50G",
		"SUBCORE_VM_FIXED0475_E5_50G",
		"SUBCORE_VM_FIXED0500_E5_50G",
		"SUBCORE_VM_FIXED0525_E5_50G",
		"SUBCORE_VM_FIXED0550_E5_50G",
		"SUBCORE_VM_FIXED0575_E5_50G",
		"SUBCORE_VM_FIXED0600_E5_50G",
		"SUBCORE_VM_FIXED0625_E5_50G",
		"SUBCORE_VM_FIXED0650_E5_50G",
		"SUBCORE_VM_FIXED0675_E5_50G",
		"SUBCORE_VM_FIXED0700_E5_50G",
		"SUBCORE_VM_FIXED0725_E5_50G",
		"SUBCORE_VM_FIXED0750_E5_50G",
		"SUBCORE_VM_FIXED0775_E5_50G",
		"SUBCORE_VM_FIXED0800_E5_50G",
		"SUBCORE_VM_FIXED0825_E5_50G",
		"SUBCORE_VM_FIXED0850_E5_50G",
		"SUBCORE_VM_FIXED0875_E5_50G",
		"SUBCORE_VM_FIXED0900_E5_50G",
		"SUBCORE_VM_FIXED0925_E5_50G",
		"SUBCORE_VM_FIXED0950_E5_50G",
		"SUBCORE_VM_FIXED0975_E5_50G",
		"SUBCORE_VM_FIXED1000_E5_50G",
		"SUBCORE_VM_FIXED1025_E5_50G",
		"SUBCORE_VM_FIXED1050_E5_50G",
		"SUBCORE_VM_FIXED1075_E5_50G",
		"SUBCORE_VM_FIXED1100_E5_50G",
		"SUBCORE_VM_FIXED1125_E5_50G",
		"SUBCORE_VM_FIXED1150_E5_50G",
		"SUBCORE_VM_FIXED1175_E5_50G",
		"SUBCORE_VM_FIXED1200_E5_50G",
		"SUBCORE_VM_FIXED1225_E5_50G",
		"SUBCORE_VM_FIXED1250_E5_50G",
		"SUBCORE_VM_FIXED1275_E5_50G",
		"SUBCORE_VM_FIXED1300_E5_50G",
		"SUBCORE_VM_FIXED1325_E5_50G",
		"SUBCORE_VM_FIXED1350_E5_50G",
		"SUBCORE_VM_FIXED1375_E5_50G",
		"SUBCORE_VM_FIXED1400_E5_50G",
		"SUBCORE_VM_FIXED1425_E5_50G",
		"SUBCORE_VM_FIXED1450_E5_50G",
		"SUBCORE_VM_FIXED1475_E5_50G",
		"SUBCORE_VM_FIXED1500_E5_50G",
		"SUBCORE_VM_FIXED1525_E5_50G",
		"SUBCORE_VM_FIXED1550_E5_50G",
		"SUBCORE_VM_FIXED1575_E5_50G",
		"SUBCORE_VM_FIXED1600_E5_50G",
		"SUBCORE_VM_FIXED1625_E5_50G",
		"SUBCORE_VM_FIXED1650_E5_50G",
		"SUBCORE_VM_FIXED1675_E5_50G",
		"SUBCORE_VM_FIXED1700_E5_50G",
		"SUBCORE_VM_FIXED1725_E5_50G",
		"SUBCORE_VM_FIXED1750_E5_50G",
		"SUBCORE_VM_FIXED1775_E5_50G",
		"SUBCORE_VM_FIXED1800_E5_50G",
		"SUBCORE_VM_FIXED1825_E5_50G",
		"SUBCORE_VM_FIXED1850_E5_50G",
		"SUBCORE_VM_FIXED1875_E5_50G",
		"SUBCORE_VM_FIXED1900_E5_50G",
		"SUBCORE_VM_FIXED1925_E5_50G",
		"SUBCORE_VM_FIXED1950_E5_50G",
		"SUBCORE_VM_FIXED1975_E5_50G",
		"SUBCORE_VM_FIXED2000_E5_50G",
		"SUBCORE_VM_FIXED2025_E5_50G",
		"SUBCORE_VM_FIXED2050_E5_50G",
		"SUBCORE_VM_FIXED2075_E5_50G",
		"SUBCORE_VM_FIXED2100_E5_50G",
		"SUBCORE_VM_FIXED2125_E5_50G",
		"SUBCORE_VM_FIXED2150_E5_50G",
		"SUBCORE_VM_FIXED2175_E5_50G",
		"SUBCORE_VM_FIXED2200_E5_50G",
		"SUBCORE_VM_FIXED2225_E5_50G",
		"SUBCORE_VM_FIXED2250_E5_50G",
		"SUBCORE_VM_FIXED2275_E5_50G",
		"SUBCORE_VM_FIXED2300_E5_50G",
		"SUBCORE_VM_FIXED2325_E5_50G",
		"SUBCORE_VM_FIXED2350_E5_50G",
		"SUBCORE_VM_FIXED2375_E5_50G",
		"SUBCORE_VM_FIXED2400_E5_50G",
		"SUBCORE_VM_FIXED2425_E5_50G",
		"SUBCORE_VM_FIXED2450_E5_50G",
		"SUBCORE_VM_FIXED2475_E5_50G",
		"SUBCORE_VM_FIXED2500_E5_50G",
		"SUBCORE_VM_FIXED2525_E5_50G",
		"SUBCORE_VM_FIXED2550_E5_50G",
		"SUBCORE_VM_FIXED2575_E5_50G",
		"SUBCORE_VM_FIXED2600_E5_50G",
		"SUBCORE_VM_FIXED2625_E5_50G",
		"SUBCORE_VM_FIXED2650_E5_50G",
		"SUBCORE_VM_FIXED2675_E5_50G",
		"SUBCORE_VM_FIXED2700_E5_50G",
		"SUBCORE_VM_FIXED2725_E5_50G",
		"SUBCORE_VM_FIXED2750_E5_50G",
		"SUBCORE_VM_FIXED2775_E5_50G",
		"SUBCORE_VM_FIXED2800_E5_50G",
		"SUBCORE_VM_FIXED2825_E5_50G",
		"SUBCORE_VM_FIXED2850_E5_50G",
		"SUBCORE_VM_FIXED2875_E5_50G",
		"SUBCORE_VM_FIXED2900_E5_50G",
		"SUBCORE_VM_FIXED2925_E5_50G",
		"SUBCORE_VM_FIXED2950_E5_50G",
		"SUBCORE_VM_FIXED2975_E5_50G",
		"SUBCORE_VM_FIXED3000_E5_50G",
		"SUBCORE_VM_FIXED3025_E5_50G",
		"SUBCORE_VM_FIXED3050_E5_50G",
		"SUBCORE_VM_FIXED3075_E5_50G",
		"SUBCORE_VM_FIXED3100_E5_50G",
		"SUBCORE_VM_FIXED3125_E5_50G",
		"SUBCORE_VM_FIXED3150_E5_50G",
		"SUBCORE_VM_FIXED3175_E5_50G",
		"SUBCORE_VM_FIXED3200_E5_50G",
		"SUBCORE_VM_FIXED3225_E5_50G",
		"SUBCORE_VM_FIXED3250_E5_50G",
		"SUBCORE_VM_FIXED3275_E5_50G",
		"SUBCORE_VM_FIXED3300_E5_50G",
		"SUBCORE_VM_FIXED3325_E5_50G",
		"SUBCORE_VM_FIXED3350_E5_50G",
		"SUBCORE_VM_FIXED3375_E5_50G",
		"SUBCORE_VM_FIXED3400_E5_50G",
		"SUBCORE_VM_FIXED3425_E5_50G",
		"SUBCORE_VM_FIXED3450_E5_50G",
		"SUBCORE_VM_FIXED3475_E5_50G",
		"SUBCORE_VM_FIXED3500_E5_50G",
		"SUBCORE_VM_FIXED3525_E5_50G",
		"SUBCORE_VM_FIXED3550_E5_50G",
		"SUBCORE_VM_FIXED3575_E5_50G",
		"SUBCORE_VM_FIXED3600_E5_50G",
		"SUBCORE_VM_FIXED3625_E5_50G",
		"SUBCORE_VM_FIXED3650_E5_50G",
		"SUBCORE_VM_FIXED3675_E5_50G",
		"SUBCORE_VM_FIXED3700_E5_50G",
		"SUBCORE_VM_FIXED3725_E5_50G",
		"SUBCORE_VM_FIXED3750_E5_50G",
		"SUBCORE_VM_FIXED3775_E5_50G",
		"SUBCORE_VM_FIXED3800_E5_50G",
		"SUBCORE_VM_FIXED3825_E5_50G",
		"SUBCORE_VM_FIXED3850_E5_50G",
		"SUBCORE_VM_FIXED3875_E5_50G",
		"SUBCORE_VM_FIXED3900_E5_50G",
		"SUBCORE_VM_FIXED3925_E5_50G",
		"SUBCORE_VM_FIXED3950_E5_50G",
		"SUBCORE_VM_FIXED3975_E5_50G",
		"SUBCORE_VM_FIXED4000_E5_50G",
		"SUBCORE_VM_FIXED4025_E5_50G",
		"SUBCORE_VM_FIXED4050_E5_50G",
		"SUBCORE_VM_FIXED4075_E5_50G",
		"SUBCORE_VM_FIXED4100_E5_50G",
		"SUBCORE_VM_FIXED4125_E5_50G",
		"SUBCORE_VM_FIXED4150_E5_50G",
		"SUBCORE_VM_FIXED4175_E5_50G",
		"SUBCORE_VM_FIXED4200_E5_50G",
		"SUBCORE_VM_FIXED4225_E5_50G",
		"SUBCORE_VM_FIXED4250_E5_50G",
		"SUBCORE_VM_FIXED4275_E5_50G",
		"SUBCORE_VM_FIXED4300_E5_50G",
		"SUBCORE_VM_FIXED4325_E5_50G",
		"SUBCORE_VM_FIXED4350_E5_50G",
		"SUBCORE_VM_FIXED4375_E5_50G",
		"SUBCORE_VM_FIXED4400_E5_50G",
		"SUBCORE_VM_FIXED4425_E5_50G",
		"SUBCORE_VM_FIXED4450_E5_50G",
		"SUBCORE_VM_FIXED4475_E5_50G",
		"SUBCORE_VM_FIXED4500_E5_50G",
		"SUBCORE_VM_FIXED4525_E5_50G",
		"SUBCORE_VM_FIXED4550_E5_50G",
		"SUBCORE_VM_FIXED4575_E5_50G",
		"SUBCORE_VM_FIXED4600_E5_50G",
		"SUBCORE_VM_FIXED4625_E5_50G",
		"SUBCORE_VM_FIXED4650_E5_50G",
		"SUBCORE_VM_FIXED4675_E5_50G",
		"SUBCORE_VM_FIXED4700_E5_50G",
		"SUBCORE_VM_FIXED4725_E5_50G",
		"SUBCORE_VM_FIXED4750_E5_50G",
		"SUBCORE_VM_FIXED4775_E5_50G",
		"SUBCORE_VM_FIXED4800_E5_50G",
		"SUBCORE_VM_FIXED4825_E5_50G",
		"SUBCORE_VM_FIXED4850_E5_50G",
		"SUBCORE_VM_FIXED4875_E5_50G",
		"SUBCORE_VM_FIXED4900_E5_50G",
		"SUBCORE_VM_FIXED4925_E5_50G",
		"SUBCORE_VM_FIXED4950_E5_50G",
		"SUBCORE_VM_FIXED4975_E5_50G",
		"SUBCORE_VM_FIXED5000_E5_50G",
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
		"DYNAMIC_X7_25G",
		"FIXED0040_X7_25G",
		"STANDARD_VM_FIXED0100_X7_25G",
		"STANDARD_VM_FIXED0200_X7_25G",
		"STANDARD_VM_FIXED0300_X7_25G",
		"STANDARD_VM_FIXED0400_X7_25G",
		"STANDARD_VM_FIXED0500_X7_25G",
		"STANDARD_VM_FIXED0600_X7_25G",
		"STANDARD_VM_FIXED0700_X7_25G",
		"STANDARD_VM_FIXED0800_X7_25G",
		"STANDARD_VM_FIXED0900_X7_25G",
		"STANDARD_VM_FIXED1000_X7_25G",
		"STANDARD_VM_FIXED1100_X7_25G",
		"STANDARD_VM_FIXED1200_X7_25G",
		"STANDARD_VM_FIXED1300_X7_25G",
		"STANDARD_VM_FIXED1400_X7_25G",
		"STANDARD_VM_FIXED1500_X7_25G",
		"STANDARD_VM_FIXED1600_X7_25G",
		"STANDARD_VM_FIXED1700_X7_25G",
		"STANDARD_VM_FIXED1800_X7_25G",
		"STANDARD_VM_FIXED1900_X7_25G",
		"STANDARD_VM_FIXED2000_X7_25G",
		"STANDARD_VM_FIXED2100_X7_25G",
		"STANDARD_VM_FIXED2200_X7_25G",
		"STANDARD_VM_FIXED2300_X7_25G",
		"STANDARD_VM_FIXED2400_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0025_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0050_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0075_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0100_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0125_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0150_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0175_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0200_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0225_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0250_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0275_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0300_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0325_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0350_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0375_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0400_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0425_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0450_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0475_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0500_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0525_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0550_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0575_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0600_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0625_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0650_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0675_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0700_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0725_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0750_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0775_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0800_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0825_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0850_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0875_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0900_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0925_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0950_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED0975_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1000_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1025_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1050_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1075_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1100_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1125_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1150_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1175_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1200_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1225_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1250_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1275_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1300_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1325_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1350_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1375_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1400_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1425_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1450_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1475_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1500_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1525_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1550_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1575_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1600_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1625_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1650_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1675_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1700_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1725_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1750_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1775_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1800_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1825_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1850_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1875_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1900_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1925_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1950_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED1975_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED2000_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED2025_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED2050_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED2075_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED2100_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED2125_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED2150_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED2175_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED2200_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED2225_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED2250_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED2275_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED2300_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED2325_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED2350_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED2375_X7_25G",
		"SUBCORE_STANDARD_VM_FIXED2400_X7_25G",
		"ENTIREHOST_X7_25G",
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
		"A10_GPU_VM_FIXED2400_X9_50G",
		"A10_GPU_VM_FIXED4800_X9_50G",
		"VNUMA_VM_FIXED0100_1_X9_50G",
		"VNUMA_VM_FIXED0200_2_X9_50G",
		"VNUMA_VM_FIXED0300_3_X9_50G",
		"VNUMA_VM_FIXED0400_4_X9_50G",
		"VNUMA_VM_FIXED0500_5_X9_50G",
		"VNUMA_VM_FIXED0600_6_X9_50G",
		"VNUMA_VM_FIXED0700_7_X9_50G",
		"VNUMA_VM_FIXED0800_8_X9_50G",
		"VNUMA_VM_FIXED0900_9_X9_50G",
		"VNUMA_VM_FIXED1000_10_X9_50G",
		"VNUMA_VM_FIXED1100_11_X9_50G",
		"VNUMA_VM_FIXED1200_12_X9_50G",
		"VNUMA_VM_FIXED1300_13_X9_50G",
		"VNUMA_VM_FIXED1400_14_X9_50G",
		"VNUMA_VM_FIXED1500_15_X9_50G",
		"VNUMA_VM_FIXED1600_16_X9_50G",
		"VNUMA_VM_FIXED1700_17_X9_50G",
		"VNUMA_VM_FIXED1800_18_X9_50G",
		"VNUMA_VM_FIXED1900_19_X9_50G",
		"VNUMA_VM_FIXED2000_20_X9_50G",
		"VNUMA_VM_FIXED2100_21_X9_50G",
		"VNUMA_VM_FIXED2200_22_X9_50G",
		"VNUMA_VM_FIXED2300_23_X9_50G",
		"VNUMA_VM_FIXED2400_24_X9_50G",
		"VNUMA_VM_FIXED2500_25_X9_50G",
		"VNUMA_VM_FIXED2600_26_X9_50G",
		"VNUMA_VM_FIXED2700_27_X9_50G",
		"VNUMA_VM_FIXED2800_28_X9_50G",
		"VNUMA_VM_FIXED2900_29_X9_50G",
		"VNUMA_VM_FIXED3000_30_X9_50G",
		"VNUMA_VM_FIXED3100_31_X9_50G",
		"VNUMA_VM_FIXED3200_32_X9_50G",
		"VNUMA_VM_FIXED3200_33_X9_50G",
		"VNUMA_VM_FIXED3200_34_X9_50G",
		"VNUMA_VM_FIXED3200_35_X9_50G",
		"VNUMA_VM_FIXED3200_36_X9_50G",
		"VNUMA_VM_FIXED3200_37_X9_50G",
		"VNUMA_VM_FIXED3200_38_X9_50G",
		"VNUMA_VM_FIXED3200_39_X9_50G",
		"VNUMA_VM_FIXED3200_40_X9_50G",
		"VNUMA_VM_FIXED3200_41_X9_50G",
		"VNUMA_VM_FIXED3200_42_X9_50G",
		"VNUMA_VM_FIXED3200_43_X9_50G",
		"VNUMA_VM_FIXED3200_44_X9_50G",
		"VNUMA_VM_FIXED3200_45_X9_50G",
		"VNUMA_VM_FIXED3200_46_X9_50G",
		"VNUMA_VM_FIXED3200_47_X9_50G",
		"VNUMA_VM_FIXED3200_48_X9_50G",
		"VNUMA_VM_FIXED3200_49_X9_50G",
		"VNUMA_VM_FIXED3200_50_X9_50G",
		"VNUMA_VM_FIXED3200_51_X9_50G",
		"VNUMA_VM_FIXED3200_52_X9_50G",
		"VNUMA_VM_FIXED3200_53_X9_50G",
		"VNUMA_VM_FIXED3200_54_X9_50G",
		"VNUMA_VM_FIXED3200_55_X9_50G",
		"VNUMA_VM_FIXED3200_56_X9_50G",
		"VNUMA_VM_FIXED3200_57_X9_50G",
		"VNUMA_VM_FIXED3200_58_X9_50G",
		"VNUMA_VM_FIXED3200_59_X9_50G",
		"VNUMA_VM_FIXED3200_60_X9_50G",
		"VNUMA_VM_FIXED3200_61_X9_50G",
		"VNUMA_VM_FIXED3200_62_X9_50G",
		"VNUMA_VM_FIXED3200_63_X9_50G",
		"VNUMA_VM_FIXED3200_64_X9_50G",
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

// GetMappingUpdateVnicShapeDetailsVnicShapeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateVnicShapeDetailsVnicShapeEnum(val string) (UpdateVnicShapeDetailsVnicShapeEnum, bool) {
	enum, ok := mappingUpdateVnicShapeDetailsVnicShapeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
