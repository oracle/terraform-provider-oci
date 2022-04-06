// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, policies, and identity domains.
//

package identity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateSaml2IdentityProviderDetails The representation of CreateSaml2IdentityProviderDetails
type CreateSaml2IdentityProviderDetails struct {

	// The OCID of your tenancy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name you assign to the `IdentityProvider` during creation.
	// The name must be unique across all `IdentityProvider` objects in the
	// tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// The description you assign to the `IdentityProvider` during creation.
	// Does not have to be unique, and it's changeable.
	Description *string `mandatory:"true" json:"description"`

	// The URL for retrieving the identity provider's metadata,
	// which contains information required for federating.
	MetadataUrl *string `mandatory:"true" json:"metadataUrl"`

	// The XML that contains the information required for federating.
	Metadata *string `mandatory:"true" json:"metadata"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Extra name value pairs associated with this identity provider.
	// Example: `{"clientId": "app_sf3kdjf3"}`
	FreeformAttributes map[string]string `mandatory:"false" json:"freeformAttributes"`

	// The identity provider service or product.
	// Supported identity providers are Oracle Identity Cloud Service (IDCS) and Microsoft
	// Active Directory Federation Services (ADFS).
	// Example: `IDCS`
	ProductType CreateIdentityProviderDetailsProductTypeEnum `mandatory:"true" json:"productType"`
}

//GetCompartmentId returns CompartmentId
func (m CreateSaml2IdentityProviderDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetName returns Name
func (m CreateSaml2IdentityProviderDetails) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m CreateSaml2IdentityProviderDetails) GetDescription() *string {
	return m.Description
}

//GetProductType returns ProductType
func (m CreateSaml2IdentityProviderDetails) GetProductType() CreateIdentityProviderDetailsProductTypeEnum {
	return m.ProductType
}

//GetFreeformTags returns FreeformTags
func (m CreateSaml2IdentityProviderDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateSaml2IdentityProviderDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateSaml2IdentityProviderDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSaml2IdentityProviderDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateIdentityProviderDetailsProductTypeEnum(string(m.ProductType)); !ok && m.ProductType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProductType: %s. Supported values are: %s.", m.ProductType, strings.Join(GetCreateIdentityProviderDetailsProductTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateSaml2IdentityProviderDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateSaml2IdentityProviderDetails CreateSaml2IdentityProviderDetails
	s := struct {
		DiscriminatorParam string `json:"protocol"`
		MarshalTypeCreateSaml2IdentityProviderDetails
	}{
		"SAML2",
		(MarshalTypeCreateSaml2IdentityProviderDetails)(m),
	}

	return json.Marshal(&s)
}
