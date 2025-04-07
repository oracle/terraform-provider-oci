package kms

import (
	"context"
	"fmt"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_kms "github.com/oracle/oci-go-sdk/v65/keymanagement"
)

func KmsVaultReplicationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createKmsVaultReplica,
		Read:     readKmsVaultReplica,
		Update:   updateKmsVaultReplica,
		Delete:   deleteKmsVaultReplica,
		Schema: map[string]*schema.Schema{
			// Required
			"vault_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"replica_region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"replica_vault_metadata": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"idcs_account_name_url": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"vault_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"private_endpoint_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},

			// Computed
		},
	}
}

func createKmsVaultReplica(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultReplicaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KmsVaultClient()

	return tfresource.CreateResource(d, sync)
}

func readKmsVaultReplica(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultReplicaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KmsVaultClient()

	return tfresource.ReadResource(sync)
}

func updateKmsVaultReplica(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultReplicaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KmsVaultClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteKmsVaultReplica(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultReplicaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KmsVaultClient()

	return tfresource.DeleteResource(d, sync)
}

type KmsVaultReplicaResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_kms.KmsVaultClient
	Res                    *oci_kms.ListVaultReplicasResponse
	DisableNotFoundRetries bool
}

func (s *KmsVaultReplicaResourceCrud) ID() string {
	return *s.Res.OpcRequestId
}

func (s *KmsVaultReplicaResourceCrud) Create() error {
	replicaRegionStr := ""
	if replicaRegion, ok := s.D.GetOkExists("replica_region"); ok {
		tmp := replicaRegion.(string)
		replicaRegionStr = tmp
	}

	vaultIdStr := ""
	if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
		tmp := vaultId.(string)
		vaultIdStr = tmp
	}

	return s.createVaultReplicaHelper(vaultIdStr, replicaRegionStr)
}

func (s *KmsVaultReplicaResourceCrud) Get() error {
	request := oci_kms.ListVaultReplicasRequest{}

	if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
		tmp := vaultId.(string)
		request.VaultId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.ListVaultReplicas(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *KmsVaultReplicaResourceCrud) Update() error {

	// Update is only supported for the change in replica region. All others are a forceNew
	if s.D.HasChange("replica_region") {

		oldRaw, newRaw := s.D.GetChange("replica_region")
		oldReplicaRegionName := oldRaw.(string)
		newReplicaRegionName := newRaw.(string)

		vaultIdStr := ""
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			vaultIdStr = tmp
		}

		// delete replica in the old region for the primary vault
		err := s.deleteVaultReplicaHelper(vaultIdStr, oldReplicaRegionName)
		if err != nil {
			return err
		}

		// Create replica in the new region for the primary vault after deletion is completed
		return s.createVaultReplicaHelper(vaultIdStr, newReplicaRegionName)
	}
	return nil
}

func (s *KmsVaultReplicaResourceCrud) Delete() error {
	replicaRegionStr := ""
	if replicaRegion, ok := s.D.GetOkExists("replica_region"); ok {
		tmp := replicaRegion.(string)
		replicaRegionStr = tmp
	}

	vaultIdStr := ""
	if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
		tmp := vaultId.(string)
		vaultIdStr = tmp
	}

	return s.deleteVaultReplicaHelper(vaultIdStr, replicaRegionStr)
}

func (s *KmsVaultReplicaResourceCrud) createVaultReplicaHelper(vaultId string, replicaRegion string) error {
	request := oci_kms.CreateVaultReplicaRequest{}

	if len(strings.TrimSpace(vaultId)) != 0 {
		request.VaultId = &vaultId
	}

	if len(strings.TrimSpace(replicaRegion)) != 0 {
		request.ReplicaRegion = &replicaRegion
	}

	if replicaVaultMetadata, ok := s.D.GetOkExists("replica_vault_metadata"); ok {
		if tmpList := replicaVaultMetadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "replica_vault_metadata", 0)
			tmp, err := s.mapToReplicaVaultMetadata(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ReplicaVaultMetadata = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	_, err := s.Client.CreateVaultReplica(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.Items[0].Status == oci_kms.VaultReplicaSummaryStatusCreated }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutCreate))
}

func (s *KmsVaultReplicaResourceCrud) deleteVaultReplicaHelper(vaultId string, replicaRegion string) error {
	request := oci_kms.DeleteVaultReplicaRequest{}

	if len(strings.TrimSpace(vaultId)) != 0 {
		request.VaultId = &vaultId
	}

	if len(strings.TrimSpace(replicaRegion)) != 0 {
		request.ReplicaRegion = &replicaRegion
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	_, err := s.Client.DeleteVaultReplica(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool {
		return (len(s.Res.Items) == 0 || s.Res.Items[0].Status == oci_kms.VaultReplicaSummaryStatusDeleted)
	}
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutDelete))
}

func (s *KmsVaultReplicaResourceCrud) SetData() error {
	return nil
}

// Necessary to have. Otherwise cause "Could not set resource state" error from setState() in crud helper
func (s *KmsVaultReplicaResourceCrud) State() oci_kms.VaultReplicaSummaryStatusEnum {
	if len(s.Res.Items) > 0 {
		return s.Res.Items[0].Status
	}
	return ""
}

func (s *KmsVaultReplicaResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_kms.VaultReplicaSummaryStatusCreating),
	}
}

func (s *KmsVaultReplicaResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_kms.VaultReplicaSummaryStatusCreated),
	}
}

func (s *KmsVaultReplicaResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_kms.VaultReplicaSummaryStatusDeleting),
	}
}

func (s *KmsVaultReplicaResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_kms.VaultReplicaSummaryStatusDeleted),
	}
}

func (s *KmsVaultReplicaResourceCrud) mapToReplicaVaultMetadata(fieldKeyFormat string) (oci_kms.ReplicaExternalVaultMetadata, error) {
	result := oci_kms.ReplicaExternalVaultMetadata{}

	if privateEndpointId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_endpoint_id")); ok {
		tmp := privateEndpointId.(string)
		result.PrivateEndpointId = &tmp
	}

	if idcsAccountNameUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "idcs_account_name_url")); ok {
		tmp := idcsAccountNameUrl.(string)
		result.IdcsAccountNameUrl = &tmp
	}

	return result, nil
}
