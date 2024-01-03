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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ModelDetails Possible model types
type ModelDetails interface {

	// supported language default value is en
	GetLanguageCode() *string
}

type modeldetails struct {
	JsonData     []byte
	LanguageCode *string `mandatory:"false" json:"languageCode"`
	ModelType    string  `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *modeldetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermodeldetails modeldetails
	s := struct {
		Model Unmarshalermodeldetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.LanguageCode = s.Model.LanguageCode
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *modeldetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "PRE_TRAINED_KEYPHRASE_EXTRACTION":
		mm := PreTrainedKeyPhraseExtractionModelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRE_TRAINED_HEALTH_NLU":
		mm := PreTrainedHealthNluModelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRE_TRAINED_UNIVERSAL":
		mm := PreTrainedUniversalModel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NAMED_ENTITY_RECOGNITION":
		mm := NamedEntityRecognitionModelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRE_TRAINED_LANGUAGE_DETECTION":
		mm := PreTrainedLanguageDetectionModelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRE_TRAINED_NAMED_ENTITY_RECOGNITION":
		mm := PreTrainedNamedEntityRecognitionModelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRE_TRAINED_SENTIMENT_ANALYSIS":
		mm := PreTrainedSentimentAnalysisModelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRE_TRAINED_TEXT_CLASSIFICATION":
		mm := PreTrainedTextClassificationModelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TEXT_CLASSIFICATION":
		mm := TextClassificationModelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRE_TRAINED_SUMMARIZATION":
		mm := PreTrainedSummarization{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRE_TRAINED_PII":
		mm := PreTrainedPiiModelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ModelDetails: %s.", m.ModelType)
		return *m, nil
	}
}

// GetLanguageCode returns LanguageCode
func (m modeldetails) GetLanguageCode() *string {
	return m.LanguageCode
}

func (m modeldetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m modeldetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ModelDetailsModelTypeEnum Enum with underlying type: string
type ModelDetailsModelTypeEnum string

// Set of constants representing the allowable values for ModelDetailsModelTypeEnum
const (
	ModelDetailsModelTypeNamedEntityRecognition           ModelDetailsModelTypeEnum = "NAMED_ENTITY_RECOGNITION"
	ModelDetailsModelTypeTextClassification               ModelDetailsModelTypeEnum = "TEXT_CLASSIFICATION"
	ModelDetailsModelTypePreTrainedNamedEntityRecognition ModelDetailsModelTypeEnum = "PRE_TRAINED_NAMED_ENTITY_RECOGNITION"
	ModelDetailsModelTypePreTrainedTextClassification     ModelDetailsModelTypeEnum = "PRE_TRAINED_TEXT_CLASSIFICATION"
	ModelDetailsModelTypePreTrainedSentimentAnalysis      ModelDetailsModelTypeEnum = "PRE_TRAINED_SENTIMENT_ANALYSIS"
	ModelDetailsModelTypePreTrainedKeyphraseExtraction    ModelDetailsModelTypeEnum = "PRE_TRAINED_KEYPHRASE_EXTRACTION"
	ModelDetailsModelTypePreTrainedLanguageDetection      ModelDetailsModelTypeEnum = "PRE_TRAINED_LANGUAGE_DETECTION"
	ModelDetailsModelTypePreTrainedPii                    ModelDetailsModelTypeEnum = "PRE_TRAINED_PII"
	ModelDetailsModelTypePreTrainedTranslation            ModelDetailsModelTypeEnum = "PRE_TRAINED_TRANSLATION"
	ModelDetailsModelTypePreTrainedHealthNlu              ModelDetailsModelTypeEnum = "PRE_TRAINED_HEALTH_NLU"
	ModelDetailsModelTypePreTrainedSummarization          ModelDetailsModelTypeEnum = "PRE_TRAINED_SUMMARIZATION"
	ModelDetailsModelTypePreTrainedUniversal              ModelDetailsModelTypeEnum = "PRE_TRAINED_UNIVERSAL"
)

var mappingModelDetailsModelTypeEnum = map[string]ModelDetailsModelTypeEnum{
	"NAMED_ENTITY_RECOGNITION":             ModelDetailsModelTypeNamedEntityRecognition,
	"TEXT_CLASSIFICATION":                  ModelDetailsModelTypeTextClassification,
	"PRE_TRAINED_NAMED_ENTITY_RECOGNITION": ModelDetailsModelTypePreTrainedNamedEntityRecognition,
	"PRE_TRAINED_TEXT_CLASSIFICATION":      ModelDetailsModelTypePreTrainedTextClassification,
	"PRE_TRAINED_SENTIMENT_ANALYSIS":       ModelDetailsModelTypePreTrainedSentimentAnalysis,
	"PRE_TRAINED_KEYPHRASE_EXTRACTION":     ModelDetailsModelTypePreTrainedKeyphraseExtraction,
	"PRE_TRAINED_LANGUAGE_DETECTION":       ModelDetailsModelTypePreTrainedLanguageDetection,
	"PRE_TRAINED_PII":                      ModelDetailsModelTypePreTrainedPii,
	"PRE_TRAINED_TRANSLATION":              ModelDetailsModelTypePreTrainedTranslation,
	"PRE_TRAINED_HEALTH_NLU":               ModelDetailsModelTypePreTrainedHealthNlu,
	"PRE_TRAINED_SUMMARIZATION":            ModelDetailsModelTypePreTrainedSummarization,
	"PRE_TRAINED_UNIVERSAL":                ModelDetailsModelTypePreTrainedUniversal,
}

var mappingModelDetailsModelTypeEnumLowerCase = map[string]ModelDetailsModelTypeEnum{
	"named_entity_recognition":             ModelDetailsModelTypeNamedEntityRecognition,
	"text_classification":                  ModelDetailsModelTypeTextClassification,
	"pre_trained_named_entity_recognition": ModelDetailsModelTypePreTrainedNamedEntityRecognition,
	"pre_trained_text_classification":      ModelDetailsModelTypePreTrainedTextClassification,
	"pre_trained_sentiment_analysis":       ModelDetailsModelTypePreTrainedSentimentAnalysis,
	"pre_trained_keyphrase_extraction":     ModelDetailsModelTypePreTrainedKeyphraseExtraction,
	"pre_trained_language_detection":       ModelDetailsModelTypePreTrainedLanguageDetection,
	"pre_trained_pii":                      ModelDetailsModelTypePreTrainedPii,
	"pre_trained_translation":              ModelDetailsModelTypePreTrainedTranslation,
	"pre_trained_health_nlu":               ModelDetailsModelTypePreTrainedHealthNlu,
	"pre_trained_summarization":            ModelDetailsModelTypePreTrainedSummarization,
	"pre_trained_universal":                ModelDetailsModelTypePreTrainedUniversal,
}

// GetModelDetailsModelTypeEnumValues Enumerates the set of values for ModelDetailsModelTypeEnum
func GetModelDetailsModelTypeEnumValues() []ModelDetailsModelTypeEnum {
	values := make([]ModelDetailsModelTypeEnum, 0)
	for _, v := range mappingModelDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetModelDetailsModelTypeEnumStringValues Enumerates the set of values in String for ModelDetailsModelTypeEnum
func GetModelDetailsModelTypeEnumStringValues() []string {
	return []string{
		"NAMED_ENTITY_RECOGNITION",
		"TEXT_CLASSIFICATION",
		"PRE_TRAINED_NAMED_ENTITY_RECOGNITION",
		"PRE_TRAINED_TEXT_CLASSIFICATION",
		"PRE_TRAINED_SENTIMENT_ANALYSIS",
		"PRE_TRAINED_KEYPHRASE_EXTRACTION",
		"PRE_TRAINED_LANGUAGE_DETECTION",
		"PRE_TRAINED_PII",
		"PRE_TRAINED_TRANSLATION",
		"PRE_TRAINED_HEALTH_NLU",
		"PRE_TRAINED_SUMMARIZATION",
		"PRE_TRAINED_UNIVERSAL",
	}
}

// GetMappingModelDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelDetailsModelTypeEnum(val string) (ModelDetailsModelTypeEnum, bool) {
	enum, ok := mappingModelDetailsModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
