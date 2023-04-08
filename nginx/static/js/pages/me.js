import { findTokenFull } from "../utils/token.js";
"use strict"

const name = document.getElementById("me-name");
const email = document.getElementById("me-email");
init();

async function init() {
    const data = await findTokenFull();
    name.textContent = data.name;
    email.textContent = data.email;
}