package core

import (
	"errors"
	"fmt"
	"log/slog"
	"os/exec"
)

const ffmpegName = "ffmpeg"

func DoMergeOperations(videoPath, audioPath, outputPath string) error {
	// set cmd
	cmd := exec.Command("./"+ffmpegName, "-y", "-i", videoPath, "-i", audioPath, "-c:v", "copy", "-c:a", "copy", outputPath)

	// check path valid
	_, err := exec.LookPath("./" + ffmpegName)
	if err != nil && !errors.Is(err, exec.ErrNotFound) {
		return err
	} else if errors.Is(err, exec.ErrNotFound) {
		slog.Info("can not find ffmpeg file in this dir, searching %PATH%")
		cmd = exec.Command(ffmpegName, "-y", "-i", videoPath, "-i", audioPath, "-c:v", "copy", "-c:a", "copy", outputPath)
	}

	// run command and get output (stdout, stderr)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	// print output string
	slog.Info(fmt.Sprintf("%s\n", string(out)))
	return nil
}
