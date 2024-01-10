// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateOpensearchClusterDetails The configuration to update on an existing OpenSearch cluster. Software version
// and security config are not allowed to be updated at the same time.
type UpdateOpensearchClusterDetails struct {

	// The name of the cluster. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	SoftwareVersion *string `mandatory:"false" json:"softwareVersion"`

	// The security mode of the cluster.
	SecurityMode SecurityModeEnum `mandatory:"false" json:"securityMode,omitempty"`

	// The name of the master user that are used to manage security config
	SecurityMasterUserName *string `mandatory:"false" json:"securityMasterUserName"`

	// The password hash of the master user that are used to manage security config
	SecurityMasterUserPasswordHash *string `mandatory:"false" json:"securityMasterUserPasswordHash"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateOpensearchClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOpensearchClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSecurityModeEnum(string(m.SecurityMode)); !ok && m.SecurityMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityMode: %s. Supported values are: %s.", m.SecurityMode, strings.Join(GetSecurityModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
