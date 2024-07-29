// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package capacity_management

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CapacityManagementOccHandoverResourceBlocksDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCapacityManagementOccHandoverResourceBlocks,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"handover_date_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"handover_date_less_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"handover_resource_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"occ_handover_resource_block_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"occ_handover_resource_block_collection": {
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
									"associated_capacity_requests": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"handover_quantity": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"occ_capacity_request_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"handover_date": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"handover_resource_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"occ_customer_group_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"placement_details": {
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
												"block": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"building": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"region": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"room": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"workload_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"total_handover_quantity": {
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

func readCapacityManagementOccHandoverResourceBlocks(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccHandoverResourceBlocksDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()

	return tfresource.ReadResource(sync)
}

type CapacityManagementOccHandoverResourceBlocksDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_capacity_management.CapacityManagementClient
	Res    *oci_capacity_management.ListOccHandoverResourceBlocksResponse
}

func (s *CapacityManagementOccHandoverResourceBlocksDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CapacityManagementOccHandoverResourceBlocksDataSourceCrud) Get() error {
	request := oci_capacity_management.ListOccHandoverResourceBlocksRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if handoverDateGreaterThanOrEqualTo, ok := s.D.GetOkExists("handover_date_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, handoverDateGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.HandoverDateGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if handoverDateLessThanOrEqualTo, ok := s.D.GetOkExists("handover_date_less_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, handoverDateLessThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.HandoverDateLessThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if handoverResourceName, ok := s.D.GetOkExists("handover_resource_name"); ok {
		tmp := handoverResourceName.(string)
		request.HandoverResourceName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		request.Namespace = oci_capacity_management.ListOccHandoverResourceBlocksNamespaceEnum(namespace.(string))
	}

	if occHandoverResourceBlockId, ok := s.D.GetOkExists("occ_handover_resource_block_id"); ok {
		tmp := occHandoverResourceBlockId.(string)
		request.OccHandoverResourceBlockId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "capacity_management")

	response, err := s.Client.ListOccHandoverResourceBlocks(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOccHandoverResourceBlocks(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CapacityManagementOccHandoverResourceBlocksDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CapacityManagementOccHandoverResourceBlocksDataSource-", CapacityManagementOccHandoverResourceBlocksDataSource(), s.D))
	resources := []map[string]interface{}{}
	occHandoverResourceBlock := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OccHandoverResourceBlockSummaryToMap(item))
	}
	occHandoverResourceBlock["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CapacityManagementOccHandoverResourceBlocksDataSource().Schema["occ_handover_resource_block_collection"].Elem.(*schema.Resource).Schema)
		occHandoverResourceBlock["items"] = items
	}

	resources = append(resources, occHandoverResourceBlock)
	if err := s.D.Set("occ_handover_resource_block_collection", resources); err != nil {
		return err
	}

	return nil
}

func AssociatedCapacityRequestDetailsToMap(obj oci_capacity_management.AssociatedCapacityRequestDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.HandoverQuantity != nil {
		result["handover_quantity"] = strconv.FormatInt(*obj.HandoverQuantity, 10)
	}

	if obj.OccCapacityRequestId != nil {
		result["occ_capacity_request_id"] = string(*obj.OccCapacityRequestId)
	}

	return result
}

func OccHandoverResourceBlockSummaryToMap(obj oci_capacity_management.OccHandoverResourceBlockSummary) map[string]interface{} {
	result := map[string]interface{}{}

	associatedCapacityRequests := []interface{}{}
	for _, item := range obj.AssociatedCapacityRequests {
		associatedCapacityRequests = append(associatedCapacityRequests, AssociatedCapacityRequestDetailsToMap(item))
	}
	result["associated_capacity_requests"] = associatedCapacityRequests

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.HandoverDate != nil {
		result["handover_date"] = obj.HandoverDate.String()
	}

	if obj.HandoverResourceName != nil {
		result["handover_resource_name"] = string(*obj.HandoverResourceName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["namespace"] = string(obj.Namespace)

	if obj.OccCustomerGroupId != nil {
		result["occ_customer_group_id"] = string(*obj.OccCustomerGroupId)
	}

	if obj.PlacementDetails != nil {
		result["placement_details"] = []interface{}{PlacementDetailsToMap(obj.PlacementDetails)}
	}

	if obj.TotalHandoverQuantity != nil {
		result["total_handover_quantity"] = strconv.FormatInt(*obj.TotalHandoverQuantity, 10)
	}

	return result
}

func PlacementDetailsToMap(obj *oci_capacity_management.PlacementDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.Block != nil {
		result["block"] = string(*obj.Block)
	}

	if obj.Building != nil {
		result["building"] = string(*obj.Building)
	}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	if obj.Room != nil {
		result["room"] = string(*obj.Room)
	}

	if obj.WorkloadType != nil {
		result["workload_type"] = string(*obj.WorkloadType)
	}

	return result
}
