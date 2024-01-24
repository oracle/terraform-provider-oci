// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package computeinstanceagent

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_computeinstanceagent "github.com/oracle/oci-go-sdk/v65/computeinstanceagent"
)

func ComputeinstanceagentInstanceAgentPluginsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readComputeinstanceagentInstanceAgentPlugins,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"instanceagent_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_agent_plugins": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"message": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_last_updated_utc": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readComputeinstanceagentInstanceAgentPlugins(d *schema.ResourceData, m interface{}) error {
	sync := &ComputeinstanceagentInstanceAgentPluginsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PluginClient()

	return tfresource.ReadResource(sync)
}

type ComputeinstanceagentInstanceAgentPluginsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_computeinstanceagent.PluginClient
	Res    *oci_computeinstanceagent.ListInstanceAgentPluginsResponse
}

func (s *ComputeinstanceagentInstanceAgentPluginsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ComputeinstanceagentInstanceAgentPluginsDataSourceCrud) Get() error {
	request := oci_computeinstanceagent.ListInstanceAgentPluginsRequest{}

	if instanceagentId, ok := s.D.GetOkExists("instanceagent_id"); ok {
		tmp := instanceagentId.(string)
		request.InstanceagentId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_computeinstanceagent.ListInstanceAgentPluginsStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "computeinstanceagent")

	response, err := s.Client.ListInstanceAgentPlugins(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ComputeinstanceagentInstanceAgentPluginsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ComputeinstanceagentInstanceAgentPluginsDataSource-", ComputeinstanceagentInstanceAgentPluginsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		instanceAgentPlugin := map[string]interface{}{}

		if r.Name != nil {
			instanceAgentPlugin["name"] = *r.Name
		}

		instanceAgentPlugin["status"] = r.Status

		if r.TimeLastUpdatedUtc != nil {
			instanceAgentPlugin["time_last_updated_utc"] = r.TimeLastUpdatedUtc.String()
		}

		resources = append(resources, instanceAgentPlugin)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, ComputeinstanceagentInstanceAgentPluginsDataSource().Schema["instance_agent_plugins"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("instance_agent_plugins", resources); err != nil {
		return err
	}

	return nil
}
