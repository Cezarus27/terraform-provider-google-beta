// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccHealthcareFhirStore_healthcareFhirStoreBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckHealthcareFhirStoreDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcareFhirStore_healthcareFhirStoreBasicExample(context),
			},
			{
				ResourceName:            "google_healthcare_fhir_store.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"self_link", "dataset"},
			},
		},
	})
}

func testAccHealthcareFhirStore_healthcareFhirStoreBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_healthcare_fhir_store" "default" {
  name    = "tf-test-example-fhir-store%{random_suffix}"
  dataset = google_healthcare_dataset.dataset.id
  version = "R4"

  enable_update_create          = false
  disable_referential_integrity = false
  disable_resource_versioning   = false
  enable_history_import         = false

  notification_config {
    pubsub_topic = google_pubsub_topic.topic.id
  }

  labels = {
    label1 = "labelvalue1"
  }
}

resource "google_pubsub_topic" "topic" {
  name     = "tf-test-fhir-notifications%{random_suffix}"
}

resource "google_healthcare_dataset" "dataset" {
  name     = "tf-test-example-dataset%{random_suffix}"
  location = "us-central1"
}
`, context)
}

func TestAccHealthcareFhirStore_healthcareFhirStoreStreamingConfigExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"policyChanged": acctest.BootstrapPSARoles(t, "service-", "gcp-sa-healthcare", []string{"roles/bigquery.dataEditor", "roles/bigquery.jobUser"}),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckHealthcareFhirStoreDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcareFhirStore_healthcareFhirStoreStreamingConfigExample(context),
			},
			{
				ResourceName:            "google_healthcare_fhir_store.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"self_link", "dataset"},
			},
		},
	})
}

func testAccHealthcareFhirStore_healthcareFhirStoreStreamingConfigExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_healthcare_fhir_store" "default" {
  name    = "tf-test-example-fhir-store%{random_suffix}"
  dataset = google_healthcare_dataset.dataset.id
  version = "R4"

  enable_update_create          = false
  disable_referential_integrity = false
  disable_resource_versioning   = false
  enable_history_import         = false

  labels = {
    label1 = "labelvalue1"
  }

  stream_configs {
    resource_types = ["Observation"]
    bigquery_destination {
      dataset_uri = "bq://${google_bigquery_dataset.bq_dataset.project}.${google_bigquery_dataset.bq_dataset.dataset_id}"
      schema_config {
        recursive_structure_depth = 3
      }
    }
  }
}

resource "google_pubsub_topic" "topic" {
  name     = "tf-test-fhir-notifications%{random_suffix}"
}

resource "google_healthcare_dataset" "dataset" {
  name     = "tf-test-example-dataset%{random_suffix}"
  location = "us-central1"
}

resource "google_bigquery_dataset" "bq_dataset" {
  dataset_id    = "tf_test_bq_example_dataset%{random_suffix}"
  friendly_name = "test"
  description   = "This is a test description"
  location      = "US"
  delete_contents_on_destroy = true
}
`, context)
}

func TestAccHealthcareFhirStore_healthcareFhirStoreNotificationConfigExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckHealthcareFhirStoreDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcareFhirStore_healthcareFhirStoreNotificationConfigExample(context),
			},
			{
				ResourceName:            "google_healthcare_fhir_store.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"self_link", "dataset"},
			},
		},
	})
}

func testAccHealthcareFhirStore_healthcareFhirStoreNotificationConfigExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_healthcare_fhir_store" "default" {
  name    = "tf-test-example-fhir-store%{random_suffix}"
  dataset = google_healthcare_dataset.dataset.id
  version = "R4"

  enable_update_create          = false
  disable_referential_integrity = false
  disable_resource_versioning   = false
  enable_history_import         = false

  labels = {
    label1 = "labelvalue1"
  }

  notification_config {
    pubsub_topic = "${google_pubsub_topic.topic.id}"
  }
}

resource "google_pubsub_topic" "topic" {
  name = "tf-test-fhir-notifications%{random_suffix}"
}

resource "google_healthcare_dataset" "dataset" {
  name     = "tf-test-example-dataset%{random_suffix}"
  location = "us-central1"
}
`, context)
}

func TestAccHealthcareFhirStore_healthcareFhirStoreNotificationConfigsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckHealthcareFhirStoreDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcareFhirStore_healthcareFhirStoreNotificationConfigsExample(context),
			},
			{
				ResourceName:            "google_healthcare_fhir_store.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"self_link", "dataset"},
			},
		},
	})
}

func testAccHealthcareFhirStore_healthcareFhirStoreNotificationConfigsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_healthcare_fhir_store" "default" {
  provider = google-beta
  name     = "tf-test-example-fhir-store%{random_suffix}"
  dataset  = google_healthcare_dataset.dataset.id
  version  = "R4"

  enable_update_create          = false
  disable_referential_integrity = false
  disable_resource_versioning   = false
  enable_history_import         = false

  labels = {
    label1 = "labelvalue1"
  }

  notification_configs {
    pubsub_topic       = "${google_pubsub_topic.topic.id}"
    send_full_resource = true
  }
}

resource "google_pubsub_topic" "topic" {
  provider = google-beta
  name     = "tf-test-fhir-notifications%{random_suffix}"
}

resource "google_healthcare_dataset" "dataset" {
  provider = google-beta
  name     = "tf-test-example-dataset%{random_suffix}"
  location = "us-central1"
}
`, context)
}

func testAccCheckHealthcareFhirStoreDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_healthcare_fhir_store" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{HealthcareBasePath}}{{dataset}}/fhirStores/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("HealthcareFhirStore still exists at %s", url)
			}
		}

		return nil
	}
}
