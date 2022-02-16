// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// VisionService API
//
// A description of the VisionService API.
//

package aivision

import (
	"strings"
)

// DocumentTypeEnum Enum with underlying type: string
type DocumentTypeEnum string

// Set of constants representing the allowable values for DocumentTypeEnum
const (
	DocumentTypeInvoice       DocumentTypeEnum = "INVOICE"
	DocumentTypeReceipt       DocumentTypeEnum = "RECEIPT"
	DocumentTypeResume        DocumentTypeEnum = "RESUME"
	DocumentTypeTaxForm       DocumentTypeEnum = "TAX_FORM"
	DocumentTypeDriverLicense DocumentTypeEnum = "DRIVER_LICENSE"
	DocumentTypePassport      DocumentTypeEnum = "PASSPORT"
	DocumentTypeBankStatement DocumentTypeEnum = "BANK_STATEMENT"
	DocumentTypeCheck         DocumentTypeEnum = "CHECK"
	DocumentTypePayslip       DocumentTypeEnum = "PAYSLIP"
	DocumentTypeOthers        DocumentTypeEnum = "OTHERS"
)

var mappingDocumentTypeEnum = map[string]DocumentTypeEnum{
	"INVOICE":        DocumentTypeInvoice,
	"RECEIPT":        DocumentTypeReceipt,
	"RESUME":         DocumentTypeResume,
	"TAX_FORM":       DocumentTypeTaxForm,
	"DRIVER_LICENSE": DocumentTypeDriverLicense,
	"PASSPORT":       DocumentTypePassport,
	"BANK_STATEMENT": DocumentTypeBankStatement,
	"CHECK":          DocumentTypeCheck,
	"PAYSLIP":        DocumentTypePayslip,
	"OTHERS":         DocumentTypeOthers,
}

// GetDocumentTypeEnumValues Enumerates the set of values for DocumentTypeEnum
func GetDocumentTypeEnumValues() []DocumentTypeEnum {
	values := make([]DocumentTypeEnum, 0)
	for _, v := range mappingDocumentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDocumentTypeEnumStringValues Enumerates the set of values in String for DocumentTypeEnum
func GetDocumentTypeEnumStringValues() []string {
	return []string{
		"INVOICE",
		"RECEIPT",
		"RESUME",
		"TAX_FORM",
		"DRIVER_LICENSE",
		"PASSPORT",
		"BANK_STATEMENT",
		"CHECK",
		"PAYSLIP",
		"OTHERS",
	}
}

// GetMappingDocumentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDocumentTypeEnum(val string) (DocumentTypeEnum, bool) {
	mappingDocumentTypeEnumIgnoreCase := make(map[string]DocumentTypeEnum)
	for k, v := range mappingDocumentTypeEnum {
		mappingDocumentTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDocumentTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
