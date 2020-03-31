// Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

func init() {
	RegisterOracleClient("oci_object_storage.ObjectStorageClient", &OracleClient{initClientFn: initObjectstorageObjectStorageClient})
}

func initObjectstorageObjectStorageClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_object_storage.NewObjectStorageClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) objectStorageClient() *oci_object_storage.ObjectStorageClient {
	return m.GetClient("oci_object_storage.ObjectStorageClient").(*oci_object_storage.ObjectStorageClient)
}
