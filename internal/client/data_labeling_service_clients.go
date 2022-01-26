// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_data_labeling_service "github.com/oracle/oci-go-sdk/v56/datalabelingservice"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
)

func init() {
	RegisterOracleClient("oci_data_labeling_service.DataLabelingManagementClient", &OracleClient{InitClientFn: initDatalabelingserviceDataLabelingManagementClient})
}

func initDatalabelingserviceDataLabelingManagementClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_data_labeling_service.NewDataLabelingManagementClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) DataLabelingManagementClient() *oci_data_labeling_service.DataLabelingManagementClient {
	return m.GetClient("oci_data_labeling_service.DataLabelingManagementClient").(*oci_data_labeling_service.DataLabelingManagementClient)
}
