// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatascienceMlApplicationImplementationVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatascienceMlApplicationImplementationVersions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"ml_application_implementation_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ml_application_implementation_version_collection": {
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
									"allowed_migration_destinations": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"application_components": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"application_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"component_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"job_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"model_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"pipeline_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"resource_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"configuration_schema": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"default_value": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"description": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"is_mandatory": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"key_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"sample_value": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"validation_regexp": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"value_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"description": {
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
									"ml_application_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ml_application_implementation_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ml_application_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ml_application_package_arguments": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"arguments": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"description": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"is_mandatory": {
																Type:     schema.TypeBool,
																Computed: true,
															},
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
									"package_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"time_created": {
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

func readDatascienceMlApplicationImplementationVersions(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceMlApplicationImplementationVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceMlApplicationImplementationVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.ListMlApplicationImplementationVersionsResponse
}

func (s *DatascienceMlApplicationImplementationVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceMlApplicationImplementationVersionsDataSourceCrud) Get() error {
	request := oci_datascience.ListMlApplicationImplementationVersionsRequest{}

	if mlApplicationImplementationId, ok := s.D.GetOkExists("ml_application_implementation_id"); ok {
		tmp := mlApplicationImplementationId.(string)
		request.MlApplicationImplementationId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_datascience.MlApplicationImplementationVersionLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.ListMlApplicationImplementationVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMlApplicationImplementationVersions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatascienceMlApplicationImplementationVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatascienceMlApplicationImplementationVersionsDataSource-", DatascienceMlApplicationImplementationVersionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	mlApplicationImplementationVersion := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MlApplicationImplementationVersionSummaryToMap(item))
	}
	mlApplicationImplementationVersion["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatascienceMlApplicationImplementationVersionsDataSource().Schema["ml_application_implementation_version_collection"].Elem.(*schema.Resource).Schema)
		mlApplicationImplementationVersion["items"] = items
	}

	resources = append(resources, mlApplicationImplementationVersion)
	if err := s.D.Set("ml_application_implementation_version_collection", resources); err != nil {
		return err
	}

	return nil
}

func MlApplicationImplementationVersionSummaryToMap(obj oci_datascience.MlApplicationImplementationVersionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["allowed_migration_destinations"] = obj.AllowedMigrationDestinations

	configurationSchema := []interface{}{}
	for _, item := range obj.ConfigurationSchema {
		configurationSchema = append(configurationSchema, ConfigurationPropertySchemaToMap(item))
	}
	result["configuration_schema"] = configurationSchema

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.MlApplicationId != nil {
		result["ml_application_id"] = string(*obj.MlApplicationId)
	}

	if obj.MlApplicationImplementationId != nil {
		result["ml_application_implementation_id"] = string(*obj.MlApplicationImplementationId)
	}

	if obj.MlApplicationName != nil {
		result["ml_application_name"] = string(*obj.MlApplicationName)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.PackageVersion != nil {
		result["package_version"] = string(*obj.PackageVersion)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}
