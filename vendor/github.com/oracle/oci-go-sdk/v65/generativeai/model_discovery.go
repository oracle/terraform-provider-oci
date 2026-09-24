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

// ModelDiscovery Metadata describing a specific model
type ModelDiscovery struct {

	// A unique identifier for the model
	ModelId *string `mandatory:"true" json:"modelId"`

	// The vendor that offers the model.
	Vendor *string `mandatory:"true" json:"vendor"`

	// Describes what this model can be used for.
	Capabilities []ModelCapabilityEnum `mandatory:"true" json:"capabilities"`

	// The specific API endpoints that this model supports. For example, a chat model may support OPENAI_V1_CHAT_COMPLETIONS and/or OPENAI_V1_RESPONSES. If empty, the model has not yet been annotated with API capabilities.
	ApiCapability []string `mandatory:"false" json:"apiCapability"`

	// The supported input-to-output modality transformations for this model. For example, a model can support
	// TEXT to VIDEO or AUDIO to VIDEO.
	ModalitySupport []ModelModalitySupport `mandatory:"false" json:"modalitySupport"`

	// The access level required to use the model, which can be either HOSTED (the model is hosted by the provider and can be accessed via API calls) or PROXY (the model is not directly accessible and requires going through a proxy).
	ModelAccess ModelDiscoveryModelAccessEnum `mandatory:"false" json:"modelAccess,omitempty"`

	// The list of configurable parameters supported by the model. For example, temperature and max_tokens for a text generation model.
	Parameters []Parameter `mandatory:"false" json:"parameters"`

	// The list of availability details for the model across different regions and deployment modes.
	Availability []Availability `mandatory:"false" json:"availability"`
}

func (m ModelDiscovery) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ModelDiscovery) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingModelDiscoveryModelAccessEnum(string(m.ModelAccess)); !ok && m.ModelAccess != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelAccess: %s. Supported values are: %s.", m.ModelAccess, strings.Join(GetModelDiscoveryModelAccessEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ModelDiscoveryModelAccessEnum Enum with underlying type: string
type ModelDiscoveryModelAccessEnum string

// Set of constants representing the allowable values for ModelDiscoveryModelAccessEnum
const (
	ModelDiscoveryModelAccessHosted ModelDiscoveryModelAccessEnum = "HOSTED"
	ModelDiscoveryModelAccessProxy  ModelDiscoveryModelAccessEnum = "PROXY"
)

var mappingModelDiscoveryModelAccessEnum = map[string]ModelDiscoveryModelAccessEnum{
	"HOSTED": ModelDiscoveryModelAccessHosted,
	"PROXY":  ModelDiscoveryModelAccessProxy,
}

var mappingModelDiscoveryModelAccessEnumLowerCase = map[string]ModelDiscoveryModelAccessEnum{
	"hosted": ModelDiscoveryModelAccessHosted,
	"proxy":  ModelDiscoveryModelAccessProxy,
}

// GetModelDiscoveryModelAccessEnumValues Enumerates the set of values for ModelDiscoveryModelAccessEnum
func GetModelDiscoveryModelAccessEnumValues() []ModelDiscoveryModelAccessEnum {
	values := make([]ModelDiscoveryModelAccessEnum, 0)
	for _, v := range mappingModelDiscoveryModelAccessEnum {
		values = append(values, v)
	}
	return values
}

// GetModelDiscoveryModelAccessEnumStringValues Enumerates the set of values in String for ModelDiscoveryModelAccessEnum
func GetModelDiscoveryModelAccessEnumStringValues() []string {
	return []string{
		"HOSTED",
		"PROXY",
	}
}

// GetMappingModelDiscoveryModelAccessEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelDiscoveryModelAccessEnum(val string) (ModelDiscoveryModelAccessEnum, bool) {
	enum, ok := mappingModelDiscoveryModelAccessEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
