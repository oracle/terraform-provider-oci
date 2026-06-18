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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ManagedSandboxOutboundNetworkingConfig Outbound networking configuration that uses Oracle-managed networking for sandboxes in this project.
type ManagedSandboxOutboundNetworkingConfig struct {
	NetworkPolicy SandboxNetworkPolicy `mandatory:"false" json:"networkPolicy"`
}

func (m ManagedSandboxOutboundNetworkingConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedSandboxOutboundNetworkingConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ManagedSandboxOutboundNetworkingConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeManagedSandboxOutboundNetworkingConfig ManagedSandboxOutboundNetworkingConfig
	s := struct {
		DiscriminatorParam string `json:"networkMode"`
		MarshalTypeManagedSandboxOutboundNetworkingConfig
	}{
		"MANAGED",
		(MarshalTypeManagedSandboxOutboundNetworkingConfig)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ManagedSandboxOutboundNetworkingConfig) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		NetworkPolicy sandboxnetworkpolicy `json:"networkPolicy"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.NetworkPolicy.UnmarshalPolymorphicJSON(model.NetworkPolicy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.NetworkPolicy = nn.(SandboxNetworkPolicy)
	} else {
		m.NetworkPolicy = nil
	}

	return
}
