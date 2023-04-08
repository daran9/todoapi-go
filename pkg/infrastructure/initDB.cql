CREATE KEYSPACE tododb
  WITH REPLICATION = { 
   'class' : 'NetworkTopologyStrategy',
   'datacenter1' : 3 
  };

CREATE TABLE tododb.tasks_by_id (
    id UUID,
    title text,
    completed boolean,
    created timestamp,
    PRIMARY KEY ((id), created)
    );