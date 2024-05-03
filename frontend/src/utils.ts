const domain: string = "http://localhost:8000";

const handleResponseStatus = (response: Response, errMsg: string) => {
    const { status, ok } = response;

    if (status === 401) {
        localStorage.removeItem("token");
        window.location.reload();
        return;
    }

    if (!ok) {
        throw new Error(errMsg);
    }
};

export const login = (credential: { username: string; password: string }) => {
    const url = `${domain}/signin/`;
    return fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(credential),
    })
        .then((response) => {
            if (!response.ok) {
                throw new Error("Failed to sign in!");
            }

            return response.text();
        })
        .then((token) => {
            localStorage.setItem("authToken", token);
        });
};

export const register = (credential: {
    username: string;
    password: string;
}) => {
    const url = `${domain}/signup/`;
    return fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(credential),
    }).then((response) => {
        handleResponseStatus(response, "Failed to sign up!");
    });
};

export const uploadApp = (
    data: {
        title: string;
        description: string;
        price: number;
    },
    file: File
) => {
    const authToken = localStorage.getItem("authToken");
    const url = `${domain}/upload/`;
    const formData = new FormData();
    formData.append("title", data.title);
    formData.append("description", data.description);
    formData.append("price", data.price.toString());
    formData.append("media_file", file);

    return fetch(url, {
        method: "POST",
        headers: {
            Authorization: `Bearer ${authToken}`,
        },
        body: formData,
    }).then((response) => {
        handleResponseStatus(response, "Failed to upload the app!");

        return response.json();
    });
};

export const checkout = (appId: number) => {
    const authToken = localStorage.getItem("authToken");
    const url = `${domain}/checkout/?appID=${appId}`;
    return fetch(url, {
        method: "POST",
        headers: {
            Authorization: `Bearer ${authToken}`,
            "Content-Type": "application/json",
        },
    })
        .then((response) => {
            handleResponseStatus(response, "Failed to checkout!");

            return response.text();
        })
        .then((redirectUrl) => {
            window.location.href = redirectUrl;
        });
};
