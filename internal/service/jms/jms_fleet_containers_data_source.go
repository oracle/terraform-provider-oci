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

func JmsFleetContainersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsFleetContainers,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"application_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"jre_security_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"jre_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_started_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_started_less_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"container_collection": {
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
									"application_key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"application_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"container_key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"image_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"java_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"jre_key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"jre_security_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"managed_instance_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"node_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"pod_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_started": {
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

func readJmsFleetContainers(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetContainersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsFleetContainersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.ListContainersResponse
}

func (s *JmsFleetContainersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsFleetContainersDataSourceCrud) Get() error {
	request := oci_jms.ListContainersRequest{}

	if applicationName, ok := s.D.GetOkExists("application_name"); ok {
		tmp := applicationName.(string)
		request.ApplicationName = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if jreSecurityStatus, ok := s.D.GetOkExists("jre_security_status"); ok {
		request.JreSecurityStatus = oci_jms.ListContainersJreSecurityStatusEnum(jreSecurityStatus.(string))
	}

	if jreVersion, ok := s.D.GetOkExists("jre_version"); ok {
		tmp := jreVersion.(string)
		request.JreVersion = &tmp
	}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	if timeStartedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_started_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStartedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeStartedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeStartedLessThanOrEqualTo, ok := s.D.GetOkExists("time_started_less_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStartedLessThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeStartedLessThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

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

func (s *JmsFleetContainersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsFleetContainersDataSource-", JmsFleetContainersDataSource(), s.D))
	resources := []map[string]interface{}{}
	fleetContainer := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ContainerSummaryToMap(item))
	}
	fleetContainer["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsFleetContainersDataSource().Schema["container_collection"].Elem.(*schema.Resource).Schema)
		fleetContainer["items"] = items
	}

	resources = append(resources, fleetContainer)
	if err := s.D.Set("container_collection", resources); err != nil {
		return err
	}

	return nil
}

func ContainerSummaryToMap(obj oci_jms.ContainerSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApplicationKey != nil {
		result["application_key"] = string(*obj.ApplicationKey)
	}

	if obj.ApplicationName != nil {
		result["application_name"] = string(*obj.ApplicationName)
	}

	if obj.ContainerKey != nil {
		result["container_key"] = string(*obj.ContainerKey)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.ImageName != nil {
		result["image_name"] = string(*obj.ImageName)
	}

	if obj.JavaVersion != nil {
		result["java_version"] = string(*obj.JavaVersion)
	}

	if obj.JreKey != nil {
		result["jre_key"] = string(*obj.JreKey)
	}

	result["jre_security_status"] = string(obj.JreSecurityStatus)

	if obj.ManagedInstanceId != nil {
		result["managed_instance_id"] = string(*obj.ManagedInstanceId)
	}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	if obj.NodeName != nil {
		result["node_name"] = string(*obj.NodeName)
	}

	if obj.PodName != nil {
		result["pod_name"] = string(*obj.PodName)
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	return result
}
