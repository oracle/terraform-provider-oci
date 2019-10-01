// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	oci_audit "github.com/oracle/oci-go-sdk/audit"
	oci_auto_scaling "github.com/oracle/oci-go-sdk/autoscaling"
	oci_budget "github.com/oracle/oci-go-sdk/budget"
	oci_containerengine "github.com/oracle/oci-go-sdk/containerengine"
	oci_core "github.com/oracle/oci-go-sdk/core"
	oci_database "github.com/oracle/oci-go-sdk/database"
	oci_dns "github.com/oracle/oci-go-sdk/dns"
	oci_email "github.com/oracle/oci-go-sdk/email"
	oci_events "github.com/oracle/oci-go-sdk/events"
	oci_file_storage "github.com/oracle/oci-go-sdk/filestorage"
	oci_functions "github.com/oracle/oci-go-sdk/functions"
	oci_health_checks "github.com/oracle/oci-go-sdk/healthchecks"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
	oci_kms "github.com/oracle/oci-go-sdk/keymanagement"
	oci_limits "github.com/oracle/oci-go-sdk/limits"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
	oci_monitoring "github.com/oracle/oci-go-sdk/monitoring"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"
	oci_ons "github.com/oracle/oci-go-sdk/ons"
	oci_streaming "github.com/oracle/oci-go-sdk/streaming"
	oci_waas "github.com/oracle/oci-go-sdk/waas"
	oci_work_requests "github.com/oracle/oci-go-sdk/workrequests"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

type OracleClients struct {
	configuration                  map[string]string
	auditClient                    *oci_audit.AuditClient
	autoScalingClient              *oci_auto_scaling.AutoScalingClient
	blockstorageClient             *oci_core.BlockstorageClient
	budgetClient                   *oci_budget.BudgetClient
	computeClient                  *oci_core.ComputeClient
	computeManagementClient        *oci_core.ComputeManagementClient
	containerEngineClient          *oci_containerengine.ContainerEngineClient
	databaseClient                 *oci_database.DatabaseClient
	dnsClient                      *oci_dns.DnsClient
	emailClient                    *oci_email.EmailClient
	eventsClient                   *oci_events.EventsClient
	fileStorageClient              *oci_file_storage.FileStorageClient
	functionsInvokeClient          *oci_functions.FunctionsInvokeClient
	functionsManagementClient      *oci_functions.FunctionsManagementClient
	healthChecksClient             *oci_health_checks.HealthChecksClient
	identityClient                 *oci_identity.IdentityClient
	kmsCryptoClient                *oci_kms.KmsCryptoClient
	kmsManagementClient            *oci_kms.KmsManagementClient
	kmsVaultClient                 *oci_kms.KmsVaultClient
	limitsClient                   *oci_limits.LimitsClient
	loadBalancerClient             *oci_load_balancer.LoadBalancerClient
	monitoringClient               *oci_monitoring.MonitoringClient
	notificationControlPlaneClient *oci_ons.NotificationControlPlaneClient
	notificationDataPlaneClient    *oci_ons.NotificationDataPlaneClient
	objectStorageClient            *oci_object_storage.ObjectStorageClient
	quotasClient                   *oci_limits.QuotasClient
	streamAdminClient              *oci_streaming.StreamAdminClient
	virtualNetworkClient           *oci_core.VirtualNetworkClient
	waasClient                     *oci_waas.WaasClient
	workRequestClient              *oci_work_requests.WorkRequestClient
}

func (m *OracleClients) FunctionsInvokeClient(endpoint string) (*oci_functions.FunctionsInvokeClient, error) {
	if client, err := oci_functions.NewFunctionsInvokeClientWithConfigurationProvider(*m.functionsInvokeClient.ConfigurationProvider(), endpoint); err == nil {
		if err = configureClient(&client.BaseClient); err != nil {
			return nil, err
		}
		return &client, nil
	} else {
		return nil, err
	}
}

func (m *OracleClients) KmsCryptoClient(endpoint string) (*oci_kms.KmsCryptoClient, error) {
	if client, err := oci_kms.NewKmsCryptoClientWithConfigurationProvider(*m.kmsCryptoClient.ConfigurationProvider(), endpoint); err == nil {
		if err = configureClient(&client.BaseClient); err != nil {
			return nil, err
		}
		return &client, nil
	} else {
		return nil, err
	}
}

func (m *OracleClients) KmsManagementClient(endpoint string) (*oci_kms.KmsManagementClient, error) {
	if client, err := oci_kms.NewKmsManagementClientWithConfigurationProvider(*m.kmsManagementClient.ConfigurationProvider(), endpoint); err == nil {
		if err = configureClient(&client.BaseClient); err != nil {
			return nil, err
		}
		return &client, nil
	} else {
		return nil, err
	}
}

func createSDKClients(clients *OracleClients, configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (err error) {

	auditClient, err := oci_audit.NewAuditClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&auditClient.BaseClient)
	if err != nil {
		return
	}
	clients.auditClient = &auditClient

	autoScalingClient, err := oci_auto_scaling.NewAutoScalingClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&autoScalingClient.BaseClient)
	if err != nil {
		return
	}
	clients.autoScalingClient = &autoScalingClient

	blockstorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&blockstorageClient.BaseClient)
	if err != nil {
		return
	}
	clients.blockstorageClient = &blockstorageClient

	budgetClient, err := oci_budget.NewBudgetClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&budgetClient.BaseClient)
	if err != nil {
		return
	}
	clients.budgetClient = &budgetClient

	computeClient, err := oci_core.NewComputeClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&computeClient.BaseClient)
	if err != nil {
		return
	}
	clients.computeClient = &computeClient

	computeManagementClient, err := oci_core.NewComputeManagementClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&computeManagementClient.BaseClient)
	if err != nil {
		return
	}
	clients.computeManagementClient = &computeManagementClient

	containerEngineClient, err := oci_containerengine.NewContainerEngineClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&containerEngineClient.BaseClient)
	if err != nil {
		return
	}
	clients.containerEngineClient = &containerEngineClient

	databaseClient, err := oci_database.NewDatabaseClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&databaseClient.BaseClient)
	if err != nil {
		return
	}
	clients.databaseClient = &databaseClient

	dnsClient, err := oci_dns.NewDnsClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&dnsClient.BaseClient)
	if err != nil {
		return
	}
	clients.dnsClient = &dnsClient

	emailClient, err := oci_email.NewEmailClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&emailClient.BaseClient)
	if err != nil {
		return
	}
	clients.emailClient = &emailClient

	eventsClient, err := oci_events.NewEventsClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&eventsClient.BaseClient)
	if err != nil {
		return
	}
	clients.eventsClient = &eventsClient

	fileStorageClient, err := oci_file_storage.NewFileStorageClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&fileStorageClient.BaseClient)
	if err != nil {
		return
	}
	clients.fileStorageClient = &fileStorageClient

	functionsInvokeClient, err := oci_functions.NewFunctionsInvokeClientWithConfigurationProvider(configProvider, "DUMMY_ENDPOINT")
	if err != nil {
		return
	}
	err = configureClient(&functionsInvokeClient.BaseClient)
	if err != nil {
		return
	}
	clients.functionsInvokeClient = &functionsInvokeClient

	functionsManagementClient, err := oci_functions.NewFunctionsManagementClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&functionsManagementClient.BaseClient)
	if err != nil {
		return
	}
	clients.functionsManagementClient = &functionsManagementClient

	healthChecksClient, err := oci_health_checks.NewHealthChecksClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&healthChecksClient.BaseClient)
	if err != nil {
		return
	}
	clients.healthChecksClient = &healthChecksClient

	identityClient, err := oci_identity.NewIdentityClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&identityClient.BaseClient)
	if err != nil {
		return
	}
	clients.identityClient = &identityClient

	kmsCryptoClient, err := oci_kms.NewKmsCryptoClientWithConfigurationProvider(configProvider, "DUMMY_ENDPOINT")
	if err != nil {
		return
	}
	err = configureClient(&kmsCryptoClient.BaseClient)
	if err != nil {
		return
	}
	clients.kmsCryptoClient = &kmsCryptoClient

	kmsManagementClient, err := oci_kms.NewKmsManagementClientWithConfigurationProvider(configProvider, "DUMMY_ENDPOINT")
	if err != nil {
		return
	}
	err = configureClient(&kmsManagementClient.BaseClient)
	if err != nil {
		return
	}
	clients.kmsManagementClient = &kmsManagementClient

	kmsVaultClient, err := oci_kms.NewKmsVaultClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&kmsVaultClient.BaseClient)
	if err != nil {
		return
	}
	clients.kmsVaultClient = &kmsVaultClient

	limitsClient, err := oci_limits.NewLimitsClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&limitsClient.BaseClient)
	if err != nil {
		return
	}
	clients.limitsClient = &limitsClient

	loadBalancerClient, err := oci_load_balancer.NewLoadBalancerClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&loadBalancerClient.BaseClient)
	if err != nil {
		return
	}
	clients.loadBalancerClient = &loadBalancerClient

	monitoringClient, err := oci_monitoring.NewMonitoringClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&monitoringClient.BaseClient)
	if err != nil {
		return
	}
	clients.monitoringClient = &monitoringClient

	notificationControlPlaneClient, err := oci_ons.NewNotificationControlPlaneClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&notificationControlPlaneClient.BaseClient)
	if err != nil {
		return
	}
	clients.notificationControlPlaneClient = &notificationControlPlaneClient

	notificationDataPlaneClient, err := oci_ons.NewNotificationDataPlaneClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&notificationDataPlaneClient.BaseClient)
	if err != nil {
		return
	}
	clients.notificationDataPlaneClient = &notificationDataPlaneClient

	objectStorageClient, err := oci_object_storage.NewObjectStorageClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&objectStorageClient.BaseClient)
	if err != nil {
		return
	}
	clients.objectStorageClient = &objectStorageClient

	quotasClient, err := oci_limits.NewQuotasClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&quotasClient.BaseClient)
	if err != nil {
		return
	}
	clients.quotasClient = &quotasClient

	streamAdminClient, err := oci_streaming.NewStreamAdminClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&streamAdminClient.BaseClient)
	if err != nil {
		return
	}
	clients.streamAdminClient = &streamAdminClient

	virtualNetworkClient, err := oci_core.NewVirtualNetworkClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&virtualNetworkClient.BaseClient)
	if err != nil {
		return
	}
	clients.virtualNetworkClient = &virtualNetworkClient

	waasClient, err := oci_waas.NewWaasClientWithConfigurationProvider(configProvider)
	if err != nil {
		return
	}
	err = configureClient(&waasClient.BaseClient)
	if err != nil {
		return
	}
	clients.waasClient = &waasClient

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
