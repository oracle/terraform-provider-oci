// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseDbNodeSnapshotsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseDbNodeSnapshots,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_dbnode_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dbnode_snapshots": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseDbNodeSnapshotResource()),
			},
		},
	}
}

func readDatabaseDbNodeSnapshots(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbNodeSnapshotsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDbNodeSnapshotsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbnodeSnapshotsResponse
}

func (s *DatabaseDbNodeSnapshotsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbNodeSnapshotsDataSourceCrud) Get() error {
	request := oci_database.ListDbnodeSnapshotsRequest{}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if sourceDbnodeId, ok := s.D.GetOkExists("source_dbnode_id"); ok {
		tmp := sourceDbnodeId.(string)
		request.SourceDbnodeId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.DbnodeSnapshotLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListDbnodeSnapshots(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDbnodeSnapshots(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseDbNodeSnapshotsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseDbNodeSnapshotsDataSource-", DatabaseDbNodeSnapshotsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbNodeSnapshot := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.ClusterId != nil {
			dbNodeSnapshot["cluster_id"] = *r.ClusterId
		}

		if r.DefinedTags != nil {
			dbNodeSnapshot["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		dbNodeSnapshot["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			dbNodeSnapshot["id"] = *r.Id
			dbNodeSnapshot["dbnode_snapshot_id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			dbNodeSnapshot["lifecycle_details"] = *r.LifecycleDetails
		}

		mountPoints := []interface{}{}
		for _, item := range r.MountPoints {
			mountPoints = append(mountPoints, MountPointDetailsToMap(item))
		}
		dbNodeSnapshot["mount_points"] = mountPoints

		if len(mountPoints) > 0 {
			mountPoint0 := mountPoints[0].(map[string]interface{})
			if tmp, exist := mountPoint0["db_node_id"]; exist {
				dbNodeId := tmp.(string)
				dbNodeSnapshot["mount_dbnode_id"] = dbNodeId
			}
		} else {
			dbNodeSnapshot["mount_dbnode_id"] = "null"
		}

		if r.Name != nil {
			dbNodeSnapshot["name"] = *r.Name
		}

		if r.SourceDbnodeId != nil {
			dbNodeSnapshot["source_dbnode_id"] = *r.SourceDbnodeId
		}

		dbNodeSnapshot["state"] = r.LifecycleState

		if r.SystemTags != nil {
			dbNodeSnapshot["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			dbNodeSnapshot["time_created"] = r.TimeCreated.String()
		}

		volumes := []interface{}{}
		for _, item := range r.Volumes {
			volumes = append(volumes, VolumeDetailsToMap(item))
		}
		dbNodeSnapshot["volumes"] = volumes

		resources = append(resources, dbNodeSnapshot)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseDbNodeSnapshotsDataSource().Schema["dbnode_snapshots"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("dbnode_snapshots", resources); err != nil {
		return err
	}

	return nil
}

func MountPointDetailsToMap(obj oci_database.MountPointDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DbNodeId != nil {
		result["db_node_id"] = string(*obj.DbNodeId)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func VolumeDetailsToMap(obj oci_database.VolumeDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Size != nil {
		result["size"] = int(*obj.Size)
	}

	return result
}
