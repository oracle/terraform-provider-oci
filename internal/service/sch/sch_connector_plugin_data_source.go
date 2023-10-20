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

func SchConnectorPluginDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularSchConnectorPlugin,
		Schema: map[string]*schema.Schema{
			"connector_plugin_name": {
				Type:     schema.TypeString,
				Required: true,
			},
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
	}
}

func readSingularSchConnectorPlugin(d *schema.ResourceData, m interface{}) error {
	sync := &SchConnectorPluginDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ConnectorPluginsClient()

	return tfresource.ReadResource(sync)
}

type SchConnectorPluginDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_sch.ConnectorPluginsClient
	Res    *oci_sch.GetConnectorPluginResponse
}

func (s *SchConnectorPluginDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *SchConnectorPluginDataSourceCrud) Get() error {
	request := oci_sch.GetConnectorPluginRequest{}

	if connectorPluginName, ok := s.D.GetOkExists("connector_plugin_name"); ok {
		tmp := connectorPluginName.(string)
		request.ConnectorPluginName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "sch")

	response, err := s.Client.GetConnectorPlugin(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *SchConnectorPluginDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("SchConnectorPluginDataSource-", SchConnectorPluginDataSource(), s.D))
	switch v := (s.Res.ConnectorPlugin).(type) {
	case oci_sch.SourceConnectorPlugin:
		s.D.Set("kind", "SOURCE")

		if v.MaxRetention != nil {
			s.D.Set("max_retention", *v.MaxRetention)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("estimated_throughput", v.EstimatedThroughput)

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.Schema != nil {
			s.D.Set("schema", *v.Schema)
		}

		s.D.Set("state", v.LifecycleState)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}
	case oci_sch.TargetConnectorPlugin:
		s.D.Set("kind", "TARGET")

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("estimated_throughput", v.EstimatedThroughput)

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.Schema != nil {
			s.D.Set("schema", *v.Schema)
		}

		s.D.Set("state", v.LifecycleState)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}
	default:
		log.Printf("[WARN] Received 'kind' of unknown type %v", s.Res.ConnectorPlugin)
		return nil
	}

	return nil
}
