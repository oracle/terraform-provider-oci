// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// TestBdsObjectStorageConnectionDetails Test access to specified Object Storage bucket using the API key.
type TestBdsObjectStorageConnectionDetails struct {

	// An Oracle Cloud Infrastructure URI to which this connection must be attempted. See https://docs.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat.
	ObjectStorageUri *string `mandatory:"true" json:"objectStorageUri"`

	// Base64 passphrase used to secure the private key which will be created on user behalf.
	Passphrase *string `mandatory:"true" json:"passphrase"`

	// The name of the region to establish the Object Storage endpoint. Example us-phoenix-1 .
	ObjectStorageRegion *string `mandatory:"false" json:"objectStorageRegion"`
}

func (m TestBdsObjectStorageConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TestBdsObjectStorageConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
