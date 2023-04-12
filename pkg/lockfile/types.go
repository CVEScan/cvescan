package lockfile

import "github.com/cvescan/cvescan/internal"

type Ecosystem = internal.Ecosystem
type PackageDetails = internal.PackageDetails
type PackageDetailsParser = func(pathToLockfile string) ([]PackageDetails, error)
