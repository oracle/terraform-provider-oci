// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InternalPrivateIpMapping ptivate Ip Mapping
type InternalPrivateIpMapping struct {

	// Unique identifier of the VNIC the floating private IP is or was mapped to
	VnicId *string `mandatory:"true" json:"vnicId"`

	// The current state of the floating private IP to VNIC mapping
	Status InternalPrivateIpMappingStatusEnum `mandatory:"true" json:"status"`

	// The Network Address Translated IP to communicate with internal services
	NatIp *string `mandatory:"false" json:"natIp"`
}

func (m InternalPrivateIpMapping) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternalPrivateIpMapping) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInternalPrivateIpMappingStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetInternalPrivateIpMappingStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InternalPrivateIpMappingStatusEnum Enum with underlying type: string
type InternalPrivateIpMappingStatusEnum string

// Set of constants representing the allowable values for InternalPrivateIpMappingStatusEnum
const (
	InternalPrivateIpMappingStatusProvisioning InternalPrivateIpMappingStatusEnum = "PROVISIONING"
	InternalPrivateIpMappingStatusAvailable    InternalPrivateIpMappingStatusEnum = "AVAILABLE"
	InternalPrivateIpMappingStatusTerminating  InternalPrivateIpMappingStatusEnum = "TERMINATING"
)

var mappingInternalPrivateIpMappingStatusEnum = map[string]InternalPrivateIpMappingStatusEnum{
	"PROVISIONING": InternalPrivateIpMappingStatusProvisioning,
	"AVAILABLE":    InternalPrivateIpMappingStatusAvailable,
	"TERMINATING":  InternalPrivateIpMappingStatusTerminating,
}

var mappingInternalPrivateIpMappingStatusEnumLowerCase = map[string]InternalPrivateIpMappingStatusEnum{
	"provisioning": InternalPrivateIpMappingStatusProvisioning,
	"available":    InternalPrivateIpMappingStatusAvailable,
	"terminating":  InternalPrivateIpMappingStatusTerminating,
}

// GetInternalPrivateIpMappingStatusEnumValues Enumerates the set of values for InternalPrivateIpMappingStatusEnum
func GetInternalPrivateIpMappingStatusEnumValues() []InternalPrivateIpMappingStatusEnum {
	values := make([]InternalPrivateIpMappingStatusEnum, 0)
	for _, v := range mappingInternalPrivateIpMappingStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetInternalPrivateIpMappingStatusEnumStringValues Enumerates the set of values in String for InternalPrivateIpMappingStatusEnum
func GetInternalPrivateIpMappingStatusEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
	}
}

// GetMappingInternalPrivateIpMappingStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternalPrivateIpMappingStatusEnum(val string) (InternalPrivateIpMappingStatusEnum, bool) {
	enum, ok := mappingInternalPrivateIpMappingStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
