# GameOfStonks

[Live Demo](TODO)

## Inspiration
Have you heard of the person that traded his way from a single red paperclip all the way to a house? Well, itβs a true story and it inspired us to build a fun game about trading digital items among friends (donβt worry, we did not work on NFTs π). Besides building something fun, we believe that financial education is incredibly important and set out to code a digital exchange that entertains people from all ages while they learn about market mechanisms through playful interaction. We are calling it **GameOfStonks**.

## What it Does
GameOfStonks is a multiplayer game that runs in your browser, where you and your friends trade the so-called βstonksβ with each other and with the machine. You enter the game by logging in on our website and wait in the lobby until enough other players have joined so that the game can begin. Now you have 10min to place βbidβ and βaskβ orders on the exchange and trade to increase your net worth. In the humble beginnings of your career you have no more than a few paper clips and some coins to your name. But if you have the feel for the market, if you are a true wolf of wall street, these paper clips are not just paper clips to you, they are your chance to climb up the ladder! As soon as the timer hits zero, the game ends and the rich winners are celebrated. Are you up for the challenge?

## Folder Structure

```
.
βββ LICENSE
βββ README.md
βββ backend
βΒ Β  βββ go.mod
βΒ Β  βββ go.sum
βΒ Β  βββ gotsrpc.yaml
βΒ Β  βββ main.go
βΒ Β  βββ matcher
βΒ Β  βΒ Β  βββ matcher.go              # logic to match buy & sell orders
βΒ Β  βββ services
βΒ Β  βΒ Β  βββ stonks
βΒ Β  βΒ Β      βββ converters.go
βΒ Β  βΒ Β      βββ gotsrpc_gen.go
βΒ Β  βΒ Β      βββ gotsrpcclient_gen.go
βΒ Β  βΒ Β      βββ price.go
βΒ Β  βΒ Β      βββ service.go
βΒ Β  βΒ Β      βββ stonks.go
βΒ Β  βΒ Β      βββ utils.go
βΒ Β  βββ store
βΒ Β      βββ match.go                    
βΒ Β      βββ memorymatch.go          # the match functionality
βΒ Β      βββ memorymatch_test.go
βΒ Β      βββ memoryorder.go          # the order functionality
βΒ Β      βββ order.go
βββ frontend
βΒ Β  βββ README.md
βΒ Β  βββ package-lock.json
βΒ Β  βββ package.json
βΒ Β  βββ postcss.config.js
βΒ Β  βββ public
βΒ Β  βββ src
βΒ Β  βΒ Β  βββ App.test.tsx
βΒ Β  βΒ Β  βββ App.tsx
βΒ Β  βΒ Β  βββ assets
βΒ Β  βΒ Β  βββ components
βΒ Β  βΒ Β  βΒ Β  βββ Container.tsx
βΒ Β  βΒ Β  βΒ Β  βββ Currency.tsx
βΒ Β  βΒ Β  βΒ Β  βββ buttons
βΒ Β  βΒ Β  βΒ Β  βββ card
βΒ Β  βΒ Β  βΒ Β  βββ footer
βΒ Β  βΒ Β  βΒ Β  βββ graphs
βΒ Β  βΒ Β  βΒ Β  βββ header
βΒ Β  βΒ Β  βΒ Β  βββ inputs
βΒ Β  βΒ Β  βΒ Β  βββ listItems
βΒ Β  βΒ Β  βΒ Β  βββ spinner
βΒ Β  βΒ Β  βββ icons
βΒ Β  βΒ Β  βββ index.css
βΒ Β  βΒ Β  βββ index.tsx
βΒ Β  βΒ Β  βββ model
βΒ Β  βΒ Β  βββ react-app-env.d.ts
βΒ Β  βΒ Β  βββ reportWebVitals.ts
βΒ Β  βΒ Β  βββ router
βΒ Β  βΒ Β  βββ services
βΒ Β  βΒ Β  βββ setupTests.ts
βΒ Β  βΒ Β  βββ views
βΒ Β  βΒ Β      βββ Detail.tsx
βΒ Β  βΒ Β      βββ Home.tsx
βΒ Β  βΒ Β      βββ Lobby.tsx
βΒ Β  βΒ Β      βββ Onboard.tsx
βΒ Β  βΒ Β      βββ Result.tsx
βΒ Β  βΒ Β      βββ Search.tsx
βΒ Β  βΒ Β      βββ StartStocks.tsx
βΒ Β  βΒ Β      βββ Trade.tsx
βΒ Β  βββ tailwind.config.js
βΒ Β  βββ tsconfig.json
βββ scripts
    βββ gen.sh
```