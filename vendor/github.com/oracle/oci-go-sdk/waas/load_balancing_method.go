// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
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
		return *m, nil
	}
}

func (m loadbalancingmethod) String() string {
	return common.PointerString(m)
}

// LoadBalancingMethodMethodEnum Enum with underlying type: string
type LoadBalancingMethodMethodEnum string

// Set of constants representing the allowable values for LoadBalancingMethodMethodEnum
const (
	LoadBalancingMethodMethodIpHash       LoadBalancingMethodMethodEnum = "IP_HASH"
	LoadBalancingMethodMethodRoundRobin   LoadBalancingMethodMethodEnum = "ROUND_ROBIN"
	LoadBalancingMethodMethodStickyCookie LoadBalancingMethodMethodEnum = "STICKY_COOKIE"
)

var mappingLoadBalancingMethodMethod = map[string]LoadBalancingMethodMethodEnum{
	"IP_HASH":       LoadBalancingMethodMethodIpHash,
	"ROUND_ROBIN":   LoadBalancingMethodMethodRoundRobin,
	"STICKY_COOKIE": LoadBalancingMethodMethodStickyCookie,
}

// GetLoadBalancingMethodMethodEnumValues Enumerates the set of values for LoadBalancingMethodMethodEnum
func GetLoadBalancingMethodMethodEnumValues() []LoadBalancingMethodMethodEnum {
	values := make([]LoadBalancingMethodMethodEnum, 0)
	for _, v := range mappingLoadBalancingMethodMethod {
		values = append(values, v)
	}
	return values
}
