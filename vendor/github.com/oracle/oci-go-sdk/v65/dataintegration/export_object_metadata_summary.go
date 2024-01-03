// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExportObjectMetadataSummary Details of the exported object
type ExportObjectMetadataSummary struct {

	// Key of the object
	Key *string `mandatory:"false" json:"key"`

	// Name of the object
	Name *string `mandatory:"false" json:"name"`

	// Object identifier
	Identifier *string `mandatory:"false" json:"identifier"`

	// Object type
	ObjectType *string `mandatory:"false" json:"objectType"`

	// Object version
	ObjectVersion *string `mandatory:"false" json:"objectVersion"`

	// Aggregator key
	AggregatorKey *string `mandatory:"false" json:"aggregatorKey"`

	// Object name path
	NamePath *string `mandatory:"false" json:"namePath"`

	// time at which this object was last updated.
	TimeUpdatedInMillis *int64 `mandatory:"false" json:"timeUpdatedInMillis"`
}

func (m ExportObjectMetadataSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExportObjectMetadataSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
