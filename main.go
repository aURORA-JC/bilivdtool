package main

import (
	"bilivdtool/core"
	"flag"
	"fmt"
	"log/slog"
	"os"
)

const version = ""

func main() {
	fmt.Println("bilivdtool v" + version)
	videoPath := flag.String("v", "", "video file path")
	audioPath := flag.String("a", "", "audio file path")
	outputPath := flag.String("o", "", "output file path")
	flag.Parse()

	if *videoPath == "" || *audioPath == "" || *outputPath == "" {
		fmt.Println("\tUsage: bilivdtool -v [video_path] -a [audio_path] -o [output_path]")
		fmt.Println("\t\tMore Detail: https://github.com/aURORA-JC/bilivdtool")
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
