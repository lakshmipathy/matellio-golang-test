package config

func GetUploadFilePath() string {
	return getConfigString("upload.file_name")
}
