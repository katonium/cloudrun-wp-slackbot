# Secret Manager secrets for Slack Bot
resource "google_secret_manager_secret" "slack_bot_token" {
  project   = var.project_id
  secret_id = "catbot-prod-slack-bot-token"

  replication {
    auto {}
  }
}

resource "google_secret_manager_secret" "slack_app_token" {
  project   = var.project_id
  secret_id = "catbot-prod-slack-app-token"

  replication {
    auto {}
  }
}