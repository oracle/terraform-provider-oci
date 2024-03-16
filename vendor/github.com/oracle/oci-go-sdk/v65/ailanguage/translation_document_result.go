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

// TranslationDocumentResult The document response for translation call.
type TranslationDocumentResult struct {

	// Document unique identifier defined by the user.
	Key *string `mandatory:"true" json:"key"`

	// Translated text in selected target language.
	TranslatedText *string `mandatory:"true" json:"translatedText"`

	// Language code supported
	// Automatically detect language - auto
	// Arabic - ar
	// Brazilian Portuguese -  pt-BR
	// Canadian French - fr-CA
	// Croatian - hr
	// Czech - cs
	// Danish - da
	// Dutch - nl
	// English - en
	// Finnish - fi
	// French - fr
	// German - de
	// Greek - el
	// Hebrew - he
	// Hungarian - hu
	// Italian - it
	// Japanese - ja
	// Korean - ko
	// Norwegian - no
	// Polish - pl
	// Portuguese - pt
	// Romanian - ro
	// Russian - ru
	// Simplified Chinese - zh-CN
	// Slovak - sk
	// Slovenian - sl
	// Spanish - es
	// Swedish - sv
	// Thai - th
	// Traditional Chinese - zh-TW
	// Turkish - tr
	// Vietnamese - vi
	SourceLanguageCode *string `mandatory:"true" json:"sourceLanguageCode"`

	// Language code supported
	// Arabic - ar
	// Brazilian Portuguese -  pt-BR
	// Canadian French - fr-CA
	// Croatian - hr
	// Czech - cs
	// Danish - da
	// Dutch - nl
	// English - en
	// Finnish - fi
	// French - fr
	// German - de
	// Greek - el
	// Hebrew - he
	// Hungarian - hu
	// Italian - it
	// Japanese - ja
	// Korean - ko
	// Norwegian - no
	// Polish - pl
	// Portuguese - pt
	// Romanian - ro
	// Russian - ru
	// Simplified Chinese - zh-CN
	// Slovak - sk
	// Slovenian - sl
	// Spanish - es
	// Swedish - sv
	// Thai - th
	// Traditional Chinese - zh-TW
	// Turkish - tr
	// Vietnamese - vi
	TargetLanguageCode *string `mandatory:"true" json:"targetLanguageCode"`
}

func (m TranslationDocumentResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TranslationDocumentResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
