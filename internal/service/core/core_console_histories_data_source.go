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

func CoreConsoleHistoriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreConsoleHistories,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"console_histories": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreConsoleHistoryResource()),
			},
		},
	}
}

func readCoreConsoleHistories(d *schema.ResourceData, m interface{}) error {
	sync := &CoreConsoleHistoriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreConsoleHistoriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListConsoleHistoriesResponse
}

func (s *CoreConsoleHistoriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreConsoleHistoriesDataSourceCrud) Get() error {
	request := oci_core.ListConsoleHistoriesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.ConsoleHistoryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListConsoleHistories(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListConsoleHistories(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreConsoleHistoriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreConsoleHistoriesDataSource-", CoreConsoleHistoriesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		consoleHistory := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailabilityDomain != nil {
			consoleHistory["availability_domain"] = *r.AvailabilityDomain
		}

		if r.DefinedTags != nil {
			consoleHistory["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			consoleHistory["display_name"] = *r.DisplayName
		}

		consoleHistory["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			consoleHistory["id"] = *r.Id
		}

		if r.InstanceId != nil {
			consoleHistory["instance_id"] = *r.InstanceId
		}

		consoleHistory["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			consoleHistory["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, consoleHistory)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreConsoleHistoriesDataSource().Schema["console_histories"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("console_histories", resources); err != nil {
		return err
	}

	return nil
}
