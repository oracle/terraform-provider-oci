// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JavaMigrationAnalysisResult Result of the Java migration analysis. The analysis result is stored in an Object Storage bucket.
type JavaMigrationAnalysisResult struct {

	// The OCID of the migration analysis report.
	Id *string `mandatory:"true" json:"id"`

	// The fleet OCID.
	FleetId *string `mandatory:"true" json:"fleetId"`

	// The name of the application for which the Java migration analysis was performed.
	ApplicationName *string `mandatory:"true" json:"applicationName"`

	// The installation path of the application for which the Java migration analysis was performed.
	ApplicationPath *string `mandatory:"true" json:"applicationPath"`

	// Execution type of the application for an application type, such as WAR and EAR, that is deployed or installed.
	ApplicationExecutionType ApplicationExecutionTypeEnum `mandatory:"true" json:"applicationExecutionType"`

	// The source JDK version of the application that's currently running.
	SourceJdkVersion *string `mandatory:"true" json:"sourceJdkVersion"`

	// The target JDK version of the application to be migrated.
	TargetJdkVersion *string `mandatory:"true" json:"targetJdkVersion"`

	// The object storage namespace that contains the results of the migration analysis.
	Namespace *string `mandatory:"true" json:"namespace"`

	// The name of the object storage bucket that contains the results of the migration analysis.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// The directory path of the object storage bucket that contains the results of the migration analysis.
	ObjectStorageUploadDirPath *string `mandatory:"true" json:"objectStorageUploadDirPath"`

	// The names of the object storage objects that contain the results of the migration analysis.
	ObjectList []string `mandatory:"true" json:"objectList"`

	// Additional info reserved for future use.
	Metadata *string `mandatory:"true" json:"metadata"`

	// The OCID of the work request of this analysis.
	WorkRequestId *string `mandatory:"false" json:"workRequestId"`

	// The unique key that identifies the application.
	ApplicationKey *string `mandatory:"false" json:"applicationKey"`

	// The managed instance OCID.
	ManagedInstanceId *string `mandatory:"false" json:"managedInstanceId"`

	// The hostname of the managed instance that hosts the application for which the Java migration analysis was performed.
	HostName *string `mandatory:"false" json:"hostName"`

	// The time the result is compiled.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m JavaMigrationAnalysisResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JavaMigrationAnalysisResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingApplicationExecutionTypeEnum(string(m.ApplicationExecutionType)); !ok && m.ApplicationExecutionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ApplicationExecutionType: %s. Supported values are: %s.", m.ApplicationExecutionType, strings.Join(GetApplicationExecutionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
