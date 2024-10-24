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

// ProductStackConfigCategoryDetails ProductStack Config Category Details.
// Defines suite or stack of products on which applications hosted by the resources are built and need to be managed.
type ProductStackConfigCategoryDetails struct {

	// Products that belong to the stack.
	// For example, Oracle WebLogic and Java for the Oracle Fusion Middleware product stack.
	Products []ConfigAssociationDetails `mandatory:"true" json:"products"`

	SubCategoryDetails ProductStackSubCategoryDetails `mandatory:"false" json:"subCategoryDetails"`
}

func (m ProductStackConfigCategoryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProductStackConfigCategoryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ProductStackConfigCategoryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeProductStackConfigCategoryDetails ProductStackConfigCategoryDetails
	s := struct {
		DiscriminatorParam string `json:"configCategory"`
		MarshalTypeProductStackConfigCategoryDetails
	}{
		"PRODUCT_STACK",
		(MarshalTypeProductStackConfigCategoryDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ProductStackConfigCategoryDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		SubCategoryDetails productstacksubcategorydetails `json:"subCategoryDetails"`
		Products           []ConfigAssociationDetails     `json:"products"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.SubCategoryDetails.UnmarshalPolymorphicJSON(model.SubCategoryDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SubCategoryDetails = nn.(ProductStackSubCategoryDetails)
	} else {
		m.SubCategoryDetails = nil
	}

	m.Products = make([]ConfigAssociationDetails, len(model.Products))
	copy(m.Products, model.Products)
	return
}
