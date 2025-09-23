// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TaskDetails The minimum details of a task.
type TaskDetails interface {
}

type taskdetails struct {
	JsonData []byte
	TaskType string `json:"taskType"`
}

// UnmarshalJSON unmarshals json
func (m *taskdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertaskdetails taskdetails
	s := struct {
		Model Unmarshalertaskdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TaskType = s.Model.TaskType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *taskdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TaskType {
	case "DEPLOYED_APPLICATION_MIGRATION":
		mm := DeployedApplicationMigrationTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REMOVE_INSTALLATION_SITE":
		mm := RemoveInstallationSiteTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CRYPTO":
		mm := CryptoTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "JAVA_MIGRATION":
		mm := JavaMigrationTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SCAN_LIBRARY":
		mm := ScanLibraryTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PERFORMANCE_TUNING":
		mm := PerformanceTuningTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SCAN_JAVA_SERVER":
		mm := ScanJavaServerTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "JFR":
		mm := JfrTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ADD_INSTALLATION_SITE":
		mm := AddInstallationSiteTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for TaskDetails: %s.", m.TaskType)
		return *m, nil
	}
}

func (m taskdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m taskdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TaskDetailsTaskTypeEnum Enum with underlying type: string
type TaskDetailsTaskTypeEnum string

// Set of constants representing the allowable values for TaskDetailsTaskTypeEnum
const (
	TaskDetailsTaskTypeCrypto                       TaskDetailsTaskTypeEnum = "CRYPTO"
	TaskDetailsTaskTypeJfr                          TaskDetailsTaskTypeEnum = "JFR"
	TaskDetailsTaskTypeScanLibrary                  TaskDetailsTaskTypeEnum = "SCAN_LIBRARY"
	TaskDetailsTaskTypeScanJavaServer               TaskDetailsTaskTypeEnum = "SCAN_JAVA_SERVER"
	TaskDetailsTaskTypeJavaMigration                TaskDetailsTaskTypeEnum = "JAVA_MIGRATION"
	TaskDetailsTaskTypeDeployedApplicationMigration TaskDetailsTaskTypeEnum = "DEPLOYED_APPLICATION_MIGRATION"
	TaskDetailsTaskTypePerformanceTuning            TaskDetailsTaskTypeEnum = "PERFORMANCE_TUNING"
	TaskDetailsTaskTypeAddInstallationSite          TaskDetailsTaskTypeEnum = "ADD_INSTALLATION_SITE"
	TaskDetailsTaskTypeRemoveInstallationSite       TaskDetailsTaskTypeEnum = "REMOVE_INSTALLATION_SITE"
)

var mappingTaskDetailsTaskTypeEnum = map[string]TaskDetailsTaskTypeEnum{
	"CRYPTO":                         TaskDetailsTaskTypeCrypto,
	"JFR":                            TaskDetailsTaskTypeJfr,
	"SCAN_LIBRARY":                   TaskDetailsTaskTypeScanLibrary,
	"SCAN_JAVA_SERVER":               TaskDetailsTaskTypeScanJavaServer,
	"JAVA_MIGRATION":                 TaskDetailsTaskTypeJavaMigration,
	"DEPLOYED_APPLICATION_MIGRATION": TaskDetailsTaskTypeDeployedApplicationMigration,
	"PERFORMANCE_TUNING":             TaskDetailsTaskTypePerformanceTuning,
	"ADD_INSTALLATION_SITE":          TaskDetailsTaskTypeAddInstallationSite,
	"REMOVE_INSTALLATION_SITE":       TaskDetailsTaskTypeRemoveInstallationSite,
}

var mappingTaskDetailsTaskTypeEnumLowerCase = map[string]TaskDetailsTaskTypeEnum{
	"crypto":                         TaskDetailsTaskTypeCrypto,
	"jfr":                            TaskDetailsTaskTypeJfr,
	"scan_library":                   TaskDetailsTaskTypeScanLibrary,
	"scan_java_server":               TaskDetailsTaskTypeScanJavaServer,
	"java_migration":                 TaskDetailsTaskTypeJavaMigration,
	"deployed_application_migration": TaskDetailsTaskTypeDeployedApplicationMigration,
	"performance_tuning":             TaskDetailsTaskTypePerformanceTuning,
	"add_installation_site":          TaskDetailsTaskTypeAddInstallationSite,
	"remove_installation_site":       TaskDetailsTaskTypeRemoveInstallationSite,
}

// GetTaskDetailsTaskTypeEnumValues Enumerates the set of values for TaskDetailsTaskTypeEnum
func GetTaskDetailsTaskTypeEnumValues() []TaskDetailsTaskTypeEnum {
	values := make([]TaskDetailsTaskTypeEnum, 0)
	for _, v := range mappingTaskDetailsTaskTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskDetailsTaskTypeEnumStringValues Enumerates the set of values in String for TaskDetailsTaskTypeEnum
func GetTaskDetailsTaskTypeEnumStringValues() []string {
	return []string{
		"CRYPTO",
		"JFR",
		"SCAN_LIBRARY",
		"SCAN_JAVA_SERVER",
		"JAVA_MIGRATION",
		"DEPLOYED_APPLICATION_MIGRATION",
		"PERFORMANCE_TUNING",
		"ADD_INSTALLATION_SITE",
		"REMOVE_INSTALLATION_SITE",
	}
}

// GetMappingTaskDetailsTaskTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskDetailsTaskTypeEnum(val string) (TaskDetailsTaskTypeEnum, bool) {
	enum, ok := mappingTaskDetailsTaskTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
