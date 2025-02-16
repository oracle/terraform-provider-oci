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

// LabelInfo Object to define source or destination label info.
type LabelInfo struct {
	Label *LabelDetails `mandatory:"false" json:"label"`

	// CIDR of source or destination label.
	Cidr *string `mandatory:"false" json:"cidr"`

	// The type of cidr.
	// CIDR_BLOCK - Signals that the specified cidr is a ipv6 / ipv4 cidr.
	// SERVICE_CIDR_BLOCK - Signals that the specified cidr is an OSN cidr block.
	// ALL_CIDR_BLOCK - Signals that the specified cidr is for all (ipv4 and ipv6) ip addresses.
	CidrType LabelInfoCidrTypeEnum `mandatory:"false" json:"cidrType,omitempty"`
}

func (m LabelInfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LabelInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLabelInfoCidrTypeEnum(string(m.CidrType)); !ok && m.CidrType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CidrType: %s. Supported values are: %s.", m.CidrType, strings.Join(GetLabelInfoCidrTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LabelInfoCidrTypeEnum Enum with underlying type: string
type LabelInfoCidrTypeEnum string

// Set of constants representing the allowable values for LabelInfoCidrTypeEnum
const (
	LabelInfoCidrTypeCidrBlock        LabelInfoCidrTypeEnum = "CIDR_BLOCK"
	LabelInfoCidrTypeServiceCidrBlock LabelInfoCidrTypeEnum = "SERVICE_CIDR_BLOCK"
	LabelInfoCidrTypeAllCidrBlock     LabelInfoCidrTypeEnum = "ALL_CIDR_BLOCK"
)

var mappingLabelInfoCidrTypeEnum = map[string]LabelInfoCidrTypeEnum{
	"CIDR_BLOCK":         LabelInfoCidrTypeCidrBlock,
	"SERVICE_CIDR_BLOCK": LabelInfoCidrTypeServiceCidrBlock,
	"ALL_CIDR_BLOCK":     LabelInfoCidrTypeAllCidrBlock,
}

var mappingLabelInfoCidrTypeEnumLowerCase = map[string]LabelInfoCidrTypeEnum{
	"cidr_block":         LabelInfoCidrTypeCidrBlock,
	"service_cidr_block": LabelInfoCidrTypeServiceCidrBlock,
	"all_cidr_block":     LabelInfoCidrTypeAllCidrBlock,
}

// GetLabelInfoCidrTypeEnumValues Enumerates the set of values for LabelInfoCidrTypeEnum
func GetLabelInfoCidrTypeEnumValues() []LabelInfoCidrTypeEnum {
	values := make([]LabelInfoCidrTypeEnum, 0)
	for _, v := range mappingLabelInfoCidrTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLabelInfoCidrTypeEnumStringValues Enumerates the set of values in String for LabelInfoCidrTypeEnum
func GetLabelInfoCidrTypeEnumStringValues() []string {
	return []string{
		"CIDR_BLOCK",
		"SERVICE_CIDR_BLOCK",
		"ALL_CIDR_BLOCK",
	}
}

// GetMappingLabelInfoCidrTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLabelInfoCidrTypeEnum(val string) (LabelInfoCidrTypeEnum, bool) {
	enum, ok := mappingLabelInfoCidrTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
