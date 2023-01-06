// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EstimatedPatchingTime The estimated total time required in minutes for all patching operations (database server, storage server, and network switch patching).
type EstimatedPatchingTime struct {

	// The estimated total time required in minutes for all patching operations.
	TotalEstimatedPatchingTime *int `mandatory:"false" json:"totalEstimatedPatchingTime"`

	// The estimated time required in minutes for database server patching.
	EstimatedDbServerPatchingTime *int `mandatory:"false" json:"estimatedDbServerPatchingTime"`

	// The estimated time required in minutes for storage server patching.
	EstimatedStorageServerPatchingTime *int `mandatory:"false" json:"estimatedStorageServerPatchingTime"`

	// The estimated time required in minutes for network switch patching.
	EstimatedNetworkSwitchesPatchingTime *int `mandatory:"false" json:"estimatedNetworkSwitchesPatchingTime"`
}

func (m EstimatedPatchingTime) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EstimatedPatchingTime) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
