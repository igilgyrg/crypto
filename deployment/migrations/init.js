conn = new Mongo();
db = conn.getDB("crypto");

db.sol.insert({"max": 41.44, "min": 41.05, "open": 41.41, "closed": 41.29, "datetime": ISODate("2022-08-06T00:00:00.000Z")});