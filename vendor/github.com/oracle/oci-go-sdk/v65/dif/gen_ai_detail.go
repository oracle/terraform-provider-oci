// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataIntelligences Control Plane API
//
// Use the DataIntelligences Control Plane API to manage dataIntelligences.
//

package dif

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GenAiDetail GenAI details required to provision dedicated clusters.
type GenAiDetail struct {

	// Id for the GGCS instance to be provisioned.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// The dedicated AI cluster type.
	ClusterType ClusterTypeEnum `mandatory:"true" json:"clusterType"`

	// Name of the base model.
	BaseModel *string `mandatory:"true" json:"baseModel"`

	// Region on which the cluster end endpoint will be provisioned.
	OciRegion *string `mandatory:"true" json:"ociRegion"`

	// No of replicas of base model to be used for hosting.
	UnitCount *int `mandatory:"true" json:"unitCount"`

	// List of endpoints to provision for the GENAI cluster.
	Endpoints []EndpointDetails `mandatory:"false" json:"endpoints"`
}

func (m GenAiDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenAiDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingClusterTypeEnum(string(m.ClusterType)); !ok && m.ClusterType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClusterType: %s. Supported values are: %s.", m.ClusterType, strings.Join(GetClusterTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
