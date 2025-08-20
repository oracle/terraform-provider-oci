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

// KvModelSubType The Kv model sub type
type KvModelSubType struct {

	// The model sub type for KEY_VALUE_EXTRACTION
	// The allowed values are:
	// - `RECEIPT`
	// - `INVOICE`
	// - `PASSPORT`
	// - `DRIVER_LICENSE`
	// - `HEALTH_INSURANCE_ID`
	ModelSubType KvModelSubTypeModelSubTypeEnum `mandatory:"true" json:"modelSubType"`
}

func (m KvModelSubType) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KvModelSubType) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingKvModelSubTypeModelSubTypeEnum(string(m.ModelSubType)); !ok && m.ModelSubType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelSubType: %s. Supported values are: %s.", m.ModelSubType, strings.Join(GetKvModelSubTypeModelSubTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m KvModelSubType) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeKvModelSubType KvModelSubType
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeKvModelSubType
	}{
		"PRE_TRAINED_KEY_VALUE_EXTRACTION",
		(MarshalTypeKvModelSubType)(m),
	}

	return json.Marshal(&s)
}

// KvModelSubTypeModelSubTypeEnum Enum with underlying type: string
type KvModelSubTypeModelSubTypeEnum string

// Set of constants representing the allowable values for KvModelSubTypeModelSubTypeEnum
const (
	KvModelSubTypeModelSubTypeReceipt           KvModelSubTypeModelSubTypeEnum = "RECEIPT"
	KvModelSubTypeModelSubTypeInvoice           KvModelSubTypeModelSubTypeEnum = "INVOICE"
	KvModelSubTypeModelSubTypePassport          KvModelSubTypeModelSubTypeEnum = "PASSPORT"
	KvModelSubTypeModelSubTypeDriverLicense     KvModelSubTypeModelSubTypeEnum = "DRIVER_LICENSE"
	KvModelSubTypeModelSubTypeHealthInsuranceId KvModelSubTypeModelSubTypeEnum = "HEALTH_INSURANCE_ID"
)

var mappingKvModelSubTypeModelSubTypeEnum = map[string]KvModelSubTypeModelSubTypeEnum{
	"RECEIPT":             KvModelSubTypeModelSubTypeReceipt,
	"INVOICE":             KvModelSubTypeModelSubTypeInvoice,
	"PASSPORT":            KvModelSubTypeModelSubTypePassport,
	"DRIVER_LICENSE":      KvModelSubTypeModelSubTypeDriverLicense,
	"HEALTH_INSURANCE_ID": KvModelSubTypeModelSubTypeHealthInsuranceId,
}

var mappingKvModelSubTypeModelSubTypeEnumLowerCase = map[string]KvModelSubTypeModelSubTypeEnum{
	"receipt":             KvModelSubTypeModelSubTypeReceipt,
	"invoice":             KvModelSubTypeModelSubTypeInvoice,
	"passport":            KvModelSubTypeModelSubTypePassport,
	"driver_license":      KvModelSubTypeModelSubTypeDriverLicense,
	"health_insurance_id": KvModelSubTypeModelSubTypeHealthInsuranceId,
}

// GetKvModelSubTypeModelSubTypeEnumValues Enumerates the set of values for KvModelSubTypeModelSubTypeEnum
func GetKvModelSubTypeModelSubTypeEnumValues() []KvModelSubTypeModelSubTypeEnum {
	values := make([]KvModelSubTypeModelSubTypeEnum, 0)
	for _, v := range mappingKvModelSubTypeModelSubTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetKvModelSubTypeModelSubTypeEnumStringValues Enumerates the set of values in String for KvModelSubTypeModelSubTypeEnum
func GetKvModelSubTypeModelSubTypeEnumStringValues() []string {
	return []string{
		"RECEIPT",
		"INVOICE",
		"PASSPORT",
		"DRIVER_LICENSE",
		"HEALTH_INSURANCE_ID",
	}
}

// GetMappingKvModelSubTypeModelSubTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKvModelSubTypeModelSubTypeEnum(val string) (KvModelSubTypeModelSubTypeEnum, bool) {
	enum, ok := mappingKvModelSubTypeModelSubTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
