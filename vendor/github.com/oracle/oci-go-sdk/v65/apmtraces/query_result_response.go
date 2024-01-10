// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Trace Explorer API
//
// Use the Application Performance Monitoring Trace Explorer API to query traces and associated spans in Trace Explorer. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmtraces

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// QueryResultResponse A response containing a collection of query rows (selected attributes and aggregations) filtered, grouped and
// sorted by the specified criteria from the query that is run, and the associated summary describing the corresponding
// query result metadata.
type QueryResultResponse struct {
	QueryResultMetadataSummary *QueryResultMetadataSummary `mandatory:"true" json:"queryResultMetadataSummary"`

	// A collection of objects with each object representing an individual row of the query result set.  The total number of objects
	// returned in this collection correspond to the total number of rows returned by the actual query that is run against
	// the queried entity.
	QueryResultRows []QueryResultRow `mandatory:"true" json:"queryResultRows"`
}

func (m QueryResultResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QueryResultResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
