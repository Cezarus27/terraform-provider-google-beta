// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccComputeResourcePolicy_resourcePolicyBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeResourcePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeResourcePolicy_resourcePolicyBasicExample(context),
			},
		},
	})
}

func testAccComputeResourcePolicy_resourcePolicyBasicExample(context map[string]interface{}) string {
	return Nprintf(`
provider "google-beta" {
  region = "us-central1"
  zone   = "us-central1-a"
}

resource "google_compute_resource_policy" "foo" {
  provider = "google-beta"
  name = "policy-%{random_suffix}"
  region = "us-central1"
  snapshot_schedule_policy {
    schedule {
      daily_schedule {
        days_in_cycle = 1
        start_time = "04:00"
      }
    }
  }
}
`, context)
}

func TestAccComputeResourcePolicy_resourcePolicyFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeResourcePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeResourcePolicy_resourcePolicyFullExample(context),
			},
		},
	})
}

func testAccComputeResourcePolicy_resourcePolicyFullExample(context map[string]interface{}) string {
	return Nprintf(`
provider "google-beta" {
  region = "us-central1"
  zone   = "us-central1-a"
}

resource "google_compute_resource_policy" "bar" {
  provider = "google-beta"
  name = "policy-%{random_suffix}"
  region = "us-central1"
  snapshot_schedule_policy {
    schedule {
      hourly_schedule {
        hours_in_cycle = 20
        start_time = "23:00"
      }
    }
    retention_policy {
      max_retention_days = 10
      on_source_disk_delete = "KEEP_AUTO_SNAPSHOTS"
    }
    snapshot_properties {
      labels = {
        my_label = "value"
      }
      storage_locations = ["us"]
      guest_flush = true
    }
  }
}
`, context)
}

func testAccCheckComputeResourcePolicyDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_resource_policy" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/resourcePolicies/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeResourcePolicy still exists at %s", url)
		}
	}

	return nil
}
