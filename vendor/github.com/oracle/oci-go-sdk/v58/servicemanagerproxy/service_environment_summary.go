// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Manager Proxy API
//
// Use the Service Manager Proxy API to obtain information about SaaS environments provisioned by Service Manager.
// You can get information such as service types and service environment URLs.
//

package servicemanagerproxy

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ServiceEnvironmentSummary Summary of service environment details.
type ServiceEnvironmentSummary struct {

	// Unqiue identifier for the entitlement related to the environment.
	// **Note:** Not an OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"true" json:"id"`

	// The unique subscription ID associated with the service environment ID.
	// **Note:** Not an OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SubscriptionId *string `mandatory:"true" json:"subscriptionId"`

	// Status of the entitlement registration for the service.
	Status ServiceEntitlementRegistrationStatusEnum `mandatory:"true" json:"status"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) for the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ServiceDefinition *ServiceDefinition `mandatory:"true" json:"serviceDefinition"`

	// The URL for the console.
	ConsoleUrl *string `mandatory:"false" json:"consoleUrl"`

	// Array of service environment end points.
	ServiceEnvironmentEndpoints []ServiceEnvironmentEndPointOverview `mandatory:"false" json:"serviceEnvironmentEndpoints"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"CostCenter": "42"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m ServiceEnvironmentSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ServiceEnvironmentSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingServiceEntitlementRegistrationStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetServiceEntitlementRegistrationStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
