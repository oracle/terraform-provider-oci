// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PreemptibleInstanceConfigDetails Configuration options for preemptible instances.
type PreemptibleInstanceConfigDetails struct {
	PreemptionAction PreemptionAction `mandatory:"true" json:"preemptionAction"`
}

func (m PreemptibleInstanceConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PreemptibleInstanceConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *PreemptibleInstanceConfigDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		PreemptionAction preemptionaction `json:"preemptionAction"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.PreemptionAction.UnmarshalPolymorphicJSON(model.PreemptionAction.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PreemptionAction = nn.(PreemptionAction)
	} else {
		m.PreemptionAction = nil
	}

	return
}
