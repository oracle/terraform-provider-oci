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

// ApiKeyAiModelAuthDetailsSummary Summary of API key authentication details for an AI Model connection.
type ApiKeyAiModelAuthDetailsSummary struct {

	// Base URL of the AI model endpoint.
	// If not specified, the default base URL for the selected AI provider will be used.
	BaseUrl *string `mandatory:"false" json:"baseUrl"`

	// API key secret OCID for the AI model connection.
	ApiKeySecretId *string `mandatory:"false" json:"apiKeySecretId"`
}

func (m ApiKeyAiModelAuthDetailsSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApiKeyAiModelAuthDetailsSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ApiKeyAiModelAuthDetailsSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeApiKeyAiModelAuthDetailsSummary ApiKeyAiModelAuthDetailsSummary
	s := struct {
		DiscriminatorParam string `json:"authType"`
		MarshalTypeApiKeyAiModelAuthDetailsSummary
	}{
		"API_KEY",
		(MarshalTypeApiKeyAiModelAuthDetailsSummary)(m),
	}

	return json.Marshal(&s)
}
