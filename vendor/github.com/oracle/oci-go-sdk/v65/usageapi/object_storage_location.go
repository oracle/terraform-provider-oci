// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the dimension of your choosing. The Usage API is used by the Cost Analysis tool in the Console. Also see Using the Usage API (https://docs.cloud.oracle.com/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ObjectStorageLocation The object storage location where usage or cost CSVs will be uploaded.
type ObjectStorageLocation struct {

	// The destination Object Store Region specified by the customer.
	Region *string `mandatory:"true" json:"region"`

	// The namespace needed to determine the object storage bucket.
	Namespace *string `mandatory:"true" json:"namespace"`

	// The bucket name where usage or cost CSVs will be uploaded.
	BucketName *string `mandatory:"true" json:"bucketName"`
}

func (m ObjectStorageLocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectStorageLocation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ObjectStorageLocation) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageLocation ObjectStorageLocation
	s := struct {
		DiscriminatorParam string `json:"locationType"`
		MarshalTypeObjectStorageLocation
	}{
		"OBJECT_STORAGE",
		(MarshalTypeObjectStorageLocation)(m),
	}

	return json.Marshal(&s)
}
