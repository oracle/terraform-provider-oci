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

func JmsJmsPluginDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["jms_plugin_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(JmsJmsPluginResource(), fieldMap, readSingularJmsJmsPlugin)
}

func readSingularJmsJmsPlugin(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJmsPluginDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsJmsPluginDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.GetJmsPluginResponse
}

func (s *JmsJmsPluginDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsJmsPluginDataSourceCrud) Get() error {
	request := oci_jms.GetJmsPluginRequest{}

	if jmsPluginId, ok := s.D.GetOkExists("jms_plugin_id"); ok {
		tmp := jmsPluginId.(string)
		request.JmsPluginId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.GetJmsPlugin(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsJmsPluginDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AgentId != nil {
		s.D.Set("agent_id", *s.Res.AgentId)
	}

	s.D.Set("agent_type", s.Res.AgentType)

	s.D.Set("availability_status", s.Res.AvailabilityStatus)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.FleetId != nil {
		s.D.Set("fleet_id", *s.Res.FleetId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	if s.Res.OsArchitecture != nil {
		s.D.Set("os_architecture", *s.Res.OsArchitecture)
	}

	if s.Res.OsDistribution != nil {
		s.D.Set("os_distribution", *s.Res.OsDistribution)
	}

	s.D.Set("os_family", s.Res.OsFamily)

	if s.Res.PluginVersion != nil {
		s.D.Set("plugin_version", *s.Res.PluginVersion)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeLastSeen != nil {
		s.D.Set("time_last_seen", s.Res.TimeLastSeen.String())
	}

	if s.Res.TimeRegistered != nil {
		s.D.Set("time_registered", s.Res.TimeRegistered.String())
	}

	return nil
}
