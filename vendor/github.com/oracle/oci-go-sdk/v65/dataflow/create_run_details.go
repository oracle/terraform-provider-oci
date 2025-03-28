// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateRunDetails The create run details. The following properties are optional and override the default values
// set in the associated application:
//   - applicationId
//   - archiveUri
//   - applicationLogConfig
//   - arguments
//   - configuration
//   - definedTags
//   - displayName
//   - driverShape
//   - execute
//   - executorShape
//   - freeformTags
//   - logsBucketUri
//   - metastoreId
//   - numExecutors
//   - parameters
//   - sparkVersion
//   - warehouseBucketUri
//
// It is expected that either the applicationId or the execute parameter is specified; but not both.
// If both or none are set, a Bad Request (HTTP 400) status will be sent as the response.
// If an appicationId is not specified, then a value for the execute parameter is expected.
// Using data parsed from the value, a new application will be created and assicated with the new run.
// See information on the execute parameter for details on the format of this parameter.
// The optional parameter spark version can only be specified when using the execute parameter.  If it
// is not specified when using the execute parameter, the latest version will be used as default.
// If the execute parameter is not used, the spark version will be taken from the associated application.
// If displayName is not specified, it will be derived from the displayName of associated application or
// set by API using fileUri's application file name.
// Once a run is created, its properties (except for definedTags and freeformTags) cannot be changed.
// If the parent application's properties (including definedTags and freeformTags) are updated,
// the corresponding properties of the run will not update.
type CreateRunDetails struct {

	// The OCID of a compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ApplicationLogConfig *ApplicationLogConfig `mandatory:"false" json:"applicationLogConfig"`

	// The OCID of the associated application. If this value is set, then no value for the execute parameter is required. If this value is not set, then a value for the execute parameter is required, and a new application is created and associated with the new run.
	ApplicationId *string `mandatory:"false" json:"applicationId"`

	// A comma separated list of one or more archive files as Oracle Cloud Infrastructure URIs. For example, ``oci://path/to/a.zip,oci://path/to/b.zip``. An Oracle Cloud Infrastructure URI of an archive.zip file containing custom dependencies that may be used to support the execution of a Python, Java, or Scala application.
	// See https://docs.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat.
	ArchiveUri *string `mandatory:"false" json:"archiveUri"`

	// The arguments passed to the running application as command line arguments.  An argument is
	// either a plain text or a placeholder. Placeholders are replaced using values from the parameters
	// map.  Each placeholder specified must be represented in the parameters map else the request
	// (POST or PUT) will fail with a HTTP 400 status code.  Placeholders are specified as
	// `Service Api Spec`, where `name` is the name of the parameter.
	// Example:  `[ "--input", "${input_file}", "--name", "John Doe" ]`
	// If "input_file" has a value of "mydata.xml", then the value above will be translated to
	// `--input mydata.xml --name "John Doe"`
	Arguments []string `mandatory:"false" json:"arguments"`

	// The Spark configuration passed to the running process.
	// See https://spark.apache.org/docs/latest/configuration.html#available-properties.
	// Example: { "spark.app.name" : "My App Name", "spark.shuffle.io.maxRetries" : "4" }
	// Note: Not all Spark properties are permitted to be set.  Attempting to set a property that is
	// not allowed to be overwritten will cause a 400 status to be returned.
	Configuration map[string]string `mandatory:"false" json:"configuration"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name that does not have to be unique. Avoid entering confidential information. If this value is not specified, it will be derived from the associated application's displayName or set by API using fileUri's application file name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The VM shape for the driver. Sets the driver cores and memory.
	DriverShape *string `mandatory:"false" json:"driverShape"`

	DriverShapeConfig *ShapeConfig `mandatory:"false" json:"driverShapeConfig"`

	// The input used for spark-submit command. For more details see https://spark.apache.org/docs/latest/submitting-applications.html#launching-applications-with-spark-submit.
	// Supported options include ``--class``, ``--file``, ``--jars``, ``--conf``, ``--py-files``, and main application file with arguments.
	// Example: ``--jars oci://path/to/a.jar,oci://path/to/b.jar --files oci://path/to/a.json,oci://path/to/b.csv --py-files oci://path/to/a.py,oci://path/to/b.py --conf spark.sql.crossJoin.enabled=true --class org.apache.spark.examples.SparkPi oci://path/to/main.jar 10``
	// Note: If execute is specified together with applicationId, className, configuration, fileUri, language, arguments, parameters during application create/update, or run create/submit,
	// Data Flow service will use derived information from execute input only.
	Execute *string `mandatory:"false" json:"execute"`

	// The VM shape for the executors. Sets the executor cores and memory.
	ExecutorShape *string `mandatory:"false" json:"executorShape"`

	ExecutorShapeConfig *ShapeConfig `mandatory:"false" json:"executorShapeConfig"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// An Oracle Cloud Infrastructure URI of the bucket where the Spark job logs are to be uploaded.
	// See https://docs.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat.
	LogsBucketUri *string `mandatory:"false" json:"logsBucketUri"`

	// The OCID of OCI Hive Metastore.
	MetastoreId *string `mandatory:"false" json:"metastoreId"`

	// The number of executor VMs requested.
	NumExecutors *int `mandatory:"false" json:"numExecutors"`

	// An array of name/value pairs used to fill placeholders found in properties like
	// `Application.arguments`.  The name must be a string of one or more word characters
	// (a-z, A-Z, 0-9, _).  The value can be a string of 0 or more characters of any kind.
	// Example:  [ { name: "iterations", value: "10"}, { name: "input_file", value: "mydata.xml" }, { name: "variable_x", value: "${x}"} ]
	Parameters []ApplicationParameter `mandatory:"false" json:"parameters"`

	// The OCID of a pool. Unique Id to indentify a dataflow pool resource.
	PoolId *string `mandatory:"false" json:"poolId"`

	// The Spark version utilized to run the application. This value may be set if applicationId is not since the Spark version will be taken from the associated application.
	SparkVersion *string `mandatory:"false" json:"sparkVersion"`

	// The Spark application processing type.
	Type ApplicationTypeEnum `mandatory:"false" json:"type,omitempty"`

	// An Oracle Cloud Infrastructure URI of the bucket to be used as default warehouse directory
	// for BATCH SQL runs.
	// See https://docs.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat.
	WarehouseBucketUri *string `mandatory:"false" json:"warehouseBucketUri"`

	// The maximum duration in minutes for which an Application should run. Data Flow Run would be terminated
	// once it reaches this duration from the time it transitions to `IN_PROGRESS` state.
	MaxDurationInMinutes *int64 `mandatory:"false" json:"maxDurationInMinutes"`

	// The timeout value in minutes used to manage Runs. A Run would be stopped after inactivity for this amount of time period.
	// Note: This parameter is currently only applicable for Runs of type `SESSION`. Default value is 2880 minutes (2 days)
	IdleTimeoutInMinutes *int64 `mandatory:"false" json:"idleTimeoutInMinutes"`
}

func (m CreateRunDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateRunDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingApplicationTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetApplicationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
