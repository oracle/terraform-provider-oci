// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"github.com/oracle/oci-go-sdk/common"
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

	// The date and time a application was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time a application was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-04-03T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// An Oracle Cloud Infrastructure URI of an archive.zip file containing custom dependencies that may be used to support the execution a Python, Java, or Scala application.
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

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The detailed messages about the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// An Oracle Cloud Infrastructure URI of the bucket where the Spark job logs are to be uploaded.
	// See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat.
	LogsBucketUri *string `mandatory:"false" json:"logsBucketUri"`

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

	// An Oracle Cloud Infrastructure URI of the bucket to be used as default warehouse directory
	// for BATCH SQL runs.
	// See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat.
	WarehouseBucketUri *string `mandatory:"false" json:"warehouseBucketUri"`
}

func (m Run) String() string {
	return common.PointerString(m)
}
