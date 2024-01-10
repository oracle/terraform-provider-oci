// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateBdsApiKeyDetails API key created on user's behalf.
type CreateBdsApiKeyDetails struct {

	// The OCID of the user for whom this new generated API key pair will be created.
	UserId *string `mandatory:"true" json:"userId"`

	// Base64 passphrase used to secure the private key which will be created on user behalf.
	Passphrase *string `mandatory:"true" json:"passphrase"`

	// User friendly identifier used to uniquely differentiate between different API keys associated with this Big Data Service cluster.
	// Only ASCII alphanumeric characters with no spaces allowed.
	KeyAlias *string `mandatory:"true" json:"keyAlias"`

	// The name of the region to establish the Object Storage endpoint. See https://docs.oracle.com/en-us/iaas/api/#/en/identity/20160918/Region/
	// for additional information.
	DefaultRegion *string `mandatory:"false" json:"defaultRegion"`
}

func (m CreateBdsApiKeyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateBdsApiKeyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
