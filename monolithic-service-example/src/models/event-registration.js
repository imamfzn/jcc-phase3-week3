const mongoose = require('mongoose');

const EventRegistration = new mongoose.Schema({
  eventId: {
    type: String,
    required: true
  },
  ticketId: {
    type: String,
    required: true
  },
  userId: {
    type: String,
    required: true
  },
},
{
  timestamps: {
    createdAt: 'created_at',
    updatedAt: 'updated_at',
  },
});

EventRegistration.set('toJSON', {
  virtuals: true,
  versionKey: false,
  transform(doc, ret) {
    delete ret._id;
  },
});

module.exports = mongoose.model('EventRegistration', EventRegistration);
