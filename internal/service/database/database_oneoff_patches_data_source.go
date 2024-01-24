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

func DatabaseOneoffPatchesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseOneoffPatches,
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oneoff_patches": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseOneoffPatchResource()),
			},
		},
	}
}

func readDatabaseOneoffPatches(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseOneoffPatchesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseOneoffPatchesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListOneoffPatchesResponse
}

func (s *DatabaseOneoffPatchesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseOneoffPatchesDataSourceCrud) Get() error {
	request := oci_database.ListOneoffPatchesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.OneoffPatchSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListOneoffPatches(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOneoffPatches(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseOneoffPatchesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseOneoffPatchesDataSource-", DatabaseOneoffPatchesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		oneoffPatch := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DbVersion != nil {
			oneoffPatch["db_version"] = *r.DbVersion
		}

		if r.DefinedTags != nil {
			oneoffPatch["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			oneoffPatch["display_name"] = *r.DisplayName
		}

		oneoffPatch["freeform_tags"] = r.FreeformTags
		oneoffPatch["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			oneoffPatch["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			oneoffPatch["lifecycle_details"] = *r.LifecycleDetails
		}

		oneoffPatch["one_off_patches"] = r.OneOffPatches
		oneoffPatch["one_off_patches"] = r.OneOffPatches

		if r.ReleaseUpdate != nil {
			oneoffPatch["release_update"] = *r.ReleaseUpdate
		}

		if r.Sha256Sum != nil {
			oneoffPatch["sha256sum"] = *r.Sha256Sum
		}

		if r.SizeInKBs != nil {
			oneoffPatch["size_in_kbs"] = *r.SizeInKBs
		}

		oneoffPatch["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			oneoffPatch["time_created"] = r.TimeCreated.String()
		}

		if r.TimeOfExpiration != nil {
			oneoffPatch["time_of_expiration"] = r.TimeOfExpiration.String()
		}

		if r.TimeUpdated != nil {
			oneoffPatch["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, oneoffPatch)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseOneoffPatchesDataSource().Schema["oneoff_patches"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("oneoff_patches", resources); err != nil {
		return err
	}

	return nil
}
