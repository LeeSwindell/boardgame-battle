build:
	cd backend && go build -ldflags="-X main.appEnv=prod"
	cd lobbymanager && go build -ldflags="-X main.appEnv=prod"
	cd frontend && npm run build

testbuild: 
	cd backend && go build
	cd lobbymanager && go build
	cd frontend && npm run testbuild
	
run:
	osascript -e 'tell app "Terminal" to do script "cd coding/hogwarts && go run ./backend -ldflags=\"-X main.appEnv=prod\""'
	osascript -e 'tell app "Terminal" to do script "cd coding/hogwarts && go run ./lobbymanager -ldflags=\"-X main.appEnv=prod\""'
	cd frontend && npm run build && cd ..
	caddy run

startgame:
	cd backend && go build -ldflags="-X main.appEnv=prod" && ./backend

startlm: 
	cd lobbymanager && go build -ldflags="-X main.appEnv=prod" && ./lobbymanager
