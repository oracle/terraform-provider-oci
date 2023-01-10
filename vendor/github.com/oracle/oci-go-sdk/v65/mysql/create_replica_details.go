// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateReplicaDetails Details required to create a read replica.
type CreateReplicaDetails struct {

	// The OCID of the DB System the read replica is associated with.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// The user-friendly name for the read replica. It does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// User provided description of the read replica.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Specifies whether the read replica can be deleted. Set to true to prevent deletion, false (default) to allow.
	// Note that if a read replica is delete protected it also prevents the entire DB System from being deleted. If
	// the DB System is delete protected, read replicas can still be deleted individually if they are not delete
	// protected themselves.
	IsDeleteProtected *bool `mandatory:"false" json:"isDeleteProtected"`
}

func (m CreateReplicaDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateReplicaDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
