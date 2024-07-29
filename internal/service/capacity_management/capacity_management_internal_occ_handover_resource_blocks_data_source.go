// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package capacity_management

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CapacityManagementInternalOccHandoverResourceBlocksDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCapacityManagementInternalOccHandoverResourceBlocks,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
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
				Required: true,
			},
			"occ_customer_group_id": {
				Type:     schema.TypeString,
				Required: true,
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

func readCapacityManagementInternalOccHandoverResourceBlocks(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementInternalOccHandoverResourceBlocksDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()

	return tfresource.ReadResource(sync)
}

type CapacityManagementInternalOccHandoverResourceBlocksDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_capacity_management.CapacityManagementClient
	Res    *oci_capacity_management.ListInternalOccHandoverResourceBlocksResponse
}

func (s *CapacityManagementInternalOccHandoverResourceBlocksDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CapacityManagementInternalOccHandoverResourceBlocksDataSourceCrud) Get() error {
	request := oci_capacity_management.ListInternalOccHandoverResourceBlocksRequest{}

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
		request.Namespace = oci_capacity_management.ListInternalOccHandoverResourceBlocksNamespaceEnum(namespace.(string))
	}

	if occCustomerGroupId, ok := s.D.GetOkExists("occ_customer_group_id"); ok {
		tmp := occCustomerGroupId.(string)
		request.OccCustomerGroupId = &tmp
	}

	if occHandoverResourceBlockId, ok := s.D.GetOkExists("occ_handover_resource_block_id"); ok {
		tmp := occHandoverResourceBlockId.(string)
		request.OccHandoverResourceBlockId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "capacity_management")

	response, err := s.Client.ListInternalOccHandoverResourceBlocks(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInternalOccHandoverResourceBlocks(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CapacityManagementInternalOccHandoverResourceBlocksDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CapacityManagementInternalOccHandoverResourceBlocksDataSource-", CapacityManagementInternalOccHandoverResourceBlocksDataSource(), s.D))
	resources := []map[string]interface{}{}
	internalOccHandoverResourceBlock := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OccHandoverResourceBlockSummaryToMap(item))
	}
	internalOccHandoverResourceBlock["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CapacityManagementInternalOccHandoverResourceBlocksDataSource().Schema["occ_handover_resource_block_collection"].Elem.(*schema.Resource).Schema)
		internalOccHandoverResourceBlock["items"] = items
	}

	resources = append(resources, internalOccHandoverResourceBlock)
	if err := s.D.Set("occ_handover_resource_block_collection", resources); err != nil {
		return err
	}

	return nil
}
