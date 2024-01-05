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

// OkeBlueGreenStrategy Specifies the required blue-green release strategy for OKE deployment.
type OkeBlueGreenStrategy interface {
}

type okebluegreenstrategy struct {
	JsonData     []byte
	StrategyType string `json:"strategyType"`
}

// UnmarshalJSON unmarshals json
func (m *okebluegreenstrategy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerokebluegreenstrategy okebluegreenstrategy
	s := struct {
		Model Unmarshalerokebluegreenstrategy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.StrategyType = s.Model.StrategyType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *okebluegreenstrategy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StrategyType {
	case "NGINX_BLUE_GREEN_STRATEGY":
		mm := NginxBlueGreenStrategy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for OkeBlueGreenStrategy: %s.", m.StrategyType)
		return *m, nil
	}
}

func (m okebluegreenstrategy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m okebluegreenstrategy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OkeBlueGreenStrategyStrategyTypeEnum Enum with underlying type: string
type OkeBlueGreenStrategyStrategyTypeEnum string

// Set of constants representing the allowable values for OkeBlueGreenStrategyStrategyTypeEnum
const (
	OkeBlueGreenStrategyStrategyTypeNginxBlueGreenStrategy OkeBlueGreenStrategyStrategyTypeEnum = "NGINX_BLUE_GREEN_STRATEGY"
)

var mappingOkeBlueGreenStrategyStrategyTypeEnum = map[string]OkeBlueGreenStrategyStrategyTypeEnum{
	"NGINX_BLUE_GREEN_STRATEGY": OkeBlueGreenStrategyStrategyTypeNginxBlueGreenStrategy,
}

var mappingOkeBlueGreenStrategyStrategyTypeEnumLowerCase = map[string]OkeBlueGreenStrategyStrategyTypeEnum{
	"nginx_blue_green_strategy": OkeBlueGreenStrategyStrategyTypeNginxBlueGreenStrategy,
}

// GetOkeBlueGreenStrategyStrategyTypeEnumValues Enumerates the set of values for OkeBlueGreenStrategyStrategyTypeEnum
func GetOkeBlueGreenStrategyStrategyTypeEnumValues() []OkeBlueGreenStrategyStrategyTypeEnum {
	values := make([]OkeBlueGreenStrategyStrategyTypeEnum, 0)
	for _, v := range mappingOkeBlueGreenStrategyStrategyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOkeBlueGreenStrategyStrategyTypeEnumStringValues Enumerates the set of values in String for OkeBlueGreenStrategyStrategyTypeEnum
func GetOkeBlueGreenStrategyStrategyTypeEnumStringValues() []string {
	return []string{
		"NGINX_BLUE_GREEN_STRATEGY",
	}
}

// GetMappingOkeBlueGreenStrategyStrategyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOkeBlueGreenStrategyStrategyTypeEnum(val string) (OkeBlueGreenStrategyStrategyTypeEnum, bool) {
	enum, ok := mappingOkeBlueGreenStrategyStrategyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
