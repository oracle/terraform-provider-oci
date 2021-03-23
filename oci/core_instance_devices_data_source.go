// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v37/core"
)

func init() {
	RegisterDatasource("oci_core_instance_devices", CoreInstanceDevicesDataSource())
}

func CoreInstanceDevicesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreInstanceDevices,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_available": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"devices": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_available": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCoreInstanceDevices(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceDevicesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient()

	return ReadResource(sync)
}

type CoreInstanceDevicesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListInstanceDevicesResponse
}

func (s *CoreInstanceDevicesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreInstanceDevicesDataSourceCrud) Get() error {
	request := oci_core.ListInstanceDevicesRequest{}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	if isAvailable, ok := s.D.GetOkExists("is_available"); ok {
		tmp := isAvailable.(bool)
		request.IsAvailable = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListInstanceDevices(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInstanceDevices(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreInstanceDevicesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("CoreInstanceDevicesDataSource-", CoreInstanceDevicesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		instanceDevice := map[string]interface{}{}

		if r.IsAvailable != nil {
			instanceDevice["is_available"] = *r.IsAvailable
		}

		if r.Name != nil {
			instanceDevice["name"] = *r.Name
		}

		resources = append(resources, instanceDevice)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, CoreInstanceDevicesDataSource().Schema["devices"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("devices", resources); err != nil {
		return err
	}

	return nil
}
