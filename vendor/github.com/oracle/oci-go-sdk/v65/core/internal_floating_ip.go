// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InternalFloatingIp An internal floating IP
type InternalFloatingIp struct {

	// ID of the compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Unique identifier of a floating IP
	Id *string `mandatory:"true" json:"id"`

	// The IP address of the floating IP
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// The current state of the floating IP
	State InternalFloatingIpStateEnum `mandatory:"true" json:"state"`

	// User friendly name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Unique identifier of a VNIC the floating IP is mapped to
	MappedVnicId *string `mandatory:"false" json:"mappedVnicId"`

	// The current state of the floating IP to VNIC mapping
	MappingStatus InternalFloatingIpMappingStatusEnum `mandatory:"false" json:"mappingStatus,omitempty"`

	// Creation time of the entity
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m InternalFloatingIp) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternalFloatingIp) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInternalFloatingIpStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetInternalFloatingIpStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingInternalFloatingIpMappingStatusEnum(string(m.MappingStatus)); !ok && m.MappingStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MappingStatus: %s. Supported values are: %s.", m.MappingStatus, strings.Join(GetInternalFloatingIpMappingStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InternalFloatingIpMappingStatusEnum Enum with underlying type: string
type InternalFloatingIpMappingStatusEnum string

// Set of constants representing the allowable values for InternalFloatingIpMappingStatusEnum
const (
	InternalFloatingIpMappingStatusAttaching InternalFloatingIpMappingStatusEnum = "ATTACHING"
	InternalFloatingIpMappingStatusAttached  InternalFloatingIpMappingStatusEnum = "ATTACHED"
	InternalFloatingIpMappingStatusDetaching InternalFloatingIpMappingStatusEnum = "DETACHING"
	InternalFloatingIpMappingStatusDetached  InternalFloatingIpMappingStatusEnum = "DETACHED"
)

var mappingInternalFloatingIpMappingStatusEnum = map[string]InternalFloatingIpMappingStatusEnum{
	"ATTACHING": InternalFloatingIpMappingStatusAttaching,
	"ATTACHED":  InternalFloatingIpMappingStatusAttached,
	"DETACHING": InternalFloatingIpMappingStatusDetaching,
	"DETACHED":  InternalFloatingIpMappingStatusDetached,
}

var mappingInternalFloatingIpMappingStatusEnumLowerCase = map[string]InternalFloatingIpMappingStatusEnum{
	"attaching": InternalFloatingIpMappingStatusAttaching,
	"attached":  InternalFloatingIpMappingStatusAttached,
	"detaching": InternalFloatingIpMappingStatusDetaching,
	"detached":  InternalFloatingIpMappingStatusDetached,
}

// GetInternalFloatingIpMappingStatusEnumValues Enumerates the set of values for InternalFloatingIpMappingStatusEnum
func GetInternalFloatingIpMappingStatusEnumValues() []InternalFloatingIpMappingStatusEnum {
	values := make([]InternalFloatingIpMappingStatusEnum, 0)
	for _, v := range mappingInternalFloatingIpMappingStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetInternalFloatingIpMappingStatusEnumStringValues Enumerates the set of values in String for InternalFloatingIpMappingStatusEnum
func GetInternalFloatingIpMappingStatusEnumStringValues() []string {
	return []string{
		"ATTACHING",
		"ATTACHED",
		"DETACHING",
		"DETACHED",
	}
}

// GetMappingInternalFloatingIpMappingStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternalFloatingIpMappingStatusEnum(val string) (InternalFloatingIpMappingStatusEnum, bool) {
	enum, ok := mappingInternalFloatingIpMappingStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// InternalFloatingIpStateEnum Enum with underlying type: string
type InternalFloatingIpStateEnum string

// Set of constants representing the allowable values for InternalFloatingIpStateEnum
const (
	InternalFloatingIpStateProvisioning InternalFloatingIpStateEnum = "PROVISIONING"
	InternalFloatingIpStateAvailable    InternalFloatingIpStateEnum = "AVAILABLE"
	InternalFloatingIpStateTerminating  InternalFloatingIpStateEnum = "TERMINATING"
	InternalFloatingIpStateTerminated   InternalFloatingIpStateEnum = "TERMINATED"
)

var mappingInternalFloatingIpStateEnum = map[string]InternalFloatingIpStateEnum{
	"PROVISIONING": InternalFloatingIpStateProvisioning,
	"AVAILABLE":    InternalFloatingIpStateAvailable,
	"TERMINATING":  InternalFloatingIpStateTerminating,
	"TERMINATED":   InternalFloatingIpStateTerminated,
}

var mappingInternalFloatingIpStateEnumLowerCase = map[string]InternalFloatingIpStateEnum{
	"provisioning": InternalFloatingIpStateProvisioning,
	"available":    InternalFloatingIpStateAvailable,
	"terminating":  InternalFloatingIpStateTerminating,
	"terminated":   InternalFloatingIpStateTerminated,
}

// GetInternalFloatingIpStateEnumValues Enumerates the set of values for InternalFloatingIpStateEnum
func GetInternalFloatingIpStateEnumValues() []InternalFloatingIpStateEnum {
	values := make([]InternalFloatingIpStateEnum, 0)
	for _, v := range mappingInternalFloatingIpStateEnum {
		values = append(values, v)
	}
	return values
}

// GetInternalFloatingIpStateEnumStringValues Enumerates the set of values in String for InternalFloatingIpStateEnum
func GetInternalFloatingIpStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
	}
}

// GetMappingInternalFloatingIpStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternalFloatingIpStateEnum(val string) (InternalFloatingIpStateEnum, bool) {
	enum, ok := mappingInternalFloatingIpStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
