// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PoolConfig An object containing the details about the compute shapes and number of compute instances to provison.
type PoolConfig struct {

	// The compute shape of the resources you would like to provision.
	Shape *string `mandatory:"false" json:"shape"`

	ShapeConfig *ShapeConfig `mandatory:"false" json:"shapeConfig"`

	// Minimum number of compute instances in the pool for a given compute shape.
	Min *int `mandatory:"false" json:"min"`

	// Maximum number of compute instances in the pool for a given compute shape.
	Max *int `mandatory:"false" json:"max"`
}

func (m PoolConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PoolConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
