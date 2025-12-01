// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	"net/http"
	"strconv"
	"strings"

	oci_resource_analytics "github.com/oracle/oci-go-sdk/v65/resourceanalytics"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_resource_analytics.MonitoredRegionClient", &OracleClient{InitClientFn: initResourceanalyticsMonitoredRegionClient})
	RegisterOracleClient("oci_resource_analytics.ResourceAnalyticsInstanceClient", &OracleClient{InitClientFn: initResourceanalyticsResourceAnalyticsInstanceClient})
	RegisterOracleClient("oci_resource_analytics.TenancyAttachmentClient", &OracleClient{InitClientFn: initResourceanalyticsTenancyAttachmentClient})
}

func initResourceanalyticsMonitoredRegionClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_resource_analytics.NewMonitoredRegionClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) MonitoredRegionClient() *oci_resource_analytics.MonitoredRegionClient {
	return m.GetClient("oci_resource_analytics.MonitoredRegionClient").(*oci_resource_analytics.MonitoredRegionClient)
}

func initResourceanalyticsResourceAnalyticsInstanceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_resource_analytics.NewResourceAnalyticsInstanceClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	originalInterceptor := client.Interceptor
	//Add interceptor to ensure Content-Length header is set for POST requests without body
	//This fixes the issue where disableOac calls fail due to missing Content-Length header
	client.Interceptor = func(request *http.Request) error {
		if request.Method == "POST" && request.ContentLength <= 0 {
			// For POST requests without a body (like disableOac), set Content-Length to 0
			if request.Header.Get("Content-Length") == "" {
				request.Header.Set("Content-Length", "0")
				request.ContentLength = 0
			}
		}
		// Also handle cases where we have a body but Content-Length is missing
		if request.Method == "POST" && request.Body != nil {
			if request.Header.Get("Content-Length") == "" && request.ContentLength > 0 {
				request.Header.Set("Content-Length", strconv.FormatInt(request.ContentLength, 10))
			}
		}
		// Handle URL path specific cases for disableOac
		if request.Method == "POST" && strings.Contains(request.URL.Path, "disableOac") {
			if request.Header.Get("Content-Length") == "" {
				request.Header.Set("Content-Length", "0")
				request.ContentLength = 0
			}
		}

		err := originalInterceptor(request)
		if err != nil {
			return err
		}
		return nil
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) ResourceAnalyticsInstanceClient() *oci_resource_analytics.ResourceAnalyticsInstanceClient {
	return m.GetClient("oci_resource_analytics.ResourceAnalyticsInstanceClient").(*oci_resource_analytics.ResourceAnalyticsInstanceClient)
}

func initResourceanalyticsTenancyAttachmentClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_resource_analytics.NewTenancyAttachmentClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) TenancyAttachmentClient() *oci_resource_analytics.TenancyAttachmentClient {
	return m.GetClient("oci_resource_analytics.TenancyAttachmentClient").(*oci_resource_analytics.TenancyAttachmentClient)
}
