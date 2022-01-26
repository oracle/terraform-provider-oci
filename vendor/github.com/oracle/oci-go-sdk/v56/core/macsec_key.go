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

// MacsecKey An object defining the Secrets-in-Vault OCIDs representing the MACsec key.
type MacsecKey struct {

	// Secret OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) containing the Connectivity association Key Name (CKN) of this MACsec key.
	ConnectivityAssociationNameSecretId *string `mandatory:"true" json:"connectivityAssociationNameSecretId"`

	// Secret OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) containing the Connectivity Association Key (CAK) of this MACsec key.
	ConnectivityAssociationKeySecretId *string `mandatory:"true" json:"connectivityAssociationKeySecretId"`

	// The secret version of the connectivity association name secret in Vault.
	ConnectivityAssociationNameSecretVersion *int64 `mandatory:"false" json:"connectivityAssociationNameSecretVersion"`

	// The secret version of the `connectivityAssociationKey` secret in Vault.
	ConnectivityAssociationKeySecretVersion *int64 `mandatory:"false" json:"connectivityAssociationKeySecretVersion"`
}

func (m MacsecKey) String() string {
	return common.PointerString(m)
}
