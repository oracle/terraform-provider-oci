// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_jms "github.com/oracle/oci-go-sdk/v58/jms"
)

func JmsListJreUsageDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularJmsListJreUsage,
		Schema: map[string]*schema.Schema{
			"application_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"application_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_end": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_start": {
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
						"approximate_application_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"approximate_installation_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"approximate_managed_instance_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"distribution": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"end_of_support_life_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fleet_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"managed_instance_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"operating_systems": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"architecture": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"family": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"version": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"release_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"security_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_end": {
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
						"time_start": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vendor": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularJmsListJreUsage(d *schema.ResourceData, m interface{}) error {
	sync := &JmsListJreUsageDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsListJreUsageDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.ListJreUsageResponse
}

func (s *JmsListJreUsageDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsListJreUsageDataSourceCrud) Get() error {
	request := oci_jms.ListJreUsageRequest{}

	if applicationId, ok := s.D.GetOkExists("application_id"); ok {
		tmp := applicationId.(string)
		request.ApplicationId = &tmp
	}

	if applicationName, ok := s.D.GetOkExists("application_name"); ok {
		tmp := applicationName.(string)
		request.ApplicationName = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if hostId, ok := s.D.GetOkExists("host_id"); ok {
		tmp := hostId.(string)
		request.HostId = &tmp
	}

	if timeEnd, ok := s.D.GetOkExists("time_end"); ok {
		tmp, err := time.Parse(time.RFC3339, timeEnd.(string))
		if err != nil {
			return err
		}
		request.TimeEnd = &oci_common.SDKTime{Time: tmp}
	}

	if timeStart, ok := s.D.GetOkExists("time_start"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStart.(string))
		if err != nil {
			return err
		}
		request.TimeStart = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.ListJreUsage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsListJreUsageDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsListJreUsageDataSource-", JmsListJreUsageDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, JreUsageToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func JreUsageToMap(obj oci_jms.JreUsage) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApproximateApplicationCount != nil {
		result["approximate_application_count"] = int(*obj.ApproximateApplicationCount)
	}

	if obj.ApproximateInstallationCount != nil {
		result["approximate_installation_count"] = int(*obj.ApproximateInstallationCount)
	}

	if obj.ApproximateManagedInstanceCount != nil {
		result["approximate_managed_instance_count"] = int(*obj.ApproximateManagedInstanceCount)
	}

	if obj.Distribution != nil {
		result["distribution"] = string(*obj.Distribution)
	}

	if obj.EndOfSupportLifeDate != nil {
		result["end_of_support_life_date"] = obj.EndOfSupportLifeDate.String()
	}

	if obj.FleetId != nil {
		result["fleet_id"] = string(*obj.FleetId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.ManagedInstanceId != nil {
		result["managed_instance_id"] = string(*obj.ManagedInstanceId)
	}

	operatingSystems := []interface{}{}
	for _, item := range obj.OperatingSystems {
		operatingSystems = append(operatingSystems, JmsOperatingSystemToMap(item))
	}
	result["operating_systems"] = operatingSystems

	if obj.ReleaseDate != nil {
		result["release_date"] = obj.ReleaseDate.String()
	}

	result["security_status"] = string(obj.SecurityStatus)

	if obj.TimeEnd != nil {
		result["time_end"] = obj.TimeEnd.String()
	}

	if obj.TimeFirstSeen != nil {
		result["time_first_seen"] = obj.TimeFirstSeen.String()
	}

	if obj.TimeLastSeen != nil {
		result["time_last_seen"] = obj.TimeLastSeen.String()
	}

	if obj.TimeStart != nil {
		result["time_start"] = obj.TimeStart.String()
	}

	if obj.Vendor != nil {
		result["vendor"] = string(*obj.Vendor)
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}

func JmsOperatingSystemToMap(obj oci_jms.OperatingSystem) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Architecture != nil {
		result["architecture"] = string(*obj.Architecture)
	}

	result["family"] = string(obj.Family)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}
