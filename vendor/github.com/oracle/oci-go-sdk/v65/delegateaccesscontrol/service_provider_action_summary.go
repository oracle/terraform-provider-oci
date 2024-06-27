// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Delegate Access Control API
//
// Oracle Delegate Access Control allows ExaCC and ExaCS customers to delegate management of their Exadata resources operators outside their tenancies.
// With Delegate Access Control, Support Providers can deliver managed services using comprehensive and robust tooling built on the OCI platform.
// Customers maintain control over who has access to the delegated resources in their tenancy and what actions can be taken.
// Enterprises managing resources across multiple tenants can use Delegate Access Control to streamline management tasks.
// Using logging service, customers can view a near real-time audit report of all actions performed by a Service Provider operator.
//

package delegateaccesscontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ServiceProviderActionSummary Details of the Service Provider Action. Service provider actions are pre-defined set of commands available to the support operator on different layers of the infrastructure.
type ServiceProviderActionSummary struct {

	// Unique identifier assigned by Oracle to a Service Provider Action.
	Id *string `mandatory:"true" json:"id"`

	// Name of the Service Provider Action.
	Name *string `mandatory:"true" json:"name"`

	// Display Name of the Service Provider Action.
	CustomerDisplayName *string `mandatory:"false" json:"customerDisplayName"`

	// Name of the component for which the Service Provider Action is applicable.
	Component *string `mandatory:"false" json:"component"`

	// resourceType for which the ServiceProviderAction is applicable
	ResourceType DelegationControlResourceTypeEnum `mandatory:"false" json:"resourceType,omitempty"`

	// List of Service Provider Service Types that this Service Provider Action is applicable to.
	ServiceProviderServiceTypes []ServiceProviderServiceTypeEnum `mandatory:"false" json:"serviceProviderServiceTypes,omitempty"`

	// The current lifecycle state of the Service Provider Action.
	LifecycleState ServiceProviderActionLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Description of the Service Provider Action in terms of associated risk profile, and characteristics of the operating system commands made
	// available to the support operator under this Service Provider Action.
	Description *string `mandatory:"false" json:"description"`
}

func (m ServiceProviderActionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ServiceProviderActionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDelegationControlResourceTypeEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetDelegationControlResourceTypeEnumStringValues(), ",")))
	}
	for _, val := range m.ServiceProviderServiceTypes {
		if _, ok := GetMappingServiceProviderServiceTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceProviderServiceTypes: %s. Supported values are: %s.", val, strings.Join(GetServiceProviderServiceTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingServiceProviderActionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetServiceProviderActionLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
