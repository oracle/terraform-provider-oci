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

// GcpKeyDetails Details for GCP encryption key.
type GcpKeyDetails struct {

	// GCP key name
	KeyName *string `mandatory:"true" json:"keyName"`

	// GCP project name
	Project *string `mandatory:"true" json:"project"`

	// GCP key ring location
	Location *string `mandatory:"true" json:"location"`

	// GCP key ring
	KeyRing *string `mandatory:"true" json:"keyRing"`

	// GCP kms REST API endpoint
	KmsRestEndpoint *string `mandatory:"false" json:"kmsRestEndpoint"`
}

func (m GcpKeyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GcpKeyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GcpKeyDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGcpKeyDetails GcpKeyDetails
	s := struct {
		DiscriminatorParam string `json:"provider"`
		MarshalTypeGcpKeyDetails
	}{
		"GCP",
		(MarshalTypeGcpKeyDetails)(m),
	}

	return json.Marshal(&s)
}
