// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// HuggingFaceModel Configuration for importing a model from Hugging Face. Requires the model ID
// and a reference to the token stored in a vault for authenticated access.
type HuggingFaceModel struct {

	// The full model OCID from Hugging Face, typically in the format
	// "org/model-name" (e.g., "meta-llama/Llama-2-7b").
	ModelId *string `mandatory:"true" json:"modelId"`

	// Hugging Face access token to authenticate requests for restricted models.
	// This token will be securely stored in OCI Vault.
	AccessToken *string `mandatory:"false" json:"accessToken"`

	// The name of the branch in the Hugging Face repository to import the model from.
	// If not specified, "main" will be used by default.
	// If you provide both a branch and a commit hash, the model will be imported from the specified commit.
	Branch *string `mandatory:"false" json:"branch"`

	// The commit hash in the Hugging Face repository to import the model from.
	// If both a branch and a commit are provided, the commit hash will be used.
	Commit *string `mandatory:"false" json:"commit"`
}

func (m HuggingFaceModel) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HuggingFaceModel) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HuggingFaceModel) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHuggingFaceModel HuggingFaceModel
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeHuggingFaceModel
	}{
		"HUGGING_FACE_MODEL",
		(MarshalTypeHuggingFaceModel)(m),
	}

	return json.Marshal(&s)
}
