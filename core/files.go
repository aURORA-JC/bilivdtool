package core

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"os"
	"regexp"
)

const (
	unencryptedFileHeadPattern = `^\x00{3}.+$ftypiso5`
	paddingCharacterPattern    = `0`
	tempFilePattern            = "bvdttmp"
	chunkSize                  = 4096
)

func DoFileOperations(filePath string) error {
	// open file
	file, err := os.OpenFile(filePath, os.O_RDWR, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	// get head 20 byte
	var headStr string
	headStr, err = getHeaderString(file)
	if err != nil {
		return err
	}

	// calculate padding
	paddingLength := 0
	if !checkFilePrefix(headStr) {
		paddingLength = getPaddingLength(headStr)
	}

	// remove padding
	err = removePadding(file, paddingLength)
	if err != nil {
		return err
	}

	return nil
}

func getHeaderString(file *os.File) (string, error) {
	fileHead := make([]byte, 20)
	_, err := file.ReadAt(fileHead, 0)
	if err != nil {
		slog.Error("failed to read 20 byte from the start of this file")
		return "", err
	}

	fileHeadStr := string(fileHead)
	slog.Info("file start with: " + fileHeadStr)
	return fileHeadStr, nil
}

func checkFilePrefix(headStr string) bool {
	ufhPattern := regexp.MustCompile(unencryptedFileHeadPattern)
	matches := ufhPattern.FindString(headStr)

	return len(matches) != 0
}

func getPaddingLength(headStr string) int {
	pcPattern := regexp.MustCompile(paddingCharacterPattern)
	matches := pcPattern.FindAllString(headStr, -1)

	slog.Info(fmt.Sprintf("matched padding: %v", matches))
	return len(matches)
}

func removePadding(file *os.File, paddingLength int) error {
	slog.Info("padding length: " + fmt.Sprint(paddingLength))
	if paddingLength == 0 {
		slog.Info("padding length is 0, skipped remove operation")
		return nil
	}

	// create a tmp file
	tmpFile, err := os.CreateTemp("", tempFilePattern)
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())
	slog.Info("created tmp file at " + tmpFile.Name())

	// copy file to tmp
	rw := bufio.NewReadWriter(bufio.NewReader(file), bufio.NewWriter(tmpFile))
	if err = copyFile(rw, true, paddingLength); err != nil {
		return err
	}
	slog.Info("copied data to tmp")

	// move pointers to the start of files and copy from tmp to origin file
	file.Seek(0, 0)
	tmpFile.Seek(0, 0)
	rw = bufio.NewReadWriter(bufio.NewReader(tmpFile), bufio.NewWriter(file))
	if err = copyFile(rw, false, 0); err != nil {
		return err
	}
	slog.Info("moved pointers and copied data from tmp to origin")

	return nil
}

func copyFile(rw *bufio.ReadWriter, skipPadding bool, paddingLength int) error {
	// if read source has padding, skip padding
	if skipPadding {
		for i := 0; i < paddingLength; i++ {
			_, err := rw.ReadByte()
			if err != nil {
				return err
			}
		}
		slog.Info("skipped paddings")
	}

	// create chunk with CHUNK_SIZE, copy data chunk by chunk
	chunk := make([]byte, chunkSize)
	for {
		n, err := rw.Read(chunk)
		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			break
		}

		_, err = rw.Write(chunk[:n])
		if err != nil {
			return err
		}
	}

	// clear the buffer and write to disk
	err := rw.Flush()
	if err != nil {
		return err
	}
	slog.Info("data copied")

	return nil
}
