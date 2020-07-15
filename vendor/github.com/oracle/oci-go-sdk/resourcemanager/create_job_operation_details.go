// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// API for the Resource Manager service.
// Use this API to install, configure, and manage resources via the "infrastructure-as-code" model.
// For more information, see
// Overview of Resource Manager (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm).
//

package resourcemanager

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateJobOperationDetails Job details that are specific to the operation type.
type CreateJobOperationDetails interface {
}

type createjoboperationdetails struct {
	JsonData  []byte
	Operation string `json:"operation"`
}

// UnmarshalJSON unmarshals json
func (m *createjoboperationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatejoboperationdetails createjoboperationdetails
	s := struct {
		Model Unmarshalercreatejoboperationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Operation = s.Model.Operation

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createjoboperationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Operation {
	case "IMPORT_TF_STATE":
		mm := CreateImportTfStateJobOperationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APPLY":
		mm := CreateApplyJobOperationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PLAN":
		mm := CreatePlanJobOperationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DESTROY":
		mm := CreateDestroyJobOperationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m createjoboperationdetails) String() string {
	return common.PointerString(m)
}
