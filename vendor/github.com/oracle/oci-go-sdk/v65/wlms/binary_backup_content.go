// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// WebLogic Management Service API
//
// WebLogic Management Service is an OCI service that enables a unified view and management of WebLogic domains
// in Oracle Cloud Infrastructure. Features include on-demand patching of WebLogic domains, rollback of the
// last applied patch, discovery and management of WebLogic instances on a compute host.
//

package wlms

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BinaryBackupContent Information of the content of a binary backup. A binary backup contains middleware binaries.
type BinaryBackupContent struct {
	Middleware *MiddlewareBinaryBackupContent `mandatory:"true" json:"middleware"`
}

func (m BinaryBackupContent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BinaryBackupContent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BinaryBackupContent) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBinaryBackupContent BinaryBackupContent
	s := struct {
		DiscriminatorParam string `json:"contentType"`
		MarshalTypeBinaryBackupContent
	}{
		"BINARY",
		(MarshalTypeBinaryBackupContent)(m),
	}

	return json.Marshal(&s)
}
