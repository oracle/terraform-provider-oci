// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// SendCommonDrgExportPolicyUpdateDetails Details for Sending common export policy update
type SendCommonDrgExportPolicyUpdateDetails struct {

	// The attachment type.
	AttachmentType SendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnum `mandatory:"true" json:"attachmentType"`
}

func (m SendCommonDrgExportPolicyUpdateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SendCommonDrgExportPolicyUpdateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnum(string(m.AttachmentType)); !ok && m.AttachmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttachmentType: %s. Supported values are: %s.", m.AttachmentType, strings.Join(GetSendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnum Enum with underlying type: string
type SendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnum string

// Set of constants representing the allowable values for SendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnum
const (
	SendCommonDrgExportPolicyUpdateDetailsAttachmentTypeVirtualCircuit SendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnum = "VIRTUAL_CIRCUIT"
	SendCommonDrgExportPolicyUpdateDetailsAttachmentTypeIpsecTunnel    SendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnum = "IPSEC_TUNNEL"
)

var mappingSendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnum = map[string]SendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnum{
	"VIRTUAL_CIRCUIT": SendCommonDrgExportPolicyUpdateDetailsAttachmentTypeVirtualCircuit,
	"IPSEC_TUNNEL":    SendCommonDrgExportPolicyUpdateDetailsAttachmentTypeIpsecTunnel,
}

var mappingSendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnumLowerCase = map[string]SendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnum{
	"virtual_circuit": SendCommonDrgExportPolicyUpdateDetailsAttachmentTypeVirtualCircuit,
	"ipsec_tunnel":    SendCommonDrgExportPolicyUpdateDetailsAttachmentTypeIpsecTunnel,
}

// GetSendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnumValues Enumerates the set of values for SendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnum
func GetSendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnumValues() []SendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnum {
	values := make([]SendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnum, 0)
	for _, v := range mappingSendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnumStringValues Enumerates the set of values in String for SendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnum
func GetSendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnumStringValues() []string {
	return []string{
		"VIRTUAL_CIRCUIT",
		"IPSEC_TUNNEL",
	}
}

// GetMappingSendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnum(val string) (SendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnum, bool) {
	enum, ok := mappingSendCommonDrgExportPolicyUpdateDetailsAttachmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
