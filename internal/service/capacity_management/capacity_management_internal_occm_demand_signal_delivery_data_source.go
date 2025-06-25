// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package capacity_management

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CapacityManagementInternalOccmDemandSignalDeliveryDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["occm_demand_signal_delivery_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CapacityManagementInternalOccmDemandSignalDeliveryResource(), fieldMap, readSingularCapacityManagementInternalOccmDemandSignalDelivery)
}

func readSingularCapacityManagementInternalOccmDemandSignalDelivery(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementInternalOccmDemandSignalDeliveryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InternalDemandSignalClient()

	return tfresource.ReadResource(sync)
}

type CapacityManagementInternalOccmDemandSignalDeliveryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_capacity_management.InternalDemandSignalClient
	Res    *oci_capacity_management.GetInternalOccmDemandSignalDeliveryResponse
}

func (s *CapacityManagementInternalOccmDemandSignalDeliveryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CapacityManagementInternalOccmDemandSignalDeliveryDataSourceCrud) Get() error {
	request := oci_capacity_management.GetInternalOccmDemandSignalDeliveryRequest{}

	if occmDemandSignalDeliveryId, ok := s.D.GetOkExists("occm_demand_signal_delivery_id"); ok {
		tmp := occmDemandSignalDeliveryId.(string)
		request.OccmDemandSignalDeliveryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "capacity_management")

	response, err := s.Client.GetInternalOccmDemandSignalDelivery(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CapacityManagementInternalOccmDemandSignalDeliveryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AcceptedQuantity != nil {
		s.D.Set("accepted_quantity", strconv.FormatInt(*s.Res.AcceptedQuantity, 10))
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DemandSignalId != nil {
		s.D.Set("demand_signal_id", *s.Res.DemandSignalId)
	}

	if s.Res.DemandSignalItemId != nil {
		s.D.Set("demand_signal_item_id", *s.Res.DemandSignalItemId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Justification != nil {
		s.D.Set("justification", *s.Res.Justification)
	}

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	if s.Res.Notes != nil {
		s.D.Set("notes", *s.Res.Notes)
	}

	if s.Res.OccCustomerGroupId != nil {
		s.D.Set("occ_customer_group_id", *s.Res.OccCustomerGroupId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeDelivered != nil {
		s.D.Set("time_delivered", s.Res.TimeDelivered.String())
	}

	return nil
}
