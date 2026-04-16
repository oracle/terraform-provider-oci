// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogPipelineDestinationManagedKafkaResponse Response for OCI Managed Kafka destination.
// `privateEndpointMetadata` is response-only and present only for GET/LIST/cache payloads.
type LogPipelineDestinationManagedKafkaResponse struct {

	// Managed Kafka endpoint.
	ManagedKafkaEndpoint *string `mandatory:"true" json:"managedKafkaEndpoint"`

	// Topic name of the Managed Kafka.
	TopicName *string `mandatory:"true" json:"topicName"`

	DlqConfig *DlqConfig `mandatory:"true" json:"dlqConfig"`

	SecurityConfig *SecurityConfig `mandatory:"true" json:"securityConfig"`

	// Producer configuration for the Managed Kafka destination as a free-form key/value map
	// (for example serializer/deserializer settings, acks, retries).
	ProducerConfig map[string]string `mandatory:"true" json:"producerConfig"`

	// Name of Log Pipeline destination.
	Name *string `mandatory:"false" json:"name"`

	PrivateEndpointMetadata *PrivateEndpointMetadata `mandatory:"false" json:"privateEndpointMetadata"`

	// Endpoint Mode for connecting to Managed Kafka
	EndpointMode EndpointModeEnum `mandatory:"true" json:"endpointMode"`
}

// GetName returns Name
func (m LogPipelineDestinationManagedKafkaResponse) GetName() *string {
	return m.Name
}

func (m LogPipelineDestinationManagedKafkaResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogPipelineDestinationManagedKafkaResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingEndpointModeEnum(string(m.EndpointMode)); !ok && m.EndpointMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EndpointMode: %s. Supported values are: %s.", m.EndpointMode, strings.Join(GetEndpointModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m LogPipelineDestinationManagedKafkaResponse) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLogPipelineDestinationManagedKafkaResponse LogPipelineDestinationManagedKafkaResponse
	s := struct {
		DiscriminatorParam string `json:"pipelineDestinationType"`
		MarshalTypeLogPipelineDestinationManagedKafkaResponse
	}{
		"MANAGED_KAFKA",
		(MarshalTypeLogPipelineDestinationManagedKafkaResponse)(m),
	}

	return json.Marshal(&s)
}
