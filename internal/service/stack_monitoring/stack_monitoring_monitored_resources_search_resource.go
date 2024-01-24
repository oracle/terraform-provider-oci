// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringMonitoredResourcesSearchResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createStackMonitoringMonitoredResourcesSearch,
		Read:     readStackMonitoringMonitoredResourcesSearch,
		Delete:   deleteStackMonitoringMonitoredResourcesSearch,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"exclude_fields": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"fields": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"host_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"host_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"license": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"management_agent_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name_contains": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"property_equals": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"resource_time_zone": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"time_created_greater_than_or_equal_to": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},
			"time_created_less_than": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},
			"time_updated_greater_than_or_equal_to": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},
			"time_updated_less_than": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
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
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
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
						"external_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"host_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"license": {
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
						"properties": {
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
									"value": {
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
						"system_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_updated": {
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
		},
	}
}

func createStackMonitoringMonitoredResourcesSearch(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoredResourcesSearchResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.CreateResource(d, sync)
}

func readStackMonitoringMonitoredResourcesSearch(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteStackMonitoringMonitoredResourcesSearch(d *schema.ResourceData, m interface{}) error {
	return nil
}

type StackMonitoringMonitoredResourcesSearchResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_stack_monitoring.StackMonitoringClient
	Res                    *oci_stack_monitoring.MonitoredResourceCollection
	DisableNotFoundRetries bool
}

func (s *StackMonitoringMonitoredResourcesSearchResourceCrud) ID() string {
	var id = "/monitoredResources/actions/associateMonitoredResources"

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		id = id + "/compartmentId" + compartmentId.(string)
	}
	if hostName, ok := s.D.GetOkExists("host_name"); ok {
		id = id + "/hostName" + hostName.(string)
	}
	if hostNameContains, ok := s.D.GetOkExists("host_name_contains"); ok {
		id = id + "/hostNameContains" + hostNameContains.(string)
	}
	if managementAgentId, ok := s.D.GetOkExists("management_agent_id"); ok {
		id = id + "/managementAgentId" + managementAgentId.(string)
	}
	if name, ok := s.D.GetOkExists("name"); ok {
		id = id + "/name" + name.(string)
	}
	if nameContains, ok := s.D.GetOkExists("name_contains"); ok {
		id = id + "/nameContains" + nameContains.(string)
	}
	if type_, ok := s.D.GetOkExists("type"); ok {
		id = id + "/type" + type_.(string)
	}

	return id
}

func (s *StackMonitoringMonitoredResourcesSearchResourceCrud) Create() error {
	request := oci_stack_monitoring.SearchMonitoredResourcesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if excludeFields, ok := s.D.GetOkExists("exclude_fields"); ok {
		interfaces := excludeFields.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("exclude_fields") {
			request.ExcludeFields = tmp
		}
	}

	if externalId, ok := s.D.GetOkExists("external_id"); ok {
		tmp := externalId.(string)
		request.ExternalId = &tmp
	}

	if fields, ok := s.D.GetOkExists("fields"); ok {
		interfaces := fields.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("fields") {
			request.Fields = tmp
		}
	}

	if hostName, ok := s.D.GetOkExists("host_name"); ok {
		tmp := hostName.(string)
		request.HostName = &tmp
	}

	if hostNameContains, ok := s.D.GetOkExists("host_name_contains"); ok {
		tmp := hostNameContains.(string)
		request.HostNameContains = &tmp
	}

	if license, ok := s.D.GetOkExists("license"); ok {
		request.License = oci_stack_monitoring.LicenseTypeEnum(license.(string))
	}

	if managementAgentId, ok := s.D.GetOkExists("management_agent_id"); ok {
		tmp := managementAgentId.(string)
		request.ManagementAgentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if nameContains, ok := s.D.GetOkExists("name_contains"); ok {
		tmp := nameContains.(string)
		request.NameContains = &tmp
	}

	if propertyEquals, ok := s.D.GetOkExists("property_equals"); ok {
		request.PropertyEquals = tfresource.ObjectMapToStringMap(propertyEquals.(map[string]interface{}))
	}

	if resourceTimeZone, ok := s.D.GetOkExists("resource_time_zone"); ok {
		tmp := resourceTimeZone.(string)
		request.ResourceTimeZone = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_stack_monitoring.ResourceLifecycleStateEnum(state.(string))
	}

	if timeCreatedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_created_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeCreatedLessThan, ok := s.D.GetOkExists("time_created_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	if timeUpdatedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_updated_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeUpdatedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeUpdatedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeUpdatedLessThan, ok := s.D.GetOkExists("time_updated_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeUpdatedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeUpdatedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		tmp := type_.(string)
		request.Type = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.SearchMonitoredResources(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MonitoredResourceCollection
	return nil
}

func (s *StackMonitoringMonitoredResourcesSearchResourceCrud) SetData() error {
	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MonitoredResourceSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func MonitoredResourceSummaryToMap(obj oci_stack_monitoring.MonitoredResourceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.ExternalId != nil {
		result["external_id"] = string(*obj.ExternalId)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["license"] = string(obj.License)

	if obj.ManagementAgentId != nil {
		result["management_agent_id"] = string(*obj.ManagementAgentId)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	properties := []interface{}{}
	for _, item := range obj.Properties {
		properties = append(properties, MonitoredResourcePropertyToMap(item))
	}
	result["properties"] = properties

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}
