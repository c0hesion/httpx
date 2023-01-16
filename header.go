package httpx

// Standard HTTP Headers as derived from:
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers
const (
	WWWAuthenticateHeader    = "WWW-Authenticate"
	AuthorizationHeader      = "Authorization"
	ProxyAuthenticateHeader  = "Proxy-Authenticate"
	ProxyAuthorizationHeader = "Proxy-Authorization"

	AgeHeader           = "Age"
	CacheControlHeader  = "Cache-Control"
	ClearSiteDataHeader = "Clear-Site-Data"
	ExpiresHeader       = "Expires"
	PragmaHeader        = "Pragma"
	WarningHeader       = "Warning"

	DownlinkHeader = "Downlink"
	ECTHeader      = "ECT"
	RTTHeader      = "RTT"

	LastModifiedHeader      = "Last-Modified"
	ETagHeader              = "ETag"
	IfMatchHeader           = "If-Match"
	IfNoneMatchHeader       = "If-None-Match"
	IfModifiedSinceHeader   = "If-Modified-Since"
	IfUnmodifiedSinceHeader = "If-Unmodified-Since"
	VaryHeader              = "Vary"

	ConnectionHeader = "Connection"
	KeepAliveHeader  = "Keep-Alive"

	AcceptHeader         = "Accept"
	AcceptEncodingHeader = "Accept-Encoding"
	AcceptLanguageHeader = "Accept-Language"

	ExpectHeader      = "Expect"
	MaxForwardsHeader = "Max-Forwards"

	CookieHeader    = "Cookie"
	SetCookieHeader = "Set-Cookie"

	AccessControlAllowOriginHeader      = "Access-Control-Allow-Origin"
	AccessControlAllowCredentialsHeader = "Access-Control-Allow-Credentials"
	AccessControlAllowHeadersHeader     = "Access-Control-Allow-Headers"
	AccessControlAllowMethodsHeader     = "Access-Control-Allow-Methods"
	AccessControlExposeHeadersHeader    = "Access-Control-Expose-Headers"
	AccessControlMaxAgeHeader           = "Access-Control-Max-Age"
	AccessControlRequestHeadersHeader   = "Access-Control-Request-Headers"
	AccessControlRequestMethodHeader    = "Access-Control-Request-Method"
	OriginHeader                        = "Origin"
	TimingAllowOriginHeader             = "Timing-Allow-Origin"

	ContentDispositionHeader = "Content-Disposition"
	ContentLengthHeader      = "Content-Length"
	ContentTypeHeader        = "Content-Type"
	ContentEncodingHeader    = "Content-Encoding"
	ContentLanguageHeader    = "Content-Language"
	ContentLocationHeader    = "Content-Location"

	ForwardedHeader = "Forwarded"
	ViaHeader       = "Via"

	LocationHeader = "Location"
	RefreshHeader  = "Refresh"

	FromHeader           = "From"
	HostHeader           = "Host"
	ReferrerHeader       = "Referrer"
	ReferrerPolicyHeader = "Referrer-Policy"
	UserAgentHeader      = "User-Agent"

	AllowHeader  = "Allow"
	ServerHeader = "Server"

	AcceptRangesHeader = "Accept-Ranges"
	RangeHeader        = "Range"
	IfRangeHeader      = "If-Range"
	ContentRangeHeader = "Content-Range"

	CrossOriginEmbedderPolicyHeader       = "Cross-Origin-Embedder-Policy"
	CrossOriginOpenerPolicyHeader         = "Cross-Origin-Opener-Policy"
	CrossOriginResourcePolicyHeader       = "Cross-Origin-Resource-Policy"
	ContentSecurityPolicyHeader           = "Content-Security-Policy"
	ContentSecurityPolicyReportOnlyHeader = "Content-Security-Policy-Report-Only"
	ExpectCTHeader                        = "Expect-CT"
	PermissionsPolicyHeader               = "Permissions-Policy"
	StrictTransportSecurityHeader         = "Strict-Transport-Security"
	UpgradeInsecureRequestsHeader         = "Upgrade-Insecure-Requests"
	XContentTypeOptionsHeader             = "X-Content-Type-Options"
	XDownloadOptionsHeader                = "X-Download-Options"
	XFrameOptionHeader                    = "X-Frame-Option"
	XPermittedCrossDomainPoliciesHeader   = "X-Permitted-Cross-Domain-Policies"
	XPoweredByHeader                      = "X-Powered-By"
	XXSSProtectionHeader                  = "X-XSS-Protection"

	SecFetchSiteHeader                   = "Sec-Fetch-Site"
	SecFetchModeHeader                   = "Sec-Fetch-Mode"
	SecFetchUserHeader                   = "Sec-Fetch-User"
	SecFetchDestHeader                   = "Sec-Fetch-Dest"
	ServiceWorkerNavigationPreloadHeader = "Service-Worker-Navigation-Preload"

	LastEventIDHeader = "Last-Event-ID"
	PingFromHeader    = "Ping-From"
	PingToHeader      = "Ping-To"
	ReportToHeader    = "Report-To"

	TransferEncodingHeader = "Transfer-Encoding"
	TEHeader               = "TE"
	TrailerHeader          = "Trailer"

	SecWebSocketKeyHeader        = "Sec-WebSocket-Key"
	SecWebSocketExtensionsHeader = "Sec-WebSocket-Extensions"
	SecWebSocketAcceptHeader     = "Sec-WebSocket-Accept"
	SecWebSocketProtocolHeader   = "Sec-WebSocket-Protocol"
	SecWebSocketVersionHeader    = "Sec-WebSocket-Version"

	AltSvcHeader               = "Alt-Svc"
	DateHeader                 = "Date"
	LargeAllocationHeader      = "Large-Allocation"
	LinkHeader                 = "Link"
	RetryAfterHeader           = "Retry-After"
	ServerTimingHeader         = "Server-Timing"
	ServiceWorkerAllowedHeader = "Service-Worker-Allowed"
	SourceMapHeader            = "SourceMap"
	UpgradeHeader              = "Upgrade"
	XDNSPrefetchControlHeader  = "X-DNS-Prefetch-Control"
	XRequestedWithHeader       = "X-Requested-With"
)
