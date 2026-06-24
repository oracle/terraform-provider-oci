package core

import (
	"testing"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func TestBlockVolumeReplicaInfoToMapSetsXrrKmsKeyIdFromKmsKeyId(t *testing.T) {
	kmsKeyId := "ocid1.key.oc1..exampleuniqueID"

	got := BlockVolumeReplicaInfoToMap(oci_core.BlockVolumeReplicaInfo{
		KmsKeyId: &kmsKeyId,
	})

	if got["kms_key_id"] != kmsKeyId {
		t.Fatalf("kms_key_id = %v, want %q", got["kms_key_id"], kmsKeyId)
	}
	if got["xrr_kms_key_id"] != kmsKeyId {
		t.Fatalf("xrr_kms_key_id = %v, want %q", got["xrr_kms_key_id"], kmsKeyId)
	}
}

func TestBlockVolumeReplicaInfoToMapDoesNotSetXrrKmsKeyIdWithoutKmsKeyId(t *testing.T) {
	got := BlockVolumeReplicaInfoToMap(oci_core.BlockVolumeReplicaInfo{})

	if _, ok := got["xrr_kms_key_id"]; ok {
		t.Fatalf("xrr_kms_key_id set without KmsKeyId: %v", got["xrr_kms_key_id"])
	}
}

func TestVolumeGroupReplicaInfoToMapSetsXrrKmsKeyIdFromKmsKeyId(t *testing.T) {
	kmsKeyId := "ocid1.key.oc1..exampleuniqueID"

	got := VolumeGroupReplicaInfoToMap(oci_core.VolumeGroupReplicaInfo{
		KmsKeyId: &kmsKeyId,
	})

	if got["xrr_kms_key_id"] != kmsKeyId {
		t.Fatalf("xrr_kms_key_id = %v, want %q", got["xrr_kms_key_id"], kmsKeyId)
	}
}

func TestVolumeGroupReplicaInfoToMapDoesNotSetXrrKmsKeyIdWithoutKmsKeyId(t *testing.T) {
	got := VolumeGroupReplicaInfoToMap(oci_core.VolumeGroupReplicaInfo{})

	if _, ok := got["xrr_kms_key_id"]; ok {
		t.Fatalf("xrr_kms_key_id set without KmsKeyId: %v", got["xrr_kms_key_id"])
	}
}
