// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Multicloud API
//
// Use the Oracle Multicloud API to retrieve resource anchors and network anchors, and the metadata mappings related a Cloud Service Provider. For more information, see <link to docs>.
//

package multicloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OciNetworkSubnet Oracle Cloud Infrastructure network subnet object.
type OciNetworkSubnet struct {

	// Defines if the subnet is the primary or backup for the network
	Type OciNetworkSubnetTypeEnum `mandatory:"true" json:"type"`

	// OCID for existing the subnet. CSP can not set this property.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// Subnet label. CSP can set this property
	Label *string `mandatory:"false" json:"label"`
}

func (m OciNetworkSubnet) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciNetworkSubnet) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOciNetworkSubnetTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetOciNetworkSubnetTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OciNetworkSubnetTypeEnum Enum with underlying type: string
type OciNetworkSubnetTypeEnum string

// Set of constants representing the allowable values for OciNetworkSubnetTypeEnum
const (
	OciNetworkSubnetTypeClient OciNetworkSubnetTypeEnum = "CLIENT"
	OciNetworkSubnetTypeBackup OciNetworkSubnetTypeEnum = "BACKUP"
)

var mappingOciNetworkSubnetTypeEnum = map[string]OciNetworkSubnetTypeEnum{
	"CLIENT": OciNetworkSubnetTypeClient,
	"BACKUP": OciNetworkSubnetTypeBackup,
}

var mappingOciNetworkSubnetTypeEnumLowerCase = map[string]OciNetworkSubnetTypeEnum{
	"client": OciNetworkSubnetTypeClient,
	"backup": OciNetworkSubnetTypeBackup,
}

// GetOciNetworkSubnetTypeEnumValues Enumerates the set of values for OciNetworkSubnetTypeEnum
func GetOciNetworkSubnetTypeEnumValues() []OciNetworkSubnetTypeEnum {
	values := make([]OciNetworkSubnetTypeEnum, 0)
	for _, v := range mappingOciNetworkSubnetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOciNetworkSubnetTypeEnumStringValues Enumerates the set of values in String for OciNetworkSubnetTypeEnum
func GetOciNetworkSubnetTypeEnumStringValues() []string {
	return []string{
		"CLIENT",
		"BACKUP",
	}
}

// GetMappingOciNetworkSubnetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOciNetworkSubnetTypeEnum(val string) (OciNetworkSubnetTypeEnum, bool) {
	enum, ok := mappingOciNetworkSubnetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
