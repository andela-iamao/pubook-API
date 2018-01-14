### PUBOOK(API)
The Pubook API exposes an interface that enables users add information about books to the platform.

### Why?
Pubook seeks to create a large collection of books by all sorts of authors and even further seeks
to enable prospective authors publish their works on the platform and make it accessible to everyone
with access to the API

### Endpoints

##### GET /api/v1/books
###### Response Body
```
{
   "data": [
       {
           "id": int,
           "title": string,
           "author": string
       }
   ],
   "metadata": {}
}
```

##### POST /api/v1/books
###### Request body
```
{
    "title": string,
    "author": string
}
```

###### Response body
```
{
    "title": string,
    "author": string
}
```

##### GET /api/v1/book/:ID
###### Response Status
    200
###### Response body
```
{
    "book": {
        "id": int
        "title": string,
        "author": string
    }
}
```

##### PUT /api/v1/book/:ID
###### Response Status
    204 success
###### Response body
```
{}
```

##### DELETE /api/v1/book/:ID
###### Response Status
    204 success
###### Response body
```
{}
```



