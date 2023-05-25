# boardgame-battle
Online multiplayer board game platform. Soon to be self-hosted

Sure! Here's a basic README file for your React and Go board game platform project:

# Board Game Platform

Welcome to my Board Game Platform! This project is a web-based platform built using React (Vite) and Go, designed to host and facilitate online board games.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Usage](#usage)
- [Features](#features)
- [Upcoming](#upcoming)

### Prerequisites

Make sure you have the following software installed:

- Node.js (version >= 12.0.0)
- Go (version >= 1.18)

## Usage

1. Start the client development server (from root directory after git cloning):

   ```bash
   cd frontend
   npm run dev
   ```

   This will run the Vite development server on `http://localhost:5173`.

2. Start the lobby manager (from root directory):

   ```bash
   go run github.com/leeswindell/boardgame-battle/lobbymanager
   ```

   This will start the Go lobbymanager on `http://localhost:8000`. When

3. Open your browser and visit `http://localhost:5173` to access the Board Game Platform.

## Features

The Board Game Platform currently supports the following features:

- Password free user login
- Lobby system for creating/joining game rooms
- Multiplayer support with real-time game synchronization

## Upcoming

- A self-hosted version of the site available at www.gamewithyourfriends.dev
- Graceful shutdown of completed or abandoned games.
- A complete and playable game - currently it's just playable
- A better Readme with some pictures. 
