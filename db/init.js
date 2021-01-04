// connect to admin database to create users
db = new Mongo().getDB("admin");
// create admin user
db.createUser({
  user: "adminCluster",
  pwd: "secret-pass",
  roles: [
    {
      role: "clusterAdmin",
      db: "admin",
    },
  ],
});
// authenticate with admin user
db.auth("adminCluster", "secret-pass");
