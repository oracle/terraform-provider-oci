// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	"fmt"
	"strings"

	oci_functions "github.com/oracle/oci-go-sdk/v58/functions"

	oci_kms "github.com/oracle/oci-go-sdk/v58/keymanagement"

	"github.com/terraform-providers/terraform-provider-oci/internal/globalvar"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_work_requests "github.com/oracle/oci-go-sdk/v58/workrequests"

	utils "github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var OracleClientRegistrationsVar *OracleClientRegistrations // This is a global registration for all oracle clients. This is invariant information about all clients regardless of region

func RegisterOracleClient(name string, client *OracleClient) {
	if OracleClientRegistrationsVar == nil {
		OracleClientRegistrationsVar = &OracleClientRegistrations{
			RegisteredClients: make(map[string]*OracleClient),
		}
	}
	OracleClientRegistrationsVar.RegisteredClients[name] = client
}

type ConfigureClient func(client *oci_common.BaseClient) error

var ConfigureClientVar ConfigureClient // global fn ref used to configure all clients initially and others later on

type InitSdkClientFn func(oci_common.ConfigurationProvider, ConfigureClient, ServiceClientOverrides) (interface{}, error)

type OracleClientRegistrations struct {
	RegisteredClients map[string]*OracleClient
}

type ServiceClientOverrides struct {
	HostUrlOverride string
}

type OracleClient struct {
	InitClientFn InitSdkClientFn
}

type OracleClients struct {
	Configuration     map[string]string
	SdkClientMap      map[string]interface{}
	WorkRequestClient *oci_work_requests.WorkRequestClient
}

func (m *OracleClients) GetClient(name string) interface{} {
	return m.SdkClientMap[name]
}

// The following clients require special endpoint information that is only known at Terraform apply time; so they
// Create duplicate clients reusing the same Configuration provider as the initialized client and adding the endpoint
// here.
func (m *OracleClients) FunctionsInvokeClientWithEndpoint(endpoint string) (*oci_functions.FunctionsInvokeClient, error) {
	if client, err := oci_functions.NewFunctionsInvokeClientWithConfigurationProvider(*m.FunctionsInvokeClient().ConfigurationProvider(), endpoint); err == nil {
		if err = ConfigureClientVar(&client.BaseClient); err != nil {
			return nil, err
		}
		return &client, nil
	} else {
		return nil, err
	}
}
func (m *OracleClients) KmsCryptoClientWithEndpoint(endpoint string) (*oci_kms.KmsCryptoClient, error) {
	if client, err := oci_kms.NewKmsCryptoClientWithConfigurationProvider(*m.KmsCryptoClient().ConfigurationProvider(), endpoint); err == nil {
		if err = ConfigureClientVar(&client.BaseClient); err != nil {
			return nil, err
		}
		return &client, nil
	} else {
		return nil, err
	}
}

func (m *OracleClients) KmsManagementClientWithEndpoint(endpoint string) (*oci_kms.KmsManagementClient, error) {
	if client, err := oci_kms.NewKmsManagementClientWithConfigurationProvider(*m.KmsManagementClient().ConfigurationProvider(), endpoint); err == nil {
		if err = ConfigureClientVar(&client.BaseClient); err != nil {
			return nil, err
		}
		return &client, nil
	} else {
		return nil, err
	}
}

func getClientHostOverrides() map[string]string {
	// Get the host URL override for clients
	clientHostOverrides := make(map[string]string)
	clientHostOverridesString := utils.GetEnvSettingWithBlankDefault(globalvar.ClientHostOverridesEnv)
	if clientHostOverridesString == "" {
		return clientHostOverrides
	}

	clientHostFlags := strings.Split(clientHostOverridesString, globalvar.ColonDelimiter)
	for _, item := range clientHostFlags {
		clientNameHost := strings.Split(item, globalvar.EqualToOperatorDelimiter)
		if clientNameHost == nil || len(clientNameHost) != 2 {
			continue
		}
		clientHostOverrides[clientNameHost[0]] = clientNameHost[1]
	}
	return clientHostOverrides
}

func CreateSDKClients(clients *OracleClients, configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (err error) {
	if OracleClientRegistrationsVar == nil || len(OracleClientRegistrationsVar.RegisteredClients) == 0 {
		return fmt.Errorf("there are no clients to Create")
	}

	clientHostOverrides := getClientHostOverrides()
	for serviceName, clientRegistration := range OracleClientRegistrationsVar.RegisteredClients {
		if clientRegistration.InitClientFn != nil {
			serviceClientOverrides := ServiceClientOverrides{}
			// apply client host override
			if host, ok := clientHostOverrides[serviceName]; ok {
				serviceClientOverrides.HostUrlOverride = host
			}

			clients.SdkClientMap[serviceName], err = clientRegistration.InitClientFn(configProvider, configureClient, serviceClientOverrides)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("unable to initialize '%s' client", serviceName)
		}
	}

	workRequestClient, err := oci_work_requests.NewWorkRequestClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&workRequestClient.BaseClient)
	if err != nil {
		return
	}
	clients.WorkRequestClient = &workRequestClient

	return
}
