// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_announcements_service "github.com/oracle/oci-go-sdk/v65/announcementsservice"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_announcements_service.AnnouncementSubscriptionClient", &OracleClient{InitClientFn: initAnnouncementsserviceAnnouncementSubscriptionClient})
}

func initAnnouncementsserviceAnnouncementSubscriptionClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_announcements_service.NewAnnouncementSubscriptionClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) AnnouncementSubscriptionClient() *oci_announcements_service.AnnouncementSubscriptionClient {
	return m.GetClient("oci_announcements_service.AnnouncementSubscriptionClient").(*oci_announcements_service.AnnouncementSubscriptionClient)
}
