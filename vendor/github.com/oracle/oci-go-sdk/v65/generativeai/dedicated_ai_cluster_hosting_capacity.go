// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Management API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service management API to create and manage DedicatedAiCluster, Endpoint, Model, and WorkRequest in the Generative AI service. For example, create a custom model by fine-tuning an out-of-the-box model using your own data, on a fine-tuning dedicated AI cluster. Then, create a hosting dedicated AI cluster with an endpoint to host your custom model.
// To access your custom model endpoints, or to try the out-of-the-box models to generate text, summarize, and create text embeddings see the Generative AI Inference API (https://docs.cloud.oracle.com/iaas/api/#/en/generative-ai-inference/latest/).
// To learn more about the service, see the Generative AI documentation (https://docs.cloud.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeai

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DedicatedAiClusterHostingCapacity The capacity of a hosting type dedicated AI cluster.
type DedicatedAiClusterHostingCapacity struct {

	// The total number of endpoints that can be hosted on this dedicated AI cluster.
	TotalEndpointCapacity *int `mandatory:"false" json:"totalEndpointCapacity"`

	// The number of endpoints hosted on this dedicated AI cluster.
	UsedEndpointCapacity *int `mandatory:"false" json:"usedEndpointCapacity"`
}

func (m DedicatedAiClusterHostingCapacity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DedicatedAiClusterHostingCapacity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DedicatedAiClusterHostingCapacity) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDedicatedAiClusterHostingCapacity DedicatedAiClusterHostingCapacity
	s := struct {
		DiscriminatorParam string `json:"capacityType"`
		MarshalTypeDedicatedAiClusterHostingCapacity
	}{
		"HOSTING_CAPACITY",
		(MarshalTypeDedicatedAiClusterHostingCapacity)(m),
	}

	return json.Marshal(&s)
}
