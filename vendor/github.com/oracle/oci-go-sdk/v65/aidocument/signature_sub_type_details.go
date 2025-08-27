// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Document Understanding API
//
// Document AI helps customers perform various analysis on their documents. If a customer has lots of documents, they can process them in batch using asynchronous API endpoints.
//

package aidocument

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SignatureSubTypeDetails Signature Document Element Extraction model sub type details
type SignatureSubTypeDetails struct {

	// Sub type model based on the model type.
	// The allowed values are:
	// - `PRE_TRAINED_KEY_VALUE_EXTRACTION`
	// - `PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION`
	ModelType SignatureSubTypeDetailsModelTypeEnum `mandatory:"true" json:"modelType"`

	// The model sub type for DOCUMENT_ELEMENTS_EXTRACTION.
	// The allowed values are:
	//   - `QR_BAR_CODE`
	//   - `SIGNATURE`
	ModelSubType SignatureSubTypeDetailsModelSubTypeEnum `mandatory:"true" json:"modelSubType"`
}

func (m SignatureSubTypeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SignatureSubTypeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSignatureSubTypeDetailsModelTypeEnum(string(m.ModelType)); !ok && m.ModelType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelType: %s. Supported values are: %s.", m.ModelType, strings.Join(GetSignatureSubTypeDetailsModelTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSignatureSubTypeDetailsModelSubTypeEnum(string(m.ModelSubType)); !ok && m.ModelSubType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelSubType: %s. Supported values are: %s.", m.ModelSubType, strings.Join(GetSignatureSubTypeDetailsModelSubTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SignatureSubTypeDetailsModelTypeEnum Enum with underlying type: string
type SignatureSubTypeDetailsModelTypeEnum string

// Set of constants representing the allowable values for SignatureSubTypeDetailsModelTypeEnum
const (
	SignatureSubTypeDetailsModelTypeKeyValueExtraction         SignatureSubTypeDetailsModelTypeEnum = "PRE_TRAINED_KEY_VALUE_EXTRACTION"
	SignatureSubTypeDetailsModelTypeDocumentElementsExtraction SignatureSubTypeDetailsModelTypeEnum = "PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION"
)

var mappingSignatureSubTypeDetailsModelTypeEnum = map[string]SignatureSubTypeDetailsModelTypeEnum{
	"PRE_TRAINED_KEY_VALUE_EXTRACTION":         SignatureSubTypeDetailsModelTypeKeyValueExtraction,
	"PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION": SignatureSubTypeDetailsModelTypeDocumentElementsExtraction,
}

var mappingSignatureSubTypeDetailsModelTypeEnumLowerCase = map[string]SignatureSubTypeDetailsModelTypeEnum{
	"pre_trained_key_value_extraction":         SignatureSubTypeDetailsModelTypeKeyValueExtraction,
	"pre_trained_document_elements_extraction": SignatureSubTypeDetailsModelTypeDocumentElementsExtraction,
}

// GetSignatureSubTypeDetailsModelTypeEnumValues Enumerates the set of values for SignatureSubTypeDetailsModelTypeEnum
func GetSignatureSubTypeDetailsModelTypeEnumValues() []SignatureSubTypeDetailsModelTypeEnum {
	values := make([]SignatureSubTypeDetailsModelTypeEnum, 0)
	for _, v := range mappingSignatureSubTypeDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSignatureSubTypeDetailsModelTypeEnumStringValues Enumerates the set of values in String for SignatureSubTypeDetailsModelTypeEnum
func GetSignatureSubTypeDetailsModelTypeEnumStringValues() []string {
	return []string{
		"PRE_TRAINED_KEY_VALUE_EXTRACTION",
		"PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION",
	}
}

// GetMappingSignatureSubTypeDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSignatureSubTypeDetailsModelTypeEnum(val string) (SignatureSubTypeDetailsModelTypeEnum, bool) {
	enum, ok := mappingSignatureSubTypeDetailsModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SignatureSubTypeDetailsModelSubTypeEnum Enum with underlying type: string
type SignatureSubTypeDetailsModelSubTypeEnum string

// Set of constants representing the allowable values for SignatureSubTypeDetailsModelSubTypeEnum
const (
	SignatureSubTypeDetailsModelSubTypeQrBarCode SignatureSubTypeDetailsModelSubTypeEnum = "QR_BAR_CODE"
	SignatureSubTypeDetailsModelSubTypeSignature SignatureSubTypeDetailsModelSubTypeEnum = "SIGNATURE"
)

var mappingSignatureSubTypeDetailsModelSubTypeEnum = map[string]SignatureSubTypeDetailsModelSubTypeEnum{
	"QR_BAR_CODE": SignatureSubTypeDetailsModelSubTypeQrBarCode,
	"SIGNATURE":   SignatureSubTypeDetailsModelSubTypeSignature,
}

var mappingSignatureSubTypeDetailsModelSubTypeEnumLowerCase = map[string]SignatureSubTypeDetailsModelSubTypeEnum{
	"qr_bar_code": SignatureSubTypeDetailsModelSubTypeQrBarCode,
	"signature":   SignatureSubTypeDetailsModelSubTypeSignature,
}

// GetSignatureSubTypeDetailsModelSubTypeEnumValues Enumerates the set of values for SignatureSubTypeDetailsModelSubTypeEnum
func GetSignatureSubTypeDetailsModelSubTypeEnumValues() []SignatureSubTypeDetailsModelSubTypeEnum {
	values := make([]SignatureSubTypeDetailsModelSubTypeEnum, 0)
	for _, v := range mappingSignatureSubTypeDetailsModelSubTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSignatureSubTypeDetailsModelSubTypeEnumStringValues Enumerates the set of values in String for SignatureSubTypeDetailsModelSubTypeEnum
func GetSignatureSubTypeDetailsModelSubTypeEnumStringValues() []string {
	return []string{
		"QR_BAR_CODE",
		"SIGNATURE",
	}
}

// GetMappingSignatureSubTypeDetailsModelSubTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSignatureSubTypeDetailsModelSubTypeEnum(val string) (SignatureSubTypeDetailsModelSubTypeEnum, bool) {
	enum, ok := mappingSignatureSubTypeDetailsModelSubTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
