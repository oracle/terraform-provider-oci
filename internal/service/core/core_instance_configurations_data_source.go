// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreInstanceConfigurationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreInstanceConfigurations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_configurations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreInstanceConfigurationResource()),
			},
		},
	}
}

func readCoreInstanceConfigurations(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceConfigurationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeManagementClient()

	return tfresource.ReadResource(sync)
}

type CoreInstanceConfigurationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeManagementClient
	Res    *oci_core.ListInstanceConfigurationsResponse
}

func (s *CoreInstanceConfigurationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreInstanceConfigurationsDataSourceCrud) Get() error {
	request := oci_core.ListInstanceConfigurationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListInstanceConfigurations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInstanceConfigurations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreInstanceConfigurationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreInstanceConfigurationsDataSource-", CoreInstanceConfigurationsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		instanceConfiguration := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			instanceConfiguration["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			instanceConfiguration["display_name"] = *r.DisplayName
		}

		instanceConfiguration["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			instanceConfiguration["id"] = *r.Id
		}

		if r.TimeCreated != nil {
			instanceConfiguration["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, instanceConfiguration)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreInstanceConfigurationsDataSource().Schema["instance_configurations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("instance_configurations", resources); err != nil {
		return err
	}

	return nil
}
