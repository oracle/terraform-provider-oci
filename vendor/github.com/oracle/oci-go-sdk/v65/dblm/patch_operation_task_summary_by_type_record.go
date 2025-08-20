// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchOperationTaskSummaryByTypeRecord Patch operation task summary by type record.
type PatchOperationTaskSummaryByTypeRecord struct {

	// Patch task type.
	Type PatchTaskTypeEnum `mandatory:"false" json:"type,omitempty"`

	// Summary of patch tasks of the associated task type grouped by status.
	TaskSummaryByStatus []PatchOperationTaskSummaryByStatusRecord `mandatory:"false" json:"taskSummaryByStatus"`
}

func (m PatchOperationTaskSummaryByTypeRecord) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchOperationTaskSummaryByTypeRecord) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPatchTaskTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetPatchTaskTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
