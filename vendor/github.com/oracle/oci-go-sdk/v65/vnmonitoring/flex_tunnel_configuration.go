// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FlexTunnelConfiguration The tunnel configuration of a flex tunnel.
type FlexTunnelConfiguration interface {

	// IP address of your end of the tunnel.
	GetCustomerTunnelIp() *string

	// IP address of the oracle end of the tunnel.
	GetOracleTunnelIp() *string

	GetBgpSession() *FlexTunnelBgpSession
}

type flextunnelconfiguration struct {
	JsonData         []byte
	CustomerTunnelIp *string               `mandatory:"true" json:"customerTunnelIp"`
	OracleTunnelIp   *string               `mandatory:"true" json:"oracleTunnelIp"`
	BgpSession       *FlexTunnelBgpSession `mandatory:"true" json:"bgpSession"`
	TunnelType       string                `json:"tunnelType"`
}

// UnmarshalJSON unmarshals json
func (m *flextunnelconfiguration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerflextunnelconfiguration flextunnelconfiguration
	s := struct {
		Model Unmarshalerflextunnelconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CustomerTunnelIp = s.Model.CustomerTunnelIp
	m.OracleTunnelIp = s.Model.OracleTunnelIp
	m.BgpSession = s.Model.BgpSession
	m.TunnelType = s.Model.TunnelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *flextunnelconfiguration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TunnelType {
	case "GRE":
		mm := GreFlexTunnelConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for FlexTunnelConfiguration: %s.", m.TunnelType)
		return *m, nil
	}
}

// GetCustomerTunnelIp returns CustomerTunnelIp
func (m flextunnelconfiguration) GetCustomerTunnelIp() *string {
	return m.CustomerTunnelIp
}

// GetOracleTunnelIp returns OracleTunnelIp
func (m flextunnelconfiguration) GetOracleTunnelIp() *string {
	return m.OracleTunnelIp
}

// GetBgpSession returns BgpSession
func (m flextunnelconfiguration) GetBgpSession() *FlexTunnelBgpSession {
	return m.BgpSession
}

func (m flextunnelconfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m flextunnelconfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FlexTunnelConfigurationTunnelTypeEnum Enum with underlying type: string
type FlexTunnelConfigurationTunnelTypeEnum string

// Set of constants representing the allowable values for FlexTunnelConfigurationTunnelTypeEnum
const (
	FlexTunnelConfigurationTunnelTypeGre FlexTunnelConfigurationTunnelTypeEnum = "GRE"
)

var mappingFlexTunnelConfigurationTunnelTypeEnum = map[string]FlexTunnelConfigurationTunnelTypeEnum{
	"GRE": FlexTunnelConfigurationTunnelTypeGre,
}

var mappingFlexTunnelConfigurationTunnelTypeEnumLowerCase = map[string]FlexTunnelConfigurationTunnelTypeEnum{
	"gre": FlexTunnelConfigurationTunnelTypeGre,
}

// GetFlexTunnelConfigurationTunnelTypeEnumValues Enumerates the set of values for FlexTunnelConfigurationTunnelTypeEnum
func GetFlexTunnelConfigurationTunnelTypeEnumValues() []FlexTunnelConfigurationTunnelTypeEnum {
	values := make([]FlexTunnelConfigurationTunnelTypeEnum, 0)
	for _, v := range mappingFlexTunnelConfigurationTunnelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFlexTunnelConfigurationTunnelTypeEnumStringValues Enumerates the set of values in String for FlexTunnelConfigurationTunnelTypeEnum
func GetFlexTunnelConfigurationTunnelTypeEnumStringValues() []string {
	return []string{
		"GRE",
	}
}

// GetMappingFlexTunnelConfigurationTunnelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFlexTunnelConfigurationTunnelTypeEnum(val string) (FlexTunnelConfigurationTunnelTypeEnum, bool) {
	enum, ok := mappingFlexTunnelConfigurationTunnelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
