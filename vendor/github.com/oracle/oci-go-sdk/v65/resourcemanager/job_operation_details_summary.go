// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JobOperationDetailsSummary A summary of job details that is specific to the operation type.
type JobOperationDetailsSummary interface {
}

type joboperationdetailssummary struct {
	JsonData  []byte
	Operation string `json:"operation"`
}

// UnmarshalJSON unmarshals json
func (m *joboperationdetailssummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerjoboperationdetailssummary joboperationdetailssummary
	s := struct {
		Model Unmarshalerjoboperationdetailssummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Operation = s.Model.Operation

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *joboperationdetailssummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Operation {
	case "APPLY_ROLLBACK":
		mm := ApplyRollbackJobOperationDetailsSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IMPORT_TF_STATE":
		mm := ImportTfStateJobOperationDetailsSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PLAN_ROLLBACK":
		mm := PlanRollbackJobOperationDetailsSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PLAN":
		mm := PlanJobOperationDetailsSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DESTROY":
		mm := DestroyJobOperationDetailsSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APPLY":
		mm := ApplyJobOperationDetailsSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for JobOperationDetailsSummary: %s.", m.Operation)
		return *m, nil
	}
}

func (m joboperationdetailssummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m joboperationdetailssummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
