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

// WorkerMonitorList Details of the monitor assigned to an On-premise vantage point worker.
type WorkerMonitorList struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the monitor.
	Id *string `mandatory:"true" json:"id"`

	// Unique name that can be edited. The name should not contain any confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Type of monitor.
	MonitorType MonitorTypesEnum `mandatory:"true" json:"monitorType"`

	// If isRunNow is enabled, then the monitor will run immediately.
	IsRunNow *bool `mandatory:"false" json:"isRunNow"`

	// The time the resource was last assigned to an On-premise vantage point worker, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-12T22:47:12.613Z`
	TimeAssigned *common.SDKTime `mandatory:"false" json:"timeAssigned"`
}

func (m WorkerMonitorList) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkerMonitorList) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMonitorTypesEnum(string(m.MonitorType)); !ok && m.MonitorType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MonitorType: %s. Supported values are: %s.", m.MonitorType, strings.Join(GetMonitorTypesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
