package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"bilivdtool/core"
)

var version string

func main() {
	fmt.Println("bilivdtool v" + version)
	videoPath := flag.String("v", "", "video file path")
	audioPath := flag.String("a", "", "audio file path")
	outputPath := flag.String("o", "", "output file path")
	flag.Parse()

	if *videoPath == "" || *audioPath == "" || *outputPath == "" {
		fmt.Println("Usage: bilivdtool -v [video_path] -a [audio_path] -o [output_path]")
		fmt.Println("More Detail: https://github.com/aURORA-JC/bilivdtool")
		os.Exit(1)
	}

	if err := core.DoFileOperations(*videoPath); err != nil {
		slog.Error(err.Error())
		os.Exit(2)
	}

	if err := core.DoFileOperations(*audioPath); err != nil {
		slog.Error(err.Error())
		os.Exit(2)
	}

	if err := core.DoMergeOperations(*videoPath, *audioPath, *outputPath); err != nil {
		slog.Error(err.Error())
		os.Exit(3)
	}

	slog.Info("all operations executed successfully")
	slog.Info("files merged successfully")
}
