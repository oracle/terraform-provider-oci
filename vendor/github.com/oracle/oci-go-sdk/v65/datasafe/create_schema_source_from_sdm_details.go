// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// CreateSchemaSourceFromSdmDetails Details of the sensitive data model that's used as the source of subsetting schemas
type CreateSchemaSourceFromSdmDetails struct {

	// The OCID of the sensitive data model that's used as the source of subsetting schemas
	SensitiveDataModelId *string `mandatory:"true" json:"sensitiveDataModelId"`
}

func (m CreateSchemaSourceFromSdmDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSchemaSourceFromSdmDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateSchemaSourceFromSdmDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateSchemaSourceFromSdmDetails CreateSchemaSourceFromSdmDetails
	s := struct {
		DiscriminatorParam string `json:"schemaSource"`
		MarshalTypeCreateSchemaSourceFromSdmDetails
	}{
		"SENSITIVE_DATA_MODEL",
		(MarshalTypeCreateSchemaSourceFromSdmDetails)(m),
	}

	return json.Marshal(&s)
}
