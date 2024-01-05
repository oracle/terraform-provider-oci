// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LibraryMaskingFormatEntry A library masking format to be used for masking. It can be either a
// predefined or a user-defined library masking format. It enables reuse
// of an existing library masking format and helps avoid defining the masking
// logic again. Use the ListLibraryMaskingFormats operation to view the
// existing library masking formats.
type LibraryMaskingFormatEntry struct {

	// The OCID of the library masking format.
	LibraryMaskingFormatId *string `mandatory:"true" json:"libraryMaskingFormatId"`

	// The description of the format entry.
	Description *string `mandatory:"false" json:"description"`
}

// GetDescription returns Description
func (m LibraryMaskingFormatEntry) GetDescription() *string {
	return m.Description
}

func (m LibraryMaskingFormatEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LibraryMaskingFormatEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m LibraryMaskingFormatEntry) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLibraryMaskingFormatEntry LibraryMaskingFormatEntry
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeLibraryMaskingFormatEntry
	}{
		"LIBRARY_MASKING_FORMAT",
		(MarshalTypeLibraryMaskingFormatEntry)(m),
	}

	return json.Marshal(&s)
}
