// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatascienceContainersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatascienceContainers,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"container_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_latest": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tag_query_param": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_workload": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"usage_query_param": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"containers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"container_name": {
							Type:     schema.TypeString,
							Computed: true,
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
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"family_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"is_latest": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag_configuration_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"key": {
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
						"target_workloads": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"usages": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"workload_configuration_details_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"additional_configurations": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"cmd": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"health_check_port": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"server_port": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"use_case_configuration": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"additional_configurations": {
													Type:     schema.TypeMap,
													Computed: true,
													Elem:     schema.TypeString,
												},
												"use_case_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"workload_type": {
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

func readDatascienceContainers(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceContainersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceContainersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.ListContainersResponse
}

func (s *DatascienceContainersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceContainersDataSourceCrud) Get() error {
	request := oci_datascience.ListContainersRequest{}

	if containerName, ok := s.D.GetOkExists("container_name"); ok {
		tmp := containerName.(string)
		request.ContainerName = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if isLatest, ok := s.D.GetOkExists("is_latest"); ok {
		tmp := isLatest.(bool)
		request.IsLatest = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_datascience.ListContainersLifecycleStateEnum(state.(string))
	}

	if tagQueryParam, ok := s.D.GetOkExists("tag_query_param"); ok {
		tmp := tagQueryParam.(string)
		request.TagQueryParam = &tmp
	}

	if targetWorkload, ok := s.D.GetOkExists("target_workload"); ok {
		request.TargetWorkload = oci_datascience.ListContainersTargetWorkloadEnum(targetWorkload.(string))
	}

	if usageQueryParam, ok := s.D.GetOkExists("usage_query_param"); ok {
		request.UsageQueryParam = oci_datascience.ListContainersUsageQueryParamEnum(usageQueryParam.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.ListContainers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListContainers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatascienceContainersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatascienceContainersDataSource-", DatascienceContainersDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		container := map[string]interface{}{}

		if r.ContainerName != nil {
			container["container_name"] = *r.ContainerName
		}

		if r.DefinedTags != nil {
			container["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			container["description"] = *r.Description
		}

		if r.DisplayName != nil {
			container["display_name"] = *r.DisplayName
		}

		if r.FamilyName != nil {
			container["family_name"] = *r.FamilyName
		}

		container["freeform_tags"] = r.FreeformTags

		if r.IsLatest != nil {
			container["is_latest"] = *r.IsLatest
		}

		container["state"] = r.LifecycleState

		if r.Tag != nil {
			container["tag"] = *r.Tag
		}

		tagConfigurationList := []interface{}{}
		for _, item := range r.TagConfigurationList {
			tagConfigurationList = append(tagConfigurationList, TagConfigurationToMap(item))
		}
		container["tag_configuration_list"] = tagConfigurationList

		container["target_workloads"] = r.TargetWorkloads

		container["usages"] = r.Usages

		workloadConfigurationDetailsList := []interface{}{}
		for _, item := range r.WorkloadConfigurationDetailsList {
			workloadConfigurationDetailsList = append(workloadConfigurationDetailsList, WorkloadConfigurationDetailsToMap(item))
		}
		container["workload_configuration_details_list"] = workloadConfigurationDetailsList

		resources = append(resources, container)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatascienceContainersDataSource().Schema["containers"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("containers", resources); err != nil {
		return err
	}

	return nil
}

func JobRunUseCaseConfigurationDetailsToMap(obj *oci_datascience.JobRunUseCaseConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_datascience.GenericJobRunUseCaseConfigurationDetails:
		result["use_case_type"] = "GENERIC"

		result["additional_configurations"] = v.AdditionalConfigurations
	default:
		log.Printf("[WARN] Received 'use_case_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func TagConfigurationToMap(obj oci_datascience.TagConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func WorkloadConfigurationDetailsToMap(obj oci_datascience.WorkloadConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_datascience.JobRunWorkloadConfigurationDetails:
		result["workload_type"] = "JOB_RUN"

		if v.UseCaseConfiguration != nil {
			useCaseConfigurationArray := []interface{}{}
			if useCaseConfigurationMap := JobRunUseCaseConfigurationDetailsToMap(&v.UseCaseConfiguration); useCaseConfigurationMap != nil {
				useCaseConfigurationArray = append(useCaseConfigurationArray, useCaseConfigurationMap)
			}
			result["use_case_configuration"] = useCaseConfigurationArray
		}
	case oci_datascience.ModelDeployWorkloadConfigurationDetails:
		result["workload_type"] = "MODEL_DEPLOYMENT"

		result["additional_configurations"] = v.AdditionalConfigurations

		if v.Cmd != nil {
			result["cmd"] = string(*v.Cmd)
		}

		if v.HealthCheckPort != nil {
			result["health_check_port"] = int(*v.HealthCheckPort)
		}

		if v.ServerPort != nil {
			result["server_port"] = int(*v.ServerPort)
		}
	default:
		log.Printf("[WARN] Received 'workload_type' of unknown type %v", obj)
		return nil
	}

	return result
}
