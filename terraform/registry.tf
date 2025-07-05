# Artifact Registry repository for CatBot
resource "google_artifact_registry_repository" "catbot" {
  location      = var.region
  project       = var.project_id
  repository_id = "catbot"
  description   = "Docker repository for CatBot Slack application"
  format        = "DOCKER"
}