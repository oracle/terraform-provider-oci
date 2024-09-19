// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Management API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service management API to create and manage DedicatedAiCluster, Endpoint, Model, and WorkRequest in the Generative AI service. For example, create a custom model by fine-tuning an out-of-the-box model using your own data, on a fine-tuning dedicated AI cluster. Then, create a hosting dedicated AI cluster with an endpoint to host your custom model.
// To access your custom model endpoints, or to try the out-of-the-box models to generate text, summarize, and create text embeddings see the Generative AI Inference API (https://docs.cloud.oracle.com/iaas/api/#/en/generative-ai-inference/latest/).
// To learn more about the service, see the Generative AI documentation (https://docs.cloud.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeai

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DedicatedAiClusterCapacity The total capacity for a dedicated AI cluster.
type DedicatedAiClusterCapacity interface {
}

type dedicatedaiclustercapacity struct {
	JsonData     []byte
	CapacityType string `json:"capacityType"`
}

// UnmarshalJSON unmarshals json
func (m *dedicatedaiclustercapacity) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdedicatedaiclustercapacity dedicatedaiclustercapacity
	s := struct {
		Model Unmarshalerdedicatedaiclustercapacity
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CapacityType = s.Model.CapacityType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dedicatedaiclustercapacity) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.CapacityType {
	case "HOSTING_CAPACITY":
		mm := DedicatedAiClusterHostingCapacity{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DedicatedAiClusterCapacity: %s.", m.CapacityType)
		return *m, nil
	}
}

func (m dedicatedaiclustercapacity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dedicatedaiclustercapacity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DedicatedAiClusterCapacityCapacityTypeEnum Enum with underlying type: string
type DedicatedAiClusterCapacityCapacityTypeEnum string

// Set of constants representing the allowable values for DedicatedAiClusterCapacityCapacityTypeEnum
const (
	DedicatedAiClusterCapacityCapacityTypeHostingCapacity DedicatedAiClusterCapacityCapacityTypeEnum = "HOSTING_CAPACITY"
)

var mappingDedicatedAiClusterCapacityCapacityTypeEnum = map[string]DedicatedAiClusterCapacityCapacityTypeEnum{
	"HOSTING_CAPACITY": DedicatedAiClusterCapacityCapacityTypeHostingCapacity,
}

var mappingDedicatedAiClusterCapacityCapacityTypeEnumLowerCase = map[string]DedicatedAiClusterCapacityCapacityTypeEnum{
	"hosting_capacity": DedicatedAiClusterCapacityCapacityTypeHostingCapacity,
}

// GetDedicatedAiClusterCapacityCapacityTypeEnumValues Enumerates the set of values for DedicatedAiClusterCapacityCapacityTypeEnum
func GetDedicatedAiClusterCapacityCapacityTypeEnumValues() []DedicatedAiClusterCapacityCapacityTypeEnum {
	values := make([]DedicatedAiClusterCapacityCapacityTypeEnum, 0)
	for _, v := range mappingDedicatedAiClusterCapacityCapacityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDedicatedAiClusterCapacityCapacityTypeEnumStringValues Enumerates the set of values in String for DedicatedAiClusterCapacityCapacityTypeEnum
func GetDedicatedAiClusterCapacityCapacityTypeEnumStringValues() []string {
	return []string{
		"HOSTING_CAPACITY",
	}
}

// GetMappingDedicatedAiClusterCapacityCapacityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDedicatedAiClusterCapacityCapacityTypeEnum(val string) (DedicatedAiClusterCapacityCapacityTypeEnum, bool) {
	enum, ok := mappingDedicatedAiClusterCapacityCapacityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
