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

// DirectArchiveSourceDetails The details for the Archive source of the Function from Direct Archive. This allows the API caller to directly upload the function code as a base64-encoded archive file.
// It is useful when the caller wants to provide the code inline during function creation.
type DirectArchiveSourceDetails struct {
}

func (m DirectArchiveSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DirectArchiveSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DirectArchiveSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDirectArchiveSourceDetails DirectArchiveSourceDetails
	s := struct {
		DiscriminatorParam string `json:"archiveSourceType"`
		MarshalTypeDirectArchiveSourceDetails
	}{
		"DIRECT_ARCHIVE",
		(MarshalTypeDirectArchiveSourceDetails)(m),
	}

	return json.Marshal(&s)
}
