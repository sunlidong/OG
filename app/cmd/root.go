// Copyright © 2019 Annchain Authors <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bytes"
	"fmt"
	"github.com/annchain/OG/common/goroutine"
	"github.com/annchain/OG/consensus/annsensus"
	"github.com/annchain/OG/og"
	"github.com/annchain/OG/og/downloader"
	"github.com/annchain/OG/og/fetcher"
	"github.com/annchain/OG/og/syncer"
	"github.com/annchain/OG/p2p"
	"github.com/rifflock/lfshook"
	"io/ioutil"
	"os"
	"runtime/debug"
	"time"

	"github.com/annchain/OG/mylog"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"path"
	"path/filepath"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "OG",
	Short: "OG: The next generation of DLT",
	Long:  `OG to da moon`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	defer DumpStack()
	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Fatalf("Fatal error occurred. Program will exit")
		os.Exit(1)
	}
}

func DumpStack() {
	if err := recover(); err != nil {
		logrus.WithField("obj", err).Error("Fatal error occurred. Program will exit")
		var buf bytes.Buffer
		stack := debug.Stack()
		buf.WriteString(fmt.Sprintf("Panic: %v\n", err))
		buf.Write(stack)
		dumpName := "dump_" + time.Now().Format("20060102-150405")
		nerr := ioutil.WriteFile(dumpName, buf.Bytes(), 0644)
		if nerr != nil {
			fmt.Println("write dump file error", nerr)
			fmt.Println(buf.String())
		}
		logrus.WithField("stack ", buf.String()).Error("panic")
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringP("datadir", "d", "data", fmt.Sprintf("Runtime directory for storage and configurations"))
	rootCmd.PersistentFlags().StringP("config", "c", "config.toml", "Path for configuration file or url of config server")
	rootCmd.PersistentFlags().BoolP("genkey", "k", false, "Automatically generate a private key if the privkey is missing.")
	rootCmd.PersistentFlags().StringP("log_dir", "l", "", "Path for configuration file. Not enabled by default")
	rootCmd.PersistentFlags().BoolP("log_stdout", "s", false, "Whether the log will be printed to stdout")
	rootCmd.PersistentFlags().StringP("log_level", "v", "debug", "Logging verbosity, possible values:[panic, fatal, error, warn, info, debug]")
	rootCmd.PersistentFlags().BoolP("log_line_number", "n", false, "log_line_number")
	rootCmd.PersistentFlags().BoolP("multifile_by_level", "m", false, "multifile_by_level")
	rootCmd.PersistentFlags().BoolP("multifile_by_module", "M", false, "multifile_by_module")

	_ = viper.BindPFlag("datadir", rootCmd.PersistentFlags().Lookup("datadir"))
	_ = viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	_ = viper.BindPFlag("genkey", rootCmd.PersistentFlags().Lookup("genkey"))
	_ = viper.BindPFlag("log.log_dir", rootCmd.PersistentFlags().Lookup("log_dir"))
	_ = viper.BindPFlag("log_line_number", rootCmd.PersistentFlags().Lookup("log_line_number"))
	_ = viper.BindPFlag("multifile_by_level", rootCmd.PersistentFlags().Lookup("multifile_by_level"))
	_ = viper.BindPFlag("multifile_by_module", rootCmd.PersistentFlags().Lookup("multifile_by_module"))
	//viper.BindPFlag("log_stdout", rootCmd.PersistentFlags().Lookup("log_stdout"))
	_ = viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("log_level"))

	viper.SetDefault("hub.outgoing_buffer_size", 100)
	viper.SetDefault("hub.incoming_buffer_size", 100)
	viper.SetDefault("hub.message_cache_expiration_seconds", 60)
	viper.SetDefault("hub.message_cache_max_size", 30000)
	viper.SetDefault("crypto.algorithm", "ed25519")
	viper.SetDefault("tx_buffer.new_tx_queue_size", 1)

	viper.SetDefault("max_tx_hash", "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")
	viper.SetDefault("max_mined_hash", "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")

	viper.SetDefault("debug.node_id", 0)
}

func panicIfError(err error, message string) {
	if err != nil {
		fmt.Println(message)
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

// initLogger uses viper to get the log path and level. It should be called by all other commands
func initLogger() {
	logdir := viper.GetString("log.log_dir")
	stdout := viper.GetBool("log_stdout")

	var writer io.Writer

	if logdir != "" {
		folderPath, err := filepath.Abs(logdir)
		panicIfError(err, fmt.Sprintf("Error on parsing log path: %s", logdir))

		abspath, err := filepath.Abs(path.Join(logdir, "run"))
		panicIfError(err, fmt.Sprintf("Error on parsing log file path: %s", logdir))

		err = os.MkdirAll(folderPath, os.ModePerm)
		panicIfError(err, fmt.Sprintf("Error on creating log dir: %s", folderPath))

		if stdout {
			logFile, err := os.OpenFile(abspath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			panicIfError(err, fmt.Sprintf("Error on creating log file: %s", abspath))
			abspath += ".log"
			fmt.Println("Will be logged to stdout and ", abspath)
			writer = io.MultiWriter(os.Stdout, logFile)
		} else {
			fmt.Println("Will be logged to ", abspath+".log")
			writer = mylog.RotateLog(abspath)
		}
	} else {
		// stdout only
		fmt.Println("Will be logged to stdout")
		writer = os.Stdout
	}

	logrus.SetOutput(writer)

	// Only log the warning severity or above.
	switch viper.GetString("log.level") {
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	default:
		fmt.Println("Unknown level", viper.GetString("log.level"), "Set to INFO")
		logrus.SetLevel(logrus.InfoLevel)
	}

	Formatter := new(logrus.TextFormatter)
	Formatter.ForceColors = logdir == ""
	//Formatter.DisableColors = true
	Formatter.TimestampFormat = "2006-01-02 15:04:05.000000"
	Formatter.FullTimestamp = true

	logrus.SetFormatter(Formatter)

	// redirect standard log to logrus
	//log.SetOutput(logrus.StandardLogger().Writer())
	//log.Println("Standard logger. Am I here?")
	lineNum := viper.GetBool("log_line_number")
	if lineNum {
		//filenameHook := filename.NewHook()
		//filenameHook.Field = "line"
		//logrus.AddHook(filenameHook)
		logrus.SetReportCaller(true)
	}
	byLevel := viper.GetBool("multifile_by_level")
	if byLevel && logdir != "" {
		panicLog, _ := filepath.Abs(path.Join(logdir, "panic"))
		fatalLog, _ := filepath.Abs(path.Join(logdir, "fatal"))
		warnLog, _ := filepath.Abs(path.Join(logdir, "warn"))
		errorLog, _ := filepath.Abs(path.Join(logdir, "error"))
		infoLog, _ := filepath.Abs(path.Join(logdir, "info"))
		debugLog, _ := filepath.Abs(path.Join(logdir, "debug"))
		traceLog, _ := filepath.Abs(path.Join(logdir, "trace"))
		writerMap := lfshook.WriterMap{
			logrus.PanicLevel: mylog.RotateLog(panicLog),
			logrus.FatalLevel: mylog.RotateLog(fatalLog),
			logrus.WarnLevel:  mylog.RotateLog(warnLog),
			logrus.ErrorLevel: mylog.RotateLog(errorLog),
			logrus.InfoLevel:  mylog.RotateLog(infoLog),
			logrus.DebugLevel: mylog.RotateLog(debugLog),
			logrus.TraceLevel: mylog.RotateLog(traceLog),
		}
		logrus.AddHook(lfshook.NewHook(
			writerMap,
			Formatter,
		))
	}
	logger := logrus.StandardLogger()
	logrus.Debug("Logger initialized.")
	byModule := viper.GetBool("multifile_by_module")
	if !byModule {
		logdir = ""
	}

	downloader.InitLoggers(logger, logdir)
	fetcher.InitLoggers(logger, logdir)
	p2p.InitLoggers(logger, logdir)
	og.InitLoggers(logger, logdir)
	syncer.InitLoggers(logger, logdir)
	annsensus.InitLoggers(logger, logdir)

}

func startPerformanceMonitor() {
	function := func() {
		logrus.WithField("port", viper.GetString("profiling.port")).Info("Performance monitor started")
		log.Println(http.ListenAndServe("0.0.0.0:"+viper.GetString("profiling.port"), nil))
	}
	goroutine.New(function)
}
