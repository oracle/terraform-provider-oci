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

// InvoiceKvModelSubTypeDetails Invoice KV model sub type details
type InvoiceKvModelSubTypeDetails struct {

	// Sub type model based on the model type.
	// The allowed values are:
	// - `PRE_TRAINED_KEY_VALUE_EXTRACTION`
	// - `PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION`
	ModelType InvoiceKvModelSubTypeDetailsModelTypeEnum `mandatory:"true" json:"modelType"`

	// The model sub type for KEY_VALUE_EXTRACTION
	// The allowed values are:
	// - `RECEIPT`
	// - `INVOICE`
	// - `PASSPORT`
	// - `DRIVER_LICENSE`
	// - `HEALTH_INSURANCE_ID`
	ModelSubType InvoiceKvModelSubTypeDetailsModelSubTypeEnum `mandatory:"true" json:"modelSubType"`
}

func (m InvoiceKvModelSubTypeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InvoiceKvModelSubTypeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInvoiceKvModelSubTypeDetailsModelTypeEnum(string(m.ModelType)); !ok && m.ModelType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelType: %s. Supported values are: %s.", m.ModelType, strings.Join(GetInvoiceKvModelSubTypeDetailsModelTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingInvoiceKvModelSubTypeDetailsModelSubTypeEnum(string(m.ModelSubType)); !ok && m.ModelSubType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelSubType: %s. Supported values are: %s.", m.ModelSubType, strings.Join(GetInvoiceKvModelSubTypeDetailsModelSubTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InvoiceKvModelSubTypeDetailsModelTypeEnum Enum with underlying type: string
type InvoiceKvModelSubTypeDetailsModelTypeEnum string

// Set of constants representing the allowable values for InvoiceKvModelSubTypeDetailsModelTypeEnum
const (
	InvoiceKvModelSubTypeDetailsModelTypeKeyValueExtraction         InvoiceKvModelSubTypeDetailsModelTypeEnum = "PRE_TRAINED_KEY_VALUE_EXTRACTION"
	InvoiceKvModelSubTypeDetailsModelTypeDocumentElementsExtraction InvoiceKvModelSubTypeDetailsModelTypeEnum = "PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION"
)

var mappingInvoiceKvModelSubTypeDetailsModelTypeEnum = map[string]InvoiceKvModelSubTypeDetailsModelTypeEnum{
	"PRE_TRAINED_KEY_VALUE_EXTRACTION":         InvoiceKvModelSubTypeDetailsModelTypeKeyValueExtraction,
	"PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION": InvoiceKvModelSubTypeDetailsModelTypeDocumentElementsExtraction,
}

var mappingInvoiceKvModelSubTypeDetailsModelTypeEnumLowerCase = map[string]InvoiceKvModelSubTypeDetailsModelTypeEnum{
	"pre_trained_key_value_extraction":         InvoiceKvModelSubTypeDetailsModelTypeKeyValueExtraction,
	"pre_trained_document_elements_extraction": InvoiceKvModelSubTypeDetailsModelTypeDocumentElementsExtraction,
}

// GetInvoiceKvModelSubTypeDetailsModelTypeEnumValues Enumerates the set of values for InvoiceKvModelSubTypeDetailsModelTypeEnum
func GetInvoiceKvModelSubTypeDetailsModelTypeEnumValues() []InvoiceKvModelSubTypeDetailsModelTypeEnum {
	values := make([]InvoiceKvModelSubTypeDetailsModelTypeEnum, 0)
	for _, v := range mappingInvoiceKvModelSubTypeDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInvoiceKvModelSubTypeDetailsModelTypeEnumStringValues Enumerates the set of values in String for InvoiceKvModelSubTypeDetailsModelTypeEnum
func GetInvoiceKvModelSubTypeDetailsModelTypeEnumStringValues() []string {
	return []string{
		"PRE_TRAINED_KEY_VALUE_EXTRACTION",
		"PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION",
	}
}

// GetMappingInvoiceKvModelSubTypeDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInvoiceKvModelSubTypeDetailsModelTypeEnum(val string) (InvoiceKvModelSubTypeDetailsModelTypeEnum, bool) {
	enum, ok := mappingInvoiceKvModelSubTypeDetailsModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// InvoiceKvModelSubTypeDetailsModelSubTypeEnum Enum with underlying type: string
type InvoiceKvModelSubTypeDetailsModelSubTypeEnum string

// Set of constants representing the allowable values for InvoiceKvModelSubTypeDetailsModelSubTypeEnum
const (
	InvoiceKvModelSubTypeDetailsModelSubTypeReceipt           InvoiceKvModelSubTypeDetailsModelSubTypeEnum = "RECEIPT"
	InvoiceKvModelSubTypeDetailsModelSubTypeInvoice           InvoiceKvModelSubTypeDetailsModelSubTypeEnum = "INVOICE"
	InvoiceKvModelSubTypeDetailsModelSubTypePassport          InvoiceKvModelSubTypeDetailsModelSubTypeEnum = "PASSPORT"
	InvoiceKvModelSubTypeDetailsModelSubTypeDriverLicense     InvoiceKvModelSubTypeDetailsModelSubTypeEnum = "DRIVER_LICENSE"
	InvoiceKvModelSubTypeDetailsModelSubTypeHealthInsuranceId InvoiceKvModelSubTypeDetailsModelSubTypeEnum = "HEALTH_INSURANCE_ID"
)

var mappingInvoiceKvModelSubTypeDetailsModelSubTypeEnum = map[string]InvoiceKvModelSubTypeDetailsModelSubTypeEnum{
	"RECEIPT":             InvoiceKvModelSubTypeDetailsModelSubTypeReceipt,
	"INVOICE":             InvoiceKvModelSubTypeDetailsModelSubTypeInvoice,
	"PASSPORT":            InvoiceKvModelSubTypeDetailsModelSubTypePassport,
	"DRIVER_LICENSE":      InvoiceKvModelSubTypeDetailsModelSubTypeDriverLicense,
	"HEALTH_INSURANCE_ID": InvoiceKvModelSubTypeDetailsModelSubTypeHealthInsuranceId,
}

var mappingInvoiceKvModelSubTypeDetailsModelSubTypeEnumLowerCase = map[string]InvoiceKvModelSubTypeDetailsModelSubTypeEnum{
	"receipt":             InvoiceKvModelSubTypeDetailsModelSubTypeReceipt,
	"invoice":             InvoiceKvModelSubTypeDetailsModelSubTypeInvoice,
	"passport":            InvoiceKvModelSubTypeDetailsModelSubTypePassport,
	"driver_license":      InvoiceKvModelSubTypeDetailsModelSubTypeDriverLicense,
	"health_insurance_id": InvoiceKvModelSubTypeDetailsModelSubTypeHealthInsuranceId,
}

// GetInvoiceKvModelSubTypeDetailsModelSubTypeEnumValues Enumerates the set of values for InvoiceKvModelSubTypeDetailsModelSubTypeEnum
func GetInvoiceKvModelSubTypeDetailsModelSubTypeEnumValues() []InvoiceKvModelSubTypeDetailsModelSubTypeEnum {
	values := make([]InvoiceKvModelSubTypeDetailsModelSubTypeEnum, 0)
	for _, v := range mappingInvoiceKvModelSubTypeDetailsModelSubTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInvoiceKvModelSubTypeDetailsModelSubTypeEnumStringValues Enumerates the set of values in String for InvoiceKvModelSubTypeDetailsModelSubTypeEnum
func GetInvoiceKvModelSubTypeDetailsModelSubTypeEnumStringValues() []string {
	return []string{
		"RECEIPT",
		"INVOICE",
		"PASSPORT",
		"DRIVER_LICENSE",
		"HEALTH_INSURANCE_ID",
	}
}

// GetMappingInvoiceKvModelSubTypeDetailsModelSubTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInvoiceKvModelSubTypeDetailsModelSubTypeEnum(val string) (InvoiceKvModelSubTypeDetailsModelSubTypeEnum, bool) {
	enum, ok := mappingInvoiceKvModelSubTypeDetailsModelSubTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
