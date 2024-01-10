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

// OpensearchClusterSummary The summary of information about an OpenSearch cluster.
type OpensearchClusterSummary struct {

	// The OCID of the cluster.
	Id *string `mandatory:"true" json:"id"`

	// The OCID for the compartment where the cluster is located.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The software version the cluster is running.
	SoftwareVersion *string `mandatory:"true" json:"softwareVersion"`

	// The total amount of storage in GB, for the cluster.
	TotalStorageGB *int `mandatory:"true" json:"totalStorageGB"`

	// The name of the cluster. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The date and time the cluster was created. Format defined
	// by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the cluster was updated. Format defined
	// by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Additional information about the current lifecycle state of the cluster.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The current state of the cluster.
	LifecycleState OpensearchClusterLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The availability domains to distribute the cluser nodes across.
	AvailabilityDomains []string `mandatory:"false" json:"availabilityDomains"`

	// The security mode of the cluster.
	SecurityMode SecurityModeEnum `mandatory:"false" json:"securityMode,omitempty"`
}

func (m OpensearchClusterSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OpensearchClusterSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOpensearchClusterLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOpensearchClusterLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSecurityModeEnum(string(m.SecurityMode)); !ok && m.SecurityMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityMode: %s. Supported values are: %s.", m.SecurityMode, strings.Join(GetSecurityModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
