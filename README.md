# Basic Chat Application in Go

## Overview

This is a basic chat application built using Go. The application leverages WebSockets for real-time communication, the Pat router for HTTP routing, and Jet for template rendering.

## Features

- Real-time messaging using WebSockets (Gorilla WebSocket)
- User online status display
- Automatic removal of users when they leave
- Automatic reconnection for users in case of disconnection

## Prerequisites

Ensure you have the following installed on your machine:

- Go (version 1.16 or later)
- Git

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/sanjay-xdr/GoChat.git
    cd GoChat
    ```

2. Install the necessary Go dependencies:

    ```sh
    go get github.com/gorilla/websocket
    go get github.com/bmizerany/pat
    go get github.com/CloudyKit/jet/v6
    ```

## Running the Application

To run the chat application, use the following command:

```sh
go run cmd/web/*.go
