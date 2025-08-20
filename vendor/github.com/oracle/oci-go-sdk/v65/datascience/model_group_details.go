// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ModelGroupDetails The model group details.
type ModelGroupDetails interface {

	// An array of custom metadata details for the model group.
	GetCustomMetadataList() []CustomMetadata
}

type modelgroupdetails struct {
	JsonData           []byte
	CustomMetadataList []CustomMetadata `mandatory:"false" json:"customMetadataList"`
	Type               string           `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *modelgroupdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermodelgroupdetails modelgroupdetails
	s := struct {
		Model Unmarshalermodelgroupdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CustomMetadataList = s.Model.CustomMetadataList
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *modelgroupdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "HETEROGENEOUS":
		mm := HeterogeneousModelGroupDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STACKED":
		mm := StackedModelGroupDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOMOGENEOUS":
		mm := HomogeneousModelGroupDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ModelGroupDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetCustomMetadataList returns CustomMetadataList
func (m modelgroupdetails) GetCustomMetadataList() []CustomMetadata {
	return m.CustomMetadataList
}

func (m modelgroupdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m modelgroupdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
