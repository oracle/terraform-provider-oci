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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ApiKeyItem The ApiKey item.
type ApiKeyItem struct {

	// The key name.
	KeyName *string `mandatory:"true" json:"keyName"`

	// The masked key.
	KeyMask *string `mandatory:"true" json:"keyMask"`

	// The date and time that the key was created in the format of an RFC3339 datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time when the key would be expired, if not provided it would be 90 days, in the format defined by RFC 3339.
	TimeExpiry *common.SDKTime `mandatory:"true" json:"timeExpiry"`

	// The current state of the API key item.
	State ApiKeyItemStateEnum `mandatory:"true" json:"state"`

	// The key.
	Key *string `mandatory:"false" json:"key"`

	// The date and time that the key is activated in the format of an RFC3339 datetime string.
	TimeActivated *common.SDKTime `mandatory:"false" json:"timeActivated"`

	// The date and time that the key is deactivated in the format of an RFC3339 datetime string.
	TimeDeactivated *common.SDKTime `mandatory:"false" json:"timeDeactivated"`

	// The date and time that the key is revoked in the format of an RFC3339 datetime string.
	TimeRevoked *common.SDKTime `mandatory:"false" json:"timeRevoked"`

	// The date and time that the key is last used in the format of an RFC3339 datetime string.
	TimeLastUsed *common.SDKTime `mandatory:"false" json:"timeLastUsed"`
}

func (m ApiKeyItem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApiKeyItem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingApiKeyItemStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetApiKeyItemStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApiKeyItemStateEnum Enum with underlying type: string
type ApiKeyItemStateEnum string

// Set of constants representing the allowable values for ApiKeyItemStateEnum
const (
	ApiKeyItemStateActive   ApiKeyItemStateEnum = "ACTIVE"
	ApiKeyItemStateInactive ApiKeyItemStateEnum = "INACTIVE"
	ApiKeyItemStateRevoked  ApiKeyItemStateEnum = "REVOKED"
	ApiKeyItemStateExpired  ApiKeyItemStateEnum = "EXPIRED"
	ApiKeyItemStateDeleted  ApiKeyItemStateEnum = "DELETED"
)

var mappingApiKeyItemStateEnum = map[string]ApiKeyItemStateEnum{
	"ACTIVE":   ApiKeyItemStateActive,
	"INACTIVE": ApiKeyItemStateInactive,
	"REVOKED":  ApiKeyItemStateRevoked,
	"EXPIRED":  ApiKeyItemStateExpired,
	"DELETED":  ApiKeyItemStateDeleted,
}

var mappingApiKeyItemStateEnumLowerCase = map[string]ApiKeyItemStateEnum{
	"active":   ApiKeyItemStateActive,
	"inactive": ApiKeyItemStateInactive,
	"revoked":  ApiKeyItemStateRevoked,
	"expired":  ApiKeyItemStateExpired,
	"deleted":  ApiKeyItemStateDeleted,
}

// GetApiKeyItemStateEnumValues Enumerates the set of values for ApiKeyItemStateEnum
func GetApiKeyItemStateEnumValues() []ApiKeyItemStateEnum {
	values := make([]ApiKeyItemStateEnum, 0)
	for _, v := range mappingApiKeyItemStateEnum {
		values = append(values, v)
	}
	return values
}

// GetApiKeyItemStateEnumStringValues Enumerates the set of values in String for ApiKeyItemStateEnum
func GetApiKeyItemStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"REVOKED",
		"EXPIRED",
		"DELETED",
	}
}

// GetMappingApiKeyItemStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApiKeyItemStateEnum(val string) (ApiKeyItemStateEnum, bool) {
	enum, ok := mappingApiKeyItemStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
