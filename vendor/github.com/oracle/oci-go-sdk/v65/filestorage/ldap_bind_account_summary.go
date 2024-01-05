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

// LdapBindAccountSummary Summary information for the LDAP bind account used by the outbound connector.
type LdapBindAccountSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the outbound connector.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the outbound connector.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My outbound connector`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time the outbound connector was created
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Array of server endpoints to use when connecting with the LDAP bind account.
	Endpoints []Endpoint `mandatory:"true" json:"endpoints"`

	// The LDAP Distinguished Name of the account.
	BindDistinguishedName *string `mandatory:"true" json:"bindDistinguishedName"`

	// The availability domain the outbound connector is in. May be unset
	// as a blank or NULL value.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Free-form tags for this resource. Each tag is a simple key-value pair
	//  with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The current state of this outbound connector.
	LifecycleState OutboundConnectorSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetAvailabilityDomain returns AvailabilityDomain
func (m LdapBindAccountSummary) GetAvailabilityDomain() *string {
	return m.AvailabilityDomain
}

// GetCompartmentId returns CompartmentId
func (m LdapBindAccountSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetId returns Id
func (m LdapBindAccountSummary) GetId() *string {
	return m.Id
}

// GetLifecycleState returns LifecycleState
func (m LdapBindAccountSummary) GetLifecycleState() OutboundConnectorSummaryLifecycleStateEnum {
	return m.LifecycleState
}

// GetDisplayName returns DisplayName
func (m LdapBindAccountSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetTimeCreated returns TimeCreated
func (m LdapBindAccountSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetFreeformTags returns FreeformTags
func (m LdapBindAccountSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m LdapBindAccountSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m LdapBindAccountSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LdapBindAccountSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOutboundConnectorSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOutboundConnectorSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m LdapBindAccountSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLdapBindAccountSummary LdapBindAccountSummary
	s := struct {
		DiscriminatorParam string `json:"connectorType"`
		MarshalTypeLdapBindAccountSummary
	}{
		"LDAPBIND",
		(MarshalTypeLdapBindAccountSummary)(m),
	}

	return json.Marshal(&s)
}
