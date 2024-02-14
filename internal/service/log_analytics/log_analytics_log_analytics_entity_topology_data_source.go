// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"
)

func LogAnalyticsLogAnalyticsEntityTopologyDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularLogAnalyticsLogAnalyticsEntityTopology,
		Schema: map[string]*schema.Schema{
			"log_analytics_entity_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"metadata_equals": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"links": {
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
												"destination_entity_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"source_entity_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"nodes": {
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
												"are_logs_collected": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"cloud_resource_id": {
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
												"entity_type_internal_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"entity_type_name": {
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
												"lifecycle_details": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"management_agent_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"source_id": {
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
												"timezone_region": {
													Type:     schema.TypeString,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"are_logs_collected": {
																Type:     schema.TypeBool,
																Computed: true,
															},
															"cloud_resource_id": {
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
															"entity_type_internal_name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"entity_type_name": {
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
															"lifecycle_details": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"management_agent_id": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"metadata": {
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
																					"name": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"type": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"value": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"source_id": {
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
															"time_last_discovered": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"time_updated": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"timezone_region": {
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
							},
						},
					},
				},
			},
		},
	}
}

func readSingularLogAnalyticsLogAnalyticsEntityTopology(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsEntityTopologyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsLogAnalyticsEntityTopologyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.ListLogAnalyticsEntityTopologyResponse
}

func (s *LogAnalyticsLogAnalyticsEntityTopologyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsLogAnalyticsEntityTopologyDataSourceCrud) Get() error {
	request := oci_log_analytics.ListLogAnalyticsEntityTopologyRequest{}

	if logAnalyticsEntityId, ok := s.D.GetOkExists("log_analytics_entity_id"); ok {
		tmp := logAnalyticsEntityId.(string)
		request.LogAnalyticsEntityId = &tmp
	}

	if metadataEquals, ok := s.D.GetOkExists("metadata_equals"); ok {
		interfaces := metadataEquals.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("metadata_equals") {
			request.MetadataEquals = tmp
		}
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_log_analytics.ListLogAnalyticsEntityTopologyLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.ListLogAnalyticsEntityTopology(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LogAnalyticsLogAnalyticsEntityTopologyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LogAnalyticsLogAnalyticsEntityTopologyDataSource-", LogAnalyticsLogAnalyticsEntityTopologyDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, LogAnalyticsEntityTopologySummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func LogAnalyticsEntityCollectionToMap(obj *oci_log_analytics.LogAnalyticsEntityCollection) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, LogAnalyticsEntitySummaryToMap(item))
	}
	result["items"] = items

	return result
}

func LogAnalyticsEntityTopologyLinkToMap(obj oci_log_analytics.LogAnalyticsEntityTopologyLink) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DestinationEntityId != nil {
		result["destination_entity_id"] = string(*obj.DestinationEntityId)
	}

	if obj.SourceEntityId != nil {
		result["source_entity_id"] = string(*obj.SourceEntityId)
	}

	return result
}

func LogAnalyticsEntityTopologyLinkCollectionToMap(obj *oci_log_analytics.LogAnalyticsEntityTopologyLinkCollection) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, LogAnalyticsEntityTopologyLinkToMap(item))
	}
	result["items"] = items

	return result
}

func LogAnalyticsEntityTopologySummaryToMap(obj oci_log_analytics.LogAnalyticsEntityTopologySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Links != nil {
		result["links"] = []interface{}{LogAnalyticsEntityTopologyLinkCollectionToMap(obj.Links)}
	}

	if obj.Nodes != nil {
		result["nodes"] = []interface{}{LogAnalyticsEntityCollectionToMap(obj.Nodes)}
	}

	return result
}
