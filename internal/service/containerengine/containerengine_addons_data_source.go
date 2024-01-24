// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ContainerengineAddonsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readContainerengineAddons,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"addons": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ContainerengineAddonResource(),
			},
		},
	}
}

func readContainerengineAddons(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineAddonsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

type ContainerengineAddonsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.ListAddonsResponse
}

func (s *ContainerengineAddonsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineAddonsDataSourceCrud) Get() error {
	request := oci_containerengine.ListAddonsRequest{}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerengine")

	response, err := s.Client.ListAddons(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAddons(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ContainerengineAddonsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ContainerengineAddonsDataSource-", ContainerengineAddonsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		addon := map[string]interface{}{}

		if r.AddonError != nil {
			addon["addon_error"] = []interface{}{AddonErrorToMap(r.AddonError)}
		} else {
			addon["addon_error"] = nil
		}

		if r.CurrentInstalledVersion != nil {
			addon["current_installed_version"] = *r.CurrentInstalledVersion
		}

		if r.Name != nil {
			addon["addon_name"] = *r.Name
		}

		addon["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			addon["time_created"] = r.TimeCreated.String()
		}

		if r.Version != nil {
			addon["version"] = *r.Version
		}

		resources = append(resources, addon)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, ContainerengineAddonsDataSource().Schema["addons"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("addons", resources); err != nil {
		return err
	}

	return nil
}
