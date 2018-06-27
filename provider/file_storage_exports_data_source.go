// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_file_storage "github.com/oracle/oci-go-sdk/filestorage"

	"github.com/oracle/terraform-provider-oci/crud"
)

func ExportsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readExports,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"export_set_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"file_system_id": {
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
			"exports": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ExportResource(),
			},
		},
	}
}

func readExports(d *schema.ResourceData, m interface{}) error {
	sync := &ExportsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return crud.ReadResource(sync)
}

type ExportsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_file_storage.FileStorageClient
	Res    *oci_file_storage.ListExportsResponse
}

func (s *ExportsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ExportsDataSourceCrud) Get() error {
	request := oci_file_storage.ListExportsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if exportSetId, ok := s.D.GetOkExists("export_set_id"); ok {
		tmp := exportSetId.(string)
		request.ExportSetId = &tmp
	}

	if fileSystemId, ok := s.D.GetOkExists("file_system_id"); ok {
		tmp := fileSystemId.(string)
		request.FileSystemId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_file_storage.ListExportsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "file_storage")

	response, err := s.Client.ListExports(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExports(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ExportsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		export := map[string]interface{}{}

		if r.ExportSetId != nil {
			export["export_set_id"] = *r.ExportSetId
		}

		if r.FileSystemId != nil {
			export["file_system_id"] = *r.FileSystemId
		}

		if r.Id != nil {
			export["id"] = *r.Id
		}

		if r.Path != nil {
			export["path"] = *r.Path
		}

		export["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			export["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, export)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, ExportsDataSource().Schema["exports"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("exports", resources); err != nil {
		panic(err)
	}

	return
}
