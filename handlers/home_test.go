package handlers

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotFoundPage(t *testing.T) {
	// Use html/template to match the type of Templates
	Templates = template.Must(template.New("404.html").Parse(`
		<html>
			<head><title>404 Not Found</title></head>
			<body>
				<div>
					<h1>404 - Page Not Found</h1>
					<p>The page you are looking for does not exist.</p>
				</div>
			</body>
		</html>
	`))

	req := httptest.NewRequest("GET", "/404", nil)
	rr := httptest.NewRecorder()

	NotFoundHandler(rr, req)

	// Assert the response status code
	assert.Equal(t, http.StatusNotFound, rr.Code, "Expected status 404")

	body := rr.Body.String()
	assert.Contains(t, body, "404 - Page Not Found", "Response should contain the 404 message")
	assert.Contains(t, body, "The page you are looking for does not exist.", "Response should contain the additional message")
}
