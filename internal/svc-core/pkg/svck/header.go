package svck

const (
	// Subject metadata

	SubjectIdHeader       = "x-subject-id"
	SubjectFullNameHeader = "x-subject-name"
	SubjectRoleHeader     = "x-subject-role"

	// Reverse proxy headers

	RealIPHeader           = "x-real-ip"
	ForwardedForHeader     = "x-forwarded-for"
	HeaderGatewayUserAgent = "x-gw-user-agent"
	HeaderUserAgent        = "user-agent"

	// Request metadata headers

	HeaderRequestId = "x-request-id"
)
