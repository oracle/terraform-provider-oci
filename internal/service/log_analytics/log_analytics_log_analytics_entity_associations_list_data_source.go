// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LogAnalyticsLogAnalyticsEntityAssociationsListDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLogAnalyticsLogAnalyticsEntityAssociationsList,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"direct_or_all_associations": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_analytics_entity_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"log_analytics_entity_collection": {
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
									"associated_sources_count": {
										Type:     schema.TypeInt,
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
	}
}

func readLogAnalyticsLogAnalyticsEntityAssociationsList(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsEntityAssociationsListDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsLogAnalyticsEntityAssociationsListDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.ListEntityAssociationsResponse
}

func (s *LogAnalyticsLogAnalyticsEntityAssociationsListDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsLogAnalyticsEntityAssociationsListDataSourceCrud) Get() error {
	request := oci_log_analytics.ListEntityAssociationsRequest{}

	if directOrAllAssociations, ok := s.D.GetOkExists("direct_or_all_associations"); ok {
		request.DirectOrAllAssociations = oci_log_analytics.ListEntityAssociationsDirectOrAllAssociationsEnum(directOrAllAssociations.(string))
	}

	if logAnalyticsEntityId, ok := s.D.GetOkExists("log_analytics_entity_id"); ok {
		tmp := logAnalyticsEntityId.(string)
		request.LogAnalyticsEntityId = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.ListEntityAssociations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListEntityAssociations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LogAnalyticsLogAnalyticsEntityAssociationsListDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LogAnalyticsLogAnalyticsEntityAssociationsListDataSource-", LogAnalyticsLogAnalyticsEntityAssociationsListDataSource(), s.D))
	resources := []map[string]interface{}{}
	logAnalyticsEntityAssociationsList := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, LogAnalyticsEntitySummaryToMap(item))
	}
	logAnalyticsEntityAssociationsList["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, LogAnalyticsLogAnalyticsEntityAssociationsListDataSource().Schema["log_analytics_entity_collection"].Elem.(*schema.Resource).Schema)
		logAnalyticsEntityAssociationsList["items"] = items
	}

	resources = append(resources, logAnalyticsEntityAssociationsList)
	if err := s.D.Set("log_analytics_entity_collection", resources); err != nil {
		return err
	}

	return nil
}
