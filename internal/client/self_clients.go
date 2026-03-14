// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_self "github.com/oracle/oci-go-sdk/v65/self"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_self.PartnerIntegerationClient", &OracleClient{InitClientFn: initSelfPartnerIntegerationClient})
	RegisterOracleClient("oci_self.SelfSubscriptionClient", &OracleClient{InitClientFn: initSelfSubscriptionClient})
}

func initSelfPartnerIntegerationClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_self.NewPartnerIntegerationClientWithConfigurationProvider(configProvider)
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
	client.Host = "https://self.us-ashburn-1.oci.oc-test.com/"
	return &client, nil
}

func (m *OracleClients) PartnerIntegerationClient() *oci_self.PartnerIntegerationClient {
	return m.GetClient("oci_self.PartnerIntegerationClient").(*oci_self.PartnerIntegerationClient)
}

func initSelfSubscriptionClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_self.NewSubscriptionClientWithConfigurationProvider(configProvider)
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
	client.Host = "https://self.us-ashburn-1.oci.oc-test.com/"
	return &client, nil
}

func (m *OracleClients) SelfSubscriptionClient() *oci_self.SubscriptionClient {
	return m.GetClient("oci_self.SelfSubscriptionClient").(*oci_self.SubscriptionClient)
}
