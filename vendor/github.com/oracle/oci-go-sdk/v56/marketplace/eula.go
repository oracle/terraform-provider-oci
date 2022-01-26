// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Eula A base object for all types of end user license agreements.
type Eula interface {
}

type eula struct {
	JsonData []byte
	EulaType string `json:"eulaType"`
}

// UnmarshalJSON unmarshals json
func (m *eula) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalereula eula
	s := struct {
		Model Unmarshalereula
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.EulaType = s.Model.EulaType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *eula) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EulaType {
	case "TEXT":
		mm := TextBasedEula{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m eula) String() string {
	return common.PointerString(m)
}
