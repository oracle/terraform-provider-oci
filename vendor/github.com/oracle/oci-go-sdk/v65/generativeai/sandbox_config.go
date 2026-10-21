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

// SandboxConfig Sandbox configuration under a project.
// Use `sandboxOutboundNetworkingConfig` to choose whether sandbox outbound traffic originates from an Oracle-managed network or from a customer subnet.
// The `warmPoolConfig` field is a read-only projection of reserved MANAGED default-pool capacities.
// Project create and update requests must omit `warmPoolConfig`; supplying it, including null or an empty object, is rejected.
// Newly enabled projects start with zero default-pool capacity. Use the named warm-pool APIs to change capacities or manage CUSTOM pools.
// On Project Update, omit sandboxConfig to retain the existing configuration. When supplying
// sandboxConfig, include isEnabled. Sandbox can be enabled after project creation but cannot
// be disabled after it is enabled. After enablement, isPersistenceEnabled and custom outbound
// networking cannot change; managed network policy can be updated.
// Omitted isPersistenceEnabled retains its current value.
// Credential profiles require unique labels and can be added, removed, or changed on Project
// Update. Omit credentialProfiles to retain them, or supply the complete desired list; an
// empty array removes all profiles from the Project. When credentialProfiles is supplied,
// omitting sandboxOutboundNetworkingConfig retains the current outbound networking configuration.
// Updates apply to subsequently created
// Sandboxes after Project configuration propagation. Existing Sandboxes retain their selected
// profile snapshot, including when that profile is changed or removed from the Project.
// Updating a secret's CURRENT version in Vault does not change the Project configuration and
// remains visible to existing Sandboxes that reference that secret.
type SandboxConfig struct {

	// Indicates whether sandbox config is enabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// Indicates whether persistent block volume storage is enabled for sandbox workspaces in this project.
	// When omitted, the service-defined default is used.
	IsPersistenceEnabled *bool `mandatory:"false" json:"isPersistenceEnabled"`

	WarmPoolConfig *SandboxWarmPoolConfig `mandatory:"false" json:"warmPoolConfig"`

	SandboxOutboundNetworkingConfig SandboxOutboundNetworkingConfig `mandatory:"false" json:"sandboxOutboundNetworkingConfig"`

	// Named credential profiles available for Sandbox creation. A Sandbox can explicitly
	// select one profile using credentialProfileLabel; omitting that selector selects none.
	// On Project Update, omission retains existing profiles and a supplied list replaces
	// the complete list. An empty array removes all profiles from the Project. Adding,
	// removing, or changing profiles affects subsequently created Sandboxes after configuration
	// propagation. Existing Sandboxes retain their selected configuration snapshot.
	CredentialProfiles []SandboxCredentialProfile `mandatory:"false" json:"credentialProfiles"`
}

func (m SandboxConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SandboxConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *SandboxConfig) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		IsPersistenceEnabled            *bool                           `json:"isPersistenceEnabled"`
		WarmPoolConfig                  *SandboxWarmPoolConfig          `json:"warmPoolConfig"`
		SandboxOutboundNetworkingConfig sandboxoutboundnetworkingconfig `json:"sandboxOutboundNetworkingConfig"`
		CredentialProfiles              []SandboxCredentialProfile      `json:"credentialProfiles"`
		IsEnabled                       *bool                           `json:"isEnabled"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.IsPersistenceEnabled = model.IsPersistenceEnabled

	m.WarmPoolConfig = model.WarmPoolConfig

	nn, e = model.SandboxOutboundNetworkingConfig.UnmarshalPolymorphicJSON(model.SandboxOutboundNetworkingConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SandboxOutboundNetworkingConfig = nn.(SandboxOutboundNetworkingConfig)
	} else {
		m.SandboxOutboundNetworkingConfig = nil
	}

	m.CredentialProfiles = make([]SandboxCredentialProfile, len(model.CredentialProfiles))
	copy(m.CredentialProfiles, model.CredentialProfiles)
	m.IsEnabled = model.IsEnabled

	return
}
