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

// QueryResultRowTypeSummary Summary of the datatype, unit and related metadata of an individual row element of a query result row that is returned.
type QueryResultRowTypeSummary struct {

	// Datatype of the query result row element.
	DataType *string `mandatory:"false" json:"dataType"`

	// Granular unit in which the query result row element's data is represented.
	Unit *string `mandatory:"false" json:"unit"`

	// Alias name if an alias is used for the query result row element or an assigned display name from the query language
	// in some default cases.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Actual show expression in the user typed query that produced this column.
	Expression *string `mandatory:"false" json:"expression"`

	// A query result row type summary object that represents a nested table structure.
	QueryResultRowTypeSummaries []QueryResultRowTypeSummary `mandatory:"false" json:"queryResultRowTypeSummaries"`
}

func (m QueryResultRowTypeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QueryResultRowTypeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
