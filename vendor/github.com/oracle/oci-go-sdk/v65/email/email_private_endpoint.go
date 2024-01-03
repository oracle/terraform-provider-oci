// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// Use the Email Delivery API to do the necessary set up to send high-volume and application-generated emails through the OCI Email Delivery service.
// For more information, see Overview of the Email Delivery Service (https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//  **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
//  If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EmailPrivateEndpoint A private endpoint allows for connections to service from customer and reverse over a private network.
type EmailPrivateEndpoint struct {

	// Unique identifier that is immutable
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Subnet Identifier
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Private endpoint type, to differentiate between requirements of each.
	// We plan to add additional types in the future so clients should tolerate unrecognized values.
	PrivateEndpointType EmailPrivateEndpointPrivateEndpointTypeEnum `mandatory:"true" json:"privateEndpointType"`

	// The current state of the private endpoint resource.
	LifecycleState EmailPrivateEndpointLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Fully qualified DNS domain name configured in the customer VCN's private DNS that is used to connect to Email Delivery service using the private endpoint.
	// For a submission endpoint type, this works the same as Email Delivery's public SMTP endpoint in the region.
	EndpointFqdn *string `mandatory:"false" json:"endpointFqdn"`

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this resource belongs to.
	// For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - The nsgIds array is optional.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Optional IP within the provided subnet for the Private Endpoint to use.
	// If not provided, an IP will be chosen from the given subnet.
	PrivateEndpointIp *string `mandatory:"false" json:"privateEndpointIp"`

	// Optional IP within the provided subnet to use for the reverse connection source IP to customer smarthost.
	// If not provided, an IP will be chosen from the given subnet.
	PrivateEndpointSourceIp *string `mandatory:"false" json:"privateEndpointSourceIp"`

	// Customer provided smart host that will route emails from Email Delivery in the case where delivery is needed in addition to submission.
	// Used when the privateEndpoint type is SUBMISSION
	SmartHost *string `mandatory:"false" json:"smartHost"`

	// The id of the Virtual Network Interface Card (VNIC) created in the customer's subnet for this private endpoint.
	PrivateEndpointVnicId *string `mandatory:"false" json:"privateEndpointVnicId"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a
	// resource in 'Failed' state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Private Endpoint display name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the private endpoint.
	Description *string `mandatory:"false" json:"description"`

	// The time the private endpoint was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the private endpoint was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m EmailPrivateEndpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EmailPrivateEndpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEmailPrivateEndpointPrivateEndpointTypeEnum(string(m.PrivateEndpointType)); !ok && m.PrivateEndpointType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrivateEndpointType: %s. Supported values are: %s.", m.PrivateEndpointType, strings.Join(GetEmailPrivateEndpointPrivateEndpointTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEmailPrivateEndpointLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEmailPrivateEndpointLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EmailPrivateEndpointPrivateEndpointTypeEnum Enum with underlying type: string
type EmailPrivateEndpointPrivateEndpointTypeEnum string

// Set of constants representing the allowable values for EmailPrivateEndpointPrivateEndpointTypeEnum
const (
	EmailPrivateEndpointPrivateEndpointTypeSubmission EmailPrivateEndpointPrivateEndpointTypeEnum = "SUBMISSION"
	EmailPrivateEndpointPrivateEndpointTypeBounce     EmailPrivateEndpointPrivateEndpointTypeEnum = "BOUNCE"
)

var mappingEmailPrivateEndpointPrivateEndpointTypeEnum = map[string]EmailPrivateEndpointPrivateEndpointTypeEnum{
	"SUBMISSION": EmailPrivateEndpointPrivateEndpointTypeSubmission,
	"BOUNCE":     EmailPrivateEndpointPrivateEndpointTypeBounce,
}

var mappingEmailPrivateEndpointPrivateEndpointTypeEnumLowerCase = map[string]EmailPrivateEndpointPrivateEndpointTypeEnum{
	"submission": EmailPrivateEndpointPrivateEndpointTypeSubmission,
	"bounce":     EmailPrivateEndpointPrivateEndpointTypeBounce,
}

// GetEmailPrivateEndpointPrivateEndpointTypeEnumValues Enumerates the set of values for EmailPrivateEndpointPrivateEndpointTypeEnum
func GetEmailPrivateEndpointPrivateEndpointTypeEnumValues() []EmailPrivateEndpointPrivateEndpointTypeEnum {
	values := make([]EmailPrivateEndpointPrivateEndpointTypeEnum, 0)
	for _, v := range mappingEmailPrivateEndpointPrivateEndpointTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEmailPrivateEndpointPrivateEndpointTypeEnumStringValues Enumerates the set of values in String for EmailPrivateEndpointPrivateEndpointTypeEnum
func GetEmailPrivateEndpointPrivateEndpointTypeEnumStringValues() []string {
	return []string{
		"SUBMISSION",
		"BOUNCE",
	}
}

// GetMappingEmailPrivateEndpointPrivateEndpointTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmailPrivateEndpointPrivateEndpointTypeEnum(val string) (EmailPrivateEndpointPrivateEndpointTypeEnum, bool) {
	enum, ok := mappingEmailPrivateEndpointPrivateEndpointTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// EmailPrivateEndpointLifecycleStateEnum Enum with underlying type: string
type EmailPrivateEndpointLifecycleStateEnum string

// Set of constants representing the allowable values for EmailPrivateEndpointLifecycleStateEnum
const (
	EmailPrivateEndpointLifecycleStateCreating EmailPrivateEndpointLifecycleStateEnum = "CREATING"
	EmailPrivateEndpointLifecycleStateUpdating EmailPrivateEndpointLifecycleStateEnum = "UPDATING"
	EmailPrivateEndpointLifecycleStateActive   EmailPrivateEndpointLifecycleStateEnum = "ACTIVE"
	EmailPrivateEndpointLifecycleStateInactive EmailPrivateEndpointLifecycleStateEnum = "INACTIVE"
	EmailPrivateEndpointLifecycleStateDeleting EmailPrivateEndpointLifecycleStateEnum = "DELETING"
	EmailPrivateEndpointLifecycleStateDeleted  EmailPrivateEndpointLifecycleStateEnum = "DELETED"
	EmailPrivateEndpointLifecycleStateFailed   EmailPrivateEndpointLifecycleStateEnum = "FAILED"
)

var mappingEmailPrivateEndpointLifecycleStateEnum = map[string]EmailPrivateEndpointLifecycleStateEnum{
	"CREATING": EmailPrivateEndpointLifecycleStateCreating,
	"UPDATING": EmailPrivateEndpointLifecycleStateUpdating,
	"ACTIVE":   EmailPrivateEndpointLifecycleStateActive,
	"INACTIVE": EmailPrivateEndpointLifecycleStateInactive,
	"DELETING": EmailPrivateEndpointLifecycleStateDeleting,
	"DELETED":  EmailPrivateEndpointLifecycleStateDeleted,
	"FAILED":   EmailPrivateEndpointLifecycleStateFailed,
}

var mappingEmailPrivateEndpointLifecycleStateEnumLowerCase = map[string]EmailPrivateEndpointLifecycleStateEnum{
	"creating": EmailPrivateEndpointLifecycleStateCreating,
	"updating": EmailPrivateEndpointLifecycleStateUpdating,
	"active":   EmailPrivateEndpointLifecycleStateActive,
	"inactive": EmailPrivateEndpointLifecycleStateInactive,
	"deleting": EmailPrivateEndpointLifecycleStateDeleting,
	"deleted":  EmailPrivateEndpointLifecycleStateDeleted,
	"failed":   EmailPrivateEndpointLifecycleStateFailed,
}

// GetEmailPrivateEndpointLifecycleStateEnumValues Enumerates the set of values for EmailPrivateEndpointLifecycleStateEnum
func GetEmailPrivateEndpointLifecycleStateEnumValues() []EmailPrivateEndpointLifecycleStateEnum {
	values := make([]EmailPrivateEndpointLifecycleStateEnum, 0)
	for _, v := range mappingEmailPrivateEndpointLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetEmailPrivateEndpointLifecycleStateEnumStringValues Enumerates the set of values in String for EmailPrivateEndpointLifecycleStateEnum
func GetEmailPrivateEndpointLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingEmailPrivateEndpointLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmailPrivateEndpointLifecycleStateEnum(val string) (EmailPrivateEndpointLifecycleStateEnum, bool) {
	enum, ok := mappingEmailPrivateEndpointLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
