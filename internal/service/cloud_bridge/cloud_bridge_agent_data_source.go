// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_bridge

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_bridge "github.com/oracle/oci-go-sdk/v65/cloudbridge"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudBridgeAgentDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["agent_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CloudBridgeAgentResource(), fieldMap, readSingularCloudBridgeAgent)
}

func readSingularCloudBridgeAgent(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAgentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OcbAgentSvcClient()

	return tfresource.ReadResource(sync)
}

type CloudBridgeAgentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_bridge.OcbAgentSvcClient
	Res    *oci_cloud_bridge.GetAgentResponse
}

func (s *CloudBridgeAgentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudBridgeAgentDataSourceCrud) Get() error {
	request := oci_cloud_bridge.GetAgentRequest{}

	if agentId, ok := s.D.GetOkExists("agent_id"); ok {
		tmp := agentId.(string)
		request.AgentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_bridge")

	response, err := s.Client.GetAgent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CloudBridgeAgentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Agent.Id)

	if s.Res.AgentPubKey != nil {
		s.D.Set("agent_pub_key", *s.Res.AgentPubKey)
	}

	s.D.Set("agent_type", s.Res.AgentType)

	if s.Res.AgentVersion != nil {
		s.D.Set("agent_version", *s.Res.AgentVersion)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EnvironmentId != nil {
		s.D.Set("environment_id", *s.Res.EnvironmentId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("heart_beat_status", s.Res.HeartBeatStatus)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.OsVersion != nil {
		s.D.Set("os_version", *s.Res.OsVersion)
	}

	pluginList := []interface{}{}
	for _, item := range s.Res.PluginList {
		pluginList = append(pluginList, PluginSummaryToMap(item))
	}
	s.D.Set("plugin_list", pluginList)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeExpireAgentKeyInMs != nil {
		s.D.Set("time_expire_agent_key_in_ms", s.Res.TimeExpireAgentKeyInMs.String())
	}

	if s.Res.TimeLastSyncReceived != nil {
		s.D.Set("time_last_sync_received", s.Res.TimeLastSyncReceived.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
