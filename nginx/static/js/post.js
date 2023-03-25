"use strict"
import {url} from "../config.json"

async function sendType(type) {
    const resp = await fetch(url, {
        method: "POST",
        body: type
    });
    const text = await resp.text();
    alert(text);
}