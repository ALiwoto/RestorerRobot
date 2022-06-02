package backupUtils

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoConfig"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoStyle"
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

// GenerateCaption generates caption for the backup using the specified options.
func GenerateCaption(opts *GenerateCaptionOptions) wotoStyle.WStyle {
	md := wotoStyle.GetBold("Config name: ").Mono(opts.ConfigName)
	md.Bold("\nType: ").Mono(opts.BackupInitType)
	md.Bold("\nInitiated by: ").Mono(opts.InitiatedBy)
	if opts.UserId != 0 {
		md.Bold("\nID: ").Mono(ssg.ToBase10(opts.UserId))
	}

	if !opts.DateTime.IsZero() {
		// format should be like: Wed-01-06-2022 11:39 AM
		md.Bold("\nDate Time: ").Mono(opts.DateTime.Format("Mon-01-02-2006 03:04 PM"))
	}

	if opts.FileSize != "" {
		startingTitle := "File"
		if opts.BackupFormat != "" {
			startingTitle = opts.BackupFormat
		}

		md.Bold("\n" + startingTitle + " size: ").Mono(opts.FileSize)
	}

	return md
}

func FormatFileSize(size int64) string {
	var sizeSuffix string
	var sizeValue float64

	if size > 1024*1024*1024 {
		sizeSuffix = "GB"
		sizeValue = float64(size) / 1024 / 1024 / 1024
	} else if size > 1024*1024 {
		sizeSuffix = "MB"
		sizeValue = float64(size) / 1024 / 1024
	} else if size > 1024 {
		sizeSuffix = "KB"
		sizeValue = float64(size) / 1024
	} else {
		sizeSuffix = "B"
		sizeValue = float64(size)
	}

	return fmt.Sprintf("%.4f", sizeValue) + " " + sizeSuffix
}

// GenerateFileNameFromOrigin creates a filename from the origin,
// in "VALUE-backup-2022-5-16--15-30-59" format.
func GenerateFileNameFromValue(value string) string {
	return value + "-backup-" + time.Now().Format("2006-01-02--15-04-05")
}
