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

func JmsPluginErrorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsPluginErrors,
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
			"managed_instance_id": {
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
			"plugin_error_collection": {
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
									"agent_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
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
									"host_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"managed_instance_id": {
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

func readJmsPluginErrors(d *schema.ResourceData, m interface{}) error {
	sync := &JmsPluginErrorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsPluginErrorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.ListPluginErrorsResponse
}

func (s *JmsPluginErrorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsPluginErrorsDataSourceCrud) Get() error {
	request := oci_jms.ListPluginErrorsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
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

	response, err := s.Client.ListPluginErrors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPluginErrors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsPluginErrorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsPluginErrorsDataSource-", JmsPluginErrorsDataSource(), s.D))
	resources := []map[string]interface{}{}
	pluginError := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PluginErrorSummaryToMap(item))
	}
	pluginError["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsPluginErrorsDataSource().Schema["plugin_error_collection"].Elem.(*schema.Resource).Schema)
		pluginError["items"] = items
	}

	resources = append(resources, pluginError)
	if err := s.D.Set("plugin_error_collection", resources); err != nil {
		return err
	}

	return nil
}

func PluginErrorDetailsToMap(obj oci_jms.PluginErrorDetails) map[string]interface{} {
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

func PluginErrorSummaryToMap(obj oci_jms.PluginErrorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["agent_type"] = string(obj.AgentType)

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	errors := []interface{}{}
	for _, item := range obj.Errors {
		errors = append(errors, PluginErrorDetailsToMap(item))
	}
	result["errors"] = errors

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	if obj.ManagedInstanceId != nil {
		result["managed_instance_id"] = string(*obj.ManagedInstanceId)
	}

	if obj.TimeFirstSeen != nil {
		result["time_first_seen"] = obj.TimeFirstSeen.String()
	}

	if obj.TimeLastSeen != nil {
		result["time_last_seen"] = obj.TimeLastSeen.String()
	}

	return result
}
