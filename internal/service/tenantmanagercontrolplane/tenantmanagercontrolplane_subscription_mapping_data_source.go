// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tenantmanagercontrolplane

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_tenantmanagercontrolplane "github.com/oracle/oci-go-sdk/v65/tenantmanagercontrolplane"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TenantmanagercontrolplaneSubscriptionMappingDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["subscription_mapping_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(TenantmanagercontrolplaneSubscriptionMappingResource(), fieldMap, readSingularTenantmanagercontrolplaneSubscriptionMapping)
}

func readSingularTenantmanagercontrolplaneSubscriptionMapping(d *schema.ResourceData, m interface{}) error {
	sync := &TenantmanagercontrolplaneSubscriptionMappingDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OrganizationsSubscriptionClient()

	return tfresource.ReadResource(sync)
}

type TenantmanagercontrolplaneSubscriptionMappingDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_tenantmanagercontrolplane.SubscriptionClient
	Res    *oci_tenantmanagercontrolplane.GetSubscriptionMappingResponse
}

func (s *TenantmanagercontrolplaneSubscriptionMappingDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *TenantmanagercontrolplaneSubscriptionMappingDataSourceCrud) Get() error {
	request := oci_tenantmanagercontrolplane.GetSubscriptionMappingRequest{}

	if subscriptionMappingId, ok := s.D.GetOkExists("subscription_mapping_id"); ok {
		tmp := subscriptionMappingId.(string)
		request.SubscriptionMappingId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "tenantmanagercontrolplane")

	response, err := s.Client.GetSubscriptionMapping(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *TenantmanagercontrolplaneSubscriptionMappingDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.IsExplicitlyAssigned != nil {
		s.D.Set("is_explicitly_assigned", *s.Res.IsExplicitlyAssigned)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubscriptionId != nil {
		s.D.Set("subscription_id", *s.Res.SubscriptionId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeTerminated != nil {
		s.D.Set("time_terminated", s.Res.TimeTerminated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
