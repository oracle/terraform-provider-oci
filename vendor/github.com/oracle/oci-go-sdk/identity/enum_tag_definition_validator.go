// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// EnumTagDefinitionValidator Validates the 'value' set for a definedTag is contained in the list of allowable 'values'.
// If the 'validatorType' is 'ENUM', then at least one valid value must be specified in the 'values' array.
type EnumTagDefinitionValidator struct {

	// The list of allowed values for a definedTag value.
	Values []string `mandatory:"false" json:"values"`
}

func (m EnumTagDefinitionValidator) String() string {
	return common.PointerString(m)
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
