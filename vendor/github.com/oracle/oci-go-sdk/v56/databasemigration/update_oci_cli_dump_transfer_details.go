// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdateOciCliDumpTransferDetails Optional dump transfer details for OCI-CLI-based dump transfer in source or target host.
type UpdateOciCliDumpTransferDetails struct {

	// Path to the OCI CLI installation in the node.
	OciHome *string `mandatory:"true" json:"ociHome"`
}

func (m UpdateOciCliDumpTransferDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateOciCliDumpTransferDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateOciCliDumpTransferDetails UpdateOciCliDumpTransferDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeUpdateOciCliDumpTransferDetails
	}{
		"OCI_CLI",
		(MarshalTypeUpdateOciCliDumpTransferDetails)(m),
	}

	return json.Marshal(&s)
}
