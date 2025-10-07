// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Internet of Things API
//
// Use the Internet of Things (IoT) API to manage IoT domain groups, domains, and digital twin resources including models, adapters, instances, and relationships.
// For more information, see Internet of Things (https://docs.oracle.com/iaas/Content/internet-of-things/home.htm).
//

package iot

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ApexDataAccessDetails This contains configuration for apex data access.
type ApexDataAccessDetails struct {

	// Initial admin password for APEX workspace associated with the IoT domain.
	DbWorkspaceAdminInitialPassword *string `mandatory:"true" json:"dbWorkspaceAdminInitialPassword"`
}

func (m ApexDataAccessDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApexDataAccessDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ApexDataAccessDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeApexDataAccessDetails ApexDataAccessDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeApexDataAccessDetails
	}{
		"APEX",
		(MarshalTypeApexDataAccessDetails)(m),
	}

	return json.Marshal(&s)
}
