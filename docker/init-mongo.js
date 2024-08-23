db = db.getSiblingDB('Playground');

db.createCollection('user');

db.user.insertMany([
  {
    name: 'John Doe',
    email: 'john@example.com',
    age: 30
  },
  {
    name: 'Jane Smith',
    email: 'jane@example.com',
    age: 28
  }
  // Add more documents as needed
]);

print('Data initialization completed');