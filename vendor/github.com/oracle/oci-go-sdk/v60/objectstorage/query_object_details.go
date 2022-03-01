// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Common set of Object Storage and Archive Storage APIs for managing buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.cloud.oracle.com/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v60/common"
	"strings"
)

// QueryObjectDetails The parameters required by Object Storage to process a query request on an object.
// To use any of the API operations, you must be authorized in an IAM policy. If you are not authorized,
// talk to an administrator. If you are an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type QueryObjectDetails struct {

	// The name of the object on which the query will be ran.
	ObjectName *string `mandatory:"true" json:"objectName"`

	// Columns to be returned in the response.
	Columns *string `mandatory:"true" json:"columns"`

	// Field to specify the logical expression to be used to filter rows.
	Predicates *string `mandatory:"true" json:"predicates"`

	// Field to specify the column(s) to be used for aggregation.
	Aggregations *string `mandatory:"true" json:"aggregations"`

	// Optional field to specify the version ID of the object.
	VersionId *string `mandatory:"false" json:"versionId"`

	// Optional field to specify the number of rows to be returned.
	Limit *int `mandatory:"false" json:"limit"`

	// Optional field to specify the data format of the target object. By default, this is autodetected.
	DataFormat *string `mandatory:"false" json:"dataFormat"`

	// Optional field to specify the data format of the results. The default ResultFormat is JSON.
	ResultFormat *string `mandatory:"false" json:"resultFormat"`
}

func (m QueryObjectDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QueryObjectDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
