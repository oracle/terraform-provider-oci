// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"fmt"

	oci_apigateway "github.com/oracle/oci-go-sdk/apigateway"
	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_functions "github.com/oracle/oci-go-sdk/functions"
	oci_kms "github.com/oracle/oci-go-sdk/keymanagement"
	oci_work_requests "github.com/oracle/oci-go-sdk/workrequests"
)

var oracleClients *OracleClients

func RegisterOracleClient(name string, client *OracleClient) {
	if oracleClients == nil {
		oracleClients = &OracleClients{
			configuration: make(map[string]string),
			clientMap:     make(map[string]*OracleClient),
		}
	}
	oracleClients.clientMap[name] = client
}

type InitSdkClientFn func(oci_common.ConfigurationProvider, ConfigureClient) (interface{}, error)

type OracleClient struct {
	sdkClient    interface{}
	initClientFn InitSdkClientFn
}

type OracleClients struct {
	configuration             map[string]string
	clientMap                 map[string]*OracleClient
	gatewayWorkRequestsClient *oci_apigateway.WorkRequestsClient
	workRequestClient         *oci_work_requests.WorkRequestClient
}

func (m *OracleClients) GetClient(name string) interface{} {
	return m.clientMap[name].sdkClient
}

// The following clients require special endpoint information that is only known at Terraform apply time; so they
// create duplicate clients reusing the same configuration provider as the initialized client and adding the endpoint
// here.
func (m *OracleClients) FunctionsInvokeClient(endpoint string) (*oci_functions.FunctionsInvokeClient, error) {
	if client, err := oci_functions.NewFunctionsInvokeClientWithConfigurationProvider(*m.functionsInvokeClient().ConfigurationProvider(), endpoint); err == nil {
		if err = configureClient(&client.BaseClient); err != nil {
			return nil, err
		}
		return &client, nil
	} else {
		return nil, err
	}
}

func (m *OracleClients) KmsCryptoClient(endpoint string) (*oci_kms.KmsCryptoClient, error) {
	if client, err := oci_kms.NewKmsCryptoClientWithConfigurationProvider(*m.kmsCryptoClient().ConfigurationProvider(), endpoint); err == nil {
		if err = configureClient(&client.BaseClient); err != nil {
			return nil, err
		}
		return &client, nil
	} else {
		return nil, err
	}
}

func (m *OracleClients) KmsManagementClient(endpoint string) (*oci_kms.KmsManagementClient, error) {
	if client, err := oci_kms.NewKmsManagementClientWithConfigurationProvider(*m.kmsManagementClient().ConfigurationProvider(), endpoint); err == nil {
		if err = configureClient(&client.BaseClient); err != nil {
			return nil, err
		}
		return &client, nil
	} else {
		return nil, err
	}
}

func createSDKClients(clients *OracleClients, configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (err error) {
	if clients == nil || len(clients.clientMap) == 0 {
		return fmt.Errorf("there are no clients to create")
	}

	for serviceName, client := range clients.clientMap {
		if client.initClientFn != nil {
			client.sdkClient, err = client.initClientFn(configProvider, configureClient)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("unable to initialize '%s' client", serviceName)
		}
	}

	gatewayWorkRequestsClient, err := oci_apigateway.NewWorkRequestsClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&gatewayWorkRequestsClient.BaseClient)
	if err != nil {
		return
	}
	clients.gatewayWorkRequestsClient = &gatewayWorkRequestsClient

	workRequestClient, err := oci_work_requests.NewWorkRequestClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&workRequestClient.BaseClient)
	if err != nil {
		return
	}
	clients.workRequestClient = &workRequestClient

	return
}
