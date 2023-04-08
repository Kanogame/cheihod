"use strict"
import postManager from "../scripts/postManager.js";

const placeform = document.getElementById("new-place-form");

placeform.addEventListener("submit", async e => {
    e.preventDefault();
    const data = new FormData(e.target);
    const postBody = {
        name: data.get("name"),
        place: data.get("place"),
        date: data.get("date"),
        time: data.get("time"),
        capacity: data.get("capacity")
    }
    const post = new postManager();
    const resp = await post.SendPostJson("place/add", postBody);
    console.log(resp);
    if (resp === true) {
        alert("success");
    } else {
        alert("error while handling post");
    }
});