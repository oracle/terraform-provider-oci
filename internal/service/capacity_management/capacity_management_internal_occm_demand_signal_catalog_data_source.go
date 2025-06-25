// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package capacity_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CapacityManagementInternalOccmDemandSignalCatalogDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCapacityManagementInternalOccmDemandSignalCatalog,
		Schema: map[string]*schema.Schema{
			"occm_demand_signal_catalog_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"occ_customer_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCapacityManagementInternalOccmDemandSignalCatalog(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementInternalOccmDemandSignalCatalogDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InternalDemandSignalClient()

	return tfresource.ReadResource(sync)
}

type CapacityManagementInternalOccmDemandSignalCatalogDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_capacity_management.InternalDemandSignalClient
	Res    *oci_capacity_management.GetInternalOccmDemandSignalCatalogResponse
}

func (s *CapacityManagementInternalOccmDemandSignalCatalogDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CapacityManagementInternalOccmDemandSignalCatalogDataSourceCrud) Get() error {
	request := oci_capacity_management.GetInternalOccmDemandSignalCatalogRequest{}

	if occmDemandSignalCatalogId, ok := s.D.GetOkExists("occm_demand_signal_catalog_id"); ok {
		tmp := occmDemandSignalCatalogId.(string)
		request.OccmDemandSignalCatalogId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "capacity_management")

	response, err := s.Client.GetInternalOccmDemandSignalCatalog(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CapacityManagementInternalOccmDemandSignalCatalogDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

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
