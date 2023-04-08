import findToken from "../utils/token.js"

const mobilePanel = document.getElementById("mobile-panel");
const account = document.getElementById("account");
const login = document.getElementById("login");
const name = document.getElementById("name");

const bars = document.getElementById("three-bars");
bars.addEventListener("click", () => {
    if (mobilePanel.style.display === "") {
        mobilePanel.style.display = "block";
    } else {
        mobilePanel.style.display = "";
    }
});

findToken(account, login, name);