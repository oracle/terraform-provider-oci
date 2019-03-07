// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	oci_audit "github.com/oracle/oci-go-sdk/audit"
	oci_autoscaling "github.com/oracle/oci-go-sdk/autoscaling"
	oci_containerengine "github.com/oracle/oci-go-sdk/containerengine"
	oci_core "github.com/oracle/oci-go-sdk/core"
	oci_database "github.com/oracle/oci-go-sdk/database"
	oci_dns "github.com/oracle/oci-go-sdk/dns"
	oci_email "github.com/oracle/oci-go-sdk/email"
	oci_file_storage "github.com/oracle/oci-go-sdk/filestorage"
	oci_health_checks "github.com/oracle/oci-go-sdk/healthchecks"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
	oci_kms "github.com/oracle/oci-go-sdk/keymanagement"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
	oci_monitoring "github.com/oracle/oci-go-sdk/monitoring"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"
	oci_ons "github.com/oracle/oci-go-sdk/ons"
	oci_streaming "github.com/oracle/oci-go-sdk/streaming"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

type ConfigureClient func(client *oci_common.BaseClient) error

var configureClient ConfigureClient

func setGoSDKClients(clients *OracleClients, officialSdkConfigProvider oci_common.ConfigurationProvider, httpClient *http.Client, userAgent string) (err error) {
	// Official Go SDK clients:

	auditClient, err := oci_audit.NewAuditClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}
	autoScalingClient, err := oci_autoscaling.NewAutoScalingClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}
	blockstorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}
	computeClient, err := oci_core.NewComputeClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}
	computeManagementClient, err := oci_core.NewComputeManagementClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}
	containerEngineClient, err := oci_containerengine.NewContainerEngineClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}
	databaseClient, err := oci_database.NewDatabaseClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}
	dnsClient, err := oci_dns.NewDnsClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}
	emailClient, err := oci_email.NewEmailClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}
	fileStorageClient, err := oci_file_storage.NewFileStorageClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}
	healthChecksClient, err := oci_health_checks.NewHealthChecksClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}
	identityClient, err := oci_identity.NewIdentityClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}
	kmsCryptoClient, err := oci_kms.NewKmsCryptoClientWithConfigurationProvider(officialSdkConfigProvider, "DUMMY_ENDPOINT")
	if err != nil {
		return
	}
	kmsManagementClient, err := oci_kms.NewKmsManagementClientWithConfigurationProvider(officialSdkConfigProvider, "DUMMY_ENDPOINT")
	if err != nil {
		return
	}
	kmsVaultClient, err := oci_kms.NewKmsVaultClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}
	loadBalancerClient, err := oci_load_balancer.NewLoadBalancerClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}
	monitoringClient, err := oci_monitoring.NewMonitoringClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}
	notificationControlPlaneClient, err := oci_ons.NewNotificationControlPlaneClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}
	notificationDataPlaneClient, err := oci_ons.NewNotificationDataPlaneClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}
	objectStorageClient, err := oci_object_storage.NewObjectStorageClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}
	streamAdminClient, err := oci_streaming.NewStreamAdminClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}
	virtualNetworkClient, err := oci_core.NewVirtualNetworkClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}

	useOboToken, err := strconv.ParseBool(getEnvSettingWithDefault("use_obo_token", "false"))
	if err != nil {
		return
	}

	simulateDb, _ := strconv.ParseBool(getEnvSettingWithDefault("simulate_db", "false"))

	requestSigner := oci_common.DefaultRequestSigner(officialSdkConfigProvider)
	var oboTokenProvider OboTokenProvider
	oboTokenProvider = emptyOboTokenProvider{}
	if useOboToken {
		// Add Obo token to the default list and update the signer
		httpHeadersToSign := append(oci_common.DefaultGenericHeaders(), requestHeaderOpcOboToken)
		requestSigner = oci_common.RequestSigner(officialSdkConfigProvider, httpHeadersToSign, oci_common.DefaultBodyHeaders())
		oboTokenProvider = oboTokenProviderFromEnv{}
	}

	configureClient = func(client *oci_common.BaseClient) error {
		client.HTTPClient = httpClient
		client.UserAgent = userAgent
		client.Signer = requestSigner
		client.Interceptor = func(r *http.Request) error {
			if oboToken, err := oboTokenProvider.OboToken(); err == nil && oboToken != "" {
				r.Header.Set(requestHeaderOpcOboToken, oboToken)
			}

			if simulateDb {
				if r.Method == http.MethodPost && (strings.Contains(r.URL.Path, "/dbSystems") || strings.Contains(r.URL.Path, "/autonomousData")) {
					r.Header.Set(requestHeaderOpcHostSerial, "FAKEHOSTSERIAL")
				}
			}
			return nil
		}

		// R1, et al Support
		domainNameOverride := getEnvSettingWithBlankDefault(domainNameOverrideEnv)
		r1DomainName := getEnvSettingWithBlankDefault(oracleR1DomainNameEnv)
		customCertLoc := getEnvSettingWithBlankDefault(customCertLocationEnv)
		r1CertLoc := getEnvSettingWithBlankDefault(r1CertLocationEnv)

		if domainNameOverride != "" || r1DomainName != "" {
			if domainNameOverride != "" && r1DomainName != "" {
				return fmt.Errorf("conflicting environment variables (domain_name_override and oracle_r1_domain_name) resulting in domain name ambiguity:  %s and %s", domainNameOverride, r1DomainName)
			}

			region, _ := officialSdkConfigProvider.Region()
			service := strings.Split(client.Host, ".")[0]
			client.Host = fmt.Sprintf("%s.%s.%s", service, strings.ToLower(region), domainNameOverride+r1DomainName)
		}

		if customCertLoc != "" || r1CertLoc != "" {
			if customCertLoc != "" && r1CertLoc != "" {
				return fmt.Errorf("conflicting environment variables (custom_cert_location and R1_CERT_LOCATION) resulting in certificate locations ambiguity: %s and %s", customCertLoc, r1CertLoc)
			}

			cert, err := ioutil.ReadFile(customCertLoc + r1CertLoc)
			if err != nil {
				return err
			}
			pool := x509.NewCertPool()
			if ok := pool.AppendCertsFromPEM(cert); !ok {
				return fmt.Errorf("failed to append R1 cert to the cert pool")
			}
			// install the certificates in the client
			if h, ok := client.HTTPClient.(*http.Client); ok {
				tr := &http.Transport{TLSClientConfig: &tls.Config{RootCAs: pool}}
				h.Transport = tr
			} else {
				return fmt.Errorf("the client dispatcher is not of http.Client type. can not patch the tls config")
			}
		}

		if r1DomainName != "" && r1CertLoc == "" || r1DomainName == "" && r1CertLoc != "" {
			return fmt.Errorf("both certificate location and domain name must be specified to target r1")
		}

		return nil
	}

	err = configureClient(&auditClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&autoScalingClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&blockstorageClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&computeClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&computeManagementClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&containerEngineClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&databaseClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&dnsClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&emailClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&fileStorageClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&healthChecksClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&identityClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&kmsCryptoClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&kmsManagementClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&kmsVaultClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&loadBalancerClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&monitoringClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&notificationControlPlaneClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&notificationDataPlaneClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&objectStorageClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&streamAdminClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&virtualNetworkClient.BaseClient)
	if err != nil {
		return
	}

	clients.auditClient = &auditClient
	clients.autoScalingClient = &autoScalingClient
	clients.blockstorageClient = &blockstorageClient
	clients.computeClient = &computeClient
	clients.computeManagementClient = &computeManagementClient
	clients.containerEngineClient = &containerEngineClient
	clients.databaseClient = &databaseClient
	clients.dnsClient = &dnsClient
	clients.emailClient = &emailClient
	clients.fileStorageClient = &fileStorageClient
	clients.healthChecksClient = &healthChecksClient
	clients.identityClient = &identityClient
	clients.kmsCryptoClient = &kmsCryptoClient
	clients.kmsManagementClient = &kmsManagementClient
	clients.kmsVaultClient = &kmsVaultClient
	clients.loadBalancerClient = &loadBalancerClient
	clients.monitoringClient = &monitoringClient
	clients.notificationControlPlaneClient = &notificationControlPlaneClient
	clients.notificationDataPlaneClient = &notificationDataPlaneClient
	clients.objectStorageClient = &objectStorageClient
	clients.streamAdminClient = &streamAdminClient
	clients.virtualNetworkClient = &virtualNetworkClient

	return
}

type OracleClients struct {
	auditClient                    *oci_audit.AuditClient
	autoScalingClient              *oci_autoscaling.AutoScalingClient
	blockstorageClient             *oci_core.BlockstorageClient
	computeClient                  *oci_core.ComputeClient
	computeManagementClient        *oci_core.ComputeManagementClient
	containerEngineClient          *oci_containerengine.ContainerEngineClient
	databaseClient                 *oci_database.DatabaseClient
	dnsClient                      *oci_dns.DnsClient
	emailClient                    *oci_email.EmailClient
	fileStorageClient              *oci_file_storage.FileStorageClient
	healthChecksClient             *oci_health_checks.HealthChecksClient
	identityClient                 *oci_identity.IdentityClient
	kmsCryptoClient                *oci_kms.KmsCryptoClient
	kmsManagementClient            *oci_kms.KmsManagementClient
	kmsVaultClient                 *oci_kms.KmsVaultClient
	loadBalancerClient             *oci_load_balancer.LoadBalancerClient
	monitoringClient               *oci_monitoring.MonitoringClient
	notificationControlPlaneClient *oci_ons.NotificationControlPlaneClient
	notificationDataPlaneClient    *oci_ons.NotificationDataPlaneClient
	objectStorageClient            *oci_object_storage.ObjectStorageClient
	streamAdminClient              *oci_streaming.StreamAdminClient
	virtualNetworkClient           *oci_core.VirtualNetworkClient
	configuration                  map[string]string
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
