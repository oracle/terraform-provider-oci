// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OciControlCenterCp API
//
// A description of the OciControlCenterCp API
//

package capacitymanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchOccCapacityRequestDetails CapacityRequest details to patch.
type PatchOccCapacityRequestDetails struct {

	// List of patch instructions.
	Items []PatchInstruction `mandatory:"false" json:"items"`
}

func (m PatchOccCapacityRequestDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchOccCapacityRequestDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *PatchOccCapacityRequestDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Items []patchinstruction `json:"items"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Items = make([]PatchInstruction, len(model.Items))
	for i, n := range model.Items {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Items[i] = nn.(PatchInstruction)
		} else {
			m.Items[i] = nil
		}
	}
	return
}
