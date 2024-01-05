// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// EM Warehouse API
//
// Use the EM Warehouse API to manage EM Warehouse data collection.
//

package emwarehouse

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResourceUsage The resource usage information.
type ResourceUsage struct {

	// operations Insights Warehouse Identifier
	OperationsInsightsWarehouseId *string `mandatory:"true" json:"operationsInsightsWarehouseId"`

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// EmInstanceCount
	EmInstanceCount *int `mandatory:"false" json:"emInstanceCount"`

	// EmInstance Target count
	TargetsCount *int `mandatory:"false" json:"targetsCount"`

	// List of emInstances
	EmInstances []EmInstancesDetails `mandatory:"false" json:"emInstances"`

	// schema name
	SchemaName *string `mandatory:"false" json:"schemaName"`
}

func (m ResourceUsage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourceUsage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
