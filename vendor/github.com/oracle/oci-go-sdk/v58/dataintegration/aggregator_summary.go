// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// AggregatorSummary A summary type containing information about the object's aggregator including its type, key, name and description.
type AggregatorSummary struct {

	// The type of the aggregator.
	Type *string `mandatory:"false" json:"type"`

	// The key of the aggregator object.
	Key *string `mandatory:"false" json:"key"`

	// The name of the aggregator.
	Name *string `mandatory:"false" json:"name"`

	// The identifier of the aggregator.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The description of the aggregator.
	Description *string `mandatory:"false" json:"description"`
}

func (m AggregatorSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AggregatorSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
