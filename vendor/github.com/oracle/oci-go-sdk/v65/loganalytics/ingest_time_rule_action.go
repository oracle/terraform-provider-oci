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

// IngestTimeRuleAction The action to be performed if the ingest time rule condition(s) are satisfied.
type IngestTimeRuleAction interface {
}

type ingesttimeruleaction struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *ingesttimeruleaction) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleringesttimeruleaction ingesttimeruleaction
	s := struct {
		Model Unmarshaleringesttimeruleaction
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *ingesttimeruleaction) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "METRIC_EXTRACTION":
		mm := IngestTimeRuleMetricExtractionAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for IngestTimeRuleAction: %s.", m.Type)
		return *m, nil
	}
}

func (m ingesttimeruleaction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ingesttimeruleaction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IngestTimeRuleActionTypeEnum Enum with underlying type: string
type IngestTimeRuleActionTypeEnum string

// Set of constants representing the allowable values for IngestTimeRuleActionTypeEnum
const (
	IngestTimeRuleActionTypeMetricExtraction IngestTimeRuleActionTypeEnum = "METRIC_EXTRACTION"
)

var mappingIngestTimeRuleActionTypeEnum = map[string]IngestTimeRuleActionTypeEnum{
	"METRIC_EXTRACTION": IngestTimeRuleActionTypeMetricExtraction,
}

var mappingIngestTimeRuleActionTypeEnumLowerCase = map[string]IngestTimeRuleActionTypeEnum{
	"metric_extraction": IngestTimeRuleActionTypeMetricExtraction,
}

// GetIngestTimeRuleActionTypeEnumValues Enumerates the set of values for IngestTimeRuleActionTypeEnum
func GetIngestTimeRuleActionTypeEnumValues() []IngestTimeRuleActionTypeEnum {
	values := make([]IngestTimeRuleActionTypeEnum, 0)
	for _, v := range mappingIngestTimeRuleActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIngestTimeRuleActionTypeEnumStringValues Enumerates the set of values in String for IngestTimeRuleActionTypeEnum
func GetIngestTimeRuleActionTypeEnumStringValues() []string {
	return []string{
		"METRIC_EXTRACTION",
	}
}

// GetMappingIngestTimeRuleActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIngestTimeRuleActionTypeEnum(val string) (IngestTimeRuleActionTypeEnum, bool) {
	enum, ok := mappingIngestTimeRuleActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
