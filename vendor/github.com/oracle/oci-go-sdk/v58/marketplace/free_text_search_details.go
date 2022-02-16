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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FreeTextSearchDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMatchingContextTypeEnumEnum(string(m.MatchingContextType)); !ok && m.MatchingContextType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MatchingContextType: %s. Supported values are: %s.", m.MatchingContextType, strings.Join(GetMatchingContextTypeEnumEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
