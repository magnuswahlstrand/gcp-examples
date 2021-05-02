terraform {
  required_providers {}
}

locals {
  namespace = "mbw"
  project_name = "explorer"
  env = "prod"

  owner = "Magnus"
  location = "us-west1"

  required_prefix = "${local.namespace}-${local.project_name}-${local.env}"

  gcp_project = "${local.namespace}-${local.project_name}-${local.env}"
}

provider "google" {
  project = local.gcp_project
}

resource "google_project" "project" {
  name = local.gcp_project
  project_id = local.gcp_project

  billing_account = var.gcp_billing_account
}

// --- Enable APIs ---

resource "google_project_service" "run" {
  service = "run.googleapis.com"
  disable_on_destroy = false
  disable_dependent_services = true
}

resource "google_project_service" "containerregistry" {
  service = "containerregistry.googleapis.com"
  disable_on_destroy = false
  disable_dependent_services = true
}

resource "google_container_registry" "registry" {
  location = "EU"
  project = google_project.project.name
}

// --- Service Accounts ---

resource "google_service_account" "login_service_account" {
  account_id = "svc-login"
  display_name = "Login Service"
}

resource "google_service_account" "registration_service_account" {
  account_id = "svc-registration"
  display_name = "Registration Service"
}

// --- Services ---

resource "google_cloud_run_service" "login" {
  name = "${local.required_prefix}-run-login"
  location = local.location

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
      service_account_name = google_service_account.login_service_account.email
    }
  }

  traffic {
    percent = 100
    latest_revision = true
  }
}

resource "google_cloud_run_service" "registration" {
  name = "${local.required_prefix}-run-registration"
  location = local.location

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
      service_account_name = google_service_account.registration_service_account.email
    }
  }

  traffic {
    percent = 100
    latest_revision = true
  }
}