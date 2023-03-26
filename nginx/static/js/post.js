"use strict"

class PostConnection {
    constructor(url, data) {
        this.url = url;
        this.data = data;
    }

    async SendData() {
        const resp = await fetch(this.url, {
            method: "POST",
            mode: "cors",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(this.data),
        });
        const text = await resp.text();
        console.log(resp);
        console.log(text);
    }
}

export default PostConnection;