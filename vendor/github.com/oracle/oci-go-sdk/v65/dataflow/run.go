// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// Run A run object.
type Run struct {

	// The application ID.
	ApplicationId *string `mandatory:"true" json:"applicationId"`

	// The OCID of a compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The VM shape for the driver. Sets the driver cores and memory.
	DriverShape *string `mandatory:"true" json:"driverShape"`

	// The VM shape for the executors. Sets the executor cores and memory.
	ExecutorShape *string `mandatory:"true" json:"executorShape"`

	// An Oracle Cloud Infrastructure URI of the file containing the application to execute.
	// See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat.
	FileUri *string `mandatory:"true" json:"fileUri"`

	// The ID of a run.
	Id *string `mandatory:"true" json:"id"`

	// The Spark language.
	Language ApplicationLanguageEnum `mandatory:"true" json:"language"`

	// The current state of this run.
	LifecycleState RunLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The number of executor VMs requested.
	NumExecutors *int `mandatory:"true" json:"numExecutors"`

	// The Spark version utilized to run the application.
	SparkVersion *string `mandatory:"true" json:"sparkVersion"`

	// The date and time the resource was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the resource was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-04-03T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// A comma separated list of one or more archive files as Oracle Cloud Infrastructure URIs. For example, ``oci://path/to/a.zip,oci://path/to/b.zip``. An Oracle Cloud Infrastructure URI of an archive.zip file containing custom dependencies that may be used to support the execution of a Python, Java, or Scala application.
	// See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat.
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

	ApplicationLogConfig *ApplicationLogConfig `mandatory:"false" json:"applicationLogConfig"`

	// The class for the application.
	ClassName *string `mandatory:"false" json:"className"`

	// The Spark configuration passed to the running process.
	// See https://spark.apache.org/docs/latest/configuration.html#available-properties.
	// Example: { "spark.app.name" : "My App Name", "spark.shuffle.io.maxRetries" : "4" }
	// Note: Not all Spark properties are permitted to be set.  Attempting to set a property that is
	// not allowed to be overwritten will cause a 400 status to be returned.
	Configuration map[string]string `mandatory:"false" json:"configuration"`

	// The data read by the run in bytes.
	DataReadInBytes *int64 `mandatory:"false" json:"dataReadInBytes"`

	// The data written by the run in bytes.
	DataWrittenInBytes *int64 `mandatory:"false" json:"dataWrittenInBytes"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. This name is not necessarily unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	DriverShapeConfig *ShapeConfig `mandatory:"false" json:"driverShapeConfig"`

	// The input used for spark-submit command. For more details see https://spark.apache.org/docs/latest/submitting-applications.html#launching-applications-with-spark-submit.
	// Supported options include ``--class``, ``--file``, ``--jars``, ``--conf``, ``--py-files``, and main application file with arguments.
	// Example: ``--jars oci://path/to/a.jar,oci://path/to/b.jar --files oci://path/to/a.json,oci://path/to/b.csv --py-files oci://path/to/a.py,oci://path/to/b.py --conf spark.sql.crossJoin.enabled=true --class org.apache.spark.examples.SparkPi oci://path/to/main.jar 10``
	// Note: If execute is specified together with applicationId, className, configuration, fileUri, language, arguments, parameters during application create/update, or run create/submit,
	// Data Flow service will use derived information from execute input only.
	Execute *string `mandatory:"false" json:"execute"`

	ExecutorShapeConfig *ShapeConfig `mandatory:"false" json:"executorShapeConfig"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The detailed messages about the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// An Oracle Cloud Infrastructure URI of the bucket where the Spark job logs are to be uploaded.
	// See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat.
	LogsBucketUri *string `mandatory:"false" json:"logsBucketUri"`

	// The OCID of OCI Hive Metastore.
	MetastoreId *string `mandatory:"false" json:"metastoreId"`

	// Unique Oracle assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" json:"opcRequestId"`

	// The OCID of the user who created the resource.
	OwnerPrincipalId *string `mandatory:"false" json:"ownerPrincipalId"`

	// The username of the user who created the resource.  If the username of the owner does not exist,
	// `null` will be returned and the caller should refer to the ownerPrincipalId value instead.
	OwnerUserName *string `mandatory:"false" json:"ownerUserName"`

	// An array of name/value pairs used to fill placeholders found in properties like
	// `Application.arguments`.  The name must be a string of one or more word characters
	// (a-z, A-Z, 0-9, _).  The value can be a string of 0 or more characters of any kind.
	// Example:  [ { name: "iterations", value: "10"}, { name: "input_file", value: "mydata.xml" }, { name: "variable_x", value: "${x}"} ]
	Parameters []ApplicationParameter `mandatory:"false" json:"parameters"`

	// The OCID of a pool. Unique Id to indentify a dataflow pool resource.
	PoolId *string `mandatory:"false" json:"poolId"`

	// An array of DNS zone names.
	// Example: `[ "app.examplecorp.com", "app.examplecorp2.com" ]`
	PrivateEndpointDnsZones []string `mandatory:"false" json:"privateEndpointDnsZones"`

	// The maximum number of hosts to be accessed through the private endpoint. This value is used
	// to calculate the relevant CIDR block and should be a multiple of 256.  If the value is not a
	// multiple of 256, it is rounded up to the next multiple of 256. For example, 300 is rounded up
	// to 512.
	PrivateEndpointMaxHostCount *int `mandatory:"false" json:"privateEndpointMaxHostCount"`

	// An array of network security group OCIDs.
	PrivateEndpointNsgIds []string `mandatory:"false" json:"privateEndpointNsgIds"`

	// The OCID of a private endpoint.
	PrivateEndpointId *string `mandatory:"false" json:"privateEndpointId"`

	// The OCID of a subnet.
	PrivateEndpointSubnetId *string `mandatory:"false" json:"privateEndpointSubnetId"`

	// The duration of the run in milliseconds.
	RunDurationInMilliseconds *int64 `mandatory:"false" json:"runDurationInMilliseconds"`

	// The total number of oCPU requested by the run.
	TotalOCpu *int `mandatory:"false" json:"totalOCpu"`

	// The Spark application processing type.
	Type ApplicationTypeEnum `mandatory:"false" json:"type,omitempty"`

	// An Oracle Cloud Infrastructure URI of the bucket to be used as default warehouse directory
	// for BATCH SQL runs.
	// See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat.
	WarehouseBucketUri *string `mandatory:"false" json:"warehouseBucketUri"`

	// The maximum duration in minutes for which an Application should run. Data Flow Run would be terminated
	// once it reaches this duration from the time it transitions to `IN_PROGRESS` state.
	MaxDurationInMinutes *int64 `mandatory:"false" json:"maxDurationInMinutes"`

	// The timeout value in minutes used to manage Runs. A Run would be stopped after inactivity for this amount of time period.
	// Note: This parameter is currently only applicable for Runs of type `SESSION`. Default value is 2880 minutes (2 days)
	IdleTimeoutInMinutes *int64 `mandatory:"false" json:"idleTimeoutInMinutes"`
}

func (m Run) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Run) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingApplicationLanguageEnum(string(m.Language)); !ok && m.Language != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Language: %s. Supported values are: %s.", m.Language, strings.Join(GetApplicationLanguageEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRunLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetRunLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingApplicationTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetApplicationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
