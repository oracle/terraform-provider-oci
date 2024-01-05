// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateDataTransferMediumDetails Data Transfer Medium details for the Migration.
// Only one type of data transfer medium can be specified, except for the case of Amazon RDS Oracle as source,
// where Object Storage Details along with AwsS3Details are required.
// If an empty object is specified, the stored Data Transfer Medium details will be removed.
type UpdateDataTransferMediumDetails struct {
	DatabaseLinkDetails *UpdateDatabaseLinkDetails `mandatory:"false" json:"databaseLinkDetails"`

	ObjectStorageDetails *UpdateObjectStoreBucket `mandatory:"false" json:"objectStorageDetails"`

	AwsS3Details *UpdateAwsS3Details `mandatory:"false" json:"awsS3Details"`
}

func (m UpdateDataTransferMediumDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDataTransferMediumDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
