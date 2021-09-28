# Basic Chat Api
This is a basic chat api that I made with golang. I use golang/fiber as a framework and JWT for authentication/Authorization.

## Build and Run
you need to create .env file with this content:
```
DB=<DB_NAME>
DB_URI=<DB_URI>
PORT=<PORT_OF_YOUR_SERVER>
```
after that you can run the server with
-   `docker-compose up --build -d`


## Endpoints

### SignUp
-   method: `POST`
-   path: `/api/user/signup`
-   body: 
    ```js
    {
        "username": string,
        "password": string
    }
    ```
-   response:
    ```js
    {
    "data": {
        "token": string
    },
    "error": string,
    "message": string,
    "success": boolean
    }
    ```

### Login
-   method: `POST`
-   path: `/api/user/login`
-   body: 
    ```js
    {
        "username": string,
        "password": string
    }
    ```
-   response:
    ```js
    {
    "data": {
        "token": string
    },
    "error": string,
    "message": string,
    "success": boolean
    }
    ```
### Block User
-   method: `POST`
-   path: `/api/user/block`
   * Header:
  
        |  Name | Description                           | Type   |
        |:---------:|---------------------------------------|--------|
        | Authorization | authentication token of the user  | String |
-   body: 
    ```js
    {
        "username": string,
    }
    ```
-   response:
    ```js
    {
    "data": null,
    "error": string,
    "message": string,
    "success": boolean
    }
    ```

### Send Message
-   method: `POST`
-   path: `/api/message/send`
   * Header:
  
        |  Name | Description                           | Type   |
        |:---------:|---------------------------------------|--------|
        | Authorization | authentication token of the user  | String |
-   body: 
    ```js
    {
        "to": string,
        "content": string
    }
    ```
-   response:
    ```js
    {
    "data": null,
    "error": string,
    "message": string,
    "success": boolean
    }
    ```
### Get Message History
-   method: `GET`
-   path: `/api/message/history`
   * Header:
  
        |  Name | Description                           | Type   |
        |:---------:|---------------------------------------|--------|
        | Authorization | authentication token of the user  | String |
-   response:
    ```js
    {
    "data": {
        "messages": []Message
    },
    "error": string,
    "message": string,
    "success": boolean
    }
    ```
### Get Message History By User
-   method: `GET`
-   path: `/api/message/historyByUser/:username`
   * Header:
  
        |  Name | Description                           | Type   |
        |:---------:|---------------------------------------|--------|
        | Authorization | authentication token of the user  | String |
-   response:
    ```js
    {
    "data": {
        "messages": []Message
    },
    "error": string,
    "message": string,
    "success": boolean
    }
    ```
