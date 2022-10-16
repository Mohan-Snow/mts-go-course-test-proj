package main

import "testing"

func Test_shrinkString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{name: "englishStraight", args: args{input: "aaabbbbAAAAA"}, wantResult: "A5a3b4"},
		{name: "englishMixed", args: args{input: "cccaaabbbbaacc"}, wantResult: "a5b4c5"},
		{name: "russian", args: args{input: "аааАААААБББББАААААаа"}, wantResult: "А10Б5а5"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := shrinkString(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("shrinkString() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
