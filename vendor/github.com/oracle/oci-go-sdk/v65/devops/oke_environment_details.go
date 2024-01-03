// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OkeEnvironmentDetails Specifies Devops Deploy Environment.
type OkeEnvironmentDetails interface {
}

type okeenvironmentdetails struct {
	JsonData        []byte
	EnvironmentType string `json:"environmentType"`
}

// UnmarshalJSON unmarshals json
func (m *okeenvironmentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerokeenvironmentdetails okeenvironmentdetails
	s := struct {
		Model Unmarshalerokeenvironmentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.EnvironmentType = s.Model.EnvironmentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *okeenvironmentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EnvironmentType {
	case "CLUSTER_NAMESPACE":
		mm := OkeClusterNamespace{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_CLUSTER":
		mm := OkeCluster{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for OkeEnvironmentDetails: %s.", m.EnvironmentType)
		return *m, nil
	}
}

func (m okeenvironmentdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m okeenvironmentdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OkeEnvironmentDetailsEnvironmentTypeEnum Enum with underlying type: string
type OkeEnvironmentDetailsEnvironmentTypeEnum string

// Set of constants representing the allowable values for OkeEnvironmentDetailsEnvironmentTypeEnum
const (
	OkeEnvironmentDetailsEnvironmentTypeOkeCluster       OkeEnvironmentDetailsEnvironmentTypeEnum = "OKE_CLUSTER"
	OkeEnvironmentDetailsEnvironmentTypeClusterNamespace OkeEnvironmentDetailsEnvironmentTypeEnum = "CLUSTER_NAMESPACE"
)

var mappingOkeEnvironmentDetailsEnvironmentTypeEnum = map[string]OkeEnvironmentDetailsEnvironmentTypeEnum{
	"OKE_CLUSTER":       OkeEnvironmentDetailsEnvironmentTypeOkeCluster,
	"CLUSTER_NAMESPACE": OkeEnvironmentDetailsEnvironmentTypeClusterNamespace,
}

var mappingOkeEnvironmentDetailsEnvironmentTypeEnumLowerCase = map[string]OkeEnvironmentDetailsEnvironmentTypeEnum{
	"oke_cluster":       OkeEnvironmentDetailsEnvironmentTypeOkeCluster,
	"cluster_namespace": OkeEnvironmentDetailsEnvironmentTypeClusterNamespace,
}

// GetOkeEnvironmentDetailsEnvironmentTypeEnumValues Enumerates the set of values for OkeEnvironmentDetailsEnvironmentTypeEnum
func GetOkeEnvironmentDetailsEnvironmentTypeEnumValues() []OkeEnvironmentDetailsEnvironmentTypeEnum {
	values := make([]OkeEnvironmentDetailsEnvironmentTypeEnum, 0)
	for _, v := range mappingOkeEnvironmentDetailsEnvironmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOkeEnvironmentDetailsEnvironmentTypeEnumStringValues Enumerates the set of values in String for OkeEnvironmentDetailsEnvironmentTypeEnum
func GetOkeEnvironmentDetailsEnvironmentTypeEnumStringValues() []string {
	return []string{
		"OKE_CLUSTER",
		"CLUSTER_NAMESPACE",
	}
}

// GetMappingOkeEnvironmentDetailsEnvironmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOkeEnvironmentDetailsEnvironmentTypeEnum(val string) (OkeEnvironmentDetailsEnvironmentTypeEnum, bool) {
	enum, ok := mappingOkeEnvironmentDetailsEnvironmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
