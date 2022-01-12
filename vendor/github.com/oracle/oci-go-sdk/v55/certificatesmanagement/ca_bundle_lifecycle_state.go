// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

// CaBundleLifecycleStateEnum Enum with underlying type: string
type CaBundleLifecycleStateEnum string

// Set of constants representing the allowable values for CaBundleLifecycleStateEnum
const (
	CaBundleLifecycleStateCreating CaBundleLifecycleStateEnum = "CREATING"
	CaBundleLifecycleStateActive   CaBundleLifecycleStateEnum = "ACTIVE"
	CaBundleLifecycleStateUpdating CaBundleLifecycleStateEnum = "UPDATING"
	CaBundleLifecycleStateDeleting CaBundleLifecycleStateEnum = "DELETING"
	CaBundleLifecycleStateDeleted  CaBundleLifecycleStateEnum = "DELETED"
	CaBundleLifecycleStateFailed   CaBundleLifecycleStateEnum = "FAILED"
)

var mappingCaBundleLifecycleState = map[string]CaBundleLifecycleStateEnum{
	"CREATING": CaBundleLifecycleStateCreating,
	"ACTIVE":   CaBundleLifecycleStateActive,
	"UPDATING": CaBundleLifecycleStateUpdating,
	"DELETING": CaBundleLifecycleStateDeleting,
	"DELETED":  CaBundleLifecycleStateDeleted,
	"FAILED":   CaBundleLifecycleStateFailed,
}

// GetCaBundleLifecycleStateEnumValues Enumerates the set of values for CaBundleLifecycleStateEnum
func GetCaBundleLifecycleStateEnumValues() []CaBundleLifecycleStateEnum {
	values := make([]CaBundleLifecycleStateEnum, 0)
	for _, v := range mappingCaBundleLifecycleState {
		values = append(values, v)
	}
	return values
}
