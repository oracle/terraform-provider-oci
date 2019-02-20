// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	oci_kms "github.com/oracle/oci-go-sdk/keymanagement"
)

func KmsVaultResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createKmsVault,
		Read:     readKmsVault,
		Update:   updateKmsVault,
		Delete:   deleteKmsVault,
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

func createKmsVault(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).kmsVaultClient

	return CreateResource(d, sync)
}

func readKmsVault(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).kmsVaultClient

	return ReadResource(sync)
}

func updateKmsVault(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).kmsVaultClient

	return UpdateResource(d, sync)
}

func deleteKmsVault(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).kmsVaultClient

	return DeleteResource(d, sync)
}

type KmsVaultResourceCrud struct {
	BaseCrud
	Client                 *oci_kms.KmsVaultClient
	Res                    *oci_kms.Vault
	DisableNotFoundRetries bool
}

func (s *KmsVaultResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *KmsVaultResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_kms.VaultLifecycleStateCreating),
	}
}

func (s *KmsVaultResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_kms.VaultLifecycleStateActive),
	}
}

func (s *KmsVaultResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_kms.VaultLifecycleStateDeleting),
		string(oci_kms.VaultLifecycleStateSchedulingDeletion),
	}
}

func (s *KmsVaultResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_kms.VaultLifecycleStateDeleted),
		string(oci_kms.VaultLifecycleStatePendingDeletion),
	}
}

func (s *KmsVaultResourceCrud) Create() error {
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

func (s *KmsVaultResourceCrud) Get() error {
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

func (s *KmsVaultResourceCrud) Update() error {
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

func (s *KmsVaultResourceCrud) Delete() error {
	request := oci_kms.ScheduleVaultDeletionRequest{}

	tmp := s.D.Id()
	request.VaultId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "kms")

	_, err := s.Client.ScheduleVaultDeletion(context.Background(), request)
	return err
}

func (s *KmsVaultResourceCrud) SetData() error {
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
