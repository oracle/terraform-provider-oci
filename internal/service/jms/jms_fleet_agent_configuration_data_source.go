// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsFleetAgentConfigurationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["fleet_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(JmsFleetAgentConfigurationResource(), fieldMap, readSingularJmsFleetAgentConfiguration)
}

func readSingularJmsFleetAgentConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetAgentConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsFleetAgentConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.GetFleetAgentConfigurationResponse
}

func (s *JmsFleetAgentConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsFleetAgentConfigurationDataSourceCrud) Get() error {
	request := oci_jms.GetFleetAgentConfigurationRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.GetFleetAgentConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsFleetAgentConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsFleetAgentConfigurationDataSource-", JmsFleetAgentConfigurationDataSource(), s.D))

	if s.Res.AgentPollingIntervalInMinutes != nil {
		s.D.Set("agent_polling_interval_in_minutes", *s.Res.AgentPollingIntervalInMinutes)
	}

	if s.Res.IsCapturingIpAddressAndFqdnEnabled != nil {
		s.D.Set("is_capturing_ip_address_and_fqdn_enabled", *s.Res.IsCapturingIpAddressAndFqdnEnabled)
	}

	if s.Res.IsCollectingManagedInstanceMetricsEnabled != nil {
		s.D.Set("is_collecting_managed_instance_metrics_enabled", *s.Res.IsCollectingManagedInstanceMetricsEnabled)
	}

	if s.Res.IsCollectingUsernamesEnabled != nil {
		s.D.Set("is_collecting_usernames_enabled", *s.Res.IsCollectingUsernamesEnabled)
	}

	if s.Res.IsLibrariesScanEnabled != nil {
		s.D.Set("is_libraries_scan_enabled", *s.Res.IsLibrariesScanEnabled)
	}

	if s.Res.JavaUsageTrackerProcessingFrequencyInMinutes != nil {
		s.D.Set("java_usage_tracker_processing_frequency_in_minutes", *s.Res.JavaUsageTrackerProcessingFrequencyInMinutes)
	}

	if s.Res.JreScanFrequencyInMinutes != nil {
		s.D.Set("jre_scan_frequency_in_minutes", *s.Res.JreScanFrequencyInMinutes)
	}

	if s.Res.LinuxConfiguration != nil {
		s.D.Set("linux_configuration", []interface{}{FleetAgentOsConfigurationToMap(s.Res.LinuxConfiguration)})
	} else {
		s.D.Set("linux_configuration", nil)
	}

	if s.Res.MacOsConfiguration != nil {
		s.D.Set("mac_os_configuration", []interface{}{FleetAgentOsConfigurationToMap(s.Res.MacOsConfiguration)})
	} else {
		s.D.Set("mac_os_configuration", nil)
	}

	if s.Res.TimeLastModified != nil {
		s.D.Set("time_last_modified", s.Res.TimeLastModified.String())
	}

	if s.Res.WindowsConfiguration != nil {
		s.D.Set("windows_configuration", []interface{}{FleetAgentOsConfigurationToMap(s.Res.WindowsConfiguration)})
	} else {
		s.D.Set("windows_configuration", nil)
	}

	if s.Res.WorkRequestValidityPeriodInDays != nil {
		s.D.Set("work_request_validity_period_in_days", *s.Res.WorkRequestValidityPeriodInDays)
	}

	return nil
}
