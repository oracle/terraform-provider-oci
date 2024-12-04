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

// UnifyRoutesDetails Details for unifying routes
type UnifyRoutesDetails struct {

	// The attachment type.
	AttachmentType UnifyRoutesDetailsAttachmentTypeEnum `mandatory:"true" json:"attachmentType"`
}

func (m UnifyRoutesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnifyRoutesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUnifyRoutesDetailsAttachmentTypeEnum(string(m.AttachmentType)); !ok && m.AttachmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttachmentType: %s. Supported values are: %s.", m.AttachmentType, strings.Join(GetUnifyRoutesDetailsAttachmentTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnifyRoutesDetailsAttachmentTypeEnum Enum with underlying type: string
type UnifyRoutesDetailsAttachmentTypeEnum string

// Set of constants representing the allowable values for UnifyRoutesDetailsAttachmentTypeEnum
const (
	UnifyRoutesDetailsAttachmentTypeVirtualCircuit UnifyRoutesDetailsAttachmentTypeEnum = "VIRTUAL_CIRCUIT"
	UnifyRoutesDetailsAttachmentTypeIpsecTunnel    UnifyRoutesDetailsAttachmentTypeEnum = "IPSEC_TUNNEL"
)

var mappingUnifyRoutesDetailsAttachmentTypeEnum = map[string]UnifyRoutesDetailsAttachmentTypeEnum{
	"VIRTUAL_CIRCUIT": UnifyRoutesDetailsAttachmentTypeVirtualCircuit,
	"IPSEC_TUNNEL":    UnifyRoutesDetailsAttachmentTypeIpsecTunnel,
}

var mappingUnifyRoutesDetailsAttachmentTypeEnumLowerCase = map[string]UnifyRoutesDetailsAttachmentTypeEnum{
	"virtual_circuit": UnifyRoutesDetailsAttachmentTypeVirtualCircuit,
	"ipsec_tunnel":    UnifyRoutesDetailsAttachmentTypeIpsecTunnel,
}

// GetUnifyRoutesDetailsAttachmentTypeEnumValues Enumerates the set of values for UnifyRoutesDetailsAttachmentTypeEnum
func GetUnifyRoutesDetailsAttachmentTypeEnumValues() []UnifyRoutesDetailsAttachmentTypeEnum {
	values := make([]UnifyRoutesDetailsAttachmentTypeEnum, 0)
	for _, v := range mappingUnifyRoutesDetailsAttachmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUnifyRoutesDetailsAttachmentTypeEnumStringValues Enumerates the set of values in String for UnifyRoutesDetailsAttachmentTypeEnum
func GetUnifyRoutesDetailsAttachmentTypeEnumStringValues() []string {
	return []string{
		"VIRTUAL_CIRCUIT",
		"IPSEC_TUNNEL",
	}
}

// GetMappingUnifyRoutesDetailsAttachmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnifyRoutesDetailsAttachmentTypeEnum(val string) (UnifyRoutesDetailsAttachmentTypeEnum, bool) {
	enum, ok := mappingUnifyRoutesDetailsAttachmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
