// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// CreateMlApplicationDetails The information about new MlApplication.
type CreateMlApplicationDetails struct {

	// The name of MlApplication. It is unique in a given tenancy.
	Name *string `mandatory:"true" json:"name"`

	PredictionContract *PredictionContract `mandatory:"true" json:"predictionContract"`

	// List of application components (OCI resources shared for all MlApplicationInstances).
	ApplicationComponents []ApplicationComponent `mandatory:"true" json:"applicationComponents"`

	// The OCID of the compartment where the MlApplication is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Optional description of the ML Application
	Description *string `mandatory:"false" json:"description"`

	// Schema of configuration which needs to be provided for each ML Application Instance
	ConfigurationSchema []ConfigurationPropertySchema `mandatory:"false" json:"configurationSchema"`

	// Vault ID used for secure persisting possible secrets from configuration
	VaultId *string `mandatory:"false" json:"vaultId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateMlApplicationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMlApplicationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateMlApplicationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description           *string                           `json:"description"`
		ConfigurationSchema   []ConfigurationPropertySchema     `json:"configurationSchema"`
		VaultId               *string                           `json:"vaultId"`
		FreeformTags          map[string]string                 `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{} `json:"definedTags"`
		Name                  *string                           `json:"name"`
		PredictionContract    *PredictionContract               `json:"predictionContract"`
		ApplicationComponents []applicationcomponent            `json:"applicationComponents"`
		CompartmentId         *string                           `json:"compartmentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.ConfigurationSchema = make([]ConfigurationPropertySchema, len(model.ConfigurationSchema))
	for i, n := range model.ConfigurationSchema {
		m.ConfigurationSchema[i] = n
	}

	m.VaultId = model.VaultId

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Name = model.Name

	m.PredictionContract = model.PredictionContract

	m.ApplicationComponents = make([]ApplicationComponent, len(model.ApplicationComponents))
	for i, n := range model.ApplicationComponents {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ApplicationComponents[i] = nn.(ApplicationComponent)
		} else {
			m.ApplicationComponents[i] = nil
		}
	}

	m.CompartmentId = model.CompartmentId

	return
}
