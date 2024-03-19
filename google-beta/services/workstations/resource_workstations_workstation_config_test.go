// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package workstations_test

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccWorkstationsWorkstationConfig_basic(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_basic(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccWorkstationsWorkstationConfig_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  provider   		      	 = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location   		         = "us-central1"

  labels = {
	foo = "bar"
  }
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location   		     = "us-central1"

  labels = {
	foo = "bar"
  }
}
`, context)
}

func TestAccWorkstationsWorkstationConfig_displayName(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"display_name":  "Display Name N",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_displayName(context, ""),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "labels", "terraform_labels"},
			},
			{
				Config: testAccWorkstationsWorkstationConfig_displayName(context, "2"),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccWorkstationsWorkstationConfig_displayName(context map[string]interface{}, update string) string {
	context["display_name"] = context["display_name"].(string) + update
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  provider   		      	 = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location   		         = "us-central1"

  labels = {
	foo = "bar"
  }
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location   		     = "us-central1"
  display_name           = "%{display_name} %{random_suffix}"

  labels = {
	foo = "bar"
  }
}
`, context)
}

func TestAccWorkstationsWorkstationConfig_persistentDirectories(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_persistentDirectories(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccWorkstationsWorkstationConfig_persistentDirectories(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  provider   		      	 = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location   		         = "us-central1"

  labels = {
	foo = "bar"
  }
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location   		     = "us-central1"

  persistent_directories {
	mount_path = "/home"
  }

  labels = {
	foo = "bar"
  }
}
`, context)
}

func TestAccWorkstationsWorkstationConfig_ephemeralDirectories(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_ephemeralDirectories(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccWorkstationsWorkstationConfig_ephemeralDirectories(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_compute_disk" "test_source_disk" {
  provider = google-beta
  name     = "tf-test-workstation-source-disk%{random_suffix}"
  size     = 10
  type     = "pd-ssd"
  zone     = "us-central1-a"
}

resource "google_compute_snapshot" "test_source_snapshot" {
  provider    = google-beta
  name        = "tf-test-workstation-source-snapshot%{random_suffix}"
  source_disk = google_compute_disk.test_source_disk.name
  zone        = "us-central1-a"
}

resource "google_workstations_workstation_cluster" "default" {
  provider   		      	 = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location   		         = "us-central1"

  labels = {
	foo = "bar"
  }
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location   		     = "us-central1"

  ephemeral_directories {
        mount_path = "/cache"
	gce_pd {
	      source_snapshot = google_compute_snapshot.test_source_snapshot.id
	      read_only  = true
    	}
  }

  labels = {
	foo = "bar"
  }
}
`, context)
}

func TestAccWorkstationsWorkstationConfig_ephemeralDirectories_withSourceImage(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_ephemeralDirectories_withSourceImage(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccWorkstationsWorkstationConfig_ephemeralDirectories_withSourceImage(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_compute_disk" "test_source_disk" {
  provider = google-beta
  name     = "tf-test-workstation-source-disk%{random_suffix}"
  size     = 10
  type     = "pd-ssd"
  zone     = "us-central1-a"
}

resource "google_compute_image" "test_source_image" {
  provider    = google-beta
  name        = "tf-test-workstation-source-image%{random_suffix}"
  source_disk = google_compute_disk.test_source_disk.name
  storage_locations = ["us-central1"]
}

resource "google_workstations_workstation_cluster" "default" {
  provider   		      	 = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location   		         = "us-central1"

  labels = {
	foo = "bar"
  }
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location   		     = "us-central1"

  ephemeral_directories {
        mount_path = "/cache"
	gce_pd {
 	      disk_type = "pd-standard"
	      source_image = google_compute_image.test_source_image.id
	      read_only  = true
    	}
  }

  labels = {
	foo = "bar"
  }
}
`, context)
}

func TestAccWorkstationsWorkstationConfig_serviceAccount(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_serviceAccount(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
		},
	})
}

func testAccWorkstationsWorkstationConfig_serviceAccount(context map[string]interface{}) string {
	return acctest.Nprintf(`
  resource "google_compute_network" "default" {
    provider                = google-beta
    name                    = "tf-test-workstation-cluster%{random_suffix}"
    auto_create_subnetworks = false
  }
  
  resource "google_compute_subnetwork" "default" {
    provider      = google-beta
    name          = "tf-test-workstation-cluster%{random_suffix}"
    ip_cidr_range = "10.0.0.0/24"
    region        = "us-central1"
    network       = google_compute_network.default.name
  }
  
  resource "google_workstations_workstation_cluster" "default" {
    provider                   = google-beta
    workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
    network                    = google_compute_network.default.id
    subnetwork                 = google_compute_subnetwork.default.id
    location                   = "us-central1"
  }
  
  resource "google_service_account" "default" {
    provider = google-beta
  
    account_id   = "tf-test-my-account%{random_suffix}"
    display_name = "Service Account"
  }
  
  resource "google_workstations_workstation_config" "default" {
    provider               = google-beta
    workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
    workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
    location               = "us-central1"

    enable_audit_agent     = true
  
    host {
      gce_instance {  
        service_account             = google_service_account.default.email
        service_account_scopes      = ["https://www.googleapis.com/auth/cloud-platform"]
      }
    }
  }
`, context)
}

func TestAccWorkstationsWorkstationConfig_boost(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_boost(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
		},
	})
}

func testAccWorkstationsWorkstationConfig_boost(context map[string]interface{}) string {
	return acctest.Nprintf(`
  resource "google_compute_network" "default" {
    provider                = google-beta
    name                    = "tf-test-workstation-cluster%{random_suffix}"
    auto_create_subnetworks = false
  }
  
  resource "google_compute_subnetwork" "default" {
    provider      = google-beta
    name          = "tf-test-workstation-cluster%{random_suffix}"
    ip_cidr_range = "10.0.0.0/24"
    region        = "us-central1"
    network       = google_compute_network.default.name
  }
  
  resource "google_workstations_workstation_cluster" "default" {
    provider                   = google-beta
    workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
    network                    = google_compute_network.default.id
    subnetwork                 = google_compute_subnetwork.default.id
    location                   = "us-central1"
  }

  resource "google_workstations_workstation_config" "default" {
    provider               = google-beta
    workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
    workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
    location               = "us-central1"
    host {
      gce_instance {
        boost_configs {
          id           = "boost-1"
          machine_type = "n1-standard-2"
          accelerators {
            type  = "nvidia-tesla-t4"
            count = 1
          }
        }
        boost_configs {
          id           = "boost-1"
          machine_type = "e2-standard-2"
        }
      }
    }
  }
`, context)
}

func TestAccWorkstationsWorkstationConfig_disableTcpConnections(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_disableTcpConnections(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
		},
	})
}

func testAccWorkstationsWorkstationConfig_disableTcpConnections(context map[string]interface{}) string {
	return acctest.Nprintf(`
  resource "google_compute_network" "default" {
    provider                = google-beta
    name                    = "tf-test-workstation-cluster%{random_suffix}"
    auto_create_subnetworks = false
  }
  
  resource "google_compute_subnetwork" "default" {
    provider      = google-beta
    name          = "tf-test-workstation-cluster%{random_suffix}"
    ip_cidr_range = "10.0.0.0/24"
    region        = "us-central1"
    network       = google_compute_network.default.name
  }
  
  resource "google_workstations_workstation_cluster" "default" {
    provider                   = google-beta
    workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
    network                    = google_compute_network.default.id
    subnetwork                 = google_compute_subnetwork.default.id
    location                   = "us-central1"
  }
  
  resource "google_service_account" "default" {
    provider = google-beta
  
    account_id   = "tf-test-my-account%{random_suffix}"
    display_name = "Service Account"
  }
  
  resource "google_workstations_workstation_config" "default" {
    provider               = google-beta
    workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
    workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
    location               = "us-central1"

    disable_tcp_connections = true
  
    host {
      gce_instance {  
        service_account             = google_service_account.default.email
        service_account_scopes      = ["https://www.googleapis.com/auth/cloud-platform"]
      }
    }
  }
`, context)
}

func TestAccWorkstationsWorkstationConfig_readinessChecks(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_readinessChecks(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
		},
	})
}

func testAccWorkstationsWorkstationConfig_readinessChecks(context map[string]interface{}) string {
	return acctest.Nprintf(`
  resource "google_compute_network" "default" {
    provider                = google-beta
    name                    = "tf-test-workstation-cluster%{random_suffix}"
    auto_create_subnetworks = false
  }
  
  resource "google_compute_subnetwork" "default" {
    provider      = google-beta
    name          = "tf-test-workstation-cluster%{random_suffix}"
    ip_cidr_range = "10.0.0.0/24"
    region        = "us-central1"
    network       = google_compute_network.default.name
  }
  
  resource "google_workstations_workstation_cluster" "default" {
    provider                   = google-beta
    workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
    network                    = google_compute_network.default.id
    subnetwork                 = google_compute_subnetwork.default.id
    location                   = "us-central1"
  }
  
  resource "google_service_account" "default" {
    provider = google-beta
  
    account_id   = "tf-test-my-account%{random_suffix}"
    display_name = "Service Account"
  }
  
  resource "google_workstations_workstation_config" "default" {
    provider               = google-beta
    workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
    workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
    location               = "us-central1"

    readiness_checks {
      path = "/"
      port = 80
    }
  
    host {
      gce_instance {  
        service_account             = google_service_account.default.email
        service_account_scopes      = ["https://www.googleapis.com/auth/cloud-platform"]
      }
    }
  }
`, context)
}

func TestAccWorkstationsWorkstationConfig_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_workstationConfigBasicExample(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "annotations", "labels", "terraform_labels"},
			},
			{
				Config: testAccWorkstationsWorkstationConfig_update(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "annotations", "labels", "terraform_labels"},
			},
			{
				Config: testAccWorkstationsWorkstationConfig_workstationConfigBasicExample(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "annotations", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccWorkstationsWorkstationConfig_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  provider   		      	 = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location   		         = "us-central1"

  labels = {
	foo = "bar"
  }
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location   		     = "us-central1"

  host {
    gce_instance {
      machine_type                 = "n1-standard-4"
      boot_disk_size_gb            = 35
      disable_public_ip_addresses  = true
      enable_nested_virtualization = true
    }
  }

  labels = {
	foo = "bar"
  }

  lifecycle {
    prevent_destroy = true
  }
}
`, context)
}

func TestAccWorkstationsWorkstationConfig_updateHostDetails(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_updateHostDetailsDefault(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
			{
				Config: testAccWorkstationsWorkstationConfig_updateHostDetailsUpdated(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
			{
				Config: testAccWorkstationsWorkstationConfig_updateHostDetailsUnsetInstanceConfigs(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
		},
	})
}

func testAccWorkstationsWorkstationConfig_updateHostDetailsDefault(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  provider                   = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location                   = "us-central1"
}

resource "google_service_account" "default" {
  provider = google-beta

  account_id   = "tf-test-my-account%{random_suffix}"
  display_name = "Service Account"
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location               = "us-central1"

  host {
    gce_instance {
      machine_type                = "e2-standard-2"
      boot_disk_size_gb           = 35
      pool_size                   = 0

      service_account             = google_service_account.default.email
      disable_public_ip_addresses = false

      shielded_instance_config {
        enable_secure_boot          = false
        enable_vtpm                 = false
        enable_integrity_monitoring = false
      }

      confidential_instance_config {
        enable_confidential_compute = false
      }
    }
  }
}
`, context)
}

func testAccWorkstationsWorkstationConfig_updateHostDetailsUpdated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  provider                   = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location                   = "us-central1"
}

# No longer explicitly used in google_workstations_workstation_config resource block below, but the
# service account needs to keep existing to allow the field to default from the API without error
resource "google_service_account" "default" {
  provider = google-beta

  account_id   = "tf-test-my-account%{random_suffix}"
  display_name = "Service Account"
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location               = "us-central1"

  host {
    gce_instance {
      machine_type                = "n2d-standard-2"
      boot_disk_size_gb           = 35
      pool_size                   = 1

      disable_public_ip_addresses = true
      tags = ["foo", "bar"]

      shielded_instance_config {
        enable_secure_boot          = true
        enable_vtpm                 = true
        enable_integrity_monitoring = true
      }

      confidential_instance_config {
        enable_confidential_compute = true
      }
 
      boost_configs {
        id           = "boost-1"
        machine_type = "n2d-standard-2"
      }
    }
  }
}
`, context)
}

func testAccWorkstationsWorkstationConfig_updateHostDetailsUnsetInstanceConfigs(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  provider                   = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location                   = "us-central1"
}

# No longer explicitly used in google_workstations_workstation_config resource block below, but the
# service account needs to keep existing to allow the field to default from the API without error
resource "google_service_account" "default" {
  provider = google-beta

  account_id   = "tf-test-my-account%{random_suffix}"
  display_name = "Service Account"
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location               = "us-central1"

  host {
    gce_instance {
      machine_type                = "n2d-standard-2"
      boot_disk_size_gb           = 35
      pool_size                   = 1

      disable_public_ip_addresses = true
      tags = ["foo", "bar"]

      shielded_instance_config {}
      confidential_instance_config {}
    }
  }
}
`, context)
}

func TestAccWorkstationsWorkstationConfig_updateWorkingDir(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_withCustomWorkingDir(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
			{
				Config: testAccWorkstationsWorkstationConfig_unsetWorkingDir(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
		},
	})
}

func testAccWorkstationsWorkstationConfig_withCustomWorkingDir(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  provider                   = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location                   = "us-central1"
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location               = "us-central1"

  container {
    image       = "us-central1-docker.pkg.dev/cloud-workstations-images/predefined/code-oss:latest"
    working_dir = "/test"
  }
}
`, context)
}

func testAccWorkstationsWorkstationConfig_unsetWorkingDir(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  provider                   = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location                   = "us-central1"
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location               = "us-central1"

  container {
    image       = "us-central1-docker.pkg.dev/cloud-workstations-images/predefined/code-oss:latest"
  }
}
`, context)
}

func TestAccWorkstationsWorkstationConfig_updatePersistentDirectorySourceSnapshot(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_withSourceDiskSnapshot(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
			{
				Config: testAccWorkstationsWorkstationConfig_withUpdatedSourceDiskSnapshot(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
		},
	})
}

func testAccWorkstationsWorkstationConfig_withSourceDiskSnapshot(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_compute_disk" "test_source_disk" {
  provider = google-beta
  name     = "tf-test-workstation-source-disk%{random_suffix}"
  size     = 10
  type     = "pd-ssd"
  zone     = "us-central1-a"
}

resource "google_compute_snapshot" "test_source_snapshot" {
  provider    = google-beta
  name        = "tf-test-workstation-source-snapshot%{random_suffix}"
  source_disk = google_compute_disk.test_source_disk.name
  zone        = "us-central1-a"
}

resource "google_workstations_workstation_cluster" "default" {
  provider                   = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location                   = "us-central1"
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location               = "us-central1"

  persistent_directories {
    mount_path = "/home"

    gce_pd {
      source_snapshot = google_compute_snapshot.test_source_snapshot.id
      reclaim_policy  = "DELETE"
    }
  }
}
`, context)
}

func testAccWorkstationsWorkstationConfig_withUpdatedSourceDiskSnapshot(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_compute_disk" "test_source_disk" {
  provider = google-beta
  name     = "tf-test-workstation-source-disk%{random_suffix}"
  size     = 10
  type     = "pd-ssd"
  zone     = "us-central1-a"
}

resource "google_compute_snapshot" "test_source_snapshot" {
  provider    = google-beta
  name        = "tf-test-workstation-source-snapshot%{random_suffix}"
  source_disk = google_compute_disk.test_source_disk.name
  zone        = "us-central1-a"
}

resource "google_compute_snapshot" "test_source_snapshot2" {
  provider    = google-beta
  name        = "tf-test-workstation-source-snapshot2%{random_suffix}"
  source_disk = google_compute_disk.test_source_disk.name
  zone        = "us-central1-a"
}

resource "google_workstations_workstation_cluster" "default" {
  provider                   = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location                   = "us-central1"
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location               = "us-central1"

  persistent_directories {
    mount_path = "/home"

    gce_pd {
      source_snapshot = google_compute_snapshot.test_source_snapshot2.id
      reclaim_policy  = "RETAIN"
    }
  }
}
`, context)
}
