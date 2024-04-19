# boardgame-battle

## What is this?

This is a fullstack app for playing the boardgame Hogwarts Battle. My wonderful
co-founder over at Restart Round and I started playing this game on a trip, and
didn't finish it. So I spent months recreating the entire game's rules in Go,
and built a React frontend to play the game. We finally beat it in time for an
expansion to be released! Technically this is also a platform that could support
any game engine/API with the lobbies.

I hosted the game on a home server using Caddy (an Nginx alternative). I would love to host it for you all to play, but if you want to play, you'll have to buy the game yourself! What is included here is all of the code For the Frontend, Game engine, and lobby system. There's some neat tech related to concurrency, multi-threading, pub/sub, and websockets for a synchronized live-view of the frontend. If you want to run this game yourself, there's instructions below.

### Prerequisites

Make sure you have the following software installed:

- Go (version >= 1.18)

## Usage

1. Clone the repo and start the client development server:

   ```bash
   git clone https://github.com/LeeSwindell/boardgame-battle
   cd frontend
   npm run dev
   ```

   This will run the Vite development server on `http://localhost:5173`.

2. Start the lobby manager (from root directory in a new terminal tab/window):

   ```bash
   go run ./lobbymanager
   ```

   This will start the Go lobbymanager on `http://localhost:8000`.

3. Start the game engine (from root directory in a new terminal tab/window):

   ```bash
   go run ./backend
   ```

   This will start the game engine on 'http://localhost:8080'.

4. Open your browser and visit `http://localhost:5173` to access the Board Game Platform.

5. You will need a copy of the game. Once you have that, scan the cards, and upload them to the /frontend/public/images directory. You can view the necessary names in the /backend/cards.go file under the ImgPath field - e.g. 'ImgPath: "/images/starters/alohomora.jpg"'

## Features

1. The entire base game is available, as well as one expansion.
2. There's a server for connecting and hosting lobbies to join.
3. Everything in the frontend is included except card images.
4. There's an 'undo' feature, which keeps track of the game state at each decision point. Basically any time you take an action you can undo it, all the way back to the start of the game.
5. If you want to host this yourself after importing cards, check out nginx/caddy. If you are going to do this I'm assuming you know what you're doing.

### Challenges

Some of the biggest challenges were maintaining a consistent game state. If multiple players can all queue actions asynchronously, how do we decide what actually happens? Typically we would have an async/await style backend, but here I chose to use multi threading with a shared memory model. This way I could use a mutex to lock the shared state, and perform many checks to ensure the 'right' action was being processed each time. In addition, there are a number of side-effects that happen in this game. One single card might trigger 12 different events. It's kind of obnoxious in person to play this game because of that, but coding it was even worse (not really though... I loved it). I used go routines to 'queue' side effects to happen, and maintained a loop that would check for certain coditions. This way, if the loop caught something, it could grab the lock and squeeze into the side effect queue. This also required creating a pub/sub or event handling system.
