// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OciProviderSetKeyVersionDetails Used for setting KMS key versions
type OciProviderSetKeyVersionDetails struct {

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances.
	KmsKeyVersionId *string `mandatory:"true" json:"kmsKeyVersionId"`
}

func (m OciProviderSetKeyVersionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciProviderSetKeyVersionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OciProviderSetKeyVersionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOciProviderSetKeyVersionDetails OciProviderSetKeyVersionDetails
	s := struct {
		DiscriminatorParam string `json:"provider"`
		MarshalTypeOciProviderSetKeyVersionDetails
	}{
		"OCI",
		(MarshalTypeOciProviderSetKeyVersionDetails)(m),
	}

	return json.Marshal(&s)
}
