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

// CreateOciFssExportPathDetails Path used to access the associated OCI File Storage Service (FSS) File System.
type CreateOciFssExportPathDetails struct {

	// Path used to access the associated OCI File Storage Service (FSS) File System.
	Path *string `mandatory:"true" json:"path"`
}

func (m CreateOciFssExportPathDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOciFssExportPathDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateOciFssExportPathDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateOciFssExportPathDetails CreateOciFssExportPathDetails
	s := struct {
		DiscriminatorParam string `json:"ociFssExportType"`
		MarshalTypeCreateOciFssExportPathDetails
	}{
		"PATH",
		(MarshalTypeCreateOciFssExportPathDetails)(m),
	}

	return json.Marshal(&s)
}
