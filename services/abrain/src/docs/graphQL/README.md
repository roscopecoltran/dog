### This document contains graphQL specification used in this project.

Design a structure of sample input and output in graphQL format, keeping in mind our expectations.
Lets say for song sugg.
We want to give input something like this.
```
{
  user: ankur
  suggest: song_suggestion
  return: song.name, artist
}
```
and gives output something like this.
```
{
  results: [
          {
                "song.name": "song1",
                "artist": "abc"
          },
          {
                 "song.name": "song2",
                 "artist": "xyz"
          }
  ]
}
```

I am thinking of graphQL query for this should be like below.
```
{
    User(name: "ankur", action: "song_suggestion") {
         song.name
         artist
    }
}
```
