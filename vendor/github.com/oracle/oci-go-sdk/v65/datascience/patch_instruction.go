// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchInstruction A single instruction to be included as part of Patch request content.
type PatchInstruction interface {
}

type patchinstruction struct {
	JsonData  []byte
	Operation string `json:"operation"`
}

// UnmarshalJSON unmarshals json
func (m *patchinstruction) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpatchinstruction patchinstruction
	s := struct {
		Model Unmarshalerpatchinstruction
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Operation = s.Model.Operation

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *patchinstruction) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Operation {
	case "REMOVE":
		mm := PatchRemoveMemberModels{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INSERT":
		mm := PatchInsertNewMemberModels{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for PatchInstruction: %s.", m.Operation)
		return *m, nil
	}
}

func (m patchinstruction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m patchinstruction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
