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

// ServiceProvider Details of the Service Provider. Service provider offers services to the customer to support the delegated resources.
type ServiceProvider struct {

	// Unique identifier for the Service Provider.
	Id *string `mandatory:"true" json:"id"`

	// Unique name of the Service Provider.
	Name *string `mandatory:"true" json:"name"`

	// Service Provider type.
	ServiceProviderType ServiceProviderServiceProviderTypeEnum `mandatory:"true" json:"serviceProviderType"`

	// Types of services offered by this provider.
	ServiceTypes []ServiceProviderServiceTypeEnum `mandatory:"true" json:"serviceTypes"`

	// Resource types for which this provider will provide service. Default to all if not specified.
	SupportedResourceTypes []DelegationControlResourceTypeEnum `mandatory:"true" json:"supportedResourceTypes"`

	// The OCID of the compartment that contains the Delegation Control.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Description of the Service Provider.
	Description *string `mandatory:"false" json:"description"`

	// The current lifecycle state of the Service Provider.
	LifecycleState ServiceProviderLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Description of the current lifecycle state in more detail.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// Time when the Service Provider was created expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z'
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time when the Service Provider was last modified expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z'
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ServiceProvider) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ServiceProvider) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingServiceProviderServiceProviderTypeEnum(string(m.ServiceProviderType)); !ok && m.ServiceProviderType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceProviderType: %s. Supported values are: %s.", m.ServiceProviderType, strings.Join(GetServiceProviderServiceProviderTypeEnumStringValues(), ",")))
	}
	for _, val := range m.ServiceTypes {
		if _, ok := GetMappingServiceProviderServiceTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceTypes: %s. Supported values are: %s.", val, strings.Join(GetServiceProviderServiceTypeEnumStringValues(), ",")))
		}
	}

	for _, val := range m.SupportedResourceTypes {
		if _, ok := GetMappingDelegationControlResourceTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SupportedResourceTypes: %s. Supported values are: %s.", val, strings.Join(GetDelegationControlResourceTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingServiceProviderLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetServiceProviderLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ServiceProviderServiceProviderTypeEnum Enum with underlying type: string
type ServiceProviderServiceProviderTypeEnum string

// Set of constants representing the allowable values for ServiceProviderServiceProviderTypeEnum
const (
	ServiceProviderServiceProviderTypeOracleProvided ServiceProviderServiceProviderTypeEnum = "ORACLE_PROVIDED"
)

var mappingServiceProviderServiceProviderTypeEnum = map[string]ServiceProviderServiceProviderTypeEnum{
	"ORACLE_PROVIDED": ServiceProviderServiceProviderTypeOracleProvided,
}

var mappingServiceProviderServiceProviderTypeEnumLowerCase = map[string]ServiceProviderServiceProviderTypeEnum{
	"oracle_provided": ServiceProviderServiceProviderTypeOracleProvided,
}

// GetServiceProviderServiceProviderTypeEnumValues Enumerates the set of values for ServiceProviderServiceProviderTypeEnum
func GetServiceProviderServiceProviderTypeEnumValues() []ServiceProviderServiceProviderTypeEnum {
	values := make([]ServiceProviderServiceProviderTypeEnum, 0)
	for _, v := range mappingServiceProviderServiceProviderTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetServiceProviderServiceProviderTypeEnumStringValues Enumerates the set of values in String for ServiceProviderServiceProviderTypeEnum
func GetServiceProviderServiceProviderTypeEnumStringValues() []string {
	return []string{
		"ORACLE_PROVIDED",
	}
}

// GetMappingServiceProviderServiceProviderTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingServiceProviderServiceProviderTypeEnum(val string) (ServiceProviderServiceProviderTypeEnum, bool) {
	enum, ok := mappingServiceProviderServiceProviderTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ServiceProviderLifecycleStateEnum Enum with underlying type: string
type ServiceProviderLifecycleStateEnum string

// Set of constants representing the allowable values for ServiceProviderLifecycleStateEnum
const (
	ServiceProviderLifecycleStateCreating ServiceProviderLifecycleStateEnum = "CREATING"
	ServiceProviderLifecycleStateActive   ServiceProviderLifecycleStateEnum = "ACTIVE"
	ServiceProviderLifecycleStateUpdating ServiceProviderLifecycleStateEnum = "UPDATING"
	ServiceProviderLifecycleStateDeleting ServiceProviderLifecycleStateEnum = "DELETING"
	ServiceProviderLifecycleStateDeleted  ServiceProviderLifecycleStateEnum = "DELETED"
	ServiceProviderLifecycleStateFailed   ServiceProviderLifecycleStateEnum = "FAILED"
)

var mappingServiceProviderLifecycleStateEnum = map[string]ServiceProviderLifecycleStateEnum{
	"CREATING": ServiceProviderLifecycleStateCreating,
	"ACTIVE":   ServiceProviderLifecycleStateActive,
	"UPDATING": ServiceProviderLifecycleStateUpdating,
	"DELETING": ServiceProviderLifecycleStateDeleting,
	"DELETED":  ServiceProviderLifecycleStateDeleted,
	"FAILED":   ServiceProviderLifecycleStateFailed,
}

var mappingServiceProviderLifecycleStateEnumLowerCase = map[string]ServiceProviderLifecycleStateEnum{
	"creating": ServiceProviderLifecycleStateCreating,
	"active":   ServiceProviderLifecycleStateActive,
	"updating": ServiceProviderLifecycleStateUpdating,
	"deleting": ServiceProviderLifecycleStateDeleting,
	"deleted":  ServiceProviderLifecycleStateDeleted,
	"failed":   ServiceProviderLifecycleStateFailed,
}

// GetServiceProviderLifecycleStateEnumValues Enumerates the set of values for ServiceProviderLifecycleStateEnum
func GetServiceProviderLifecycleStateEnumValues() []ServiceProviderLifecycleStateEnum {
	values := make([]ServiceProviderLifecycleStateEnum, 0)
	for _, v := range mappingServiceProviderLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetServiceProviderLifecycleStateEnumStringValues Enumerates the set of values in String for ServiceProviderLifecycleStateEnum
func GetServiceProviderLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingServiceProviderLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingServiceProviderLifecycleStateEnum(val string) (ServiceProviderLifecycleStateEnum, bool) {
	enum, ok := mappingServiceProviderLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
