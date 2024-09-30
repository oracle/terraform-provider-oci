// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Security Attribute API
//
// Use the Security Attributes API to manage security attributes and security attribute namespaces. For more information, see the documentation for Security Attributes (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attributes.htm) and Security Attribute Nampespaces (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attribute-namespaces.htm).
//

package securityattribute

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SecurityAttributeNamespaceSummary A container for security attributes.
type SecurityAttributeNamespaceSummary struct {

	// The OCID of the security attribute namespace.
	Id *string `mandatory:"false" json:"id"`

	// The OCID of the compartment that contains the security attribute namespace.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The name of the security attribute namespace. It must be unique across all security attribute namespaces in the tenancy and cannot be changed.
	Name *string `mandatory:"false" json:"name"`

	// A description you create for the security attribute namespace to help you identify it.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Indicates whether the security attribute namespace is retired.
	IsRetired *bool `mandatory:"false" json:"isRetired"`

	// Indicates possible modes the security attributes in the namespace can be set to.
	// This is not accepted from the user. Currently the supported values are enforce and audit.
	Mode []string `mandatory:"false" json:"mode"`

	// The security attribute namespace's current state. After creating a security attribute namespace, make sure its `lifecycleState` is ACTIVE before using it. After retiring a security attribute namespace, make sure its `lifecycleState` is INACTIVE.
	LifecycleState SecurityAttributeNamespaceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Date and time the security attribute namespace was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m SecurityAttributeNamespaceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecurityAttributeNamespaceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSecurityAttributeNamespaceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSecurityAttributeNamespaceLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
