// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package computeinstanceagent

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_computeinstanceagent "github.com/oracle/oci-go-sdk/v58/computeinstanceagent"
)

func ComputeinstanceagentInstanceAgentPluginDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularComputeinstanceagentInstanceAgentPlugin,
		Schema: map[string]*schema.Schema{
			"instanceagent_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"plugin_name": {
				Type:     schema.TypeString,
				Required: true,
			},
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
	}
}

func readSingularComputeinstanceagentInstanceAgentPlugin(d *schema.ResourceData, m interface{}) error {
	sync := &ComputeinstanceagentInstanceAgentPluginDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PluginClient()

	return tfresource.ReadResource(sync)
}

type ComputeinstanceagentInstanceAgentPluginDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_computeinstanceagent.PluginClient
	Res    *oci_computeinstanceagent.GetInstanceAgentPluginResponse
}

func (s *ComputeinstanceagentInstanceAgentPluginDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ComputeinstanceagentInstanceAgentPluginDataSourceCrud) Get() error {
	request := oci_computeinstanceagent.GetInstanceAgentPluginRequest{}

	if instanceagentId, ok := s.D.GetOkExists("instanceagent_id"); ok {
		tmp := instanceagentId.(string)
		request.InstanceagentId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if pluginName, ok := s.D.GetOkExists("plugin_name"); ok {
		tmp := pluginName.(string)
		request.PluginName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "computeinstanceagent")

	response, err := s.Client.GetInstanceAgentPlugin(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ComputeinstanceagentInstanceAgentPluginDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ComputeinstanceagentInstanceAgentPluginDataSource-", ComputeinstanceagentInstanceAgentPluginDataSource(), s.D))

	if s.Res.Message != nil {
		s.D.Set("message", *s.Res.Message)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("status", s.Res.Status)

	if s.Res.TimeLastUpdatedUtc != nil {
		s.D.Set("time_last_updated_utc", s.Res.TimeLastUpdatedUtc.String())
	}

	return nil
}
