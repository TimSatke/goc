package goc

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Cfg is the globally used configuration.
	Cfg *viper.Viper
)

func init() {
	rootCmd.AddCommand(prefsCmd)

	Cfg = loadCfg()
}

var prefsCmd = &cobra.Command{
	Use:     "prefs",
	Aliases: []string{"preferences", "settings", "opts", "options"},
	Short:   "TODO",
	Long:    "TODO",
}

func loadCfg() *viper.Viper {
	cfg := viper.New()

	cfg.SetDefault("cmd.output.directory", os.Getenv("HOME")+"/Development/tools")
	cfg.SetDefault("cmd.define.editor", "vim")
	cfg.SetDefault("cmd.undefine.prompt", true)

	cd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	cfg.AddConfigPath(cd)
	cfg.SetConfigFile("define.yaml")

	err = cfg.ReadInConfig()
	if err != nil {
		Printf("Error while reading: %v\n", err)
	}

	err = cfg.WriteConfig()
	if err != nil {
		panic(err)
	}

	return cfg
}
