// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Key Management Service API
//
// API for managing and performing operations with keys and vaults.
//

package keymanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ScheduleVaultDeletionDetails Details for scheduling Vault deletion
type ScheduleVaultDeletionDetails struct {

	// An optional property to indicate the deletion time of the Vault.
	// The time format should comply with RFC-3339 standards. This time must be between 7 to 30 days from the time
	// when the request is received. If the property is missing, it will be set to 30 days from request time by default.
	TimeOfDeletion *common.SDKTime `mandatory:"false" json:"timeOfDeletion"`
}

func (m ScheduleVaultDeletionDetails) String() string {
	return common.PointerString(m)
}
