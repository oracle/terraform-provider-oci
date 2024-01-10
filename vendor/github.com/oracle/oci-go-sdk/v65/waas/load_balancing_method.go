// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LoadBalancingMethod The representation of LoadBalancingMethod
type LoadBalancingMethod interface {
}

type loadbalancingmethod struct {
	JsonData []byte
	Method   string `json:"method"`
}

// UnmarshalJSON unmarshals json
func (m *loadbalancingmethod) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerloadbalancingmethod loadbalancingmethod
	s := struct {
		Model Unmarshalerloadbalancingmethod
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Method = s.Model.Method

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *loadbalancingmethod) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Method {
	case "ROUND_ROBIN":
		mm := RoundRobinLoadBalancingMethod{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STICKY_COOKIE":
		mm := StickyCookieLoadBalancingMethod{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IP_HASH":
		mm := IpHashLoadBalancingMethod{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for LoadBalancingMethod: %s.", m.Method)
		return *m, nil
	}
}

func (m loadbalancingmethod) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m loadbalancingmethod) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LoadBalancingMethodMethodEnum Enum with underlying type: string
type LoadBalancingMethodMethodEnum string

// Set of constants representing the allowable values for LoadBalancingMethodMethodEnum
const (
	LoadBalancingMethodMethodIpHash       LoadBalancingMethodMethodEnum = "IP_HASH"
	LoadBalancingMethodMethodRoundRobin   LoadBalancingMethodMethodEnum = "ROUND_ROBIN"
	LoadBalancingMethodMethodStickyCookie LoadBalancingMethodMethodEnum = "STICKY_COOKIE"
)

var mappingLoadBalancingMethodMethodEnum = map[string]LoadBalancingMethodMethodEnum{
	"IP_HASH":       LoadBalancingMethodMethodIpHash,
	"ROUND_ROBIN":   LoadBalancingMethodMethodRoundRobin,
	"STICKY_COOKIE": LoadBalancingMethodMethodStickyCookie,
}

var mappingLoadBalancingMethodMethodEnumLowerCase = map[string]LoadBalancingMethodMethodEnum{
	"ip_hash":       LoadBalancingMethodMethodIpHash,
	"round_robin":   LoadBalancingMethodMethodRoundRobin,
	"sticky_cookie": LoadBalancingMethodMethodStickyCookie,
}

// GetLoadBalancingMethodMethodEnumValues Enumerates the set of values for LoadBalancingMethodMethodEnum
func GetLoadBalancingMethodMethodEnumValues() []LoadBalancingMethodMethodEnum {
	values := make([]LoadBalancingMethodMethodEnum, 0)
	for _, v := range mappingLoadBalancingMethodMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetLoadBalancingMethodMethodEnumStringValues Enumerates the set of values in String for LoadBalancingMethodMethodEnum
func GetLoadBalancingMethodMethodEnumStringValues() []string {
	return []string{
		"IP_HASH",
		"ROUND_ROBIN",
		"STICKY_COOKIE",
	}
}

// GetMappingLoadBalancingMethodMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLoadBalancingMethodMethodEnum(val string) (LoadBalancingMethodMethodEnum, bool) {
	enum, ok := mappingLoadBalancingMethodMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
