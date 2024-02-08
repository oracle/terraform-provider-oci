// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// BatchLanguageTranslationDetails The documents details for translation call.
type BatchLanguageTranslationDetails struct {

	// List of documents for translation.
	Documents []TextDocument `mandatory:"true" json:"documents"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that calls the API, inference will be served from pre trained model
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Language code supported
	// - auto : Automatically detect language
	// - ar : Arabic
	// - pt-BR : Brazilian Portuguese
	// - cs : Czech
	// - da : Danish
	// - nl : Dutch
	// - en : English
	// - fi : Finnish
	// - fr : French
	// - fr-CA : Canadian French
	// - de : German
	// - it : Italian
	// - ja : Japanese
	// - ko : Korean
	// - no : Norwegian
	// - pl : Polish
	// - ro : Romanian
	// - zh-CN : Simplified Chinese
	// - es : Spanish
	// - sv : Swedish
	// - zh-TW : Traditional Chinese
	// - tr : Turkish
	// - el : Greek
	// - he : Hebrew
	TargetLanguageCode *string `mandatory:"false" json:"targetLanguageCode"`
}

func (m BatchLanguageTranslationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BatchLanguageTranslationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
