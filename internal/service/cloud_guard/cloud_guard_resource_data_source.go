// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

//resource not exposed to user through Terraform, but generated.
//Hence TF team suggested to keep the file commented as codeGen patch build fails if file not present

package cloud_guard

/*import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudGuardResourceDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCloudGuardResource,
		Schema: map[string]*schema.Schema{
			"resource_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"additional_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"os_info": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"open_ports_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"problem_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"risk_level": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"target_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_first_monitored": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_monitored": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vulnerability_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readSingularCloudGuardResource(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardResourceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

type CloudGuardResourceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.GetResourceResponse
}

func (s *CloudGuardResourceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardResourceDataSourceCrud) Get() error {
	request := oci_cloud_guard.GetResourceRequest{}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_guard")

	response, err := s.Client.GetResource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CloudGuardResourceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AdditionalDetails != nil {
		s.D.Set("additional_details", []interface{}{ResourceAdditionalDetailsToMap(s.Res.AdditionalDetails)})
	} else {
		s.D.Set("additional_details", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.OpenPortsCount != nil {
		s.D.Set("open_ports_count", *s.Res.OpenPortsCount)
	}

	if s.Res.ProblemCount != nil {
		s.D.Set("problem_count", *s.Res.ProblemCount)
	}

	if s.Res.Region != nil {
		s.D.Set("region", *s.Res.Region)
	}

	if s.Res.ResourceName != nil {
		s.D.Set("resource_name", *s.Res.ResourceName)
	}

	if s.Res.ResourceType != nil {
		s.D.Set("resource_type", *s.Res.ResourceType)
	}

	s.D.Set("risk_level", s.Res.RiskLevel)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.TargetName != nil {
		s.D.Set("target_name", *s.Res.TargetName)
	}

	if s.Res.TimeFirstMonitored != nil {
		s.D.Set("time_first_monitored", s.Res.TimeFirstMonitored.String())
	}

	if s.Res.TimeLastMonitored != nil {
		s.D.Set("time_last_monitored", s.Res.TimeLastMonitored.String())
	}

	if s.Res.VulnerabilityCount != nil {
		s.D.Set("vulnerability_count", *s.Res.VulnerabilityCount)
	}

	return nil
}*/
