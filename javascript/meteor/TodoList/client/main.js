import { Template } from 'meteor/templating';
import { Todos } from '../libs/Todos.js';
import './main.html';

Template.main.helpers({
  todos() {
    return Todos.find({});
  },
})
 