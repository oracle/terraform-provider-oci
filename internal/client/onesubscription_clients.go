// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_onesubscription "github.com/oracle/oci-go-sdk/v65/onesubscription"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_onesubscription.BillingScheduleRegionalClient", &OracleClient{InitClientFn: initOnesubscriptionBillingScheduleClient})
	RegisterOracleClient("oci_onesubscription.CommitmentRegionalClient", &OracleClient{InitClientFn: initOnesubscriptionCommitmentClient})
	RegisterOracleClient("oci_onesubscription.ComputedUsageRegionalClient", &OracleClient{InitClientFn: initOnesubscriptionComputedUsageClient})
	RegisterOracleClient("oci_onesubscription.InvoiceSummaryRegionalClient", &OracleClient{InitClientFn: initOnesubscriptionInvoiceSummaryClient})
	RegisterOracleClient("oci_onesubscription.OrganizationSubscriptionRegionalClient", &OracleClient{InitClientFn: initOnesubscriptionOrganizationSubscriptionClient})
	RegisterOracleClient("oci_onesubscription.RatecardRegionalClient", &OracleClient{InitClientFn: initOnesubscriptionRatecardClient})
	RegisterOracleClient("oci_onesubscription.SubscribedServiceRegionalClient", &OracleClient{InitClientFn: initOnesubscriptionSubscribedServiceClient})
	RegisterOracleClient("oci_onesubscription.SubscriptionRegionalClient", &OracleClient{InitClientFn: initOnesubscriptionSubscriptionClient})
}

func initOnesubscriptionBillingScheduleClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_onesubscription.NewBillingScheduleClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) BillingScheduleRegionalClient() *oci_onesubscription.BillingScheduleClient {
	return m.GetClient("oci_onesubscription.BillingScheduleRegionalClient").(*oci_onesubscription.BillingScheduleClient)
}

func initOnesubscriptionCommitmentClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_onesubscription.NewCommitmentClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) CommitmentRegionalClient() *oci_onesubscription.CommitmentClient {
	return m.GetClient("oci_onesubscription.CommitmentRegionalClient").(*oci_onesubscription.CommitmentClient)
}

func initOnesubscriptionComputedUsageClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_onesubscription.NewComputedUsageClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ComputedUsageRegionalClient() *oci_onesubscription.ComputedUsageClient {
	return m.GetClient("oci_onesubscription.ComputedUsageRegionalClient").(*oci_onesubscription.ComputedUsageClient)
}

func initOnesubscriptionInvoiceSummaryClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_onesubscription.NewInvoiceSummaryClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) InvoiceSummaryRegionalClient() *oci_onesubscription.InvoiceSummaryClient {
	return m.GetClient("oci_onesubscription.InvoiceSummaryRegionalClient").(*oci_onesubscription.InvoiceSummaryClient)
}

func initOnesubscriptionOrganizationSubscriptionClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_onesubscription.NewOrganizationSubscriptionClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) OrganizationSubscriptionRegionalClient() *oci_onesubscription.OrganizationSubscriptionClient {
	return m.GetClient("oci_onesubscription.OrganizationSubscriptionRegionalClient").(*oci_onesubscription.OrganizationSubscriptionClient)
}

func initOnesubscriptionRatecardClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_onesubscription.NewRatecardClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) RatecardRegionalClient() *oci_onesubscription.RatecardClient {
	return m.GetClient("oci_onesubscription.RatecardRegionalClient").(*oci_onesubscription.RatecardClient)
}

func initOnesubscriptionSubscribedServiceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_onesubscription.NewSubscribedServiceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) SubscribedServiceRegionalClient() *oci_onesubscription.SubscribedServiceClient {
	return m.GetClient("oci_onesubscription.SubscribedServiceRegionalClient").(*oci_onesubscription.SubscribedServiceClient)
}

func initOnesubscriptionSubscriptionClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_onesubscription.NewSubscriptionClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) SubscriptionRegionalClient() *oci_onesubscription.SubscriptionClient {
	return m.GetClient("oci_onesubscription.SubscriptionRegionalClient").(*oci_onesubscription.SubscriptionClient)
}
