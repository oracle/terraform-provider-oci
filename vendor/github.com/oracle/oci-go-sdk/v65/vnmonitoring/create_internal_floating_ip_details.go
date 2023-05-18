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

// CreateInternalFloatingIpDetails Details to create internal floating Ip
type CreateInternalFloatingIpDetails struct {

	// ID of the compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// User friendly name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// ID of the entity owning the floating IP
	OwnerId *string `mandatory:"false" json:"ownerId"`

	// OCID of the pool object created by the current tenancy
	PublicIpPoolId *string `mandatory:"false" json:"publicIpPoolId"`

	// Only provided when no publicIpPoolId is specified.
	InternalPoolName CreateInternalFloatingIpDetailsInternalPoolNameEnum `mandatory:"false" json:"internalPoolName,omitempty"`
}

func (m CreateInternalFloatingIpDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateInternalFloatingIpDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateInternalFloatingIpDetailsInternalPoolNameEnum(string(m.InternalPoolName)); !ok && m.InternalPoolName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InternalPoolName: %s. Supported values are: %s.", m.InternalPoolName, strings.Join(GetCreateInternalFloatingIpDetailsInternalPoolNameEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateInternalFloatingIpDetailsInternalPoolNameEnum Enum with underlying type: string
type CreateInternalFloatingIpDetailsInternalPoolNameEnum string

// Set of constants representing the allowable values for CreateInternalFloatingIpDetailsInternalPoolNameEnum
const (
	CreateInternalFloatingIpDetailsInternalPoolNameExternal   CreateInternalFloatingIpDetailsInternalPoolNameEnum = "EXTERNAL"
	CreateInternalFloatingIpDetailsInternalPoolNameSociEgress CreateInternalFloatingIpDetailsInternalPoolNameEnum = "SOCI_EGRESS"
)

var mappingCreateInternalFloatingIpDetailsInternalPoolNameEnum = map[string]CreateInternalFloatingIpDetailsInternalPoolNameEnum{
	"EXTERNAL":    CreateInternalFloatingIpDetailsInternalPoolNameExternal,
	"SOCI_EGRESS": CreateInternalFloatingIpDetailsInternalPoolNameSociEgress,
}

var mappingCreateInternalFloatingIpDetailsInternalPoolNameEnumLowerCase = map[string]CreateInternalFloatingIpDetailsInternalPoolNameEnum{
	"external":    CreateInternalFloatingIpDetailsInternalPoolNameExternal,
	"soci_egress": CreateInternalFloatingIpDetailsInternalPoolNameSociEgress,
}

// GetCreateInternalFloatingIpDetailsInternalPoolNameEnumValues Enumerates the set of values for CreateInternalFloatingIpDetailsInternalPoolNameEnum
func GetCreateInternalFloatingIpDetailsInternalPoolNameEnumValues() []CreateInternalFloatingIpDetailsInternalPoolNameEnum {
	values := make([]CreateInternalFloatingIpDetailsInternalPoolNameEnum, 0)
	for _, v := range mappingCreateInternalFloatingIpDetailsInternalPoolNameEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateInternalFloatingIpDetailsInternalPoolNameEnumStringValues Enumerates the set of values in String for CreateInternalFloatingIpDetailsInternalPoolNameEnum
func GetCreateInternalFloatingIpDetailsInternalPoolNameEnumStringValues() []string {
	return []string{
		"EXTERNAL",
		"SOCI_EGRESS",
	}
}

// GetMappingCreateInternalFloatingIpDetailsInternalPoolNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateInternalFloatingIpDetailsInternalPoolNameEnum(val string) (CreateInternalFloatingIpDetailsInternalPoolNameEnum, bool) {
	enum, ok := mappingCreateInternalFloatingIpDetailsInternalPoolNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
