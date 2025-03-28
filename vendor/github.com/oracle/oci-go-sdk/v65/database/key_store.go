// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KeyStore A key store to connect to an on-premise encryption key appliance like Oracle Key Vault.
type KeyStore struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the key store. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the key store.
	LifecycleState KeyStoreLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	TypeDetails KeyStoreTypeDetails `mandatory:"true" json:"typeDetails"`

	// The date and time that the key store was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// List of databases associated with the key store.
	AssociatedDatabases []KeyStoreAssociatedDatabaseDetails `mandatory:"false" json:"associatedDatabases"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m KeyStore) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KeyStore) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingKeyStoreLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetKeyStoreLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *KeyStore) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeCreated         *common.SDKTime                     `json:"timeCreated"`
		LifecycleDetails    *string                             `json:"lifecycleDetails"`
		AssociatedDatabases []KeyStoreAssociatedDatabaseDetails `json:"associatedDatabases"`
		FreeformTags        map[string]string                   `json:"freeformTags"`
		DefinedTags         map[string]map[string]interface{}   `json:"definedTags"`
		Id                  *string                             `json:"id"`
		CompartmentId       *string                             `json:"compartmentId"`
		DisplayName         *string                             `json:"displayName"`
		LifecycleState      KeyStoreLifecycleStateEnum          `json:"lifecycleState"`
		TypeDetails         keystoretypedetails                 `json:"typeDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeCreated = model.TimeCreated

	m.LifecycleDetails = model.LifecycleDetails

	m.AssociatedDatabases = make([]KeyStoreAssociatedDatabaseDetails, len(model.AssociatedDatabases))
	copy(m.AssociatedDatabases, model.AssociatedDatabases)
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.LifecycleState = model.LifecycleState

	nn, e = model.TypeDetails.UnmarshalPolymorphicJSON(model.TypeDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TypeDetails = nn.(KeyStoreTypeDetails)
	} else {
		m.TypeDetails = nil
	}

	return
}

// KeyStoreLifecycleStateEnum Enum with underlying type: string
type KeyStoreLifecycleStateEnum string

// Set of constants representing the allowable values for KeyStoreLifecycleStateEnum
const (
	KeyStoreLifecycleStateActive         KeyStoreLifecycleStateEnum = "ACTIVE"
	KeyStoreLifecycleStateDeleted        KeyStoreLifecycleStateEnum = "DELETED"
	KeyStoreLifecycleStateNeedsAttention KeyStoreLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingKeyStoreLifecycleStateEnum = map[string]KeyStoreLifecycleStateEnum{
	"ACTIVE":          KeyStoreLifecycleStateActive,
	"DELETED":         KeyStoreLifecycleStateDeleted,
	"NEEDS_ATTENTION": KeyStoreLifecycleStateNeedsAttention,
}

var mappingKeyStoreLifecycleStateEnumLowerCase = map[string]KeyStoreLifecycleStateEnum{
	"active":          KeyStoreLifecycleStateActive,
	"deleted":         KeyStoreLifecycleStateDeleted,
	"needs_attention": KeyStoreLifecycleStateNeedsAttention,
}

// GetKeyStoreLifecycleStateEnumValues Enumerates the set of values for KeyStoreLifecycleStateEnum
func GetKeyStoreLifecycleStateEnumValues() []KeyStoreLifecycleStateEnum {
	values := make([]KeyStoreLifecycleStateEnum, 0)
	for _, v := range mappingKeyStoreLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetKeyStoreLifecycleStateEnumStringValues Enumerates the set of values in String for KeyStoreLifecycleStateEnum
func GetKeyStoreLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingKeyStoreLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKeyStoreLifecycleStateEnum(val string) (KeyStoreLifecycleStateEnum, bool) {
	enum, ok := mappingKeyStoreLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
