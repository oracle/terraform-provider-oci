// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_file_storage "github.com/oracle/oci-go-sdk/v56/filestorage"
)

func FileStorageExportsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFileStorageExports,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
				Elem:     tfresource.GetDataSourceItemSchema(FileStorageExportResource()),
			},
		},
	}
}

func readFileStorageExports(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageExportsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.ReadResource(sync)
}

type FileStorageExportsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_file_storage.FileStorageClient
	Res    *oci_file_storage.ListExportsResponse
}

func (s *FileStorageExportsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FileStorageExportsDataSourceCrud) Get() error {
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "file_storage")

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

func (s *FileStorageExportsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FileStorageExportsDataSource-", FileStorageExportsDataSource(), s.D))
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
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, FileStorageExportsDataSource().Schema["exports"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("exports", resources); err != nil {
		return err
	}

	return nil
}
