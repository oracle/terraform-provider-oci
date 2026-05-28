// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateApiKeyAiModelAuthDetails The information to update API key authentication details for an AI Model connection.
type UpdateApiKeyAiModelAuthDetails struct {

	// Base URL of the AI model endpoint.
	// If not specified, the default base URL for the selected AI provider will be used.
	BaseUrl *string `mandatory:"false" json:"baseUrl"`

	// API key for the AI model connection.
	// Deprecated: This field is deprecated and replaced by "apiKeySecretId".
	// This change follows the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	ApiKey *string `mandatory:"false" json:"apiKey"`

	// API key secret OCID for the AI model connection.
	ApiKeySecretId *string `mandatory:"false" json:"apiKeySecretId"`
}

func (m UpdateApiKeyAiModelAuthDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateApiKeyAiModelAuthDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateApiKeyAiModelAuthDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateApiKeyAiModelAuthDetails UpdateApiKeyAiModelAuthDetails
	s := struct {
		DiscriminatorParam string `json:"authType"`
		MarshalTypeUpdateApiKeyAiModelAuthDetails
	}{
		"API_KEY",
		(MarshalTypeUpdateApiKeyAiModelAuthDetails)(m),
	}

	return json.Marshal(&s)
}
