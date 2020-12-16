// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_object_storage "github.com/oracle/oci-go-sdk/v31/objectstorage"

	oci_common "github.com/oracle/oci-go-sdk/v31/common"
)

func init() {
	RegisterOracleClient("oci_object_storage.ObjectStorageClient", &OracleClient{initClientFn: initObjectstorageObjectStorageClient})
}

func initObjectstorageObjectStorageClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_object_storage.NewObjectStorageClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) objectStorageClient() *oci_object_storage.ObjectStorageClient {
	return m.GetClient("oci_object_storage.ObjectStorageClient").(*oci_object_storage.ObjectStorageClient)
}
