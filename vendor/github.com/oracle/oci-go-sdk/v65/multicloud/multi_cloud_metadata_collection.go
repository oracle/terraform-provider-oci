// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Multicloud API
//
// Use the Oracle Multicloud API to retrieve resource anchors and network anchors, and the metadata mappings related a Cloud Service Provider. For more information, see Oracle Multicloud Hub (https://docs.oracle.com/iaas/Content/multicloud-hub/home.htm).
//

package multicloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MultiCloudMetadataCollection Multicloud metadata for Multicloud subscriptions in the indicated compartment.
// For more information, see
// Listing Multicloud Metadata for a Subscription (https://docs.oracle.com/iaas/Content/multicloud-hub/list-subscription-metadata.htm).
type MultiCloudMetadataCollection struct {

	// List of MultiCloudMetadata.
	Items []MultiCloudMetadataSummary `mandatory:"true" json:"items"`
}

func (m MultiCloudMetadataCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MultiCloudMetadataCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
