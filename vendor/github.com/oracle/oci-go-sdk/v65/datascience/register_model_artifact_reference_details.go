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

// RegisterModelArtifactReferenceDetails Parameters that are required to register a model artifact reference collection.
type RegisterModelArtifactReferenceDetails struct {

	// A list of model artifact references to register.
	ModelArtifactReferences []ModelArtifactReferenceDetails `mandatory:"true" json:"modelArtifactReferences"`
}

func (m RegisterModelArtifactReferenceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RegisterModelArtifactReferenceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *RegisterModelArtifactReferenceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ModelArtifactReferences []modelartifactreferencedetails `json:"modelArtifactReferences"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ModelArtifactReferences = make([]ModelArtifactReferenceDetails, len(model.ModelArtifactReferences))
	for i, n := range model.ModelArtifactReferences {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ModelArtifactReferences[i] = nn.(ModelArtifactReferenceDetails)
		} else {
			m.ModelArtifactReferences[i] = nil
		}
	}
	return
}
