// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

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

var mappingCertificateAuthorityLifecycleState = map[string]CertificateAuthorityLifecycleStateEnum{
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

// GetCertificateAuthorityLifecycleStateEnumValues Enumerates the set of values for CertificateAuthorityLifecycleStateEnum
func GetCertificateAuthorityLifecycleStateEnumValues() []CertificateAuthorityLifecycleStateEnum {
	values := make([]CertificateAuthorityLifecycleStateEnum, 0)
	for _, v := range mappingCertificateAuthorityLifecycleState {
		values = append(values, v)
	}
	return values
}
