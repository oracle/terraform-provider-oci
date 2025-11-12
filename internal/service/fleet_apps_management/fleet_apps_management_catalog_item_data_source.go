// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementCatalogItemDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["catalog_item_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FleetAppsManagementCatalogItemResource(), fieldMap, readSingularFleetAppsManagementCatalogItem)
}

func readSingularFleetAppsManagementCatalogItem(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementCatalogItemDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementCatalogClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementCatalogItemDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementCatalogClient
	Res    *oci_fleet_apps_management.GetCatalogItemResponse
}

func (s *FleetAppsManagementCatalogItemDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementCatalogItemDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.GetCatalogItemRequest{}

	if catalogItemId, ok := s.D.GetOkExists("catalog_item_id"); ok {
		tmp := catalogItemId.(string)
		request.CatalogItemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.GetCatalogItem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FleetAppsManagementCatalogItemDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CatalogResultPayload != nil {
		catalogResultPayloadArray := []interface{}{}
		if catalogResultPayloadMap := CatalogResultPayloadToMap(&s.Res.CatalogResultPayload); catalogResultPayloadMap != nil {
			catalogResultPayloadArray = append(catalogResultPayloadArray, catalogResultPayloadMap)
		}
		s.D.Set("catalog_result_payload", catalogResultPayloadArray)
	} else {
		s.D.Set("catalog_result_payload", nil)
	}

	if s.Res.CatalogSourcePayload != nil {
		catalogSourcePayloadArray := []interface{}{}
		if catalogSourcePayloadMap := CatalogSourcePayloadToMap(&s.Res.CatalogSourcePayload); catalogSourcePayloadMap != nil {
			catalogSourcePayloadArray = append(catalogSourcePayloadArray, catalogSourcePayloadMap)
		}
		s.D.Set("catalog_source_payload", catalogSourcePayloadArray)
	} else {
		s.D.Set("catalog_source_payload", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("config_source_type", s.Res.ConfigSourceType)

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

	if s.Res.IsItemLocked != nil {
		s.D.Set("is_item_locked", *s.Res.IsItemLocked)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ListingId != nil {
		s.D.Set("listing_id", *s.Res.ListingId)
	}

	if s.Res.ListingVersion != nil {
		s.D.Set("listing_version", *s.Res.ListingVersion)
	}

	s.D.Set("package_type", s.Res.PackageType)

	if s.Res.ShortDescription != nil {
		s.D.Set("short_description", *s.Res.ShortDescription)
	}

	if s.Res.ShouldListPublicItems != nil {
		s.D.Set("should_list_public_items", *s.Res.ShouldListPublicItems)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeBackfillLastChecked != nil {
		s.D.Set("time_backfill_last_checked", s.Res.TimeBackfillLastChecked.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastChecked != nil {
		s.D.Set("time_last_checked", s.Res.TimeLastChecked.String())
	}

	if s.Res.TimeReleased != nil {
		s.D.Set("time_released", s.Res.TimeReleased.Format(time.RFC3339Nano))
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VersionDescription != nil {
		s.D.Set("version_description", *s.Res.VersionDescription)
	}

	return nil
}
