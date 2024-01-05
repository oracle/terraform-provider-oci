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

// TrafficNode Defines the configuration of the OCI entity that represents a traffic node in `PathAnalysisResult`.
type TrafficNode interface {
	GetEgressTraffic() *EgressTrafficSpec

	GetNextHopRoutingAction() RoutingAction

	GetEgressSecurityAction() SecurityAction

	GetIngressSecurityAction() SecurityAction
}

type trafficnode struct {
	JsonData              []byte
	EgressTraffic         *EgressTrafficSpec `mandatory:"false" json:"egressTraffic"`
	NextHopRoutingAction  routingaction      `mandatory:"false" json:"nextHopRoutingAction"`
	EgressSecurityAction  securityaction     `mandatory:"false" json:"egressSecurityAction"`
	IngressSecurityAction securityaction     `mandatory:"false" json:"ingressSecurityAction"`
	Type                  string             `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *trafficnode) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertrafficnode trafficnode
	s := struct {
		Model Unmarshalertrafficnode
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.EgressTraffic = s.Model.EgressTraffic
	m.NextHopRoutingAction = s.Model.NextHopRoutingAction
	m.EgressSecurityAction = s.Model.EgressSecurityAction
	m.IngressSecurityAction = s.Model.IngressSecurityAction
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *trafficnode) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "VISIBLE":
		mm := VisibleTrafficNode{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ACCESS_DENIED":
		mm := AccessDeniedTrafficNode{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for TrafficNode: %s.", m.Type)
		return *m, nil
	}
}

// GetEgressTraffic returns EgressTraffic
func (m trafficnode) GetEgressTraffic() *EgressTrafficSpec {
	return m.EgressTraffic
}

// GetNextHopRoutingAction returns NextHopRoutingAction
func (m trafficnode) GetNextHopRoutingAction() routingaction {
	return m.NextHopRoutingAction
}

// GetEgressSecurityAction returns EgressSecurityAction
func (m trafficnode) GetEgressSecurityAction() securityaction {
	return m.EgressSecurityAction
}

// GetIngressSecurityAction returns IngressSecurityAction
func (m trafficnode) GetIngressSecurityAction() securityaction {
	return m.IngressSecurityAction
}

func (m trafficnode) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m trafficnode) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TrafficNodeTypeEnum Enum with underlying type: string
type TrafficNodeTypeEnum string

// Set of constants representing the allowable values for TrafficNodeTypeEnum
const (
	TrafficNodeTypeVisible      TrafficNodeTypeEnum = "VISIBLE"
	TrafficNodeTypeAccessDenied TrafficNodeTypeEnum = "ACCESS_DENIED"
)

var mappingTrafficNodeTypeEnum = map[string]TrafficNodeTypeEnum{
	"VISIBLE":       TrafficNodeTypeVisible,
	"ACCESS_DENIED": TrafficNodeTypeAccessDenied,
}

var mappingTrafficNodeTypeEnumLowerCase = map[string]TrafficNodeTypeEnum{
	"visible":       TrafficNodeTypeVisible,
	"access_denied": TrafficNodeTypeAccessDenied,
}

// GetTrafficNodeTypeEnumValues Enumerates the set of values for TrafficNodeTypeEnum
func GetTrafficNodeTypeEnumValues() []TrafficNodeTypeEnum {
	values := make([]TrafficNodeTypeEnum, 0)
	for _, v := range mappingTrafficNodeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTrafficNodeTypeEnumStringValues Enumerates the set of values in String for TrafficNodeTypeEnum
func GetTrafficNodeTypeEnumStringValues() []string {
	return []string{
		"VISIBLE",
		"ACCESS_DENIED",
	}
}

// GetMappingTrafficNodeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTrafficNodeTypeEnum(val string) (TrafficNodeTypeEnum, bool) {
	enum, ok := mappingTrafficNodeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
