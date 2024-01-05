// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DefaultVirtualServiceRoutingPolicy Routing policy for the virtual service.
type DefaultVirtualServiceRoutingPolicy struct {

	// Type of the virtual service routing policy.
	Type DefaultVirtualServiceRoutingPolicyTypeEnum `mandatory:"true" json:"type"`
}

func (m DefaultVirtualServiceRoutingPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DefaultVirtualServiceRoutingPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDefaultVirtualServiceRoutingPolicyTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetDefaultVirtualServiceRoutingPolicyTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DefaultVirtualServiceRoutingPolicyTypeEnum Enum with underlying type: string
type DefaultVirtualServiceRoutingPolicyTypeEnum string

// Set of constants representing the allowable values for DefaultVirtualServiceRoutingPolicyTypeEnum
const (
	DefaultVirtualServiceRoutingPolicyTypeUniform DefaultVirtualServiceRoutingPolicyTypeEnum = "UNIFORM"
	DefaultVirtualServiceRoutingPolicyTypeDeny    DefaultVirtualServiceRoutingPolicyTypeEnum = "DENY"
)

var mappingDefaultVirtualServiceRoutingPolicyTypeEnum = map[string]DefaultVirtualServiceRoutingPolicyTypeEnum{
	"UNIFORM": DefaultVirtualServiceRoutingPolicyTypeUniform,
	"DENY":    DefaultVirtualServiceRoutingPolicyTypeDeny,
}

var mappingDefaultVirtualServiceRoutingPolicyTypeEnumLowerCase = map[string]DefaultVirtualServiceRoutingPolicyTypeEnum{
	"uniform": DefaultVirtualServiceRoutingPolicyTypeUniform,
	"deny":    DefaultVirtualServiceRoutingPolicyTypeDeny,
}

// GetDefaultVirtualServiceRoutingPolicyTypeEnumValues Enumerates the set of values for DefaultVirtualServiceRoutingPolicyTypeEnum
func GetDefaultVirtualServiceRoutingPolicyTypeEnumValues() []DefaultVirtualServiceRoutingPolicyTypeEnum {
	values := make([]DefaultVirtualServiceRoutingPolicyTypeEnum, 0)
	for _, v := range mappingDefaultVirtualServiceRoutingPolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDefaultVirtualServiceRoutingPolicyTypeEnumStringValues Enumerates the set of values in String for DefaultVirtualServiceRoutingPolicyTypeEnum
func GetDefaultVirtualServiceRoutingPolicyTypeEnumStringValues() []string {
	return []string{
		"UNIFORM",
		"DENY",
	}
}

// GetMappingDefaultVirtualServiceRoutingPolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDefaultVirtualServiceRoutingPolicyTypeEnum(val string) (DefaultVirtualServiceRoutingPolicyTypeEnum, bool) {
	enum, ok := mappingDefaultVirtualServiceRoutingPolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
