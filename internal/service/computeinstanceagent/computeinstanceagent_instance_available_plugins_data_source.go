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

func ComputeinstanceagentInstanceAvailablePluginsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readComputeinstanceagentInstanceAvailablePlugins,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"os_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"os_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"available_plugins": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_enabled_by_default": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_supported": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"summary": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readComputeinstanceagentInstanceAvailablePlugins(d *schema.ResourceData, m interface{}) error {
	sync := &ComputeinstanceagentInstanceAvailablePluginsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PluginconfigClient()

	return tfresource.ReadResource(sync)
}

type ComputeinstanceagentInstanceAvailablePluginsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_computeinstanceagent.PluginconfigClient
	Res    *oci_computeinstanceagent.ListInstanceagentAvailablePluginsResponse
}

func (s *ComputeinstanceagentInstanceAvailablePluginsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ComputeinstanceagentInstanceAvailablePluginsDataSourceCrud) Get() error {
	request := oci_computeinstanceagent.ListInstanceagentAvailablePluginsRequest{}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if osName, ok := s.D.GetOkExists("os_name"); ok {
		tmp := osName.(string)
		request.OsName = &tmp
	}

	if osVersion, ok := s.D.GetOkExists("os_version"); ok {
		tmp := osVersion.(string)
		request.OsVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "computeinstanceagent")

	response, err := s.Client.ListInstanceagentAvailablePlugins(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ComputeinstanceagentInstanceAvailablePluginsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ComputeinstanceagentInstanceAvailablePluginsDataSource-", ComputeinstanceagentInstanceAvailablePluginsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		instanceAvailablePlugin := map[string]interface{}{}

		if r.IsEnabledByDefault != nil {
			instanceAvailablePlugin["is_enabled_by_default"] = *r.IsEnabledByDefault
		}

		if r.IsSupported != nil {
			instanceAvailablePlugin["is_supported"] = *r.IsSupported
		}

		if r.Name != nil {
			instanceAvailablePlugin["name"] = *r.Name
		}

		if r.Summary != nil {
			instanceAvailablePlugin["summary"] = *r.Summary
		}

		resources = append(resources, instanceAvailablePlugin)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, ComputeinstanceagentInstanceAvailablePluginsDataSource().Schema["available_plugins"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("available_plugins", resources); err != nil {
		return err
	}

	return nil
}
