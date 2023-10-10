package main

import (
    "fmt"
)

// EmailService defines the interface for sending emails.
type EmailService interface {
    SendEmail(to, subject, body string) error
}

// SMTPEmailService is an implementation of the EmailService using SMTP.
type SMTPEmailService struct {
    SMTPServer string
    Port       int
    Username   string
    Password   string
}

func (s SMTPEmailService) SendEmail(to, subject, body string) error {
    // Implementation for sending email using SMTP
    fmt.Printf("Sending email to %s via SMTP...\n", to)
    return nil
}

// HTTPAPIEmailService is an implementation of the EmailService using an HTTP API.
type HTTPAPIEmailService struct {
    APIEndpoint string
    APIKey      string
}

func (s HTTPAPIEmailService) SendEmail(to, subject, body string) error {
    // Implementation for sending email using an HTTP API
    fmt.Printf("Sending email to %s via HTTP API...\n", to)
    return nil
}

// AppService is a service that depends on an EmailService.
type AppService struct {
    EmailService EmailService
}

func (app AppService) SendWelcomeEmail(userEmail string) {
    subject := "Welcome to our application"
    body := "Thank you for signing up!"
    
    // Use the injected EmailService to send the email
    err := app.EmailService.SendEmail(userEmail, subject, body)
    if err != nil {
        fmt.Printf("Error sending email: %v\n", err)
    }
}

func main() {
    smtpEmailService := SMTPEmailService{
        SMTPServer: "smtp.example.com",
        Port:       587,
        Username:   "your_username",
        Password:   "your_password",
    }
    
    httpAPIEmailService := HTTPAPIEmailService{
        APIEndpoint: "https://api.example.com/send-email",
        APIKey:      "your_api_key",
    }
    
    app := AppService{EmailService: smtpEmailService}
    app.SendWelcomeEmail("user@example.com")
    
    // To switch to a different email service, simply change the injected service.
    app.EmailService = httpAPIEmailService
    app.SendWelcomeEmail("anotheruser@example.com")
}
