// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Streaming with Apache Kafka (OSAK) API
//
// Use Oracle Streaming with Apache Kafka Control Plane API to create/update/delete managed Kafka clusters.
//

package managedkafka

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KafkaClusterConfig A shared configuration object used by 0 or more kafka clusters.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type KafkaClusterConfig struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KafkaClusterConfig.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the KafkaClusterConfig was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the KafkaClusterConfig.
	LifecycleState KafkaClusterConfigLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	LatestConfig *KafkaClusterConfigVersion `mandatory:"true" json:"latestConfig"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The date and time the KafkaClusterConfig was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message that describes the current state of the KafkaClusterConfig in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m KafkaClusterConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KafkaClusterConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingKafkaClusterConfigLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetKafkaClusterConfigLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// KafkaClusterConfigLifecycleStateEnum Enum with underlying type: string
type KafkaClusterConfigLifecycleStateEnum string

// Set of constants representing the allowable values for KafkaClusterConfigLifecycleStateEnum
const (
	KafkaClusterConfigLifecycleStateCreating KafkaClusterConfigLifecycleStateEnum = "CREATING"
	KafkaClusterConfigLifecycleStateActive   KafkaClusterConfigLifecycleStateEnum = "ACTIVE"
	KafkaClusterConfigLifecycleStateUpdating KafkaClusterConfigLifecycleStateEnum = "UPDATING"
	KafkaClusterConfigLifecycleStateDeleted  KafkaClusterConfigLifecycleStateEnum = "DELETED"
)

var mappingKafkaClusterConfigLifecycleStateEnum = map[string]KafkaClusterConfigLifecycleStateEnum{
	"CREATING": KafkaClusterConfigLifecycleStateCreating,
	"ACTIVE":   KafkaClusterConfigLifecycleStateActive,
	"UPDATING": KafkaClusterConfigLifecycleStateUpdating,
	"DELETED":  KafkaClusterConfigLifecycleStateDeleted,
}

var mappingKafkaClusterConfigLifecycleStateEnumLowerCase = map[string]KafkaClusterConfigLifecycleStateEnum{
	"creating": KafkaClusterConfigLifecycleStateCreating,
	"active":   KafkaClusterConfigLifecycleStateActive,
	"updating": KafkaClusterConfigLifecycleStateUpdating,
	"deleted":  KafkaClusterConfigLifecycleStateDeleted,
}

// GetKafkaClusterConfigLifecycleStateEnumValues Enumerates the set of values for KafkaClusterConfigLifecycleStateEnum
func GetKafkaClusterConfigLifecycleStateEnumValues() []KafkaClusterConfigLifecycleStateEnum {
	values := make([]KafkaClusterConfigLifecycleStateEnum, 0)
	for _, v := range mappingKafkaClusterConfigLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetKafkaClusterConfigLifecycleStateEnumStringValues Enumerates the set of values in String for KafkaClusterConfigLifecycleStateEnum
func GetKafkaClusterConfigLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETED",
	}
}

// GetMappingKafkaClusterConfigLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKafkaClusterConfigLifecycleStateEnum(val string) (KafkaClusterConfigLifecycleStateEnum, bool) {
	enum, ok := mappingKafkaClusterConfigLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
