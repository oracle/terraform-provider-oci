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

// CertificateAuthorityLifecycleStateEnum Enum with underlying type: string
type CertificateAuthorityLifecycleStateEnum string

// Set of constants representing the allowable values for CertificateAuthorityLifecycleStateEnum
const (
	CertificateAuthorityLifecycleStateCreating           CertificateAuthorityLifecycleStateEnum = "CREATING"
	CertificateAuthorityLifecycleStateActive             CertificateAuthorityLifecycleStateEnum = "ACTIVE"
	CertificateAuthorityLifecycleStateUpdating           CertificateAuthorityLifecycleStateEnum = "UPDATING"
	CertificateAuthorityLifecycleStateDeleting           CertificateAuthorityLifecycleStateEnum = "DELETING"
	CertificateAuthorityLifecycleStateDeleted            CertificateAuthorityLifecycleStateEnum = "DELETED"
	CertificateAuthorityLifecycleStateSchedulingDeletion CertificateAuthorityLifecycleStateEnum = "SCHEDULING_DELETION"
	CertificateAuthorityLifecycleStatePendingDeletion    CertificateAuthorityLifecycleStateEnum = "PENDING_DELETION"
	CertificateAuthorityLifecycleStateCancellingDeletion CertificateAuthorityLifecycleStateEnum = "CANCELLING_DELETION"
	CertificateAuthorityLifecycleStateFailed             CertificateAuthorityLifecycleStateEnum = "FAILED"
)

var mappingCertificateAuthorityLifecycleStateEnum = map[string]CertificateAuthorityLifecycleStateEnum{
	"CREATING":            CertificateAuthorityLifecycleStateCreating,
	"ACTIVE":              CertificateAuthorityLifecycleStateActive,
	"UPDATING":            CertificateAuthorityLifecycleStateUpdating,
	"DELETING":            CertificateAuthorityLifecycleStateDeleting,
	"DELETED":             CertificateAuthorityLifecycleStateDeleted,
	"SCHEDULING_DELETION": CertificateAuthorityLifecycleStateSchedulingDeletion,
	"PENDING_DELETION":    CertificateAuthorityLifecycleStatePendingDeletion,
	"CANCELLING_DELETION": CertificateAuthorityLifecycleStateCancellingDeletion,
	"FAILED":              CertificateAuthorityLifecycleStateFailed,
}

var mappingCertificateAuthorityLifecycleStateEnumLowerCase = map[string]CertificateAuthorityLifecycleStateEnum{
	"creating":            CertificateAuthorityLifecycleStateCreating,
	"active":              CertificateAuthorityLifecycleStateActive,
	"updating":            CertificateAuthorityLifecycleStateUpdating,
	"deleting":            CertificateAuthorityLifecycleStateDeleting,
	"deleted":             CertificateAuthorityLifecycleStateDeleted,
	"scheduling_deletion": CertificateAuthorityLifecycleStateSchedulingDeletion,
	"pending_deletion":    CertificateAuthorityLifecycleStatePendingDeletion,
	"cancelling_deletion": CertificateAuthorityLifecycleStateCancellingDeletion,
	"failed":              CertificateAuthorityLifecycleStateFailed,
}

// GetCertificateAuthorityLifecycleStateEnumValues Enumerates the set of values for CertificateAuthorityLifecycleStateEnum
func GetCertificateAuthorityLifecycleStateEnumValues() []CertificateAuthorityLifecycleStateEnum {
	values := make([]CertificateAuthorityLifecycleStateEnum, 0)
	for _, v := range mappingCertificateAuthorityLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCertificateAuthorityLifecycleStateEnumStringValues Enumerates the set of values in String for CertificateAuthorityLifecycleStateEnum
func GetCertificateAuthorityLifecycleStateEnumStringValues() []string {
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

// GetMappingCertificateAuthorityLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCertificateAuthorityLifecycleStateEnum(val string) (CertificateAuthorityLifecycleStateEnum, bool) {
	enum, ok := mappingCertificateAuthorityLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
