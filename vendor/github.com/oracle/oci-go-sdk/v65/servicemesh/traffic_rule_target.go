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

// TrafficRuleTarget Target of the traffic router rule.
type TrafficRuleTarget interface {
}

type trafficruletarget struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *trafficruletarget) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertrafficruletarget trafficruletarget
	s := struct {
		Model Unmarshalertrafficruletarget
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *trafficruletarget) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "VIRTUAL_DEPLOYMENT":
		mm := VirtualDeploymentTrafficRuleTarget{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VIRTUAL_SERVICE":
		mm := VirtualServiceTrafficRuleTarget{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for TrafficRuleTarget: %s.", m.Type)
		return *m, nil
	}
}

func (m trafficruletarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m trafficruletarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TrafficRuleTargetTypeEnum Enum with underlying type: string
type TrafficRuleTargetTypeEnum string

// Set of constants representing the allowable values for TrafficRuleTargetTypeEnum
const (
	TrafficRuleTargetTypeDeployment TrafficRuleTargetTypeEnum = "VIRTUAL_DEPLOYMENT"
	TrafficRuleTargetTypeService    TrafficRuleTargetTypeEnum = "VIRTUAL_SERVICE"
)

var mappingTrafficRuleTargetTypeEnum = map[string]TrafficRuleTargetTypeEnum{
	"VIRTUAL_DEPLOYMENT": TrafficRuleTargetTypeDeployment,
	"VIRTUAL_SERVICE":    TrafficRuleTargetTypeService,
}

var mappingTrafficRuleTargetTypeEnumLowerCase = map[string]TrafficRuleTargetTypeEnum{
	"virtual_deployment": TrafficRuleTargetTypeDeployment,
	"virtual_service":    TrafficRuleTargetTypeService,
}

// GetTrafficRuleTargetTypeEnumValues Enumerates the set of values for TrafficRuleTargetTypeEnum
func GetTrafficRuleTargetTypeEnumValues() []TrafficRuleTargetTypeEnum {
	values := make([]TrafficRuleTargetTypeEnum, 0)
	for _, v := range mappingTrafficRuleTargetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTrafficRuleTargetTypeEnumStringValues Enumerates the set of values in String for TrafficRuleTargetTypeEnum
func GetTrafficRuleTargetTypeEnumStringValues() []string {
	return []string{
		"VIRTUAL_DEPLOYMENT",
		"VIRTUAL_SERVICE",
	}
}

// GetMappingTrafficRuleTargetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTrafficRuleTargetTypeEnum(val string) (TrafficRuleTargetTypeEnum, bool) {
	enum, ok := mappingTrafficRuleTargetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
