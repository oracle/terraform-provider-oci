// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Agents Management API
//
// OCI Generative AI Agents is a fully managed service that combines the power of large language models (LLMs) with an intelligent retrieval system to create contextually relevant answers by searching your knowledge base, making your AI applications smart and efficient.
// OCI Generative AI Agents supports several ways to onboard your data and then allows you and your customers to interact with your data using a chat interface or API.
// Use the Generative AI Agents API to create and manage agents, knowledge bases, data sources, endpoints, data ingestion jobs, and work requests.
// For creating and managing client chat sessions see the /EN/generative-ai-agents-client/latest/.
// To learn more about the service, see the Generative AI Agents documentation (https://docs.oracle.com/iaas/Content/generative-ai-agents/home.htm).
//

package generativeaiagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ContentModerationGuardrailConfig The configuration details about whether to apply the content moderation feature to input and output. Content moderation removes toxic and biased content from responses. It is recommended to use content moderation.
type ContentModerationGuardrailConfig struct {

	// An input guardrail mode for content moderation.
	InputGuardrailMode GuardrailModeEnum `mandatory:"false" json:"inputGuardrailMode,omitempty"`

	// An output guardrail mode for content moderation.
	OutputGuardrailMode GuardrailModeEnum `mandatory:"false" json:"outputGuardrailMode,omitempty"`
}

func (m ContentModerationGuardrailConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ContentModerationGuardrailConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingGuardrailModeEnum(string(m.InputGuardrailMode)); !ok && m.InputGuardrailMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InputGuardrailMode: %s. Supported values are: %s.", m.InputGuardrailMode, strings.Join(GetGuardrailModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGuardrailModeEnum(string(m.OutputGuardrailMode)); !ok && m.OutputGuardrailMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OutputGuardrailMode: %s. Supported values are: %s.", m.OutputGuardrailMode, strings.Join(GetGuardrailModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
