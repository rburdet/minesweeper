# Minesweeper

Hey, this is my implemenation of the challenge deviget proposed. It consists on a backend developed in golang and now deployed in [heroku](https://rburdet-minesweeper-api.herokuapp.com/api) and on a client, based on CRA, deployed on [now](https://minesweeper-pink.now.sh/). 

## Project structure
- [Client](#/client) Has all information related to the client project
- [Server](#/server) Talks about the server, the api documentation is also in its readme

## Disclaimers

I must admit that i cheated on the time, when i had any free time, i thought about the main structures, the most difficult algorithms, and the architecture overall, that time helped me to get the hands on the code much faster and with the idea ready on my head, of course it had a lot of tweaks, but the most important building blocks were already designed. 

I've also used as a base another challenge i had solved for cabify (you can find it [here](https://gitlab.com/rburdet/cabify)) and tried to improve some parts, like being nearer to a CLEAN architecture, or using _gin_ instead of _echo_ 

## Assumptions

You'll find project related assumptions in their folders. Overall, i think the time given to this challenge its not enough and many things should be avoided

- [x] Design and implement a documented RESTful API for the game (think of a mobile app for your API)
- [x] Implement an API client library for the API designed above. Ideally, in a different language, of your preference, to the one used for the API
- [x] When a cell with no adjacent mines is revealed, all adjacent squares will be revealed (and repeat)
- [x] Ability to 'flag' a cell with a question mark or red flag (only on the backend)
- [x] Detect when game is over 
- [x] Persistence (Its a dictionary on the server)
- [ ] Time tracking
- [x] Ability to start a new game and preserve/resume the old ones
- [x] Ability to select the game parameters: number of rows, columns, and mines
- [ ] Ability to support multiple users/accounts

