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

// CreateFlexTunnelConfigurationDetails The create details for tunnel configuration.
type CreateFlexTunnelConfigurationDetails interface {

	// IP address of your end of the tunnel.
	GetCustomerTunnelIp() *string

	// IP address of the oracle end of the tunnel.
	GetOracleTunnelIp() *string

	GetBgpSession() *CreateFlexTunnelBgpSessionDetails
}

type createflextunnelconfigurationdetails struct {
	JsonData         []byte
	CustomerTunnelIp *string                            `mandatory:"true" json:"customerTunnelIp"`
	OracleTunnelIp   *string                            `mandatory:"true" json:"oracleTunnelIp"`
	BgpSession       *CreateFlexTunnelBgpSessionDetails `mandatory:"true" json:"bgpSession"`
	TunnelType       string                             `json:"tunnelType"`
}

// UnmarshalJSON unmarshals json
func (m *createflextunnelconfigurationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateflextunnelconfigurationdetails createflextunnelconfigurationdetails
	s := struct {
		Model Unmarshalercreateflextunnelconfigurationdetails
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
func (m *createflextunnelconfigurationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TunnelType {
	case "GRE":
		mm := CreateGreFlexTunnelConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateFlexTunnelConfigurationDetails: %s.", m.TunnelType)
		return *m, nil
	}
}

// GetCustomerTunnelIp returns CustomerTunnelIp
func (m createflextunnelconfigurationdetails) GetCustomerTunnelIp() *string {
	return m.CustomerTunnelIp
}

// GetOracleTunnelIp returns OracleTunnelIp
func (m createflextunnelconfigurationdetails) GetOracleTunnelIp() *string {
	return m.OracleTunnelIp
}

// GetBgpSession returns BgpSession
func (m createflextunnelconfigurationdetails) GetBgpSession() *CreateFlexTunnelBgpSessionDetails {
	return m.BgpSession
}

func (m createflextunnelconfigurationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createflextunnelconfigurationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateFlexTunnelConfigurationDetailsTunnelTypeEnum Enum with underlying type: string
type CreateFlexTunnelConfigurationDetailsTunnelTypeEnum string

// Set of constants representing the allowable values for CreateFlexTunnelConfigurationDetailsTunnelTypeEnum
const (
	CreateFlexTunnelConfigurationDetailsTunnelTypeGre CreateFlexTunnelConfigurationDetailsTunnelTypeEnum = "GRE"
)

var mappingCreateFlexTunnelConfigurationDetailsTunnelTypeEnum = map[string]CreateFlexTunnelConfigurationDetailsTunnelTypeEnum{
	"GRE": CreateFlexTunnelConfigurationDetailsTunnelTypeGre,
}

var mappingCreateFlexTunnelConfigurationDetailsTunnelTypeEnumLowerCase = map[string]CreateFlexTunnelConfigurationDetailsTunnelTypeEnum{
	"gre": CreateFlexTunnelConfigurationDetailsTunnelTypeGre,
}

// GetCreateFlexTunnelConfigurationDetailsTunnelTypeEnumValues Enumerates the set of values for CreateFlexTunnelConfigurationDetailsTunnelTypeEnum
func GetCreateFlexTunnelConfigurationDetailsTunnelTypeEnumValues() []CreateFlexTunnelConfigurationDetailsTunnelTypeEnum {
	values := make([]CreateFlexTunnelConfigurationDetailsTunnelTypeEnum, 0)
	for _, v := range mappingCreateFlexTunnelConfigurationDetailsTunnelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateFlexTunnelConfigurationDetailsTunnelTypeEnumStringValues Enumerates the set of values in String for CreateFlexTunnelConfigurationDetailsTunnelTypeEnum
func GetCreateFlexTunnelConfigurationDetailsTunnelTypeEnumStringValues() []string {
	return []string{
		"GRE",
	}
}

// GetMappingCreateFlexTunnelConfigurationDetailsTunnelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateFlexTunnelConfigurationDetailsTunnelTypeEnum(val string) (CreateFlexTunnelConfigurationDetailsTunnelTypeEnum, bool) {
	enum, ok := mappingCreateFlexTunnelConfigurationDetailsTunnelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
