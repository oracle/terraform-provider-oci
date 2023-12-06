// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_ocvp.ClusterClient", &OracleClient{InitClientFn: initOcvpClusterClient})
	RegisterOracleClient("oci_ocvp.EsxiHostClient", &OracleClient{InitClientFn: initOcvpEsxiHostClient})
	RegisterOracleClient("oci_ocvp.WorkRequestClient", &OracleClient{InitClientFn: initOcvpWorkRequestClient})
	RegisterOracleClient("oci_ocvp.SddcClient", &OracleClient{InitClientFn: initOcvpSddcClient})
}

func initOcvpClusterClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_ocvp.NewClusterClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ClusterClient() *oci_ocvp.ClusterClient {
	return m.GetClient("oci_ocvp.ClusterClient").(*oci_ocvp.ClusterClient)
}

func initOcvpEsxiHostClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_ocvp.NewEsxiHostClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) EsxiHostClient() *oci_ocvp.EsxiHostClient {
	return m.GetClient("oci_ocvp.EsxiHostClient").(*oci_ocvp.EsxiHostClient)
}

func initOcvpWorkRequestClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_ocvp.NewWorkRequestClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) OcvpWorkRequestClient() *oci_ocvp.WorkRequestClient {
	return m.GetClient("oci_ocvp.WorkRequestClient").(*oci_ocvp.WorkRequestClient)
}

func initOcvpSddcClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_ocvp.NewSddcClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) SddcClient() *oci_ocvp.SddcClient {
	return m.GetClient("oci_ocvp.SddcClient").(*oci_ocvp.SddcClient)
}
