// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Anomaly Detection API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InlineInputJobDetails This is the specialized JSON format with an additional field for 'locationType'. This is a required field
// used for deciding if it is an inline location or contains object-storage location.
type InlineInputJobDetails struct {

	// Inline input details.
	Message *string `mandatory:"true" json:"message"`
}

func (m InlineInputJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InlineInputJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m InlineInputJobDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInlineInputJobDetails InlineInputJobDetails
	s := struct {
		DiscriminatorParam string `json:"inputType"`
		MarshalTypeInlineInputJobDetails
	}{
		"INLINE",
		(MarshalTypeInlineInputJobDetails)(m),
	}

	return json.Marshal(&s)
}
