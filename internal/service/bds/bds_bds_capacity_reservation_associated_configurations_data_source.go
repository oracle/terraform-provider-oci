// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BdsBdsCapacityReservationAssociatedConfigurationsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readBdsBdsCapacityReservationAssociatedConfigurationsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"bds_capacity_reservation_id": {
				Type:     schema.TypeString,
				Required: true,
			},
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
			"bds_capacity_reservation_associated_configuration_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bds_instance_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
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
							},
						},
					},
				},
			},
		},
	}
}

func readBdsBdsCapacityReservationAssociatedConfigurationsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BdsBdsCapacityReservationAssociatedConfigurationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type BdsBdsCapacityReservationAssociatedConfigurationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.ListBdsCapacityReservationAssociatedConfigurationsResponse
}

func (s *BdsBdsCapacityReservationAssociatedConfigurationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsCapacityReservationAssociatedConfigurationsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_bds.ListBdsCapacityReservationAssociatedConfigurationsRequest{}

	if bdsCapacityReservationId, ok := s.D.GetOkExists("bds_capacity_reservation_id"); ok {
		tmp := bdsCapacityReservationId.(string)
		request.BdsCapacityReservationId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_bds.BdsCapacityReservationConfigurationLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.ListBdsCapacityReservationAssociatedConfigurations(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBdsCapacityReservationAssociatedConfigurations(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BdsBdsCapacityReservationAssociatedConfigurationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BdsBdsCapacityReservationAssociatedConfigurationsDataSource-", BdsBdsCapacityReservationAssociatedConfigurationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	bdsCapacityReservationAssociatedConfiguration := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, BdsCapacityReservationAssociatedConfigurationSummaryToMap(item))
	}
	bdsCapacityReservationAssociatedConfiguration["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, BdsBdsCapacityReservationAssociatedConfigurationsDataSource().Schema["bds_capacity_reservation_associated_configuration_collection"].Elem.(*schema.Resource).Schema)
		bdsCapacityReservationAssociatedConfiguration["items"] = items
	}

	resources = append(resources, bdsCapacityReservationAssociatedConfiguration)
	if err := s.D.Set("bds_capacity_reservation_associated_configuration_collection", resources); err != nil {
		return err
	}

	return nil
}

func BdsCapacityReservationAssociatedConfigurationSummaryToMap(obj oci_bds.BdsCapacityReservationAssociatedConfigurationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BdsInstanceId != nil {
		result["bds_instance_id"] = string(*obj.BdsInstanceId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
