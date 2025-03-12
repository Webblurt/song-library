# Song Library API

song library is a Go-based RESTful API for managing a song library. It allows you to store, retrieve, create, update, and delete song information, including lyrics, group details, and release dates.

## Features

*   **Retrieve Song Lyrics:** Fetch song lyrics based on various filters (song ID, group name, song name, pagination).
*   **Browse Song Library:** Get a list of songs in the library with filtering capabilities (group, song, release date).
*   **Create New Songs:** Add new songs to the library with details like group name and song name.
*   **Update Existing Songs:** Modify details of existing songs, including group, song name, release date, and external links.
*   **Delete Songs:** Remove songs from the library.
*   **External API Integration:** Integrates with an external API to fetch initial song information.
*   **Database Persistence:** Uses PostgreSQL as the backend database for storing song data.
*   **RESTful API:** Follows REST principles for a clear and consistent API design.
*   **Logging:** Includes basic logs to debug application.

## Technologies Used

*   **Go:** Programming language.
*   **net/http:** Go's standard library for HTTP.
*   **encoding/json:** Go's standard library for JSON encoding/decoding.
*   **pgx:** PostgreSQL driver for Go.
*   **PostgreSQL:** Database.
*   **golang-migrate:** Database migration tool.

## API Endpoints

### Song Lyrics

*   **GET /api/v1/library/songs?song_id={song_id}**
    *   Retrieves the lyrics for a specific song by ID.
    *   Example: `/api/v1/library/songs?song_id=123`
*   **GET /api/v1/library/songs?group={group}&song={song}&page={page}&limit={limit}**
    *   Retrieves song lyrics based on group and song name, with pagination.
    *   `group`: Name of the music group (e.g., "The Beatles").
    *   `song`: Name of the song (e.g., "Hey Jude").
    *   `page`: Page number to retrieve.
    *   `limit`: Limit of results on page.
    *   Example: `/api/v1/library/songs?group=The%20Beatles&song=Hey%20Jude&page=1&limit=10`
*   **Error Handling:**
    *   Returns `400 Bad Request` if required parameters (song_id or both group and song) are missing.
    *   Returns `404 Not Found` if the song is not found.
    *   Returns `500 Internal Server Error` if there is an issue encoding the response.

### Song Management

*   **POST /api/v1/library/songs**
    *   Creates a new song.
    *   **Request Body (JSON):**
        ```json
        {
          "group": "New Group",
          "song": "New Song"
        }
        ```
    *   **Response:**
        *   `201 Created` on success, with a message indicating successful creation.
        *   `500 Internal Server Error` if there is an error while song creation.
        *   `400 Bad Request` if request body is not valid.
*   **PUT /api/v1/library/songs**
    *   Updates an existing song.
    *   **Request Body (JSON):**
        ```json
        {
          "id": "123",
          "group": "Updated Group",
          "song": "Updated Song",
          "release_date": "2023-10-27",
          "link": "https://example.com/updated-link"
        }
        ```
        *   `release_date` must be in `YYYY-MM-DD` format.
    *   **Response:**
        *   `200 OK` on success, with a message indicating successful update.
        *   `500 Internal Server Error` if there is an error while song updating.
        *   `400 Bad Request` if request body is not valid.

*   **DELETE /api/v1/library/songs?song_id={song_id}**
    *   Deletes an existing song.
     *   `song_id`: ID of song to delete.
    *   Example: `/api/v1/library/songs?song_id=123`
    *   **Response:**
        *   `200 OK` on success, with a message indicating successful deletion.
        *   `500 Internal Server Error` if there is an error while song deleting.
        *   `400 Bad Request` if `song_id` is not provided.

### Library Data

*   **GET /api/v1/library?group={group}&song={song}&release_date={release_date}&limit={limit}&offset={offset}**
    *   Retrieves a list of songs from the library based on various filters.
    *   `group`: Filter by group name.
    *   `song`: Filter by song name.
    *   `release_date`: Filter by release date.
    *   `limit`: Limit of results on page.
    *   `offset`: Offset of the results.
    *   Example: `/api/v1/library?group=The%20Beatles&limit=20&offset=10`
*   **Error Handling:**
    *   Returns `404 Not Found` if library data is not found.
    *   Returns `500 Internal Server Error` if there is an issue encoding the response.

## Getting Started

### Prerequisites

*   **Go:** Make sure you have Go installed on your system.
*   **PostgreSQL:** You need a running PostgreSQL database instance.
* **External API**: There is an External API client, you should implement it.

### Configuration

The application is configured using environment variables, stored in a `.env` file in the root directory of the project.

1.  **Create `.env`:** Create a new file named `.env` in the project's root directory.
2.  **Populate `.env`:** Add the following environment variables to your `.env` file, adjusting the values as needed:

    ```properties
    DB_NAME=mydb
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=user
    DB_PASSWORD=secret
    DB_SSLMODE=disable
    MIGRATIONS_PATH=C:/song-library # You should change it to correct path.
    SERVER_PORT=8080
    API_NAME=example_api
    API_URL=https://api.example.com
    ENABLE_DEBUG=true
    ```

    *   **`DB_NAME`:** The name of your PostgreSQL database.
    *   **`DB_HOST`:** The hostname or IP address of your PostgreSQL server.
    *   **`DB_PORT`:** The port your PostgreSQL server is listening on (default: 5432).
    *   **`DB_USER`:** The username for connecting to your PostgreSQL database.
    *   **`DB_PASSWORD`:** The password for the database user.
    *   **`DB_SSLMODE`:** The SSL mode for the database connection (e.g., `disable`, `require`).
    * **`MIGRATIONS_PATH`**: Path to project root directory.
    *   **`SERVER_PORT`:** The port the API server will run on (default: 8080).
    *   **`API_NAME`:** Name of your API, can be any.
    *   **`API_URL`:** URL to you API, can be any.
    *   **`ENABLE_DEBUG`:** Enable debug logs if `true`.

### Installation

1.  **Clone the repository:**
    ```bash
    git clone <repository-url>
    cd sing-lib
    ```

2.  **Install dependencies:**
    ```bash
    go mod download
    ```

3.  **Run the application:**
    ```bash
    go run cmd/song-library/main.go
    ```


