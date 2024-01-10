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

// TrafficRuleTargetDetails Target of the traffic router rule.
type TrafficRuleTargetDetails interface {
}

type trafficruletargetdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *trafficruletargetdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertrafficruletargetdetails trafficruletargetdetails
	s := struct {
		Model Unmarshalertrafficruletargetdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *trafficruletargetdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "VIRTUAL_DEPLOYMENT":
		mm := VirtualDeploymentTrafficRuleTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VIRTUAL_SERVICE":
		mm := VirtualServiceTrafficRuleTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for TrafficRuleTargetDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m trafficruletargetdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m trafficruletargetdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TrafficRuleTargetDetailsTypeEnum Enum with underlying type: string
type TrafficRuleTargetDetailsTypeEnum string

// Set of constants representing the allowable values for TrafficRuleTargetDetailsTypeEnum
const (
	TrafficRuleTargetDetailsTypeDeployment TrafficRuleTargetDetailsTypeEnum = "VIRTUAL_DEPLOYMENT"
	TrafficRuleTargetDetailsTypeService    TrafficRuleTargetDetailsTypeEnum = "VIRTUAL_SERVICE"
)

var mappingTrafficRuleTargetDetailsTypeEnum = map[string]TrafficRuleTargetDetailsTypeEnum{
	"VIRTUAL_DEPLOYMENT": TrafficRuleTargetDetailsTypeDeployment,
	"VIRTUAL_SERVICE":    TrafficRuleTargetDetailsTypeService,
}

var mappingTrafficRuleTargetDetailsTypeEnumLowerCase = map[string]TrafficRuleTargetDetailsTypeEnum{
	"virtual_deployment": TrafficRuleTargetDetailsTypeDeployment,
	"virtual_service":    TrafficRuleTargetDetailsTypeService,
}

// GetTrafficRuleTargetDetailsTypeEnumValues Enumerates the set of values for TrafficRuleTargetDetailsTypeEnum
func GetTrafficRuleTargetDetailsTypeEnumValues() []TrafficRuleTargetDetailsTypeEnum {
	values := make([]TrafficRuleTargetDetailsTypeEnum, 0)
	for _, v := range mappingTrafficRuleTargetDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTrafficRuleTargetDetailsTypeEnumStringValues Enumerates the set of values in String for TrafficRuleTargetDetailsTypeEnum
func GetTrafficRuleTargetDetailsTypeEnumStringValues() []string {
	return []string{
		"VIRTUAL_DEPLOYMENT",
		"VIRTUAL_SERVICE",
	}
}

// GetMappingTrafficRuleTargetDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTrafficRuleTargetDetailsTypeEnum(val string) (TrafficRuleTargetDetailsTypeEnum, bool) {
	enum, ok := mappingTrafficRuleTargetDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
