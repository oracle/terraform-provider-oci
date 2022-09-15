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

// CreateCompartmentsServiceSettingDetails The raw data structure for creating and updating compartments delete service setting
type CreateCompartmentsServiceSettingDetails struct {

	// Phonebook id of the service team
	PhoneBook *string `mandatory:"true" json:"phone_book"`

	// The endpoint to be called for compartments delete
	UrlPattern *string `mandatory:"true" json:"url_pattern"`

	// The type of service endpoint. Available action [LIST, SKIP, GET, POST]
	Action CreateCompartmentsServiceSettingDetailsActionEnum `mandatory:"true" json:"action"`

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

func (m CreateCompartmentsServiceSettingDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCompartmentsServiceSettingDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateCompartmentsServiceSettingDetailsActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetCreateCompartmentsServiceSettingDetailsActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateCompartmentsServiceSettingDetailsActionEnum Enum with underlying type: string
type CreateCompartmentsServiceSettingDetailsActionEnum string

// Set of constants representing the allowable values for CreateCompartmentsServiceSettingDetailsActionEnum
const (
	CreateCompartmentsServiceSettingDetailsActionList CreateCompartmentsServiceSettingDetailsActionEnum = "LIST"
	CreateCompartmentsServiceSettingDetailsActionSkip CreateCompartmentsServiceSettingDetailsActionEnum = "SKIP"
	CreateCompartmentsServiceSettingDetailsActionGet  CreateCompartmentsServiceSettingDetailsActionEnum = "GET"
	CreateCompartmentsServiceSettingDetailsActionPost CreateCompartmentsServiceSettingDetailsActionEnum = "POST"
)

var mappingCreateCompartmentsServiceSettingDetailsActionEnum = map[string]CreateCompartmentsServiceSettingDetailsActionEnum{
	"LIST": CreateCompartmentsServiceSettingDetailsActionList,
	"SKIP": CreateCompartmentsServiceSettingDetailsActionSkip,
	"GET":  CreateCompartmentsServiceSettingDetailsActionGet,
	"POST": CreateCompartmentsServiceSettingDetailsActionPost,
}

var mappingCreateCompartmentsServiceSettingDetailsActionEnumLowerCase = map[string]CreateCompartmentsServiceSettingDetailsActionEnum{
	"list": CreateCompartmentsServiceSettingDetailsActionList,
	"skip": CreateCompartmentsServiceSettingDetailsActionSkip,
	"get":  CreateCompartmentsServiceSettingDetailsActionGet,
	"post": CreateCompartmentsServiceSettingDetailsActionPost,
}

// GetCreateCompartmentsServiceSettingDetailsActionEnumValues Enumerates the set of values for CreateCompartmentsServiceSettingDetailsActionEnum
func GetCreateCompartmentsServiceSettingDetailsActionEnumValues() []CreateCompartmentsServiceSettingDetailsActionEnum {
	values := make([]CreateCompartmentsServiceSettingDetailsActionEnum, 0)
	for _, v := range mappingCreateCompartmentsServiceSettingDetailsActionEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateCompartmentsServiceSettingDetailsActionEnumStringValues Enumerates the set of values in String for CreateCompartmentsServiceSettingDetailsActionEnum
func GetCreateCompartmentsServiceSettingDetailsActionEnumStringValues() []string {
	return []string{
		"LIST",
		"SKIP",
		"GET",
		"POST",
	}
}

// GetMappingCreateCompartmentsServiceSettingDetailsActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateCompartmentsServiceSettingDetailsActionEnum(val string) (CreateCompartmentsServiceSettingDetailsActionEnum, bool) {
	enum, ok := mappingCreateCompartmentsServiceSettingDetailsActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
