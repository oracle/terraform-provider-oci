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

// QrBarCodeSubTypeDetails Bar Code / QR Code Document Element Extraction model sub type details
type QrBarCodeSubTypeDetails struct {

	// Sub type model based on the model type.
	// The allowed values are:
	// - `PRE_TRAINED_KEY_VALUE_EXTRACTION`
	// - `PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION`
	ModelType QrBarCodeSubTypeDetailsModelTypeEnum `mandatory:"true" json:"modelType"`

	// The model sub type for DOCUMENT_ELEMENTS_EXTRACTION.
	// The allowed values are:
	//   - `QR_BAR_CODE`
	//   - `SIGNATURE`
	ModelSubType QrBarCodeSubTypeDetailsModelSubTypeEnum `mandatory:"true" json:"modelSubType"`
}

func (m QrBarCodeSubTypeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QrBarCodeSubTypeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingQrBarCodeSubTypeDetailsModelTypeEnum(string(m.ModelType)); !ok && m.ModelType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelType: %s. Supported values are: %s.", m.ModelType, strings.Join(GetQrBarCodeSubTypeDetailsModelTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingQrBarCodeSubTypeDetailsModelSubTypeEnum(string(m.ModelSubType)); !ok && m.ModelSubType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelSubType: %s. Supported values are: %s.", m.ModelSubType, strings.Join(GetQrBarCodeSubTypeDetailsModelSubTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// QrBarCodeSubTypeDetailsModelTypeEnum Enum with underlying type: string
type QrBarCodeSubTypeDetailsModelTypeEnum string

// Set of constants representing the allowable values for QrBarCodeSubTypeDetailsModelTypeEnum
const (
	QrBarCodeSubTypeDetailsModelTypeKeyValueExtraction         QrBarCodeSubTypeDetailsModelTypeEnum = "PRE_TRAINED_KEY_VALUE_EXTRACTION"
	QrBarCodeSubTypeDetailsModelTypeDocumentElementsExtraction QrBarCodeSubTypeDetailsModelTypeEnum = "PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION"
)

var mappingQrBarCodeSubTypeDetailsModelTypeEnum = map[string]QrBarCodeSubTypeDetailsModelTypeEnum{
	"PRE_TRAINED_KEY_VALUE_EXTRACTION":         QrBarCodeSubTypeDetailsModelTypeKeyValueExtraction,
	"PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION": QrBarCodeSubTypeDetailsModelTypeDocumentElementsExtraction,
}

var mappingQrBarCodeSubTypeDetailsModelTypeEnumLowerCase = map[string]QrBarCodeSubTypeDetailsModelTypeEnum{
	"pre_trained_key_value_extraction":         QrBarCodeSubTypeDetailsModelTypeKeyValueExtraction,
	"pre_trained_document_elements_extraction": QrBarCodeSubTypeDetailsModelTypeDocumentElementsExtraction,
}

// GetQrBarCodeSubTypeDetailsModelTypeEnumValues Enumerates the set of values for QrBarCodeSubTypeDetailsModelTypeEnum
func GetQrBarCodeSubTypeDetailsModelTypeEnumValues() []QrBarCodeSubTypeDetailsModelTypeEnum {
	values := make([]QrBarCodeSubTypeDetailsModelTypeEnum, 0)
	for _, v := range mappingQrBarCodeSubTypeDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetQrBarCodeSubTypeDetailsModelTypeEnumStringValues Enumerates the set of values in String for QrBarCodeSubTypeDetailsModelTypeEnum
func GetQrBarCodeSubTypeDetailsModelTypeEnumStringValues() []string {
	return []string{
		"PRE_TRAINED_KEY_VALUE_EXTRACTION",
		"PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION",
	}
}

// GetMappingQrBarCodeSubTypeDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQrBarCodeSubTypeDetailsModelTypeEnum(val string) (QrBarCodeSubTypeDetailsModelTypeEnum, bool) {
	enum, ok := mappingQrBarCodeSubTypeDetailsModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// QrBarCodeSubTypeDetailsModelSubTypeEnum Enum with underlying type: string
type QrBarCodeSubTypeDetailsModelSubTypeEnum string

// Set of constants representing the allowable values for QrBarCodeSubTypeDetailsModelSubTypeEnum
const (
	QrBarCodeSubTypeDetailsModelSubTypeQrBarCode QrBarCodeSubTypeDetailsModelSubTypeEnum = "QR_BAR_CODE"
	QrBarCodeSubTypeDetailsModelSubTypeSignature QrBarCodeSubTypeDetailsModelSubTypeEnum = "SIGNATURE"
)

var mappingQrBarCodeSubTypeDetailsModelSubTypeEnum = map[string]QrBarCodeSubTypeDetailsModelSubTypeEnum{
	"QR_BAR_CODE": QrBarCodeSubTypeDetailsModelSubTypeQrBarCode,
	"SIGNATURE":   QrBarCodeSubTypeDetailsModelSubTypeSignature,
}

var mappingQrBarCodeSubTypeDetailsModelSubTypeEnumLowerCase = map[string]QrBarCodeSubTypeDetailsModelSubTypeEnum{
	"qr_bar_code": QrBarCodeSubTypeDetailsModelSubTypeQrBarCode,
	"signature":   QrBarCodeSubTypeDetailsModelSubTypeSignature,
}

// GetQrBarCodeSubTypeDetailsModelSubTypeEnumValues Enumerates the set of values for QrBarCodeSubTypeDetailsModelSubTypeEnum
func GetQrBarCodeSubTypeDetailsModelSubTypeEnumValues() []QrBarCodeSubTypeDetailsModelSubTypeEnum {
	values := make([]QrBarCodeSubTypeDetailsModelSubTypeEnum, 0)
	for _, v := range mappingQrBarCodeSubTypeDetailsModelSubTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetQrBarCodeSubTypeDetailsModelSubTypeEnumStringValues Enumerates the set of values in String for QrBarCodeSubTypeDetailsModelSubTypeEnum
func GetQrBarCodeSubTypeDetailsModelSubTypeEnumStringValues() []string {
	return []string{
		"QR_BAR_CODE",
		"SIGNATURE",
	}
}

// GetMappingQrBarCodeSubTypeDetailsModelSubTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQrBarCodeSubTypeDetailsModelSubTypeEnum(val string) (QrBarCodeSubTypeDetailsModelSubTypeEnum, bool) {
	enum, ok := mappingQrBarCodeSubTypeDetailsModelSubTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
