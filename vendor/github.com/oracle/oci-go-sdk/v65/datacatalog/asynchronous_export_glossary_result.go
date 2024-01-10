// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AsynchronousExportGlossaryResult Details about the job which performs an export.
type AsynchronousExportGlossaryResult struct {

	// Display name of the export job.
	JobDefinitionName *string `mandatory:"false" json:"jobDefinitionName"`

	// Unique key of the export job definition.
	JobDefinitionKey *string `mandatory:"false" json:"jobDefinitionKey"`

	// Unique key of the export job.
	JobKey *string `mandatory:"false" json:"jobKey"`

	// Unique key of the job execution.
	JobExecutionKey *string `mandatory:"false" json:"jobExecutionKey"`

	// Unique key of the object being exported.
	SourceKey *string `mandatory:"false" json:"sourceKey"`
}

func (m AsynchronousExportGlossaryResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AsynchronousExportGlossaryResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
