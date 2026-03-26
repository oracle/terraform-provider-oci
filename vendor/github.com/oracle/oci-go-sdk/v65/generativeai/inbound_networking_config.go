// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Management API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service management API to create and manage DedicatedAiCluster, Endpoint, Model, and WorkRequest in the Generative AI service. For example, create a custom model by fine-tuning an out-of-the-box model using your own data, on a fine-tuning dedicated AI cluster. Then, create a hosting dedicated AI cluster with an endpoint to host your custom model.
// To access your custom model endpoints, or to try the out-of-the-box models to generate text, summarize, and create text embeddings see the Generative AI Inference API (https://docs.oracle.com/iaas/api/#/en/generative-ai-inference/latest/).
// To learn more about the service, see the Generative AI documentation (https://docs.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeai

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InboundNetworkingConfig Inbound Networking configuration.
type InboundNetworkingConfig struct {

	// inbounding from public or private endpoint.
	EndpointMode InboundNetworkingConfigEndpointModeEnum `mandatory:"true" json:"endpointMode"`

	// The [OCID] of Private Endpoint when endpointMode=Private
	PrivateEndpointId *string `mandatory:"false" json:"privateEndpointId"`
}

func (m InboundNetworkingConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InboundNetworkingConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInboundNetworkingConfigEndpointModeEnum(string(m.EndpointMode)); !ok && m.EndpointMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EndpointMode: %s. Supported values are: %s.", m.EndpointMode, strings.Join(GetInboundNetworkingConfigEndpointModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InboundNetworkingConfigEndpointModeEnum Enum with underlying type: string
type InboundNetworkingConfigEndpointModeEnum string

// Set of constants representing the allowable values for InboundNetworkingConfigEndpointModeEnum
const (
	InboundNetworkingConfigEndpointModePublic  InboundNetworkingConfigEndpointModeEnum = "PUBLIC"
	InboundNetworkingConfigEndpointModePrivate InboundNetworkingConfigEndpointModeEnum = "PRIVATE"
)

var mappingInboundNetworkingConfigEndpointModeEnum = map[string]InboundNetworkingConfigEndpointModeEnum{
	"PUBLIC":  InboundNetworkingConfigEndpointModePublic,
	"PRIVATE": InboundNetworkingConfigEndpointModePrivate,
}

var mappingInboundNetworkingConfigEndpointModeEnumLowerCase = map[string]InboundNetworkingConfigEndpointModeEnum{
	"public":  InboundNetworkingConfigEndpointModePublic,
	"private": InboundNetworkingConfigEndpointModePrivate,
}

// GetInboundNetworkingConfigEndpointModeEnumValues Enumerates the set of values for InboundNetworkingConfigEndpointModeEnum
func GetInboundNetworkingConfigEndpointModeEnumValues() []InboundNetworkingConfigEndpointModeEnum {
	values := make([]InboundNetworkingConfigEndpointModeEnum, 0)
	for _, v := range mappingInboundNetworkingConfigEndpointModeEnum {
		values = append(values, v)
	}
	return values
}

// GetInboundNetworkingConfigEndpointModeEnumStringValues Enumerates the set of values in String for InboundNetworkingConfigEndpointModeEnum
func GetInboundNetworkingConfigEndpointModeEnumStringValues() []string {
	return []string{
		"PUBLIC",
		"PRIVATE",
	}
}

// GetMappingInboundNetworkingConfigEndpointModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInboundNetworkingConfigEndpointModeEnum(val string) (InboundNetworkingConfigEndpointModeEnum, bool) {
	enum, ok := mappingInboundNetworkingConfigEndpointModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
