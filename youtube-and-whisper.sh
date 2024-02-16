#!/bin/zsh
# change the shebang above to bash if needed
# USAGE: youtube-and-whisper.sh {youtube-link}
# e.g. ./youtube-and-whisper.sh https://www.youtube.com/watch?v=wcBpuiefGaw

# Check if yt-dlp is installed
if ! command -v yt-dlp &> /dev/null; then
    echo "Error: yt-dlp is not installed. Please install it first."
    exit 1
fi

# Check if whisper is installed
if ! command -v whisper &> /dev/null; then
    echo "Error: whisper is not installed. Please install it first."
    exit 1
fi

# Check if URL argument is provided
if [ -z "$1" ]; then
    echo "Usage: $0 <YouTube URL>"
    exit 1
fi

# Download audio using yt-dlp
yt-dlp --extract-audio --audio-format mp3 --audio-quality 10 -o '%(id)s.%(ext)s' "$1"

# Check if download was successful
if [ $? -ne 0 ]; then
    echo "Error: Failed to download audio from the provided URL."
    exit 1
fi

# Pass filename to whisper command
filename=$(yt-dlp --get-filename --output '%(id)s.mp3' "$1")
whisper "$filename" --model small --language Chinese | tee  "$filename"-chinese.txt
whisper "$filename" --model small --language English | tee  "$filename"-english.txt

