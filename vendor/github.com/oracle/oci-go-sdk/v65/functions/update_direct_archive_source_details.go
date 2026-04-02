// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateDirectArchiveSourceDetails Source details required for updating an archive based function from Direct Archive.
type UpdateDirectArchiveSourceDetails struct {

	// The base64-encoded archive file of the function code. The archive file must contain all the files for the function. Please refer to functions documentation for maximum allowed size and supported archive formats.
	ArchiveFile []byte `mandatory:"true" json:"archiveFile"`
}

func (m UpdateDirectArchiveSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDirectArchiveSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateDirectArchiveSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateDirectArchiveSourceDetails UpdateDirectArchiveSourceDetails
	s := struct {
		DiscriminatorParam string `json:"archiveSourceType"`
		MarshalTypeUpdateDirectArchiveSourceDetails
	}{
		"DIRECT_ARCHIVE",
		(MarshalTypeUpdateDirectArchiveSourceDetails)(m),
	}

	return json.Marshal(&s)
}
