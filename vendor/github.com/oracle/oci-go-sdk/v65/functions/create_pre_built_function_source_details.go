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

// CreatePreBuiltFunctionSourceDetails Source details for creating a function from a Pre‑Built Functions listing (PbfListing).
type CreatePreBuiltFunctionSourceDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the PbfListing this
	// function is sourced from.
	PbfListingId *string `mandatory:"true" json:"pbfListingId"`
}

func (m CreatePreBuiltFunctionSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePreBuiltFunctionSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreatePreBuiltFunctionSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreatePreBuiltFunctionSourceDetails CreatePreBuiltFunctionSourceDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeCreatePreBuiltFunctionSourceDetails
	}{
		"PRE_BUILT_FUNCTIONS",
		(MarshalTypeCreatePreBuiltFunctionSourceDetails)(m),
	}

	return json.Marshal(&s)
}
