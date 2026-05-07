// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateOciFssMountTargetIdDetails The OCID of the OCI File Storage Service (FSS) Mount Target. The user must have read permission for mount-targets.
type CreateOciFssMountTargetIdDetails struct {

	// The OCID of the OCI File Storage Service (FSS) Mount Target.
	Id *string `mandatory:"true" json:"id"`
}

func (m CreateOciFssMountTargetIdDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOciFssMountTargetIdDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateOciFssMountTargetIdDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateOciFssMountTargetIdDetails CreateOciFssMountTargetIdDetails
	s := struct {
		DiscriminatorParam string `json:"ociFssMountTargetType"`
		MarshalTypeCreateOciFssMountTargetIdDetails
	}{
		"OCID",
		(MarshalTypeCreateOciFssMountTargetIdDetails)(m),
	}

	return json.Marshal(&s)
}
