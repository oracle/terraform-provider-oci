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

func CapacityManagementInternalNamespaceOccOverviewsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCapacityManagementInternalNamespaceOccOverviews,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"from": {
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
			"to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"workload_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"occ_overview_collection": {
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
									"capacity_requests_blob": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"period_value": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_available": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_cancelled": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_demanded": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_rejected": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_supplied": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_unfulfilled": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"unit": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"workload_type_breakdown_blob": {
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

func readCapacityManagementInternalNamespaceOccOverviews(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementInternalNamespaceOccOverviewsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()

	return tfresource.ReadResource(sync)
}

type CapacityManagementInternalNamespaceOccOverviewsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_capacity_management.CapacityManagementClient
	Res    *oci_capacity_management.ListInternalNamespaceOccOverviewsResponse
}

func (s *CapacityManagementInternalNamespaceOccOverviewsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CapacityManagementInternalNamespaceOccOverviewsDataSourceCrud) Get() error {
	request := oci_capacity_management.ListInternalNamespaceOccOverviewsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if from, ok := s.D.GetOkExists("from"); ok {
		tmp, err := time.Parse(time.RFC3339, from.(string))
		if err != nil {
			return err
		}
		request.From = &oci_common.SDKTime{Time: tmp}
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		request.Namespace = oci_capacity_management.ListInternalNamespaceOccOverviewsNamespaceEnum(namespace.(string))
	}

	if occCustomerGroupId, ok := s.D.GetOkExists("occ_customer_group_id"); ok {
		tmp := occCustomerGroupId.(string)
		request.OccCustomerGroupId = &tmp
	}

	if to, ok := s.D.GetOkExists("to"); ok {
		tmp, err := time.Parse(time.RFC3339, to.(string))
		if err != nil {
			return err
		}
		request.To = &oci_common.SDKTime{Time: tmp}
	}

	if workloadType, ok := s.D.GetOkExists("workload_type"); ok {
		tmp := workloadType.(string)
		request.WorkloadType = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "capacity_management")

	response, err := s.Client.ListInternalNamespaceOccOverviews(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInternalNamespaceOccOverviews(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CapacityManagementInternalNamespaceOccOverviewsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CapacityManagementInternalNamespaceOccOverviewsDataSource-", CapacityManagementInternalNamespaceOccOverviewsDataSource(), s.D))
	resources := []map[string]interface{}{}
	internalNamespaceOccOverview := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OccOverviewSummaryToMap(item))
	}
	internalNamespaceOccOverview["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CapacityManagementInternalNamespaceOccOverviewsDataSource().Schema["occ_overview_collection"].Elem.(*schema.Resource).Schema)
		internalNamespaceOccOverview["items"] = items
	}

	resources = append(resources, internalNamespaceOccOverview)
	if err := s.D.Set("occ_overview_collection", resources); err != nil {
		return err
	}

	return nil
}
