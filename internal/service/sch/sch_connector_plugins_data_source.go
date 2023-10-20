// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package sch

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_sch "github.com/oracle/oci-go-sdk/v65/sch"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func SchConnectorPluginsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSchConnectorPlugins,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"connector_plugin_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"estimated_throughput": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"kind": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"max_retention": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"schema": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readSchConnectorPlugins(d *schema.ResourceData, m interface{}) error {
	sync := &SchConnectorPluginsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ConnectorPluginsClient()

	return tfresource.ReadResource(sync)
}

type SchConnectorPluginsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_sch.ConnectorPluginsClient
	Res    *oci_sch.ListConnectorPluginsResponse
}

func (s *SchConnectorPluginsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *SchConnectorPluginsDataSourceCrud) Get() error {
	request := oci_sch.ListConnectorPluginsRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_sch.ListConnectorPluginsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "sch")

	response, err := s.Client.ListConnectorPlugins(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListConnectorPlugins(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *SchConnectorPluginsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("SchConnectorPluginsDataSource-", SchConnectorPluginsDataSource(), s.D))
	resources := []map[string]interface{}{}
	connectorPlugin := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ConnectorPluginSummaryToMap(item))
	}
	connectorPlugin["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, SchConnectorPluginsDataSource().Schema["connector_plugin_collection"].Elem.(*schema.Resource).Schema)
		connectorPlugin["items"] = items
	}

	resources = append(resources, connectorPlugin)
	if err := s.D.Set("connector_plugin_collection", resources); err != nil {
		return err
	}

	return nil
}

func ConnectorPluginSummaryToMap(obj oci_sch.ConnectorPluginSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_sch.SourceConnectorPluginSummary:
		result["kind"] = "SOURCE"

		if v.MaxRetention != nil {
			result["max_retention"] = string(*v.MaxRetention)
		}
	case oci_sch.TargetConnectorPluginSummary:
		result["kind"] = "TARGET"
	default:
		log.Printf("[WARN] Received 'kind' of unknown type %v", obj)
		return nil
	}

	return result
}
