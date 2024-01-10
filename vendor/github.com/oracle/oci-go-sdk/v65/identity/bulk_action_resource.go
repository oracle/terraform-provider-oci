// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BulkActionResource The bulk action resource entity.
type BulkActionResource struct {

	// The resource OCID.
	Identifier *string `mandatory:"true" json:"identifier"`

	// The resource-type. To get the list of supported resource-types use
	// ListBulkActionResourceTypes.
	EntityType *string `mandatory:"true" json:"entityType"`

	// Additional information that helps to identity the resource for bulk action.
	// The APIs to delete and move most resource types only require the resource identifier (ocid).
	// But some resource-types require additional identifying information.
	// This information is provided in the resource's public API document. It is also
	// available through the
	// ListBulkActionResourceTypes.
	// **Example**:
	// The APIs to delete or move the `buckets` resource-type require `namespaceName` and `bucketName` to identify the resource, as
	// shown in the APIs, DeleteBucket and
	// UpdateBucket.
	// To add a bucket for bulk actions, specify `namespaceName` and `bucketName` in
	// the metadata property as shown in this example
	//     {
	//       "identifier": "<OCID_of_bucket>"
	//       "entityType": "bucket",
	//       "metadata":
	//       {
	//         "namespaceName": "sampleNamespace",
	//         "bucketName": "sampleBucket"
	//       }
	//     }
	Metadata map[string]string `mandatory:"false" json:"metadata"`
}

func (m BulkActionResource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkActionResource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
