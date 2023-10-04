// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// BenchmarkDetails Benchmark test configuration detail.
type BenchmarkDetails struct {

	// Private endpoint of cluster to be tested.
	ClusterPrivateEndpoint *string `mandatory:"false" json:"clusterPrivateEndpoint"`

	// UserId of cluster trying to access.
	ClusterUserId *string `mandatory:"false" json:"clusterUserId"`

	// Password of cluster trying to access.
	ClusterPassword *string `mandatory:"false" json:"clusterPassword"`

	// Subnet of cluster(Preferably customer-vcn and public-subnet)
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// Id of custom image for instance creation
	CustomImageId *string `mandatory:"false" json:"customImageId"`

	// CompartmentId that cluster is on
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Throughput of ingestion test
	ThroughputIngestion *int `mandatory:"false" json:"throughputIngestion"`

	// Throughput of query
	ThroughputQuery *int `mandatory:"false" json:"throughputQuery"`

	// Ramp up time it takes to reach throughput level
	RampUpDurationIngestion *int `mandatory:"false" json:"rampUpDurationIngestion"`

	// Ramp up time it takes to reach throughput level
	RampUpDurationQuery *int `mandatory:"false" json:"rampUpDurationQuery"`

	// Size of batch query to be returned for query testing
	QueryBatchSize *int `mandatory:"false" json:"queryBatchSize"`

	// Duration of Time it takes to run for Ingestion test
	ExecutionTimeIngestion *int `mandatory:"false" json:"executionTimeIngestion"`

	// Duration of Time it takes to run for Query Test
	ExecutionTimeQuery *int `mandatory:"false" json:"executionTimeQuery"`

	// Run Ingestion testing
	IsRunningIngestionTest *bool `mandatory:"false" json:"isRunningIngestionTest"`

	// Run Query testing
	IsRunningQueryTest *bool `mandatory:"false" json:"isRunningQueryTest"`
}

func (m BenchmarkDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BenchmarkDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
