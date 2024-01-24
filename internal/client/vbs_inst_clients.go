// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_vbs_inst "github.com/oracle/oci-go-sdk/v65/vbsinst"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_vbs_inst.VbsInstanceClient", &OracleClient{InitClientFn: initVbsinstVbsInstanceClient})
}

func initVbsinstVbsInstanceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_vbs_inst.NewVbsInstanceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) VbsInstanceClient() *oci_vbs_inst.VbsInstanceClient {
	return m.GetClient("oci_vbs_inst.VbsInstanceClient").(*oci_vbs_inst.VbsInstanceClient)
}
