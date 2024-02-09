// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LoadBalancer load balancing algorithm to use when picking an upstream host in the cluster.
type LoadBalancer interface {
}

type loadbalancer struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *loadbalancer) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerloadbalancer loadbalancer
	s := struct {
		Model Unmarshalerloadbalancer
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *loadbalancer) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "RING_HASH":
		mm := RingHashLoadBalancer{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ROUND_ROBIN":
		mm := RoundRobinLoadBalancer{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RANDOM":
		mm := RandomLoadBalancer{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MAGLEV":
		mm := MaglevLoadBalancer{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LEAST_REQUEST":
		mm := LeastRequestLoadBalancer{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for LoadBalancer: %s.", m.Type)
		return *m, nil
	}
}

func (m loadbalancer) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m loadbalancer) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LoadBalancerTypeEnum Enum with underlying type: string
type LoadBalancerTypeEnum string

// Set of constants representing the allowable values for LoadBalancerTypeEnum
const (
	LoadBalancerTypeRoundRobin   LoadBalancerTypeEnum = "ROUND_ROBIN"
	LoadBalancerTypeLeastRequest LoadBalancerTypeEnum = "LEAST_REQUEST"
	LoadBalancerTypeRandom       LoadBalancerTypeEnum = "RANDOM"
	LoadBalancerTypeRingHash     LoadBalancerTypeEnum = "RING_HASH"
	LoadBalancerTypeMaglev       LoadBalancerTypeEnum = "MAGLEV"
)

var mappingLoadBalancerTypeEnum = map[string]LoadBalancerTypeEnum{
	"ROUND_ROBIN":   LoadBalancerTypeRoundRobin,
	"LEAST_REQUEST": LoadBalancerTypeLeastRequest,
	"RANDOM":        LoadBalancerTypeRandom,
	"RING_HASH":     LoadBalancerTypeRingHash,
	"MAGLEV":        LoadBalancerTypeMaglev,
}

var mappingLoadBalancerTypeEnumLowerCase = map[string]LoadBalancerTypeEnum{
	"round_robin":   LoadBalancerTypeRoundRobin,
	"least_request": LoadBalancerTypeLeastRequest,
	"random":        LoadBalancerTypeRandom,
	"ring_hash":     LoadBalancerTypeRingHash,
	"maglev":        LoadBalancerTypeMaglev,
}

// GetLoadBalancerTypeEnumValues Enumerates the set of values for LoadBalancerTypeEnum
func GetLoadBalancerTypeEnumValues() []LoadBalancerTypeEnum {
	values := make([]LoadBalancerTypeEnum, 0)
	for _, v := range mappingLoadBalancerTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLoadBalancerTypeEnumStringValues Enumerates the set of values in String for LoadBalancerTypeEnum
func GetLoadBalancerTypeEnumStringValues() []string {
	return []string{
		"ROUND_ROBIN",
		"LEAST_REQUEST",
		"RANDOM",
		"RING_HASH",
		"MAGLEV",
	}
}

// GetMappingLoadBalancerTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLoadBalancerTypeEnum(val string) (LoadBalancerTypeEnum, bool) {
	enum, ok := mappingLoadBalancerTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
