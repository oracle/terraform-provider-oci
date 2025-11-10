// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// VcnNetworkSource Matches a specific Virtual Cloud Network, or a set of Virtual Cloud Networks.
type VcnNetworkSource struct {

	// The network type to match.
	NetworkSourceType VcnNetworkSourceNetworkSourceTypeEnum `mandatory:"true" json:"networkSourceType"`

	// The ID of the VCN to match, or "ALL" to match all VCNs in the specified compartment.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The VCN must exist in the specified compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The network traffic must originate from the specified IP range, expressed in CIDR notation, to match.
	// Currently, only IPv4 addresses are supported.
	SourceIpAddress *string `mandatory:"false" json:"sourceIpAddress"`
}

func (m VcnNetworkSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VcnNetworkSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVcnNetworkSourceNetworkSourceTypeEnum(string(m.NetworkSourceType)); !ok && m.NetworkSourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetworkSourceType: %s. Supported values are: %s.", m.NetworkSourceType, strings.Join(GetVcnNetworkSourceNetworkSourceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VcnNetworkSourceNetworkSourceTypeEnum Enum with underlying type: string
type VcnNetworkSourceNetworkSourceTypeEnum string

// Set of constants representing the allowable values for VcnNetworkSourceNetworkSourceTypeEnum
const (
	VcnNetworkSourceNetworkSourceTypeVcn      VcnNetworkSourceNetworkSourceTypeEnum = "VCN"
	VcnNetworkSourceNetworkSourceTypePe       VcnNetworkSourceNetworkSourceTypeEnum = "PE"
	VcnNetworkSourceNetworkSourceTypeInternet VcnNetworkSourceNetworkSourceTypeEnum = "INTERNET"
	VcnNetworkSourceNetworkSourceTypeSgw      VcnNetworkSourceNetworkSourceTypeEnum = "SGW"
	VcnNetworkSourceNetworkSourceTypeAny      VcnNetworkSourceNetworkSourceTypeEnum = "ANY"
)

var mappingVcnNetworkSourceNetworkSourceTypeEnum = map[string]VcnNetworkSourceNetworkSourceTypeEnum{
	"VCN":      VcnNetworkSourceNetworkSourceTypeVcn,
	"PE":       VcnNetworkSourceNetworkSourceTypePe,
	"INTERNET": VcnNetworkSourceNetworkSourceTypeInternet,
	"SGW":      VcnNetworkSourceNetworkSourceTypeSgw,
	"ANY":      VcnNetworkSourceNetworkSourceTypeAny,
}

var mappingVcnNetworkSourceNetworkSourceTypeEnumLowerCase = map[string]VcnNetworkSourceNetworkSourceTypeEnum{
	"vcn":      VcnNetworkSourceNetworkSourceTypeVcn,
	"pe":       VcnNetworkSourceNetworkSourceTypePe,
	"internet": VcnNetworkSourceNetworkSourceTypeInternet,
	"sgw":      VcnNetworkSourceNetworkSourceTypeSgw,
	"any":      VcnNetworkSourceNetworkSourceTypeAny,
}

// GetVcnNetworkSourceNetworkSourceTypeEnumValues Enumerates the set of values for VcnNetworkSourceNetworkSourceTypeEnum
func GetVcnNetworkSourceNetworkSourceTypeEnumValues() []VcnNetworkSourceNetworkSourceTypeEnum {
	values := make([]VcnNetworkSourceNetworkSourceTypeEnum, 0)
	for _, v := range mappingVcnNetworkSourceNetworkSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetVcnNetworkSourceNetworkSourceTypeEnumStringValues Enumerates the set of values in String for VcnNetworkSourceNetworkSourceTypeEnum
func GetVcnNetworkSourceNetworkSourceTypeEnumStringValues() []string {
	return []string{
		"VCN",
		"PE",
		"INTERNET",
		"SGW",
		"ANY",
	}
}

// GetMappingVcnNetworkSourceNetworkSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVcnNetworkSourceNetworkSourceTypeEnum(val string) (VcnNetworkSourceNetworkSourceTypeEnum, bool) {
	enum, ok := mappingVcnNetworkSourceNetworkSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
