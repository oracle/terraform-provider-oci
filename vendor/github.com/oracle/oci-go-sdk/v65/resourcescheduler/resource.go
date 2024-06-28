// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Scheduler API
//
// Use the Resource scheduler API to manage schedules, to perform actions on a collection of resources.
//

package resourcescheduler

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Resource This is the schedule resource entity.
type Resource struct {

	// This is the resource OCID.
	Id *string `mandatory:"true" json:"id"`

	// This is additional information that helps to identity the resource for the schedule.
	//     {
	//       "id": "<OCID_of_bucket>"
	//       "metadata":
	//       {
	//         "namespaceName": "sampleNamespace",
	//         "bucketName": "sampleBucket"
	//       }
	//     }
	Metadata map[string]string `mandatory:"false" json:"metadata"`
}

func (m Resource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Resource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
