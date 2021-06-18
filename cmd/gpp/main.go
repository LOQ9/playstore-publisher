package main

import (
	"fmt"
	"go-playstore-publisher/playpublisher"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	apkFilePath            string
	serviceAccountFilePath string
	packageNameID          string

	rootCmd = &cobra.Command{
		Use:   "playstore-publisher",
		Short: "Publish applications to PlayStore",
		Long:  "Allow publishing applications to the Google Play Store",
	}
)

func main() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&serviceAccountFilePath, "serviceAccountFile", "s", "", "The Play publisher service account file")
	rootCmd.PersistentFlags().StringVarP(&packageNameID, "packageNameID", "p", "", "The package name ID")
	rootCmd.PersistentFlags().StringVarP(&apkFilePath, "apkFile", "a", "", "The path to the APK to upload")
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))

	rootCmd.MarkPersistentFlagRequired("serviceAccountFile")
	rootCmd.MarkPersistentFlagRequired("packageNameID")
	rootCmd.MarkPersistentFlagRequired("apkFile")

	var listApks = &cobra.Command{
		Use:   "list",
		Short: "List upload APK into the application in the Play Store",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := playpublisher.NewClient(serviceAccountFilePath)
			if err != nil {
				return err
			}

			return client.ListService.List(packageNameID)
		},
	}

	var uploadApk = &cobra.Command{
		Use:   "upload",
		Short: "Upload APK binary to the PlayStore",
		RunE: func(cmd *cobra.Command, args []string) error {
			// fmt.Printf("Inside subCmd Run with args: %v\n", args)
			client, err := playpublisher.NewClient(serviceAccountFilePath)
			if err != nil {
				return err
			}

			file, err := os.Open(apkFilePath)
			if err != nil {
				return err
			}
			defer file.Close()

			return client.UploadService.Upload(packageNameID, file, "alpha")
		},
	}

	rootCmd.AddCommand(listApks)
	rootCmd.AddCommand(uploadApk)

	rootCmd.Execute()
}

func initConfig() {
	if serviceAccountFilePath != "" {
		// Use config file from the flag.
		viper.SetConfigFile(serviceAccountFilePath)
	}

	viper.SetEnvPrefix("PP")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
