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

// DefaultPhaseOneParameters Phase One Parameters
type DefaultPhaseOneParameters struct {

	// Default Phase One Encryption Algorithms
	DefaultEncryptionAlgorithms []string `mandatory:"false" json:"defaultEncryptionAlgorithms"`

	// Default Phase One Authentication Algorithms
	DefaultAuthenticationAlgorithms []string `mandatory:"false" json:"defaultAuthenticationAlgorithms"`

	// Default DH Groups
	DefaultDhGroups []string `mandatory:"false" json:"defaultDhGroups"`
}

func (m DefaultPhaseOneParameters) String() string {
	return common.PointerString(m)
}
