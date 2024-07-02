// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SqlText SQL Text type object.
type SqlText struct {

	// Unique SQL_ID for a SQL Statement.
	SqlIdentifier *string `mandatory:"true" json:"sqlIdentifier"`

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// SQL command
	// Example: `"SELECT"`
	SqlCommand *string `mandatory:"true" json:"sqlCommand"`

	// Full SQL Text
	// Example: `"SELECT username,profile,default_tablespace,temporary_tablespace FROM dba_users"`
	// Disclaimer: SQL text being uploaded explicitly via APIs is not masked. Any sensitive literals contained in the sqlFullText column should be masked prior to ingestion.
	SqlFullText *string `mandatory:"true" json:"sqlFullText"`

	// Version
	// Example: `1`
	Version *float32 `mandatory:"false" json:"version"`

	// Exact matching signature
	// Example: `"18067345456756876713"`
	ExactMatchingSignature *string `mandatory:"false" json:"exactMatchingSignature"`

	// Force matching signature
	// Example: `"18067345456756876713"`
	ForceMatchingSignature *string `mandatory:"false" json:"forceMatchingSignature"`
}

func (m SqlText) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlText) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
