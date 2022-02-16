// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_osp_gateway "github.com/oracle/oci-go-sdk/v58/ospgateway"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
)

func init() {
	RegisterOracleClient("oci_osp_gateway.InvoiceServiceClient", &OracleClient{InitClientFn: initOspgatewayInvoiceServiceClient})
	RegisterOracleClient("oci_osp_gateway.SubscriptionServiceClient", &OracleClient{InitClientFn: initOspgatewaySubscriptionServiceClient})
}

func initOspgatewayInvoiceServiceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_osp_gateway.NewInvoiceServiceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) InvoiceServiceClient() *oci_osp_gateway.InvoiceServiceClient {
	return m.GetClient("oci_osp_gateway.InvoiceServiceClient").(*oci_osp_gateway.InvoiceServiceClient)
}

func initOspgatewaySubscriptionServiceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_osp_gateway.NewSubscriptionServiceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) SubscriptionServiceClient() *oci_osp_gateway.SubscriptionServiceClient {
	return m.GetClient("oci_osp_gateway.SubscriptionServiceClient").(*oci_osp_gateway.SubscriptionServiceClient)
}
