package batch

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_batch "github.com/oracle/oci-go-sdk/v65/batch"
)

// Offline fixture only; no OCI calls are made by this test.
const smokeReservationID = "synthetic-reservation-id"

func reservationSmokeData(t *testing.T, compartment, subnet string) *schema.ResourceData {
	t.Helper()
	return schema.TestResourceDataRaw(t, BatchBatchContextResource().Schema, map[string]interface{}{
		"compartment_id": compartment,
		"fleets": []interface{}{map[string]interface{}{
			"name": "reservation-smoke-cpu", "type": "SERVICE_MANAGED_FLEET", "max_concurrent_tasks": 1,
			"shape":       []interface{}{map[string]interface{}{"shape_name": "VM.Standard.E5.Flex", "ocpus": 1, "memory_in_gbs": 12}},
			"reservation": []interface{}{map[string]interface{}{"type": "COMPUTE_CAPACITY_RESERVATION", "capacity_reservation_ids": []interface{}{smokeReservationID}}},
		}},
		"network": []interface{}{map[string]interface{}{"subnet_id": subnet}},
	})
}

func checkReservationSmokeDetails(s *BatchBatchContextResourceCrud, id, compartment, subnet string) error {
	if s.Res == nil || s.Res.Id == nil || *s.Res.Id != id {
		return fmt.Errorf("GET did not return the created BC %s", id)
	}
	if s.Res.LifecycleState != oci_batch.BatchContextLifecycleStateCreating && s.Res.LifecycleState != oci_batch.BatchContextLifecycleStateFailed {
		return fmt.Errorf("unexpected state %q for synthetic reservation", s.Res.LifecycleState)
	}
	if err := s.SetData(); err != nil {
		return err
	}
	want := map[string]interface{}{
		"compartment_id": compartment, "network.0.subnet_id": subnet,
		"fleets.#": 1, "fleets.0.name": "reservation-smoke-cpu", "fleets.0.type": "SERVICE_MANAGED_FLEET",
		"fleets.0.max_concurrent_tasks": 1, "fleets.0.shape.0.shape_name": "VM.Standard.E5.Flex",
		"fleets.0.shape.0.ocpus": 1, "fleets.0.shape.0.memory_in_gbs": 12,
		"fleets.0.reservation.#": 1, "fleets.0.reservation.0.type": "COMPUTE_CAPACITY_RESERVATION",
		"fleets.0.reservation.0.capacity_reservation_ids.#": 1,
		"fleets.0.reservation.0.capacity_reservation_ids.0": smokeReservationID,
	}
	for key, expected := range want {
		if actual := s.D.Get(key); !reflect.DeepEqual(actual, expected) {
			return fmt.Errorf("%s = %#v; want %#v", key, actual, expected)
		}
	}
	return nil
}

// Offline checks exercise request expansion, SDK response decoding and state
// flattening, including a missing-reservation negative control.
func TestBatchBatchContextReservationSmokeMapping(t *testing.T) {
	s := &BatchBatchContextResourceCrud{}
	s.D = reservationSmokeData(t, "compartment", "subnet")
	fleet, err := s.mapToCreateFleetDetails("fleets.0.%s")
	if err != nil {
		t.Fatal(err)
	}
	encoded, err := json.Marshal(fleet)
	if err != nil {
		t.Fatal(err)
	}
	var payload map[string]interface{}
	if err := json.Unmarshal(encoded, &payload); err != nil {
		t.Fatal(err)
	}
	reservation, ok := payload["reservation"].(map[string]interface{})
	if !ok || reservation["type"] != "COMPUTE_CAPACITY_RESERVATION" || !reflect.DeepEqual(reservation["capacityReservationIds"], []interface{}{smokeReservationID}) {
		t.Fatalf("request reservation lost: %s", encoded)
	}
	for _, state := range []string{"CREATING", "FAILED"} {
		for _, missing := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/missing=%t", state, missing), func(t *testing.T) {
				var decodedFleet interface{}
				if err := json.Unmarshal(encoded, &decodedFleet); err != nil {
					t.Fatal(err)
				}
				if missing {
					delete(decodedFleet.(map[string]interface{}), "reservation")
				}
				body, err := json.Marshal(map[string]interface{}{
					"id": "bc", "compartmentId": "compartment", "lifecycleState": state,
					"network": map[string]interface{}{"subnetId": "subnet"}, "fleets": []interface{}{decodedFleet},
				})
				if err != nil {
					t.Fatal(err)
				}
				var bc oci_batch.BatchContext
				if err := json.Unmarshal(body, &bc); err != nil {
					t.Fatal(err)
				}
				read := &BatchBatchContextResourceCrud{Res: &bc}
				read.D = schema.TestResourceDataRaw(t, BatchBatchContextResource().Schema, nil)
				err = checkReservationSmokeDetails(read, "bc", "compartment", "subnet")
				if missing && err == nil {
					t.Fatal("missing reservation incorrectly passed")
				}
				if !missing && err != nil {
					t.Fatal(err)
				}
			})
		}
	}
}
