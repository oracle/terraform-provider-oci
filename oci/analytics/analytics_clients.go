// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_analytics "github.com/oracle/oci-go-sdk/v49/analytics"

	tf_common "github.com/terraform-providers/terraform-provider-oci/oci"

	oci_common "github.com/oracle/oci-go-sdk/v49/common"

	tf_common "github.com/terraform-providers/terraform-provider-oci/oci"
)

func init() {
	tf_common.RegisterOracleClient("oci_analytics.AnalyticsClient", &tf_common.OracleClient{InitClientFn: initAnalyticsAnalyticsClient})
}

type OracleAnalyticsClients struct {
	*tf_common.OracleClients
}

func initAnalyticsAnalyticsClient(configProvider oci_common.ConfigurationProvider, configureClient tf_common.ConfigureClient, serviceClientOverrides tf_common.ServiceClientOverrides) (interface{}, error) {
	client, err := oci_analytics.NewAnalyticsClientWithConfigurationProvider(configProvider)
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

//func (m *OracleAnalyticsClients) analyticsClient() *oci_analytics.AnalyticsClient {
//	return m.GetClient("oci_analytics.AnalyticsClient").(*oci_analytics.AnalyticsClient)
//}
