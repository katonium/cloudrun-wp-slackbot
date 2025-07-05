# cloudrun-wp-slackbot
A Slack Bot template running on Google Cloud Run Worker Pools

## Slack Bot Setup

### Prerequisites
- Go 1.21+
- Docker
- Slack Bot Token and App Token

### Environment Variables
Create a `.env` file with:
```
SLACK_BOT_TOKEN=xoxb-your-bot-token
SLACK_APP_TOKEN=xapp-your-app-token
```

### Running the Application

**Build the container:**
```bash
mise run slackbot-build
```

**Run locally:**
```bash
mise run slackbot-run
```

**Publish to registry:**
```bash
mise run slackbot-publish
```

**Direct Go run:**
```bash
cd slackbot
go run main.go
```

The bot responds to messages containing "hello" with a greeting.

## Emojis for Commit messages and PRs

Inspired by [gitmoji.dev](https://gitmoji.dev/) and  [GitCommitEmoji.md](https://gist.github.com/parmentf/035de27d6ed1dce0b36a)

| Commit or PR type                                   | Emoji             |
|-----------------------------------------------------|------------------|
| Improve structure / format of the code.             | :art: `:art:`       |
| Improve performance.                                | :zap: `:zap:`       |
| Remove code or files.                               | :fire: `:fire:`      |
| Fix a bug.                                          | :bug: `:bug:`       |
| Refactor code.                                      | :recycle: `:recycle:`   |
| Critical hotfix.                                    | :ambulance: `:ambulance:` |
| Introduce new features.                             | :sparkles: `:sparkles:`   |
| Add or update documentation.                        | :memo: `:memo:`      |
| Documenting source code                             |	:bulb: `:bulb:`     |
| Deploy stuff.                                       | :rocket: `:rocket:`    |
| Begin a project.                                    | :tada: `:tada:`      |
| Add, update, or pass tests.                         | :white_check_mark: `:white_check_mark:` |
| Add or update secrets.                              | :closed_lock_with_key: `:closed_lock_with_key:` |
| Release / Version tags.                             | :bookmark: `:bookmark:`  |
| Work in progress.                                   | :construction: `:construction:` |
| Downgrade dependencies.                             | :arrow_down: `:arrow_down:` |
| Upgrade dependencies.                               | :arrow_up: `:arrow_up:`  |
| Pin dependencies to specific versions.              | :pushpin: `:pushpin:`   |
| Add or update CI build system.                      | :construction_worker: `:construction_worker:` |
| Fix CI Build.                                       | :green_heart: `:green_heart:` |
| Add a dependency.                                   | :heavy_plus_sign: `:heavy_plus_sign:` |
| Remove a dependency.                                | :heavy_minus_sign: `:heavy_minus_sign:` |
| Add or update configuration files.                  | :wrench: `:wrench:`    |
| Add or update development scripts.                  | :hammer: `:hammer:`    |
| Internationalization and localization.              | :globe_with_meridians: `:globe_with_meridians:` |
| Fix typos.                                          | :pencil2: `:pencil2:`   |
| Revert changes.                                     | :rewind: `:rewind:`    |
| Merge branches.                                     | :twisted_rightwards_arrows: `:twisted_rightwards_arrows:` |
| Add or update compiled files or packages.           | :package: `:package:`  |
| Update code due to external API changes.            | :alien: `:alien:`    |
| Move or rename resources (e.g.: files, paths, routes). | :truck: `:truck:` |
| Introduce breaking changes.                         | :boom: `:boom:`      |
| Add or update assets.                               | :bento: `:bento:`     |
| Add or update text and literals.                    | :speech_balloon: `:speech_balloon:` |
| Perform database related changes.                   | :card_file_box: `:card_file_box:` |
| Add or update logs.                                 | :loud_sound: `:loud_sound:` |
| Remove logs.                                        | :mute: `:mute:`      |
| Make architectural changes.                         | :building_construction: `:building_construction:` |
| Mock things.                                        | :clown_face: `:clown_face:` |
| Perform experiments.                                | :alembic: `:alembic:` |
| Add, update, or remove feature flags.               | :triangular_flag_on_post: `:triangular_flag_on_post:` |
| Work on code related to authorization, roles and permissions. | :passport_control: `:passport_control:` |
| Simple fix for a non-critical issue.                | :adhesive_bandage: `:adhesive_bandage:` |
| Add a failing test.                                 | :test_tube: `:test_tube:` |
| Add or update business logic.                       | :necktie: `:necktie:` |
| Add or update healthcheck.                          | :stethoscope: `:stethoscope:` |
| Infrastructure related changes.                     | :bricks: `:bricks:` |
| Improve developer experience.                       | :technologist: `:technologist:` |
| Add or update code related to multithreading or concurrency. | :thread: `:thread:` |
| Add or update code related to validation.           | :safety_vest: `:safety_vest:` |

