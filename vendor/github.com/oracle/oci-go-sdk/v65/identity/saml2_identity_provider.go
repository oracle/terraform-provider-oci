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

// Saml2IdentityProvider A special type of IdentityProvider that
// supports the SAML 2.0 protocol. For more information, see
// Identity Providers and Federation (https://docs.oracle.com/iaas/Content/Identity/Concepts/federation.htm).
type Saml2IdentityProvider struct {

	// The OCID of the `IdentityProvider`.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy containing the `IdentityProvider`.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name you assign to the `IdentityProvider` during creation. The name
	// must be unique across all `IdentityProvider` objects in the tenancy and
	// cannot be changed. This is the name federated users see when choosing
	// which identity provider to use when signing in to the Oracle Cloud Infrastructure
	// Console.
	Name *string `mandatory:"true" json:"name"`

	// The description you assign to the `IdentityProvider` during creation. Does
	// not have to be unique, and it's changeable.
	Description *string `mandatory:"true" json:"description"`

	// The identity provider service or product.
	// Supported identity providers are Oracle Identity Cloud Service (IDCS) and Microsoft
	// Active Directory Federation Services (ADFS).
	// Allowed values are:
	// - `ADFS`
	// - `IDCS`
	// Example: `IDCS`
	ProductType *string `mandatory:"true" json:"productType"`

	// Date and time the `IdentityProvider` was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The URL for retrieving the identity provider's metadata, which
	// contains information required for federating.
	MetadataUrl *string `mandatory:"true" json:"metadataUrl"`

	// The identity provider's signing certificate used by the IAM Service
	// to validate the SAML2 token.
	SigningCertificate *string `mandatory:"true" json:"signingCertificate"`

	// The URL to redirect federated users to for authentication with the
	// identity provider.
	RedirectUrl *string `mandatory:"true" json:"redirectUrl"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int64 `mandatory:"false" json:"inactiveStatus"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The XML that contains the information required for federating Identity with SAML2 Identity Provider.
	Metadata *string `mandatory:"false" json:"metadata"`

	// Extra name value pairs associated with this identity provider.
	// Example: `{"clientId": "app_sf3kdjf3"}`
	FreeformAttributes map[string]string `mandatory:"false" json:"freeformAttributes"`

	// The current state. After creating an `IdentityProvider`, make sure its
	// `lifecycleState` changes from CREATING to ACTIVE before using it.
	LifecycleState IdentityProviderLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m Saml2IdentityProvider) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m Saml2IdentityProvider) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetName returns Name
func (m Saml2IdentityProvider) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m Saml2IdentityProvider) GetDescription() *string {
	return m.Description
}

// GetProductType returns ProductType
func (m Saml2IdentityProvider) GetProductType() *string {
	return m.ProductType
}

// GetTimeCreated returns TimeCreated
func (m Saml2IdentityProvider) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetLifecycleState returns LifecycleState
func (m Saml2IdentityProvider) GetLifecycleState() IdentityProviderLifecycleStateEnum {
	return m.LifecycleState
}

// GetInactiveStatus returns InactiveStatus
func (m Saml2IdentityProvider) GetInactiveStatus() *int64 {
	return m.InactiveStatus
}

// GetFreeformTags returns FreeformTags
func (m Saml2IdentityProvider) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m Saml2IdentityProvider) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m Saml2IdentityProvider) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Saml2IdentityProvider) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingIdentityProviderLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetIdentityProviderLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m Saml2IdentityProvider) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSaml2IdentityProvider Saml2IdentityProvider
	s := struct {
		DiscriminatorParam string `json:"protocol"`
		MarshalTypeSaml2IdentityProvider
	}{
		"SAML2",
		(MarshalTypeSaml2IdentityProvider)(m),
	}

	return json.Marshal(&s)
}
