import { findTokenFull } from "../utils/token.js"
import WOW from '../../node_modules/wow.js/src/WOW.js';

"use strict"
const wow = new WOW({
  boxClass: 'wow',
  animateClass: 'animated',
  offset: 0,
  live: true
});
wow.init();

const mobilePanel = document.getElementById("mobile-panel");
const account = document.getElementById("account");
const login = document.getElementById("login");
const accountM = document.getElementById("account-mobile");
const loginM = document.getElementById("login-mobile");
const name = document.getElementById("name");

const bars = document.getElementById("three-bars");
bars.addEventListener("click", () => {
    if (mobilePanel.style.display === "") {
        mobilePanel.style.display = "block";
    } else {
        mobilePanel.style.display = "";
    }
});

const resp = await findTokenFull();
try {
if (resp.success === "true") {
    account.classList.remove("disabled");
    login.classList.add("disabled");
    accountM.classList.remove("disabled");
    loginM.classList.add("disabled");
    name.textContent = resp.name
}} catch (e) { accountM.classList.add("disabled"); }

document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', function (e) {
        e.preventDefault();
        
        document.querySelector(this.getAttribute('href')).scrollIntoView({
            behavior: 'smooth'
        });
    });
});
