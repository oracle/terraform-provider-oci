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

// UpdateVolumeKmsKeyDetails The representation of UpdateVolumeKmsKeyDetails
type UpdateVolumeKmsKeyDetails struct {

	// The new kms key which will be used to protect the specific volume.
	// This key has to be a valid kms key ocid, and user must have key delegation policy to allow them to access this key.
	// Even if this new kms key is the same as the previous kms key id, block storage service will use it to regenerate a new volume encryption key.
	// Example: `{"kmsKeyId": "ocid1.key.region1.sea.afnl2n7daag4s.abzwkljs6uevhlgcznhmh7oiatyrxngrywc3tje3uk3g77hzmewqiieuk75f"}`
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`
}

func (m UpdateVolumeKmsKeyDetails) String() string {
	return common.PointerString(m)
}
