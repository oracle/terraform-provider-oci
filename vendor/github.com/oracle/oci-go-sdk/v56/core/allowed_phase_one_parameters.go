// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// AllowedPhaseOneParameters Phase One Parameters
type AllowedPhaseOneParameters struct {

	// Phase One Encryption Algorithms
	EncryptionAlgorithms []string `mandatory:"false" json:"encryptionAlgorithms"`

	// Phase One Authentication Algorithms
	AuthenticationAlgorithms []string `mandatory:"false" json:"authenticationAlgorithms"`

	// DH Groups
	DhGroups []string `mandatory:"false" json:"dhGroups"`
}

func (m AllowedPhaseOneParameters) String() string {
	return common.PointerString(m)
}
