# Disclaimers

I must admit that i cheat in the time, in my free times, like when i was having lunch or having a bath i thought of the main structures, the most difficult algorithms, and the architecture overall, that time helped me to get the hands on the code much faster and with the idea ready on my head, of course it had a lot of tweaks, but the most important building blocks were already designed. 

I've also used as a base another challenge i had solved for cabify (you can find it [here](https://gitlab.com/rburdet/cabify)) and tried to improve some parts, like being nearer to a CLEAN architecture, or using _gin_ instead of _echo_ 
#Assumptions

- Code has a lot of boilerplate to show the use of DI, and how easy could be to change one implementation to other
- Thanks to DI, service and repositories are easily testable, because of the time constraints i wont implement tests for any of them.
- Random name creation could have a lot of collisions, for massive usage i would use another random generation technique
- For security reasons, number of adjancetMines shouldnt be sent for all the board but for those the click operation opened, for time constraints i couldnt get it done
- In memory data base is certainly not the best approach, im using a dictionary with an instance and not a pointer to keep the references. Im doing this because it let me debug easier, and to make refresh flow similar to a non relational db like mongo.
- Im not using default values on the server and the max and mins are handled by go struct binding reflection 
- Error responses should be handled in a more sophisticated way as they are being hardcoded. Again, this is a time restriction
- More testing are needed on the sum adjacents and click algorithms
- Client is not handling error codes efficiently

# Libaries

- [Http framework](https://github.com/gin-gonic/gin)
- [For cors](https://github.com/gin-contrib/cors)
- [Random generator](https://github.com/Pallinder/go-randomdata)
