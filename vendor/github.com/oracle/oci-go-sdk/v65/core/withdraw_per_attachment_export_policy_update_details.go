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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WithdrawPerAttachmentExportPolicyUpdateDetails Details for withdrawing attachment specific export policy update
type WithdrawPerAttachmentExportPolicyUpdateDetails struct {

	// The attachment type.
	AttachmentType WithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum `mandatory:"false" json:"attachmentType,omitempty"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the
	// DRG attachment.
	DrgAttachmentId *string `mandatory:"false" json:"drgAttachmentId"`
}

func (m WithdrawPerAttachmentExportPolicyUpdateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WithdrawPerAttachmentExportPolicyUpdateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingWithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum(string(m.AttachmentType)); !ok && m.AttachmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttachmentType: %s. Supported values are: %s.", m.AttachmentType, strings.Join(GetWithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum Enum with underlying type: string
type WithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum string

// Set of constants representing the allowable values for WithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum
const (
	WithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeVirtualCircuit WithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum = "VIRTUAL_CIRCUIT"
	WithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeIpsecTunnel    WithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum = "IPSEC_TUNNEL"
)

var mappingWithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum = map[string]WithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum{
	"VIRTUAL_CIRCUIT": WithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeVirtualCircuit,
	"IPSEC_TUNNEL":    WithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeIpsecTunnel,
}

var mappingWithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnumLowerCase = map[string]WithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum{
	"virtual_circuit": WithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeVirtualCircuit,
	"ipsec_tunnel":    WithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeIpsecTunnel,
}

// GetWithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnumValues Enumerates the set of values for WithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum
func GetWithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnumValues() []WithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum {
	values := make([]WithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum, 0)
	for _, v := range mappingWithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnumStringValues Enumerates the set of values in String for WithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum
func GetWithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnumStringValues() []string {
	return []string{
		"VIRTUAL_CIRCUIT",
		"IPSEC_TUNNEL",
	}
}

// GetMappingWithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum(val string) (WithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum, bool) {
	enum, ok := mappingWithdrawPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
