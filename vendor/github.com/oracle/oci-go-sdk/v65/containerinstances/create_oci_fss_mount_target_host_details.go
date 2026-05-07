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

// CreateOciFssMountTargetHostDetails IP Address or Fully Qualified Domain Name (FQDN) that can be used as host while mounting the OCI File Storage Service (FSS) File System to Containers.
type CreateOciFssMountTargetHostDetails struct {

	// Either IP Address or FQDN associated with the OCI File Storage Service (FSS) Mount Target.
	Host *string `mandatory:"true" json:"host"`
}

func (m CreateOciFssMountTargetHostDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOciFssMountTargetHostDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateOciFssMountTargetHostDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateOciFssMountTargetHostDetails CreateOciFssMountTargetHostDetails
	s := struct {
		DiscriminatorParam string `json:"ociFssMountTargetType"`
		MarshalTypeCreateOciFssMountTargetHostDetails
	}{
		"HOST",
		(MarshalTypeCreateOciFssMountTargetHostDetails)(m),
	}

	return json.Marshal(&s)
}
