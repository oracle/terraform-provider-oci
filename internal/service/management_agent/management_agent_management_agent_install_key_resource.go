// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package management_agent

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
	oci_management_agent "github.com/oracle/oci-go-sdk/v56/managementagent"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

func ManagementAgentManagementAgentInstallKeyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createManagementAgentManagementAgentInstallKey,
		Read:     readManagementAgentManagementAgentInstallKey,
		Update:   updateManagementAgentManagementAgentInstallKey,
		Delete:   deleteManagementAgentManagementAgentInstallKey,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"allowed_key_install_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"time_expires": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: utils.TimeDiffSuppressFunction,
			},

			// Computed
			"created_by_principal_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"current_key_install_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
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
		},
	}
}

func createManagementAgentManagementAgentInstallKey(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentInstallKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.CreateResource(d, sync)
}

func readManagementAgentManagementAgentInstallKey(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentInstallKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.ReadResource(sync)
}

func updateManagementAgentManagementAgentInstallKey(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentInstallKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteManagementAgentManagementAgentInstallKey(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentInstallKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ManagementAgentManagementAgentInstallKeyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_management_agent.ManagementAgentClient
	Res                    *oci_management_agent.ManagementAgentInstallKey
	DisableNotFoundRetries bool
}

func (s *ManagementAgentManagementAgentInstallKeyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ManagementAgentManagementAgentInstallKeyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_management_agent.LifecycleStatesCreating),
	}
}

func (s *ManagementAgentManagementAgentInstallKeyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_management_agent.LifecycleStatesActive),
	}
}

func (s *ManagementAgentManagementAgentInstallKeyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_management_agent.LifecycleStatesDeleting),
	}
}

func (s *ManagementAgentManagementAgentInstallKeyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_management_agent.LifecycleStatesTerminated),
		string(oci_management_agent.LifecycleStatesDeleted),
	}
}

func (s *ManagementAgentManagementAgentInstallKeyResourceCrud) Create() error {
	request := oci_management_agent.CreateManagementAgentInstallKeyRequest{}

	if allowedKeyInstallCount, ok := s.D.GetOkExists("allowed_key_install_count"); ok {
		tmp := allowedKeyInstallCount.(int)
		request.AllowedKeyInstallCount = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if timeExpires, ok := s.D.GetOkExists("time_expires"); ok {
		tmp, err := time.Parse(time.RFC3339Nano, timeExpires.(string))
		if err != nil {
			return err
		}
		request.TimeExpires = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_agent")

	response, err := s.Client.CreateManagementAgentInstallKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagementAgentInstallKey
	return nil
}

func (s *ManagementAgentManagementAgentInstallKeyResourceCrud) Get() error {
	request := oci_management_agent.GetManagementAgentInstallKeyRequest{}

	tmp := s.D.Id()
	request.ManagementAgentInstallKeyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_agent")

	response, err := s.Client.GetManagementAgentInstallKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagementAgentInstallKey
	return nil
}

func (s *ManagementAgentManagementAgentInstallKeyResourceCrud) Update() error {
	request := oci_management_agent.UpdateManagementAgentInstallKeyRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if isKeyActive, ok := s.D.GetOkExists("is_key_active"); ok {
		tmp := isKeyActive.(bool)
		request.IsKeyActive = &tmp
	}

	tmp := s.D.Id()
	request.ManagementAgentInstallKeyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_agent")

	response, err := s.Client.UpdateManagementAgentInstallKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagementAgentInstallKey
	return nil
}

func (s *ManagementAgentManagementAgentInstallKeyResourceCrud) Delete() error {
	request := oci_management_agent.DeleteManagementAgentInstallKeyRequest{}

	tmp := s.D.Id()
	request.ManagementAgentInstallKeyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_agent")

	_, err := s.Client.DeleteManagementAgentInstallKey(context.Background(), request)
	return err
}

func (s *ManagementAgentManagementAgentInstallKeyResourceCrud) SetData() error {
	if s.Res.AllowedKeyInstallCount != nil {
		s.D.Set("allowed_key_install_count", *s.Res.AllowedKeyInstallCount)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedByPrincipalId != nil {
		s.D.Set("created_by_principal_id", *s.Res.CreatedByPrincipalId)
	}

	if s.Res.CurrentKeyInstallCount != nil {
		s.D.Set("current_key_install_count", *s.Res.CurrentKeyInstallCount)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeExpires != nil {
		s.D.Set("time_expires", s.Res.TimeExpires.Format(time.RFC3339Nano))
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
