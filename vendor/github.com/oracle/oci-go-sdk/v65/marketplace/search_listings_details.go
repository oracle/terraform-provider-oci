// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SearchListingsDetails A base request type that contains common criteria for Marketplace Search Listings details.
type SearchListingsDetails interface {

	// The type of matching context returned in the response. If you specify HIGHLIGHTS, then the service will highlight fragments in its response. The default value is NONE.
	GetMatchingContextType() MatchingContextTypeEnumEnum
}

type searchlistingsdetails struct {
	JsonData            []byte
	MatchingContextType MatchingContextTypeEnumEnum `mandatory:"false" json:"matchingContextType,omitempty"`
	Type                string                      `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *searchlistingsdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersearchlistingsdetails searchlistingsdetails
	s := struct {
		Model Unmarshalersearchlistingsdetails
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
func (m *searchlistingsdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

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
		common.Logf("Recieved unsupported enum value for SearchListingsDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetMatchingContextType returns MatchingContextType
func (m searchlistingsdetails) GetMatchingContextType() MatchingContextTypeEnumEnum {
	return m.MatchingContextType
}

func (m searchlistingsdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m searchlistingsdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMatchingContextTypeEnumEnum(string(m.MatchingContextType)); !ok && m.MatchingContextType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MatchingContextType: %s. Supported values are: %s.", m.MatchingContextType, strings.Join(GetMatchingContextTypeEnumEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
