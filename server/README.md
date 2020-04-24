# Minesweeper api

You can find the application running on heroku: https://rburdet-minesweeper-api.herokuapp.com/api/game

But you can try it locally if you want
```
#go version go1.14.2 darwin/amd64
go run main.go
```

The API documentation can be found [here](#API.md)

## Deployment

This server was deployed to heroku via git, doing `git push heroku master`, i was trying to do it on `zeit now` but since the v2, they dont support docker deployments anymore :( 

## Assumptions

- Server code has a lot of boilerplate to show the use of DI, and how easy could be to change one implementation to other
- Thanks to DI, service and repositories are easily testable, because of the time constraints i wont implement tests for any of them.
- Random name creation could have a lot of collisions, for massive usage i would use another random generation technique
- For security reasons, number of adjancetMines shouldnt be sent for all the board but for those the click operation opened, for time constraints i couldnt get it done
- In memory data base is certainly not the best approach, im using a dictionary with an instance and not a pointer to keep the references. Im doing this because it let me debug easier, and to make refresh flow similar to a non relational db like mongo.
- Im not using default values on the server and the max and mins are handled by go struct binding reflection 
- Error responses should be handled in a more sophisticated way as they are being hardcoded. Again, this is a time restriction
- More testing are needed on the sum adjacents and click algorithms

# Libaries
- [Http framework](https://github.com/gin-gonic/gin)
- [For cors](https://github.com/gin-contrib/cors)
- [Random generator](https://github.com/Pallinder/go-randomdata)

