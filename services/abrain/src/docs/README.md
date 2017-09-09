### This document contains folder architechture for this project.

#### knowledge (This folder will contain drivers and api for talking to graph database)
1. This directory will potentially contain knowledge graph apis and may to ES apis for creating and fetching knowledge to and from disk. (At present we will only implement graph apis and later think about ES apis)
2. It should have a config file that will contain knowledge representation config for graph and elasticsearch.
3. Config file should act as some kind of schema for graphdb.
4. Config file should be a simple YML file easily configurable by anyone so that its easy for anyone to design a knowledge store.

#### actions (This folder will contain ML scripts that should be executed on either stimulus data or knowledge data)
1. This directory will contain all the action scripts which can be either voluntary scripts, involuntary scripts or both.
2. Voluntary scripts are those scipts which will need external triggers to execute them.
3. Involuntary scripts are those which are executed on internal triggers. i.e. (by other actions or time based crons)
4. this folder should also have a config file which should be used to define these voluntary and involuntary actions.
5. each action will have a specific format for input (probably a json format). that would be passed by system as a command line argument.

#### triggers (here there should be code for different triggers to the system)
1. This will contain code for cron based triggers.
2. It place will also contain custom triggers which can be specified by user.
3. Again we will have a specific config file which will specify when will that trigger be fired and some related data.


#### orchestrator (cmd directory) (This will be the main entry and exit point to system)
1. Based on various stimulus it will invoke proper triggers.
2. It will manage proper data flow between modules.
3. It will keep everything in synergy with each other.
4. After all the operations it will provide proper response in the form it is desired by the user.
