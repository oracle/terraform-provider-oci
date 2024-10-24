// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ProductStackSubCategoryDetails ProductStack Config Category Details.
type ProductStackSubCategoryDetails interface {
}

type productstacksubcategorydetails struct {
	JsonData    []byte
	SubCategory string `json:"subCategory"`
}

// UnmarshalJSON unmarshals json
func (m *productstacksubcategorydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerproductstacksubcategorydetails productstacksubcategorydetails
	s := struct {
		Model Unmarshalerproductstacksubcategorydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SubCategory = s.Model.SubCategory

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *productstacksubcategorydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SubCategory {
	case "PRODUCT_STACK_AS_PRODUCT":
		mm := ProductStackAsProductSubCategoryDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRODUCT_STACK_GENERIC":
		mm := ProductStackGenericSubCategoryDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ProductStackSubCategoryDetails: %s.", m.SubCategory)
		return *m, nil
	}
}

func (m productstacksubcategorydetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m productstacksubcategorydetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ProductStackSubCategoryDetailsSubCategoryEnum Enum with underlying type: string
type ProductStackSubCategoryDetailsSubCategoryEnum string

// Set of constants representing the allowable values for ProductStackSubCategoryDetailsSubCategoryEnum
const (
	ProductStackSubCategoryDetailsSubCategoryGeneric   ProductStackSubCategoryDetailsSubCategoryEnum = "PRODUCT_STACK_GENERIC"
	ProductStackSubCategoryDetailsSubCategoryAsProduct ProductStackSubCategoryDetailsSubCategoryEnum = "PRODUCT_STACK_AS_PRODUCT"
)

var mappingProductStackSubCategoryDetailsSubCategoryEnum = map[string]ProductStackSubCategoryDetailsSubCategoryEnum{
	"PRODUCT_STACK_GENERIC":    ProductStackSubCategoryDetailsSubCategoryGeneric,
	"PRODUCT_STACK_AS_PRODUCT": ProductStackSubCategoryDetailsSubCategoryAsProduct,
}

var mappingProductStackSubCategoryDetailsSubCategoryEnumLowerCase = map[string]ProductStackSubCategoryDetailsSubCategoryEnum{
	"product_stack_generic":    ProductStackSubCategoryDetailsSubCategoryGeneric,
	"product_stack_as_product": ProductStackSubCategoryDetailsSubCategoryAsProduct,
}

// GetProductStackSubCategoryDetailsSubCategoryEnumValues Enumerates the set of values for ProductStackSubCategoryDetailsSubCategoryEnum
func GetProductStackSubCategoryDetailsSubCategoryEnumValues() []ProductStackSubCategoryDetailsSubCategoryEnum {
	values := make([]ProductStackSubCategoryDetailsSubCategoryEnum, 0)
	for _, v := range mappingProductStackSubCategoryDetailsSubCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetProductStackSubCategoryDetailsSubCategoryEnumStringValues Enumerates the set of values in String for ProductStackSubCategoryDetailsSubCategoryEnum
func GetProductStackSubCategoryDetailsSubCategoryEnumStringValues() []string {
	return []string{
		"PRODUCT_STACK_GENERIC",
		"PRODUCT_STACK_AS_PRODUCT",
	}
}

// GetMappingProductStackSubCategoryDetailsSubCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProductStackSubCategoryDetailsSubCategoryEnum(val string) (ProductStackSubCategoryDetailsSubCategoryEnum, bool) {
	enum, ok := mappingProductStackSubCategoryDetailsSubCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
