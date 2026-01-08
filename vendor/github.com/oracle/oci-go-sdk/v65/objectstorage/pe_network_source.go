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

// PeNetworkSource Matches a specific Private Endpoint, or a set of Private Endpoints.
type PeNetworkSource struct {

	// The network type to match.
	NetworkSourceType PeNetworkSourceNetworkSourceTypeEnum `mandatory:"true" json:"networkSourceType"`

	// The ID of the PE to match, or "ALL" to match all PEs in the specified compartment.
	PeId *string `mandatory:"true" json:"peId"`

	// The PE must exist in the specified compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The network traffic must originate from the specified IP range, expressed in CIDR notation, to match.
	// Currently, only IPv4 addresses are supported.
	SourceIpAddress *string `mandatory:"false" json:"sourceIpAddress"`
}

func (m PeNetworkSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PeNetworkSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPeNetworkSourceNetworkSourceTypeEnum(string(m.NetworkSourceType)); !ok && m.NetworkSourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetworkSourceType: %s. Supported values are: %s.", m.NetworkSourceType, strings.Join(GetPeNetworkSourceNetworkSourceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PeNetworkSourceNetworkSourceTypeEnum Enum with underlying type: string
type PeNetworkSourceNetworkSourceTypeEnum string

// Set of constants representing the allowable values for PeNetworkSourceNetworkSourceTypeEnum
const (
	PeNetworkSourceNetworkSourceTypeVcn      PeNetworkSourceNetworkSourceTypeEnum = "VCN"
	PeNetworkSourceNetworkSourceTypePe       PeNetworkSourceNetworkSourceTypeEnum = "PE"
	PeNetworkSourceNetworkSourceTypeInternet PeNetworkSourceNetworkSourceTypeEnum = "INTERNET"
	PeNetworkSourceNetworkSourceTypeSgw      PeNetworkSourceNetworkSourceTypeEnum = "SGW"
	PeNetworkSourceNetworkSourceTypeAny      PeNetworkSourceNetworkSourceTypeEnum = "ANY"
)

var mappingPeNetworkSourceNetworkSourceTypeEnum = map[string]PeNetworkSourceNetworkSourceTypeEnum{
	"VCN":      PeNetworkSourceNetworkSourceTypeVcn,
	"PE":       PeNetworkSourceNetworkSourceTypePe,
	"INTERNET": PeNetworkSourceNetworkSourceTypeInternet,
	"SGW":      PeNetworkSourceNetworkSourceTypeSgw,
	"ANY":      PeNetworkSourceNetworkSourceTypeAny,
}

var mappingPeNetworkSourceNetworkSourceTypeEnumLowerCase = map[string]PeNetworkSourceNetworkSourceTypeEnum{
	"vcn":      PeNetworkSourceNetworkSourceTypeVcn,
	"pe":       PeNetworkSourceNetworkSourceTypePe,
	"internet": PeNetworkSourceNetworkSourceTypeInternet,
	"sgw":      PeNetworkSourceNetworkSourceTypeSgw,
	"any":      PeNetworkSourceNetworkSourceTypeAny,
}

// GetPeNetworkSourceNetworkSourceTypeEnumValues Enumerates the set of values for PeNetworkSourceNetworkSourceTypeEnum
func GetPeNetworkSourceNetworkSourceTypeEnumValues() []PeNetworkSourceNetworkSourceTypeEnum {
	values := make([]PeNetworkSourceNetworkSourceTypeEnum, 0)
	for _, v := range mappingPeNetworkSourceNetworkSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPeNetworkSourceNetworkSourceTypeEnumStringValues Enumerates the set of values in String for PeNetworkSourceNetworkSourceTypeEnum
func GetPeNetworkSourceNetworkSourceTypeEnumStringValues() []string {
	return []string{
		"VCN",
		"PE",
		"INTERNET",
		"SGW",
		"ANY",
	}
}

// GetMappingPeNetworkSourceNetworkSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPeNetworkSourceNetworkSourceTypeEnum(val string) (PeNetworkSourceNetworkSourceTypeEnum, bool) {
	enum, ok := mappingPeNetworkSourceNetworkSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
