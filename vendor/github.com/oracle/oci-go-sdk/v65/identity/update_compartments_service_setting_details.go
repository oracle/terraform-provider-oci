// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, policies, and identity domains.
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateCompartmentsServiceSettingDetails The raw data structure for creating and updating compartments delete service setting
type UpdateCompartmentsServiceSettingDetails struct {

	// Service name that has been onboarded to Identity
	ServiceName *string `mandatory:"true" json:"service_name"`

	// Resource kind name that has been onboarded to Identity
	ResourceName *string `mandatory:"true" json:"resource_name"`

	// Phonebook id of the service team
	PhoneBook *string `mandatory:"true" json:"phone_book"`

	// The endpoint to be called for compartments delete
	UrlPattern *string `mandatory:"true" json:"url_pattern"`

	// The type of service endpoint. Available action [LIST, SKIP, GET, POST]
	Action UpdateCompartmentsServiceSettingDetailsActionEnum `mandatory:"true" json:"action"`

	// True if this resource is GAed
	ResourceGa *bool `mandatory:"true" json:"resource_ga"`

	// The compartment id that is used for authz
	AuthzCompartment *string `mandatory:"true" json:"authz_compartment"`

	// If present, it overrides the url pattern for particular regions
	UrlOverrides map[string]string `mandatory:"false" json:"url_overrides"`

	// If true, the response should contain a json array, otherwise a json object
	UseListContainer *bool `mandatory:"false" json:"use_list_container"`

	// If a compartment delete found the resource with the specified states, then this compartment can be safely deleted. Eg. ["DELETED"]
	TerminalStates []string `mandatory:"false" json:"terminal_states"`

	// If a service team wants to do a special filtering on the resources, contact compartments team beforehand
	ResponseFilter *string `mandatory:"false" json:"response_filter"`

	// Query pattern
	Query *string `mandatory:"false" json:"query"`

	// The compartment id that for validation. This compartment should contain your resource(s)
	CompartmentWithResource *string `mandatory:"false" json:"compartment_with_resource"`
}

func (m UpdateCompartmentsServiceSettingDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateCompartmentsServiceSettingDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpdateCompartmentsServiceSettingDetailsActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetUpdateCompartmentsServiceSettingDetailsActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateCompartmentsServiceSettingDetailsActionEnum Enum with underlying type: string
type UpdateCompartmentsServiceSettingDetailsActionEnum string

// Set of constants representing the allowable values for UpdateCompartmentsServiceSettingDetailsActionEnum
const (
	UpdateCompartmentsServiceSettingDetailsActionList UpdateCompartmentsServiceSettingDetailsActionEnum = "LIST"
	UpdateCompartmentsServiceSettingDetailsActionSkip UpdateCompartmentsServiceSettingDetailsActionEnum = "SKIP"
	UpdateCompartmentsServiceSettingDetailsActionGet  UpdateCompartmentsServiceSettingDetailsActionEnum = "GET"
	UpdateCompartmentsServiceSettingDetailsActionPost UpdateCompartmentsServiceSettingDetailsActionEnum = "POST"
)

var mappingUpdateCompartmentsServiceSettingDetailsActionEnum = map[string]UpdateCompartmentsServiceSettingDetailsActionEnum{
	"LIST": UpdateCompartmentsServiceSettingDetailsActionList,
	"SKIP": UpdateCompartmentsServiceSettingDetailsActionSkip,
	"GET":  UpdateCompartmentsServiceSettingDetailsActionGet,
	"POST": UpdateCompartmentsServiceSettingDetailsActionPost,
}

var mappingUpdateCompartmentsServiceSettingDetailsActionEnumLowerCase = map[string]UpdateCompartmentsServiceSettingDetailsActionEnum{
	"list": UpdateCompartmentsServiceSettingDetailsActionList,
	"skip": UpdateCompartmentsServiceSettingDetailsActionSkip,
	"get":  UpdateCompartmentsServiceSettingDetailsActionGet,
	"post": UpdateCompartmentsServiceSettingDetailsActionPost,
}

// GetUpdateCompartmentsServiceSettingDetailsActionEnumValues Enumerates the set of values for UpdateCompartmentsServiceSettingDetailsActionEnum
func GetUpdateCompartmentsServiceSettingDetailsActionEnumValues() []UpdateCompartmentsServiceSettingDetailsActionEnum {
	values := make([]UpdateCompartmentsServiceSettingDetailsActionEnum, 0)
	for _, v := range mappingUpdateCompartmentsServiceSettingDetailsActionEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateCompartmentsServiceSettingDetailsActionEnumStringValues Enumerates the set of values in String for UpdateCompartmentsServiceSettingDetailsActionEnum
func GetUpdateCompartmentsServiceSettingDetailsActionEnumStringValues() []string {
	return []string{
		"LIST",
		"SKIP",
		"GET",
		"POST",
	}
}

// GetMappingUpdateCompartmentsServiceSettingDetailsActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateCompartmentsServiceSettingDetailsActionEnum(val string) (UpdateCompartmentsServiceSettingDetailsActionEnum, bool) {
	enum, ok := mappingUpdateCompartmentsServiceSettingDetailsActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
