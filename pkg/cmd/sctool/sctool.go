// Copyright (C) 2017 ScyllaDB

package main

import (
	"io"
	"log"
	"os"

	"github.com/scylladb/scylla-manager/pkg/command/backup"
	"github.com/scylladb/scylla-manager/pkg/command/backup/backupdelete"
	"github.com/scylladb/scylla-manager/pkg/command/backup/backupfiles"
	"github.com/scylladb/scylla-manager/pkg/command/backup/backuplist"
	"github.com/scylladb/scylla-manager/pkg/command/backup/backupvalidate"
	"github.com/scylladb/scylla-manager/pkg/command/cluster/clusteradd"
	"github.com/scylladb/scylla-manager/pkg/command/cluster/clusterdelete"
	"github.com/scylladb/scylla-manager/pkg/command/cluster/clusterlist"
	"github.com/scylladb/scylla-manager/pkg/command/cluster/clusterupdate"
	"github.com/scylladb/scylla-manager/pkg/command/repair"
	"github.com/scylladb/scylla-manager/pkg/command/repair/repaircontrol"
	"github.com/scylladb/scylla-manager/pkg/command/resume"
	"github.com/scylladb/scylla-manager/pkg/command/status"
	"github.com/scylladb/scylla-manager/pkg/command/suspend"
	"github.com/scylladb/scylla-manager/pkg/command/task/taskdelete"
	"github.com/scylladb/scylla-manager/pkg/command/task/taskhistory"
	"github.com/scylladb/scylla-manager/pkg/command/task/tasklist"
	"github.com/scylladb/scylla-manager/pkg/command/task/taskprogress"
	"github.com/scylladb/scylla-manager/pkg/command/task/taskstart"
	"github.com/scylladb/scylla-manager/pkg/command/task/taskstop"
	"github.com/scylladb/scylla-manager/pkg/command/task/taskupdate"
	"github.com/scylladb/scylla-manager/pkg/command/version"
	"github.com/scylladb/scylla-manager/pkg/managerclient"
	"github.com/spf13/cobra"
)

func main() {
	log.SetOutput(io.Discard)

	cmd := buildCommand()
	if err := cmd.Execute(); err != nil {
		managerclient.PrintError(cmd.OutOrStderr(), err)
		os.Exit(1)
	}

	os.Exit(0)
}

func buildCommand() *cobra.Command {
	var client managerclient.Client

	backupCmd := backup.NewCommand(&client)
	backupCmd.AddCommand(
		backupdelete.NewCommand(&client),
		backupfiles.NewCommand(&client),
		backuplist.NewCommand(&client),
		backupvalidate.NewCommand(&client),
	)

	clusterCmd := &cobra.Command{
		Use:   "cluster",
		Short: "Add or delete clusters",
	}
	clusterCmd.AddCommand(
		clusteradd.NewCommand(&client),
		clusterdelete.NewCommand(&client),
		clusterlist.NewCommand(&client),
		clusterupdate.NewCommand(&client),
	)

	repairCmd := repair.NewCommand(&client)
	repairCmd.AddCommand(repaircontrol.NewCommand(&client))

	taskCmd := &cobra.Command{
		Use:   "task",
		Short: "Start, stop and track task progress",
	}
	taskCmd.AddCommand(
		taskdelete.NewCommand(&client),
		taskhistory.NewCommand(&client),
		tasklist.NewCommand(&client),
		taskprogress.NewCommand(&client),
		taskstart.NewCommand(&client),
		taskstop.NewCommand(&client),
		taskupdate.NewCommand(&client),
	)

	rootCmd := newRootCommand(&client)
	rootCmd.AddCommand(
		backupCmd,
		clusterCmd,
		repairCmd,
		resume.NewCommand(&client),
		status.NewCommand(&client),
		suspend.NewCommand(&client),
		taskCmd,
		version.NewCommand(&client),
	)

	setCommandDefaults(rootCmd)
	return rootCmd
}

func setCommandDefaults(cmd *cobra.Command) {
	// By default do not accept any arguments
	if cmd.Args == nil {
		cmd.Args = cobra.NoArgs
	}
	// Do not print errors, error printing is handled in main
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true

	// Call recursively.
	for _, c := range cmd.Commands() {
		setCommandDefaults(c)
	}
}
