package distributed_database

import "testing"

func TestValidateInsertActionValueAllowsExistingCompositeRaftShardSpaceWithoutCreateSizing(t *testing.T) {
	testCases := []struct {
		name       string
		autonomous bool
		value      string
	}{
		{
			name:       "autonomous",
			autonomous: true,
			value: `{
  "name": "US_1",
  "raftClusters": [{
    "name": "RW_US_1",
    "databases": [{
      "source": "ADBD_EXISTING_CLUSTER",
      "adminPassword": "test-password",
      "computeCount": 16,
      "dataStorageSizeInGbs": 32,
      "isAutoScalingEnabled": true,
      "cloudAutonomousVmClusterId": "ocid1.cloudautonomousvmcluster.oc1.phx.example",
      "vaultId": "ocid1.vault.oc1.eu-frankfurt-1.example",
      "kmsKeyId": "ocid1.key.oc1.eu-frankfurt-1.example"
    }]
  }]
}`,
		},
		{
			name:       "exascale",
			autonomous: false,
			value: `{
  "name": "asia",
  "raftClusters": [{
    "name": "east",
    "databases": [{
      "source": "XS_EXISTING_CLUSTER",
      "vmClusterId": "ocid1.exadbvmcluster.oc1.iad.example",
      "adminPasswordSecretId": "ocid1.vaultsecret.oc1.uk-london-1.example",
      "vaultId": "ocid1.vault.oc1.eu-frankfurt-1.example",
      "kmsKeyId": "ocid1.key.oc1.eu-frankfurt-1.example"
    }]
  }]
}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateInsertActionValue("compositeRaftShardSpaces", tc.value, tc.autonomous)
			if err != nil {
				t.Fatalf("existing Composite RAFT shard-space INSERT must not require create-only sizing fields: %v", err)
			}
		})
	}
}
