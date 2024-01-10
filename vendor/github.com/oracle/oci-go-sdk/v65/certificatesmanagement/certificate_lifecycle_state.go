// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

import (
	"strings"
)

// CertificateLifecycleStateEnum Enum with underlying type: string
type CertificateLifecycleStateEnum string

// Set of constants representing the allowable values for CertificateLifecycleStateEnum
const (
	CertificateLifecycleStateCreating           CertificateLifecycleStateEnum = "CREATING"
	CertificateLifecycleStateActive             CertificateLifecycleStateEnum = "ACTIVE"
	CertificateLifecycleStateUpdating           CertificateLifecycleStateEnum = "UPDATING"
	CertificateLifecycleStateDeleting           CertificateLifecycleStateEnum = "DELETING"
	CertificateLifecycleStateDeleted            CertificateLifecycleStateEnum = "DELETED"
	CertificateLifecycleStateSchedulingDeletion CertificateLifecycleStateEnum = "SCHEDULING_DELETION"
	CertificateLifecycleStatePendingDeletion    CertificateLifecycleStateEnum = "PENDING_DELETION"
	CertificateLifecycleStateCancellingDeletion CertificateLifecycleStateEnum = "CANCELLING_DELETION"
	CertificateLifecycleStateFailed             CertificateLifecycleStateEnum = "FAILED"
)

var mappingCertificateLifecycleStateEnum = map[string]CertificateLifecycleStateEnum{
	"CREATING":            CertificateLifecycleStateCreating,
	"ACTIVE":              CertificateLifecycleStateActive,
	"UPDATING":            CertificateLifecycleStateUpdating,
	"DELETING":            CertificateLifecycleStateDeleting,
	"DELETED":             CertificateLifecycleStateDeleted,
	"SCHEDULING_DELETION": CertificateLifecycleStateSchedulingDeletion,
	"PENDING_DELETION":    CertificateLifecycleStatePendingDeletion,
	"CANCELLING_DELETION": CertificateLifecycleStateCancellingDeletion,
	"FAILED":              CertificateLifecycleStateFailed,
}

var mappingCertificateLifecycleStateEnumLowerCase = map[string]CertificateLifecycleStateEnum{
	"creating":            CertificateLifecycleStateCreating,
	"active":              CertificateLifecycleStateActive,
	"updating":            CertificateLifecycleStateUpdating,
	"deleting":            CertificateLifecycleStateDeleting,
	"deleted":             CertificateLifecycleStateDeleted,
	"scheduling_deletion": CertificateLifecycleStateSchedulingDeletion,
	"pending_deletion":    CertificateLifecycleStatePendingDeletion,
	"cancelling_deletion": CertificateLifecycleStateCancellingDeletion,
	"failed":              CertificateLifecycleStateFailed,
}

// GetCertificateLifecycleStateEnumValues Enumerates the set of values for CertificateLifecycleStateEnum
func GetCertificateLifecycleStateEnumValues() []CertificateLifecycleStateEnum {
	values := make([]CertificateLifecycleStateEnum, 0)
	for _, v := range mappingCertificateLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCertificateLifecycleStateEnumStringValues Enumerates the set of values in String for CertificateLifecycleStateEnum
func GetCertificateLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"SCHEDULING_DELETION",
		"PENDING_DELETION",
		"CANCELLING_DELETION",
		"FAILED",
	}
}

// GetMappingCertificateLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCertificateLifecycleStateEnum(val string) (CertificateLifecycleStateEnum, bool) {
	enum, ok := mappingCertificateLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
