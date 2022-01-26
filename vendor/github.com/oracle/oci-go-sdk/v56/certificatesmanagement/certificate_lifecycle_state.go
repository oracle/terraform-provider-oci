// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

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

var mappingCertificateLifecycleState = map[string]CertificateLifecycleStateEnum{
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

// GetCertificateLifecycleStateEnumValues Enumerates the set of values for CertificateLifecycleStateEnum
func GetCertificateLifecycleStateEnumValues() []CertificateLifecycleStateEnum {
	values := make([]CertificateLifecycleStateEnum, 0)
	for _, v := range mappingCertificateLifecycleState {
		values = append(values, v)
	}
	return values
}
