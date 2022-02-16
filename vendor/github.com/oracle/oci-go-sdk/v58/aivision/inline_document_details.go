// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// VisionService API
//
// A description of the VisionService API.
//

package aivision

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// InlineDocumentDetails Document incorporated in the request payload.
type InlineDocumentDetails struct {

	// Document raw data.
	Data []byte `mandatory:"true" json:"data"`
}

func (m InlineDocumentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InlineDocumentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m InlineDocumentDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInlineDocumentDetails InlineDocumentDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeInlineDocumentDetails
	}{
		"INLINE",
		(MarshalTypeInlineDocumentDetails)(m),
	}

	return json.Marshal(&s)
}
