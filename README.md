# UniConnect Backend

UniConnect is a platform that connects users(students) with local businesses. This repository contains the backend API built with Go, which powers the user and business management features, including registration, login, listing, and interactions between users and businesses.

## Features

### User Features
- **User Registration**: Users can create an account by providing their name, email address, username, and password.
- **User Login**: Users can log in using their email address and password.
  
### Business Features
- **Business Registration**: Users can add businesses by submitting details such as business name, contact information, description, location, logo/image, business type, and website.
- **View Business Listings**: Users can browse a list of businesses registered on the platform.
- **View Business Details**: Users can view a particular business's details, including contact information, description, location, and website (if available).
- **Interact with Businesses**: Users can like and comment on business listings to interact with business owners.
- **Contact Business**: Users can contact businesses via their provided WhatsApp number.

## Non-functional Requirements

- **Performance**: The platform is designed to handle a large number of concurrent users without significant downtime or slowdown.
- **Security**: Token-based authentication is implemented to secure user login, business login, and sensitive data transmission.
- **Usability**: The platform is designed to be user-friendly, with clear navigation and a simple interface.
- **Scalability**: The backend is designed with scalability in mind, making it easy to accommodate future growth and feature additions.

## Tech Stack

- **Backend Language**: Go (Golang)
- **Database**: PostgreSQL (or any SQL database)
- **Authentication**: JSON Web Tokens (JWT) for secure authentication
- **API**: RESTful API for communication between frontend and backend
- **Deployment**: Suitable for cloud deployment on platforms like DigitalOcean, AWS, or Google Cloud

## Getting Started

### Prerequisites

- Go 1.17+ installed on your machine
- PostgreSQL or a compatible SQL database
- Git for version control

### Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/uniconnect-backend.git
   cd uniconnect-backend

2. Install dependencies:
   ```bash
   go mod tidy

3. Set up your environment variables. Create a `.env` file in the root directory and add the following:
   ```bash
   DB_HOST=your-db-host
   DB_USER=your-db-user
   DB_PASSWORD=your-db-password
   DB_NAME=your-db-name
   JWT_SECRET=your-secret-key

4. Run database migrations:
   ```bash
   go run migrations/migrate.go

5. Start the server:
   ```bash
   go run main.go

