package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// List of common audiobook and audio extensions to look for.
var supportedExtensions = map[string]bool{
	".mp3":  true,
	".m4a":  true,
	".m4b":  true,
	".aac":  true,
	".ogg":  true,
	".flac": true,
	".wav":  true,
}

func main() {
	// 1. Define command-line flags.
	// Bind the "-t" flag to the 'splitTime' variable, set its default to 300,
	// and provide a description for the help menu.
	var splitTime int
	flag.IntVar(&splitTime, "t", 300, "Split time in seconds (default is 300 seconds / 5 minutes)")

	// Override the default help output. I'm using my own custom help messages.
	// Go's flag package automatically triggers this function if the user passes "-h" or "--help".
	flag.Usage = func() {
		fmt.Println("Audiobook Splitter: Easily slice audiobooks into segments without re-encoding.")
		fmt.Println("\nUsage:")
		fmt.Println("  audio-splitter [options]")
		fmt.Println("\n  Run this program within the same directory as a bunch of audio files (mp3, m4a, m4b, aac, ogg, flac, wav), and it will output split versions of those files into an 'output' folder, formatted as [original_file_name]__[number].[extension].")
		fmt.Println("\n\nOptions:")
		fmt.Println("  -h, --help     Show this help message and exit")
		fmt.Printf("  -t <seconds>   Set the splitting time per file in seconds (default: 300)\n\n")
		fmt.Println("\nExample:")
		fmt.Println("  audio-splitter -t 600   // Splits audio into 10-minute increments")
	}

	// Parse the flags. This MUST be called after defining the flags and before using them.
	flag.Parse()

	// 2. Check if ffmpeg is installed and available on the system PATH.
	_, err := exec.LookPath("ffmpeg")
	if err != nil {
		fmt.Println("Error: 'ffmpeg' was not found on your system PATH.")
		fmt.Println("Please install ffmpeg and ensure it is accessible from the command line.")
		os.Exit(1)
	}

	// 3. Create the "output" directory.
	outputDir := "output"
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// 4. Read all files in the current directory.
	entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("Error reading current directory: %v\n", err)
		os.Exit(1)
	}

	filesProcessed := 0

	// Convert our integer splitTime into a string so we can pass it to the exec.Command arguments.
	segmentTimeStr := fmt.Sprintf("%d", splitTime)

	// 5. Iterate through the directory entries.
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		fileName := entry.Name()
		ext := strings.ToLower(filepath.Ext(fileName))

		if supportedExtensions[ext] {
			fmt.Printf("Processing: %s (Splitting every %s seconds)...\n", fileName, segmentTimeStr)

			baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
			outputPattern := filepath.Join(outputDir, fmt.Sprintf("%s__%%03d%s", baseName, ext))

			// 6. Construct the FFmpeg command using the custom segmentTimeStr.
			cmd := exec.Command("ffmpeg",
				"-i", fileName,
				"-f", "segment",
				"-segment_time", segmentTimeStr, // This now dynamically uses the user's input (or 300)
				"-segment_start_number", "1",
				"-c", "copy",
				outputPattern,
			)

			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("Failed to process %s.\nFFmpeg Error: %v\nOutput: %s\n", fileName, err, string(output))
				continue
			}

			filesProcessed++
			fmt.Printf("Successfully split: %s\n", fileName)
		}
	}

	if filesProcessed == 0 {
		fmt.Println("No supported audio files found in the current directory.")
	} else {
		fmt.Printf("Done! Successfully processed %d file(s).\n", filesProcessed)
	}
}
