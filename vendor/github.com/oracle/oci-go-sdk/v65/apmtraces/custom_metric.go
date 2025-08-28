// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Trace Explorer API
//
// Use the Application Performance Monitoring Trace Explorer API to query traces and associated spans in Trace Explorer. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmtraces

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CustomMetric Definition of the Custom Metric.
type CustomMetric struct {

	// Name of the Custom Metric.
	Name *string `mandatory:"true" json:"name"`

	// Namespace in the Custom Metric. It defaults to `oracle_apm_custom` if not specified.
	// If specified, the necessary OCI policies should be set to allow APM to write to that namespace.
	Namespace *string `mandatory:"false" json:"namespace"`

	// Description of the Custom Metric.
	Description *string `mandatory:"false" json:"description"`

	// Resource Group of the Custom Metric.
	ResourceGroup *string `mandatory:"false" json:"resourceGroup"`

	// Indicates whether anomaly Detection should be performed on the generated metric.
	IsAnomalyDetectionEnabled *bool `mandatory:"false" json:"isAnomalyDetectionEnabled"`

	// Compartment of the Monitoring Service. It defaults to the APM domain's compartment if not specified.
	// If specified, the necessary OCI policies should be set to allow APM to write to that compartment.
	Compartment *string `mandatory:"false" json:"compartment"`

	// Unit in which the metric value is reported. For example 'ms'.
	Unit *string `mandatory:"false" json:"unit"`

	// Used in conjunction with the dry run header.  When the dry run header is set and the isPublishMetric flag is set to true, the
	// scheduled query is not created, but validations happen to check if the right OCI policies have been set to write to the specified
	// namespace/compartment.
	IsMetricPublished *bool `mandatory:"false" json:"isMetricPublished"`
}

func (m CustomMetric) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CustomMetric) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
