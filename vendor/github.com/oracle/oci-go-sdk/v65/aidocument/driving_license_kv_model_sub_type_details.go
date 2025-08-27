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

// DrivingLicenseKvModelSubTypeDetails Driving licence KV model sub type details
type DrivingLicenseKvModelSubTypeDetails struct {

	// Sub type model based on the model type.
	// The allowed values are:
	// - `PRE_TRAINED_KEY_VALUE_EXTRACTION`
	// - `PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION`
	ModelType DrivingLicenseKvModelSubTypeDetailsModelTypeEnum `mandatory:"true" json:"modelType"`

	// The model sub type for KEY_VALUE_EXTRACTION
	// The allowed values are:
	// - `RECEIPT`
	// - `INVOICE`
	// - `PASSPORT`
	// - `DRIVER_LICENSE`
	// - `HEALTH_INSURANCE_ID`
	ModelSubType DrivingLicenseKvModelSubTypeDetailsModelSubTypeEnum `mandatory:"true" json:"modelSubType"`
}

func (m DrivingLicenseKvModelSubTypeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrivingLicenseKvModelSubTypeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDrivingLicenseKvModelSubTypeDetailsModelTypeEnum(string(m.ModelType)); !ok && m.ModelType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelType: %s. Supported values are: %s.", m.ModelType, strings.Join(GetDrivingLicenseKvModelSubTypeDetailsModelTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDrivingLicenseKvModelSubTypeDetailsModelSubTypeEnum(string(m.ModelSubType)); !ok && m.ModelSubType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelSubType: %s. Supported values are: %s.", m.ModelSubType, strings.Join(GetDrivingLicenseKvModelSubTypeDetailsModelSubTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DrivingLicenseKvModelSubTypeDetailsModelTypeEnum Enum with underlying type: string
type DrivingLicenseKvModelSubTypeDetailsModelTypeEnum string

// Set of constants representing the allowable values for DrivingLicenseKvModelSubTypeDetailsModelTypeEnum
const (
	DrivingLicenseKvModelSubTypeDetailsModelTypeKeyValueExtraction         DrivingLicenseKvModelSubTypeDetailsModelTypeEnum = "PRE_TRAINED_KEY_VALUE_EXTRACTION"
	DrivingLicenseKvModelSubTypeDetailsModelTypeDocumentElementsExtraction DrivingLicenseKvModelSubTypeDetailsModelTypeEnum = "PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION"
)

var mappingDrivingLicenseKvModelSubTypeDetailsModelTypeEnum = map[string]DrivingLicenseKvModelSubTypeDetailsModelTypeEnum{
	"PRE_TRAINED_KEY_VALUE_EXTRACTION":         DrivingLicenseKvModelSubTypeDetailsModelTypeKeyValueExtraction,
	"PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION": DrivingLicenseKvModelSubTypeDetailsModelTypeDocumentElementsExtraction,
}

var mappingDrivingLicenseKvModelSubTypeDetailsModelTypeEnumLowerCase = map[string]DrivingLicenseKvModelSubTypeDetailsModelTypeEnum{
	"pre_trained_key_value_extraction":         DrivingLicenseKvModelSubTypeDetailsModelTypeKeyValueExtraction,
	"pre_trained_document_elements_extraction": DrivingLicenseKvModelSubTypeDetailsModelTypeDocumentElementsExtraction,
}

// GetDrivingLicenseKvModelSubTypeDetailsModelTypeEnumValues Enumerates the set of values for DrivingLicenseKvModelSubTypeDetailsModelTypeEnum
func GetDrivingLicenseKvModelSubTypeDetailsModelTypeEnumValues() []DrivingLicenseKvModelSubTypeDetailsModelTypeEnum {
	values := make([]DrivingLicenseKvModelSubTypeDetailsModelTypeEnum, 0)
	for _, v := range mappingDrivingLicenseKvModelSubTypeDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrivingLicenseKvModelSubTypeDetailsModelTypeEnumStringValues Enumerates the set of values in String for DrivingLicenseKvModelSubTypeDetailsModelTypeEnum
func GetDrivingLicenseKvModelSubTypeDetailsModelTypeEnumStringValues() []string {
	return []string{
		"PRE_TRAINED_KEY_VALUE_EXTRACTION",
		"PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION",
	}
}

// GetMappingDrivingLicenseKvModelSubTypeDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrivingLicenseKvModelSubTypeDetailsModelTypeEnum(val string) (DrivingLicenseKvModelSubTypeDetailsModelTypeEnum, bool) {
	enum, ok := mappingDrivingLicenseKvModelSubTypeDetailsModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DrivingLicenseKvModelSubTypeDetailsModelSubTypeEnum Enum with underlying type: string
type DrivingLicenseKvModelSubTypeDetailsModelSubTypeEnum string

// Set of constants representing the allowable values for DrivingLicenseKvModelSubTypeDetailsModelSubTypeEnum
const (
	DrivingLicenseKvModelSubTypeDetailsModelSubTypeReceipt           DrivingLicenseKvModelSubTypeDetailsModelSubTypeEnum = "RECEIPT"
	DrivingLicenseKvModelSubTypeDetailsModelSubTypeInvoice           DrivingLicenseKvModelSubTypeDetailsModelSubTypeEnum = "INVOICE"
	DrivingLicenseKvModelSubTypeDetailsModelSubTypePassport          DrivingLicenseKvModelSubTypeDetailsModelSubTypeEnum = "PASSPORT"
	DrivingLicenseKvModelSubTypeDetailsModelSubTypeDriverLicense     DrivingLicenseKvModelSubTypeDetailsModelSubTypeEnum = "DRIVER_LICENSE"
	DrivingLicenseKvModelSubTypeDetailsModelSubTypeHealthInsuranceId DrivingLicenseKvModelSubTypeDetailsModelSubTypeEnum = "HEALTH_INSURANCE_ID"
)

var mappingDrivingLicenseKvModelSubTypeDetailsModelSubTypeEnum = map[string]DrivingLicenseKvModelSubTypeDetailsModelSubTypeEnum{
	"RECEIPT":             DrivingLicenseKvModelSubTypeDetailsModelSubTypeReceipt,
	"INVOICE":             DrivingLicenseKvModelSubTypeDetailsModelSubTypeInvoice,
	"PASSPORT":            DrivingLicenseKvModelSubTypeDetailsModelSubTypePassport,
	"DRIVER_LICENSE":      DrivingLicenseKvModelSubTypeDetailsModelSubTypeDriverLicense,
	"HEALTH_INSURANCE_ID": DrivingLicenseKvModelSubTypeDetailsModelSubTypeHealthInsuranceId,
}

var mappingDrivingLicenseKvModelSubTypeDetailsModelSubTypeEnumLowerCase = map[string]DrivingLicenseKvModelSubTypeDetailsModelSubTypeEnum{
	"receipt":             DrivingLicenseKvModelSubTypeDetailsModelSubTypeReceipt,
	"invoice":             DrivingLicenseKvModelSubTypeDetailsModelSubTypeInvoice,
	"passport":            DrivingLicenseKvModelSubTypeDetailsModelSubTypePassport,
	"driver_license":      DrivingLicenseKvModelSubTypeDetailsModelSubTypeDriverLicense,
	"health_insurance_id": DrivingLicenseKvModelSubTypeDetailsModelSubTypeHealthInsuranceId,
}

// GetDrivingLicenseKvModelSubTypeDetailsModelSubTypeEnumValues Enumerates the set of values for DrivingLicenseKvModelSubTypeDetailsModelSubTypeEnum
func GetDrivingLicenseKvModelSubTypeDetailsModelSubTypeEnumValues() []DrivingLicenseKvModelSubTypeDetailsModelSubTypeEnum {
	values := make([]DrivingLicenseKvModelSubTypeDetailsModelSubTypeEnum, 0)
	for _, v := range mappingDrivingLicenseKvModelSubTypeDetailsModelSubTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrivingLicenseKvModelSubTypeDetailsModelSubTypeEnumStringValues Enumerates the set of values in String for DrivingLicenseKvModelSubTypeDetailsModelSubTypeEnum
func GetDrivingLicenseKvModelSubTypeDetailsModelSubTypeEnumStringValues() []string {
	return []string{
		"RECEIPT",
		"INVOICE",
		"PASSPORT",
		"DRIVER_LICENSE",
		"HEALTH_INSURANCE_ID",
	}
}

// GetMappingDrivingLicenseKvModelSubTypeDetailsModelSubTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrivingLicenseKvModelSubTypeDetailsModelSubTypeEnum(val string) (DrivingLicenseKvModelSubTypeDetailsModelSubTypeEnum, bool) {
	enum, ok := mappingDrivingLicenseKvModelSubTypeDetailsModelSubTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
