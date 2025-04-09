// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_lustre_file_storage "github.com/oracle/oci-go-sdk/v65/lustrefilestorage"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_lustre_file_storage.LustreFileStorageClient", &OracleClient{InitClientFn: initLustrefilestorageLustreFileStorageClient})
}

func initLustrefilestorageLustreFileStorageClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_lustre_file_storage.NewLustreFileStorageClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) LustreFileStorageClient() *oci_lustre_file_storage.LustreFileStorageClient {
	return m.GetClient("oci_lustre_file_storage.LustreFileStorageClient").(*oci_lustre_file_storage.LustreFileStorageClient)
}
