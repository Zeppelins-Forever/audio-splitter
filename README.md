# audio-splitter
Split audio files and audiobooks into more, smaller files. Good for vehicle media systems.

### Usage
1. Make sure you have `ffmpeg` in your system PATH.
2. Run `audio-splitter` (or `audio-splitter.exe` if on Windows) to automatically split all audio files in the current directory into 5 minute chunks, and place them in an `output` directory. This does not delete the original files.
   * If you want a different time split than 5 minutes, use the `-t <seconds>` flag to specify your value. Ex. `audio-splitter -t 600` for a 10 minute (600 second) split.
   * Get this info with the `-h` or `--help` flag. Ex. `audio-splitter -h`.
