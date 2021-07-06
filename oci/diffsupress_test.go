// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"
)

// issue-routing-tag: terraform/default
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

func TestAdDiffSuppress(t *testing.T) {
	oldValues := [8]string{"0.009999999776482582", "0.00999909999", "0.009999999776482582", "0.69999909", "0.699999988079071", "0.699999988079071", "0.7", "0.01"}
	newValues := [8]string{"0.01", "0.01", "0.001", "0.7", "0.0001", "0.8", "0.699999988079071", "0.009999999776482582"}
	assertResult := [8]bool{false, false, true, false, true, true, false, false}
	for i := 0; i < 8; i++ {
		old := oldValues[i]
		new := newValues[i]
		fmt.Println(i, " ", adDiffSuppress("", old, new, nil))
		if adDiffSuppress("", old, new, nil) == assertResult[i] {
			if assertResult[i] == false {
				t.Errorf(fmt.Sprintf("Suppress expected, old : %s, new : %s", old, new))
			} else {
				t.Errorf(fmt.Sprintf("No Suppress expected, old : %s, new : %s", old, new))
			}
		}
	}
}

// issue-routing-tag: terraform/default
func TestGiVersionDiffSuppress(t *testing.T) {
	oldValues := [7]string{"", "18.0.0.0", "19.0.0.0", "19.0.0.0", "18.0.2.0", "19.0.0.3", "18.0.0.0"}
	newValues := [7]string{"18.0.0.0", "", "18.2.0.0", "19.0.0.0", "18.0.0.0", "19.0.0.0", "18.0.3.0"}
	assertResult := [7]bool{true, true, true, false, false, false, false}
	for i := 0; i < 7; i++ {
		old := oldValues[i]
		new := newValues[i]
		if giVersionDiffSuppress("", old, new, nil) == assertResult[i] {
			if assertResult[i] == true {
				t.Errorf(fmt.Sprintf("Suppress expected, old : %s, new : %s", old, new))
			} else {
				t.Errorf(fmt.Sprintf("No Suppress expected, old : %s, new : %s", old, new))
			}
		}
	}
}
