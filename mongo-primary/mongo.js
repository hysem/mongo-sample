// rs.initiate();
// rs.add({ host: "mongo-secondary:27017", priority: 0, votes: 0 });

config = {
  _id: "testReplica",
  members: [
    { _id: 0, host: "mongo-primary:27017" },
    { _id: 1, host: "mongo-secondary:27017" },
  ],
};
rs.initiate(config);
