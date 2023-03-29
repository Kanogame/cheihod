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
findToken();
const mobilePanel = document.getElementById("mobile-panel");

const bars = document.getElementById("three-bars");
bars.addEventListener("click", () => {
    if (mobilePanel.style.display === "") {
        mobilePanel.style.display = "block";
    } else {
        mobilePanel.style.display = "";
    }
});

document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', function (e) {
        e.preventDefault();

        document.querySelector(this.getAttribute('href')).scrollIntoView({
            behavior: 'smooth'
        });
    });
});