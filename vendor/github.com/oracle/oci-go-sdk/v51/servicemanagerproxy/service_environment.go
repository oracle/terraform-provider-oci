// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Manager Proxy API
//
// API to manage Service manager proxy.
//

package servicemanagerproxy

import (
	"github.com/oracle/oci-go-sdk/v51/common"
)

// ServiceEnvironment Model describing service environment properties.
type ServiceEnvironment struct {

	// Unqiue identifier for the entitlement related to the environment.
	Id *string `mandatory:"true" json:"id"`

	// The subscription Id corresponding to the service environment Id.
	SubscriptionId *string `mandatory:"true" json:"subscriptionId"`

	// Status of the entitlement registration for the service.
	Status ServiceEntitlementRegistrationStatusEnum `mandatory:"true" json:"status"`

	// Compartment Id associated with the service.
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
