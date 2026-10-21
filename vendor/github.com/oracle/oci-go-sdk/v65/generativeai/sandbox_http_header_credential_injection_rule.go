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

// SandboxHttpHeaderCredentialInjectionRule An HTTP_HEADER rule that binds one exact HTTPS host on port 443 to exactly one
// injected HTTP request header and one OCI Vault secret. Redirects to another host do not
// carry the credential. Only properties defined for this rule type are supported.
type SandboxHttpHeaderCredentialInjectionRule struct {

	// The normalized ASCII DNS hostname that selects the rule. This value must be the exact
	// regional Generative AI Responses endpoint hostname for the Project's region and realm.
	// Do not include a scheme, path, port, wildcard, trailing dot, or IP literal.
	// Brokerage uses HTTPS on port 443.
	Host *string `mandatory:"true" json:"host"`

	// The single HTTP request header that the service replaces with the credential value.
	InjectHeaders []SandboxCredentialInjectionHeader `mandatory:"true" json:"injectHeaders"`
}

func (m SandboxHttpHeaderCredentialInjectionRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SandboxHttpHeaderCredentialInjectionRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SandboxHttpHeaderCredentialInjectionRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSandboxHttpHeaderCredentialInjectionRule SandboxHttpHeaderCredentialInjectionRule
	s := struct {
		DiscriminatorParam string `json:"ruleType"`
		MarshalTypeSandboxHttpHeaderCredentialInjectionRule
	}{
		"HTTP_HEADER",
		(MarshalTypeSandboxHttpHeaderCredentialInjectionRule)(m),
	}

	return json.Marshal(&s)
}
