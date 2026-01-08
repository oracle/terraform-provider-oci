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

// InternetNetworkSource Matches public internet traffic from a range of IP addresses.
type InternetNetworkSource struct {

	// The network type to match.
	NetworkSourceType InternetNetworkSourceNetworkSourceTypeEnum `mandatory:"true" json:"networkSourceType"`

	// The network traffic must originate from the specified IP range, expressed in CIDR notation, to match.
	// Currently, only IPv4 addresses are supported.
	SourceIpAddress *string `mandatory:"true" json:"sourceIpAddress"`
}

func (m InternetNetworkSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternetNetworkSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInternetNetworkSourceNetworkSourceTypeEnum(string(m.NetworkSourceType)); !ok && m.NetworkSourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetworkSourceType: %s. Supported values are: %s.", m.NetworkSourceType, strings.Join(GetInternetNetworkSourceNetworkSourceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InternetNetworkSourceNetworkSourceTypeEnum Enum with underlying type: string
type InternetNetworkSourceNetworkSourceTypeEnum string

// Set of constants representing the allowable values for InternetNetworkSourceNetworkSourceTypeEnum
const (
	InternetNetworkSourceNetworkSourceTypeVcn      InternetNetworkSourceNetworkSourceTypeEnum = "VCN"
	InternetNetworkSourceNetworkSourceTypePe       InternetNetworkSourceNetworkSourceTypeEnum = "PE"
	InternetNetworkSourceNetworkSourceTypeInternet InternetNetworkSourceNetworkSourceTypeEnum = "INTERNET"
	InternetNetworkSourceNetworkSourceTypeSgw      InternetNetworkSourceNetworkSourceTypeEnum = "SGW"
	InternetNetworkSourceNetworkSourceTypeAny      InternetNetworkSourceNetworkSourceTypeEnum = "ANY"
)

var mappingInternetNetworkSourceNetworkSourceTypeEnum = map[string]InternetNetworkSourceNetworkSourceTypeEnum{
	"VCN":      InternetNetworkSourceNetworkSourceTypeVcn,
	"PE":       InternetNetworkSourceNetworkSourceTypePe,
	"INTERNET": InternetNetworkSourceNetworkSourceTypeInternet,
	"SGW":      InternetNetworkSourceNetworkSourceTypeSgw,
	"ANY":      InternetNetworkSourceNetworkSourceTypeAny,
}

var mappingInternetNetworkSourceNetworkSourceTypeEnumLowerCase = map[string]InternetNetworkSourceNetworkSourceTypeEnum{
	"vcn":      InternetNetworkSourceNetworkSourceTypeVcn,
	"pe":       InternetNetworkSourceNetworkSourceTypePe,
	"internet": InternetNetworkSourceNetworkSourceTypeInternet,
	"sgw":      InternetNetworkSourceNetworkSourceTypeSgw,
	"any":      InternetNetworkSourceNetworkSourceTypeAny,
}

// GetInternetNetworkSourceNetworkSourceTypeEnumValues Enumerates the set of values for InternetNetworkSourceNetworkSourceTypeEnum
func GetInternetNetworkSourceNetworkSourceTypeEnumValues() []InternetNetworkSourceNetworkSourceTypeEnum {
	values := make([]InternetNetworkSourceNetworkSourceTypeEnum, 0)
	for _, v := range mappingInternetNetworkSourceNetworkSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInternetNetworkSourceNetworkSourceTypeEnumStringValues Enumerates the set of values in String for InternetNetworkSourceNetworkSourceTypeEnum
func GetInternetNetworkSourceNetworkSourceTypeEnumStringValues() []string {
	return []string{
		"VCN",
		"PE",
		"INTERNET",
		"SGW",
		"ANY",
	}
}

// GetMappingInternetNetworkSourceNetworkSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternetNetworkSourceNetworkSourceTypeEnum(val string) (InternetNetworkSourceNetworkSourceTypeEnum, bool) {
	enum, ok := mappingInternetNetworkSourceNetworkSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
