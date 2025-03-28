// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Language API
//
// OCI Language Service solutions can help enterprise customers integrate AI into their products immediately using our proven,
// pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BatchDetectDominantLanguageDetails The documents details for language detect call.
type BatchDetectDominantLanguageDetails struct {

	// List of Documents for detect language.
	Documents []DominantLanguageDocument `mandatory:"true" json:"documents"`

	// Unique name across user tenancy in a region to identify an endpoint to be used for inferencing.
	Alias *string `mandatory:"false" json:"alias"`

	// Specifies whether to consider or ignore transliteration. For example "hi, aap kaise ho? sab kuch teek hai? I will call you tomorrow." would be detected as English when ignore transliteration=true, Hindi when ignoreTransliteration=false.
	ShouldIgnoreTransliteration *bool `mandatory:"false" json:"shouldIgnoreTransliteration"`

	// default value is None.
	// Specifies maximum number of characters to consider for determining the dominant language.
	// If unspecified, then optimum number characters will be considered.
	// If 0 is specified then all the characters are used to determine the language.
	// If the value is greater than 0, then specified number of characters will be considered from the beginning of the text.
	CharsToConsider *int `mandatory:"false" json:"charsToConsider"`

	// The endpoint which have to be used for inferencing. If endpointId and compartmentId is provided, then inference will be served from custom model which is mapped to this Endpoint.
	EndpointId *string `mandatory:"false" json:"endpointId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that calls the API, inference will be served from pre trained model
	CompartmentId *string `mandatory:"false" json:"compartmentId"`
}

func (m BatchDetectDominantLanguageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BatchDetectDominantLanguageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
