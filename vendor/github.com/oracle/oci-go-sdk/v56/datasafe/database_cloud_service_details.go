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

// DatabaseCloudServiceDetails The details of the Oracle Database Cloud Service to be registered as a target database in Data Safe.
type DatabaseCloudServiceDetails struct {

	// The OCID of the VM cluster in which the database is running.
	VmClusterId *string `mandatory:"false" json:"vmClusterId"`

	// The OCID of the cloud database system registered as a target database in Data Safe.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

	// The database service name.
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// The infrastructure type the database is running on.
	InfrastructureType InfrastructureTypeEnum `mandatory:"true" json:"infrastructureType"`
}

//GetInfrastructureType returns InfrastructureType
func (m DatabaseCloudServiceDetails) GetInfrastructureType() InfrastructureTypeEnum {
	return m.InfrastructureType
}

func (m DatabaseCloudServiceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DatabaseCloudServiceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseCloudServiceDetails DatabaseCloudServiceDetails
	s := struct {
		DiscriminatorParam string `json:"databaseType"`
		MarshalTypeDatabaseCloudServiceDetails
	}{
		"DATABASE_CLOUD_SERVICE",
		(MarshalTypeDatabaseCloudServiceDetails)(m),
	}

	return json.Marshal(&s)
}
