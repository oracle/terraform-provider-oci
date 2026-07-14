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

// CreateOciFssSysSecurityDetails SYS security options for OCI FSS File System.
type CreateOciFssSysSecurityDetails struct {

	// Determines whether in-transit encryption needs to be enables.
	// Check https://docs.oracle.com/en-us/iaas/Content/File/Tasks/intransitencryption.htm#Using_Intransit_Encryption for more details.
	IsEncryptedInTransit *bool `mandatory:"false" json:"isEncryptedInTransit"`
}

// GetIsEncryptedInTransit returns IsEncryptedInTransit
func (m CreateOciFssSysSecurityDetails) GetIsEncryptedInTransit() *bool {
	return m.IsEncryptedInTransit
}

func (m CreateOciFssSysSecurityDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOciFssSysSecurityDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateOciFssSysSecurityDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateOciFssSysSecurityDetails CreateOciFssSysSecurityDetails
	s := struct {
		DiscriminatorParam string `json:"auth"`
		MarshalTypeCreateOciFssSysSecurityDetails
	}{
		"SYS",
		(MarshalTypeCreateOciFssSysSecurityDetails)(m),
	}

	return json.Marshal(&s)
}
