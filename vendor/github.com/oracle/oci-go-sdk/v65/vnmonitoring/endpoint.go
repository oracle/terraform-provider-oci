// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Endpoint Information describing a source or destination in a `PathAnalyzerTest` resource.
type Endpoint interface {
}

type endpoint struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *endpoint) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerendpoint endpoint
	s := struct {
		Model Unmarshalerendpoint
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *endpoint) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "NETWORK_LOAD_BALANCER_LISTENER":
		mm := NetworkLoadBalancerListenerEndpoint{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE":
		mm := ComputeInstanceEndpoint{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NETWORK_LOAD_BALANCER":
		mm := NetworkLoadBalancerEndpoint{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOAD_BALANCER":
		mm := LoadBalancerEndpoint{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VNIC":
		mm := VnicEndpoint{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IP_ADDRESS":
		mm := IpAddressEndpoint{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VLAN":
		mm := VlanEndpoint{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SUBNET":
		mm := SubnetEndpoint{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOAD_BALANCER_LISTENER":
		mm := LoadBalancerListenerEndpoint{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for Endpoint: %s.", m.Type)
		return *m, nil
	}
}

func (m endpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m endpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EndpointTypeEnum Enum with underlying type: string
type EndpointTypeEnum string

// Set of constants representing the allowable values for EndpointTypeEnum
const (
	EndpointTypeIpAddress                   EndpointTypeEnum = "IP_ADDRESS"
	EndpointTypeSubnet                      EndpointTypeEnum = "SUBNET"
	EndpointTypeComputeInstance             EndpointTypeEnum = "COMPUTE_INSTANCE"
	EndpointTypeVnic                        EndpointTypeEnum = "VNIC"
	EndpointTypeLoadBalancer                EndpointTypeEnum = "LOAD_BALANCER"
	EndpointTypeLoadBalancerListener        EndpointTypeEnum = "LOAD_BALANCER_LISTENER"
	EndpointTypeNetworkLoadBalancer         EndpointTypeEnum = "NETWORK_LOAD_BALANCER"
	EndpointTypeNetworkLoadBalancerListener EndpointTypeEnum = "NETWORK_LOAD_BALANCER_LISTENER"
	EndpointTypeVlan                        EndpointTypeEnum = "VLAN"
)

var mappingEndpointTypeEnum = map[string]EndpointTypeEnum{
	"IP_ADDRESS":                     EndpointTypeIpAddress,
	"SUBNET":                         EndpointTypeSubnet,
	"COMPUTE_INSTANCE":               EndpointTypeComputeInstance,
	"VNIC":                           EndpointTypeVnic,
	"LOAD_BALANCER":                  EndpointTypeLoadBalancer,
	"LOAD_BALANCER_LISTENER":         EndpointTypeLoadBalancerListener,
	"NETWORK_LOAD_BALANCER":          EndpointTypeNetworkLoadBalancer,
	"NETWORK_LOAD_BALANCER_LISTENER": EndpointTypeNetworkLoadBalancerListener,
	"VLAN":                           EndpointTypeVlan,
}

var mappingEndpointTypeEnumLowerCase = map[string]EndpointTypeEnum{
	"ip_address":                     EndpointTypeIpAddress,
	"subnet":                         EndpointTypeSubnet,
	"compute_instance":               EndpointTypeComputeInstance,
	"vnic":                           EndpointTypeVnic,
	"load_balancer":                  EndpointTypeLoadBalancer,
	"load_balancer_listener":         EndpointTypeLoadBalancerListener,
	"network_load_balancer":          EndpointTypeNetworkLoadBalancer,
	"network_load_balancer_listener": EndpointTypeNetworkLoadBalancerListener,
	"vlan":                           EndpointTypeVlan,
}

// GetEndpointTypeEnumValues Enumerates the set of values for EndpointTypeEnum
func GetEndpointTypeEnumValues() []EndpointTypeEnum {
	values := make([]EndpointTypeEnum, 0)
	for _, v := range mappingEndpointTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEndpointTypeEnumStringValues Enumerates the set of values in String for EndpointTypeEnum
func GetEndpointTypeEnumStringValues() []string {
	return []string{
		"IP_ADDRESS",
		"SUBNET",
		"COMPUTE_INSTANCE",
		"VNIC",
		"LOAD_BALANCER",
		"LOAD_BALANCER_LISTENER",
		"NETWORK_LOAD_BALANCER",
		"NETWORK_LOAD_BALANCER_LISTENER",
		"VLAN",
	}
}

// GetMappingEndpointTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEndpointTypeEnum(val string) (EndpointTypeEnum, bool) {
	enum, ok := mappingEndpointTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
