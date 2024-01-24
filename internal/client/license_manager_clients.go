// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_license_manager "github.com/oracle/oci-go-sdk/v65/licensemanager"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_license_manager.LicenseManagerClient", &OracleClient{InitClientFn: initLicensemanagerLicenseManagerClient})
}

func initLicensemanagerLicenseManagerClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_license_manager.NewLicenseManagerClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) LicenseManagerClient() *oci_license_manager.LicenseManagerClient {
	return m.GetClient("oci_license_manager.LicenseManagerClient").(*oci_license_manager.LicenseManagerClient)
}
