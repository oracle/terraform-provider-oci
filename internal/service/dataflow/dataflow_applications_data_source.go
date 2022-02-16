// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataflow

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_dataflow "github.com/oracle/oci-go-sdk/v58/dataflow"
)

func DataflowApplicationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataflowApplications,
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
			"display_name_starts_with": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"owner_principal_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"spark_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"applications": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DataflowApplicationResource()),
			},
		},
	}
}

func readDataflowApplications(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowApplicationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()

	return tfresource.ReadResource(sync)
}

type DataflowApplicationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dataflow.DataFlowClient
	Res    *oci_dataflow.ListApplicationsResponse
}

func (s *DataflowApplicationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataflowApplicationsDataSourceCrud) Get() error {
	request := oci_dataflow.ListApplicationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if displayNameStartsWith, ok := s.D.GetOkExists("display_name_starts_with"); ok {
		tmp := displayNameStartsWith.(string)
		request.DisplayNameStartsWith = &tmp
	}

	if ownerPrincipalId, ok := s.D.GetOkExists("owner_principal_id"); ok {
		tmp := ownerPrincipalId.(string)
		request.OwnerPrincipalId = &tmp
	}

	if sparkVersion, ok := s.D.GetOkExists("spark_version"); ok {
		tmp := sparkVersion.(string)
		request.SparkVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dataflow")

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

func (s *DataflowApplicationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataflowApplicationsDataSource-", DataflowApplicationsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		application := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			application["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			application["display_name"] = *r.DisplayName
		}

		application["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			application["id"] = *r.Id
		}

		application["language"] = r.Language

		if r.OwnerPrincipalId != nil {
			application["owner_principal_id"] = *r.OwnerPrincipalId
		}

		if r.OwnerUserName != nil {
			application["owner_user_name"] = *r.OwnerUserName
		}

		if r.SparkVersion != nil {
			application["spark_version"] = *r.SparkVersion
		}

		application["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			application["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			application["time_updated"] = r.TimeUpdated.String()
		}

		application["type"] = r.Type

		resources = append(resources, application)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DataflowApplicationsDataSource().Schema["applications"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("applications", resources); err != nil {
		return err
	}

	return nil
}
