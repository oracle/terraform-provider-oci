// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"
)

func TestDbVersionDiffSuppress(t *testing.T) {
	oldValues := [9]string{"11.2.0.4.190416", "12.1.0.2.190416", "12.2.0.1.190416", "18.6.0.0.190416", "18.6.0.0.190416", "11.2.0.4", "11.2.0.4", "11.2.0.4", "18.6.0.0.190416"}
	newValues := [9]string{"11.2.0.4", "12.1.0.2", "12.2.0.1", "18.6.0.0", "18.0.0.0", "12.1.0.2", "12.2.0.4", "11.2.0.5", "19.6.0.0.190416"}
	assertResult := [9]bool{false, false, false, false, false, true, true, true, true}
	for i := 0; i < 9; i++ {
		old := oldValues[i]
		new := newValues[i]
		if dbVersionDiffSuppress("", old, new, nil) == assertResult[i] {
			if assertResult[i] == false {
				t.Errorf(fmt.Sprintf("Suppress expected, old : %s, new : %s", old, new))
			} else {
				t.Errorf(fmt.Sprintf("No Suppress expected, old : %s, new : %s", old, new))
			}
		}
	}
}
