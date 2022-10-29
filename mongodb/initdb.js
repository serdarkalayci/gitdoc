db = db.getSiblingDB('gitdoc');

db.createCollection('users');
db.users.insert({name: 'admin', password: 'admin', role: 'admin'});

