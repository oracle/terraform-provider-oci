// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Search Service API
//
// Search for resources in your cloud network.
//

package resourcesearch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SearchDetails A base request type that contains common criteria for searching for resources.
type SearchDetails interface {

	// The type of matching context returned in the response. If you specify `HIGHLIGHTS`, then the service will highlight fragments in its response. (For more information, see ResourceSummary.searchContext and SearchContext.) The default setting is `NONE`.
	GetMatchingContextType() SearchDetailsMatchingContextTypeEnum
}

type searchdetails struct {
	JsonData            []byte
	MatchingContextType SearchDetailsMatchingContextTypeEnum `mandatory:"false" json:"matchingContextType,omitempty"`
	Type                string                               `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *searchdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersearchdetails searchdetails
	s := struct {
		Model Unmarshalersearchdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.MatchingContextType = s.Model.MatchingContextType
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *searchdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "Structured":
		mm := StructuredSearchDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FreeText":
		mm := FreeTextSearchDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for SearchDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetMatchingContextType returns MatchingContextType
func (m searchdetails) GetMatchingContextType() SearchDetailsMatchingContextTypeEnum {
	return m.MatchingContextType
}

func (m searchdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m searchdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSearchDetailsMatchingContextTypeEnum(string(m.MatchingContextType)); !ok && m.MatchingContextType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MatchingContextType: %s. Supported values are: %s.", m.MatchingContextType, strings.Join(GetSearchDetailsMatchingContextTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SearchDetailsMatchingContextTypeEnum Enum with underlying type: string
type SearchDetailsMatchingContextTypeEnum string

// Set of constants representing the allowable values for SearchDetailsMatchingContextTypeEnum
const (
	SearchDetailsMatchingContextTypeNone       SearchDetailsMatchingContextTypeEnum = "NONE"
	SearchDetailsMatchingContextTypeHighlights SearchDetailsMatchingContextTypeEnum = "HIGHLIGHTS"
)

var mappingSearchDetailsMatchingContextTypeEnum = map[string]SearchDetailsMatchingContextTypeEnum{
	"NONE":       SearchDetailsMatchingContextTypeNone,
	"HIGHLIGHTS": SearchDetailsMatchingContextTypeHighlights,
}

var mappingSearchDetailsMatchingContextTypeEnumLowerCase = map[string]SearchDetailsMatchingContextTypeEnum{
	"none":       SearchDetailsMatchingContextTypeNone,
	"highlights": SearchDetailsMatchingContextTypeHighlights,
}

// GetSearchDetailsMatchingContextTypeEnumValues Enumerates the set of values for SearchDetailsMatchingContextTypeEnum
func GetSearchDetailsMatchingContextTypeEnumValues() []SearchDetailsMatchingContextTypeEnum {
	values := make([]SearchDetailsMatchingContextTypeEnum, 0)
	for _, v := range mappingSearchDetailsMatchingContextTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSearchDetailsMatchingContextTypeEnumStringValues Enumerates the set of values in String for SearchDetailsMatchingContextTypeEnum
func GetSearchDetailsMatchingContextTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"HIGHLIGHTS",
	}
}

// GetMappingSearchDetailsMatchingContextTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSearchDetailsMatchingContextTypeEnum(val string) (SearchDetailsMatchingContextTypeEnum, bool) {
	enum, ok := mappingSearchDetailsMatchingContextTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
