// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MaintenanceNotificationDetails Notification details for the maintenance activity.
type MaintenanceNotificationDetails struct {

	// Jira tickets for on-call engineer to send customer notification
	JiraTicket *string `mandatory:"true" json:"jiraTicket"`

	// Maintenance Notification type
	MaintenanceNotificationType MaintenanceNotificationTypeEnum `mandatory:"true" json:"maintenanceNotificationType"`

	// List of tenantIds where we need to send the maintenance notifications
	TenantIds []string `mandatory:"false" json:"tenantIds"`

	// List of OpenSearch clusterIds where we need to send the maintenance notifications
	ClusterIds []string `mandatory:"false" json:"clusterIds"`

	// Start time for the notification activity in UTC like "2023-20-04 04:00:00.000Z"
	StartTime *string `mandatory:"false" json:"startTime"`

	// Start time for the notification activity in UTC like "2023-20-04 04:00:00.000Z"
	EndTime *string `mandatory:"false" json:"endTime"`
}

func (m MaintenanceNotificationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaintenanceNotificationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMaintenanceNotificationTypeEnum(string(m.MaintenanceNotificationType)); !ok && m.MaintenanceNotificationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceNotificationType: %s. Supported values are: %s.", m.MaintenanceNotificationType, strings.Join(GetMaintenanceNotificationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
