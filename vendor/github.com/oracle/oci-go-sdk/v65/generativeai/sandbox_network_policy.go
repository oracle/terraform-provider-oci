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

// SandboxNetworkPolicy Base network policy for sandboxes in a project using Oracle-managed outbound networking.
type SandboxNetworkPolicy interface {
}

type sandboxnetworkpolicy struct {
	JsonData []byte
	Mode     string `json:"mode"`
}

// UnmarshalJSON unmarshals json
func (m *sandboxnetworkpolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersandboxnetworkpolicy sandboxnetworkpolicy
	s := struct {
		Model Unmarshalersandboxnetworkpolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Mode = s.Model.Mode

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *sandboxnetworkpolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Mode {
	case "ALLOW":
		mm := SandboxAllowNetworkPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DENY_ALL":
		mm := SandboxDenyAllNetworkPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ALLOW_ALL":
		mm := SandboxAllowAllNetworkPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for SandboxNetworkPolicy: %s.", m.Mode)
		return *m, nil
	}
}

func (m sandboxnetworkpolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m sandboxnetworkpolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SandboxNetworkPolicyModeEnum Enum with underlying type: string
type SandboxNetworkPolicyModeEnum string

// Set of constants representing the allowable values for SandboxNetworkPolicyModeEnum
const (
	SandboxNetworkPolicyModeDenyAll  SandboxNetworkPolicyModeEnum = "DENY_ALL"
	SandboxNetworkPolicyModeAllowAll SandboxNetworkPolicyModeEnum = "ALLOW_ALL"
	SandboxNetworkPolicyModeAllow    SandboxNetworkPolicyModeEnum = "ALLOW"
)

var mappingSandboxNetworkPolicyModeEnum = map[string]SandboxNetworkPolicyModeEnum{
	"DENY_ALL":  SandboxNetworkPolicyModeDenyAll,
	"ALLOW_ALL": SandboxNetworkPolicyModeAllowAll,
	"ALLOW":     SandboxNetworkPolicyModeAllow,
}

var mappingSandboxNetworkPolicyModeEnumLowerCase = map[string]SandboxNetworkPolicyModeEnum{
	"deny_all":  SandboxNetworkPolicyModeDenyAll,
	"allow_all": SandboxNetworkPolicyModeAllowAll,
	"allow":     SandboxNetworkPolicyModeAllow,
}

// GetSandboxNetworkPolicyModeEnumValues Enumerates the set of values for SandboxNetworkPolicyModeEnum
func GetSandboxNetworkPolicyModeEnumValues() []SandboxNetworkPolicyModeEnum {
	values := make([]SandboxNetworkPolicyModeEnum, 0)
	for _, v := range mappingSandboxNetworkPolicyModeEnum {
		values = append(values, v)
	}
	return values
}

// GetSandboxNetworkPolicyModeEnumStringValues Enumerates the set of values in String for SandboxNetworkPolicyModeEnum
func GetSandboxNetworkPolicyModeEnumStringValues() []string {
	return []string{
		"DENY_ALL",
		"ALLOW_ALL",
		"ALLOW",
	}
}

// GetMappingSandboxNetworkPolicyModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSandboxNetworkPolicyModeEnum(val string) (SandboxNetworkPolicyModeEnum, bool) {
	enum, ok := mappingSandboxNetworkPolicyModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
