// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_file_storage "github.com/oracle/oci-go-sdk/filestorage"

	"github.com/oracle/terraform-provider-oci/crud"
)

func FileSystemsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFileSystems,
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
			"file_systems": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     FileSystemResource(),
			},
		},
	}
}

func readFileSystems(d *schema.ResourceData, m interface{}) error {
	sync := &FileSystemsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return crud.ReadResource(sync)
}

type FileSystemsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_file_storage.FileStorageClient
	Res    *oci_file_storage.ListFileSystemsResponse
}

func (s *FileSystemsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FileSystemsDataSourceCrud) Get() error {
	request := oci_file_storage.ListFileSystemsRequest{}

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
		request.LifecycleState = oci_file_storage.ListFileSystemsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "file_storage")

	response, err := s.Client.ListFileSystems(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListFileSystems(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FileSystemsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		fileSystem := map[string]interface{}{
			"availability_domain": *r.AvailabilityDomain,
			"compartment_id":      *r.CompartmentId,
		}

		if r.DisplayName != nil {
			fileSystem["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			fileSystem["id"] = *r.Id
		}

		if r.MeteredBytes != nil {
			fileSystem["metered_bytes"] = *r.MeteredBytes
		}

		fileSystem["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			fileSystem["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, fileSystem)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, FileSystemsDataSource().Schema["file_systems"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("file_systems", resources); err != nil {
		panic(err)
	}

	return
}
