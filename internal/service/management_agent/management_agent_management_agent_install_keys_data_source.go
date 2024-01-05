// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package management_agent

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_management_agent "github.com/oracle/oci-go-sdk/v65/managementagent"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ManagementAgentManagementAgentInstallKeysDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readManagementAgentManagementAgentInstallKeys,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_agent_install_keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(ManagementAgentManagementAgentInstallKeyResource()),
			},
		},
	}
}

func readManagementAgentManagementAgentInstallKeys(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentInstallKeysDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.ReadResource(sync)
}

type ManagementAgentManagementAgentInstallKeysDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_management_agent.ManagementAgentClient
	Res    *oci_management_agent.ListManagementAgentInstallKeysResponse
}

func (s *ManagementAgentManagementAgentInstallKeysDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagementAgentManagementAgentInstallKeysDataSourceCrud) Get() error {
	request := oci_management_agent.ListManagementAgentInstallKeysRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		tmp := accessLevel.(string)
		request.AccessLevel = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_management_agent.ListManagementAgentInstallKeysLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "management_agent")

	response, err := s.Client.ListManagementAgentInstallKeys(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagementAgentInstallKeys(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ManagementAgentManagementAgentInstallKeysDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ManagementAgentManagementAgentInstallKeysDataSource-", ManagementAgentManagementAgentInstallKeysDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		managementAgentInstallKey := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AllowedKeyInstallCount != nil {
			managementAgentInstallKey["allowed_key_install_count"] = *r.AllowedKeyInstallCount
		}

		if r.CreatedByPrincipalId != nil {
			managementAgentInstallKey["created_by_principal_id"] = *r.CreatedByPrincipalId
		}

		if r.CurrentKeyInstallCount != nil {
			managementAgentInstallKey["current_key_install_count"] = *r.CurrentKeyInstallCount
		}

		if r.DisplayName != nil {
			managementAgentInstallKey["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			managementAgentInstallKey["id"] = *r.Id
		}

		if r.IsUnlimited != nil {
			managementAgentInstallKey["is_unlimited"] = *r.IsUnlimited
		}

		if r.LifecycleDetails != nil {
			managementAgentInstallKey["lifecycle_details"] = *r.LifecycleDetails
		}

		managementAgentInstallKey["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			managementAgentInstallKey["time_created"] = r.TimeCreated.String()
		}

		if r.TimeExpires != nil {
			managementAgentInstallKey["time_expires"] = r.TimeExpires.Format(time.RFC3339Nano)
		}

		resources = append(resources, managementAgentInstallKey)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, ManagementAgentManagementAgentInstallKeysDataSource().Schema["management_agent_install_keys"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("management_agent_install_keys", resources); err != nil {
		return err
	}

	return nil
}
