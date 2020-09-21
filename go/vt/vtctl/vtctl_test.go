package vtctl

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateShardRanges(t *testing.T) {
	type args struct {
		shards int
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			"errors for shards less than 0",
			args{0},
			nil,
			true,
		},
		{
			"errors for shards more than 65536",
			args{65537},
			nil,
			true,
		},
		{
			"works for a single shard",
			args{1},
			[]string{"-"},
			false,
		},
		{
			"works for more than one shard",
			args{2},
			[]string{"-80", "80-"},
			false,
		},
		{
			"works for an odd number of shards",
			args{7},
			[]string{"-24", "24-49", "49-6d", "6d-92", "92-b6", "b6-db", "db-"},
			false,
		},
		{
			"works for large number of shards",
			args{256},
			[]string{"-01", "01-02", "02-03", "03-04", "04-05", "05-06", "06-07", "07-08", "08-09", "09-0a", "0a-0b", "0b-0c", "0c-0d", "0d-0e", "0e-0f", "0f-10", "10-11", "11-12", "12-13", "13-14", "14-15", "15-16", "16-17", "17-18", "18-19", "19-1a", "1a-1b", "1b-1c", "1c-1d", "1d-1e", "1e-1f", "1f-20", "20-21", "21-22", "22-23", "23-24", "24-25", "25-26", "26-27", "27-28", "28-29", "29-2a", "2a-2b", "2b-2c", "2c-2d", "2d-2e", "2e-2f", "2f-30", "30-31", "31-32", "32-33", "33-34", "34-35", "35-36", "36-37", "37-38", "38-39", "39-3a", "3a-3b", "3b-3c", "3c-3d", "3d-3e", "3e-3f", "3f-40", "40-41", "41-42", "42-43", "43-44", "44-45", "45-46", "46-47", "47-48", "48-49", "49-4a", "4a-4b", "4b-4c", "4c-4d", "4d-4e", "4e-4f", "4f-50", "50-51", "51-52", "52-53", "53-54", "54-55", "55-56", "56-57", "57-58", "58-59", "59-5a", "5a-5b", "5b-5c", "5c-5d", "5d-5e", "5e-5f", "5f-60", "60-61", "61-62", "62-63", "63-64", "64-65", "65-66", "66-67", "67-68", "68-69", "69-6a", "6a-6b", "6b-6c", "6c-6d", "6d-6e", "6e-6f", "6f-70", "70-71", "71-72", "72-73", "73-74", "74-75", "75-76", "76-77", "77-78", "78-79", "79-7a", "7a-7b", "7b-7c", "7c-7d", "7d-7e", "7e-7f", "7f-80", "80-81", "81-82", "82-83", "83-84", "84-85", "85-86", "86-87", "87-88", "88-89", "89-8a", "8a-8b", "8b-8c", "8c-8d", "8d-8e", "8e-8f", "8f-90", "90-91", "91-92", "92-93", "93-94", "94-95", "95-96", "96-97", "97-98", "98-99", "99-9a", "9a-9b", "9b-9c", "9c-9d", "9d-9e", "9e-9f", "9f-a0", "a0-a1", "a1-a2", "a2-a3", "a3-a4", "a4-a5", "a5-a6", "a6-a7", "a7-a8", "a8-a9", "a9-aa", "aa-ab", "ab-ac", "ac-ad", "ad-ae", "ae-af", "af-b0", "b0-b1", "b1-b2", "b2-b3", "b3-b4", "b4-b5", "b5-b6", "b6-b7", "b7-b8", "b8-b9", "b9-ba", "ba-bb", "bb-bc", "bc-bd", "bd-be", "be-bf", "bf-c0", "c0-c1", "c1-c2", "c2-c3", "c3-c4", "c4-c5", "c5-c6", "c6-c7", "c7-c8", "c8-c9", "c9-ca", "ca-cb", "cb-cc", "cc-cd", "cd-ce", "ce-cf", "cf-d0", "d0-d1", "d1-d2", "d2-d3", "d3-d4", "d4-d5", "d5-d6", "d6-d7", "d7-d8", "d8-d9", "d9-da", "da-db", "db-dc", "dc-dd", "dd-de", "de-df", "df-e0", "e0-e1", "e1-e2", "e2-e3", "e3-e4", "e4-e5", "e5-e6", "e6-e7", "e7-e8", "e8-e9", "e9-ea", "ea-eb", "eb-ec", "ec-ed", "ed-ee", "ee-ef", "ef-f0", "f0-f1", "f1-f2", "f2-f3", "f3-f4", "f4-f5", "f5-f6", "f6-f7", "f7-f8", "f8-f9", "f9-fa", "fa-fb", "fb-fc", "fc-fd", "fd-fe", "fe-ff", "ff-"},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateShardRanges(tt.args.shards)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, got, tt.want)
		})
	}
}

func TestShardCalculatorForShardsGreaterThan512(t *testing.T) {
	got, err := generateShardRanges(512)
	if err != nil {
		t.Errorf("listShardRanges() error = %v", err)
	}

	want := "ff80-"

	if got[511] != want {
		t.Errorf("Invalid mapping for a 512 shard keyspace. Expected %v, want %v", got[511], want)
	}
}
