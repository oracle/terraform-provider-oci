// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OciCliDumpTransferDetails Optional dump transfer details for OCI-CLI-based dump transfer in source or target host.
type OciCliDumpTransferDetails struct {

	// Directory path to OCI SSL wallet location on Db server node.
	WalletLocation *string `mandatory:"false" json:"walletLocation"`

	// Path to the OCI CLI installation in the node.
	OciHome *string `mandatory:"false" json:"ociHome"`
}

// GetWalletLocation returns WalletLocation
func (m OciCliDumpTransferDetails) GetWalletLocation() *string {
	return m.WalletLocation
}

func (m OciCliDumpTransferDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciCliDumpTransferDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OciCliDumpTransferDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOciCliDumpTransferDetails OciCliDumpTransferDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeOciCliDumpTransferDetails
	}{
		"OCI_CLI",
		(MarshalTypeOciCliDumpTransferDetails)(m),
	}

	return json.Marshal(&s)
}
