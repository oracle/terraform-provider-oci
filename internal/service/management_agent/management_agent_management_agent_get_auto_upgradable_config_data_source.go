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

func ManagementAgentManagementAgentGetAutoUpgradableConfigDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularManagementAgentManagementAgentGetAutoUpgradableConfig,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"is_agent_auto_upgradable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func readSingularManagementAgentManagementAgentGetAutoUpgradableConfig(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentGetAutoUpgradableConfigDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.ReadResource(sync)
}

type ManagementAgentManagementAgentGetAutoUpgradableConfigDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_management_agent.ManagementAgentClient
	Res    *oci_management_agent.GetAutoUpgradableConfigResponse
}

func (s *ManagementAgentManagementAgentGetAutoUpgradableConfigDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagementAgentManagementAgentGetAutoUpgradableConfigDataSourceCrud) Get() error {
	request := oci_management_agent.GetAutoUpgradableConfigRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "management_agent")

	response, err := s.Client.GetAutoUpgradableConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ManagementAgentManagementAgentGetAutoUpgradableConfigDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ManagementAgentManagementAgentGetAutoUpgradableConfigDataSource-", ManagementAgentManagementAgentGetAutoUpgradableConfigDataSource(), s.D))

	if s.Res.IsAgentAutoUpgradable != nil {
		s.D.Set("is_agent_auto_upgradable", *s.Res.IsAgentAutoUpgradable)
	}

	return nil
}
