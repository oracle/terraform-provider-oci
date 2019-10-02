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

// DefaultTagDefinitionValidator This is the default validatorType for definedTag. This is same as not setting any value on the validator field.
// By default only string value can be set for this definedTag.
type DefaultTagDefinitionValidator struct {
}

func (m DefaultTagDefinitionValidator) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DefaultTagDefinitionValidator) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDefaultTagDefinitionValidator DefaultTagDefinitionValidator
	s := struct {
		DiscriminatorParam string `json:"validatorType"`
		MarshalTypeDefaultTagDefinitionValidator
	}{
		"DEFAULT",
		(MarshalTypeDefaultTagDefinitionValidator)(m),
	}

	return json.Marshal(&s)
}
