// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_cloud_guard "github.com/oracle/oci-go-sdk/v58/cloudguard"

	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CloudGuardCloudGuardConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCloudGuardCloudGuardConfiguration,
		Read:     readCloudGuardCloudGuardConfiguration,
		Update:   updateCloudGuardCloudGuardConfiguration,
		Delete:   deleteCloudGuardCloudGuardConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"reporting_region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"self_manage_resources": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
		},
	}
}

func createCloudGuardCloudGuardConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardCloudGuardConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.CreateResource(d, sync)
}

func readCloudGuardCloudGuardConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardCloudGuardConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

func updateCloudGuardCloudGuardConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardCloudGuardConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCloudGuardCloudGuardConfiguration(d *schema.ResourceData, m interface{}) error {
	return nil
}

type CloudGuardCloudGuardConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_cloud_guard.CloudGuardClient
	Res                    *oci_cloud_guard.Configuration
	DisableNotFoundRetries bool
}

func (s *CloudGuardCloudGuardConfigurationResourceCrud) ID() string {
	return getCloudGuardConfigurationCompositeId()
}

func (s *CloudGuardCloudGuardConfigurationResourceCrud) Create() error {
	request := oci_cloud_guard.UpdateConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if reportingRegion, ok := s.D.GetOkExists("reporting_region"); ok {
		tmp := reportingRegion.(string)
		request.ReportingRegion = &tmp
	}

	if selfManageResources, ok := s.D.GetOkExists("self_manage_resources"); ok {
		tmp := selfManageResources.(bool)
		request.SelfManageResources = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_cloud_guard.CloudGuardStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.UpdateConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Configuration
	return nil
}

func (s *CloudGuardCloudGuardConfigurationResourceCrud) Get() error {
	request := oci_cloud_guard.GetConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	err := parseCloudGuardConfigurationCompositeId(s.D.Id())
	if err == nil {
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.GetConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Configuration
	return nil
}

func (s *CloudGuardCloudGuardConfigurationResourceCrud) Update() error {
	request := oci_cloud_guard.UpdateConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if reportingRegion, ok := s.D.GetOkExists("reporting_region"); ok {
		tmp := reportingRegion.(string)
		request.ReportingRegion = &tmp
	}

	if selfManageResources, ok := s.D.GetOkExists("self_manage_resources"); ok {
		tmp := selfManageResources.(bool)
		request.SelfManageResources = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_cloud_guard.CloudGuardStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.UpdateConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Configuration
	return nil
}

func (s *CloudGuardCloudGuardConfigurationResourceCrud) SetData() error {

	err := parseCloudGuardConfigurationCompositeId(s.D.Id())
	if err == nil {
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.ReportingRegion != nil {
		s.D.Set("reporting_region", *s.Res.ReportingRegion)
	}

	if s.Res.SelfManageResources != nil {
		s.D.Set("self_manage_resources", *s.Res.SelfManageResources)
	}

	s.D.Set("status", s.Res.Status)

	return nil
}

func getCloudGuardConfigurationCompositeId() string {
	compositeId := "configuration"
	return compositeId
}

func parseCloudGuardConfigurationCompositeId(compositeId string) (err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("configuration", compositeId)
	if !match || len(parts) != 1 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}

	return
}
