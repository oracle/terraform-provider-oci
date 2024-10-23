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

// ProductStackAsProductSubCategoryDetails Consider Product stack as product.To be provided if the product stack should be considered as a Product also.
// Allows compliance reporting and patching against the product stack target.
type ProductStackAsProductSubCategoryDetails struct {

	// Versions associated with the PRODUCT .
	Versions []string `mandatory:"true" json:"versions"`

	// OCID for the Credential name to be associated with the Product Stack.
	// These are useful for target discovery or lifecycle management activities, for example, Oracle WebLogic admin credentials for Oracle WebLogic Application server.
	Credentials []ConfigAssociationDetails `mandatory:"false" json:"credentials"`

	// Various components of the Product.
	// For example:The administration server or node manager can be the components of the Oracle WebLogic Application server.
	// Forms server or concurrent manager can be the components of the Oracle E-Business Suite.
	Components []string `mandatory:"false" json:"components"`

	// Patch Types associated with this Product Stack which will be considered as Product.
	PatchTypes []ConfigAssociationDetails `mandatory:"false" json:"patchTypes"`
}

func (m ProductStackAsProductSubCategoryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProductStackAsProductSubCategoryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ProductStackAsProductSubCategoryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeProductStackAsProductSubCategoryDetails ProductStackAsProductSubCategoryDetails
	s := struct {
		DiscriminatorParam string `json:"subCategory"`
		MarshalTypeProductStackAsProductSubCategoryDetails
	}{
		"PRODUCT_STACK_AS_PRODUCT",
		(MarshalTypeProductStackAsProductSubCategoryDetails)(m),
	}

	return json.Marshal(&s)
}
