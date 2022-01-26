// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// SearchListingsDetails The base model for a Search Listings details.
type SearchListingsDetails interface {

	// The type of matching context returned in the response.
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
		return *m, nil
	}
}

//GetMatchingContextType returns MatchingContextType
func (m searchlistingsdetails) GetMatchingContextType() MatchingContextTypeEnumEnum {
	return m.MatchingContextType
}

func (m searchlistingsdetails) String() string {
	return common.PointerString(m)
}
