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

// OkeCanaryStrategy Specifies the required canary release strategy for OKE deployment.
type OkeCanaryStrategy interface {
}

type okecanarystrategy struct {
	JsonData     []byte
	StrategyType string `json:"strategyType"`
}

// UnmarshalJSON unmarshals json
func (m *okecanarystrategy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerokecanarystrategy okecanarystrategy
	s := struct {
		Model Unmarshalerokecanarystrategy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.StrategyType = s.Model.StrategyType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *okecanarystrategy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StrategyType {
	case "NGINX_CANARY_STRATEGY":
		mm := NginxCanaryStrategy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for OkeCanaryStrategy: %s.", m.StrategyType)
		return *m, nil
	}
}

func (m okecanarystrategy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m okecanarystrategy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OkeCanaryStrategyStrategyTypeEnum Enum with underlying type: string
type OkeCanaryStrategyStrategyTypeEnum string

// Set of constants representing the allowable values for OkeCanaryStrategyStrategyTypeEnum
const (
	OkeCanaryStrategyStrategyTypeNginxCanaryStrategy OkeCanaryStrategyStrategyTypeEnum = "NGINX_CANARY_STRATEGY"
)

var mappingOkeCanaryStrategyStrategyTypeEnum = map[string]OkeCanaryStrategyStrategyTypeEnum{
	"NGINX_CANARY_STRATEGY": OkeCanaryStrategyStrategyTypeNginxCanaryStrategy,
}

var mappingOkeCanaryStrategyStrategyTypeEnumLowerCase = map[string]OkeCanaryStrategyStrategyTypeEnum{
	"nginx_canary_strategy": OkeCanaryStrategyStrategyTypeNginxCanaryStrategy,
}

// GetOkeCanaryStrategyStrategyTypeEnumValues Enumerates the set of values for OkeCanaryStrategyStrategyTypeEnum
func GetOkeCanaryStrategyStrategyTypeEnumValues() []OkeCanaryStrategyStrategyTypeEnum {
	values := make([]OkeCanaryStrategyStrategyTypeEnum, 0)
	for _, v := range mappingOkeCanaryStrategyStrategyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOkeCanaryStrategyStrategyTypeEnumStringValues Enumerates the set of values in String for OkeCanaryStrategyStrategyTypeEnum
func GetOkeCanaryStrategyStrategyTypeEnumStringValues() []string {
	return []string{
		"NGINX_CANARY_STRATEGY",
	}
}

// GetMappingOkeCanaryStrategyStrategyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOkeCanaryStrategyStrategyTypeEnum(val string) (OkeCanaryStrategyStrategyTypeEnum, bool) {
	enum, ok := mappingOkeCanaryStrategyStrategyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
