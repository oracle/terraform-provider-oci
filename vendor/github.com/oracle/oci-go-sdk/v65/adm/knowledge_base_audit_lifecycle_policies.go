// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.oracle.com/iaas/Content/application-dependency-management/home.htm).
//

package adm

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KnowledgeBaseAuditLifecyclePolicies The audit lifecycle policies to apply to all audits associated with the knowledge base.
// The lifecycle policy must be described per audit type.
// The default lifecycle policy for vulnerability audits is RetainLifecyclePolicy.
// The default lifecycle policy for artifact version audits is DeleteAfterLifecyclePolicy with a deleteAfterDuration of 1 day.
type KnowledgeBaseAuditLifecyclePolicies struct {
	ForVulnerabilityAudits KnowledgeBaseAuditLifecyclePolicy `mandatory:"false" json:"forVulnerabilityAudits"`

	ForArtifactVersionAudits KnowledgeBaseAuditLifecyclePolicy `mandatory:"false" json:"forArtifactVersionAudits"`
}

func (m KnowledgeBaseAuditLifecyclePolicies) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KnowledgeBaseAuditLifecyclePolicies) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *KnowledgeBaseAuditLifecyclePolicies) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ForVulnerabilityAudits   knowledgebaseauditlifecyclepolicy `json:"forVulnerabilityAudits"`
		ForArtifactVersionAudits knowledgebaseauditlifecyclepolicy `json:"forArtifactVersionAudits"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.ForVulnerabilityAudits.UnmarshalPolymorphicJSON(model.ForVulnerabilityAudits.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ForVulnerabilityAudits = nn.(KnowledgeBaseAuditLifecyclePolicy)
	} else {
		m.ForVulnerabilityAudits = nil
	}

	nn, e = model.ForArtifactVersionAudits.UnmarshalPolymorphicJSON(model.ForArtifactVersionAudits.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ForArtifactVersionAudits = nn.(KnowledgeBaseAuditLifecyclePolicy)
	} else {
		m.ForArtifactVersionAudits = nil
	}

	return
}
