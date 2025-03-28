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

// JobMetric A set of metrics are collected periodically to assess the state and performance characteristics of the execution
// instance of a job. The metrics are grouped based on their category and sub categories and aggregated based on
// their batch information.
type JobMetric struct {

	// Key of the job metric that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// Detailed description of the metric.
	Description *string `mandatory:"false" json:"description"`

	// The unique key of the parent job execution for which the job metric resource is being created.
	JobExecutionKey *string `mandatory:"false" json:"jobExecutionKey"`

	// The time the metric was logged or captured in the system where the job executed.
	// An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeInserted *common.SDKTime `mandatory:"false" json:"timeInserted"`

	// Category of this metric.
	Category *string `mandatory:"false" json:"category"`

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Sub category of this metric under the category. Used for aggregating values. May be null.
	SubCategory *string `mandatory:"false" json:"subCategory"`

	// Unit of this metric.
	Unit *string `mandatory:"false" json:"unit"`

	// Value of this metric.
	Value *string `mandatory:"false" json:"value"`

	// Batch key for grouping, may be null.
	BatchKey *string `mandatory:"false" json:"batchKey"`

	// URI to the job metric instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// The date and time the job metric was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The last time that this metric was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// OCID of the user who created the metric for this job. Usually the executor of the job instance.
	CreatedById *string `mandatory:"false" json:"createdById"`

	// OCID of the user who created the metric for this job. Usually the executor of the job instance.
	UpdatedById *string `mandatory:"false" json:"updatedById"`
}

func (m JobMetric) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JobMetric) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
