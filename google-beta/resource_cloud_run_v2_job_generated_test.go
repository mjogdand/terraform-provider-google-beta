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

func TestAccCloudRunV2Job_cloudrunv2JobBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudRunV2JobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunV2Job_cloudrunv2JobBasicExample(context),
			},
			{
				ResourceName:            "google_cloud_run_v2_job.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccCloudRunV2Job_cloudrunv2JobBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_v2_job" "default" {
  name     = "tf-test-cloudrun-job%{random_suffix}"
  location = "us-central1"

  template {
    template {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }

  lifecycle {
    ignore_changes = [
      launch_stage,
    ]
  }
}
`, context)
}

func TestAccCloudRunV2Job_cloudrunv2JobSqlExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudRunV2JobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunV2Job_cloudrunv2JobSqlExample(context),
			},
			{
				ResourceName:            "google_cloud_run_v2_job.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccCloudRunV2Job_cloudrunv2JobSqlExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_v2_job" "default" {
  name     = "tf-test-cloudrun-job%{random_suffix}"
  location = "us-central1"
  
  template {
    template{
      volumes {
        name = "cloudsql"
        cloud_sql_instance {
          instances = [google_sql_database_instance.instance.connection_name]
        }
      }

      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"

        env {
          name = "FOO"
          value = "bar"
        }
        env {
          name = "latestdclsecret"
          value_source {
            secret_key_ref {
              secret = google_secret_manager_secret.secret.secret_id
              version = "1"
            }
          }
        }
        volume_mounts {
          name = "cloudsql"
          mount_path = "/cloudsql"
        }
      }
    }
  }

  lifecycle {
    ignore_changes = [
      launch_stage,
    ]
  }
}

data "google_project" "project" {
}

resource "google_secret_manager_secret" "secret" {
  secret_id = "secret%{random_suffix}"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "secret-version-data" {
  secret = google_secret_manager_secret.secret.name
  secret_data = "secret-data"
}

resource "google_secret_manager_secret_iam_member" "secret-access" {
  secret_id = google_secret_manager_secret.secret.id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${data.google_project.project.number}-compute@developer.gserviceaccount.com"
  depends_on = [google_secret_manager_secret.secret]
}

resource "google_sql_database_instance" "instance" {
  name             = "tf-test-cloudrun-sql%{random_suffix}"
  region           = "us-central1"
  database_version = "MYSQL_5_7"
  settings {
    tier = "db-f1-micro"
  }

  deletion_protection  = "%{deletion_protection}"
}
`, context)
}

func TestAccCloudRunV2Job_cloudrunv2JobVpcaccessExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudRunV2JobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunV2Job_cloudrunv2JobVpcaccessExample(context),
			},
			{
				ResourceName:            "google_cloud_run_v2_job.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccCloudRunV2Job_cloudrunv2JobVpcaccessExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_v2_job" "default" {
  name     = "tf-test-cloudrun-job%{random_suffix}"
  location = "us-central1"

  template {
    template{
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
      vpc_access{
        connector = google_vpc_access_connector.connector.id
        egress = "ALL_TRAFFIC"
      }
    }
  }

  lifecycle {
    ignore_changes = [
      launch_stage,
    ]
  }
}

resource "google_vpc_access_connector" "connector" {
  name          = "tf-test-run-vpc%{random_suffix}"
  subnet {
    name = google_compute_subnetwork.custom_test.name
  }
  machine_type = "e2-standard-4"
  min_instances = 2
  max_instances = 3
  region        = "us-central1"
}
resource "google_compute_subnetwork" "custom_test" {
  name          = "tf-test-run-subnetwork%{random_suffix}"
  ip_cidr_range = "10.2.0.0/28"
  region        = "us-central1"
  network       = google_compute_network.custom_test.id
}
resource "google_compute_network" "custom_test" {
  name                    = "tf-test-run-network%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccCloudRunV2Job_cloudrunv2JobSecretExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudRunV2JobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunV2Job_cloudrunv2JobSecretExample(context),
			},
			{
				ResourceName:            "google_cloud_run_v2_job.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccCloudRunV2Job_cloudrunv2JobSecretExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_v2_job" "default" {
  name     = "tf-test-cloudrun-job%{random_suffix}"
  location = "us-central1"

  template {
    template {
      volumes {
        name = "a-volume"
        secret {
          secret = google_secret_manager_secret.secret.secret_id
          default_mode = 292 # 0444
          items {
            version = "1"
            path = "my-secret"
            mode = 256 # 0400
          }
        }
      }
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
        volume_mounts {
          name = "a-volume"
          mount_path = "/secrets"
        }
      }
    }
  }

  lifecycle {
    ignore_changes = [
      launch_stage,
    ]
  }
}

data "google_project" "project" {
}

resource "google_secret_manager_secret" "secret" {
  secret_id = "secret%{random_suffix}"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "secret-version-data" {
  secret = google_secret_manager_secret.secret.name
  secret_data = "secret-data"
}

resource "google_secret_manager_secret_iam_member" "secret-access" {
  secret_id = google_secret_manager_secret.secret.id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${data.google_project.project.number}-compute@developer.gserviceaccount.com"
  depends_on = [google_secret_manager_secret.secret]
}
`, context)
}

func TestAccCloudRunV2Job_cloudrunv2JobEmptydirExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckCloudRunV2JobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunV2Job_cloudrunv2JobEmptydirExample(context),
			},
			{
				ResourceName:            "google_cloud_run_v2_job.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccCloudRunV2Job_cloudrunv2JobEmptydirExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_v2_job" "default" {
  provider = google-beta
  name     = "tf-test-cloudrun-job%{random_suffix}"
  location = "us-central1"
  launch_stage = "BETA"
  template {
    template {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
	volume_mounts {
	  name = "empty-dir-volume"
	  mount_path = "/mnt"
	}
      }
      volumes {
        name = "empty-dir-volume"
	empty_dir {
	  medium = "MEMORY"
	  size_limit = "128Mi"
	}
      }
    }
  }

  lifecycle {
    ignore_changes = [
      launch_stage,
    ]
  }
}
`, context)
}

func testAccCheckCloudRunV2JobDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_cloud_run_v2_job" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{CloudRunV2BasePath}}projects/{{project}}/locations/{{location}}/jobs/{{name}}")
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
				return fmt.Errorf("CloudRunV2Job still exists at %s", url)
			}
		}

		return nil
	}
}
