// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// Use the File Storage service API to manage file systems, mount targets, and snapshots.
// For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateLdapBindAccountDetails Account details for the LDAP bind account to be used by mount targets that use this outbound connector.
type CreateLdapBindAccountDetails struct {

	// The availability domain the outbound connector is in. May be unset
	// as a blank or NULL value.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the outbound connector.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Array of server endpoints to use when connecting with the LDAP bind account.
	Endpoints []Endpoint `mandatory:"true" json:"endpoints"`

	// The LDAP Distinguished Name of the bind account.
	BindDistinguishedName *string `mandatory:"true" json:"bindDistinguishedName"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My outbound connector`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair
	//  with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the password for the LDAP bind account in the Vault.
	PasswordSecretId *string `mandatory:"false" json:"passwordSecretId"`

	// Version of the password secret in the Vault to use.
	PasswordSecretVersion *int `mandatory:"false" json:"passwordSecretVersion"`
}

// GetAvailabilityDomain returns AvailabilityDomain
func (m CreateLdapBindAccountDetails) GetAvailabilityDomain() *string {
	return m.AvailabilityDomain
}

// GetCompartmentId returns CompartmentId
func (m CreateLdapBindAccountDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m CreateLdapBindAccountDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetFreeformTags returns FreeformTags
func (m CreateLdapBindAccountDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateLdapBindAccountDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateLdapBindAccountDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateLdapBindAccountDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateLdapBindAccountDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateLdapBindAccountDetails CreateLdapBindAccountDetails
	s := struct {
		DiscriminatorParam string `json:"connectorType"`
		MarshalTypeCreateLdapBindAccountDetails
	}{
		"LDAPBIND",
		(MarshalTypeCreateLdapBindAccountDetails)(m),
	}

	return json.Marshal(&s)
}
