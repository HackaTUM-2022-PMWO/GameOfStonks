# GameOfStonks

[Live Demo](TODO)

## Inspiration
Have you heard of the person that traded his way from a single red paperclip all the way to a house? Well, it’s a true story and it inspired us to build a fun game about trading digital items among friends (don’t worry, we did not work on NFTs 😀). Besides building something fun, we believe that financial education is incredibly important and set out to code a digital exchange that entertains people from all ages while they learn about market mechanisms through playful interaction. We are calling it **GameOfStonks**.

## What it Does
GameOfStonks is a multiplayer game that runs in your browser, where you and your friends trade the so-called “stonks” with each other and with the machine. You enter the game by logging in on our website and wait in the lobby until enough other players have joined so that the game can begin. Now you have 10min to place “bid” and “ask” orders on the exchange and trade to increase your net worth. In the humble beginnings of your career you have no more than a few paper clips and some coins to your name. But if you have the feel for the market, if you are a true wolf of wall street, these paper clips are not just paper clips to you, they are your chance to climb up the ladder! As soon as the timer hits zero, the game ends and the rich winners are celebrated. Are you up for the challenge?

## Folder Structure

```
.
├── LICENSE
├── README.md
├── backend
│   ├── go.mod
│   ├── go.sum
│   ├── gotsrpc.yaml
│   ├── main.go
│   ├── matcher
│   │   └── matcher.go              # logic to match buy & sell orders
│   ├── services
│   │   └── stonks
│   │       ├── converters.go
│   │       ├── gotsrpc_gen.go
│   │       ├── gotsrpcclient_gen.go
│   │       ├── price.go
│   │       ├── service.go
│   │       ├── stonks.go
│   │       └── utils.go
│   └── store
│       ├── match.go                    
│       ├── memorymatch.go          # the match functionality
│       ├── memorymatch_test.go
│       ├── memoryorder.go          # the order functionality
│       └── order.go
├── frontend
│   ├── README.md
│   ├── package-lock.json
│   ├── package.json
│   ├── postcss.config.js
│   ├── public
│   ├── src
│   │   ├── App.test.tsx
│   │   ├── App.tsx
│   │   ├── assets
│   │   ├── components
│   │   │   ├── Container.tsx
│   │   │   ├── Currency.tsx
│   │   │   ├── buttons
│   │   │   ├── card
│   │   │   ├── footer
│   │   │   ├── graphs
│   │   │   ├── header
│   │   │   ├── inputs
│   │   │   ├── listItems
│   │   │   └── spinner
│   │   ├── icons
│   │   ├── index.css
│   │   ├── index.tsx
│   │   ├── model
│   │   ├── react-app-env.d.ts
│   │   ├── reportWebVitals.ts
│   │   ├── router
│   │   ├── services
│   │   ├── setupTests.ts
│   │   └── views
│   │       ├── Detail.tsx
│   │       ├── Home.tsx
│   │       ├── Lobby.tsx
│   │       ├── Onboard.tsx
│   │       ├── Result.tsx
│   │       ├── Search.tsx
│   │       ├── StartStocks.tsx
│   │       └── Trade.tsx
│   ├── tailwind.config.js
│   └── tsconfig.json
└── scripts
    └── gen.sh
```