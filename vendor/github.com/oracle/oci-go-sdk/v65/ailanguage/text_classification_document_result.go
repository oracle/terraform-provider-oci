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

// TextClassificationDocumentResult The document response for test classification detect call.
type TextClassificationDocumentResult struct {

	// Document unique identifier defined by the user.
	Key *string `mandatory:"true" json:"key"`

	// List of detected text classes.
	TextClassification []TextClassification `mandatory:"true" json:"textClassification"`

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
	LanguageCode *string `mandatory:"true" json:"languageCode"`
}

func (m TextClassificationDocumentResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TextClassificationDocumentResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
