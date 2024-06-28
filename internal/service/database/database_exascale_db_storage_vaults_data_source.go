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

func DatabaseExascaleDbStorageVaultsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseExascaleDbStorageVaults,
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
			"exascale_db_storage_vaults": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseExascaleDbStorageVaultResource()),
			},
		},
	}
}

func readDatabaseExascaleDbStorageVaults(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExascaleDbStorageVaultsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseExascaleDbStorageVaultsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListExascaleDbStorageVaultsResponse
}

func (s *DatabaseExascaleDbStorageVaultsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseExascaleDbStorageVaultsDataSourceCrud) Get() error {
	request := oci_database.ListExascaleDbStorageVaultsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.ExascaleDbStorageVaultLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListExascaleDbStorageVaults(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExascaleDbStorageVaults(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseExascaleDbStorageVaultsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseExascaleDbStorageVaultsDataSource-", DatabaseExascaleDbStorageVaultsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		exascaleDbStorageVault := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AdditionalFlashCacheInPercent != nil {
			exascaleDbStorageVault["additional_flash_cache_in_percent"] = *r.AdditionalFlashCacheInPercent
		}

		if r.AvailabilityDomain != nil {
			exascaleDbStorageVault["availability_domain"] = *r.AvailabilityDomain
		}

		if r.DefinedTags != nil {
			exascaleDbStorageVault["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			exascaleDbStorageVault["description"] = *r.Description
		}

		if r.DisplayName != nil {
			exascaleDbStorageVault["display_name"] = *r.DisplayName
		}

		exascaleDbStorageVault["freeform_tags"] = r.FreeformTags

		if r.HighCapacityDatabaseStorage != nil {
			exascaleDbStorageVault["high_capacity_database_storage"] = []interface{}{ExascaleDbStorageDetailsToMap(r.HighCapacityDatabaseStorage)}
		} else {
			exascaleDbStorageVault["high_capacity_database_storage"] = nil
		}

		if r.Id != nil {
			exascaleDbStorageVault["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			exascaleDbStorageVault["lifecycle_details"] = *r.LifecycleDetails
		}

		exascaleDbStorageVault["state"] = r.LifecycleState

		if r.SystemTags != nil {
			exascaleDbStorageVault["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			exascaleDbStorageVault["time_created"] = r.TimeCreated.String()
		}

		if r.TimeZone != nil {
			exascaleDbStorageVault["time_zone"] = *r.TimeZone
		}

		if r.VmClusterCount != nil {
			exascaleDbStorageVault["vm_cluster_count"] = *r.VmClusterCount
		}

		resources = append(resources, exascaleDbStorageVault)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseExascaleDbStorageVaultsDataSource().Schema["exascale_db_storage_vaults"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("exascale_db_storage_vaults", resources); err != nil {
		return err
	}

	return nil
}
