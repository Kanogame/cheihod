import WOW from '../node_modules/wow.js/src/WOW.js' 

"use strict"
const wow = new WOW({
  boxClass: 'wow',
  animateClass: 'animated',
  offset: 0,
  live: true
});
wow.init();
console.log("wow loaded");
