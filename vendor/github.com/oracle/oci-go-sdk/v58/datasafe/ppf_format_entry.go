// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// PpfFormatEntry The Post Processing Function masking format is a special masking option that
// enables you to use a custom function to further transform column values after
// they have been masked using some other masking formats. It takes the intermediate
// masked values as input and returns the final masked values. For example, you can
// use it for adding checksums or special encodings to the masked values.
// A post-processing function has the same signature as a user-defined function,
// but it passes in the masked values the masking engine generates, and returns
// the final masked values that should be used for masking. To learn more, check
// Post Processing Function in the Data Safe documentation.
type PpfFormatEntry struct {

	// The post processing function in SCHEMA_NAME.PACKAGE_NAME.FUNCTION_NAME
	// format. It can be a standalone or packaged function, so PACKAGE_NAME
	// is optional.
	PostProcessingFunction *string `mandatory:"true" json:"postProcessingFunction"`

	// The description of the format entry.
	Description *string `mandatory:"false" json:"description"`
}

//GetDescription returns Description
func (m PpfFormatEntry) GetDescription() *string {
	return m.Description
}

func (m PpfFormatEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PpfFormatEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PpfFormatEntry) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePpfFormatEntry PpfFormatEntry
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypePpfFormatEntry
	}{
		"POST_PROCESSING_FUNCTION",
		(MarshalTypePpfFormatEntry)(m),
	}

	return json.Marshal(&s)
}
