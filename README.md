# Deepfake detector DB

API endpoints: 

### GET     /api/assertions
Returns all assertions.
Query params: imageId, imageUrl, asserter, assertionId, resolved, seen

### GET     /api/notifications/:asserter
Returns the notifications for an asserter address.

### POST    /api/notifications
Sets a notification to seen.
Body type: JSON
Field: "assertionId"