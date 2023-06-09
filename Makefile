build:
	cd backend && go build -ldflags="-X main.appEnv=prod"
	cd lobbymanager && go build -ldflags="-X main.appEnv=prod"
	cd frontend && npm run build
