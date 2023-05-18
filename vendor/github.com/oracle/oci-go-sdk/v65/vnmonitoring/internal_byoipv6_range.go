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

// InternalByoipv6Range This is a ByoipRange type that only used by VCNIP to propagate a ByoipRange for IPv6 to VCNCP after validation is done.
type InternalByoipv6Range struct {

	// A user-friendly name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the ByoipRange.
	Id *string `mandatory:"false" json:"id"`

	// The OCID of the compartment that contains the ByoipRange.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The InternalByoipv6Range's current state.
	LifecycleState InternalByoipv6RangeLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The CIDR IPv6 address block of the Imported ByoipRange. The CIDR length is always /64.
	Byoipv6RangeCidrBlock *string `mandatory:"false" json:"byoipv6RangeCidrBlock"`
}

func (m InternalByoipv6Range) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternalByoipv6Range) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInternalByoipv6RangeLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetInternalByoipv6RangeLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InternalByoipv6RangeLifecycleStateEnum Enum with underlying type: string
type InternalByoipv6RangeLifecycleStateEnum string

// Set of constants representing the allowable values for InternalByoipv6RangeLifecycleStateEnum
const (
	InternalByoipv6RangeLifecycleStateProvisioning InternalByoipv6RangeLifecycleStateEnum = "PROVISIONING"
	InternalByoipv6RangeLifecycleStateAvailable    InternalByoipv6RangeLifecycleStateEnum = "AVAILABLE"
	InternalByoipv6RangeLifecycleStateTerminating  InternalByoipv6RangeLifecycleStateEnum = "TERMINATING"
	InternalByoipv6RangeLifecycleStateTerminated   InternalByoipv6RangeLifecycleStateEnum = "TERMINATED"
)

var mappingInternalByoipv6RangeLifecycleStateEnum = map[string]InternalByoipv6RangeLifecycleStateEnum{
	"PROVISIONING": InternalByoipv6RangeLifecycleStateProvisioning,
	"AVAILABLE":    InternalByoipv6RangeLifecycleStateAvailable,
	"TERMINATING":  InternalByoipv6RangeLifecycleStateTerminating,
	"TERMINATED":   InternalByoipv6RangeLifecycleStateTerminated,
}

var mappingInternalByoipv6RangeLifecycleStateEnumLowerCase = map[string]InternalByoipv6RangeLifecycleStateEnum{
	"provisioning": InternalByoipv6RangeLifecycleStateProvisioning,
	"available":    InternalByoipv6RangeLifecycleStateAvailable,
	"terminating":  InternalByoipv6RangeLifecycleStateTerminating,
	"terminated":   InternalByoipv6RangeLifecycleStateTerminated,
}

// GetInternalByoipv6RangeLifecycleStateEnumValues Enumerates the set of values for InternalByoipv6RangeLifecycleStateEnum
func GetInternalByoipv6RangeLifecycleStateEnumValues() []InternalByoipv6RangeLifecycleStateEnum {
	values := make([]InternalByoipv6RangeLifecycleStateEnum, 0)
	for _, v := range mappingInternalByoipv6RangeLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetInternalByoipv6RangeLifecycleStateEnumStringValues Enumerates the set of values in String for InternalByoipv6RangeLifecycleStateEnum
func GetInternalByoipv6RangeLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
	}
}

// GetMappingInternalByoipv6RangeLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternalByoipv6RangeLifecycleStateEnum(val string) (InternalByoipv6RangeLifecycleStateEnum, bool) {
	enum, ok := mappingInternalByoipv6RangeLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
