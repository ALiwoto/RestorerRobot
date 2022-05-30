package backupUtils

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoConfig"
	"github.com/AnimeKaizoku/ssg/ssg"
	fErrors "github.com/go-faster/errors"
)

// BackupDatabase backups a database using its url to the specified filename
// and the type.
// currently only "sql" and "dump" types are supported for the third
// argument.
func BackupDatabase(url, filename, bType string) error {
	backupCommand := wotoConfig.GetPgDumpCommand() + " -d " + url + " "
	if bType == wotoConfig.BackupTypeSQL {
		backupCommand += ">> " + filename
	} else if bType == wotoConfig.BackupTypeDump {
		backupCommand += "> " + filename
	} else {
		return errors.New("unsupported backup type")
	}

	result := ssg.RunCommand(backupCommand)
	if result.Error != nil {
		return fErrors.Wrap(result.Error, result.Stderr+result.Stdout)
	}

	return nil
}

// ZipSource converts a source file/directory to the destination zip file.
func ZipSource(source, target string) error {
	// 1. Create a ZIP file and zip.Writer
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	// 2. Go through all the files of the source
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 3. Create a local file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// set compression
		header.Method = zip.Deflate

		// 4. Set relative path of a file as the header name
		header.Name, err = filepath.Rel(filepath.Dir(source), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}

		// 5. Create writer for the file header and save content of the file
		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(headerWriter, f)
		return err
	})
}

// GenerateFileNameFromOrigin creates a filename from the origin,
// in "VALUE-backup-2022-5-16--15-30-59" format.
func GenerateFileNameFromValue(value string) string {
	return value + "-backup-" + time.Now().Format("2006-01-02--15-04-05")
}
