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

// UpdateFlexTunnelConfigurationDetails The update details for tunnel configuration.
type UpdateFlexTunnelConfigurationDetails interface {

	// IP address of your end of the tunnel.
	GetCustomerTunnelIp() *string

	// IP address of the oracle end of the tunnel.
	GetOracleTunnelIp() *string

	GetBgpSession() *UpdateFlexTunnelBgpSessionDetails
}

type updateflextunnelconfigurationdetails struct {
	JsonData         []byte
	CustomerTunnelIp *string                            `mandatory:"false" json:"customerTunnelIp"`
	OracleTunnelIp   *string                            `mandatory:"false" json:"oracleTunnelIp"`
	BgpSession       *UpdateFlexTunnelBgpSessionDetails `mandatory:"false" json:"bgpSession"`
	TunnelType       string                             `json:"tunnelType"`
}

// UnmarshalJSON unmarshals json
func (m *updateflextunnelconfigurationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateflextunnelconfigurationdetails updateflextunnelconfigurationdetails
	s := struct {
		Model Unmarshalerupdateflextunnelconfigurationdetails
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
func (m *updateflextunnelconfigurationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TunnelType {
	case "GRE":
		mm := UpdateGreFlexTunnelConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateFlexTunnelConfigurationDetails: %s.", m.TunnelType)
		return *m, nil
	}
}

// GetCustomerTunnelIp returns CustomerTunnelIp
func (m updateflextunnelconfigurationdetails) GetCustomerTunnelIp() *string {
	return m.CustomerTunnelIp
}

// GetOracleTunnelIp returns OracleTunnelIp
func (m updateflextunnelconfigurationdetails) GetOracleTunnelIp() *string {
	return m.OracleTunnelIp
}

// GetBgpSession returns BgpSession
func (m updateflextunnelconfigurationdetails) GetBgpSession() *UpdateFlexTunnelBgpSessionDetails {
	return m.BgpSession
}

func (m updateflextunnelconfigurationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateflextunnelconfigurationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateFlexTunnelConfigurationDetailsTunnelTypeEnum Enum with underlying type: string
type UpdateFlexTunnelConfigurationDetailsTunnelTypeEnum string

// Set of constants representing the allowable values for UpdateFlexTunnelConfigurationDetailsTunnelTypeEnum
const (
	UpdateFlexTunnelConfigurationDetailsTunnelTypeGre UpdateFlexTunnelConfigurationDetailsTunnelTypeEnum = "GRE"
)

var mappingUpdateFlexTunnelConfigurationDetailsTunnelTypeEnum = map[string]UpdateFlexTunnelConfigurationDetailsTunnelTypeEnum{
	"GRE": UpdateFlexTunnelConfigurationDetailsTunnelTypeGre,
}

var mappingUpdateFlexTunnelConfigurationDetailsTunnelTypeEnumLowerCase = map[string]UpdateFlexTunnelConfigurationDetailsTunnelTypeEnum{
	"gre": UpdateFlexTunnelConfigurationDetailsTunnelTypeGre,
}

// GetUpdateFlexTunnelConfigurationDetailsTunnelTypeEnumValues Enumerates the set of values for UpdateFlexTunnelConfigurationDetailsTunnelTypeEnum
func GetUpdateFlexTunnelConfigurationDetailsTunnelTypeEnumValues() []UpdateFlexTunnelConfigurationDetailsTunnelTypeEnum {
	values := make([]UpdateFlexTunnelConfigurationDetailsTunnelTypeEnum, 0)
	for _, v := range mappingUpdateFlexTunnelConfigurationDetailsTunnelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateFlexTunnelConfigurationDetailsTunnelTypeEnumStringValues Enumerates the set of values in String for UpdateFlexTunnelConfigurationDetailsTunnelTypeEnum
func GetUpdateFlexTunnelConfigurationDetailsTunnelTypeEnumStringValues() []string {
	return []string{
		"GRE",
	}
}

// GetMappingUpdateFlexTunnelConfigurationDetailsTunnelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateFlexTunnelConfigurationDetailsTunnelTypeEnum(val string) (UpdateFlexTunnelConfigurationDetailsTunnelTypeEnum, bool) {
	enum, ok := mappingUpdateFlexTunnelConfigurationDetailsTunnelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
