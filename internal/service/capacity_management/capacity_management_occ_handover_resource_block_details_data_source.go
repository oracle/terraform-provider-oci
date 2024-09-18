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

func CapacityManagementOccHandoverResourceBlockDetailsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCapacityManagementOccHandoverResourceBlockDetails,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"host_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"occ_handover_resource_block_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"occ_handover_resource_block_detail_collection": {
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
									"details": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"occ_resource_handover_block_id": {
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

func readCapacityManagementOccHandoverResourceBlockDetails(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccHandoverResourceBlockDetailsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()

	return tfresource.ReadResource(sync)
}

type CapacityManagementOccHandoverResourceBlockDetailsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_capacity_management.CapacityManagementClient
	Res    *oci_capacity_management.ListOccHandoverResourceBlockDetailsResponse
}

func (s *CapacityManagementOccHandoverResourceBlockDetailsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CapacityManagementOccHandoverResourceBlockDetailsDataSourceCrud) Get() error {
	request := oci_capacity_management.ListOccHandoverResourceBlockDetailsRequest{}

	if hostId, ok := s.D.GetOkExists("host_id"); ok {
		tmp := hostId.(string)
		request.HostId = &tmp
	}

	if occHandoverResourceBlockId, ok := s.D.GetOkExists("occ_handover_resource_block_id"); ok {
		tmp := occHandoverResourceBlockId.(string)
		request.OccHandoverResourceBlockId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "capacity_management")

	response, err := s.Client.ListOccHandoverResourceBlockDetails(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOccHandoverResourceBlockDetails(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CapacityManagementOccHandoverResourceBlockDetailsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CapacityManagementOccHandoverResourceBlockDetailsDataSource-", CapacityManagementOccHandoverResourceBlockDetailsDataSource(), s.D))
	resources := []map[string]interface{}{}
	occHandoverResourceBlockDetail := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OccHandoverResourceBlockDetailSummaryToMap(item))
	}
	occHandoverResourceBlockDetail["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CapacityManagementOccHandoverResourceBlockDetailsDataSource().Schema["occ_handover_resource_block_detail_collection"].Elem.(*schema.Resource).Schema)
		occHandoverResourceBlockDetail["items"] = items
	}

	resources = append(resources, occHandoverResourceBlockDetail)
	if err := s.D.Set("occ_handover_resource_block_detail_collection", resources); err != nil {
		return err
	}

	return nil
}

func OccHandoverResourceBlockDetailSummaryToMap(obj oci_capacity_management.OccHandoverResourceBlockDetailSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["details"] = obj.Details

	if obj.OccResourceHandoverBlockId != nil {
		result["occ_resource_handover_block_id"] = string(*obj.OccResourceHandoverBlockId)
	}

	return result
}
