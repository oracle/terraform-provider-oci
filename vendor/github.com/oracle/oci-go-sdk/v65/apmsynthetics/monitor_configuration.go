// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MonitorConfiguration Details of monitor configuration.
type MonitorConfiguration interface {

	// If isFailureRetried is enabled, then a failed call will be retried.
	GetIsFailureRetried() *bool

	GetDnsConfiguration() *DnsConfiguration
}

type monitorconfiguration struct {
	JsonData         []byte
	IsFailureRetried *bool             `mandatory:"false" json:"isFailureRetried"`
	DnsConfiguration *DnsConfiguration `mandatory:"false" json:"dnsConfiguration"`
	ConfigType       string            `json:"configType"`
}

// UnmarshalJSON unmarshals json
func (m *monitorconfiguration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermonitorconfiguration monitorconfiguration
	s := struct {
		Model Unmarshalermonitorconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.IsFailureRetried = s.Model.IsFailureRetried
	m.DnsConfiguration = s.Model.DnsConfiguration
	m.ConfigType = s.Model.ConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *monitorconfiguration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigType {
	case "DNSSEC_CONFIG":
		mm := DnsSecMonitorConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DNS_TRACE_CONFIG":
		mm := DnsTraceMonitorConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SCRIPTED_REST_CONFIG":
		mm := ScriptedRestMonitorConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DNS_SERVER_CONFIG":
		mm := DnsServerMonitorConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SCRIPTED_BROWSER_CONFIG":
		mm := ScriptedBrowserMonitorConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REST_CONFIG":
		mm := RestMonitorConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BROWSER_CONFIG":
		mm := BrowserMonitorConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NETWORK_CONFIG":
		mm := NetworkMonitorConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for MonitorConfiguration: %s.", m.ConfigType)
		return *m, nil
	}
}

// GetIsFailureRetried returns IsFailureRetried
func (m monitorconfiguration) GetIsFailureRetried() *bool {
	return m.IsFailureRetried
}

// GetDnsConfiguration returns DnsConfiguration
func (m monitorconfiguration) GetDnsConfiguration() *DnsConfiguration {
	return m.DnsConfiguration
}

func (m monitorconfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m monitorconfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MonitorConfigurationConfigTypeEnum Enum with underlying type: string
type MonitorConfigurationConfigTypeEnum string

// Set of constants representing the allowable values for MonitorConfigurationConfigTypeEnum
const (
	MonitorConfigurationConfigTypeBrowserConfig         MonitorConfigurationConfigTypeEnum = "BROWSER_CONFIG"
	MonitorConfigurationConfigTypeScriptedBrowserConfig MonitorConfigurationConfigTypeEnum = "SCRIPTED_BROWSER_CONFIG"
	MonitorConfigurationConfigTypeRestConfig            MonitorConfigurationConfigTypeEnum = "REST_CONFIG"
	MonitorConfigurationConfigTypeScriptedRestConfig    MonitorConfigurationConfigTypeEnum = "SCRIPTED_REST_CONFIG"
	MonitorConfigurationConfigTypeNetworkConfig         MonitorConfigurationConfigTypeEnum = "NETWORK_CONFIG"
	MonitorConfigurationConfigTypeDnsServerConfig       MonitorConfigurationConfigTypeEnum = "DNS_SERVER_CONFIG"
	MonitorConfigurationConfigTypeDnsTraceConfig        MonitorConfigurationConfigTypeEnum = "DNS_TRACE_CONFIG"
	MonitorConfigurationConfigTypeDnssecConfig          MonitorConfigurationConfigTypeEnum = "DNSSEC_CONFIG"
)

var mappingMonitorConfigurationConfigTypeEnum = map[string]MonitorConfigurationConfigTypeEnum{
	"BROWSER_CONFIG":          MonitorConfigurationConfigTypeBrowserConfig,
	"SCRIPTED_BROWSER_CONFIG": MonitorConfigurationConfigTypeScriptedBrowserConfig,
	"REST_CONFIG":             MonitorConfigurationConfigTypeRestConfig,
	"SCRIPTED_REST_CONFIG":    MonitorConfigurationConfigTypeScriptedRestConfig,
	"NETWORK_CONFIG":          MonitorConfigurationConfigTypeNetworkConfig,
	"DNS_SERVER_CONFIG":       MonitorConfigurationConfigTypeDnsServerConfig,
	"DNS_TRACE_CONFIG":        MonitorConfigurationConfigTypeDnsTraceConfig,
	"DNSSEC_CONFIG":           MonitorConfigurationConfigTypeDnssecConfig,
}

var mappingMonitorConfigurationConfigTypeEnumLowerCase = map[string]MonitorConfigurationConfigTypeEnum{
	"browser_config":          MonitorConfigurationConfigTypeBrowserConfig,
	"scripted_browser_config": MonitorConfigurationConfigTypeScriptedBrowserConfig,
	"rest_config":             MonitorConfigurationConfigTypeRestConfig,
	"scripted_rest_config":    MonitorConfigurationConfigTypeScriptedRestConfig,
	"network_config":          MonitorConfigurationConfigTypeNetworkConfig,
	"dns_server_config":       MonitorConfigurationConfigTypeDnsServerConfig,
	"dns_trace_config":        MonitorConfigurationConfigTypeDnsTraceConfig,
	"dnssec_config":           MonitorConfigurationConfigTypeDnssecConfig,
}

// GetMonitorConfigurationConfigTypeEnumValues Enumerates the set of values for MonitorConfigurationConfigTypeEnum
func GetMonitorConfigurationConfigTypeEnumValues() []MonitorConfigurationConfigTypeEnum {
	values := make([]MonitorConfigurationConfigTypeEnum, 0)
	for _, v := range mappingMonitorConfigurationConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMonitorConfigurationConfigTypeEnumStringValues Enumerates the set of values in String for MonitorConfigurationConfigTypeEnum
func GetMonitorConfigurationConfigTypeEnumStringValues() []string {
	return []string{
		"BROWSER_CONFIG",
		"SCRIPTED_BROWSER_CONFIG",
		"REST_CONFIG",
		"SCRIPTED_REST_CONFIG",
		"NETWORK_CONFIG",
		"DNS_SERVER_CONFIG",
		"DNS_TRACE_CONFIG",
		"DNSSEC_CONFIG",
	}
}

// GetMappingMonitorConfigurationConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMonitorConfigurationConfigTypeEnum(val string) (MonitorConfigurationConfigTypeEnum, bool) {
	enum, ok := mappingMonitorConfigurationConfigTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
