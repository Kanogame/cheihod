"use strict"

async function sendPost(url, object) {
    const resp = await fetch(url, {
        method: "POST",
        body: object
    });
    const text = await resp.text();
    alert(text);
}