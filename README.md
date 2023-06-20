# boardgame-battle
Online multiplayer board game platform. Hosted on Fly.io

# Board Game Platform

Welcome to my Board Game Platform! This project is a web-based platform built using React (ViteJS) and Go, designed to host and facilitate online board games (currently only the deckbuilder Hogwarts Battle)

## Table of Contents

- [Upcoming Dev Goals](#upcoming)
- [Prerequisites](#prerequisites)
- [Usage](#usage)
- [Features](#features)

## Upcoming

- Design Event queue for synchronizing turns and actions into groups. Needed for an 'undo' feature. Currently, all actions just modify the state directly when they hold the gamestate lock. 
- Change in-memory state management to database. Current issue is that Fly.io servers restart periodically and end the game. Hacky fix is to ping clients regularly so Fly won't shutdown, but a database for gamestate is needed eventually. 
- Test concurrency in deployed version
- A complete and playable game - currently it's just playable.
- A better Readme with some pictures. 

### Prerequisites

Make sure you have the following software installed:

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
   go run ./lobbymanager
   ```

   This will start the Go lobbymanager on `http://localhost:8000`.

3. Start the game engine (from root directory):

   ```bash
   go run ./backend
   ```

   This will start the game engine on 'http://localhost:8080'.

4. Open your browser and visit `http://localhost:5173` to access the Board Game Platform.

## Features

The Board Game Platform currently supports the following features:

- Password free user login
- Lobby system for creating/joining game rooms
- Multiplayer support with real-time game synchronization