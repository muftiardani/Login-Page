const API_BASE_URL = "http://localhost:8080/api";

async function handleResponse(response) {
    const data = await response.json();
    if (!response.ok) {
        throw new Error(data.message || "An error occurred");
    }
    return data;
}

async function fetchWithAuth(url, options = {}) {
    const fetchOptions = { ...options, credentials: 'include' };
    let response = await fetch(url, fetchOptions);
    if (response.status === 401) {
        await fetch(`${API_BASE_URL}/refresh`, { method: 'POST', credentials: 'include' });
        response = await fetch(url, fetchOptions);
    }
    return response;
}

export async function login(credentials) {
    const response = await fetch(`${API_BASE_URL}/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(credentials),
        credentials: 'include',
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

export async function changePassword(payload) {
    const response = await fetchWithAuth(`${API_BASE_URL}/user/password`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
    });
    return handleResponse(response);
}

export async function logout() {
    await fetch(`${API_BASE_URL}/logout`, {
        method: "POST",
        credentials: 'include',
    });
}

export async function getPayments() {
    const response = await fetchWithAuth(`${API_BASE_URL}/payments`);
    return handleResponse(response);
}

export async function getDashboardSummary() {
    const response = await fetchWithAuth(`${API_BASE_URL}/dashboard/summary`);
    return handleResponse(response);
}

export async function getChartData() {
    const response = await fetchWithAuth(`${API_BASE_URL}/dashboard/chart`);
    return handleResponse(response);
}