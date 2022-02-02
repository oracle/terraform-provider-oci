// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// FreeTextSearchDetails A request containing arbitrary text that must be present in the Marketplace Applications.
type FreeTextSearchDetails struct {

	// The text to search for.
	Text *string `mandatory:"true" json:"text"`

	// The type of matching context returned in the response. If you specify HIGHLIGHTS, then the service will highlight fragments in its response. The default value is NONE.
	MatchingContextType MatchingContextTypeEnumEnum `mandatory:"false" json:"matchingContextType,omitempty"`
}

//GetMatchingContextType returns MatchingContextType
func (m FreeTextSearchDetails) GetMatchingContextType() MatchingContextTypeEnumEnum {
	return m.MatchingContextType
}

func (m FreeTextSearchDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m FreeTextSearchDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFreeTextSearchDetails FreeTextSearchDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeFreeTextSearchDetails
	}{
		"FreeText",
		(MarshalTypeFreeTextSearchDetails)(m),
	}

	return json.Marshal(&s)
}
