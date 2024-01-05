// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Streaming API
//
// Use the Streaming API to produce and consume messages, create streams and stream pools, and manage related items. For more information, see Streaming (https://docs.cloud.oracle.com/Content/Streaming/Concepts/streamingoverview.htm).
//

package streaming

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StreamPool The details of a stream pool.
type StreamPool struct {

	// The OCID of the stream pool.
	Id *string `mandatory:"true" json:"id"`

	// Compartment OCID that the pool belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the stream pool.
	Name *string `mandatory:"true" json:"name"`

	// The current state of the stream pool.
	LifecycleState StreamPoolLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the stream pool was created, expressed in in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2018-04-20T00:00:07.405Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	KafkaSettings *KafkaSettings `mandatory:"true" json:"kafkaSettings"`

	CustomEncryptionKey *CustomEncryptionKey `mandatory:"true" json:"customEncryptionKey"`

	// Any additional details about the current state of the stream.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// True if the stream pool is private, false otherwise.
	// If the stream pool is private, the streams inside the stream pool can only be accessed from inside the associated subnetId.
	IsPrivate *bool `mandatory:"false" json:"isPrivate"`

	// The FQDN used to access the streams inside the stream pool (same FQDN as the messagesEndpoint attribute of a Stream object).
	// If the stream pool is private, the FQDN is customized and can only be accessed from inside the associated subnetId, otherwise the FQDN is publicly resolvable.
	// Depending on which protocol you attempt to use, you need to either prepend https or append the Kafka port.
	EndpointFqdn *string `mandatory:"false" json:"endpointFqdn"`

	PrivateEndpointSettings *PrivateEndpointSettings `mandatory:"false" json:"privateEndpointSettings"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. Exists for cross-compatibility only.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}'
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m StreamPool) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StreamPool) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStreamPoolLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetStreamPoolLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StreamPoolLifecycleStateEnum Enum with underlying type: string
type StreamPoolLifecycleStateEnum string

// Set of constants representing the allowable values for StreamPoolLifecycleStateEnum
const (
	StreamPoolLifecycleStateCreating StreamPoolLifecycleStateEnum = "CREATING"
	StreamPoolLifecycleStateActive   StreamPoolLifecycleStateEnum = "ACTIVE"
	StreamPoolLifecycleStateDeleting StreamPoolLifecycleStateEnum = "DELETING"
	StreamPoolLifecycleStateDeleted  StreamPoolLifecycleStateEnum = "DELETED"
	StreamPoolLifecycleStateFailed   StreamPoolLifecycleStateEnum = "FAILED"
	StreamPoolLifecycleStateUpdating StreamPoolLifecycleStateEnum = "UPDATING"
)

var mappingStreamPoolLifecycleStateEnum = map[string]StreamPoolLifecycleStateEnum{
	"CREATING": StreamPoolLifecycleStateCreating,
	"ACTIVE":   StreamPoolLifecycleStateActive,
	"DELETING": StreamPoolLifecycleStateDeleting,
	"DELETED":  StreamPoolLifecycleStateDeleted,
	"FAILED":   StreamPoolLifecycleStateFailed,
	"UPDATING": StreamPoolLifecycleStateUpdating,
}

var mappingStreamPoolLifecycleStateEnumLowerCase = map[string]StreamPoolLifecycleStateEnum{
	"creating": StreamPoolLifecycleStateCreating,
	"active":   StreamPoolLifecycleStateActive,
	"deleting": StreamPoolLifecycleStateDeleting,
	"deleted":  StreamPoolLifecycleStateDeleted,
	"failed":   StreamPoolLifecycleStateFailed,
	"updating": StreamPoolLifecycleStateUpdating,
}

// GetStreamPoolLifecycleStateEnumValues Enumerates the set of values for StreamPoolLifecycleStateEnum
func GetStreamPoolLifecycleStateEnumValues() []StreamPoolLifecycleStateEnum {
	values := make([]StreamPoolLifecycleStateEnum, 0)
	for _, v := range mappingStreamPoolLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetStreamPoolLifecycleStateEnumStringValues Enumerates the set of values in String for StreamPoolLifecycleStateEnum
func GetStreamPoolLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingStreamPoolLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStreamPoolLifecycleStateEnum(val string) (StreamPoolLifecycleStateEnum, bool) {
	enum, ok := mappingStreamPoolLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
