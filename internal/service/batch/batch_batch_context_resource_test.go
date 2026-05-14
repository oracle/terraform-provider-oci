package batch

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_batch "github.com/oracle/oci-go-sdk/v65/batch"
)

func TestBatchBatchContextDeletedTargetIncludesFailedOnlyDuringDelete(t *testing.T) {
	sync := &BatchBatchContextResourceCrud{}
	assertDoesNotContain(t, sync.DeletedTarget(), string(oci_batch.BatchContextLifecycleStateFailed))

	sync.DeleteTreatsFailedAsGone = true
	assertContains(t, sync.DeletedTarget(), string(oci_batch.BatchContextLifecycleStateFailed))
}

func TestBatchContextTerminalDeleteState(t *testing.T) {
	terminalStates := []oci_batch.BatchContextLifecycleStateEnum{
		oci_batch.BatchContextLifecycleStateDeleted,
		oci_batch.BatchContextLifecycleStateFailed,
	}
	for _, state := range terminalStates {
		if !batchContextTerminalDeleteState(state) {
			t.Fatalf("expected %s to be a terminal delete state", state)
		}
	}

	if batchContextTerminalDeleteState(oci_batch.BatchContextLifecycleStateActive) {
		t.Fatalf("expected ACTIVE not to be a terminal delete state")
	}
}

func TestBatchBatchContextStateMapsFailedOnlyDuringDelete(t *testing.T) {
	stateSchema := map[string]*schema.Schema{
		"state": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
	resourceData := schema.TestResourceDataRaw(t, stateSchema, map[string]interface{}{
		"state": string(oci_batch.BatchContextLifecycleStateFailed),
	})
	sync := &BatchBatchContextResourceCrud{}
	sync.D = resourceData

	if got := sync.State(); got != string(oci_batch.BatchContextLifecycleStateFailed) {
		t.Fatalf("expected FAILED outside delete flow, got %q", got)
	}

	sync.DeleteTreatsFailedAsGone = true
	if got := sync.State(); got != string(oci_batch.BatchContextLifecycleStateDeleted) {
		t.Fatalf("expected FAILED to map to DELETED during delete flow, got %q", got)
	}
}

func TestBatchBatchContextNormalizesFailedResourceStateOnlyDuringDelete(t *testing.T) {
	sync := &BatchBatchContextResourceCrud{
		Res: &oci_batch.BatchContext{
			LifecycleState: oci_batch.BatchContextLifecycleStateFailed,
		},
	}

	sync.normalizeFailedStateForDelete()
	if got := sync.Res.LifecycleState; got != oci_batch.BatchContextLifecycleStateFailed {
		t.Fatalf("expected FAILED outside delete flow, got %q", got)
	}

	sync.DeleteTreatsFailedAsGone = true
	sync.normalizeFailedStateForDelete()
	if got := sync.Res.LifecycleState; got != oci_batch.BatchContextLifecycleStateDeleted {
		t.Fatalf("expected FAILED to normalize to DELETED during delete flow, got %q", got)
	}
}

func assertContains(t *testing.T, values []string, expected string) {
	t.Helper()
	for _, value := range values {
		if value == expected {
			return
		}
	}
	t.Fatalf("expected %#v to contain %q", values, expected)
}

func assertDoesNotContain(t *testing.T, values []string, unexpected string) {
	t.Helper()
	for _, value := range values {
		if value == unexpected {
			t.Fatalf("expected %#v not to contain %q", values, unexpected)
		}
	}
}
