// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ImportOciTelemetryResourcesTaskDetails Request details for importing resources from Telemetry like resources from OCI Native Services and prometheus.
type ImportOciTelemetryResourcesTaskDetails struct {

	// Name space to be used for OCI Native service resources discovery.
	Namespace *string `mandatory:"true" json:"namespace"`

	// The resource group to use while fetching metrics from telemetry.
	// If not specified, resource group will be skipped in the list metrics request.
	ResourceGroup *string `mandatory:"false" json:"resourceGroup"`

	// Flag to indicate whether status is calculated using metrics or
	// LifeCycleState attribute of the resource in OCI service.
	ShouldUseMetricsFlowForStatus *bool `mandatory:"false" json:"shouldUseMetricsFlowForStatus"`

	// The base URL of the OCI service to which the resource belongs to.
	// Also this property is applicable only when source is OCI_TELEMETRY_NATIVE.
	ServiceBaseUrl *string `mandatory:"false" json:"serviceBaseUrl"`

	// The console path prefix to use for providing service home url page navigation.
	// For example if the prefix provided is 'security/bastion/bastions', the URL used for navigation will be
	// https://<cloudhostname>/security/bastion/bastions/<resourceOcid>. If not provided, service home page link
	// will not be shown in the stack monitoring home page.
	ConsolePathPrefix *string `mandatory:"false" json:"consolePathPrefix"`

	// Lifecycle states of the external resource which reflects the status of the resource being up.
	LifecycleStatusMappingsForUpStatus []string `mandatory:"false" json:"lifecycleStatusMappingsForUpStatus"`

	// The resource name property in the metric dimensions.
	// Resources imported will be using this property value for resource name.
	ResourceNameMapping *string `mandatory:"false" json:"resourceNameMapping"`

	// The external resource identifier property in the metric dimensions.
	// Resources imported will be using this property value for external id.
	ExternalIdMapping *string `mandatory:"false" json:"externalIdMapping"`

	// The resource type property in the metric dimensions.
	// Resources imported will be using this property value for resource type.
	// If not specified, namespace will be used for resource type.
	ResourceTypeMapping *string `mandatory:"false" json:"resourceTypeMapping"`

	// The resource name filter. Resources matching with the resource name filter will be imported.
	// Regular expressions will be accepted.
	ResourceNameFilter *string `mandatory:"false" json:"resourceNameFilter"`

	// The resource type filter. Resources matching with the resource type filter will be imported.
	// Regular expressions will be accepted.
	ResourceTypeFilter *string `mandatory:"false" json:"resourceTypeFilter"`

	// List of metrics to be used to calculate the availability of the resource.
	// Resource is considered to be up if at least one of the specified metrics is available for
	// the resource during the specified interval using the property
	// 'availabilityProxyMetricCollectionIntervalInSeconds'.
	// If no metrics are specified, availability will not be calculated for the resource.
	AvailabilityProxyMetrics []string `mandatory:"false" json:"availabilityProxyMetrics"`

	// Metrics collection interval in seconds used when calculating the availability of the
	// resource based on metrics specified using the property 'availabilityProxyMetrics'.
	AvailabilityProxyMetricCollectionInterval *int `mandatory:"false" json:"availabilityProxyMetricCollectionInterval"`

	// Source from where the metrics pushed to telemetry.
	// Possible values:
	//   * OCI_TELEMETRY_NATIVE      - The metrics are pushed to telemetry from OCI Native Services.
	//   * OCI_TELEMETRY_PROMETHEUS  - The metrics are pushed to telemetry from Prometheus.
	Source ImportOciTelemetryResourcesTaskDetailsSourceEnum `mandatory:"true" json:"source"`
}

func (m ImportOciTelemetryResourcesTaskDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImportOciTelemetryResourcesTaskDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingImportOciTelemetryResourcesTaskDetailsSourceEnum(string(m.Source)); !ok && m.Source != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Source: %s. Supported values are: %s.", m.Source, strings.Join(GetImportOciTelemetryResourcesTaskDetailsSourceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ImportOciTelemetryResourcesTaskDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeImportOciTelemetryResourcesTaskDetails ImportOciTelemetryResourcesTaskDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeImportOciTelemetryResourcesTaskDetails
	}{
		"IMPORT_OCI_TELEMETRY_RESOURCES",
		(MarshalTypeImportOciTelemetryResourcesTaskDetails)(m),
	}

	return json.Marshal(&s)
}

// ImportOciTelemetryResourcesTaskDetailsSourceEnum Enum with underlying type: string
type ImportOciTelemetryResourcesTaskDetailsSourceEnum string

// Set of constants representing the allowable values for ImportOciTelemetryResourcesTaskDetailsSourceEnum
const (
	ImportOciTelemetryResourcesTaskDetailsSourceNative     ImportOciTelemetryResourcesTaskDetailsSourceEnum = "OCI_TELEMETRY_NATIVE"
	ImportOciTelemetryResourcesTaskDetailsSourcePrometheus ImportOciTelemetryResourcesTaskDetailsSourceEnum = "OCI_TELEMETRY_PROMETHEUS"
)

var mappingImportOciTelemetryResourcesTaskDetailsSourceEnum = map[string]ImportOciTelemetryResourcesTaskDetailsSourceEnum{
	"OCI_TELEMETRY_NATIVE":     ImportOciTelemetryResourcesTaskDetailsSourceNative,
	"OCI_TELEMETRY_PROMETHEUS": ImportOciTelemetryResourcesTaskDetailsSourcePrometheus,
}

var mappingImportOciTelemetryResourcesTaskDetailsSourceEnumLowerCase = map[string]ImportOciTelemetryResourcesTaskDetailsSourceEnum{
	"oci_telemetry_native":     ImportOciTelemetryResourcesTaskDetailsSourceNative,
	"oci_telemetry_prometheus": ImportOciTelemetryResourcesTaskDetailsSourcePrometheus,
}

// GetImportOciTelemetryResourcesTaskDetailsSourceEnumValues Enumerates the set of values for ImportOciTelemetryResourcesTaskDetailsSourceEnum
func GetImportOciTelemetryResourcesTaskDetailsSourceEnumValues() []ImportOciTelemetryResourcesTaskDetailsSourceEnum {
	values := make([]ImportOciTelemetryResourcesTaskDetailsSourceEnum, 0)
	for _, v := range mappingImportOciTelemetryResourcesTaskDetailsSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetImportOciTelemetryResourcesTaskDetailsSourceEnumStringValues Enumerates the set of values in String for ImportOciTelemetryResourcesTaskDetailsSourceEnum
func GetImportOciTelemetryResourcesTaskDetailsSourceEnumStringValues() []string {
	return []string{
		"OCI_TELEMETRY_NATIVE",
		"OCI_TELEMETRY_PROMETHEUS",
	}
}

// GetMappingImportOciTelemetryResourcesTaskDetailsSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImportOciTelemetryResourcesTaskDetailsSourceEnum(val string) (ImportOciTelemetryResourcesTaskDetailsSourceEnum, bool) {
	enum, ok := mappingImportOciTelemetryResourcesTaskDetailsSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
