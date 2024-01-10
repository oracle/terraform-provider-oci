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

// IngestTimeRuleFieldCondition The field condition(s) to evaluate for an ingest time rule.
type IngestTimeRuleFieldCondition struct {

	// The field name to be evaluated.
	FieldName *string `mandatory:"true" json:"fieldName"`

	// The field value to be evaluated.
	FieldValue *string `mandatory:"true" json:"fieldValue"`

	// Optional additional condition(s) to be evaluated.
	AdditionalConditions []IngestTimeRuleAdditionalFieldCondition `mandatory:"false" json:"additionalConditions"`

	// The operator to be used for evaluating the field.
	FieldOperator IngestTimeRuleFieldConditionFieldOperatorEnum `mandatory:"true" json:"fieldOperator"`
}

func (m IngestTimeRuleFieldCondition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IngestTimeRuleFieldCondition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIngestTimeRuleFieldConditionFieldOperatorEnum(string(m.FieldOperator)); !ok && m.FieldOperator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FieldOperator: %s. Supported values are: %s.", m.FieldOperator, strings.Join(GetIngestTimeRuleFieldConditionFieldOperatorEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m IngestTimeRuleFieldCondition) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeIngestTimeRuleFieldCondition IngestTimeRuleFieldCondition
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeIngestTimeRuleFieldCondition
	}{
		"FIELD",
		(MarshalTypeIngestTimeRuleFieldCondition)(m),
	}

	return json.Marshal(&s)
}

// IngestTimeRuleFieldConditionFieldOperatorEnum Enum with underlying type: string
type IngestTimeRuleFieldConditionFieldOperatorEnum string

// Set of constants representing the allowable values for IngestTimeRuleFieldConditionFieldOperatorEnum
const (
	IngestTimeRuleFieldConditionFieldOperatorEqual IngestTimeRuleFieldConditionFieldOperatorEnum = "EQUAL"
)

var mappingIngestTimeRuleFieldConditionFieldOperatorEnum = map[string]IngestTimeRuleFieldConditionFieldOperatorEnum{
	"EQUAL": IngestTimeRuleFieldConditionFieldOperatorEqual,
}

var mappingIngestTimeRuleFieldConditionFieldOperatorEnumLowerCase = map[string]IngestTimeRuleFieldConditionFieldOperatorEnum{
	"equal": IngestTimeRuleFieldConditionFieldOperatorEqual,
}

// GetIngestTimeRuleFieldConditionFieldOperatorEnumValues Enumerates the set of values for IngestTimeRuleFieldConditionFieldOperatorEnum
func GetIngestTimeRuleFieldConditionFieldOperatorEnumValues() []IngestTimeRuleFieldConditionFieldOperatorEnum {
	values := make([]IngestTimeRuleFieldConditionFieldOperatorEnum, 0)
	for _, v := range mappingIngestTimeRuleFieldConditionFieldOperatorEnum {
		values = append(values, v)
	}
	return values
}

// GetIngestTimeRuleFieldConditionFieldOperatorEnumStringValues Enumerates the set of values in String for IngestTimeRuleFieldConditionFieldOperatorEnum
func GetIngestTimeRuleFieldConditionFieldOperatorEnumStringValues() []string {
	return []string{
		"EQUAL",
	}
}

// GetMappingIngestTimeRuleFieldConditionFieldOperatorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIngestTimeRuleFieldConditionFieldOperatorEnum(val string) (IngestTimeRuleFieldConditionFieldOperatorEnum, bool) {
	enum, ok := mappingIngestTimeRuleFieldConditionFieldOperatorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
