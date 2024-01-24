// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_osp_gateway "github.com/oracle/oci-go-sdk/v65/ospgateway"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_osp_gateway.AddressRuleServiceClient", &OracleClient{InitClientFn: initOspgatewayAddressRuleServiceClient})
	RegisterOracleClient("oci_osp_gateway.AddressServiceClient", &OracleClient{InitClientFn: initOspgatewayAddressServiceClient})
	RegisterOracleClient("oci_osp_gateway.InvoiceServiceClient", &OracleClient{InitClientFn: initOspgatewayInvoiceServiceClient})
	RegisterOracleClient("oci_osp_gateway.SubscriptionServiceClient", &OracleClient{InitClientFn: initOspgatewaySubscriptionServiceClient})
}

func initOspgatewayAddressRuleServiceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_osp_gateway.NewAddressRuleServiceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) AddressRuleServiceClient() *oci_osp_gateway.AddressRuleServiceClient {
	return m.GetClient("oci_osp_gateway.AddressRuleServiceClient").(*oci_osp_gateway.AddressRuleServiceClient)
}

func initOspgatewayAddressServiceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_osp_gateway.NewAddressServiceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) AddressServiceClient() *oci_osp_gateway.AddressServiceClient {
	return m.GetClient("oci_osp_gateway.AddressServiceClient").(*oci_osp_gateway.AddressServiceClient)
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
