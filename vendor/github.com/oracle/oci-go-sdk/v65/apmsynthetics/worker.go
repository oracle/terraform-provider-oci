// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Worker The information about an On-premise VP worker.
type Worker struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the On-premise VP worker.
	Id *string `mandatory:"true" json:"id"`

	// The runtime assigned id of the On-premise VP worker.
	RuntimeId *string `mandatory:"true" json:"runtimeId"`

	// Unique On-premise VP worker name that cannot be edited. The name should not contain any confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique permanent name of the On-premise VP worker. This is the same as the displayName.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the On-premise vantage point.
	OpvpId *string `mandatory:"true" json:"opvpId"`

	// On-premise vantage point name.
	OpvpName *string `mandatory:"true" json:"opvpName"`

	VersionDetails *OnPremiseVpWorkerVersionDetails `mandatory:"true" json:"versionDetails"`

	// Type of the On-premise VP worker.
	WorkerType OnPremiseVantagePointWorkerTypeEnum `mandatory:"true" json:"workerType"`

	// Enables or disables the On-premise VP worker.
	Status OnPremiseVantagePointWorkerStatusEnum `mandatory:"true" json:"status"`

	// Priority of the On-premise VP worker to schedule monitors.
	Priority *int `mandatory:"true" json:"priority"`

	// Configuration details of the On-premise VP worker.
	ConfigurationDetails *interface{} `mandatory:"false" json:"configurationDetails"`

	// Geographical information of the On-premise VP worker.
	GeoInfo *string `mandatory:"false" json:"geoInfo"`

	// Monitors list assigned to the On-premise VP worker.
	MonitorList []WorkerMonitorList `mandatory:"false" json:"monitorList"`

	IdentityInfo *IdentityInfoDetails `mandatory:"false" json:"identityInfo"`

	// The time the resource was last synced, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-12T22:47:12.613Z`
	TimeLastSyncUp *common.SDKTime `mandatory:"false" json:"timeLastSyncUp"`

	// The time the resource was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-12T22:47:12.613Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the resource was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-13T22:47:12.613Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Worker) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Worker) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOnPremiseVantagePointWorkerTypeEnum(string(m.WorkerType)); !ok && m.WorkerType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkerType: %s. Supported values are: %s.", m.WorkerType, strings.Join(GetOnPremiseVantagePointWorkerTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOnPremiseVantagePointWorkerStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetOnPremiseVantagePointWorkerStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
