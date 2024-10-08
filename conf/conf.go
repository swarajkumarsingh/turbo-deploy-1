package conf

import (
	"os"

	"github.com/swarajkumarsingh/turbo-deploy/constants"
)

var AWS_TOKEN string = os.Getenv("STAGE")
var AWS_REGION string = os.Getenv("STAGE")
var AWS_ACCESS_KEY string = os.Getenv("STAGE")
var AWS_SECRET_ACCESS_KEY string = os.Getenv("STAGE")

var ENV string = os.Getenv("STAGE")
var VaultKey string = os.Getenv("VAULT_KEY")
var SentryDSN string = os.Getenv("SENTRY_DSN")
var S3Bucket = os.Getenv("S3_BUCKET")
var JWTSecretKey = []byte("chotu_babu_is_not_chotu_any_more")

const (
	ENV_PROD  = constants.ENV_PROD
	ENV_UAT   = constants.ENV_UAT
	ENV_DEV   = constants.ENV_DEV
	ENV_LOCAL = constants.ENV_LOCAL
)

var ClientENV = "client"
const REGION = "ap-south-1"

const DDServiceName = "go-deployable-kyc"

// DDAgentHost is Hostname for Datadog agent
var DDAgentHost string = "172.17.0.1"