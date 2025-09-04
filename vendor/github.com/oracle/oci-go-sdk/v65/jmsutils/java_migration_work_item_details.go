// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Utilities API
//
// The APIs for Analyze Applications and other utilities of Java Management Service.
//

package jmsutils

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JavaMigrationWorkItemDetails The java migration work item details.
type JavaMigrationWorkItemDetails struct {

	// The JDK version against which the migration analysis was performed to identify effort required to move from source JDK.
	TargetJdkVersion *string `mandatory:"true" json:"targetJdkVersion"`

	// Object storage path to the input artifact/s in the form of a serialized array.
	// Example: "[\"/JMS/Utils/myartifacts1.jar\",\"/JMS/Utils/myartifacts2.war\"]"
	InputApplicationsObjectStoragePaths *string `mandatory:"true" json:"inputApplicationsObjectStoragePaths"`

	// Name of the analysis project.
	AnalysisProjectName *string `mandatory:"true" json:"analysisProjectName"`

	// The work item type.
	WorkItemType WorkItemTypeEnum `mandatory:"false" json:"workItemType,omitempty"`
}

// GetWorkItemType returns WorkItemType
func (m JavaMigrationWorkItemDetails) GetWorkItemType() WorkItemTypeEnum {
	return m.WorkItemType
}

func (m JavaMigrationWorkItemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JavaMigrationWorkItemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingWorkItemTypeEnum(string(m.WorkItemType)); !ok && m.WorkItemType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkItemType: %s. Supported values are: %s.", m.WorkItemType, strings.Join(GetWorkItemTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m JavaMigrationWorkItemDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeJavaMigrationWorkItemDetails JavaMigrationWorkItemDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeJavaMigrationWorkItemDetails
	}{
		"JAVA_MIGRATION",
		(MarshalTypeJavaMigrationWorkItemDetails)(m),
	}

	return json.Marshal(&s)
}
