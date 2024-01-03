// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NetworkChannel Specifies the configuration needed when the target OCI resource, i.e., OKE cluster, resides
//
//	in customer's private network.
type NetworkChannel interface {
}

type networkchannel struct {
	JsonData           []byte
	NetworkChannelType string `json:"networkChannelType"`
}

// UnmarshalJSON unmarshals json
func (m *networkchannel) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalernetworkchannel networkchannel
	s := struct {
		Model Unmarshalernetworkchannel
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.NetworkChannelType = s.Model.NetworkChannelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *networkchannel) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.NetworkChannelType {
	case "SERVICE_VNIC_CHANNEL":
		mm := ServiceVnicChannel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRIVATE_ENDPOINT_CHANNEL":
		mm := PrivateEndpointChannel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for NetworkChannel: %s.", m.NetworkChannelType)
		return *m, nil
	}
}

func (m networkchannel) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m networkchannel) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NetworkChannelNetworkChannelTypeEnum Enum with underlying type: string
type NetworkChannelNetworkChannelTypeEnum string

// Set of constants representing the allowable values for NetworkChannelNetworkChannelTypeEnum
const (
	NetworkChannelNetworkChannelTypePrivateEndpointChannel NetworkChannelNetworkChannelTypeEnum = "PRIVATE_ENDPOINT_CHANNEL"
	NetworkChannelNetworkChannelTypeServiceVnicChannel     NetworkChannelNetworkChannelTypeEnum = "SERVICE_VNIC_CHANNEL"
)

var mappingNetworkChannelNetworkChannelTypeEnum = map[string]NetworkChannelNetworkChannelTypeEnum{
	"PRIVATE_ENDPOINT_CHANNEL": NetworkChannelNetworkChannelTypePrivateEndpointChannel,
	"SERVICE_VNIC_CHANNEL":     NetworkChannelNetworkChannelTypeServiceVnicChannel,
}

var mappingNetworkChannelNetworkChannelTypeEnumLowerCase = map[string]NetworkChannelNetworkChannelTypeEnum{
	"private_endpoint_channel": NetworkChannelNetworkChannelTypePrivateEndpointChannel,
	"service_vnic_channel":     NetworkChannelNetworkChannelTypeServiceVnicChannel,
}

// GetNetworkChannelNetworkChannelTypeEnumValues Enumerates the set of values for NetworkChannelNetworkChannelTypeEnum
func GetNetworkChannelNetworkChannelTypeEnumValues() []NetworkChannelNetworkChannelTypeEnum {
	values := make([]NetworkChannelNetworkChannelTypeEnum, 0)
	for _, v := range mappingNetworkChannelNetworkChannelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkChannelNetworkChannelTypeEnumStringValues Enumerates the set of values in String for NetworkChannelNetworkChannelTypeEnum
func GetNetworkChannelNetworkChannelTypeEnumStringValues() []string {
	return []string{
		"PRIVATE_ENDPOINT_CHANNEL",
		"SERVICE_VNIC_CHANNEL",
	}
}

// GetMappingNetworkChannelNetworkChannelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkChannelNetworkChannelTypeEnum(val string) (NetworkChannelNetworkChannelTypeEnum, bool) {
	enum, ok := mappingNetworkChannelNetworkChannelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
