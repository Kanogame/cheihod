"use strict"
import postManager from "../scripts/postManager";

const placeform = document.getElementById("new-place-form");

placeform.addEventListener("click", e => {
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
    const resp = post.SendServerPost("place/add", postBody);
    resp = post.ParseResponceText(resp);
    console.log(resp); 
});