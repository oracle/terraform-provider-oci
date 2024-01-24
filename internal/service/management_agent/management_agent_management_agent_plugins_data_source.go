// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package management_agent

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_management_agent "github.com/oracle/oci-go-sdk/v65/managementagent"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ManagementAgentManagementAgentPluginsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readManagementAgentManagementAgentPlugins,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"agent_id": {
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
			"platform_type": {
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
			"management_agent_plugins": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_console_deployable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"supported_platform_types": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"version": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readManagementAgentManagementAgentPlugins(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentPluginsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.ReadResource(sync)
}

type ManagementAgentManagementAgentPluginsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_management_agent.ManagementAgentClient
	Res    *oci_management_agent.ListManagementAgentPluginsResponse
}

func (s *ManagementAgentManagementAgentPluginsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagementAgentManagementAgentPluginsDataSourceCrud) Get() error {
	request := oci_management_agent.ListManagementAgentPluginsRequest{}

	if agentId, ok := s.D.GetOkExists("agent_id"); ok {
		tmp := agentId.(string)
		request.AgentId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
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

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_management_agent.ListManagementAgentPluginsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "management_agent")

	response, err := s.Client.ListManagementAgentPlugins(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagementAgentPlugins(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ManagementAgentManagementAgentPluginsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ManagementAgentManagementAgentPluginsDataSource-", ManagementAgentManagementAgentPluginsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		managementAgentPlugin := map[string]interface{}{}

		if r.Description != nil {
			managementAgentPlugin["description"] = *r.Description
		}

		if r.DisplayName != nil {
			managementAgentPlugin["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			managementAgentPlugin["id"] = *r.Id
		}

		if r.IsConsoleDeployable != nil {
			managementAgentPlugin["is_console_deployable"] = *r.IsConsoleDeployable
		}

		if r.Name != nil {
			managementAgentPlugin["name"] = *r.Name
		}

		managementAgentPlugin["state"] = r.LifecycleState

		managementAgentPlugin["supported_platform_types"] = r.SupportedPlatformTypes

		if r.Version != nil {
			managementAgentPlugin["version"] = *r.Version
		}

		resources = append(resources, managementAgentPlugin)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, ManagementAgentManagementAgentPluginsDataSource().Schema["management_agent_plugins"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("management_agent_plugins", resources); err != nil {
		return err
	}

	return nil
}
