// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package capacity_management

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CapacityManagementOccCapacityRequestDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["occ_capacity_request_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CapacityManagementOccCapacityRequestResource(), fieldMap, readSingularCapacityManagementOccCapacityRequest)
}

func readSingularCapacityManagementOccCapacityRequest(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccCapacityRequestDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()

	return tfresource.ReadResource(sync)
}

type CapacityManagementOccCapacityRequestDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_capacity_management.CapacityManagementClient
	Res    *oci_capacity_management.GetOccCapacityRequestResponse
}

func (s *CapacityManagementOccCapacityRequestDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CapacityManagementOccCapacityRequestDataSourceCrud) Get() error {
	request := oci_capacity_management.GetOccCapacityRequestRequest{}

	if occCapacityRequestId, ok := s.D.GetOkExists("occ_capacity_request_id"); ok {
		tmp := occCapacityRequestId.(string)
		request.OccCapacityRequestId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "capacity_management")

	response, err := s.Client.GetOccCapacityRequest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CapacityManagementOccCapacityRequestDataSourceCrud) SetData() error {
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

	if s.Res.DateExpectedCapacityHandover != nil {
		s.D.Set("date_expected_capacity_handover", s.Res.DateExpectedCapacityHandover.Format(time.RFC3339Nano))
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	details := []interface{}{}
	for _, item := range s.Res.Details {
		details = append(details, OccCapacityRequestBaseDetailsToMap(item))
	}
	s.D.Set("details", details)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("namespace", s.Res.Namespace)

	if s.Res.OccAvailabilityCatalogId != nil {
		s.D.Set("occ_availability_catalog_id", *s.Res.OccAvailabilityCatalogId)
	}

	if s.Res.OccCustomerGroupId != nil {
		s.D.Set("occ_customer_group_id", *s.Res.OccCustomerGroupId)
	}

	if s.Res.Region != nil {
		s.D.Set("region", *s.Res.Region)
	}

	s.D.Set("request_state", s.Res.RequestState)

	s.D.Set("request_type", s.Res.RequestType)

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
