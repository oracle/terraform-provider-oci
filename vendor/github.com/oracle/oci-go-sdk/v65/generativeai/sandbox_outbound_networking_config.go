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

// SandboxOutboundNetworkingConfig Base outbound networking configuration for sandboxes in a project.
type SandboxOutboundNetworkingConfig interface {
}

type sandboxoutboundnetworkingconfig struct {
	JsonData    []byte
	NetworkMode string `json:"networkMode"`
}

// UnmarshalJSON unmarshals json
func (m *sandboxoutboundnetworkingconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersandboxoutboundnetworkingconfig sandboxoutboundnetworkingconfig
	s := struct {
		Model Unmarshalersandboxoutboundnetworkingconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.NetworkMode = s.Model.NetworkMode

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *sandboxoutboundnetworkingconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.NetworkMode {
	case "MANAGED":
		mm := ManagedSandboxOutboundNetworkingConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOM":
		mm := CustomSandboxOutboundNetworkingConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for SandboxOutboundNetworkingConfig: %s.", m.NetworkMode)
		return *m, nil
	}
}

func (m sandboxoutboundnetworkingconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m sandboxoutboundnetworkingconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SandboxOutboundNetworkingConfigNetworkModeEnum Enum with underlying type: string
type SandboxOutboundNetworkingConfigNetworkModeEnum string

// Set of constants representing the allowable values for SandboxOutboundNetworkingConfigNetworkModeEnum
const (
	SandboxOutboundNetworkingConfigNetworkModeManaged SandboxOutboundNetworkingConfigNetworkModeEnum = "MANAGED"
	SandboxOutboundNetworkingConfigNetworkModeCustom  SandboxOutboundNetworkingConfigNetworkModeEnum = "CUSTOM"
)

var mappingSandboxOutboundNetworkingConfigNetworkModeEnum = map[string]SandboxOutboundNetworkingConfigNetworkModeEnum{
	"MANAGED": SandboxOutboundNetworkingConfigNetworkModeManaged,
	"CUSTOM":  SandboxOutboundNetworkingConfigNetworkModeCustom,
}

var mappingSandboxOutboundNetworkingConfigNetworkModeEnumLowerCase = map[string]SandboxOutboundNetworkingConfigNetworkModeEnum{
	"managed": SandboxOutboundNetworkingConfigNetworkModeManaged,
	"custom":  SandboxOutboundNetworkingConfigNetworkModeCustom,
}

// GetSandboxOutboundNetworkingConfigNetworkModeEnumValues Enumerates the set of values for SandboxOutboundNetworkingConfigNetworkModeEnum
func GetSandboxOutboundNetworkingConfigNetworkModeEnumValues() []SandboxOutboundNetworkingConfigNetworkModeEnum {
	values := make([]SandboxOutboundNetworkingConfigNetworkModeEnum, 0)
	for _, v := range mappingSandboxOutboundNetworkingConfigNetworkModeEnum {
		values = append(values, v)
	}
	return values
}

// GetSandboxOutboundNetworkingConfigNetworkModeEnumStringValues Enumerates the set of values in String for SandboxOutboundNetworkingConfigNetworkModeEnum
func GetSandboxOutboundNetworkingConfigNetworkModeEnumStringValues() []string {
	return []string{
		"MANAGED",
		"CUSTOM",
	}
}

// GetMappingSandboxOutboundNetworkingConfigNetworkModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSandboxOutboundNetworkingConfigNetworkModeEnum(val string) (SandboxOutboundNetworkingConfigNetworkModeEnum, bool) {
	enum, ok := mappingSandboxOutboundNetworkingConfigNetworkModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
