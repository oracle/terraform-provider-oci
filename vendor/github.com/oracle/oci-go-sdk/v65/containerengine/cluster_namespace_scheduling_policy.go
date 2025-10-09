// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Kubernetes Engine API
//
// API for the Kubernetes Engine service (also known as the Container Engine for Kubernetes service). Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Kubernetes Engine (https://docs.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ClusterNamespaceSchedulingPolicy Cluster Namespace scheduling policy used to schedule cluster namespaces on appropriate cluster attachments
type ClusterNamespaceSchedulingPolicy interface {
}

type clusternamespaceschedulingpolicy struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *clusternamespaceschedulingpolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerclusternamespaceschedulingpolicy clusternamespaceschedulingpolicy
	s := struct {
		Model Unmarshalerclusternamespaceschedulingpolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *clusternamespaceschedulingpolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "ALL_ATTACHMENTS":
		mm := AllAttachmentsSchedulingPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LEAST_USED_ATTACHMENT":
		mm := LeastUsedAttachmentSchedulingPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ClusterNamespaceSchedulingPolicy: %s.", m.Type)
		return *m, nil
	}
}

func (m clusternamespaceschedulingpolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m clusternamespaceschedulingpolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ClusterNamespaceSchedulingPolicyTypeEnum Enum with underlying type: string
type ClusterNamespaceSchedulingPolicyTypeEnum string

// Set of constants representing the allowable values for ClusterNamespaceSchedulingPolicyTypeEnum
const (
	ClusterNamespaceSchedulingPolicyTypeLeastUsedAttachment ClusterNamespaceSchedulingPolicyTypeEnum = "LEAST_USED_ATTACHMENT"
	ClusterNamespaceSchedulingPolicyTypeAllAttachments      ClusterNamespaceSchedulingPolicyTypeEnum = "ALL_ATTACHMENTS"
)

var mappingClusterNamespaceSchedulingPolicyTypeEnum = map[string]ClusterNamespaceSchedulingPolicyTypeEnum{
	"LEAST_USED_ATTACHMENT": ClusterNamespaceSchedulingPolicyTypeLeastUsedAttachment,
	"ALL_ATTACHMENTS":       ClusterNamespaceSchedulingPolicyTypeAllAttachments,
}

var mappingClusterNamespaceSchedulingPolicyTypeEnumLowerCase = map[string]ClusterNamespaceSchedulingPolicyTypeEnum{
	"least_used_attachment": ClusterNamespaceSchedulingPolicyTypeLeastUsedAttachment,
	"all_attachments":       ClusterNamespaceSchedulingPolicyTypeAllAttachments,
}

// GetClusterNamespaceSchedulingPolicyTypeEnumValues Enumerates the set of values for ClusterNamespaceSchedulingPolicyTypeEnum
func GetClusterNamespaceSchedulingPolicyTypeEnumValues() []ClusterNamespaceSchedulingPolicyTypeEnum {
	values := make([]ClusterNamespaceSchedulingPolicyTypeEnum, 0)
	for _, v := range mappingClusterNamespaceSchedulingPolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetClusterNamespaceSchedulingPolicyTypeEnumStringValues Enumerates the set of values in String for ClusterNamespaceSchedulingPolicyTypeEnum
func GetClusterNamespaceSchedulingPolicyTypeEnumStringValues() []string {
	return []string{
		"LEAST_USED_ATTACHMENT",
		"ALL_ATTACHMENTS",
	}
}

// GetMappingClusterNamespaceSchedulingPolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClusterNamespaceSchedulingPolicyTypeEnum(val string) (ClusterNamespaceSchedulingPolicyTypeEnum, bool) {
	enum, ok := mappingClusterNamespaceSchedulingPolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
