package batch

import (
	"encoding/json"
	"fmt"
	"reflect"
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

func TestFleetToMapPreservesReservations(t *testing.T) {
	for _, includeReservation := range []bool{false, true} {
		for _, fleetType := range []string{"SERVICE_MANAGED_FLEET", "SERVICE_MANAGED_GPU_FLEET"} {
			for _, response := range []bool{false, true} {
				t.Run(fmt.Sprintf("%s/response=%t/reservation=%t", fleetType, response, includeReservation), func(t *testing.T) {
					var fleet oci_batch.Fleet
					switch {
					case fleetType == "SERVICE_MANAGED_FLEET" && response:
						fleet = &oci_batch.ServiceManagedFleet{}
					case fleetType == "SERVICE_MANAGED_FLEET":
						fleet = &oci_batch.CreateServiceManagedFleetDetails{}
					case response:
						fleet = &oci_batch.ServiceManagedGpuFleet{}
					default:
						fleet = &oci_batch.CreateServiceManagedGpuFleetDetails{}
					}
					payload := map[string]interface{}{
						"name": "test-fleet", "maxConcurrentTasks": 10,
						"state": "ACTIVE", "details": "ready",
					}
					if includeReservation {
						payload["reservation"] = map[string]interface{}{
							"type":                   "COMPUTE_CAPACITY_RESERVATION",
							"capacityReservationIds": []string{"ocid1.capacityreservation.test"},
						}
					}
					data, err := json.Marshal(payload)
					if err != nil {
						t.Fatal(err)
					}
					if err := json.Unmarshal(data, fleet); err != nil {
						t.Fatal(err)
					}
					// SDK decoding supplies value models to FleetToMap.
					fleet = reflect.ValueOf(fleet).Elem().Interface().(oci_batch.Fleet)
					got := FleetToMap(fleet)
					if got["type"] != fleetType || got["name"] != "test-fleet" || got["max_concurrent_tasks"] != 10 {
						t.Fatalf("fleet fields lost: %#v", got)
					}
					if response && (got["state"] != "ACTIVE" || got["details"] != "ready") {
						t.Fatalf("response state/details lost: %#v", got)
					}
					if !includeReservation {
						if _, exists := got["reservation"]; exists {
							t.Fatalf("unexpected reservation: %#v", got)
						}
						return
					}
					want := []interface{}{map[string]interface{}{
						"type":                     "COMPUTE_CAPACITY_RESERVATION",
						"capacity_reservation_ids": []string{"ocid1.capacityreservation.test"},
					}}
					if !reflect.DeepEqual(got["reservation"], want) {
						t.Fatalf("reservation = %#v; want %#v", got["reservation"], want)
					}
				})
			}
		}
	}
}
