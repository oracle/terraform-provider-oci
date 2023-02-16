// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CompartmentsServiceSetting Compartments delete config object for the service / resource
type CompartmentsServiceSetting struct {

	// The unique object id of the service setting
	ObjectId *string `mandatory:"false" json:"object_id"`

	// The service name of the resource
	ServiceName *string `mandatory:"false" json:"service_name"`

	// Name of the resource
	ResourceName *string `mandatory:"false" json:"resource_name"`

	// Phonebook id of the service team
	PhoneBook *string `mandatory:"false" json:"phone_book"`

	// The endpoint to be called for compartments delete
	UrlPattern *string `mandatory:"false" json:"url_pattern"`

	// The type of service endpoint. Eg. LIST or DELETE
	Action *string `mandatory:"false" json:"action"`

	// True if this resource is GAed
	ResourceGA *bool `mandatory:"false" json:"resource_GA"`

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

	// The compartment id that is used for authz
	AuthzCompartment *string `mandatory:"false" json:"authz_compartment"`
}

func (m CompartmentsServiceSetting) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CompartmentsServiceSetting) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
