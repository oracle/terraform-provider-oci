// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ExadataStorageServerStatisticsSummary Storage server details and statistics.
type ExadataStorageServerStatisticsSummary struct {
	ResourceDetails *StorageServerDetails `mandatory:"true" json:"resourceDetails"`

	CurrentStatistics *ExadataInsightResourceStatistics `mandatory:"true" json:"currentStatistics"`
}

func (m ExadataStorageServerStatisticsSummary) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ExadataStorageServerStatisticsSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExadataStorageServerStatisticsSummary ExadataStorageServerStatisticsSummary
	s := struct {
		DiscriminatorParam string `json:"exadataResourceType"`
		MarshalTypeExadataStorageServerStatisticsSummary
	}{
		"STORAGE_SERVER",
		(MarshalTypeExadataStorageServerStatisticsSummary)(m),
	}

	return json.Marshal(&s)
}
