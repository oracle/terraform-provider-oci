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

func CapacityManagementOccmDemandSignalDeliveriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCapacityManagementOccmDemandSignalDeliveries,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"occm_demand_signal_item_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"occm_demand_signal_delivery_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"accepted_quantity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"demand_signal_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"demand_signal_item_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"justification": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
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
									"time_delivered": {
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

func readCapacityManagementOccmDemandSignalDeliveries(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccmDemandSignalDeliveriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DemandSignalClient()

	return tfresource.ReadResource(sync)
}

type CapacityManagementOccmDemandSignalDeliveriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_capacity_management.DemandSignalClient
	Res    *oci_capacity_management.ListOccmDemandSignalDeliveriesResponse
}

func (s *CapacityManagementOccmDemandSignalDeliveriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CapacityManagementOccmDemandSignalDeliveriesDataSourceCrud) Get() error {
	request := oci_capacity_management.ListOccmDemandSignalDeliveriesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if occmDemandSignalItemId, ok := s.D.GetOkExists("occm_demand_signal_item_id"); ok {
		tmp := occmDemandSignalItemId.(string)
		request.OccmDemandSignalItemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "capacity_management")

	response, err := s.Client.ListOccmDemandSignalDeliveries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOccmDemandSignalDeliveries(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CapacityManagementOccmDemandSignalDeliveriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CapacityManagementOccmDemandSignalDeliveriesDataSource-", CapacityManagementOccmDemandSignalDeliveriesDataSource(), s.D))
	resources := []map[string]interface{}{}
	occmDemandSignalDelivery := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OccmDemandSignalDeliverySummaryToMap(item))
	}
	occmDemandSignalDelivery["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CapacityManagementOccmDemandSignalDeliveriesDataSource().Schema["occm_demand_signal_delivery_collection"].Elem.(*schema.Resource).Schema)
		occmDemandSignalDelivery["items"] = items
	}

	resources = append(resources, occmDemandSignalDelivery)
	if err := s.D.Set("occm_demand_signal_delivery_collection", resources); err != nil {
		return err
	}

	return nil
}

func OccmDemandSignalDeliverySummaryToMap(obj oci_capacity_management.OccmDemandSignalDeliverySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AcceptedQuantity != nil {
		result["accepted_quantity"] = strconv.FormatInt(*obj.AcceptedQuantity, 10)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DemandSignalId != nil {
		result["demand_signal_id"] = string(*obj.DemandSignalId)
	}

	if obj.DemandSignalItemId != nil {
		result["demand_signal_item_id"] = string(*obj.DemandSignalItemId)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Justification != nil {
		result["justification"] = string(*obj.Justification)
	}

	result["lifecycle_details"] = string(obj.LifecycleDetails)

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeDelivered != nil {
		result["time_delivered"] = obj.TimeDelivered.String()
	}

	return result
}
