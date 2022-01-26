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

// StructuredSearchDetails The model for a request that uses Search's structured query language.
type StructuredSearchDetails struct {

	// The structured query describing which resources to search for.
	Query *string `mandatory:"true" json:"query"`

	// The type of matching context returned in the response.
	MatchingContextType MatchingContextTypeEnumEnum `mandatory:"false" json:"matchingContextType,omitempty"`
}

//GetMatchingContextType returns MatchingContextType
func (m StructuredSearchDetails) GetMatchingContextType() MatchingContextTypeEnumEnum {
	return m.MatchingContextType
}

func (m StructuredSearchDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m StructuredSearchDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStructuredSearchDetails StructuredSearchDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeStructuredSearchDetails
	}{
		"Structured",
		(MarshalTypeStructuredSearchDetails)(m),
	}

	return json.Marshal(&s)
}
