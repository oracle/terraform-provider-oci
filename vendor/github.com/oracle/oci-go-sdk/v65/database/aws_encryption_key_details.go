// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AwsEncryptionKeyDetails Details for AWS Input
type AwsEncryptionKeyDetails struct {

	// Provide the key OCID of a registered AWS key.
	AwsEncryptionKeyId *string `mandatory:"true" json:"awsEncryptionKeyId"`
}

func (m AwsEncryptionKeyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AwsEncryptionKeyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AwsEncryptionKeyDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAwsEncryptionKeyDetails AwsEncryptionKeyDetails
	s := struct {
		DiscriminatorParam string `json:"providerType"`
		MarshalTypeAwsEncryptionKeyDetails
	}{
		"AWS",
		(MarshalTypeAwsEncryptionKeyDetails)(m),
	}

	return json.Marshal(&s)
}
