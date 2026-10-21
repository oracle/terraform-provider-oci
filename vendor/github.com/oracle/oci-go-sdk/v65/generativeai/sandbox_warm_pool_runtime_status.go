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

// SandboxWarmPoolRuntimeStatus Best-effort runtime observations returned by GetSandboxWarmPool. Runtime read failure does not fail an otherwise successful configuration read or change the pool lifecycle.
type SandboxWarmPoolRuntimeStatus struct {

	// OBSERVED means the runtime status was read. UNAVAILABLE means no usable observation is available; configuration and management lifecycle are still returned.
	ObservationStatus SandboxWarmPoolRuntimeStatusObservationStatusEnum `mandatory:"true" json:"observationStatus"`

	// Controller-reported ready idle Sandbox members, excluding adopted customer Sandboxes. Omitted when unavailable; a known zero is returned as zero. This eventually consistent count is not a reservation or a guarantee of current allocation readiness.
	ReadyCapacity *int `mandatory:"false" json:"readyCapacity"`

	// Time the runtime status was fetched, not the controller's last reconciliation time. A successful Kubernetes read can still contain an older controller observation.
	TimeRead *common.SDKTime `mandatory:"false" json:"timeRead"`

	// Concise explanation when runtime observations are unavailable.
	Reason *string `mandatory:"false" json:"reason"`
}

func (m SandboxWarmPoolRuntimeStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SandboxWarmPoolRuntimeStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSandboxWarmPoolRuntimeStatusObservationStatusEnum(string(m.ObservationStatus)); !ok && m.ObservationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObservationStatus: %s. Supported values are: %s.", m.ObservationStatus, strings.Join(GetSandboxWarmPoolRuntimeStatusObservationStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SandboxWarmPoolRuntimeStatusObservationStatusEnum Enum with underlying type: string
type SandboxWarmPoolRuntimeStatusObservationStatusEnum string

// Set of constants representing the allowable values for SandboxWarmPoolRuntimeStatusObservationStatusEnum
const (
	SandboxWarmPoolRuntimeStatusObservationStatusObserved    SandboxWarmPoolRuntimeStatusObservationStatusEnum = "OBSERVED"
	SandboxWarmPoolRuntimeStatusObservationStatusUnavailable SandboxWarmPoolRuntimeStatusObservationStatusEnum = "UNAVAILABLE"
)

var mappingSandboxWarmPoolRuntimeStatusObservationStatusEnum = map[string]SandboxWarmPoolRuntimeStatusObservationStatusEnum{
	"OBSERVED":    SandboxWarmPoolRuntimeStatusObservationStatusObserved,
	"UNAVAILABLE": SandboxWarmPoolRuntimeStatusObservationStatusUnavailable,
}

var mappingSandboxWarmPoolRuntimeStatusObservationStatusEnumLowerCase = map[string]SandboxWarmPoolRuntimeStatusObservationStatusEnum{
	"observed":    SandboxWarmPoolRuntimeStatusObservationStatusObserved,
	"unavailable": SandboxWarmPoolRuntimeStatusObservationStatusUnavailable,
}

// GetSandboxWarmPoolRuntimeStatusObservationStatusEnumValues Enumerates the set of values for SandboxWarmPoolRuntimeStatusObservationStatusEnum
func GetSandboxWarmPoolRuntimeStatusObservationStatusEnumValues() []SandboxWarmPoolRuntimeStatusObservationStatusEnum {
	values := make([]SandboxWarmPoolRuntimeStatusObservationStatusEnum, 0)
	for _, v := range mappingSandboxWarmPoolRuntimeStatusObservationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetSandboxWarmPoolRuntimeStatusObservationStatusEnumStringValues Enumerates the set of values in String for SandboxWarmPoolRuntimeStatusObservationStatusEnum
func GetSandboxWarmPoolRuntimeStatusObservationStatusEnumStringValues() []string {
	return []string{
		"OBSERVED",
		"UNAVAILABLE",
	}
}

// GetMappingSandboxWarmPoolRuntimeStatusObservationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSandboxWarmPoolRuntimeStatusObservationStatusEnum(val string) (SandboxWarmPoolRuntimeStatusObservationStatusEnum, bool) {
	enum, ok := mappingSandboxWarmPoolRuntimeStatusObservationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
