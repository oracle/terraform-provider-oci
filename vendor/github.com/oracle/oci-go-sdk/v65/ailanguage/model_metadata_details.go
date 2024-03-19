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

// ModelMetadataDetails training model details
// For this release only one model is allowed to be input here.
// One of the three modelType, ModelId, endpointId should be given other wise error will be thrown from API
type ModelMetadataDetails struct {

	// model type to used for inference allowed values are
	// - LANGUAGE_SENTIMENT_ANALYSIS
	// - LANGUAGE_DETECTION
	// - TEXT_CLASSIFICATION
	// - NAMED_ENTITY_RECOGNITION
	// - KEY_PHRASE_EXTRACTION
	// - LANGUAGE_PII_ENTITIES
	// - LANGUAGE_TRANSLATION
	ModelType *string `mandatory:"false" json:"modelType"`

	// Unique identifier model OCID that should be used for inference
	ModelId *string `mandatory:"false" json:"modelId"`

	// Unique identifier endpoint OCID that should be used for inference
	EndpointId *string `mandatory:"false" json:"endpointId"`

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
	LanguageCode *string `mandatory:"false" json:"languageCode"`

	// model configuration details
	// For PII :  < ENTITY_TYPE , ConfigurationDetails>
	// ex."ORACLE":{ "mode" : "MASK","maskingCharacter" : "&","leaveCharactersUnmasked": 3,"isUnmaskedFromEnd" : true  }
	// For language translation : { "targetLanguageCodes" : ConfigurationDetails}
	Configuration map[string]ConfigurationDetails `mandatory:"false" json:"configuration"`
}

func (m ModelMetadataDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ModelMetadataDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
