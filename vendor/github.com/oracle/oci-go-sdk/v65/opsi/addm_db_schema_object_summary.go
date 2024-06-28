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

// AddmDbSchemaObjectSummary Details for a given object id
type AddmDbSchemaObjectSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database insight.
	Id *string `mandatory:"true" json:"id"`

	// Object id (from RDBMS)
	ObjectIdentifier *int `mandatory:"true" json:"objectIdentifier"`

	// Owner of object
	Owner *string `mandatory:"true" json:"owner"`

	// Name of object
	ObjectName *string `mandatory:"true" json:"objectName"`

	// Type of the object (such as TABLE, INDEX)
	ObjectType *string `mandatory:"true" json:"objectType"`

	// Subobject name; for example, partition name
	SubObjectName *string `mandatory:"false" json:"subObjectName"`
}

func (m AddmDbSchemaObjectSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddmDbSchemaObjectSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
