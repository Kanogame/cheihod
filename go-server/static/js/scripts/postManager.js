class PostManager {
    SendServerPost = async (urlArgs, body) => {
        const resp = await fetch("http://localhost:10235/" + urlArgs, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
                },
            body: JSON.stringify(body),
        });
        return resp;
    }

    ParseResponceText = async (resp) => {
        const data = await resp.text();
    }
    
    ParseResponceJson = async (resp) => {
        const data = await resp.json();
    }
}

export default PostManager