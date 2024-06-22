package where

import "github.com/mjwhitta/safety"

var cache *safety.Map = safety.NewMap()

var defPaths map[string][]string = map[string][]string{
	"darwin": {
		"/usr/bin/core_perl",
		"/usr/bin/vendor_perl",
		"/usr/local/bin",
		"/usr/local/sbin",
		"/usr/bin",
		"/usr/sbin",
		"/bin",
		"/sbin",
		"/usr/local/osx-ndk-x86/bin",
	},
	"linux": {
		"/usr/bin/core_perl",
		"/usr/bin/vendor_perl",
		"/usr/local/bin",
		"/usr/local/sbin",
		"/usr/bin",
		"/usr/sbin",
		"/bin",
		"/sbin",
	},
	"windows": {
		"c:/program files/dotnet/",
		"c:/program files/git/cmd",
		"c:/program files/go/bin",
		"c:/program files/powershell/7",
		"c:/programdata/chocolatey/bin",
		"c:/windows/system32",
		"c:/windows",
		"c:/windows/system32/wbem",
		"c:/windows/system32/windowspowershell/v1.0/",
		"c:/windows/system32/openssh/",
	},
}

// Version is the package version.
const Version string = "1.3.3"
