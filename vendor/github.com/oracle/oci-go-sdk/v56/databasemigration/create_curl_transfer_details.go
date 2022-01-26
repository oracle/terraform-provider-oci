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

// CreateCurlTransferDetails Optional properties for Curl-based dump transfer in source or target host.
type CreateCurlTransferDetails struct {
}

func (m CreateCurlTransferDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateCurlTransferDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateCurlTransferDetails CreateCurlTransferDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeCreateCurlTransferDetails
	}{
		"CURL",
		(MarshalTypeCreateCurlTransferDetails)(m),
	}

	return json.Marshal(&s)
}
