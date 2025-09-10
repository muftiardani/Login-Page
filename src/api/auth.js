const API_BASE_URL = "http://localhost:8080/api";

async function handleResponse(response) {
    const data = await response.json();
    if (!response.ok) {
        throw new Error(data.message || "An error occurred");
    }
    return data;
}

export async function login(credentials) {
    const response = await fetch(`${API_BASE_URL}/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(credentials),
    });
    return handleResponse(response);
}

export async function register(credentials) {
    const response = await fetch(`${API_BASE_URL}/register`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(credentials),
    });
    return handleResponse(response);
}

export async function getStatus(token) {
    const response = await fetch(`${API_BASE_URL}/status`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
        },
    });
    return handleResponse(response);
}

export async function changePassword(token, payload) {
    const response = await fetch(`${API_BASE_URL}/user/password`, {
        method: "PUT",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
        },
        body: JSON.stringify(payload),
    });
    return handleResponse(response);
}