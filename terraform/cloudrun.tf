# Service account for CatBot Cloud Run service
resource "google_service_account" "catbot_service_account" {
  project      = var.project_id
  account_id   = "catbot-cloudrun-sa"
  display_name = "CatBot Cloud Run Service Account"
  description  = "Service account for CatBot Slack bot Cloud Run service"
}

# Cloud Run Worker Pools for CatBot
resource "google_cloud_run_v2_worker_pool" "catbot_worker_pool" {
  name     = "catbot-run-wp-catbot"
  location = var.region
  project  = var.project_id
  deletion_protection = false
  launch_stage = "BETA"

  template {
    service_account = google_service_account.catbot_service_account.email
    
    containers {
      image = "${var.region}-docker.pkg.dev/${var.project_id}/catbot/slackbot:latest"
      
      env {
        name = "SLACK_BOT_TOKEN"
        value_source {
          secret_key_ref {
            secret  = google_secret_manager_secret.slack_bot_token.secret_id
            version = "1"
          }
        }
      }

      env {
        name = "SLACK_APP_TOKEN"
        value_source {
          secret_key_ref {
            secret  = google_secret_manager_secret.slack_app_token.secret_id
            version = "1"
          }
        }
      }

      resources {
        limits = {
          cpu    = "1000m"
          memory = "512Mi"
        }
      }
    }
  }

  instance_splits {
    type = "INSTANCE_SPLIT_ALLOCATION_TYPE_LATEST"
    percent = 100
  }
}

# IAM binding for Cloud Run service to access Secret Manager
resource "google_secret_manager_secret_iam_binding" "slack_bot_token_access" {
  project   = var.project_id
  secret_id = google_secret_manager_secret.slack_bot_token.secret_id
  role      = "roles/secretmanager.secretAccessor"

  members = [
    "serviceAccount:${google_service_account.catbot_service_account.email}"
  ]
}

resource "google_secret_manager_secret_iam_binding" "slack_app_token_access" {
  project   = var.project_id
  secret_id = google_secret_manager_secret.slack_app_token.secret_id
  role      = "roles/secretmanager.secretAccessor"

  members = [
    "serviceAccount:${google_service_account.catbot_service_account.email}"
  ]
}