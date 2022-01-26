// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// InstalledDatabaseDetails The details of the database running on-premises or on a compute instance.
type InstalledDatabaseDetails struct {

	// The OCID of the compute instance on which the database is running.
	InstanceId *string `mandatory:"false" json:"instanceId"`

	// The list of database host IP Addresses. Fully qualified domain names can be used if connectionType is 'ONPREM_CONNECTOR'.
	IpAddresses []string `mandatory:"false" json:"ipAddresses"`

	// The port number of the database listener.
	ListenerPort *int `mandatory:"false" json:"listenerPort"`

	// The service name of the database registered as target database.
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// The infrastructure type the database is running on.
	InfrastructureType InfrastructureTypeEnum `mandatory:"true" json:"infrastructureType"`
}

//GetInfrastructureType returns InfrastructureType
func (m InstalledDatabaseDetails) GetInfrastructureType() InfrastructureTypeEnum {
	return m.InfrastructureType
}

func (m InstalledDatabaseDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InstalledDatabaseDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstalledDatabaseDetails InstalledDatabaseDetails
	s := struct {
		DiscriminatorParam string `json:"databaseType"`
		MarshalTypeInstalledDatabaseDetails
	}{
		"INSTALLED_DATABASE",
		(MarshalTypeInstalledDatabaseDetails)(m),
	}

	return json.Marshal(&s)
}
