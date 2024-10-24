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

// ProductConfigCategoryDetails Product Config Category Details.
// Defines individual products which contribute to the applications hosting on the resources that are to be managed.
type ProductConfigCategoryDetails struct {

	// Versions associated with the PRODUCT .
	Versions []string `mandatory:"true" json:"versions"`

	// OCID for the Credential name to be associated with the Product.
	// These are useful for target discovery or lifecycle management activities, for example, Oracle WebLogic admin credentials for Oracle WebLogic Application server.
	Credentials []ConfigAssociationDetails `mandatory:"false" json:"credentials"`

	// Various components of the Product.
	// For example:The administration server or node manager can be the components of the Oracle WebLogic Application server.
	// Forms server or concurrent manager can be the components of the Oracle E-Business Suite.
	Components []string `mandatory:"false" json:"components"`

	// Products compatible with this Product.
	// Provide products from the list of other products you have created that are compatible with the present one
	CompatibleProducts []ConfigAssociationDetails `mandatory:"false" json:"compatibleProducts"`

	// Patch Types associated with this Product.
	PatchTypes []ConfigAssociationDetails `mandatory:"false" json:"patchTypes"`
}

func (m ProductConfigCategoryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProductConfigCategoryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ProductConfigCategoryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeProductConfigCategoryDetails ProductConfigCategoryDetails
	s := struct {
		DiscriminatorParam string `json:"configCategory"`
		MarshalTypeProductConfigCategoryDetails
	}{
		"PRODUCT",
		(MarshalTypeProductConfigCategoryDetails)(m),
	}

	return json.Marshal(&s)
}
