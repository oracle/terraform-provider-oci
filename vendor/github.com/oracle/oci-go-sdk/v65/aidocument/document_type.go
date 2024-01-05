// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Document Understanding API
//
// Document AI helps customers perform various analysis on their documents. If a customer has lots of documents, they can process them in batch using asynchronous API endpoints.
//

package aidocument

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

var mappingDocumentTypeEnumLowerCase = map[string]DocumentTypeEnum{
	"invoice":        DocumentTypeInvoice,
	"receipt":        DocumentTypeReceipt,
	"resume":         DocumentTypeResume,
	"tax_form":       DocumentTypeTaxForm,
	"driver_license": DocumentTypeDriverLicense,
	"passport":       DocumentTypePassport,
	"bank_statement": DocumentTypeBankStatement,
	"check":          DocumentTypeCheck,
	"payslip":        DocumentTypePayslip,
	"others":         DocumentTypeOthers,
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
	enum, ok := mappingDocumentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
