"use strict"

class PostManager {
    SendPostJson = async (urlArgs, body) => {
        const resp = await fetch("http://localhost:10235/" + urlArgs, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
                },
            body: JSON.stringify(body),
        });
        const data = await resp.json();
        return data.success;
    }

    SendPostText = async (urlArgs, body) => {
        const resp = await fetch("http://localhost:10235/" + urlArgs, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
                },
            body: JSON.stringify(body),
        });
        const data = await resp.text();
        console.log(data);
        return resp;
    }
}

export default PostManager