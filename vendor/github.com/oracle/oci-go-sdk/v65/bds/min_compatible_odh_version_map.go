// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MinCompatibleOdhVersionMap A map of major ODH version to minimum ODH version that is required for current OS patch. eg. {ODH0_9: 0.9.1}
type MinCompatibleOdhVersionMap struct {

	// Type of hadoop distribution.
	OdhType BdsInstanceClusterVersionEnum `mandatory:"false" json:"odhType,omitempty"`

	// Minimum ODH version for current ODH type.
	OdhVersion *string `mandatory:"false" json:"odhVersion"`
}

func (m MinCompatibleOdhVersionMap) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MinCompatibleOdhVersionMap) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBdsInstanceClusterVersionEnum(string(m.OdhType)); !ok && m.OdhType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OdhType: %s. Supported values are: %s.", m.OdhType, strings.Join(GetBdsInstanceClusterVersionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
