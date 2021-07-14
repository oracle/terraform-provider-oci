// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v44/common"
)

// DeployEnvironmentCollection Results of a deployment environment search.
type DeployEnvironmentCollection struct {

	// Deployment environment summary items found for the search.
	Items []DeployEnvironmentSummary `mandatory:"true" json:"items"`
}

func (m DeployEnvironmentCollection) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *DeployEnvironmentCollection) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Items []deployenvironmentsummary `json:"items"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Items = make([]DeployEnvironmentSummary, len(model.Items))
	for i, n := range model.Items {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Items[i] = nn.(DeployEnvironmentSummary)
		} else {
			m.Items[i] = nil
		}
	}

	return
}
