// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Cache API
//
// Use the OCI Cache API to create and manage clusters. A cluster is a memory-based storage solution. For more information, see OCI Cache (https://docs.oracle.com/iaas/Content/ocicache/home.htm).
//

package redis

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OciCacheUserSummary OCI Cache user summary.
type OciCacheUserSummary struct {

	// OCI Cache user unique ID.
	Id *string `mandatory:"true" json:"id"`

	// OCI Cache user name.
	Name *string `mandatory:"true" json:"name"`

	// OCI Cache user compartment ID.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OCI Cache user authentication type.
	AuthenticationType AuthenticationModeAuthenticationTypeEnum `mandatory:"true" json:"authenticationType"`

	// OCI Cache user status. ON enables and OFF disables the OCI cache user to use the cluster.
	Status OciCacheUserStatusEnum `mandatory:"true" json:"status"`

	// OCI Cache user lifecycle state.
	LifecycleState OciCacheUserLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Description of OCI cache user.
	Description *string `mandatory:"true" json:"description"`

	// The date and time, when the OCI cache user was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m OciCacheUserSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciCacheUserSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAuthenticationModeAuthenticationTypeEnum(string(m.AuthenticationType)); !ok && m.AuthenticationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticationType: %s. Supported values are: %s.", m.AuthenticationType, strings.Join(GetAuthenticationModeAuthenticationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOciCacheUserStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetOciCacheUserStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOciCacheUserLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOciCacheUserLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
