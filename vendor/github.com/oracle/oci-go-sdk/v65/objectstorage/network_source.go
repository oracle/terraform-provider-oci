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

// NetworkSource Matches requests originating from the specified network type in the same region where the ACL exists.
type NetworkSource struct {

	// The network type to match.
	NetworkSourceType NetworkSourceNetworkSourceTypeEnum `mandatory:"true" json:"networkSourceType"`
}

func (m NetworkSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NetworkSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNetworkSourceNetworkSourceTypeEnum(string(m.NetworkSourceType)); !ok && m.NetworkSourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetworkSourceType: %s. Supported values are: %s.", m.NetworkSourceType, strings.Join(GetNetworkSourceNetworkSourceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NetworkSourceNetworkSourceTypeEnum Enum with underlying type: string
type NetworkSourceNetworkSourceTypeEnum string

// Set of constants representing the allowable values for NetworkSourceNetworkSourceTypeEnum
const (
	NetworkSourceNetworkSourceTypeVcn      NetworkSourceNetworkSourceTypeEnum = "VCN"
	NetworkSourceNetworkSourceTypePe       NetworkSourceNetworkSourceTypeEnum = "PE"
	NetworkSourceNetworkSourceTypeInternet NetworkSourceNetworkSourceTypeEnum = "INTERNET"
	NetworkSourceNetworkSourceTypeSgw      NetworkSourceNetworkSourceTypeEnum = "SGW"
	NetworkSourceNetworkSourceTypeAny      NetworkSourceNetworkSourceTypeEnum = "ANY"
)

var mappingNetworkSourceNetworkSourceTypeEnum = map[string]NetworkSourceNetworkSourceTypeEnum{
	"VCN":      NetworkSourceNetworkSourceTypeVcn,
	"PE":       NetworkSourceNetworkSourceTypePe,
	"INTERNET": NetworkSourceNetworkSourceTypeInternet,
	"SGW":      NetworkSourceNetworkSourceTypeSgw,
	"ANY":      NetworkSourceNetworkSourceTypeAny,
}

var mappingNetworkSourceNetworkSourceTypeEnumLowerCase = map[string]NetworkSourceNetworkSourceTypeEnum{
	"vcn":      NetworkSourceNetworkSourceTypeVcn,
	"pe":       NetworkSourceNetworkSourceTypePe,
	"internet": NetworkSourceNetworkSourceTypeInternet,
	"sgw":      NetworkSourceNetworkSourceTypeSgw,
	"any":      NetworkSourceNetworkSourceTypeAny,
}

// GetNetworkSourceNetworkSourceTypeEnumValues Enumerates the set of values for NetworkSourceNetworkSourceTypeEnum
func GetNetworkSourceNetworkSourceTypeEnumValues() []NetworkSourceNetworkSourceTypeEnum {
	values := make([]NetworkSourceNetworkSourceTypeEnum, 0)
	for _, v := range mappingNetworkSourceNetworkSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkSourceNetworkSourceTypeEnumStringValues Enumerates the set of values in String for NetworkSourceNetworkSourceTypeEnum
func GetNetworkSourceNetworkSourceTypeEnumStringValues() []string {
	return []string{
		"VCN",
		"PE",
		"INTERNET",
		"SGW",
		"ANY",
	}
}

// GetMappingNetworkSourceNetworkSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkSourceNetworkSourceTypeEnum(val string) (NetworkSourceNetworkSourceTypeEnum, bool) {
	enum, ok := mappingNetworkSourceNetworkSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
