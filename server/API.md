# Health

Used to ping the application and check if it's alive

**URL** : `/api/ping`

**Method** : `GET`

**Auth required** : NO

**Data constraints**

**Data example**

## Success Response

**Code** : `200 OK`

# Create game

Used to create a game 

**URL** : `/api/game`

**Method** : `POST`

**Auth required** : NO

**Data constraints**

```json
{
    "rows": "[greater than 2 and lower than 100]",
    "columns": "[greater than 2 and lower than 100]",
	"mines": "[greater than 1 and lower than 100]"
	"name": [optional]
}
```

**Data example**

```json
{
	"rows": 10,
	"columns": 10,
	"mines": 30,
	"name": "foo"
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "id": "foo" // Or random word if not name is provided
}
```


# Get Game

Used to get a game 

**URL** : `/api/game/:name`

**Method** : `GET`

**Auth required** : NO

**Data constraints**

**Data example**

## Success Response

**Code** : `200 OK` 

**Content example**
```json
{
    "rows": 3,
    "columns": 3,
    "mines": 2,
    "name": "foo",
    "status": "playing",
    "board": [
        [
            {
                "hasMine": false,
                "status": "",
                "clicked": false,
                "adjancetMines": 0
            },
            {
                "hasMine": false,
                "status": "",
                "clicked": false,
                "adjancetMines": 1
            },
            {
                "hasMine": false,
                "status": "",
                "clicked": false,
                "adjancetMines": 1
            }
        ],
        [
            {
                "hasMine": false,
                "status": "",
                "clicked": false,
                "adjancetMines": 1
            },
            {
                "hasMine": false,
                "status": "",
                "clicked": false,
                "adjancetMines": 2
            },
            {
                "hasMine": true,
                "status": "",
                "clicked": false,
                "adjancetMines": 0
            }
        ],
        [
            {
                "hasMine": true,
                "status": "",
                "clicked": false,
                "adjancetMines": 0
            },
            {
                "hasMine": false,
                "status": "",
                "clicked": false,
                "adjancetMines": 2
            },
            {
                "hasMine": false,
                "status": "",
                "clicked": false,
                "adjancetMines": 1
            }
        ]
    ]
}
```

## Error Response
**Code** : `404 Not Found` 

# Click

Used to do a click

**URL** : `/api/game/:name/click`

**Method** : `POST`

**Auth required** : NO

**Data constraints**

```json
{
    "i": "[valid index on board]",
	"j": "[valid index on board]",
	"action": "flag | question" [optional, no action will be interpreted as a normal click]
}
```

**Data example**

```json
{
	"i": 3,
	"j":3,
	"action": "flag"
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "rows": 3,
    "columns": 3,
    "mines": 2,
    "name": "bar",
    "status": "playing",
    "board": [
        [
            {
                "hasMine": false,
                "status": "",
                "clicked": false,
                "adjancetMines": 1
            },
            {
                "hasMine": false,
                "status": "",
                "clicked": false,
                "adjancetMines": 2
            },
            {
                "hasMine": true,
                "status": "",
                "clicked": false,
                "adjancetMines": 1
            }
        ],
        [
            {
                "hasMine": false,
                "status": "",
                "clicked": false,
                "adjancetMines": 1
            },
            {
                "hasMine": true,
                "status": "",
                "clicked": false,
                "adjancetMines": 1
            },
            {
                "hasMine": false,
                "status": "",
                "clicked": false,
                "adjancetMines": 2
            }
        ],
        [
            {
                "hasMine": false,
                "status": "",
                "clicked": false,
                "adjancetMines": 1
            },
            {
                "hasMine": false,
                "status": "",
                "clicked": false,
                "adjancetMines": 1
            },
            {
                "hasMine": false,
                "status": "flag",
                "clicked": false,
                "adjancetMines": 1
            }
        ]
    ]
}
```


## Error response

**Code**: `403 Forbidden`

**Content example** 

```json
{
    "error": "game is over"
}
```


