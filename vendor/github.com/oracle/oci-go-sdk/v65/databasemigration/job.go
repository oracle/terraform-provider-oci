// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Job Results of a Database Connection search. Contains DatabaseConnectionSummary items.
type Job struct {

	// The OCID of the Migration Job.
	Id *string `mandatory:"true" json:"id"`

	// Name of the job.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the Migration that this job belongs to.
	MigrationId *string `mandatory:"true" json:"migrationId"`

	// The job type.
	Type JobTypesEnum `mandatory:"true" json:"type"`

	// The time the Migration Job was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the migration job.
	LifecycleState JobLifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// The time the Migration Job was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	Progress *MigrationJobProgressResource `mandatory:"false" json:"progress"`

	// Database objects not supported.
	UnsupportedObjects []UnsupportedDatabaseObject `mandatory:"false" json:"unsupportedObjects"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information
	// for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Job) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Job) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJobTypesEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetJobTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingJobLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetJobLifecycleStatesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
