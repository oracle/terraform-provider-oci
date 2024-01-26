// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NamedCredentialSummary A summary of the named credential.
type NamedCredentialSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the named credential.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment
	// in which the named credential resides.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the named credential. Valid characters are uppercase or
	// lowercase letters, numbers, and "_". The name of the named credential
	// cannot be modified. It must be unique in the compartment and must begin with
	// an alphabetic character.
	Name *string `mandatory:"true" json:"name"`

	// The information specified by the user about the named credential.
	Description *string `mandatory:"true" json:"description"`

	// The scope of the named credential.
	Scope NamedCredentialScopeEnum `mandatory:"true" json:"scope"`

	// The type of resource associated with the named credential.
	Type ResourceTypeEnum `mandatory:"true" json:"type"`

	// The current lifecycle state of the named credential.
	LifecycleState LifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// The details of the lifecycle state.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// The date and time the named credential was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the named credential was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`
}

func (m NamedCredentialSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NamedCredentialSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNamedCredentialScopeEnum(string(m.Scope)); !ok && m.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", m.Scope, strings.Join(GetNamedCredentialScopeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingResourceTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
