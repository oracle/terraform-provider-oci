// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package capacity_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CapacityManagementOccAvailabilityCatalogDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["occ_availability_catalog_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CapacityManagementOccAvailabilityCatalogResource(), fieldMap, readSingularCapacityManagementOccAvailabilityCatalog)
}

func readSingularCapacityManagementOccAvailabilityCatalog(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccAvailabilityCatalogDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()

	return tfresource.ReadResource(sync)
}

type CapacityManagementOccAvailabilityCatalogDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_capacity_management.CapacityManagementClient
	Res    *oci_capacity_management.GetOccAvailabilityCatalogResponse
}

func (s *CapacityManagementOccAvailabilityCatalogDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CapacityManagementOccAvailabilityCatalogDataSourceCrud) Get() error {
	request := oci_capacity_management.GetOccAvailabilityCatalogRequest{}

	if occAvailabilityCatalogId, ok := s.D.GetOkExists("occ_availability_catalog_id"); ok {
		tmp := occAvailabilityCatalogId.(string)
		request.OccAvailabilityCatalogId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "capacity_management")

	response, err := s.Client.GetOccAvailabilityCatalog(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CapacityManagementOccAvailabilityCatalogDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("catalog_state", s.Res.CatalogState)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	details := []interface{}{}
	for _, item := range s.Res.Details {
		details = append(details, OccAvailabilitySummaryToMap(item))
	}
	s.D.Set("details", details)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MetadataDetails != nil {
		s.D.Set("metadata_details", []interface{}{MetadataDetailsToMap(s.Res.MetadataDetails)})
	} else {
		s.D.Set("metadata_details", nil)
	}

	s.D.Set("namespace", s.Res.Namespace)

	if s.Res.OccCustomerGroupId != nil {
		s.D.Set("occ_customer_group_id", *s.Res.OccCustomerGroupId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
