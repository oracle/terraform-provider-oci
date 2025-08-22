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

func DatabasePluggableDatabaseSnapshotsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabasePluggableDatabaseSnapshots,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pluggable_database_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pluggable_database_snapshots": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabasePluggableDatabaseSnapshotResource()),
			},
		},
	}
}

func readDatabasePluggableDatabaseSnapshots(d *schema.ResourceData, m interface{}) error {
	sync := &DatabasePluggableDatabaseSnapshotsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabasePluggableDatabaseSnapshotsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListPluggableDatabaseSnapshotsResponse
}

func (s *DatabasePluggableDatabaseSnapshotsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabasePluggableDatabaseSnapshotsDataSourceCrud) Get() error {
	request := oci_database.ListPluggableDatabaseSnapshotsRequest{}

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

	if pluggableDatabaseId, ok := s.D.GetOkExists("pluggable_database_id"); ok {
		tmp := pluggableDatabaseId.(string)
		request.PluggableDatabaseId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.PluggableDatabaseSnapshotLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListPluggableDatabaseSnapshots(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPluggableDatabaseSnapshots(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabasePluggableDatabaseSnapshotsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabasePluggableDatabaseSnapshotsDataSource-", DatabasePluggableDatabaseSnapshotsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		pluggableDatabaseSnapshot := map[string]interface{}{}

		if r.ClusterId != nil {
			pluggableDatabaseSnapshot["cluster_id"] = *r.ClusterId
		}

		if r.CompartmentId != nil {
			pluggableDatabaseSnapshot["compartment_id"] = *r.CompartmentId
		}

		if r.DefinedTags != nil {
			pluggableDatabaseSnapshot["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		pluggableDatabaseSnapshot["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			pluggableDatabaseSnapshot["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			pluggableDatabaseSnapshot["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.Name != nil {
			pluggableDatabaseSnapshot["name"] = *r.Name
		}

		if r.PluggableDatabaseId != nil {
			pluggableDatabaseSnapshot["pluggable_database_id"] = *r.PluggableDatabaseId
		}

		pluggableDatabaseSnapshot["state"] = r.LifecycleState

		if r.SystemTags != nil {
			pluggableDatabaseSnapshot["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			pluggableDatabaseSnapshot["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, pluggableDatabaseSnapshot)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabasePluggableDatabaseSnapshotsDataSource().Schema["pluggable_database_snapshots"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("pluggable_database_snapshots", resources); err != nil {
		return err
	}

	return nil
}
