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

func DatabaseExascaleDbStorageVaultDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["exascale_db_storage_vault_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseExascaleDbStorageVaultResource(), fieldMap, readSingularDatabaseExascaleDbStorageVault)
}

func readSingularDatabaseExascaleDbStorageVault(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExascaleDbStorageVaultDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseExascaleDbStorageVaultDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetExascaleDbStorageVaultResponse
}

func (s *DatabaseExascaleDbStorageVaultDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseExascaleDbStorageVaultDataSourceCrud) Get() error {
	request := oci_database.GetExascaleDbStorageVaultRequest{}

	if exascaleDbStorageVaultId, ok := s.D.GetOkExists("exascale_db_storage_vault_id"); ok {
		tmp := exascaleDbStorageVaultId.(string)
		request.ExascaleDbStorageVaultId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetExascaleDbStorageVault(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseExascaleDbStorageVaultDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AdditionalFlashCacheInPercent != nil {
		s.D.Set("additional_flash_cache_in_percent", *s.Res.AdditionalFlashCacheInPercent)
	}

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HighCapacityDatabaseStorage != nil {
		s.D.Set("high_capacity_database_storage", []interface{}{ExascaleDbStorageDetailsToMap(s.Res.HighCapacityDatabaseStorage)})
	} else {
		s.D.Set("high_capacity_database_storage", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeZone != nil {
		s.D.Set("time_zone", *s.Res.TimeZone)
	}

	if s.Res.VmClusterCount != nil {
		s.D.Set("vm_cluster_count", *s.Res.VmClusterCount)
	}

	s.D.Set("vm_cluster_ids", s.Res.VmClusterIds)

	return nil
}
