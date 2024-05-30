# MovieCollection

MovieCollection is an application that allows users to catalog and review their favorite movies. The project uses Go for the backend services, TypeScript for the front end, and MongoDB as the database.

## Prerequisites

Before setting up the project, make sure you have the following installed:

- [Go](https://golang.org/doc/install)
- [Node.js](https://nodejs.org/) and npm
- [MongoDB](https://docs.mongodb.com/manual/installation/)
- [Docker](https://docs.docker.com/get-docker/) (optional, for running MongoDB in a container)

## Setup Instructions

### Backend Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/vinyhi/MovieCollection.git
    cd MovieCollection
    ```

2. Set up the backend environment:
    ```bash
    make backend-setup
    ```

3. Run the backend server:
    ```bash
    make backend-run
    ```

### Frontend Setup

1. Navigate to the frontend directory:
    ```bash
    cd src
    ```

2. Set up the frontend environment:
    ```bash
    make frontend-setup
    ```

3. Start the frontend development server:
    ```bash
    make frontend-start
    ```

