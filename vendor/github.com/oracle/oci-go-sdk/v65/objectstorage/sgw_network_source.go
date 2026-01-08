// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Use Object Storage and Archive Storage APIs to manage buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.oracle.com/iaas/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.oracle.com/iaas/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SgwNetworkSource Matches a specific Service Gateway, or a set of Service Gateways.
type SgwNetworkSource struct {

	// The network type to match.
	NetworkSourceType SgwNetworkSourceNetworkSourceTypeEnum `mandatory:"true" json:"networkSourceType"`

	// The ID of the SGW to match, or "ALL" to match all SGWs in the specified compartment.
	SgwId *string `mandatory:"true" json:"sgwId"`

	// The SGW must exist in the specified compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`
}

func (m SgwNetworkSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SgwNetworkSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSgwNetworkSourceNetworkSourceTypeEnum(string(m.NetworkSourceType)); !ok && m.NetworkSourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetworkSourceType: %s. Supported values are: %s.", m.NetworkSourceType, strings.Join(GetSgwNetworkSourceNetworkSourceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SgwNetworkSourceNetworkSourceTypeEnum Enum with underlying type: string
type SgwNetworkSourceNetworkSourceTypeEnum string

// Set of constants representing the allowable values for SgwNetworkSourceNetworkSourceTypeEnum
const (
	SgwNetworkSourceNetworkSourceTypeVcn      SgwNetworkSourceNetworkSourceTypeEnum = "VCN"
	SgwNetworkSourceNetworkSourceTypePe       SgwNetworkSourceNetworkSourceTypeEnum = "PE"
	SgwNetworkSourceNetworkSourceTypeInternet SgwNetworkSourceNetworkSourceTypeEnum = "INTERNET"
	SgwNetworkSourceNetworkSourceTypeSgw      SgwNetworkSourceNetworkSourceTypeEnum = "SGW"
	SgwNetworkSourceNetworkSourceTypeAny      SgwNetworkSourceNetworkSourceTypeEnum = "ANY"
)

var mappingSgwNetworkSourceNetworkSourceTypeEnum = map[string]SgwNetworkSourceNetworkSourceTypeEnum{
	"VCN":      SgwNetworkSourceNetworkSourceTypeVcn,
	"PE":       SgwNetworkSourceNetworkSourceTypePe,
	"INTERNET": SgwNetworkSourceNetworkSourceTypeInternet,
	"SGW":      SgwNetworkSourceNetworkSourceTypeSgw,
	"ANY":      SgwNetworkSourceNetworkSourceTypeAny,
}

var mappingSgwNetworkSourceNetworkSourceTypeEnumLowerCase = map[string]SgwNetworkSourceNetworkSourceTypeEnum{
	"vcn":      SgwNetworkSourceNetworkSourceTypeVcn,
	"pe":       SgwNetworkSourceNetworkSourceTypePe,
	"internet": SgwNetworkSourceNetworkSourceTypeInternet,
	"sgw":      SgwNetworkSourceNetworkSourceTypeSgw,
	"any":      SgwNetworkSourceNetworkSourceTypeAny,
}

// GetSgwNetworkSourceNetworkSourceTypeEnumValues Enumerates the set of values for SgwNetworkSourceNetworkSourceTypeEnum
func GetSgwNetworkSourceNetworkSourceTypeEnumValues() []SgwNetworkSourceNetworkSourceTypeEnum {
	values := make([]SgwNetworkSourceNetworkSourceTypeEnum, 0)
	for _, v := range mappingSgwNetworkSourceNetworkSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSgwNetworkSourceNetworkSourceTypeEnumStringValues Enumerates the set of values in String for SgwNetworkSourceNetworkSourceTypeEnum
func GetSgwNetworkSourceNetworkSourceTypeEnumStringValues() []string {
	return []string{
		"VCN",
		"PE",
		"INTERNET",
		"SGW",
		"ANY",
	}
}

// GetMappingSgwNetworkSourceNetworkSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSgwNetworkSourceNetworkSourceTypeEnum(val string) (SgwNetworkSourceNetworkSourceTypeEnum, bool) {
	enum, ok := mappingSgwNetworkSourceNetworkSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
