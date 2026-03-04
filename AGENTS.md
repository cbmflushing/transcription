# Transcription — Agent Guidelines

<!-- lastUpdated: 2026-03-03 -->
> [!NOTE]
> If today's date is more than 7 days after the `lastUpdated` date above, ask the user: *"The AGENTS.md for this repo is more than 7 days old. Would you like me to re-generate it?"*

## Project Overview

Tooling for audio transcription of CBM sermons using OpenAI Whisper and Groq Cloud API. Focused on side-by-side English & Chinese audio processing.

## Scripts

### `youtube-and-whisper.sh`

Main workflow script: downloads audio from YouTube and transcribes it via Whisper.

```bash
./youtube-and-whisper.sh <youtube_url>
```

### `strip-chinese/`

Go utilities for bilingual text processing:
- `strip-chinese.go` — strips Chinese characters from bilingual text
- `strip-english.go` — strips English characters from bilingual text

```bash
cd strip-chinese
go run strip-chinese.go
go run strip-english.go
```

## API Key

`groq-cloud-api-key` — plaintext file containing the Groq API key (do **not** commit this if rotating keys).

## Whisper Options

| Tool | Notes |
|------|-------|
| [openai/whisper](https://github.com/openai/whisper) | Original, local |
| [whisper-ctranslate2](https://github.com/Softcatala/whisper-ctranslate2) | Faster drop-in replacement |
| [whisper-jax (HuggingFace)](https://huggingface.co/spaces/sanchit-gandhi/whisper-jax) | Fast web version, easiest to use |
| [whisper-diarization](https://github.com/MahmoudAshraf97/whisper-diarization) | Multi-speaker support |
| [whishper](https://github.com/pluja/whishper) | Local web GUI |

## Examples

See `examples/` directory for sample input/output files.
