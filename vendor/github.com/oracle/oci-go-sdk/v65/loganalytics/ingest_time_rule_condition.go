// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IngestTimeRuleCondition The condition(s) to evaluate for an ingest time rule.
type IngestTimeRuleCondition interface {
}

type ingesttimerulecondition struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *ingesttimerulecondition) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleringesttimerulecondition ingesttimerulecondition
	s := struct {
		Model Unmarshaleringesttimerulecondition
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *ingesttimerulecondition) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "FIELD":
		mm := IngestTimeRuleFieldCondition{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for IngestTimeRuleCondition: %s.", m.Kind)
		return *m, nil
	}
}

func (m ingesttimerulecondition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ingesttimerulecondition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IngestTimeRuleConditionKindEnum Enum with underlying type: string
type IngestTimeRuleConditionKindEnum string

// Set of constants representing the allowable values for IngestTimeRuleConditionKindEnum
const (
	IngestTimeRuleConditionKindField IngestTimeRuleConditionKindEnum = "FIELD"
)

var mappingIngestTimeRuleConditionKindEnum = map[string]IngestTimeRuleConditionKindEnum{
	"FIELD": IngestTimeRuleConditionKindField,
}

var mappingIngestTimeRuleConditionKindEnumLowerCase = map[string]IngestTimeRuleConditionKindEnum{
	"field": IngestTimeRuleConditionKindField,
}

// GetIngestTimeRuleConditionKindEnumValues Enumerates the set of values for IngestTimeRuleConditionKindEnum
func GetIngestTimeRuleConditionKindEnumValues() []IngestTimeRuleConditionKindEnum {
	values := make([]IngestTimeRuleConditionKindEnum, 0)
	for _, v := range mappingIngestTimeRuleConditionKindEnum {
		values = append(values, v)
	}
	return values
}

// GetIngestTimeRuleConditionKindEnumStringValues Enumerates the set of values in String for IngestTimeRuleConditionKindEnum
func GetIngestTimeRuleConditionKindEnumStringValues() []string {
	return []string{
		"FIELD",
	}
}

// GetMappingIngestTimeRuleConditionKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIngestTimeRuleConditionKindEnum(val string) (IngestTimeRuleConditionKindEnum, bool) {
	enum, ok := mappingIngestTimeRuleConditionKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
