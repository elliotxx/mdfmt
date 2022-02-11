package version

func ReleaseVersion() string {
	return versionInfo.ReleaseVersion
}

func String() string {
	return versionInfo.String()
}

func ShortString() string {
	return versionInfo.ShortString()
}

func JSON() string {
	return versionInfo.JSON()
}

func YAML() string {
	return versionInfo.YAML()
}
