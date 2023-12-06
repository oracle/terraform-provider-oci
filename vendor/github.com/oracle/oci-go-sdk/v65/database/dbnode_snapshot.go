// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DbnodeSnapshot Details of the Database Node Snapshot.
type DbnodeSnapshot struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata Database Node Snapshot.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the Database Node Snapshot. The name should be unique.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata Database Node.
	SourceDbnodeId *string `mandatory:"true" json:"sourceDbnodeId"`

	// The current state of the Exadata Database Node Snapshot.
	LifecycleState DbnodeSnapshotLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time that the Exadata Database Node Snapshot was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Additional information about the current lifecycle state of the Exadata Database Node Snapshot.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata VM cluster on Exascale Infrastructure.
	ExadbVmClusterId *string `mandatory:"false" json:"exadbVmClusterId"`

	// Details of the mount points
	MountPoints []MountPointDetails `mandatory:"false" json:"mountPoints"`

	// Details of the volumes
	Volumes []VolumeDetails `mandatory:"false" json:"volumes"`
}

func (m DbnodeSnapshot) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbnodeSnapshot) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDbnodeSnapshotLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbnodeSnapshotLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbnodeSnapshotLifecycleStateEnum Enum with underlying type: string
type DbnodeSnapshotLifecycleStateEnum string

// Set of constants representing the allowable values for DbnodeSnapshotLifecycleStateEnum
const (
	DbnodeSnapshotLifecycleStateCreating    DbnodeSnapshotLifecycleStateEnum = "CREATING"
	DbnodeSnapshotLifecycleStateAvailable   DbnodeSnapshotLifecycleStateEnum = "AVAILABLE"
	DbnodeSnapshotLifecycleStateTerminating DbnodeSnapshotLifecycleStateEnum = "TERMINATING"
	DbnodeSnapshotLifecycleStateTerminated  DbnodeSnapshotLifecycleStateEnum = "TERMINATED"
	DbnodeSnapshotLifecycleStateFailed      DbnodeSnapshotLifecycleStateEnum = "FAILED"
	DbnodeSnapshotLifecycleStateMounted     DbnodeSnapshotLifecycleStateEnum = "MOUNTED"
	DbnodeSnapshotLifecycleStateMounting    DbnodeSnapshotLifecycleStateEnum = "MOUNTING"
	DbnodeSnapshotLifecycleStateUnmounting  DbnodeSnapshotLifecycleStateEnum = "UNMOUNTING"
)

var mappingDbnodeSnapshotLifecycleStateEnum = map[string]DbnodeSnapshotLifecycleStateEnum{
	"CREATING":    DbnodeSnapshotLifecycleStateCreating,
	"AVAILABLE":   DbnodeSnapshotLifecycleStateAvailable,
	"TERMINATING": DbnodeSnapshotLifecycleStateTerminating,
	"TERMINATED":  DbnodeSnapshotLifecycleStateTerminated,
	"FAILED":      DbnodeSnapshotLifecycleStateFailed,
	"MOUNTED":     DbnodeSnapshotLifecycleStateMounted,
	"MOUNTING":    DbnodeSnapshotLifecycleStateMounting,
	"UNMOUNTING":  DbnodeSnapshotLifecycleStateUnmounting,
}

var mappingDbnodeSnapshotLifecycleStateEnumLowerCase = map[string]DbnodeSnapshotLifecycleStateEnum{
	"creating":    DbnodeSnapshotLifecycleStateCreating,
	"available":   DbnodeSnapshotLifecycleStateAvailable,
	"terminating": DbnodeSnapshotLifecycleStateTerminating,
	"terminated":  DbnodeSnapshotLifecycleStateTerminated,
	"failed":      DbnodeSnapshotLifecycleStateFailed,
	"mounted":     DbnodeSnapshotLifecycleStateMounted,
	"mounting":    DbnodeSnapshotLifecycleStateMounting,
	"unmounting":  DbnodeSnapshotLifecycleStateUnmounting,
}

// GetDbnodeSnapshotLifecycleStateEnumValues Enumerates the set of values for DbnodeSnapshotLifecycleStateEnum
func GetDbnodeSnapshotLifecycleStateEnumValues() []DbnodeSnapshotLifecycleStateEnum {
	values := make([]DbnodeSnapshotLifecycleStateEnum, 0)
	for _, v := range mappingDbnodeSnapshotLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDbnodeSnapshotLifecycleStateEnumStringValues Enumerates the set of values in String for DbnodeSnapshotLifecycleStateEnum
func GetDbnodeSnapshotLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
		"MOUNTED",
		"MOUNTING",
		"UNMOUNTING",
	}
}

// GetMappingDbnodeSnapshotLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbnodeSnapshotLifecycleStateEnum(val string) (DbnodeSnapshotLifecycleStateEnum, bool) {
	enum, ok := mappingDbnodeSnapshotLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
