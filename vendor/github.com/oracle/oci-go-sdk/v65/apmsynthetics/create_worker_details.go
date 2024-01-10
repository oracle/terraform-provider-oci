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

// CreateWorkerDetails Details of the request body used to create a new worker for an On-premise vantage point.
type CreateWorkerDetails struct {

	// Unique On-premise VP worker name that cannot be edited. The name should not contain any confidential information.
	Name *string `mandatory:"true" json:"name"`

	// Image version of the On-premise VP worker.
	Version *string `mandatory:"true" json:"version"`

	// public key for resource Principal Token based validation to be used in further calls.
	ResourcePrincipalTokenPublicKey *string `mandatory:"true" json:"resourcePrincipalTokenPublicKey"`

	// Configuration details of the On-premise VP worker.
	ConfigurationDetails *interface{} `mandatory:"false" json:"configurationDetails"`

	// Type of the On-premise VP worker.
	WorkerType OnPremiseVantagePointWorkerTypeEnum `mandatory:"false" json:"workerType,omitempty"`

	// Enables or disables the On-premise VP worker.
	Status OnPremiseVantagePointWorkerStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Priority of the On-premise VP worker to schedule monitors.
	Priority *int `mandatory:"false" json:"priority"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateWorkerDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateWorkerDetails) ValidateEnumValue() (bool, error) {
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
