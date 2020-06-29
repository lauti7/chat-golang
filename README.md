# simplechat-golang
Simple Chat Application built with GoLang and ReactJS. It is a general chat that receive **join event**, **leave event**, **typing event** and of course **message event**

##How to run ?
1. Open your terminal: `cd $GOPATH`
2. `git clone https://github.com/lauti7/simplechat-golang.git`
3. In the project folder, run: `go get ./...`
4. Then, `cd src` and run: `go run *.go`
5. Open a new terminal tab in project folder and run: `cd chatapp && npm start`

Once you're done, on port `:9000` is running Golang Backend and on port `:3000` is running React.

Golang is serving static files on ":9000/chat" route. You could run: `npm run build` and get access to the chat on `127.0.0.1:9000/chat` (or localhost).   
