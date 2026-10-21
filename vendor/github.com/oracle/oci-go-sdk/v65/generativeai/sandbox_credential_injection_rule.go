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

// SandboxCredentialInjectionRule Base type for a credential injection rule. The ruleType selects the concrete
// rule schema and its required properties. Only HTTP_HEADER is currently supported.
type SandboxCredentialInjectionRule interface {
}

type sandboxcredentialinjectionrule struct {
	JsonData []byte
	RuleType string `json:"ruleType"`
}

// UnmarshalJSON unmarshals json
func (m *sandboxcredentialinjectionrule) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersandboxcredentialinjectionrule sandboxcredentialinjectionrule
	s := struct {
		Model Unmarshalersandboxcredentialinjectionrule
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.RuleType = s.Model.RuleType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *sandboxcredentialinjectionrule) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RuleType {
	case "HTTP_HEADER":
		mm := SandboxHttpHeaderCredentialInjectionRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for SandboxCredentialInjectionRule: %s.", m.RuleType)
		return *m, nil
	}
}

func (m sandboxcredentialinjectionrule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m sandboxcredentialinjectionrule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SandboxCredentialInjectionRuleRuleTypeEnum Enum with underlying type: string
type SandboxCredentialInjectionRuleRuleTypeEnum string

// Set of constants representing the allowable values for SandboxCredentialInjectionRuleRuleTypeEnum
const (
	SandboxCredentialInjectionRuleRuleTypeHttpHeader SandboxCredentialInjectionRuleRuleTypeEnum = "HTTP_HEADER"
)

var mappingSandboxCredentialInjectionRuleRuleTypeEnum = map[string]SandboxCredentialInjectionRuleRuleTypeEnum{
	"HTTP_HEADER": SandboxCredentialInjectionRuleRuleTypeHttpHeader,
}

var mappingSandboxCredentialInjectionRuleRuleTypeEnumLowerCase = map[string]SandboxCredentialInjectionRuleRuleTypeEnum{
	"http_header": SandboxCredentialInjectionRuleRuleTypeHttpHeader,
}

// GetSandboxCredentialInjectionRuleRuleTypeEnumValues Enumerates the set of values for SandboxCredentialInjectionRuleRuleTypeEnum
func GetSandboxCredentialInjectionRuleRuleTypeEnumValues() []SandboxCredentialInjectionRuleRuleTypeEnum {
	values := make([]SandboxCredentialInjectionRuleRuleTypeEnum, 0)
	for _, v := range mappingSandboxCredentialInjectionRuleRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSandboxCredentialInjectionRuleRuleTypeEnumStringValues Enumerates the set of values in String for SandboxCredentialInjectionRuleRuleTypeEnum
func GetSandboxCredentialInjectionRuleRuleTypeEnumStringValues() []string {
	return []string{
		"HTTP_HEADER",
	}
}

// GetMappingSandboxCredentialInjectionRuleRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSandboxCredentialInjectionRuleRuleTypeEnum(val string) (SandboxCredentialInjectionRuleRuleTypeEnum, bool) {
	enum, ok := mappingSandboxCredentialInjectionRuleRuleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
