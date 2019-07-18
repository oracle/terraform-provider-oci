// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Monitoring API
//
// Use the Monitoring API to manage metric queries and alarms for assessing the health, capacity, and performance of your cloud resources.
// Endpoints vary by operation. For PostMetric, use the `telemetry-ingestion` endpoints; for all other operations, use the `telemetry` endpoints.
// For information about monitoring, see Monitoring Overview (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm).
//

package monitoring

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ChangeAlarmCompartmentDetails The configuration details for moving an alarm.
type ChangeAlarmCompartmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to move the alarm to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeAlarmCompartmentDetails) String() string {
	return common.PointerString(m)
}
