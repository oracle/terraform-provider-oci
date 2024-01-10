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

// ForwardedRoutingConfiguration Defines the type of the resource that forwarded traffic.
type ForwardedRoutingConfiguration interface {
}

type forwardedroutingconfiguration struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *forwardedroutingconfiguration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerforwardedroutingconfiguration forwardedroutingconfiguration
	s := struct {
		Model Unmarshalerforwardedroutingconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *forwardedroutingconfiguration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "VCN":
		mm := VcnRoutingConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DRG":
		mm := DrgRoutingConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ForwardedRoutingConfiguration: %s.", m.Type)
		return *m, nil
	}
}

func (m forwardedroutingconfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m forwardedroutingconfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ForwardedRoutingConfigurationTypeEnum Enum with underlying type: string
type ForwardedRoutingConfigurationTypeEnum string

// Set of constants representing the allowable values for ForwardedRoutingConfigurationTypeEnum
const (
	ForwardedRoutingConfigurationTypeVcn ForwardedRoutingConfigurationTypeEnum = "VCN"
	ForwardedRoutingConfigurationTypeDrg ForwardedRoutingConfigurationTypeEnum = "DRG"
)

var mappingForwardedRoutingConfigurationTypeEnum = map[string]ForwardedRoutingConfigurationTypeEnum{
	"VCN": ForwardedRoutingConfigurationTypeVcn,
	"DRG": ForwardedRoutingConfigurationTypeDrg,
}

var mappingForwardedRoutingConfigurationTypeEnumLowerCase = map[string]ForwardedRoutingConfigurationTypeEnum{
	"vcn": ForwardedRoutingConfigurationTypeVcn,
	"drg": ForwardedRoutingConfigurationTypeDrg,
}

// GetForwardedRoutingConfigurationTypeEnumValues Enumerates the set of values for ForwardedRoutingConfigurationTypeEnum
func GetForwardedRoutingConfigurationTypeEnumValues() []ForwardedRoutingConfigurationTypeEnum {
	values := make([]ForwardedRoutingConfigurationTypeEnum, 0)
	for _, v := range mappingForwardedRoutingConfigurationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetForwardedRoutingConfigurationTypeEnumStringValues Enumerates the set of values in String for ForwardedRoutingConfigurationTypeEnum
func GetForwardedRoutingConfigurationTypeEnumStringValues() []string {
	return []string{
		"VCN",
		"DRG",
	}
}

// GetMappingForwardedRoutingConfigurationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingForwardedRoutingConfigurationTypeEnum(val string) (ForwardedRoutingConfigurationTypeEnum, bool) {
	enum, ok := mappingForwardedRoutingConfigurationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
