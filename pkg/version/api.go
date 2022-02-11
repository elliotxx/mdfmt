package version

import (
	"runtime"
	"time"
)

func init() {
	if versionInfo == nil {
		versionInfo = &Info{
			ReleaseVersion: "develop",
			GitVersion:     "develop",
			GitCommit:      "",
			GitTreeState:   "",
			BuildTime:      time.Now().Format("2006-01-02 15:04:05"),
			Runtime: RuntimeInfo{
				GoVersion: runtime.Version(),
				GOOS:      runtime.GOOS,
				GOARCH:    runtime.GOARCH,
				NumCPU:    runtime.NumCPU(),
				Compiler:  runtime.Compiler,
			},
		}
	}
}

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
