package main

import (
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
	rootCmd.PersistentFlags().StringVarP(&serviceAccountFilePath, "serviceAccountFile", "s", "", "The Play publisher service account file")
	rootCmd.PersistentFlags().StringVarP(&packageNameID, "packageNameID", "p", "", "The package name ID")
	rootCmd.PersistentFlags().StringVarP(&apkFilePath, "apkFile", "a", "", "The path to the APK to upload")

	viper.SetEnvPrefix("pp")
	viper.AutomaticEnv()

	viper.BindEnv("serviceAccountFile")
	viper.BindEnv("packageNameID")
	viper.BindEnv("apkFile")

	viper.BindPFlag("serviceAccountFile", rootCmd.PersistentFlags().Lookup("serviceAccountFile"))
	viper.BindPFlag("packageNameID", rootCmd.PersistentFlags().Lookup("packageNameID"))
	viper.BindPFlag("apkFile", rootCmd.PersistentFlags().Lookup("apkFile"))

	var listApks = &cobra.Command{
		Use:   "list",
		Short: "List upload APK into the application in the Play Store",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := playpublisher.NewClient(viper.GetString("serviceAccountFile"))
			if err != nil {
				return err
			}

			return client.ListService.List(viper.GetString("packageNameID"))
		},
	}

	var uploadApk = &cobra.Command{
		Use:   "upload",
		Short: "Upload APK binary to the PlayStore",
		RunE: func(cmd *cobra.Command, args []string) error {
			// fmt.Printf("Inside subCmd Run with args: %v\n", args)
			client, err := playpublisher.NewClient(viper.GetString("serviceAccountFile"))
			if err != nil {
				return err
			}

			file, err := os.Open(viper.GetString("apkFile"))
			if err != nil {
				return err
			}
			defer file.Close()

			return client.UploadService.Upload(viper.GetString("packageNameID"), file, "alpha")
		},
	}

	rootCmd.AddCommand(listApks)
	rootCmd.AddCommand(uploadApk)

	rootCmd.Execute()
}
