// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// BootVolumeKmsKey Kms key id associated with this volume.
type BootVolumeKmsKey struct {

	// Kms key id associated with this volume. If volume is not using KMS, then kmsKeyId will be null string.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`
}

func (m BootVolumeKmsKey) String() string {
	return common.PointerString(m)
}
