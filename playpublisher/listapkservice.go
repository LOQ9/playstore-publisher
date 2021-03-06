package playpublisher

import (
	"fmt"
)

// ListApkService definition
type ListApkService struct {
	helper *AndroidPublisherHelper
}

// List allows to list all the APK for the target bundle identifier into the PlayStore account
func (s *ListApkService) List(packageName string) error {
	fmt.Printf("Resolving APK list for package: %s\n", packageName)

	editID, err := s.helper.insertEdit(packageName)
	if err != nil {
		return err
	}

	apks, err := s.helper.listApk(packageName, editID)
	if err != nil {
		return err
	}

	for _, v := range apks {
		fmt.Printf("\tVersion: %v - Binary sha1: %v\n", v.VersionCode, v.Binary.Sha1)
	}

	return nil
}
