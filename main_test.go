package main

import "testing"

func TestMain_MaxSequenceLength(t *testing.T) {
	t.Run("MaxSubstringLength", func(t *testing.T){
		if MaxSequenceLength("aaabccc") != 3 {
			t.Error("expected 3")
		}
		if MaxSequenceLength("aaaabccc") != 4 {
			t.Error("expected 4")
		}
		if MaxSequenceLength("aaabaaacaaaadaaaaa") != 5 {
			t.Error("expected 5")
		}
		if MaxSequenceLength("aabaaacaaaaa") != 5 {
			t.Error("expected 5")
		}
	})
}