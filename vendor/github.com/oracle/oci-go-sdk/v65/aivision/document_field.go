// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vision API
//
// Using Vision, you can upload images to detect and classify objects in them. If you have lots of images, you can process them in batch using asynchronous API endpoints. Vision's features are thematically split between Document AI for document-centric images, and Image Analysis for object and scene-based images. Pretrained models and custom models are supported.
//

package aivision

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DocumentField Form field.
type DocumentField struct {

	// The field type.
	FieldType DocumentFieldFieldTypeEnum `mandatory:"true" json:"fieldType"`

	FieldValue FieldValue `mandatory:"true" json:"fieldValue"`

	FieldLabel *FieldLabel `mandatory:"false" json:"fieldLabel"`

	FieldName *FieldName `mandatory:"false" json:"fieldName"`
}

func (m DocumentField) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DocumentField) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDocumentFieldFieldTypeEnum(string(m.FieldType)); !ok && m.FieldType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FieldType: %s. Supported values are: %s.", m.FieldType, strings.Join(GetDocumentFieldFieldTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DocumentField) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FieldLabel *FieldLabel                `json:"fieldLabel"`
		FieldName  *FieldName                 `json:"fieldName"`
		FieldType  DocumentFieldFieldTypeEnum `json:"fieldType"`
		FieldValue fieldvalue                 `json:"fieldValue"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.FieldLabel = model.FieldLabel

	m.FieldName = model.FieldName

	m.FieldType = model.FieldType

	nn, e = model.FieldValue.UnmarshalPolymorphicJSON(model.FieldValue.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.FieldValue = nn.(FieldValue)
	} else {
		m.FieldValue = nil
	}

	return
}

// DocumentFieldFieldTypeEnum Enum with underlying type: string
type DocumentFieldFieldTypeEnum string

// Set of constants representing the allowable values for DocumentFieldFieldTypeEnum
const (
	DocumentFieldFieldTypeLineItemGroup DocumentFieldFieldTypeEnum = "LINE_ITEM_GROUP"
	DocumentFieldFieldTypeLineItem      DocumentFieldFieldTypeEnum = "LINE_ITEM"
	DocumentFieldFieldTypeLineItemField DocumentFieldFieldTypeEnum = "LINE_ITEM_FIELD"
	DocumentFieldFieldTypeKeyValue      DocumentFieldFieldTypeEnum = "KEY_VALUE"
)

var mappingDocumentFieldFieldTypeEnum = map[string]DocumentFieldFieldTypeEnum{
	"LINE_ITEM_GROUP": DocumentFieldFieldTypeLineItemGroup,
	"LINE_ITEM":       DocumentFieldFieldTypeLineItem,
	"LINE_ITEM_FIELD": DocumentFieldFieldTypeLineItemField,
	"KEY_VALUE":       DocumentFieldFieldTypeKeyValue,
}

var mappingDocumentFieldFieldTypeEnumLowerCase = map[string]DocumentFieldFieldTypeEnum{
	"line_item_group": DocumentFieldFieldTypeLineItemGroup,
	"line_item":       DocumentFieldFieldTypeLineItem,
	"line_item_field": DocumentFieldFieldTypeLineItemField,
	"key_value":       DocumentFieldFieldTypeKeyValue,
}

// GetDocumentFieldFieldTypeEnumValues Enumerates the set of values for DocumentFieldFieldTypeEnum
func GetDocumentFieldFieldTypeEnumValues() []DocumentFieldFieldTypeEnum {
	values := make([]DocumentFieldFieldTypeEnum, 0)
	for _, v := range mappingDocumentFieldFieldTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDocumentFieldFieldTypeEnumStringValues Enumerates the set of values in String for DocumentFieldFieldTypeEnum
func GetDocumentFieldFieldTypeEnumStringValues() []string {
	return []string{
		"LINE_ITEM_GROUP",
		"LINE_ITEM",
		"LINE_ITEM_FIELD",
		"KEY_VALUE",
	}
}

// GetMappingDocumentFieldFieldTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDocumentFieldFieldTypeEnum(val string) (DocumentFieldFieldTypeEnum, bool) {
	enum, ok := mappingDocumentFieldFieldTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
