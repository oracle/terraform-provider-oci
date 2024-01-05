// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"
)

func OpsiOperationsInsightsWarehouseUserDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["operations_insights_warehouse_user_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OpsiOperationsInsightsWarehouseUserResource(), fieldMap, readSingularOpsiOperationsInsightsWarehouseUser)
}

func readSingularOpsiOperationsInsightsWarehouseUser(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOperationsInsightsWarehouseUserDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

type OpsiOperationsInsightsWarehouseUserDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.GetOperationsInsightsWarehouseUserResponse
}

func (s *OpsiOperationsInsightsWarehouseUserDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiOperationsInsightsWarehouseUserDataSourceCrud) Get() error {
	request := oci_opsi.GetOperationsInsightsWarehouseUserRequest{}

	if operationsInsightsWarehouseUserId, ok := s.D.GetOkExists("operations_insights_warehouse_user_id"); ok {
		tmp := operationsInsightsWarehouseUserId.(string)
		request.OperationsInsightsWarehouseUserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.GetOperationsInsightsWarehouseUser(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OpsiOperationsInsightsWarehouseUserDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionPassword != nil {
		s.D.Set("connection_password", *s.Res.ConnectionPassword)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsAwrDataAccess != nil {
		s.D.Set("is_awr_data_access", *s.Res.IsAwrDataAccess)
	}

	if s.Res.IsEmDataAccess != nil {
		s.D.Set("is_em_data_access", *s.Res.IsEmDataAccess)
	}

	if s.Res.IsOpsiDataAccess != nil {
		s.D.Set("is_opsi_data_access", *s.Res.IsOpsiDataAccess)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.OperationsInsightsWarehouseId != nil {
		s.D.Set("operations_insights_warehouse_id", *s.Res.OperationsInsightsWarehouseId)
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
