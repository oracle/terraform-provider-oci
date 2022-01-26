// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package logging

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	oci_logging "github.com/oracle/oci-go-sdk/v56/logging"
)

func LoggingUnifiedAgentConfigurationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLoggingUnifiedAgentConfigurations,
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
			"group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"log_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"unified_agent_configuration_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"compartment_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"is_enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"configuration_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"LOGGING",
										}, true),
									},
									"defined_tags": {
										Type:             schema.TypeMap,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
										Elem:             schema.TypeString,
									},
									"description": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"configuration_state": {
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
									"time_last_modified": {
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

func readLoggingUnifiedAgentConfigurations(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingUnifiedAgentConfigurationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoggingManagementClient()

	return tfresource.ReadResource(sync)
}

type LoggingUnifiedAgentConfigurationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_logging.LoggingManagementClient
	Res    *oci_logging.ListUnifiedAgentConfigurationsResponse
}

func (s *LoggingUnifiedAgentConfigurationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoggingUnifiedAgentConfigurationsDataSourceCrud) Get() error {
	request := oci_logging.ListUnifiedAgentConfigurationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if groupId, ok := s.D.GetOkExists("group_id"); ok {
		tmp := groupId.(string)
		request.GroupId = &tmp
	}

	if isCompartmentIdInSubtree, ok := s.D.GetOkExists("is_compartment_id_in_subtree"); ok {
		tmp := isCompartmentIdInSubtree.(bool)
		request.IsCompartmentIdInSubtree = &tmp
	}

	if logId, ok := s.D.GetOkExists("log_id"); ok {
		tmp := logId.(string)
		request.LogId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_logging.ListUnifiedAgentConfigurationsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "logging")

	response, err := s.Client.ListUnifiedAgentConfigurations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListUnifiedAgentConfigurations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LoggingUnifiedAgentConfigurationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LoggingUnifiedAgentConfigurationsDataSource-", LoggingUnifiedAgentConfigurationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	unifiedAgentConfiguration := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, UnifiedAgentConfigurationSummaryToMap(item))
	}
	unifiedAgentConfiguration["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, LoggingUnifiedAgentConfigurationsDataSource().Schema["unified_agent_configuration_collection"].Elem.(*schema.Resource).Schema)
		unifiedAgentConfiguration["items"] = items
	}

	resources = append(resources, unifiedAgentConfiguration)
	if err := s.D.Set("unified_agent_configuration_collection", resources); err != nil {
		return err
	}

	return nil
}
