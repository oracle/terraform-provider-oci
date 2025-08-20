// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle API Access Control
//
// This service is used to restrict the control plane service apis; so that everybody won't be
// able to access those apis.
// There are two main resouces defined as a part of this service
// 1. PrivilegedApiControl: This is created by the customer which defines which service apis are
//    controlled and who can access it.
// 2. PrivilegedApiRequest: This is a request object again created by the customer operators who           seek access to those privileged apis. After a request is obtained based on the                       PrivilegedAccessControl for which the api belongs to, either it can be approved so that the          requested person can execute the service apis or it will wait for the customer to approve it.
//

package apiaccesscontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ApiMetadata An ApiDetail contains details such as the service it belongs to, the name of the api, the type of api, and the parameters of the api if it contains.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type ApiMetadata struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ApiDetail.
	Id *string `mandatory:"true" json:"id"`

	// The operation Name of the api. The name must be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time the PrivilegedApiControl was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the ApiMetadata.
	LifecycleState ApiMetadataLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The service Name to which the api belongs to.
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// rest path of the api.
	Path *string `mandatory:"false" json:"path"`

	// ResourceType to which the apiMetadata belongs to.
	EntityType *string `mandatory:"false" json:"entityType"`

	// The name of the api to execute the api request.
	ApiName *string `mandatory:"false" json:"apiName"`

	// List of the fields that is use while calling post or put for the data.
	Fields []string `mandatory:"false" json:"fields"`

	// The date and time the PrivilegedApiControl was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The date and time the PrivilegedApiControl was marked for delete, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeDeleted *common.SDKTime `mandatory:"false" json:"timeDeleted"`

	// A message that describes the current state of the ApiMetadata in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ApiMetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApiMetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingApiMetadataLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetApiMetadataLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApiMetadataLifecycleStateEnum Enum with underlying type: string
type ApiMetadataLifecycleStateEnum string

// Set of constants representing the allowable values for ApiMetadataLifecycleStateEnum
const (
	ApiMetadataLifecycleStateCreating ApiMetadataLifecycleStateEnum = "CREATING"
	ApiMetadataLifecycleStateUpdating ApiMetadataLifecycleStateEnum = "UPDATING"
	ApiMetadataLifecycleStateActive   ApiMetadataLifecycleStateEnum = "ACTIVE"
	ApiMetadataLifecycleStateDeleting ApiMetadataLifecycleStateEnum = "DELETING"
	ApiMetadataLifecycleStateDeleted  ApiMetadataLifecycleStateEnum = "DELETED"
	ApiMetadataLifecycleStateFailed   ApiMetadataLifecycleStateEnum = "FAILED"
)

var mappingApiMetadataLifecycleStateEnum = map[string]ApiMetadataLifecycleStateEnum{
	"CREATING": ApiMetadataLifecycleStateCreating,
	"UPDATING": ApiMetadataLifecycleStateUpdating,
	"ACTIVE":   ApiMetadataLifecycleStateActive,
	"DELETING": ApiMetadataLifecycleStateDeleting,
	"DELETED":  ApiMetadataLifecycleStateDeleted,
	"FAILED":   ApiMetadataLifecycleStateFailed,
}

var mappingApiMetadataLifecycleStateEnumLowerCase = map[string]ApiMetadataLifecycleStateEnum{
	"creating": ApiMetadataLifecycleStateCreating,
	"updating": ApiMetadataLifecycleStateUpdating,
	"active":   ApiMetadataLifecycleStateActive,
	"deleting": ApiMetadataLifecycleStateDeleting,
	"deleted":  ApiMetadataLifecycleStateDeleted,
	"failed":   ApiMetadataLifecycleStateFailed,
}

// GetApiMetadataLifecycleStateEnumValues Enumerates the set of values for ApiMetadataLifecycleStateEnum
func GetApiMetadataLifecycleStateEnumValues() []ApiMetadataLifecycleStateEnum {
	values := make([]ApiMetadataLifecycleStateEnum, 0)
	for _, v := range mappingApiMetadataLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetApiMetadataLifecycleStateEnumStringValues Enumerates the set of values in String for ApiMetadataLifecycleStateEnum
func GetApiMetadataLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingApiMetadataLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApiMetadataLifecycleStateEnum(val string) (ApiMetadataLifecycleStateEnum, bool) {
	enum, ok := mappingApiMetadataLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
