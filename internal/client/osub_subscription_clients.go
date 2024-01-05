// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_osub_subscription "github.com/oracle/oci-go-sdk/v65/osubsubscription"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_osub_subscription.CommitmentClient", &OracleClient{InitClientFn: initOsubsubscriptionCommitmentClient})
	RegisterOracleClient("oci_osub_subscription.RatecardClient", &OracleClient{InitClientFn: initOsubsubscriptionRatecardClient})
	RegisterOracleClient("oci_osub_subscription.SubscriptionClient", &OracleClient{InitClientFn: initOsubsubscriptionSubscriptionClient})
}

func initOsubsubscriptionCommitmentClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_osub_subscription.NewCommitmentClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) CommitmentClient() *oci_osub_subscription.CommitmentClient {
	return m.GetClient("oci_osub_subscription.CommitmentClient").(*oci_osub_subscription.CommitmentClient)
}

func initOsubsubscriptionRatecardClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_osub_subscription.NewRatecardClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) RatecardClient() *oci_osub_subscription.RatecardClient {
	return m.GetClient("oci_osub_subscription.RatecardClient").(*oci_osub_subscription.RatecardClient)
}

func initOsubsubscriptionSubscriptionClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_osub_subscription.NewSubscriptionClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) SubscriptionClient() *oci_osub_subscription.SubscriptionClient {
	return m.GetClient("oci_osub_subscription.SubscriptionClient").(*oci_osub_subscription.SubscriptionClient)
}
