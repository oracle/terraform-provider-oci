// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Package utils contains methods for generating Request ID's
package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// GenerateOpcRequestID - Reference: https://confluence.oci.oraclecorp.com/display/DEX/Request+IDs
// Maximum segment length:	32 characters
// Allowed segment contents: regular expression pattern /^[a-zA-Z0-9]{0,32}$/
func GenerateOpcRequestID() string {
	clientId := generateUniqueID()
	stackId := generateUniqueID()
	individualId := generateUniqueID()

	opcRequestId := fmt.Sprintf("%s/%s/%s", clientId, stackId, individualId)

	return opcRequestId
}

func generateUniqueID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}

	return hex.EncodeToString(b)
}
