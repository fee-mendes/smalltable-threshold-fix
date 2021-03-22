// Copyright (C) 2017 ScyllaDB

package downloader_test

import (
	"context"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/operations"
	"github.com/scylladb/go-log"
	"github.com/scylladb/scylla-manager/pkg/downloader"
	backup "github.com/scylladb/scylla-manager/pkg/service/backup/backupspec"
	"github.com/scylladb/scylla-manager/pkg/testutils"
	"github.com/scylladb/scylla-manager/pkg/util/uuid"
	"go.uber.org/zap/zapcore"
)

func TestDownload(t *testing.T) {
	var (
		location = backup.Location{Provider: "testdata"}
		criteria = downloader.ManifestLookupCriteria{
			NodeID:      uuid.MustParse("942ba1b6-30a3-441e-ac3c-158864d8b861"),
			SnapshotTag: "sm_20210215151954UTC",
		}
	)

	tmpDir, err := ioutil.TempDir("", "scylla-manager-rclone")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	table := []struct {
		Name     string
		Decorate func(*downloader.Downloader)
	}{
		{
			Name:     "Default",
			Decorate: func(*downloader.Downloader) {},
		},
		{
			Name:     "Upload dir mode",
			Decorate: func(d *downloader.Downloader) { d.WithTableDirMode(downloader.UploadTableDirMode) },
		},
		{
			Name:     "SSTable dir mode",
			Decorate: func(d *downloader.Downloader) { d.WithTableDirMode(downloader.SSTableLoaderTableDirMode) },
		},
		{
			Name:     "Filter single table",
			Decorate: func(d *downloader.Downloader) { d.WithKeyspace([]string{"system_auth.roles"}) },
		},
		{
			Name:     "Filter many tables",
			Decorate: func(d *downloader.Downloader) { d.WithKeyspace([]string{"system_a*"}) },
		},
		{
			Name: "Clear table",
			Decorate: func(d *downloader.Downloader) {
				d.WithKeyspace([]string{"system_auth.role_permissions"})

				dir := path.Join(d.Root(), "system_auth", "role_permissions-f4d5d0c0671be202bc241807c243e80b")
				os.MkdirAll(dir, 0755)
				ioutil.WriteFile(path.Join(dir, "a"), []byte("foo"), 0755)
				ioutil.WriteFile(path.Join(dir, "b"), []byte("bar"), 0755)

				d.WithClearTables()
			},
		},
		{
			Name:     "Dry run",
			Decorate: func(d *downloader.Downloader) { d.WithDryRun() },
		},
	}

	var (
		ctx    = context.Background()
		logger = log.NewDevelopmentWithLevel(zapcore.InfoLevel)
	)
	for i := range table {
		test := table[i]
		t.Run(test.Name, func(t *testing.T) {
			dir := path.Join(tmpDir, path.Base(t.Name()))

			d, err := downloader.New(location, dir, logger)
			if err != nil {
				t.Fatal("New() error", err)
			}
			test.Decorate(d)

			m, err := d.LookupManifest(ctx, criteria)
			if err != nil {
				t.Fatal("LookupManifest() error", err)
			}

			if err := d.Download(ctx, m); err != nil {
				t.Error(err)
			}
			var files []string

			fdst, err := fs.NewFs(ctx, dir)
			if err != nil {
				t.Fatal("NewFs() error", err)
			}
			operations.ListJSON(ctx, fdst, "", &operations.ListJSONOpt{Recurse: true, FilesOnly: true}, func(item *operations.ListJSONItem) error {
				files = append(files, item.Path)
				return nil
			})
			sort.Strings(files)
			testutils.SaveGoldenJSONFileIfNeeded(t, files)
			var golden []string
			testutils.LoadGoldenJSONFile(t, &golden)
			if diff := cmp.Diff(files, golden); diff != "" {
				t.Error(diff)
			}
		})
	}
}
