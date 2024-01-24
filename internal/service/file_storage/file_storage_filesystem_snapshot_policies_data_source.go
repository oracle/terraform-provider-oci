// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FileStorageFilesystemSnapshotPoliciesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFileStorageFilesystemSnapshotPolicies,
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"filesystem_snapshot_policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(FileStorageFilesystemSnapshotPolicyResource()),
			},
		},
	}
}

func readFileStorageFilesystemSnapshotPolicies(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageFilesystemSnapshotPoliciesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.ReadResource(sync)
}

type FileStorageFilesystemSnapshotPoliciesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_file_storage.FileStorageClient
	Res    *oci_file_storage.ListFilesystemSnapshotPoliciesResponse
}

func (s *FileStorageFilesystemSnapshotPoliciesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FileStorageFilesystemSnapshotPoliciesDataSourceCrud) Get() error {
	request := oci_file_storage.ListFilesystemSnapshotPoliciesRequest{}

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
		request.LifecycleState = oci_file_storage.ListFilesystemSnapshotPoliciesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "file_storage")

	response, err := s.Client.ListFilesystemSnapshotPolicies(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListFilesystemSnapshotPolicies(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FileStorageFilesystemSnapshotPoliciesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FileStorageFilesystemSnapshotPoliciesDataSource-", FileStorageFilesystemSnapshotPoliciesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		filesystemSnapshotPolicy := map[string]interface{}{
			"availability_domain": *r.AvailabilityDomain,
			"compartment_id":      *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			filesystemSnapshotPolicy["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			filesystemSnapshotPolicy["display_name"] = *r.DisplayName
		}

		filesystemSnapshotPolicy["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			filesystemSnapshotPolicy["id"] = *r.Id
		}

		if r.PolicyPrefix != nil {
			filesystemSnapshotPolicy["policy_prefix"] = *r.PolicyPrefix
		}

		filesystemSnapshotPolicy["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			filesystemSnapshotPolicy["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, filesystemSnapshotPolicy)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, FileStorageFilesystemSnapshotPoliciesDataSource().Schema["filesystem_snapshot_policies"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("filesystem_snapshot_policies", resources); err != nil {
		return err
	}

	return nil
}
