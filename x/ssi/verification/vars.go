package verification

var ServiceTypes = []string{
	"LinkedDomains",
}

const DidMethod string = "mid"

const ClaimStatus_live = "Live"
const ClaimStatus_suspended = "Suspended"
const ClaimStatus_revoked = "Revoked"
const ClaimStatus_expired = "Expired"

func GetAcceptedCredentialClaimStatuses() []string {
	return []string{
		ClaimStatus_live,
		ClaimStatus_suspended,
		ClaimStatus_revoked,
		ClaimStatus_expired,
	}
}
