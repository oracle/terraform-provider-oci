// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	oci_kms "github.com/oracle/oci-go-sdk/keymanagement"
)

func VaultResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createVault,
		Read:     readVault,
		Update:   updateVault,
		Delete:   deleteVault,
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
			"vault_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"crypto_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_endpoint": {
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
			"time_of_deletion": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createVault(d *schema.ResourceData, m interface{}) error {
	sync := &VaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).kmsVaultClient

	return CreateResource(d, sync)
}

func readVault(d *schema.ResourceData, m interface{}) error {
	sync := &VaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).kmsVaultClient

	return ReadResource(sync)
}

func updateVault(d *schema.ResourceData, m interface{}) error {
	sync := &VaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).kmsVaultClient

	return UpdateResource(d, sync)
}

func deleteVault(d *schema.ResourceData, m interface{}) error {
	sync := &VaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).kmsVaultClient

	return DeleteResource(d, sync)
}

type VaultResourceCrud struct {
	BaseCrud
	Client                 *oci_kms.KmsVaultClient
	Res                    *oci_kms.Vault
	DisableNotFoundRetries bool
}

func (s *VaultResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *VaultResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_kms.VaultLifecycleStateCreating),
	}
}

func (s *VaultResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_kms.VaultLifecycleStateActive),
	}
}

func (s *VaultResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_kms.VaultLifecycleStateDeleting),
		string(oci_kms.VaultLifecycleStateSchedulingDeletion),
	}
}

func (s *VaultResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_kms.VaultLifecycleStateDeleted),
		string(oci_kms.VaultLifecycleStatePendingDeletion),
	}
}

func (s *VaultResourceCrud) Create() error {
	request := oci_kms.CreateVaultRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if vaultType, ok := s.D.GetOkExists("vault_type"); ok {
		request.VaultType = oci_kms.CreateVaultDetailsVaultTypeEnum(vaultType.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.CreateVault(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Vault
	return nil
}

func (s *VaultResourceCrud) Get() error {
	request := oci_kms.GetVaultRequest{}

	tmp := s.D.Id()
	request.VaultId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.GetVault(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Vault
	return nil
}

func (s *VaultResourceCrud) Update() error {
	request := oci_kms.UpdateVaultRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.VaultId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.UpdateVault(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Vault
	return nil
}

func (s *VaultResourceCrud) Delete() error {
	request := oci_kms.ScheduleVaultDeletionRequest{}

	tmp := s.D.Id()
	request.VaultId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "kms")

	_, err := s.Client.ScheduleVaultDeletion(context.Background(), request)
	return err
}

func (s *VaultResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CryptoEndpoint != nil {
		s.D.Set("crypto_endpoint", *s.Res.CryptoEndpoint)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ManagementEndpoint != nil {
		s.D.Set("management_endpoint", *s.Res.ManagementEndpoint)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfDeletion != nil {
		s.D.Set("time_of_deletion", *s.Res.TimeOfDeletion)
	}

	s.D.Set("vault_type", s.Res.VaultType)

	return nil
}
