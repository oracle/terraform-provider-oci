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

func CapacityManagementInternalOccmDemandSignalItemsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCapacityManagementInternalOccmDemandSignalItems,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"demand_signal_namespace": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"occ_customer_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"occm_demand_signal_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"internal_occm_demand_signal_item_collection": {
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
									"availability_domain": {
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
									"demand_signal_catalog_resource_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"demand_signal_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"demand_signal_namespace": {
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
									"notes": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"occ_customer_group_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"quantity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"region": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"request_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_properties": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
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
									"target_compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_needed_before": {
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

func readCapacityManagementInternalOccmDemandSignalItems(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementInternalOccmDemandSignalItemsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InternalDemandSignalClient()

	return tfresource.ReadResource(sync)
}

type CapacityManagementInternalOccmDemandSignalItemsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_capacity_management.InternalDemandSignalClient
	Res    *oci_capacity_management.ListInternalOccmDemandSignalItemsResponse
}

func (s *CapacityManagementInternalOccmDemandSignalItemsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CapacityManagementInternalOccmDemandSignalItemsDataSourceCrud) Get() error {
	request := oci_capacity_management.ListInternalOccmDemandSignalItemsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if demandSignalNamespace, ok := s.D.GetOkExists("demand_signal_namespace"); ok {
		request.DemandSignalNamespace = oci_capacity_management.ListInternalOccmDemandSignalItemsDemandSignalNamespaceEnum(demandSignalNamespace.(string))
	}

	if occCustomerGroupId, ok := s.D.GetOkExists("occ_customer_group_id"); ok {
		tmp := occCustomerGroupId.(string)
		request.OccCustomerGroupId = &tmp
	}

	if occmDemandSignalId, ok := s.D.GetOkExists("occm_demand_signal_id"); ok {
		tmp := occmDemandSignalId.(string)
		request.OccmDemandSignalId = &tmp
	}

	if resourceName, ok := s.D.GetOkExists("resource_name"); ok {
		tmp := resourceName.(string)
		request.ResourceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "capacity_management")

	response, err := s.Client.ListInternalOccmDemandSignalItems(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInternalOccmDemandSignalItems(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CapacityManagementInternalOccmDemandSignalItemsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CapacityManagementInternalOccmDemandSignalItemsDataSource-", CapacityManagementInternalOccmDemandSignalItemsDataSource(), s.D))
	resources := []map[string]interface{}{}
	internalOccmDemandSignalItem := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, InternalOccmDemandSignalItemSummaryToMap(item))
	}
	internalOccmDemandSignalItem["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CapacityManagementInternalOccmDemandSignalItemsDataSource().Schema["internal_occm_demand_signal_item_collection"].Elem.(*schema.Resource).Schema)
		internalOccmDemandSignalItem["items"] = items
	}

	resources = append(resources, internalOccmDemandSignalItem)
	if err := s.D.Set("internal_occm_demand_signal_item_collection", resources); err != nil {
		return err
	}

	return nil
}

func InternalOccmDemandSignalItemSummaryToMap(obj oci_capacity_management.InternalOccmDemandSignalItemSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DemandSignalCatalogResourceId != nil {
		result["demand_signal_catalog_resource_id"] = string(*obj.DemandSignalCatalogResourceId)
	}

	if obj.DemandSignalId != nil {
		result["demand_signal_id"] = string(*obj.DemandSignalId)
	}

	result["demand_signal_namespace"] = string(obj.DemandSignalNamespace)

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Notes != nil {
		result["notes"] = string(*obj.Notes)
	}

	if obj.OccCustomerGroupId != nil {
		result["occ_customer_group_id"] = string(*obj.OccCustomerGroupId)
	}

	if obj.Quantity != nil {
		result["quantity"] = strconv.FormatInt(*obj.Quantity, 10)
	}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	result["request_type"] = string(obj.RequestType)

	if obj.ResourceName != nil {
		result["resource_name"] = string(*obj.ResourceName)
	}

	result["resource_properties"] = obj.ResourceProperties

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TargetCompartmentId != nil {
		result["target_compartment_id"] = string(*obj.TargetCompartmentId)
	}

	if obj.TimeNeededBefore != nil {
		result["time_needed_before"] = obj.TimeNeededBefore.String()
	}

	return result
}
