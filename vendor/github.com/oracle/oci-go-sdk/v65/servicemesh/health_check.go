// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HealthCheck If specified, health check is performed on cluster members (hosts) of the Virtual Deployment. If a host fails
// the health check, traffic will not be routed to the host unless health check passes. If no configuration is
// specified, no health check will be done and all hosts will be considered healthy at all times. When using
// MESH_REGISTRY service discovery, setting health check is recommended so that traffic is not routed to unhealthy
// hosts.
type HealthCheck interface {

	// The time to wait for a health check response. If the timeout is reached the health check attempt
	// will be considered a failure.
	GetTimeoutInMs() *int64

	// The interval between health checks.
	GetIntervalInMs() *int64

	// The number of unhealthy health checks required before a host is marked unhealthy.
	GetUnhealthyThreshold() *int

	// The number of healthy health checks required before a host is marked healthy. Note that during startup,
	// only a single successful health check is required to mark a host healthy.
	GetHealthyThreshold() *int
}

type healthcheck struct {
	JsonData           []byte
	TimeoutInMs        *int64 `mandatory:"true" json:"timeoutInMs"`
	IntervalInMs       *int64 `mandatory:"true" json:"intervalInMs"`
	UnhealthyThreshold *int   `mandatory:"false" json:"unhealthyThreshold"`
	HealthyThreshold   *int   `mandatory:"false" json:"healthyThreshold"`
	Type               string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *healthcheck) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerhealthcheck healthcheck
	s := struct {
		Model Unmarshalerhealthcheck
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TimeoutInMs = s.Model.TimeoutInMs
	m.IntervalInMs = s.Model.IntervalInMs
	m.UnhealthyThreshold = s.Model.UnhealthyThreshold
	m.HealthyThreshold = s.Model.HealthyThreshold
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *healthcheck) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "TCP":
		mm := TcpHealthCheck{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HTTP":
		mm := HttpHealthCheck{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GRPC":
		mm := GrpcHealthCheck{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for HealthCheck: %s.", m.Type)
		return *m, nil
	}
}

//GetTimeoutInMs returns TimeoutInMs
func (m healthcheck) GetTimeoutInMs() *int64 {
	return m.TimeoutInMs
}

//GetIntervalInMs returns IntervalInMs
func (m healthcheck) GetIntervalInMs() *int64 {
	return m.IntervalInMs
}

//GetUnhealthyThreshold returns UnhealthyThreshold
func (m healthcheck) GetUnhealthyThreshold() *int {
	return m.UnhealthyThreshold
}

//GetHealthyThreshold returns HealthyThreshold
func (m healthcheck) GetHealthyThreshold() *int {
	return m.HealthyThreshold
}

func (m healthcheck) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m healthcheck) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HealthCheckTypeEnum Enum with underlying type: string
type HealthCheckTypeEnum string

// Set of constants representing the allowable values for HealthCheckTypeEnum
const (
	HealthCheckTypeTcp  HealthCheckTypeEnum = "TCP"
	HealthCheckTypeHttp HealthCheckTypeEnum = "HTTP"
	HealthCheckTypeGrpc HealthCheckTypeEnum = "GRPC"
)

var mappingHealthCheckTypeEnum = map[string]HealthCheckTypeEnum{
	"TCP":  HealthCheckTypeTcp,
	"HTTP": HealthCheckTypeHttp,
	"GRPC": HealthCheckTypeGrpc,
}

var mappingHealthCheckTypeEnumLowerCase = map[string]HealthCheckTypeEnum{
	"tcp":  HealthCheckTypeTcp,
	"http": HealthCheckTypeHttp,
	"grpc": HealthCheckTypeGrpc,
}

// GetHealthCheckTypeEnumValues Enumerates the set of values for HealthCheckTypeEnum
func GetHealthCheckTypeEnumValues() []HealthCheckTypeEnum {
	values := make([]HealthCheckTypeEnum, 0)
	for _, v := range mappingHealthCheckTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetHealthCheckTypeEnumStringValues Enumerates the set of values in String for HealthCheckTypeEnum
func GetHealthCheckTypeEnumStringValues() []string {
	return []string{
		"TCP",
		"HTTP",
		"GRPC",
	}
}

// GetMappingHealthCheckTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHealthCheckTypeEnum(val string) (HealthCheckTypeEnum, bool) {
	enum, ok := mappingHealthCheckTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
