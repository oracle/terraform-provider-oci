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

// CreateBaseModelGroupDetails The base create model group details.
type CreateBaseModelGroupDetails interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the modelGroup in.
	GetCompartmentId() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the modelGroup.
	GetProjectId() *string
}

type createbasemodelgroupdetails struct {
	JsonData      []byte
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
	ProjectId     *string `mandatory:"true" json:"projectId"`
	CreateType    string  `json:"createType"`
}

// UnmarshalJSON unmarshals json
func (m *createbasemodelgroupdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatebasemodelgroupdetails createbasemodelgroupdetails
	s := struct {
		Model Unmarshalercreatebasemodelgroupdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.ProjectId = s.Model.ProjectId
	m.CreateType = s.Model.CreateType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createbasemodelgroupdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.CreateType {
	case "CLONE":
		mm := CloneModelGroupDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CREATE":
		mm := CreateModelGroupDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateBaseModelGroupDetails: %s.", m.CreateType)
		return *m, nil
	}
}

// GetCompartmentId returns CompartmentId
func (m createbasemodelgroupdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetProjectId returns ProjectId
func (m createbasemodelgroupdetails) GetProjectId() *string {
	return m.ProjectId
}

func (m createbasemodelgroupdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createbasemodelgroupdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
