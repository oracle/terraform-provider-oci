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

// AllowedSecurityConfiguration Defines the allowed security configuration for the traffic.
type AllowedSecurityConfiguration interface {
}

type allowedsecurityconfiguration struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *allowedsecurityconfiguration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerallowedsecurityconfiguration allowedsecurityconfiguration
	s := struct {
		Model Unmarshalerallowedsecurityconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *allowedsecurityconfiguration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "STATEFUL_EGRESS_SECURITY_LIST":
		mm := StatefulEgressSecurityListConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NSG":
		mm := NsgConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INGRESS_SECURITY_LIST":
		mm := IngressSecurityListConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STATEFUL_INGRESS_SECURITY_LIST":
		mm := StatefulIngressSecurityListConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EGRESS_SECURITY_LIST":
		mm := EgressSecurityListConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STATEFUL_NSG":
		mm := StatefulNsgConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for AllowedSecurityConfiguration: %s.", m.Type)
		return *m, nil
	}
}

func (m allowedsecurityconfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m allowedsecurityconfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AllowedSecurityConfigurationTypeEnum Enum with underlying type: string
type AllowedSecurityConfigurationTypeEnum string

// Set of constants representing the allowable values for AllowedSecurityConfigurationTypeEnum
const (
	AllowedSecurityConfigurationTypeNsg                         AllowedSecurityConfigurationTypeEnum = "NSG"
	AllowedSecurityConfigurationTypeStatefulNsg                 AllowedSecurityConfigurationTypeEnum = "STATEFUL_NSG"
	AllowedSecurityConfigurationTypeIngressSecurityList         AllowedSecurityConfigurationTypeEnum = "INGRESS_SECURITY_LIST"
	AllowedSecurityConfigurationTypeStatefulIngressSecurityList AllowedSecurityConfigurationTypeEnum = "STATEFUL_INGRESS_SECURITY_LIST"
	AllowedSecurityConfigurationTypeEgressSecurityList          AllowedSecurityConfigurationTypeEnum = "EGRESS_SECURITY_LIST"
	AllowedSecurityConfigurationTypeStatefulEgressSecurityList  AllowedSecurityConfigurationTypeEnum = "STATEFUL_EGRESS_SECURITY_LIST"
)

var mappingAllowedSecurityConfigurationTypeEnum = map[string]AllowedSecurityConfigurationTypeEnum{
	"NSG":                            AllowedSecurityConfigurationTypeNsg,
	"STATEFUL_NSG":                   AllowedSecurityConfigurationTypeStatefulNsg,
	"INGRESS_SECURITY_LIST":          AllowedSecurityConfigurationTypeIngressSecurityList,
	"STATEFUL_INGRESS_SECURITY_LIST": AllowedSecurityConfigurationTypeStatefulIngressSecurityList,
	"EGRESS_SECURITY_LIST":           AllowedSecurityConfigurationTypeEgressSecurityList,
	"STATEFUL_EGRESS_SECURITY_LIST":  AllowedSecurityConfigurationTypeStatefulEgressSecurityList,
}

var mappingAllowedSecurityConfigurationTypeEnumLowerCase = map[string]AllowedSecurityConfigurationTypeEnum{
	"nsg":                            AllowedSecurityConfigurationTypeNsg,
	"stateful_nsg":                   AllowedSecurityConfigurationTypeStatefulNsg,
	"ingress_security_list":          AllowedSecurityConfigurationTypeIngressSecurityList,
	"stateful_ingress_security_list": AllowedSecurityConfigurationTypeStatefulIngressSecurityList,
	"egress_security_list":           AllowedSecurityConfigurationTypeEgressSecurityList,
	"stateful_egress_security_list":  AllowedSecurityConfigurationTypeStatefulEgressSecurityList,
}

// GetAllowedSecurityConfigurationTypeEnumValues Enumerates the set of values for AllowedSecurityConfigurationTypeEnum
func GetAllowedSecurityConfigurationTypeEnumValues() []AllowedSecurityConfigurationTypeEnum {
	values := make([]AllowedSecurityConfigurationTypeEnum, 0)
	for _, v := range mappingAllowedSecurityConfigurationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAllowedSecurityConfigurationTypeEnumStringValues Enumerates the set of values in String for AllowedSecurityConfigurationTypeEnum
func GetAllowedSecurityConfigurationTypeEnumStringValues() []string {
	return []string{
		"NSG",
		"STATEFUL_NSG",
		"INGRESS_SECURITY_LIST",
		"STATEFUL_INGRESS_SECURITY_LIST",
		"EGRESS_SECURITY_LIST",
		"STATEFUL_EGRESS_SECURITY_LIST",
	}
}

// GetMappingAllowedSecurityConfigurationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAllowedSecurityConfigurationTypeEnum(val string) (AllowedSecurityConfigurationTypeEnum, bool) {
	enum, ok := mappingAllowedSecurityConfigurationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
