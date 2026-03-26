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

// ScalingConfig The auto scaling configuration for the Hosted Application.
// Defines the minimum and maximum number of replicas.
// When unspecified, the service applies service-defined default scaling values.
type ScalingConfig struct {

	// scaling type for application.
	ScalingType ScalingConfigScalingTypeEnum `mandatory:"true" json:"scalingType"`

	// Minimum number of replicas to keep running.
	MinReplica *int `mandatory:"false" json:"minReplica"`

	// Maximum number of replicas allowed.
	MaxReplica *int `mandatory:"false" json:"maxReplica"`

	// Scale up if average CPU utilization exceeds this threshold.
	TargetCpuThreshold *int `mandatory:"false" json:"targetCpuThreshold"`

	// Scale up if average memory utilization exceeds this threshold.
	TargetMemoryThreshold *int `mandatory:"false" json:"targetMemoryThreshold"`

	// number of simultaneous requests that can be processed by each replica.
	TargetConcurrencyThreshold *int `mandatory:"false" json:"targetConcurrencyThreshold"`

	// requests-per-second per replica of an application.
	TargetRpsThreshold *int `mandatory:"false" json:"targetRpsThreshold"`
}

func (m ScalingConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScalingConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingScalingConfigScalingTypeEnum(string(m.ScalingType)); !ok && m.ScalingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScalingType: %s. Supported values are: %s.", m.ScalingType, strings.Join(GetScalingConfigScalingTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ScalingConfigScalingTypeEnum Enum with underlying type: string
type ScalingConfigScalingTypeEnum string

// Set of constants representing the allowable values for ScalingConfigScalingTypeEnum
const (
	ScalingConfigScalingTypeCpu               ScalingConfigScalingTypeEnum = "CPU"
	ScalingConfigScalingTypeMemory            ScalingConfigScalingTypeEnum = "MEMORY"
	ScalingConfigScalingTypeConcurrency       ScalingConfigScalingTypeEnum = "CONCURRENCY"
	ScalingConfigScalingTypeRequestsPerSecond ScalingConfigScalingTypeEnum = "REQUESTS_PER_SECOND"
)

var mappingScalingConfigScalingTypeEnum = map[string]ScalingConfigScalingTypeEnum{
	"CPU":                 ScalingConfigScalingTypeCpu,
	"MEMORY":              ScalingConfigScalingTypeMemory,
	"CONCURRENCY":         ScalingConfigScalingTypeConcurrency,
	"REQUESTS_PER_SECOND": ScalingConfigScalingTypeRequestsPerSecond,
}

var mappingScalingConfigScalingTypeEnumLowerCase = map[string]ScalingConfigScalingTypeEnum{
	"cpu":                 ScalingConfigScalingTypeCpu,
	"memory":              ScalingConfigScalingTypeMemory,
	"concurrency":         ScalingConfigScalingTypeConcurrency,
	"requests_per_second": ScalingConfigScalingTypeRequestsPerSecond,
}

// GetScalingConfigScalingTypeEnumValues Enumerates the set of values for ScalingConfigScalingTypeEnum
func GetScalingConfigScalingTypeEnumValues() []ScalingConfigScalingTypeEnum {
	values := make([]ScalingConfigScalingTypeEnum, 0)
	for _, v := range mappingScalingConfigScalingTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScalingConfigScalingTypeEnumStringValues Enumerates the set of values in String for ScalingConfigScalingTypeEnum
func GetScalingConfigScalingTypeEnumStringValues() []string {
	return []string{
		"CPU",
		"MEMORY",
		"CONCURRENCY",
		"REQUESTS_PER_SECOND",
	}
}

// GetMappingScalingConfigScalingTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScalingConfigScalingTypeEnum(val string) (ScalingConfigScalingTypeEnum, bool) {
	enum, ok := mappingScalingConfigScalingTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
