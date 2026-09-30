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

// UpdateSchemaSourceFromTargetDetails Details of the target database that's used as the source of subsetting schemas
type UpdateSchemaSourceFromTargetDetails struct {

	// The OCID of the target database that's used as the source of subsetting schemas
	TargetId *string `mandatory:"false" json:"targetId"`

	// The schemas to be subsetted
	SchemasForSubsetting []string `mandatory:"false" json:"schemasForSubsetting"`
}

func (m UpdateSchemaSourceFromTargetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateSchemaSourceFromTargetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateSchemaSourceFromTargetDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateSchemaSourceFromTargetDetails UpdateSchemaSourceFromTargetDetails
	s := struct {
		DiscriminatorParam string `json:"schemaSource"`
		MarshalTypeUpdateSchemaSourceFromTargetDetails
	}{
		"TARGET",
		(MarshalTypeUpdateSchemaSourceFromTargetDetails)(m),
	}

	return json.Marshal(&s)
}
