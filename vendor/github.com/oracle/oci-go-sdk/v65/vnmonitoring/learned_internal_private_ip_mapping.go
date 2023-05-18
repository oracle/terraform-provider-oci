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

// LearnedInternalPrivateIpMapping Learned internal private IP mapping
type LearnedInternalPrivateIpMapping struct {

	// Unique identifier of the VNIC the private IP is mapped to
	VnicId *string `mandatory:"true" json:"vnicId"`

	// Unique identifier of the private IP mapped to the VNIC
	PrivateIpId *string `mandatory:"true" json:"privateIpId"`

	// The current state of the private IP to VNIC mapping
	Status LearnedInternalPrivateIpMappingStatusEnum `mandatory:"true" json:"status"`
}

func (m LearnedInternalPrivateIpMapping) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LearnedInternalPrivateIpMapping) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLearnedInternalPrivateIpMappingStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetLearnedInternalPrivateIpMappingStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LearnedInternalPrivateIpMappingStatusEnum Enum with underlying type: string
type LearnedInternalPrivateIpMappingStatusEnum string

// Set of constants representing the allowable values for LearnedInternalPrivateIpMappingStatusEnum
const (
	LearnedInternalPrivateIpMappingStatusProvisioning LearnedInternalPrivateIpMappingStatusEnum = "PROVISIONING"
	LearnedInternalPrivateIpMappingStatusAvailable    LearnedInternalPrivateIpMappingStatusEnum = "AVAILABLE"
	LearnedInternalPrivateIpMappingStatusTerminating  LearnedInternalPrivateIpMappingStatusEnum = "TERMINATING"
)

var mappingLearnedInternalPrivateIpMappingStatusEnum = map[string]LearnedInternalPrivateIpMappingStatusEnum{
	"PROVISIONING": LearnedInternalPrivateIpMappingStatusProvisioning,
	"AVAILABLE":    LearnedInternalPrivateIpMappingStatusAvailable,
	"TERMINATING":  LearnedInternalPrivateIpMappingStatusTerminating,
}

var mappingLearnedInternalPrivateIpMappingStatusEnumLowerCase = map[string]LearnedInternalPrivateIpMappingStatusEnum{
	"provisioning": LearnedInternalPrivateIpMappingStatusProvisioning,
	"available":    LearnedInternalPrivateIpMappingStatusAvailable,
	"terminating":  LearnedInternalPrivateIpMappingStatusTerminating,
}

// GetLearnedInternalPrivateIpMappingStatusEnumValues Enumerates the set of values for LearnedInternalPrivateIpMappingStatusEnum
func GetLearnedInternalPrivateIpMappingStatusEnumValues() []LearnedInternalPrivateIpMappingStatusEnum {
	values := make([]LearnedInternalPrivateIpMappingStatusEnum, 0)
	for _, v := range mappingLearnedInternalPrivateIpMappingStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetLearnedInternalPrivateIpMappingStatusEnumStringValues Enumerates the set of values in String for LearnedInternalPrivateIpMappingStatusEnum
func GetLearnedInternalPrivateIpMappingStatusEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
	}
}

// GetMappingLearnedInternalPrivateIpMappingStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLearnedInternalPrivateIpMappingStatusEnum(val string) (LearnedInternalPrivateIpMappingStatusEnum, bool) {
	enum, ok := mappingLearnedInternalPrivateIpMappingStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
