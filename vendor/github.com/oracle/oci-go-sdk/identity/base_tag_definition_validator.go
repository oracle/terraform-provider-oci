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

// BaseTagDefinitionValidator Validates a definedTag value. Each validator performs validation steps in addition to the standard validation
// for definedTag values (See Limits on Tags (https://docs.cloud.oracle.com/Content/Identity/Concepts/taggingoverview.htm#Limits).
// If a validator is defined after a value has been set for a definedTag, then any UPDATE operation that attempts
// to change the value must pass the additional validation defined by this rule. Previously set values, that would
// fail validation, are not updated and it is possible to update other attributes of an OCI resource that contains
// a non-valid definedTag.
// To clear the validator call the UPDATE operation with DefaultTagDefinitionValidator.
type BaseTagDefinitionValidator interface {
}

type basetagdefinitionvalidator struct {
	JsonData      []byte
	ValidatorType string `json:"validatorType"`
}

// UnmarshalJSON unmarshals json
func (m *basetagdefinitionvalidator) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbasetagdefinitionvalidator basetagdefinitionvalidator
	s := struct {
		Model Unmarshalerbasetagdefinitionvalidator
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValidatorType = s.Model.ValidatorType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *basetagdefinitionvalidator) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValidatorType {
	case "DEFAULT":
		mm := DefaultTagDefinitionValidator{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ENUM":
		mm := EnumTagDefinitionValidator{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m basetagdefinitionvalidator) String() string {
	return common.PointerString(m)
}
