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

// AnyNetworkSource Matches all network source types, including those added to this API in the future.
type AnyNetworkSource struct {

	// The network type to match.
	NetworkSourceType AnyNetworkSourceNetworkSourceTypeEnum `mandatory:"true" json:"networkSourceType"`

	// The network traffic must originate from the specified IP range, expressed in CIDR notation, to match.
	// Currently, only IPv4 addresses are supported.
	SourceIpAddress *string `mandatory:"false" json:"sourceIpAddress"`
}

func (m AnyNetworkSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AnyNetworkSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAnyNetworkSourceNetworkSourceTypeEnum(string(m.NetworkSourceType)); !ok && m.NetworkSourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetworkSourceType: %s. Supported values are: %s.", m.NetworkSourceType, strings.Join(GetAnyNetworkSourceNetworkSourceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AnyNetworkSourceNetworkSourceTypeEnum Enum with underlying type: string
type AnyNetworkSourceNetworkSourceTypeEnum string

// Set of constants representing the allowable values for AnyNetworkSourceNetworkSourceTypeEnum
const (
	AnyNetworkSourceNetworkSourceTypeVcn      AnyNetworkSourceNetworkSourceTypeEnum = "VCN"
	AnyNetworkSourceNetworkSourceTypePe       AnyNetworkSourceNetworkSourceTypeEnum = "PE"
	AnyNetworkSourceNetworkSourceTypeInternet AnyNetworkSourceNetworkSourceTypeEnum = "INTERNET"
	AnyNetworkSourceNetworkSourceTypeSgw      AnyNetworkSourceNetworkSourceTypeEnum = "SGW"
	AnyNetworkSourceNetworkSourceTypeAny      AnyNetworkSourceNetworkSourceTypeEnum = "ANY"
)

var mappingAnyNetworkSourceNetworkSourceTypeEnum = map[string]AnyNetworkSourceNetworkSourceTypeEnum{
	"VCN":      AnyNetworkSourceNetworkSourceTypeVcn,
	"PE":       AnyNetworkSourceNetworkSourceTypePe,
	"INTERNET": AnyNetworkSourceNetworkSourceTypeInternet,
	"SGW":      AnyNetworkSourceNetworkSourceTypeSgw,
	"ANY":      AnyNetworkSourceNetworkSourceTypeAny,
}

var mappingAnyNetworkSourceNetworkSourceTypeEnumLowerCase = map[string]AnyNetworkSourceNetworkSourceTypeEnum{
	"vcn":      AnyNetworkSourceNetworkSourceTypeVcn,
	"pe":       AnyNetworkSourceNetworkSourceTypePe,
	"internet": AnyNetworkSourceNetworkSourceTypeInternet,
	"sgw":      AnyNetworkSourceNetworkSourceTypeSgw,
	"any":      AnyNetworkSourceNetworkSourceTypeAny,
}

// GetAnyNetworkSourceNetworkSourceTypeEnumValues Enumerates the set of values for AnyNetworkSourceNetworkSourceTypeEnum
func GetAnyNetworkSourceNetworkSourceTypeEnumValues() []AnyNetworkSourceNetworkSourceTypeEnum {
	values := make([]AnyNetworkSourceNetworkSourceTypeEnum, 0)
	for _, v := range mappingAnyNetworkSourceNetworkSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAnyNetworkSourceNetworkSourceTypeEnumStringValues Enumerates the set of values in String for AnyNetworkSourceNetworkSourceTypeEnum
func GetAnyNetworkSourceNetworkSourceTypeEnumStringValues() []string {
	return []string{
		"VCN",
		"PE",
		"INTERNET",
		"SGW",
		"ANY",
	}
}

// GetMappingAnyNetworkSourceNetworkSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAnyNetworkSourceNetworkSourceTypeEnum(val string) (AnyNetworkSourceNetworkSourceTypeEnum, bool) {
	enum, ok := mappingAnyNetworkSourceNetworkSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
