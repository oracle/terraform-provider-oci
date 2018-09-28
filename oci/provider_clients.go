// Copyright (c) 2018, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	oci_audit "github.com/oracle/oci-go-sdk/audit"
	oci_containerengine "github.com/oracle/oci-go-sdk/containerengine"
	oci_core "github.com/oracle/oci-go-sdk/core"
	oci_database "github.com/oracle/oci-go-sdk/database"
	oci_dns "github.com/oracle/oci-go-sdk/dns"
	oci_email "github.com/oracle/oci-go-sdk/email"
	oci_file_storage "github.com/oracle/oci-go-sdk/filestorage"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

func setGoSDKClients(clients *OracleClients, officialSdkConfigProvider oci_common.ConfigurationProvider, httpClient *http.Client, userAgent string) (err error) {
	// Official Go SDK clients:

	auditClient, err := oci_audit.NewAuditClientWithConfigurationProvider(officialSdkConfigProvider)
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
	identityClient, err := oci_identity.NewIdentityClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}
	loadBalancerClient, err := oci_load_balancer.NewLoadBalancerClientWithConfigurationProvider(officialSdkConfigProvider)
	if err != nil {
		return
	}
	objectStorageClient, err := oci_object_storage.NewObjectStorageClientWithConfigurationProvider(officialSdkConfigProvider)
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

	configureClient := func(client *oci_common.BaseClient) error {
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

		// R1 Support
		if region, err := officialSdkConfigProvider.Region(); err == nil && strings.ToLower(region) == "r1" {
			service := strings.Split(client.Host, ".")[0]
			domainName := getEnvSettingWithBlankDefault("oracle_r1_domain_name")
			if domainName == "" {
				return fmt.Errorf("oracle_r1_domain_name is required env setting for r1 region")
			}
			client.Host = fmt.Sprintf("%s.%s.%s", service, strings.ToLower(region), domainName)

			pool := x509.NewCertPool()
			//readCertPem reads the pem files to a []byte
			cert, err := readCertPem()
			if err != nil {
				return err
			}
			if ok := pool.AppendCertsFromPEM(cert); !ok {
				return fmt.Errorf("failed to append R1 cert to the cert pool")
			}
			//install the certificates to the client
			if h, ok := client.HTTPClient.(*http.Client); ok {
				tr := &http.Transport{TLSClientConfig: &tls.Config{RootCAs: pool}}
				h.Transport = tr
			} else {
				return fmt.Errorf("the client dispatcher is not of http.Client type. can not patch the tls config")
			}
		}
		return nil
	}

	err = configureClient(&auditClient.BaseClient)
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
	err = configureClient(&identityClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&loadBalancerClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&objectStorageClient.BaseClient)
	if err != nil {
		return
	}
	err = configureClient(&virtualNetworkClient.BaseClient)
	if err != nil {
		return
	}

	clients.auditClient = &auditClient
	clients.blockstorageClient = &blockstorageClient
	clients.computeClient = &computeClient
	clients.containerEngineClient = &containerEngineClient
	clients.databaseClient = &databaseClient
	clients.dnsClient = &dnsClient
	clients.emailClient = &emailClient
	clients.fileStorageClient = &fileStorageClient
	clients.identityClient = &identityClient
	clients.loadBalancerClient = &loadBalancerClient
	clients.objectStorageClient = &objectStorageClient
	clients.virtualNetworkClient = &virtualNetworkClient

	return
}

type OracleClients struct {
	auditClient           *oci_audit.AuditClient
	blockstorageClient    *oci_core.BlockstorageClient
	computeClient         *oci_core.ComputeClient
	containerEngineClient *oci_containerengine.ContainerEngineClient
	databaseClient        *oci_database.DatabaseClient
	dnsClient             *oci_dns.DnsClient
	emailClient           *oci_email.EmailClient
	fileStorageClient     *oci_file_storage.FileStorageClient
	identityClient        *oci_identity.IdentityClient
	loadBalancerClient    *oci_load_balancer.LoadBalancerClient
	objectStorageClient   *oci_object_storage.ObjectStorageClient
	virtualNetworkClient  *oci_core.VirtualNetworkClient
}
