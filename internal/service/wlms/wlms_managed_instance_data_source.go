// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package wlms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_wlms "github.com/oracle/oci-go-sdk/v65/wlms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func WlmsManagedInstanceDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularWlmsManagedInstance,
		Schema: map[string]*schema.Schema{
			"managed_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"configuration": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"discovery_interval": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"domain_search_paths": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"os_arch": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"os_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"plugin_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularWlmsManagedInstance(d *schema.ResourceData, m interface{}) error {
	sync := &WlmsManagedInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WeblogicManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type WlmsManagedInstanceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_wlms.WeblogicManagementServiceClient
	Res    *oci_wlms.GetManagedInstanceResponse
}

func (s *WlmsManagedInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WlmsManagedInstanceDataSourceCrud) Get() error {
	request := oci_wlms.GetManagedInstanceRequest{}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "wlms")

	response, err := s.Client.GetManagedInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *WlmsManagedInstanceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Configuration != nil {
		s.D.Set("configuration", []interface{}{ManagedInstanceConfigurationToMap(s.Res.Configuration)})
	} else {
		s.D.Set("configuration", nil)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.HostName != nil {
		s.D.Set("host_name", *s.Res.HostName)
	}

	if s.Res.OsArch != nil {
		s.D.Set("os_arch", *s.Res.OsArch)
	}

	if s.Res.OsName != nil {
		s.D.Set("os_name", *s.Res.OsName)
	}

	if s.Res.PluginStatus != nil {
		s.D.Set("plugin_status", *s.Res.PluginStatus)
	}

	if s.Res.ServerCount != nil {
		s.D.Set("server_count", *s.Res.ServerCount)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
