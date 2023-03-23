// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Use Object Storage and Archive Storage APIs to manage buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.cloud.oracle.com/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NetworkSource Matches requests originating from the specified network type in the same region where the ACL exists.
type NetworkSource interface {
}

type networksource struct {
	JsonData          []byte
	NetworkSourceType string `json:"networkSourceType"`
}

// UnmarshalJSON unmarshals json
func (m *networksource) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalernetworksource networksource
	s := struct {
		Model Unmarshalernetworksource
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.NetworkSourceType = s.Model.NetworkSourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *networksource) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.NetworkSourceType {
	case "SGW":
		mm := SgwNetworkSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ANY":
		mm := AnyNetworkSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VCN":
		mm := VcnNetworkSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PE":
		mm := PeNetworkSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INTERNET":
		mm := InternetNetworkSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for NetworkSource: %s.", m.NetworkSourceType)
		return *m, nil
	}
}

func (m networksource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m networksource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
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
