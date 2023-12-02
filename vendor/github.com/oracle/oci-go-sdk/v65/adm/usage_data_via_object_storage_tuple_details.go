// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.cloud.oracle.com/Content/application-dependency-management/home.htm).
//

package adm

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UsageDataViaObjectStorageTupleDetails Reference to an object in Object Storage as a tuple.
type UsageDataViaObjectStorageTupleDetails struct {

	// The Object Storage bucket to read the usage data from.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// The Object Storage namespace to read the usage data from.
	NamespaceName *string `mandatory:"true" json:"namespaceName"`

	// The Object Storage object name to read the usage data from.
	ObjectName *string `mandatory:"true" json:"objectName"`
}

func (m UsageDataViaObjectStorageTupleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UsageDataViaObjectStorageTupleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UsageDataViaObjectStorageTupleDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUsageDataViaObjectStorageTupleDetails UsageDataViaObjectStorageTupleDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeUsageDataViaObjectStorageTupleDetails
	}{
		"objectStorageTuple",
		(MarshalTypeUsageDataViaObjectStorageTupleDetails)(m),
	}

	return json.Marshal(&s)
}
