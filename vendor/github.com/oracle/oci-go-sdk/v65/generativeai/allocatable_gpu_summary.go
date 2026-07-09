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

// AllocatableGpuSummary Information about allocatable GPUs that can be used to launch dedicated AI clusters.
type AllocatableGpuSummary struct {

	// The compute shape for the allocatable GPUs.
	ComputeShape *string `mandatory:"true" json:"computeShape"`

	// The total number of available GPU units for the compute shape.
	TotalAvailableGpuUnits *int `mandatory:"true" json:"totalAvailableGpuUnits"`

	// The total number of used GPU units for the compute shape.
	TotalUsedGpuUnits *int `mandatory:"true" json:"totalUsedGpuUnits"`

	// The maximum number of GPU units that can be allocated to a single dedicated AI cluster for the compute shape without requiring additional provisioning.
	MaxAllocatableGpuUnits *int `mandatory:"true" json:"maxAllocatableGpuUnits"`

	// The GPU allocation details for the compute shape. Each item represents one underlying capacity allocation.
	GpuAllocations []GpuAllocation `mandatory:"true" json:"gpuAllocations"`
}

func (m AllocatableGpuSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AllocatableGpuSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
