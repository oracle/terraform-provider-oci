// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// CollectionRuleValidationResult Collection rule validation result.
type CollectionRuleValidationResult interface {

	// explanation of the validation status.
	GetStatusDescription() *string
}

type collectionrulevalidationresult struct {
	JsonData          []byte
	StatusDescription *string `mandatory:"false" json:"statusDescription"`
	Status            string  `json:"status"`
}

// UnmarshalJSON unmarshals json
func (m *collectionrulevalidationresult) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercollectionrulevalidationresult collectionrulevalidationresult
	s := struct {
		Model Unmarshalercollectionrulevalidationresult
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.StatusDescription = s.Model.StatusDescription
	m.Status = s.Model.Status

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *collectionrulevalidationresult) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Status {
	case "FAILED":
		mm := CollectionRuleFailedValidationResult{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SUCCESS":
		mm := CollectionRuleSuccessValidationResult{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CollectionRuleValidationResult: %s.", m.Status)
		return *m, nil
	}
}

// GetStatusDescription returns StatusDescription
func (m collectionrulevalidationresult) GetStatusDescription() *string {
	return m.StatusDescription
}

func (m collectionrulevalidationresult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m collectionrulevalidationresult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CollectionRuleValidationResultStatusEnum Enum with underlying type: string
type CollectionRuleValidationResultStatusEnum string

// Set of constants representing the allowable values for CollectionRuleValidationResultStatusEnum
const (
	CollectionRuleValidationResultStatusSuccess CollectionRuleValidationResultStatusEnum = "SUCCESS"
	CollectionRuleValidationResultStatusFailed  CollectionRuleValidationResultStatusEnum = "FAILED"
)

var mappingCollectionRuleValidationResultStatusEnum = map[string]CollectionRuleValidationResultStatusEnum{
	"SUCCESS": CollectionRuleValidationResultStatusSuccess,
	"FAILED":  CollectionRuleValidationResultStatusFailed,
}

var mappingCollectionRuleValidationResultStatusEnumLowerCase = map[string]CollectionRuleValidationResultStatusEnum{
	"success": CollectionRuleValidationResultStatusSuccess,
	"failed":  CollectionRuleValidationResultStatusFailed,
}

// GetCollectionRuleValidationResultStatusEnumValues Enumerates the set of values for CollectionRuleValidationResultStatusEnum
func GetCollectionRuleValidationResultStatusEnumValues() []CollectionRuleValidationResultStatusEnum {
	values := make([]CollectionRuleValidationResultStatusEnum, 0)
	for _, v := range mappingCollectionRuleValidationResultStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCollectionRuleValidationResultStatusEnumStringValues Enumerates the set of values in String for CollectionRuleValidationResultStatusEnum
func GetCollectionRuleValidationResultStatusEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"FAILED",
	}
}

// GetMappingCollectionRuleValidationResultStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCollectionRuleValidationResultStatusEnum(val string) (CollectionRuleValidationResultStatusEnum, bool) {
	enum, ok := mappingCollectionRuleValidationResultStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
