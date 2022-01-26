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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ServiceEnvironment Detailed information about a service environment.
// **Note:** Service URL formats may vary from the provided example.
type ServiceEnvironment struct {

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
}

func (m ServiceEnvironment) String() string {
	return common.PointerString(m)
}
