package auth

import (
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/gin-gonic/gin"
)

// ClerkAuthMiddleware adapts Clerk's HTTP middleware for Gin
func ClerkAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a wrapper to convert Gin to standard HTTP handler
		var handlerCalled bool
		wrappedHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handlerCalled = true
			// Extract session claims from Clerk context
			sessionClaims, ok := clerk.SessionClaimsFromContext(r.Context())
			if !ok {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "No valid session found",
				})
				return
			}
			
			// Store user ID in Gin context
			c.Set("currentUserId", sessionClaims.Subject)
			
			// Update the request in Gin context
			c.Request = r
			c.Next()
		})

		// Apply Clerk's RequireHeaderAuthorization middleware
		clerkMiddleware := clerkhttp.RequireHeaderAuthorization()
		clerkHandler := clerkMiddleware(wrappedHandler)

		// Execute the Clerk middleware
		clerkHandler.ServeHTTP(c.Writer, c.Request)

		// If the handler wasn't called, it means auth failed
		if !handlerCalled {
			c.Abort()
		}
	}
}