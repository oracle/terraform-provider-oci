// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RestoreOpensearchClusterDetails Information about the OpenSearch cluster backup to restore.
type RestoreOpensearchClusterDetails struct {

	// The OCID of the cluster backup to restore.
	OpensearchClusterBackupId *string `mandatory:"true" json:"opensearchClusterBackupId"`

	// The OCID of the compartment where the cluster backup is located.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The prefix for the indices in the cluster backup.
	Prefix *string `mandatory:"false" json:"prefix"`
}

func (m RestoreOpensearchClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RestoreOpensearchClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
