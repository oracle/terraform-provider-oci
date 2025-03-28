// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsFleetErrorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsFleetErrors,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"fleet_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_first_seen_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_first_seen_less_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_last_seen_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_last_seen_less_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fleet_error_collection": {
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
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"errors": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"details": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"reason": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_last_seen": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"fleet_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"fleet_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_first_seen": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_last_seen": {
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

func readJmsFleetErrors(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetErrorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsFleetErrorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.ListFleetErrorsResponse
}

func (s *JmsFleetErrorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsFleetErrorsDataSourceCrud) Get() error {
	request := oci_jms.ListFleetErrorsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if timeFirstSeenGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_first_seen_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeFirstSeenGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeFirstSeenGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeFirstSeenLessThanOrEqualTo, ok := s.D.GetOkExists("time_first_seen_less_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeFirstSeenLessThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeFirstSeenLessThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeLastSeenGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_last_seen_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeLastSeenGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeLastSeenGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeLastSeenLessThanOrEqualTo, ok := s.D.GetOkExists("time_last_seen_less_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeLastSeenLessThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeLastSeenLessThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.ListFleetErrors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListFleetErrors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsFleetErrorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsFleetErrorsDataSource-", JmsFleetErrorsDataSource(), s.D))
	resources := []map[string]interface{}{}
	fleetError := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, FleetErrorSummaryToMap(item))
	}
	fleetError["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsFleetErrorsDataSource().Schema["fleet_error_collection"].Elem.(*schema.Resource).Schema)
		fleetError["items"] = items
	}

	resources = append(resources, fleetError)
	if err := s.D.Set("fleet_error_collection", resources); err != nil {
		return err
	}

	return nil
}

func FleetErrorDetailsToMap(obj oci_jms.FleetErrorDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Details != nil {
		result["details"] = string(*obj.Details)
	}

	result["reason"] = string(obj.Reason)

	if obj.TimeLastSeen != nil {
		result["time_last_seen"] = obj.TimeLastSeen.String()
	}

	return result
}

func FleetErrorSummaryToMap(obj oci_jms.FleetErrorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	errors := []interface{}{}
	for _, item := range obj.Errors {
		errors = append(errors, FleetErrorDetailsToMap(item))
	}
	result["errors"] = errors

	if obj.FleetId != nil {
		result["fleet_id"] = string(*obj.FleetId)
	}

	if obj.FleetName != nil {
		result["fleet_name"] = string(*obj.FleetName)
	}

	if obj.TimeFirstSeen != nil {
		result["time_first_seen"] = obj.TimeFirstSeen.String()
	}

	if obj.TimeLastSeen != nil {
		result["time_last_seen"] = obj.TimeLastSeen.String()
	}

	return result
}
