// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package capacity_management

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CapacityManagementOccmDemandSignalItemDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["occm_demand_signal_item_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CapacityManagementOccmDemandSignalItemResource(), fieldMap, readSingularCapacityManagementOccmDemandSignalItem)
}

func readSingularCapacityManagementOccmDemandSignalItem(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccmDemandSignalItemDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DemandSignalClient()

	return tfresource.ReadResource(sync)
}

type CapacityManagementOccmDemandSignalItemDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_capacity_management.DemandSignalClient
	Res    *oci_capacity_management.GetOccmDemandSignalItemResponse
}

func (s *CapacityManagementOccmDemandSignalItemDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CapacityManagementOccmDemandSignalItemDataSourceCrud) Get() error {
	request := oci_capacity_management.GetOccmDemandSignalItemRequest{}

	if occmDemandSignalItemId, ok := s.D.GetOkExists("occm_demand_signal_item_id"); ok {
		tmp := occmDemandSignalItemId.(string)
		request.OccmDemandSignalItemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "capacity_management")

	response, err := s.Client.GetOccmDemandSignalItem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CapacityManagementOccmDemandSignalItemDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DemandQuantity != nil {
		s.D.Set("demand_quantity", strconv.FormatInt(*s.Res.DemandQuantity, 10))
	}

	if s.Res.DemandSignalCatalogResourceId != nil {
		s.D.Set("demand_signal_catalog_resource_id", *s.Res.DemandSignalCatalogResourceId)
	}

	if s.Res.DemandSignalId != nil {
		s.D.Set("demand_signal_id", *s.Res.DemandSignalId)
	}

	s.D.Set("demand_signal_namespace", s.Res.DemandSignalNamespace)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Notes != nil {
		s.D.Set("notes", *s.Res.Notes)
	}

	if s.Res.Region != nil {
		s.D.Set("region", *s.Res.Region)
	}

	s.D.Set("request_type", s.Res.RequestType)

	if s.Res.ResourceName != nil {
		s.D.Set("resource_name", *s.Res.ResourceName)
	}

	s.D.Set("resource_properties", s.Res.ResourceProperties)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetCompartmentId != nil {
		s.D.Set("target_compartment_id", *s.Res.TargetCompartmentId)
	}

	if s.Res.TimeNeededBefore != nil {
		s.D.Set("time_needed_before", s.Res.TimeNeededBefore.Format(time.RFC3339Nano))
	}

	return nil
}
