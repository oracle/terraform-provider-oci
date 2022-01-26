// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package management_agent

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_management_agent "github.com/oracle/oci-go-sdk/v56/managementagent"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ManagementAgentManagementAgentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readManagementAgentManagementAgents,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"install_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_customer_deployed": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"platform_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"plugin_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"management_agents": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(ManagementAgentManagementAgentResource()),
			},
		},
	}
}

func readManagementAgentManagementAgents(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.ReadResource(sync)
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

	if availabilityStatus, ok := s.D.GetOkExists("availability_status"); ok {
		request.AvailabilityStatus = oci_management_agent.ListManagementAgentsAvailabilityStatusEnum(availabilityStatus.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if hostId, ok := s.D.GetOkExists("host_id"); ok {
		tmp := hostId.(string)
		request.HostId = &tmp
	}

	if installType, ok := s.D.GetOkExists("install_type"); ok {
		request.InstallType = oci_management_agent.ListManagementAgentsInstallTypeEnum(installType.(string))
	}

	if isCustomerDeployed, ok := s.D.GetOkExists("is_customer_deployed"); ok {
		tmp := isCustomerDeployed.(bool)
		request.IsCustomerDeployed = &tmp
	}

	if platformType, ok := s.D.GetOkExists("platform_type"); ok {
		interfaces := platformType.([]interface{})
		tmp := make([]oci_management_agent.PlatformTypesEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_management_agent.PlatformTypesEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("platform_type") {
			request.PlatformType = tmp
		}
	}

	if pluginName, ok := s.D.GetOkExists("plugin_name"); ok {
		interfaces := pluginName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("plugin_name") {
			request.PluginName = tmp
		}
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_management_agent.ListManagementAgentsLifecycleStateEnum(state.(string))
	}

	if version, ok := s.D.GetOkExists("version"); ok {
		interfaces := version.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("version") {
			request.Version = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "management_agent")

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

	s.D.SetId(tfresource.GenerateDataSourceHashID("ManagementAgentManagementAgentsDataSource-", ManagementAgentManagementAgentsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		managementAgent := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		managementAgent["availability_status"] = r.AvailabilityStatus

		if r.DefinedTags != nil {
			managementAgent["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			managementAgent["display_name"] = *r.DisplayName
		}

		managementAgent["freeform_tags"] = r.FreeformTags

		if r.Host != nil {
			managementAgent["host"] = *r.Host
		}

		if r.HostId != nil {
			managementAgent["host_id"] = *r.HostId
		}

		if r.Id != nil {
			managementAgent["id"] = *r.Id
		}

		if r.InstallKeyId != nil {
			managementAgent["install_key_id"] = *r.InstallKeyId
		}

		managementAgent["install_type"] = r.InstallType

		if r.IsAgentAutoUpgradable != nil {
			managementAgent["is_agent_auto_upgradable"] = *r.IsAgentAutoUpgradable
		}

		if r.IsCustomerDeployed != nil {
			managementAgent["is_customer_deployed"] = *r.IsCustomerDeployed
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

		if r.ResourceArtifactVersion != nil {
			managementAgent["resource_artifact_version"] = *r.ResourceArtifactVersion
		}

		managementAgent["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			managementAgent["time_created"] = r.TimeCreated.String()
		}

		if r.TimeLastHeartbeat != nil {
			managementAgent["time_last_heartbeat"] = r.TimeLastHeartbeat.String()
		}

		if r.TimeUpdated != nil {
			managementAgent["time_updated"] = r.TimeUpdated.String()
		}

		if r.Version != nil {
			managementAgent["version"] = *r.Version
		}

		resources = append(resources, managementAgent)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, ManagementAgentManagementAgentsDataSource().Schema["management_agents"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("management_agents", resources); err != nil {
		return err
	}

	return nil
}
