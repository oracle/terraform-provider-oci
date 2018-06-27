// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_file_storage "github.com/oracle/oci-go-sdk/filestorage"

	"github.com/oracle/terraform-provider-oci/crud"
)

func SnapshotsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSnapshots,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"file_system_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"snapshots": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     SnapshotResource(),
			},
		},
	}
}

func readSnapshots(d *schema.ResourceData, m interface{}) error {
	sync := &SnapshotsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return crud.ReadResource(sync)
}

type SnapshotsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_file_storage.FileStorageClient
	Res    *oci_file_storage.ListSnapshotsResponse
}

func (s *SnapshotsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *SnapshotsDataSourceCrud) Get() error {
	request := oci_file_storage.ListSnapshotsRequest{}

	if fileSystemId, ok := s.D.GetOkExists("file_system_id"); ok {
		tmp := fileSystemId.(string)
		request.FileSystemId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_file_storage.ListSnapshotsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "file_storage")

	response, err := s.Client.ListSnapshots(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSnapshots(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *SnapshotsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		snapshot := map[string]interface{}{
			"file_system_id": *r.FileSystemId,
		}

		if r.Id != nil {
			snapshot["id"] = *r.Id
		}

		if r.Name != nil {
			snapshot["name"] = *r.Name
		}

		snapshot["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			snapshot["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, snapshot)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, SnapshotsDataSource().Schema["snapshots"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("snapshots", resources); err != nil {
		panic(err)
	}

	return
}
