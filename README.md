# ABrain
Developing a generic framework for combining graphdb, full text search and machine learning.

The Idea is to develop a framework from where its easy to integrate all above modules, something like brain architecture
design can me made easy.

## Adding some points to ponder upon while developing initial prototype

### Functionalities that I am thinking of at this point. (Will add some more as project proceeds further)
Basically some actions in brain are voluntary and some are involuntary. Perticulary involuntary actions are because of some learning behaiviours like mucle memory or like some genetically coming information like involuntary movement of esophagus.

So, I think that all these involuntary behaivours are learning models in terms of machine learning due to environment or genetic code or anything else. Like someone might learn to fight and then certain things becomes involentary which he is fighting like quick reactions and everything.

Then there are some voluntary actions like we trying to read something or when we use our hands to eat or when we try to remember anyone these are all voluntary.

I might be wrong at things point of time in certain concept. But its better to start off with something for now and then later modify it according to better understanding of the problem.

### Technology to be used in this project.
I am thinking of using Go programming language because its fast and good for systems programming.
I will probably use be using python for machine learning models.
For graph database I have decided to use dgraph because of my experience with the system.
and for full text search if required I will use elasticsearch again because of my experience with the system.

### I/O to the system.
I was thinking that input and ouput to this system should be very intuitive and should feel like if you provide some input to the system it should use background knowledge and some learned behaiviour about the subject and then generate desired output.
Here I will use term stimulus for the system input and response for system output because currently it feels more intuitive.
```
|Stimulus given   |-->|Acquiring background detail|-->|Generating desired behaviour    |-->|Mapping ML output    |-->|Reponse|
|by user or sensor|   |about subject from graphdb |   |based on backgourd details      |   |to subject knowledge |
|                 |   |and ES                     |   |and subject from learned models |   |from graphdb and ES  |
```

### Input output format.
Currently I am thinking to use json for input and output format as generally In our system one graph design is complete and finalized and after its serializer function are written we will rarely modify the structure again.
May be in future we can add some functionality in opencypher and use that in out system.

### Example case with desired input and output.
1) Lets say we have made a music recomendation system.
So, we have already made desired graph structure and involuntary behaiviour of the system is to learn from all the incoming user data (what song a user is litening) and based on user age group, region, sex etc. it learns which songs should be suggested to same type of user (kinda supervised learning). 
So, a new user if we ask like
```
{
  user: ankur
  suggest: song_suggestion
  return: song.name, artist
}
```
So, here we have triggered a voluntary action to find all required data for ankur from graphdb and use that data to find song recommendations and find its relevant info from graphdb and send response.

2) Another example can be identifying what is happening in a picture.
So, we have already made desired graph structure and involuntary behaiviour of the system is to detect objects and actions from the image source and based on that it learns whats happening in the picture. Later when we pass new picture based on previous knowledge now system can tell what is happening in the system.
```
{
  input: image
  suggest: whats_happening
  return: main_character.action
}
```

***I will add some more examples so as to clearify use case for this system***

***Note: This system will not be a db. Its a framework that will help developers create brain architecture easily and intuitively***
