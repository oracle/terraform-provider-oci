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

// SendPerAttachmentExportPolicyUpdateDetails Details for sending per attachment export policy update
type SendPerAttachmentExportPolicyUpdateDetails struct {

	// The attachment type.
	AttachmentType SendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum `mandatory:"true" json:"attachmentType"`
}

func (m SendPerAttachmentExportPolicyUpdateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SendPerAttachmentExportPolicyUpdateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum(string(m.AttachmentType)); !ok && m.AttachmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttachmentType: %s. Supported values are: %s.", m.AttachmentType, strings.Join(GetSendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum Enum with underlying type: string
type SendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum string

// Set of constants representing the allowable values for SendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum
const (
	SendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeVirtualCircuit SendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum = "VIRTUAL_CIRCUIT"
	SendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeIpsecTunnel    SendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum = "IPSEC_TUNNEL"
)

var mappingSendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum = map[string]SendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum{
	"VIRTUAL_CIRCUIT": SendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeVirtualCircuit,
	"IPSEC_TUNNEL":    SendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeIpsecTunnel,
}

var mappingSendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnumLowerCase = map[string]SendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum{
	"virtual_circuit": SendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeVirtualCircuit,
	"ipsec_tunnel":    SendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeIpsecTunnel,
}

// GetSendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnumValues Enumerates the set of values for SendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum
func GetSendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnumValues() []SendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum {
	values := make([]SendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum, 0)
	for _, v := range mappingSendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnumStringValues Enumerates the set of values in String for SendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum
func GetSendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnumStringValues() []string {
	return []string{
		"VIRTUAL_CIRCUIT",
		"IPSEC_TUNNEL",
	}
}

// GetMappingSendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum(val string) (SendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnum, bool) {
	enum, ok := mappingSendPerAttachmentExportPolicyUpdateDetailsAttachmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
