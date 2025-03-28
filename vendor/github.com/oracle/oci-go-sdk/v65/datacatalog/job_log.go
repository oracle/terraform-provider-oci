// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// JobLog Job log details. A job log is an audit log record inserted during the lifecycle of a job execution instance.
type JobLog struct {

	// Unique key of the job log that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// The unique key of the parent job execution for which the log resource was created.
	JobExecutionKey *string `mandatory:"false" json:"jobExecutionKey"`

	// OCID of the user who created the log record for this job. Usually the executor of the job instance.
	CreatedById *string `mandatory:"false" json:"createdById"`

	// OCID of the user who created the log record for this job. Usually the executor of the job instance.
	UpdatedById *string `mandatory:"false" json:"updatedById"`

	// Job log update time. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The date and time the job log was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Severity level for this log.
	Severity *string `mandatory:"false" json:"severity"`

	// Message for this job log.
	LogMessage *string `mandatory:"false" json:"logMessage"`

	// URI to the job log instance in the API.
	Uri *string `mandatory:"false" json:"uri"`
}

func (m JobLog) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JobLog) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
