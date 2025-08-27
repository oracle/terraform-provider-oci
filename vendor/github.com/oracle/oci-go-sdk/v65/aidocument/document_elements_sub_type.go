// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Document Understanding API
//
// Document AI helps customers perform various analysis on their documents. If a customer has lots of documents, they can process them in batch using asynchronous API endpoints.
//

package aidocument

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DocumentElementsSubType The Document Elements Extraction model sub type
type DocumentElementsSubType struct {

	// The model sub type for DOCUMENT_ELEMENTS_EXTRACTION.
	// The allowed values are:
	//   - `QR_BAR_CODE`
	//   - `SIGNATURE`
	ModelSubType DocumentElementsSubTypeModelSubTypeEnum `mandatory:"true" json:"modelSubType"`
}

func (m DocumentElementsSubType) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DocumentElementsSubType) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDocumentElementsSubTypeModelSubTypeEnum(string(m.ModelSubType)); !ok && m.ModelSubType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelSubType: %s. Supported values are: %s.", m.ModelSubType, strings.Join(GetDocumentElementsSubTypeModelSubTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DocumentElementsSubType) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDocumentElementsSubType DocumentElementsSubType
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDocumentElementsSubType
	}{
		"PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION",
		(MarshalTypeDocumentElementsSubType)(m),
	}

	return json.Marshal(&s)
}

// DocumentElementsSubTypeModelSubTypeEnum Enum with underlying type: string
type DocumentElementsSubTypeModelSubTypeEnum string

// Set of constants representing the allowable values for DocumentElementsSubTypeModelSubTypeEnum
const (
	DocumentElementsSubTypeModelSubTypeQrBarCode DocumentElementsSubTypeModelSubTypeEnum = "QR_BAR_CODE"
	DocumentElementsSubTypeModelSubTypeSignature DocumentElementsSubTypeModelSubTypeEnum = "SIGNATURE"
)

var mappingDocumentElementsSubTypeModelSubTypeEnum = map[string]DocumentElementsSubTypeModelSubTypeEnum{
	"QR_BAR_CODE": DocumentElementsSubTypeModelSubTypeQrBarCode,
	"SIGNATURE":   DocumentElementsSubTypeModelSubTypeSignature,
}

var mappingDocumentElementsSubTypeModelSubTypeEnumLowerCase = map[string]DocumentElementsSubTypeModelSubTypeEnum{
	"qr_bar_code": DocumentElementsSubTypeModelSubTypeQrBarCode,
	"signature":   DocumentElementsSubTypeModelSubTypeSignature,
}

// GetDocumentElementsSubTypeModelSubTypeEnumValues Enumerates the set of values for DocumentElementsSubTypeModelSubTypeEnum
func GetDocumentElementsSubTypeModelSubTypeEnumValues() []DocumentElementsSubTypeModelSubTypeEnum {
	values := make([]DocumentElementsSubTypeModelSubTypeEnum, 0)
	for _, v := range mappingDocumentElementsSubTypeModelSubTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDocumentElementsSubTypeModelSubTypeEnumStringValues Enumerates the set of values in String for DocumentElementsSubTypeModelSubTypeEnum
func GetDocumentElementsSubTypeModelSubTypeEnumStringValues() []string {
	return []string{
		"QR_BAR_CODE",
		"SIGNATURE",
	}
}

// GetMappingDocumentElementsSubTypeModelSubTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDocumentElementsSubTypeModelSubTypeEnum(val string) (DocumentElementsSubTypeModelSubTypeEnum, bool) {
	enum, ok := mappingDocumentElementsSubTypeModelSubTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
