// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateIdentityProviderDetails The representation of UpdateIdentityProviderDetails
type UpdateIdentityProviderDetails interface {

	// The description you assign to the `IdentityProvider`. Does not have to
	// be unique, and it's changeable.
	GetDescription() *string

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type updateidentityproviderdetails struct {
	JsonData     []byte
	Description  *string                           `mandatory:"false" json:"description"`
	FreeformTags map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags  map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	Protocol     string                            `json:"protocol"`
}

// UnmarshalJSON unmarshals json
func (m *updateidentityproviderdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateidentityproviderdetails updateidentityproviderdetails
	s := struct {
		Model Unmarshalerupdateidentityproviderdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Description = s.Model.Description
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.Protocol = s.Model.Protocol

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateidentityproviderdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Protocol {
	case "SAML2":
		mm := UpdateSaml2IdentityProviderDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UpdateIdentityProviderDetails: %s.", m.Protocol)
		return *m, nil
	}
}

// GetDescription returns Description
func (m updateidentityproviderdetails) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m updateidentityproviderdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m updateidentityproviderdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m updateidentityproviderdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateidentityproviderdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateIdentityProviderDetailsProtocolEnum Enum with underlying type: string
type UpdateIdentityProviderDetailsProtocolEnum string

// Set of constants representing the allowable values for UpdateIdentityProviderDetailsProtocolEnum
const (
	UpdateIdentityProviderDetailsProtocolSaml2 UpdateIdentityProviderDetailsProtocolEnum = "SAML2"
)

var mappingUpdateIdentityProviderDetailsProtocolEnum = map[string]UpdateIdentityProviderDetailsProtocolEnum{
	"SAML2": UpdateIdentityProviderDetailsProtocolSaml2,
}

var mappingUpdateIdentityProviderDetailsProtocolEnumLowerCase = map[string]UpdateIdentityProviderDetailsProtocolEnum{
	"saml2": UpdateIdentityProviderDetailsProtocolSaml2,
}

// GetUpdateIdentityProviderDetailsProtocolEnumValues Enumerates the set of values for UpdateIdentityProviderDetailsProtocolEnum
func GetUpdateIdentityProviderDetailsProtocolEnumValues() []UpdateIdentityProviderDetailsProtocolEnum {
	values := make([]UpdateIdentityProviderDetailsProtocolEnum, 0)
	for _, v := range mappingUpdateIdentityProviderDetailsProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateIdentityProviderDetailsProtocolEnumStringValues Enumerates the set of values in String for UpdateIdentityProviderDetailsProtocolEnum
func GetUpdateIdentityProviderDetailsProtocolEnumStringValues() []string {
	return []string{
		"SAML2",
	}
}

// GetMappingUpdateIdentityProviderDetailsProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateIdentityProviderDetailsProtocolEnum(val string) (UpdateIdentityProviderDetailsProtocolEnum, bool) {
	enum, ok := mappingUpdateIdentityProviderDetailsProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
