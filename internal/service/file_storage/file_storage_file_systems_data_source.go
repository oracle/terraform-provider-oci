// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import (
	"context"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_file_storage "github.com/oracle/oci-go-sdk/v56/filestorage"
)

func FileStorageFileSystemsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFileStorageFileSystems,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
			"parent_file_system_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_snapshot_id": {
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
				Elem:     tfresource.GetDataSourceItemSchema(FileStorageFileSystemResource()),
			},
		},
	}
}

func readFileStorageFileSystems(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageFileSystemsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.ReadResource(sync)
}

type FileStorageFileSystemsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_file_storage.FileStorageClient
	Res    *oci_file_storage.ListFileSystemsResponse
}

func (s *FileStorageFileSystemsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FileStorageFileSystemsDataSourceCrud) Get() error {
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

	if parentFileSystemId, ok := s.D.GetOkExists("parent_file_system_id"); ok {
		tmp := parentFileSystemId.(string)
		request.ParentFileSystemId = &tmp
	}

	if sourceSnapshotId, ok := s.D.GetOkExists("source_snapshot_id"); ok {
		tmp := sourceSnapshotId.(string)
		request.SourceSnapshotId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_file_storage.ListFileSystemsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "file_storage")

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

func (s *FileStorageFileSystemsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FileStorageFileSystemsDataSource-", FileStorageFileSystemsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		fileSystem := map[string]interface{}{
			"availability_domain": *r.AvailabilityDomain,
			"compartment_id":      *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			fileSystem["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			fileSystem["display_name"] = *r.DisplayName
		}

		fileSystem["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			fileSystem["id"] = *r.Id
		}

		if r.IsCloneParent != nil {
			fileSystem["is_clone_parent"] = *r.IsCloneParent
		}

		if r.IsHydrated != nil {
			fileSystem["is_hydrated"] = *r.IsHydrated
		}

		if r.KmsKeyId != nil {
			fileSystem["kms_key_id"] = *r.KmsKeyId
		}

		if r.LifecycleDetails != nil {
			fileSystem["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.MeteredBytes != nil {
			fileSystem["metered_bytes"] = strconv.FormatInt(*r.MeteredBytes, 10)
		}

		if r.SourceDetails != nil {
			fileSystem["source_details"] = []interface{}{FileSystemSourceDetailsToMap(r.SourceDetails)}
		} else {
			fileSystem["source_details"] = nil
		}

		fileSystem["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			fileSystem["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, fileSystem)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, FileStorageFileSystemsDataSource().Schema["file_systems"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("file_systems", resources); err != nil {
		return err
	}

	return nil
}
