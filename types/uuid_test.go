package types

import (
	"testing"
)

func TestValidUUIDS(t *testing.T) {
	t.Parallel()

	ids := [...]string{
		"07420a21-eb22-45ed-be2e-8dea2d016e56",
		"1f4113c4-cbed-49b8-ba28-29c75342191e",
		"d053167c-0e22-492d-bf83-ae03dc6ea875",
		"a4bf1531-6683-4daf-aa4f-2843602896fb",
		"a96084ca-e260-44de-b6b4-58e256977ddb",
		"23f35971-1c4b-4492-bdc2-bc6cd74e59a7",
		"9f56ab52-9646-4d15-b416-304c849cb52e",
		"7055257d-ab49-4459-acc8-ae048b6e9a9c",
		"824b522d-f62c-4313-94cf-63dcbb92db60",
		"22ce0eb8-24b3-47b4-b81e-200bb270aa86",
	}

	for _, v := range ids {
		e := UUID(v)
		if !e.IsValid() {
			t.Errorf("UUID is not valid: %s", v)
			return
		}
	}
}

func TestInvalidUUIDS(t *testing.T) {
	t.Parallel()

	ids := [...]string{
		"07420a21-eb22-45ed-be2e-8dea2d016e5_",
		"1f413c4-cbed-49b8-ba28-29c275342191e",
		"d053167c-0e22-42d-bf83-ae03dc6ea875",
		"a4bf1531-66183-4daf-aaf-2843602896fb",
		"a960824ca-e260-44de-b64-58e256977ddb",
		"23f35971-1c42b-44292-bdc2-bc6cd759a7",
		"9f56ab52-9646_4d15-b416-304c849cb52e",
		"705557d-ab49-4459-acc8-ae08b6e9a9c",
		"824522d-f62c-4313-94f-63dcbb91db60",
		"22ce0e4b8-244b3-473b4-b81e-200bb2702aa86",
		"23<35971-1c4b-4492-bdc2-bc6cd74e59a7",
		"9f56ab52-9646-4dK5-b416-304-849cb52e",
		"7055257d-ab49-44.9-acc8-ae048b6e9a9c",
		"824b5$2d-f62c-4313-94cf-63dcbb92db60",
		"22ce0eb8-24b3-47b4-b81e-20=bb270aa86",
	}

	for _, v := range ids {
		e := UUID(v)
		if e.IsValid() {
			t.Errorf("UUID should not be valid: %s", v)
			return
		}
	}
}
