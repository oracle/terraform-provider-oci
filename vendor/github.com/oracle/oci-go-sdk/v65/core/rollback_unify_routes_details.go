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

// RollbackUnifyRoutesDetails Details for rolling back unify routes
type RollbackUnifyRoutesDetails struct {

	// The attachment type.
	AttachmentType RollbackUnifyRoutesDetailsAttachmentTypeEnum `mandatory:"true" json:"attachmentType"`
}

func (m RollbackUnifyRoutesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RollbackUnifyRoutesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRollbackUnifyRoutesDetailsAttachmentTypeEnum(string(m.AttachmentType)); !ok && m.AttachmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttachmentType: %s. Supported values are: %s.", m.AttachmentType, strings.Join(GetRollbackUnifyRoutesDetailsAttachmentTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RollbackUnifyRoutesDetailsAttachmentTypeEnum Enum with underlying type: string
type RollbackUnifyRoutesDetailsAttachmentTypeEnum string

// Set of constants representing the allowable values for RollbackUnifyRoutesDetailsAttachmentTypeEnum
const (
	RollbackUnifyRoutesDetailsAttachmentTypeVirtualCircuit RollbackUnifyRoutesDetailsAttachmentTypeEnum = "VIRTUAL_CIRCUIT"
	RollbackUnifyRoutesDetailsAttachmentTypeIpsecTunnel    RollbackUnifyRoutesDetailsAttachmentTypeEnum = "IPSEC_TUNNEL"
)

var mappingRollbackUnifyRoutesDetailsAttachmentTypeEnum = map[string]RollbackUnifyRoutesDetailsAttachmentTypeEnum{
	"VIRTUAL_CIRCUIT": RollbackUnifyRoutesDetailsAttachmentTypeVirtualCircuit,
	"IPSEC_TUNNEL":    RollbackUnifyRoutesDetailsAttachmentTypeIpsecTunnel,
}

var mappingRollbackUnifyRoutesDetailsAttachmentTypeEnumLowerCase = map[string]RollbackUnifyRoutesDetailsAttachmentTypeEnum{
	"virtual_circuit": RollbackUnifyRoutesDetailsAttachmentTypeVirtualCircuit,
	"ipsec_tunnel":    RollbackUnifyRoutesDetailsAttachmentTypeIpsecTunnel,
}

// GetRollbackUnifyRoutesDetailsAttachmentTypeEnumValues Enumerates the set of values for RollbackUnifyRoutesDetailsAttachmentTypeEnum
func GetRollbackUnifyRoutesDetailsAttachmentTypeEnumValues() []RollbackUnifyRoutesDetailsAttachmentTypeEnum {
	values := make([]RollbackUnifyRoutesDetailsAttachmentTypeEnum, 0)
	for _, v := range mappingRollbackUnifyRoutesDetailsAttachmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRollbackUnifyRoutesDetailsAttachmentTypeEnumStringValues Enumerates the set of values in String for RollbackUnifyRoutesDetailsAttachmentTypeEnum
func GetRollbackUnifyRoutesDetailsAttachmentTypeEnumStringValues() []string {
	return []string{
		"VIRTUAL_CIRCUIT",
		"IPSEC_TUNNEL",
	}
}

// GetMappingRollbackUnifyRoutesDetailsAttachmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRollbackUnifyRoutesDetailsAttachmentTypeEnum(val string) (RollbackUnifyRoutesDetailsAttachmentTypeEnum, bool) {
	enum, ok := mappingRollbackUnifyRoutesDetailsAttachmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
