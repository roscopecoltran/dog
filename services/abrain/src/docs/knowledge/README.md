### This document contains specification for knowledge store and its implementation.

#### knowledge schema transformer
1. A config file will contain schema for knowledge store.
2. Above function will use config file to read schema specification and converted input data to desired schema.
3. Also it will map ouput to desired form

#### knowledge sync
1. This function will be used to sync knowledge store with the current data.
2. This will take data from schema transformer and send it to knowledge store for storing the data.

#### knowledge apis.
1. Based on config file it will transfer generic apis to specialized schema based apis to fetch knowledge from knowledge store.
2. The apis should be made in such a way that it binds with schema present in config file.

#### Idea about how generic schemas can be for knowledge store.
1. We should have fecility for having labels to make some type of collection in the knowledge store.
2. We should have some way of defining generic relationships from one node to other.
3. We can use tags for defining types for nodes and then we can define relationships from that one type of collection to other.
4. We should have some way to define where to do indexing. (We can call it key knowledge).
5. Key knowledge will be data which will be used to fetch knowledge data from graph.
6. We will all users to have weights on relationships so some more complex logic can be implemented easily.
7. One idea is we can store generic schema about knowledge in graph itself. When can serve as meta knowledge.

#### Example schema and generic representation using graph.
#####Lets take example to make music recommendation system. So, for that we need to represent following data into graph.
1. User nodes which will contain which will connect to different songs and location.
2. We will have region nodes which will connect to each users, so that we can give region based suggestions.
3. Each user node will be indexed on username because we will need it to fetch more information.
4. Each user node can have age, sex, ethnity, location etc. as properties.
5. We will have songs nodes which will be connected to users node and its edge will have many times the song has been listened.


***TODO: Example schema in YAML file format.***