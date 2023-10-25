// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.cloud.oracle.com/Content/application-dependency-management/home.htm).
//

package adm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KnowledgeBase A knowledge base is a component of Application Dependency Management (ADM) service that provides access to vulnerabilities.
type KnowledgeBase struct {

	// The Oracle Cloud Identifier (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)) of the knowledge base.
	Id *string `mandatory:"true" json:"id"`

	// The name of the knowledge base.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The creation date and time of the knowledge base (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the knowledge base was last updated (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current lifecycle state of the knowledge base.
	LifecycleState KnowledgeBaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The compartment Oracle Cloud Identifier (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)) of the knowledge base.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m KnowledgeBase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KnowledgeBase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingKnowledgeBaseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetKnowledgeBaseLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// KnowledgeBaseLifecycleStateEnum Enum with underlying type: string
type KnowledgeBaseLifecycleStateEnum string

// Set of constants representing the allowable values for KnowledgeBaseLifecycleStateEnum
const (
	KnowledgeBaseLifecycleStateCreating KnowledgeBaseLifecycleStateEnum = "CREATING"
	KnowledgeBaseLifecycleStateActive   KnowledgeBaseLifecycleStateEnum = "ACTIVE"
	KnowledgeBaseLifecycleStateUpdating KnowledgeBaseLifecycleStateEnum = "UPDATING"
	KnowledgeBaseLifecycleStateFailed   KnowledgeBaseLifecycleStateEnum = "FAILED"
	KnowledgeBaseLifecycleStateDeleting KnowledgeBaseLifecycleStateEnum = "DELETING"
	KnowledgeBaseLifecycleStateDeleted  KnowledgeBaseLifecycleStateEnum = "DELETED"
)

var mappingKnowledgeBaseLifecycleStateEnum = map[string]KnowledgeBaseLifecycleStateEnum{
	"CREATING": KnowledgeBaseLifecycleStateCreating,
	"ACTIVE":   KnowledgeBaseLifecycleStateActive,
	"UPDATING": KnowledgeBaseLifecycleStateUpdating,
	"FAILED":   KnowledgeBaseLifecycleStateFailed,
	"DELETING": KnowledgeBaseLifecycleStateDeleting,
	"DELETED":  KnowledgeBaseLifecycleStateDeleted,
}

var mappingKnowledgeBaseLifecycleStateEnumLowerCase = map[string]KnowledgeBaseLifecycleStateEnum{
	"creating": KnowledgeBaseLifecycleStateCreating,
	"active":   KnowledgeBaseLifecycleStateActive,
	"updating": KnowledgeBaseLifecycleStateUpdating,
	"failed":   KnowledgeBaseLifecycleStateFailed,
	"deleting": KnowledgeBaseLifecycleStateDeleting,
	"deleted":  KnowledgeBaseLifecycleStateDeleted,
}

// GetKnowledgeBaseLifecycleStateEnumValues Enumerates the set of values for KnowledgeBaseLifecycleStateEnum
func GetKnowledgeBaseLifecycleStateEnumValues() []KnowledgeBaseLifecycleStateEnum {
	values := make([]KnowledgeBaseLifecycleStateEnum, 0)
	for _, v := range mappingKnowledgeBaseLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetKnowledgeBaseLifecycleStateEnumStringValues Enumerates the set of values in String for KnowledgeBaseLifecycleStateEnum
func GetKnowledgeBaseLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"FAILED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingKnowledgeBaseLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKnowledgeBaseLifecycleStateEnum(val string) (KnowledgeBaseLifecycleStateEnum, bool) {
	enum, ok := mappingKnowledgeBaseLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
