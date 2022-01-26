// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Control Plane API
//

package blockchain

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// BlockchainPlatformPatchSummary Patch Details
type BlockchainPlatformPatchSummary struct {

	// patch id
	Id *string `mandatory:"false" json:"id"`

	// patch service version
	ServiceVersion *string `mandatory:"false" json:"serviceVersion"`

	// A URL for the patch specific documentation
	PatchInfoUrl *string `mandatory:"false" json:"patchInfoUrl"`

	// patch due date for customer initiated patching
	TimePatchDue *common.SDKTime `mandatory:"false" json:"timePatchDue"`
}

func (m BlockchainPlatformPatchSummary) String() string {
	return common.PointerString(m)
}
