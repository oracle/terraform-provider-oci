// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_functions "github.com/oracle/oci-go-sdk/functions"
)

func FunctionsApplicationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFunctionsApplications,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"applications": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(FunctionsApplicationResource()),
			},
		},
	}
}

func readFunctionsApplications(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsApplicationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).functionsManagementClient

	return ReadResource(sync)
}

type FunctionsApplicationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_functions.FunctionsManagementClient
	Res    *oci_functions.ListApplicationsResponse
}

func (s *FunctionsApplicationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FunctionsApplicationsDataSourceCrud) Get() error {
	request := oci_functions.ListApplicationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_functions.ApplicationLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "functions")

	response, err := s.Client.ListApplications(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListApplications(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FunctionsApplicationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		application := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			application["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			application["display_name"] = *r.DisplayName
		}

		application["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			application["id"] = *r.Id
		}

		application["state"] = r.LifecycleState

		application["subnet_ids"] = r.SubnetIds

		if r.TimeCreated != nil {
			application["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			application["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, application)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, FunctionsApplicationsDataSource().Schema["applications"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("applications", resources); err != nil {
		return err
	}

	return nil
}
