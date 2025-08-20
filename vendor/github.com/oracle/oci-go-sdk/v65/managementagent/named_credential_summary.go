// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NamedCredentialSummary A representation of a named credential in the Management Agent.
type NamedCredentialSummary struct {

	// Identifier for Named Credential. This is the NamedCredential resource OCID.
	Id *string `mandatory:"true" json:"id"`

	// Name for Named Credential. This is unique for the Management Agent.
	Name *string `mandatory:"true" json:"name"`

	// Type of the Named Credential
	Type *string `mandatory:"true" json:"type"`

	// The Management Agent this named credential is associated with. This is the ManagementAgent resource OCID.
	ManagementAgentId *string `mandatory:"true" json:"managementAgentId"`

	// The time the Named Credential was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Properties for the named credential
	Properties []NamedCredentialProperty `mandatory:"true" json:"properties"`

	// The current state of the named credential
	LifecycleState NamedCredentialLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Description of the Named Credential.
	Description *string `mandatory:"false" json:"description"`

	// The time the Named Credential data was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{ "orcl-cloud": { "free-tier-retained": "true" } }`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m NamedCredentialSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NamedCredentialSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNamedCredentialLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetNamedCredentialLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
