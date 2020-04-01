// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Secrets Management API
//
// API for managing secrets.
//

package vault

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ScheduleSecretDeletionDetails Details for scheduling the deletion of the specified secret.
type ScheduleSecretDeletionDetails struct {

	// An optional property indicating when to delete the secret version, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeOfDeletion *common.SDKTime `mandatory:"false" json:"timeOfDeletion"`
}

func (m ScheduleSecretDeletionDetails) String() string {
	return common.PointerString(m)
}
