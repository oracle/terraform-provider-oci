// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, policies, and identity domains.
//

package identity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EnumTagDefinitionValidator Used to validate the value set for a defined tag and contains the list of allowable `values`.
// You must specify at least one valid value in the `values` array. You can't have blank or
// or empty strings (`""`). Duplicate values are not allowed.
type EnumTagDefinitionValidator struct {

	// The list of allowed values for a definedTag value.
	Values []string `mandatory:"false" json:"values"`
}

func (m EnumTagDefinitionValidator) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EnumTagDefinitionValidator) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m EnumTagDefinitionValidator) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeEnumTagDefinitionValidator EnumTagDefinitionValidator
	s := struct {
		DiscriminatorParam string `json:"validatorType"`
		MarshalTypeEnumTagDefinitionValidator
	}{
		"ENUM",
		(MarshalTypeEnumTagDefinitionValidator)(m),
	}

	return json.Marshal(&s)
}
