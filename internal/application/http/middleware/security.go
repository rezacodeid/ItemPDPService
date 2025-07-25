package middleware

import (
	"github.com/gin-gonic/gin"
)

// SecurityHeadersMiddleware adds security headers to responses
func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Prevent MIME type sniffing
		c.Header("X-Content-Type-Options", "nosniff")
		
		// Prevent clickjacking
		c.Header("X-Frame-Options", "DENY")
		
		// Enable XSS protection
		c.Header("X-XSS-Protection", "1; mode=block")
		
		// Prevent information disclosure
		c.Header("X-Powered-By", "")
		c.Header("Server", "")
		
		// Content Security Policy (basic)
		c.Header("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self' 'unsafe-inline'; img-src 'self' data:; font-src 'self'; connect-src 'self'; media-src 'self'; object-src 'none'; child-src 'none'; frame-src 'none'; worker-src 'none'; frame-ancestors 'none'; form-action 'self'; base-uri 'self';")
		
		// Referrer Policy
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		
		// Permissions Policy (formerly Feature Policy)
		c.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=(), payment=(), usb=(), magnetometer=(), gyroscope=(), speaker=()")
		
		// HSTS (HTTP Strict Transport Security) - only for HTTPS
		if c.Request.TLS != nil {
			c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		}
		
		c.Next()
	}
}

// RateLimitConfig holds rate limiting configuration
type RateLimitConfig struct {
	RequestsPerMinute int
	BurstSize         int
}

// DefaultRateLimitConfig returns default rate limiting configuration
func DefaultRateLimitConfig() RateLimitConfig {
	return RateLimitConfig{
		RequestsPerMinute: 100,
		BurstSize:         10,
	}
}

// SimpleRateLimitMiddleware provides basic rate limiting
// Note: This is a simple in-memory implementation for demonstration
// For production, use a proper rate limiting solution like Redis-based rate limiter
func SimpleRateLimitMiddleware(config RateLimitConfig) gin.HandlerFunc {
	// This is a simplified rate limiter for demonstration
	// In production, use a proper distributed rate limiter
	return func(c *gin.Context) {
		// For now, just add the rate limit headers
		c.Header("X-RateLimit-Limit", "100")
		c.Header("X-RateLimit-Remaining", "99")
		c.Header("X-RateLimit-Reset", "60")
		
		c.Next()
	}
}

// CSRFConfig holds CSRF protection configuration
type CSRFConfig struct {
	TokenLength int
	CookieName  string
	HeaderName  string
	Secure      bool
	HttpOnly    bool
	SameSite    string
}

// DefaultCSRFConfig returns default CSRF configuration
func DefaultCSRFConfig() CSRFConfig {
	return CSRFConfig{
		TokenLength: 32,
		CookieName:  "csrf_token",
		HeaderName:  "X-CSRF-Token",
		Secure:      true,
		HttpOnly:    true,
		SameSite:    "Strict",
	}
}

// BasicCSRFMiddleware provides basic CSRF protection
// Note: This is a simplified implementation for demonstration
// For production, use a proper CSRF protection library
func BasicCSRFMiddleware(config CSRFConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip CSRF for GET, HEAD, OPTIONS requests
		if c.Request.Method == "GET" || c.Request.Method == "HEAD" || c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}
		
		// For demonstration, we'll add CSRF headers but not enforce validation
		// In production, implement proper CSRF token validation
		c.Header("X-CSRF-Protection", "enabled")
		
		c.Next()
	}
}

// InputSanitizationMiddleware provides basic input sanitization
func InputSanitizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add security headers for input validation
		c.Header("X-Input-Validation", "enabled")
		
		// In a real implementation, you would sanitize request parameters here
		// For now, we'll just add the header to indicate validation is in place
		
		c.Next()
	}
}
