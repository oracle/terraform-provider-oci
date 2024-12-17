// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PipelineSchemaTableCollection List of source or target schema tables of a pipeline's assigned connection.
type PipelineSchemaTableCollection struct {

	// Array of source or target schema tables of a pipeline's assigned connection.
	Items []PipelineSchemaTableSummary `mandatory:"true" json:"items"`

	// The schema name from the database connection.
	SourceSchemaName *string `mandatory:"false" json:"sourceSchemaName"`

	// The schema name from the database connection.
	TargetSchemaName *string `mandatory:"false" json:"targetSchemaName"`
}

func (m PipelineSchemaTableCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PipelineSchemaTableCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
