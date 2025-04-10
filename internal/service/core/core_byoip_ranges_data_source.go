// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func CoreByoipRangesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreByoipRanges,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
			"byoip_range_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"cidr_block": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
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
									"ip_anycast_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ipv6cidr_block": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"monitor_ip": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									// Computed
									"byoip_range_vcn_ipv6allocations": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"compartment_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
												"byoip_range_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"ipv6cidr_block": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"vcn_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"origin_asn": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"as_path_prepend_length": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"asn": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"byoasn_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_advertised": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_validated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_withdrawn": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"validation_token": {
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

func readCoreByoipRanges(d *schema.ResourceData, m interface{}) error {
	sync := &CoreByoipRangesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreByoipRangesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListByoipRangesResponse
}

func (s *CoreByoipRangesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreByoipRangesDataSourceCrud) Get() error {
	request := oci_core.ListByoipRangesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		tmp := state.(string)
		request.LifecycleState = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListByoipRanges(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListByoipRanges(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreByoipRangesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreByoipRangesDataSource-", CoreByoipRangesDataSource(), s.D))
	resources := []map[string]interface{}{}
	byoipRange := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ByoipRangeSummaryToMap(item))
	}
	byoipRange["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CoreByoipRangesDataSource().Schema["byoip_range_collection"].Elem.(*schema.Resource).Schema)
		byoipRange["items"] = items
	}

	resources = append(resources, byoipRange)
	if err := s.D.Set("byoip_range_collection", resources); err != nil {
		return err
	}

	return nil
}

func ByoipRangeOriginAsnToMap(obj *oci_core.ByoipRangeOriginAsn) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AsPathPrependLength != nil {
		result["as_path_prepend_length"] = int(*obj.AsPathPrependLength)
	}

	if obj.Asn != nil {
		result["asn"] = strconv.FormatInt(*obj.Asn, 10)
	}

	if obj.ByoasnId != nil {
		result["byoasn_id"] = string(*obj.ByoasnId)
	}

	return result
}

func ByoipRangeSummaryToMap(obj oci_core.ByoipRangeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	byoipRangeVcnIpv6Allocations := []interface{}{}
	for _, item := range obj.ByoipRangeVcnIpv6Allocations {
		byoipRangeVcnIpv6Allocations = append(byoipRangeVcnIpv6Allocations, ByoipRangeVcnIpv6AllocationSummaryToMap(item))
	}
	result["byoip_range_vcn_ipv6allocations"] = byoipRangeVcnIpv6Allocations

	if obj.CidrBlock != nil {
		result["cidr_block"] = string(*obj.CidrBlock)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Ipv6CidrBlock != nil {
		result["ipv6cidr_block"] = string(*obj.Ipv6CidrBlock)
	}

	result["lifecycle_details"] = string(obj.LifecycleDetails)

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}
