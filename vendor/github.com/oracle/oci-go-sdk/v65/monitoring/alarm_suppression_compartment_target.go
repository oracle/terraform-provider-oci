// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Monitoring API
//
// Use the Monitoring API to manage metric queries and alarms for assessing the health, capacity, and performance of your cloud resources.
// Endpoints vary by operation. For PostMetricData, use the `telemetry-ingestion` endpoints; for all other operations, use the `telemetry` endpoints.
// For more information, see
// the Monitoring documentation (https://docs.cloud.oracle.com/iaas/Content/Monitoring/home.htm).
//

package monitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AlarmSuppressionCompartmentTarget The compartment target of the alarm suppression.
type AlarmSuppressionCompartmentTarget struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment or tenancy that is the
	// target of the alarm suppression.
	// Example: `ocid1.compartment.oc1..exampleuniqueID`
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// When true, the alarm suppression targets all alarms under all compartments and subcompartments of
	// the tenancy specified. The parameter can only be set to true when compartmentId is the tenancy OCID
	// (the tenancy is the root compartment). When false, the alarm suppression targets only the alarms under
	// the specified compartment.
	CompartmentIdInSubtree *bool `mandatory:"false" json:"compartmentIdInSubtree"`
}

func (m AlarmSuppressionCompartmentTarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AlarmSuppressionCompartmentTarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AlarmSuppressionCompartmentTarget) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAlarmSuppressionCompartmentTarget AlarmSuppressionCompartmentTarget
	s := struct {
		DiscriminatorParam string `json:"targetType"`
		MarshalTypeAlarmSuppressionCompartmentTarget
	}{
		"COMPARTMENT",
		(MarshalTypeAlarmSuppressionCompartmentTarget)(m),
	}

	return json.Marshal(&s)
}
