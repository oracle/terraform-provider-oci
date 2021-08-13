// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_management_agent "github.com/oracle/oci-go-sdk/v45/managementagent"
)

func init() {
	RegisterDatasource("oci_management_agent_management_agents", ManagementAgentManagementAgentsDataSource())
}

func ManagementAgentManagementAgentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readManagementAgentManagementAgents,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_agents": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(ManagementAgentManagementAgentResource()),
			},
		},
	}
}

func readManagementAgentManagementAgents(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).managementAgentClient()

	return ReadResource(sync)
}

type ManagementAgentManagementAgentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_management_agent.ManagementAgentClient
	Res    *oci_management_agent.ListManagementAgentsResponse
}

func (s *ManagementAgentManagementAgentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagementAgentManagementAgentsDataSourceCrud) Get() error {
	request := oci_management_agent.ListManagementAgentsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_management_agent.ListManagementAgentsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "management_agent")

	response, err := s.Client.ListManagementAgents(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagementAgents(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ManagementAgentManagementAgentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("ManagementAgentManagementAgentsDataSource-", ManagementAgentManagementAgentsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		managementAgent := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		managementAgent["availability_status"] = r.AvailabilityStatus

		if r.DefinedTags != nil {
			managementAgent["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			managementAgent["display_name"] = *r.DisplayName
		}

		managementAgent["freeform_tags"] = r.FreeformTags

		if r.Host != nil {
			managementAgent["host"] = *r.Host
		}

		if r.Id != nil {
			managementAgent["id"] = *r.Id
		}

		if r.InstallKeyId != nil {
			managementAgent["install_key_id"] = *r.InstallKeyId
		}

		if r.IsAgentAutoUpgradable != nil {
			managementAgent["is_agent_auto_upgradable"] = *r.IsAgentAutoUpgradable
		}

		if r.LifecycleDetails != nil {
			managementAgent["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.PlatformName != nil {
			managementAgent["platform_name"] = *r.PlatformName
		}

		managementAgent["platform_type"] = r.PlatformType

		if r.PlatformVersion != nil {
			managementAgent["platform_version"] = *r.PlatformVersion
		}

		pluginList := []interface{}{}
		for _, item := range r.PluginList {
			pluginList = append(pluginList, ManagementAgentPluginDetailsToMap(item))
		}
		managementAgent["plugin_list"] = pluginList

		managementAgent["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			managementAgent["time_created"] = r.TimeCreated.String()
		}

		if r.TimeLastHeartbeat != nil {
			managementAgent["time_last_heartbeat"] = r.TimeLastHeartbeat.String()
		}

		if r.Version != nil {
			managementAgent["version"] = *r.Version
		}

		resources = append(resources, managementAgent)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, ManagementAgentManagementAgentsDataSource().Schema["management_agents"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("management_agents", resources); err != nil {
		return err
	}

	return nil
}
