// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_file_storage "github.com/oracle/oci-go-sdk/filestorage"

	"github.com/oracle/terraform-provider-oci/crud"
)

func ExportSetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readExportSets,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
			},
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
			"export_sets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ExportSetResource(),
			},
		},
	}
}

func readExportSets(d *schema.ResourceData, m interface{}) error {
	sync := &ExportSetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return crud.ReadResource(sync)
}

type ExportSetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_file_storage.FileStorageClient
	Res    *oci_file_storage.ListExportSetsResponse
}

func (s *ExportSetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ExportSetsDataSourceCrud) Get() error {
	request := oci_file_storage.ListExportSetsRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

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
		request.LifecycleState = oci_file_storage.ListExportSetsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "file_storage")

	response, err := s.Client.ListExportSets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExportSets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ExportSetsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		exportSet := map[string]interface{}{
			"availability_domain": *r.AvailabilityDomain,
			"compartment_id":      *r.CompartmentId,
		}

		if r.DisplayName != nil {
			exportSet["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			exportSet["id"] = *r.Id
		}

		exportSet["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			exportSet["time_created"] = r.TimeCreated.String()
		}

		if r.VcnId != nil {
			exportSet["vcn_id"] = *r.VcnId
		}

		resources = append(resources, exportSet)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, ExportSetsDataSource().Schema["export_sets"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("export_sets", resources); err != nil {
		panic(err)
	}

	return
}
