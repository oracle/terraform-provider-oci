// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_tenantmanagercontrolplane "github.com/oracle/oci-go-sdk/v65/tenantmanagercontrolplane"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_tenantmanagercontrolplane.DomainClient", &OracleClient{InitClientFn: initTenantmanagercontrolplaneDomainClient})
	RegisterOracleClient("oci_tenantmanagercontrolplane.DomainGovernanceClient", &OracleClient{InitClientFn: initTenantmanagercontrolplaneDomainGovernanceClient})
	RegisterOracleClient("oci_tenantmanagercontrolplane.LinkClient", &OracleClient{InitClientFn: initTenantmanagercontrolplaneLinkClient})
	RegisterOracleClient("oci_tenantmanagercontrolplane.OrganizationClient", &OracleClient{InitClientFn: initTenantmanagercontrolplaneOrganizationClient})
	RegisterOracleClient("oci_tenantmanagercontrolplane.RecipientInvitationClient", &OracleClient{InitClientFn: initTenantmanagercontrolplaneRecipientInvitationClient})
	RegisterOracleClient("oci_tenantmanagercontrolplane.SenderInvitationClient", &OracleClient{InitClientFn: initTenantmanagercontrolplaneSenderInvitationClient})
	RegisterOracleClient("oci_tenantmanagercontrolplane.SubscriptionClient", &OracleClient{InitClientFn: initTenantmanagercontrolplaneSubscriptionClient})
	RegisterOracleClient("oci_tenantmanagercontrolplane.WorkRequestClient", &OracleClient{InitClientFn: initTenantmanagercontrolplaneWorkRequestClient})
}

func initTenantmanagercontrolplaneDomainClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_tenantmanagercontrolplane.NewDomainClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DomainClient() *oci_tenantmanagercontrolplane.DomainClient {
	return m.GetClient("oci_tenantmanagercontrolplane.DomainClient").(*oci_tenantmanagercontrolplane.DomainClient)
}

func initTenantmanagercontrolplaneDomainGovernanceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_tenantmanagercontrolplane.NewDomainGovernanceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DomainGovernanceClient() *oci_tenantmanagercontrolplane.DomainGovernanceClient {
	return m.GetClient("oci_tenantmanagercontrolplane.DomainGovernanceClient").(*oci_tenantmanagercontrolplane.DomainGovernanceClient)
}

func initTenantmanagercontrolplaneLinkClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_tenantmanagercontrolplane.NewLinkClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) LinkClient() *oci_tenantmanagercontrolplane.LinkClient {
	return m.GetClient("oci_tenantmanagercontrolplane.LinkClient").(*oci_tenantmanagercontrolplane.LinkClient)
}

func initTenantmanagercontrolplaneOrganizationClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_tenantmanagercontrolplane.NewOrganizationClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) OrganizationClient() *oci_tenantmanagercontrolplane.OrganizationClient {
	return m.GetClient("oci_tenantmanagercontrolplane.OrganizationClient").(*oci_tenantmanagercontrolplane.OrganizationClient)
}

func initTenantmanagercontrolplaneRecipientInvitationClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_tenantmanagercontrolplane.NewRecipientInvitationClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) RecipientInvitationClient() *oci_tenantmanagercontrolplane.RecipientInvitationClient {
	return m.GetClient("oci_tenantmanagercontrolplane.RecipientInvitationClient").(*oci_tenantmanagercontrolplane.RecipientInvitationClient)
}

func initTenantmanagercontrolplaneSenderInvitationClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_tenantmanagercontrolplane.NewSenderInvitationClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) SenderInvitationClient() *oci_tenantmanagercontrolplane.SenderInvitationClient {
	return m.GetClient("oci_tenantmanagercontrolplane.SenderInvitationClient").(*oci_tenantmanagercontrolplane.SenderInvitationClient)
}

func initTenantmanagercontrolplaneSubscriptionClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_tenantmanagercontrolplane.NewSubscriptionClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) OrganizationsSubscriptionClient() *oci_tenantmanagercontrolplane.SubscriptionClient {
	return m.GetClient("oci_tenantmanagercontrolplane.SubscriptionClient").(*oci_tenantmanagercontrolplane.SubscriptionClient)
}

func initTenantmanagercontrolplaneWorkRequestClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_tenantmanagercontrolplane.NewWorkRequestClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) TenantmanagercontrolplaneWorkRequestClient() *oci_tenantmanagercontrolplane.WorkRequestClient {
	return m.GetClient("oci_tenantmanagercontrolplane.WorkRequestClient").(*oci_tenantmanagercontrolplane.WorkRequestClient)
}
