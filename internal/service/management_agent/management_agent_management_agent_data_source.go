// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package management_agent

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_management_agent "github.com/oracle/oci-go-sdk/v58/managementagent"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ManagementAgentManagementAgentDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["management_agent_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ManagementAgentManagementAgentResource(), fieldMap, readSingularManagementAgentManagementAgent)
}

func readSingularManagementAgentManagementAgent(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.ReadResource(sync)
}

type ManagementAgentManagementAgentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_management_agent.ManagementAgentClient
	Res    *oci_management_agent.GetManagementAgentResponse
}

func (s *ManagementAgentManagementAgentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagementAgentManagementAgentDataSourceCrud) Get() error {
	request := oci_management_agent.GetManagementAgentRequest{}

	if managementAgentId, ok := s.D.GetOkExists("management_agent_id"); ok {
		tmp := managementAgentId.(string)
		request.ManagementAgentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "management_agent")

	response, err := s.Client.GetManagementAgent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ManagementAgentManagementAgentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("availability_status", s.Res.AvailabilityStatus)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Host != nil {
		s.D.Set("host", *s.Res.Host)
	}

	if s.Res.HostId != nil {
		s.D.Set("host_id", *s.Res.HostId)
	}

	if s.Res.InstallKeyId != nil {
		s.D.Set("install_key_id", *s.Res.InstallKeyId)
	}

	if s.Res.InstallPath != nil {
		s.D.Set("install_path", *s.Res.InstallPath)
	}

	s.D.Set("install_type", s.Res.InstallType)

	if s.Res.IsAgentAutoUpgradable != nil {
		s.D.Set("is_agent_auto_upgradable", *s.Res.IsAgentAutoUpgradable)
	}

	if s.Res.IsCustomerDeployed != nil {
		s.D.Set("is_customer_deployed", *s.Res.IsCustomerDeployed)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.PlatformName != nil {
		s.D.Set("platform_name", *s.Res.PlatformName)
	}

	s.D.Set("platform_type", s.Res.PlatformType)

	if s.Res.PlatformVersion != nil {
		s.D.Set("platform_version", *s.Res.PlatformVersion)
	}

	pluginList := []interface{}{}
	for _, item := range s.Res.PluginList {
		pluginList = append(pluginList, ManagementAgentPluginDetailsToMap(item))
	}
	s.D.Set("plugin_list", pluginList)

	if s.Res.ResourceArtifactVersion != nil {
		s.D.Set("resource_artifact_version", *s.Res.ResourceArtifactVersion)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastHeartbeat != nil {
		s.D.Set("time_last_heartbeat", s.Res.TimeLastHeartbeat.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}
