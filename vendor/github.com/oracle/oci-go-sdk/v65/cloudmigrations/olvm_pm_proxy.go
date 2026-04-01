// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OlvmPmProxy Power management proxy types.
type OlvmPmProxy struct {

	// Pm Proxy Type
	Type OlvmPmProxyTypeEnum `mandatory:"false" json:"type,omitempty"`
}

func (m OlvmPmProxy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmPmProxy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmPmProxyTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetOlvmPmProxyTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmPmProxyTypeEnum Enum with underlying type: string
type OlvmPmProxyTypeEnum string

// Set of constants representing the allowable values for OlvmPmProxyTypeEnum
const (
	OlvmPmProxyTypeCluster OlvmPmProxyTypeEnum = "CLUSTER"
	OlvmPmProxyTypeDc      OlvmPmProxyTypeEnum = "DC"
	OlvmPmProxyTypeOtherDc OlvmPmProxyTypeEnum = "OTHER_DC"
)

var mappingOlvmPmProxyTypeEnum = map[string]OlvmPmProxyTypeEnum{
	"CLUSTER":  OlvmPmProxyTypeCluster,
	"DC":       OlvmPmProxyTypeDc,
	"OTHER_DC": OlvmPmProxyTypeOtherDc,
}

var mappingOlvmPmProxyTypeEnumLowerCase = map[string]OlvmPmProxyTypeEnum{
	"cluster":  OlvmPmProxyTypeCluster,
	"dc":       OlvmPmProxyTypeDc,
	"other_dc": OlvmPmProxyTypeOtherDc,
}

// GetOlvmPmProxyTypeEnumValues Enumerates the set of values for OlvmPmProxyTypeEnum
func GetOlvmPmProxyTypeEnumValues() []OlvmPmProxyTypeEnum {
	values := make([]OlvmPmProxyTypeEnum, 0)
	for _, v := range mappingOlvmPmProxyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmPmProxyTypeEnumStringValues Enumerates the set of values in String for OlvmPmProxyTypeEnum
func GetOlvmPmProxyTypeEnumStringValues() []string {
	return []string{
		"CLUSTER",
		"DC",
		"OTHER_DC",
	}
}

// GetMappingOlvmPmProxyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmPmProxyTypeEnum(val string) (OlvmPmProxyTypeEnum, bool) {
	enum, ok := mappingOlvmPmProxyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
