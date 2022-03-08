// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v61/common"
	"strings"
)

// DataPreview The data preview response.
type DataPreview struct {

	// Name of the entity for which data preview was requested
	EntityName *string `mandatory:"true" json:"entityName"`

	// Total number of rows taken for sampling
	SampleRowsCount *int `mandatory:"false" json:"sampleRowsCount"`

	// Array of column definition for the preview result
	Columns []Column `mandatory:"false" json:"columns"`

	// Array of rows values for the preview result
	Rows []Row `mandatory:"false" json:"rows"`
}

func (m DataPreview) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataPreview) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
