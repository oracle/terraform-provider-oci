// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ContainerSummary Information about a container in the cluster.
type ContainerSummary struct {

	// Unique identifier for the container.
	ContainerKey *string `mandatory:"true" json:"containerKey"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated managed instance of type OCMA.
	ManagedInstanceId *string `mandatory:"true" json:"managedInstanceId"`

	// The name of the container.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The container image name.
	ImageName *string `mandatory:"true" json:"imageName"`

	// The namespace of the container.
	Namespace *string `mandatory:"true" json:"namespace"`

	// The name of the node associated with the pod running this container.
	NodeName *string `mandatory:"true" json:"nodeName"`

	// The name of the pod running this container.
	PodName *string `mandatory:"true" json:"podName"`

	// Unique key that identifies the application running in the container.
	ApplicationKey *string `mandatory:"false" json:"applicationKey"`

	// The name of the application running in the container.
	ApplicationName *string `mandatory:"false" json:"applicationName"`

	// Unique key that identifies the Java runtime used to run the application in the container.
	JreKey *string `mandatory:"false" json:"jreKey"`

	// The Java runtime used to run the application in the container.
	JavaVersion *string `mandatory:"false" json:"javaVersion"`

	// The security status of the Java runtime used to run the application in the container.
	JreSecurityStatus JreSecurityStatusEnum `mandatory:"false" json:"jreSecurityStatus,omitempty"`

	// The start time of the container.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`
}

func (m ContainerSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ContainerSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingJreSecurityStatusEnum(string(m.JreSecurityStatus)); !ok && m.JreSecurityStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JreSecurityStatus: %s. Supported values are: %s.", m.JreSecurityStatus, strings.Join(GetJreSecurityStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
