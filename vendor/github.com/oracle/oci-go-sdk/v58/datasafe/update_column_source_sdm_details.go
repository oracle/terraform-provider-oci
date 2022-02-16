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

// UpdateColumnSourceSdmDetails Details of the sensitive data model to be associated as the column source with a masking policy.
type UpdateColumnSourceSdmDetails struct {

	// The OCID of the sensitive data model to be associated as the column source with the masking policy.
	SensitiveDataModelId *string `mandatory:"true" json:"sensitiveDataModelId"`
}

func (m UpdateColumnSourceSdmDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateColumnSourceSdmDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateColumnSourceSdmDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateColumnSourceSdmDetails UpdateColumnSourceSdmDetails
	s := struct {
		DiscriminatorParam string `json:"columnSource"`
		MarshalTypeUpdateColumnSourceSdmDetails
	}{
		"SENSITIVE_DATA_MODEL",
		(MarshalTypeUpdateColumnSourceSdmDetails)(m),
	}

	return json.Marshal(&s)
}
