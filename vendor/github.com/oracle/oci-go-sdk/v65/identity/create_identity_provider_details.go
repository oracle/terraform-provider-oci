// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateIdentityProviderDetails The representation of CreateIdentityProviderDetails
type CreateIdentityProviderDetails interface {

	// The OCID of your tenancy.
	GetCompartmentId() *string

	// The name you assign to the `IdentityProvider` during creation.
	// The name must be unique across all `IdentityProvider` objects in the
	// tenancy and cannot be changed.
	GetName() *string

	// The description you assign to the `IdentityProvider` during creation.
	// Does not have to be unique, and it's changeable.
	GetDescription() *string

	// The identity provider service or product.
	// Supported identity providers are Oracle Identity Cloud Service (IDCS) and Microsoft
	// Active Directory Federation Services (ADFS).
	// Example: `IDCS`
	GetProductType() CreateIdentityProviderDetailsProductTypeEnum

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createidentityproviderdetails struct {
	JsonData      []byte
	FreeformTags  map[string]string                            `mandatory:"false" json:"freeformTags"`
	DefinedTags   map[string]map[string]interface{}            `mandatory:"false" json:"definedTags"`
	CompartmentId *string                                      `mandatory:"true" json:"compartmentId"`
	Name          *string                                      `mandatory:"true" json:"name"`
	Description   *string                                      `mandatory:"true" json:"description"`
	ProductType   CreateIdentityProviderDetailsProductTypeEnum `mandatory:"true" json:"productType"`
	Protocol      string                                       `json:"protocol"`
}

// UnmarshalJSON unmarshals json
func (m *createidentityproviderdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateidentityproviderdetails createidentityproviderdetails
	s := struct {
		Model Unmarshalercreateidentityproviderdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.Name = s.Model.Name
	m.Description = s.Model.Description
	m.ProductType = s.Model.ProductType
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.Protocol = s.Model.Protocol

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createidentityproviderdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Protocol {
	case "SAML2":
		mm := CreateSaml2IdentityProviderDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateIdentityProviderDetails: %s.", m.Protocol)
		return *m, nil
	}
}

// GetFreeformTags returns FreeformTags
func (m createidentityproviderdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createidentityproviderdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetCompartmentId returns CompartmentId
func (m createidentityproviderdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetName returns Name
func (m createidentityproviderdetails) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m createidentityproviderdetails) GetDescription() *string {
	return m.Description
}

// GetProductType returns ProductType
func (m createidentityproviderdetails) GetProductType() CreateIdentityProviderDetailsProductTypeEnum {
	return m.ProductType
}

func (m createidentityproviderdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createidentityproviderdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateIdentityProviderDetailsProductTypeEnum(string(m.ProductType)); !ok && m.ProductType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProductType: %s. Supported values are: %s.", m.ProductType, strings.Join(GetCreateIdentityProviderDetailsProductTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateIdentityProviderDetailsProductTypeEnum Enum with underlying type: string
type CreateIdentityProviderDetailsProductTypeEnum string

// Set of constants representing the allowable values for CreateIdentityProviderDetailsProductTypeEnum
const (
	CreateIdentityProviderDetailsProductTypeIdcs CreateIdentityProviderDetailsProductTypeEnum = "IDCS"
	CreateIdentityProviderDetailsProductTypeAdfs CreateIdentityProviderDetailsProductTypeEnum = "ADFS"
)

var mappingCreateIdentityProviderDetailsProductTypeEnum = map[string]CreateIdentityProviderDetailsProductTypeEnum{
	"IDCS": CreateIdentityProviderDetailsProductTypeIdcs,
	"ADFS": CreateIdentityProviderDetailsProductTypeAdfs,
}

var mappingCreateIdentityProviderDetailsProductTypeEnumLowerCase = map[string]CreateIdentityProviderDetailsProductTypeEnum{
	"idcs": CreateIdentityProviderDetailsProductTypeIdcs,
	"adfs": CreateIdentityProviderDetailsProductTypeAdfs,
}

// GetCreateIdentityProviderDetailsProductTypeEnumValues Enumerates the set of values for CreateIdentityProviderDetailsProductTypeEnum
func GetCreateIdentityProviderDetailsProductTypeEnumValues() []CreateIdentityProviderDetailsProductTypeEnum {
	values := make([]CreateIdentityProviderDetailsProductTypeEnum, 0)
	for _, v := range mappingCreateIdentityProviderDetailsProductTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateIdentityProviderDetailsProductTypeEnumStringValues Enumerates the set of values in String for CreateIdentityProviderDetailsProductTypeEnum
func GetCreateIdentityProviderDetailsProductTypeEnumStringValues() []string {
	return []string{
		"IDCS",
		"ADFS",
	}
}

// GetMappingCreateIdentityProviderDetailsProductTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateIdentityProviderDetailsProductTypeEnum(val string) (CreateIdentityProviderDetailsProductTypeEnum, bool) {
	enum, ok := mappingCreateIdentityProviderDetailsProductTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateIdentityProviderDetailsProtocolEnum Enum with underlying type: string
type CreateIdentityProviderDetailsProtocolEnum string

// Set of constants representing the allowable values for CreateIdentityProviderDetailsProtocolEnum
const (
	CreateIdentityProviderDetailsProtocolSaml2 CreateIdentityProviderDetailsProtocolEnum = "SAML2"
	CreateIdentityProviderDetailsProtocolAdfs  CreateIdentityProviderDetailsProtocolEnum = "ADFS"
)

var mappingCreateIdentityProviderDetailsProtocolEnum = map[string]CreateIdentityProviderDetailsProtocolEnum{
	"SAML2": CreateIdentityProviderDetailsProtocolSaml2,
	"ADFS":  CreateIdentityProviderDetailsProtocolAdfs,
}

var mappingCreateIdentityProviderDetailsProtocolEnumLowerCase = map[string]CreateIdentityProviderDetailsProtocolEnum{
	"saml2": CreateIdentityProviderDetailsProtocolSaml2,
	"adfs":  CreateIdentityProviderDetailsProtocolAdfs,
}

// GetCreateIdentityProviderDetailsProtocolEnumValues Enumerates the set of values for CreateIdentityProviderDetailsProtocolEnum
func GetCreateIdentityProviderDetailsProtocolEnumValues() []CreateIdentityProviderDetailsProtocolEnum {
	values := make([]CreateIdentityProviderDetailsProtocolEnum, 0)
	for _, v := range mappingCreateIdentityProviderDetailsProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateIdentityProviderDetailsProtocolEnumStringValues Enumerates the set of values in String for CreateIdentityProviderDetailsProtocolEnum
func GetCreateIdentityProviderDetailsProtocolEnumStringValues() []string {
	return []string{
		"SAML2",
		"ADFS",
	}
}

// GetMappingCreateIdentityProviderDetailsProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateIdentityProviderDetailsProtocolEnum(val string) (CreateIdentityProviderDetailsProtocolEnum, bool) {
	enum, ok := mappingCreateIdentityProviderDetailsProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
