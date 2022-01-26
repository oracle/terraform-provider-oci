// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oda

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_oda "github.com/oracle/oci-go-sdk/v56/oda"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func OdaOdaInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOdaOdaInstances,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oda_instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(OdaOdaInstanceResource()),
			},
		},
	}
}

func readOdaOdaInstances(d *schema.ResourceData, m interface{}) error {
	sync := &OdaOdaInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OdaClient()

	return tfresource.ReadResource(sync)
}

type OdaOdaInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_oda.OdaClient
	Res    *oci_oda.ListOdaInstancesResponse
}

func (s *OdaOdaInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OdaOdaInstancesDataSourceCrud) Get() error {
	request := oci_oda.ListOdaInstancesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_oda.ListOdaInstancesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "oda")

	response, err := s.Client.ListOdaInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOdaInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OdaOdaInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OdaOdaInstancesDataSource-", OdaOdaInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		odaInstance := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			odaInstance["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			odaInstance["description"] = *r.Description
		}

		if r.DisplayName != nil {
			odaInstance["display_name"] = *r.DisplayName
		}

		odaInstance["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			odaInstance["id"] = *r.Id
		}

		odaInstance["lifecycle_sub_state"] = r.LifecycleSubState

		odaInstance["shape_name"] = r.ShapeName

		odaInstance["state"] = r.LifecycleState

		if r.StateMessage != nil {
			odaInstance["state_message"] = *r.StateMessage
		}

		if r.TimeCreated != nil {
			odaInstance["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			odaInstance["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, odaInstance)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OdaOdaInstancesDataSource().Schema["oda_instances"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("oda_instances", resources); err != nil {
		return err
	}

	return nil
}
